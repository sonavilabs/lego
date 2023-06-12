// Code generated by protoc-gen-goext. DO NOT EDIT.

package k8s

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClustersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListClustersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClustersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClustersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClustersResponse) SetClusters(v []*Cluster) {
	m.Clusters = v
}

func (m *ListClustersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *DeleteClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

type UpdateClusterRequest_InternetGateway = isUpdateClusterRequest_InternetGateway

func (m *UpdateClusterRequest) SetInternetGateway(v UpdateClusterRequest_InternetGateway) {
	m.InternetGateway = v
}

func (m *UpdateClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateClusterRequest) SetGatewayIpv4Address(v string) {
	m.InternetGateway = &UpdateClusterRequest_GatewayIpv4Address{
		GatewayIpv4Address: v,
	}
}

func (m *UpdateClusterRequest) SetMasterSpec(v *MasterUpdateSpec) {
	m.MasterSpec = v
}

func (m *UpdateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateClusterRequest) SetNodeServiceAccountId(v string) {
	m.NodeServiceAccountId = v
}

func (m *UpdateClusterRequest) SetNetworkPolicy(v *NetworkPolicy) {
	m.NetworkPolicy = v
}

func (m *UpdateClusterRequest) SetIpAllocationPolicy(v *IPAllocationPolicy) {
	m.IpAllocationPolicy = v
}

func (m *MasterUpdateSpec) SetVersion(v *UpdateVersionSpec) {
	m.Version = v
}

func (m *MasterUpdateSpec) SetMaintenancePolicy(v *MasterMaintenancePolicy) {
	m.MaintenancePolicy = v
}

func (m *MasterUpdateSpec) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *UpdateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

type CreateClusterRequest_InternetGateway = isCreateClusterRequest_InternetGateway

func (m *CreateClusterRequest) SetInternetGateway(v CreateClusterRequest_InternetGateway) {
	m.InternetGateway = v
}

type CreateClusterRequest_NetworkImplementation = isCreateClusterRequest_NetworkImplementation

func (m *CreateClusterRequest) SetNetworkImplementation(v CreateClusterRequest_NetworkImplementation) {
	m.NetworkImplementation = v
}

func (m *CreateClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateClusterRequest) SetMasterSpec(v *MasterSpec) {
	m.MasterSpec = v
}

func (m *CreateClusterRequest) SetIpAllocationPolicy(v *IPAllocationPolicy) {
	m.IpAllocationPolicy = v
}

func (m *CreateClusterRequest) SetGatewayIpv4Address(v string) {
	m.InternetGateway = &CreateClusterRequest_GatewayIpv4Address{
		GatewayIpv4Address: v,
	}
}

func (m *CreateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateClusterRequest) SetNodeServiceAccountId(v string) {
	m.NodeServiceAccountId = v
}

func (m *CreateClusterRequest) SetReleaseChannel(v ReleaseChannel) {
	m.ReleaseChannel = v
}

func (m *CreateClusterRequest) SetNetworkPolicy(v *NetworkPolicy) {
	m.NetworkPolicy = v
}

func (m *CreateClusterRequest) SetKmsProvider(v *KMSProvider) {
	m.KmsProvider = v
}

func (m *CreateClusterRequest) SetCilium(v *Cilium) {
	m.NetworkImplementation = &CreateClusterRequest_Cilium{
		Cilium: v,
	}
}

func (m *CreateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AutoUpgradeMasterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterOperationsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListClusterOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterNodeGroupsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterNodeGroupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterNodeGroupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterNodeGroupsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterNodeGroupsResponse) SetNodeGroups(v []*NodeGroup) {
	m.NodeGroups = v
}

func (m *ListClusterNodeGroupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterNodesRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterNodesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterNodesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterNodesResponse) SetNodes(v []*Node) {
	m.Nodes = v
}

func (m *ListClusterNodesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

type MasterSpec_MasterType = isMasterSpec_MasterType

func (m *MasterSpec) SetMasterType(v MasterSpec_MasterType) {
	m.MasterType = v
}

func (m *MasterSpec) SetZonalMasterSpec(v *ZonalMasterSpec) {
	m.MasterType = &MasterSpec_ZonalMasterSpec{
		ZonalMasterSpec: v,
	}
}

func (m *MasterSpec) SetRegionalMasterSpec(v *RegionalMasterSpec) {
	m.MasterType = &MasterSpec_RegionalMasterSpec{
		RegionalMasterSpec: v,
	}
}

func (m *MasterSpec) SetVersion(v string) {
	m.Version = v
}

func (m *MasterSpec) SetMaintenancePolicy(v *MasterMaintenancePolicy) {
	m.MaintenancePolicy = v
}

func (m *MasterSpec) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *ZonalMasterSpec) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *ZonalMasterSpec) SetInternalV4AddressSpec(v *InternalAddressSpec) {
	m.InternalV4AddressSpec = v
}

func (m *ZonalMasterSpec) SetExternalV4AddressSpec(v *ExternalAddressSpec) {
	m.ExternalV4AddressSpec = v
}

func (m *RegionalMasterSpec) SetRegionId(v string) {
	m.RegionId = v
}

func (m *RegionalMasterSpec) SetLocations(v []*MasterLocation) {
	m.Locations = v
}

func (m *RegionalMasterSpec) SetExternalV4AddressSpec(v *ExternalAddressSpec) {
	m.ExternalV4AddressSpec = v
}

func (m *RegionalMasterSpec) SetExternalV6AddressSpec(v *ExternalAddressSpec) {
	m.ExternalV6AddressSpec = v
}

func (m *InternalAddressSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *ExternalAddressSpec) SetAddress(v string) {
	m.Address = v
}

func (m *MasterLocation) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *MasterLocation) SetInternalV4AddressSpec(v *InternalAddressSpec) {
	m.InternalV4AddressSpec = v
}
