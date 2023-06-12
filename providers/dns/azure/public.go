package azure

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/sonavilabs/lego/v4/challenge/dns01"
	"github.com/sonavilabs/lego/v4/platform/config/env"
)

// DNSProviderPublic implements the challenge.Provider interface for Azure Public Zone DNS.
type DNSProviderPublic struct {
	config       *Config
	credentials  *azcore.TokenCredential
	zoneClient   *armdns.ZonesClient
	recordClient *armdns.RecordSetsClient
}

// NewDNSProviderPublic creates a DNSProviderPublic structure with intialised Azure clients.
func NewDNSProviderPublic(config *Config, credentials *azcore.TokenCredential) (*DNSProviderPublic, error) {
	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: config.Environment,
		},
	}

	zoneClient, err := armdns.NewZonesClient(config.SubscriptionID, *credentials, &options)
	if err != nil {
		return nil, fmt.Errorf("azure: failed to init zones client %w", err)
	}

	recordClient, err := armdns.NewRecordSetsClient(config.SubscriptionID, *credentials, &options)
	if err != nil {
		return nil, fmt.Errorf("azure: failed to init record set client %w", err)
	}

	dnsProvider := &DNSProviderPublic{
		config:       config,
		credentials:  credentials,
		zoneClient:   zoneClient,
		recordClient: recordClient,
	}

	return dnsProvider, nil
}

// Timeout returns the timeout and interval to use when checking for DNS propagation.
// Adjusting here to cope with spikes in propagation times.
func (d *DNSProviderPublic) Timeout() (timeout, interval time.Duration) {
	return d.config.PropagationTimeout, d.config.PollingInterval
}

// Present creates a TXT record to fulfill the dns-01 challenge.
func (d *DNSProviderPublic) Present(domain, token, keyAuth string) error {
	ctx := context.Background()
	fqdn, value := dns01.GetRecord(domain, keyAuth)

	zone, err := d.getHostedZoneID(ctx, fqdn)
	if err != nil {
		return fmt.Errorf("azure: failed to get hosted zone, %w", err)
	}

	subDomain, err := dns01.ExtractSubDomain(fqdn, zone)
	if err != nil {
		return fmt.Errorf("azure: failed to extract subdomain, %w", err)
	}

	// Get existing record set
	rset, err := d.recordClient.Get(ctx, d.config.ResourceGroup, zone, subDomain, armdns.RecordTypeTXT, nil)
	if err != nil {
		var respErr *azcore.ResponseError
		if !errors.As(err, &respErr) || respErr.StatusCode != http.StatusNotFound {
			return fmt.Errorf("azure: failed to get existing record set, %w", err)
		}
	}

	// Construct unique TXT records using map
	uniqRecords := map[string]struct{}{value: {}}
	if rset.RecordSet.Properties != nil && rset.RecordSet.Properties.TxtRecords != nil {
		for _, txtRecord := range rset.RecordSet.Properties.TxtRecords {
			// Assume Value doesn't contain multiple strings
			if len(txtRecord.Value) > 0 {
				uniqRecords[*txtRecord.Value[0]] = struct{}{}
			}
		}
	}

	var txtRecords []*armdns.TxtRecord
	for txt := range uniqRecords {
		txtRecord := txt
		txtRecords = append(txtRecords, &armdns.TxtRecord{Value: []*string{&txtRecord}})
	}

	ttlInt64 := int64(d.config.TTL)
	rec := armdns.RecordSet{
		Name: &subDomain,
		Properties: &armdns.RecordSetProperties{
			TTL:        &ttlInt64,
			TxtRecords: txtRecords,
		},
	}

	_, err = d.recordClient.CreateOrUpdate(ctx, d.config.ResourceGroup, zone, subDomain, armdns.RecordTypeTXT, rec, nil)
	if err != nil {
		return fmt.Errorf("azure: failed to create/update record, %w", err)
	}
	return nil
}

// CleanUp removes the TXT record matching the specified parameters.
func (d *DNSProviderPublic) CleanUp(domain, token, keyAuth string) error {
	ctx := context.Background()
	fqdn, _ := dns01.GetRecord(domain, keyAuth)

	zone, err := d.getHostedZoneID(ctx, fqdn)
	if err != nil {
		return fmt.Errorf("azure: failed to get hosted zone, %w", err)
	}

	subDomain, err := dns01.ExtractSubDomain(fqdn, zone)
	if err != nil {
		return fmt.Errorf("azure: failed to extract subdomain, %w", err)
	}

	_, err = d.recordClient.Delete(ctx, d.config.ResourceGroup, zone, subDomain, armdns.RecordTypeTXT, nil)
	if err != nil {
		return fmt.Errorf("azure: failed to delete record, %w", err)
	}
	return nil
}

// Checks that azure has a zone for this domain name.
func (d *DNSProviderPublic) getHostedZoneID(ctx context.Context, fqdn string) (string, error) {
	if zone := env.GetOrFile(EnvZoneName); zone != "" {
		return zone, nil
	}

	authZone, err := dns01.FindZoneByFqdn(fqdn)
	if err != nil {
		return "", fmt.Errorf("failed to get zone via fqdn, %w", err)
	}

	zone, err := d.zoneClient.Get(ctx, d.config.ResourceGroup, dns01.UnFqdn(authZone), nil)
	if err != nil {
		return "", fmt.Errorf("failed to get zone, %w", err)
	}

	// zone.Name shouldn't have a trailing dot(.)
	return strings.TrimSuffix(*zone.Name, "."), nil
}
