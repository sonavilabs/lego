// Package azure implements a DNS provider for solving the DNS-01 challenge using azure DNS.
// Azure doesn't like trailing dots on domain names, most of the acme code does.
package azure

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/sonavilabs/lego/v4/challenge"
	"github.com/sonavilabs/lego/v4/platform/config/env"
)

const defaultMetadataEndpoint = "http://169.254.169.254"

// Environment variables names.
const (
	envNamespace = "AZURE_"

	EnvEnvironment      = envNamespace + "ENVIRONMENT"
	EnvMetadataEndpoint = envNamespace + "METADATA_ENDPOINT"
	EnvSubscriptionID   = envNamespace + "SUBSCRIPTION_ID"
	EnvResourceGroup    = envNamespace + "RESOURCE_GROUP"
	EnvTenantID         = envNamespace + "TENANT_ID"
	EnvClientID         = envNamespace + "CLIENT_ID"
	EnvClientSecret     = envNamespace + "CLIENT_SECRET"
	EnvZoneName         = envNamespace + "ZONE_NAME"
	EnvPrivateZone      = envNamespace + "PRIVATE_ZONE"

	EnvTTL                = envNamespace + "TTL"
	EnvPropagationTimeout = envNamespace + "PROPAGATION_TIMEOUT"
	EnvPollingInterval    = envNamespace + "POLLING_INTERVAL"
)

// Config is used to configure the creation of the DNSProvider.
type Config struct {
	ClientID     string
	ClientSecret string
	TenantID     string

	SubscriptionID string
	ResourceGroup  string
	PrivateZone    bool

	MetadataEndpoint string
	Environment      cloud.Configuration

	PropagationTimeout time.Duration
	PollingInterval    time.Duration
	TTL                int
}

// NewDefaultConfig returns a default configuration for the DNSProvider.
func NewDefaultConfig() *Config {
	return &Config{
		TTL:                env.GetOrDefaultInt(EnvTTL, 60),
		PropagationTimeout: env.GetOrDefaultSecond(EnvPropagationTimeout, 2*time.Minute),
		PollingInterval:    env.GetOrDefaultSecond(EnvPollingInterval, 2*time.Second),
		MetadataEndpoint:   env.GetOrFile(EnvMetadataEndpoint),
		Environment:        cloud.AzurePublic,
	}
}

// DNSProvider implements the challenge.Provider interface.
type DNSProvider struct {
	provider challenge.ProviderTimeout
}

// NewDNSProvider returns a DNSProvider instance configured for azure.
// Credentials can be passed in the environment variables:
// AZURE_ENVIRONMENT, AZURE_CLIENT_ID, AZURE_CLIENT_SECRET,
// AZURE_SUBSCRIPTION_ID, AZURE_TENANT_ID, AZURE_RESOURCE_GROUP
// If the credentials are _not_ set via the environment,
// then it will attempt to get a bearer token via the instance metadata service.
// see: https://github.com/Azure/go-autorest/blob/v10.14.0/autorest/azure/auth/auth.go#L38-L42
func NewDNSProvider() (*DNSProvider, error) {
	config := NewDefaultConfig()

	environmentName := env.GetOrFile(EnvEnvironment)
	if environmentName != "" {
		switch environmentName {
		case "china":
			config.Environment = cloud.AzureChina
		case "public":
			config.Environment = cloud.AzurePublic
		case "usgovernment":
			config.Environment = cloud.AzureGovernment
		default:
			return nil, fmt.Errorf("azuredns: unknown environment %s", environmentName)
		}
	}

	config.SubscriptionID = env.GetOrFile(EnvSubscriptionID)
	config.ResourceGroup = env.GetOrFile(EnvResourceGroup)
	config.ClientSecret = env.GetOrFile(EnvClientSecret)
	config.ClientID = env.GetOrFile(EnvClientID)
	config.TenantID = env.GetOrFile(EnvTenantID)
	config.PrivateZone = env.GetOrDefaultBool(EnvPrivateZone, false)

	return NewDNSProviderConfig(config)
}

// NewDNSProviderConfig return a DNSProvider instance configured for Azure.
func NewDNSProviderConfig(config *Config) (*DNSProvider, error) {
	if config == nil {
		return nil, errors.New("azure: the configuration of the DNS provider is nil")
	}

	credentials, err := getTokenCredentials(config)
	if err != nil {
		return nil, fmt.Errorf("azure: failed to get token credentails %w", err)
	}

	if config.SubscriptionID == "" {
		subsID, err := getMetadata(config, "subscriptionId")
		if err != nil {
			return nil, fmt.Errorf("azure: failed to get subscription id from metadata: %w", err)
		}

		if subsID == "" {
			return nil, errors.New("azure: SubscriptionID is missing")
		}

		config.SubscriptionID = subsID
	}

	if config.ResourceGroup == "" {
		resGroup, err := getMetadata(config, "resourceGroupName")
		if err != nil {
			return nil, fmt.Errorf("azure: failed to get resource group from metadata:: %w", err)
		}

		if resGroup == "" {
			return nil, errors.New("azure: resourceGroup is missing")
		}
		config.ResourceGroup = resGroup
	}

	if config.PrivateZone {
		dnsProvider, err := NewDNSProviderPrivate(config, &credentials)
		if err != nil {
			return nil, fmt.Errorf("azure: failed to create private dns provider, %w", err)
		}
		return &DNSProvider{provider: dnsProvider}, nil
	}

	dnsProvider, err := NewDNSProviderPublic(config, &credentials)
	if err != nil {
		return nil, fmt.Errorf("azure: failed to create public dns provider, %w", err)
	}

	return &DNSProvider{provider: dnsProvider}, nil
}

// Timeout returns the timeout and interval to use when checking for DNS propagation.
// Adjusting here to cope with spikes in propagation times.
func (d *DNSProvider) Timeout() (timeout, interval time.Duration) {
	return d.provider.Timeout()
}

// Present creates a TXT record to fulfill the dns-01 challenge.
func (d *DNSProvider) Present(domain, token, keyAuth string) error {
	err := d.provider.Present(domain, token, keyAuth)
	if err != nil {
		return fmt.Errorf("failed to present txt record, %w", err)
	}

	return nil
}

// CleanUp removes the TXT record matching the specified parameters.
func (d *DNSProvider) CleanUp(domain, token, keyAuth string) error {
	err := d.provider.CleanUp(domain, token, keyAuth)
	if err != nil {
		return fmt.Errorf("failed to clean up dns, %w", err)
	}

	return nil
}

// Creates token credentials from config, if not set from environment.
func getTokenCredentials(config *Config) (azcore.TokenCredential, error) {
	if config.ClientID != "" && config.ClientSecret != "" && config.TenantID != "" {
		creds, err := azidentity.NewClientSecretCredential(config.TenantID, config.ClientID, config.ClientSecret, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get token from secret credentials, %w", err)
		}

		return creds, nil
	}

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get token using default credentials, %w", err)
	}

	return creds, nil
}

// Fetches metadata from environment or he instance metadata service.
// borrowed from https://github.com/Microsoft/azureimds/blob/master/imdssample.go
func getMetadata(config *Config, field string) (string, error) {
	metadataEndpoint := config.MetadataEndpoint
	if metadataEndpoint == "" {
		metadataEndpoint = defaultMetadataEndpoint
	}

	resource := fmt.Sprintf("%s/metadata/instance/compute/%s", metadataEndpoint, field)
	req, err := http.NewRequest(http.MethodGet, resource, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create metadata request, %w", err)
	}

	req.Header.Set("Metadata", "True")

	q := req.URL.Query()
	q.Add("format", "text")
	q.Add("api-version", "2017-12-01")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute metadata request, %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body from metadata request, %w", err)
	}

	return string(respBody), nil
}
