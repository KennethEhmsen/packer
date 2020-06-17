// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetInstanceRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *GetInstanceRequest) SetView(v InstanceView) {
	m.View = v
}

func (m *ListInstancesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListInstancesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListInstancesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListInstancesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListInstancesResponse) SetInstances(v []*Instance) {
	m.Instances = v
}

func (m *ListInstancesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateInstanceRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateInstanceRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateInstanceRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateInstanceRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateInstanceRequest) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *CreateInstanceRequest) SetPlatformId(v string) {
	m.PlatformId = v
}

func (m *CreateInstanceRequest) SetResourcesSpec(v *ResourcesSpec) {
	m.ResourcesSpec = v
}

func (m *CreateInstanceRequest) SetMetadata(v map[string]string) {
	m.Metadata = v
}

func (m *CreateInstanceRequest) SetBootDiskSpec(v *AttachedDiskSpec) {
	m.BootDiskSpec = v
}

func (m *CreateInstanceRequest) SetSecondaryDiskSpecs(v []*AttachedDiskSpec) {
	m.SecondaryDiskSpecs = v
}

func (m *CreateInstanceRequest) SetNetworkInterfaceSpecs(v []*NetworkInterfaceSpec) {
	m.NetworkInterfaceSpecs = v
}

func (m *CreateInstanceRequest) SetHostname(v string) {
	m.Hostname = v
}

func (m *CreateInstanceRequest) SetSchedulingPolicy(v *SchedulingPolicy) {
	m.SchedulingPolicy = v
}

func (m *CreateInstanceRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateInstanceRequest) SetNetworkSettings(v *NetworkSettings) {
	m.NetworkSettings = v
}

func (m *CreateInstanceRequest) SetPlacementPolicy(v *PlacementPolicy) {
	m.PlacementPolicy = v
}

func (m *CreateInstanceMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *UpdateInstanceRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *UpdateInstanceRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateInstanceRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateInstanceRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateInstanceRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateInstanceRequest) SetPlatformId(v string) {
	m.PlatformId = v
}

func (m *UpdateInstanceRequest) SetResourcesSpec(v *ResourcesSpec) {
	m.ResourcesSpec = v
}

func (m *UpdateInstanceRequest) SetMetadata(v map[string]string) {
	m.Metadata = v
}

func (m *UpdateInstanceRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateInstanceRequest) SetNetworkSettings(v *NetworkSettings) {
	m.NetworkSettings = v
}

func (m *UpdateInstanceRequest) SetPlacementPolicy(v *PlacementPolicy) {
	m.PlacementPolicy = v
}

func (m *UpdateInstanceMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *DeleteInstanceRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *DeleteInstanceMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *UpdateInstanceMetadataRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *UpdateInstanceMetadataRequest) SetDelete(v []string) {
	m.Delete = v
}

func (m *UpdateInstanceMetadataRequest) SetUpsert(v map[string]string) {
	m.Upsert = v
}

func (m *UpdateInstanceMetadataMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *GetInstanceSerialPortOutputRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *GetInstanceSerialPortOutputRequest) SetPort(v int64) {
	m.Port = v
}

func (m *GetInstanceSerialPortOutputResponse) SetContents(v string) {
	m.Contents = v
}

func (m *StopInstanceRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *StopInstanceMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *StartInstanceRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *StartInstanceMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *RestartInstanceRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *RestartInstanceMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *AttachInstanceDiskRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *AttachInstanceDiskRequest) SetAttachedDiskSpec(v *AttachedDiskSpec) {
	m.AttachedDiskSpec = v
}

func (m *AttachInstanceDiskMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *AttachInstanceDiskMetadata) SetDiskId(v string) {
	m.DiskId = v
}

type DetachInstanceDiskRequest_Disk = isDetachInstanceDiskRequest_Disk

func (m *DetachInstanceDiskRequest) SetDisk(v DetachInstanceDiskRequest_Disk) {
	m.Disk = v
}

func (m *DetachInstanceDiskRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *DetachInstanceDiskRequest) SetDiskId(v string) {
	m.Disk = &DetachInstanceDiskRequest_DiskId{
		DiskId: v,
	}
}

func (m *DetachInstanceDiskRequest) SetDeviceName(v string) {
	m.Disk = &DetachInstanceDiskRequest_DeviceName{
		DeviceName: v,
	}
}

func (m *DetachInstanceDiskMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *DetachInstanceDiskMetadata) SetDiskId(v string) {
	m.DiskId = v
}

func (m *AddInstanceOneToOneNatRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *AddInstanceOneToOneNatRequest) SetNetworkInterfaceIndex(v string) {
	m.NetworkInterfaceIndex = v
}

func (m *AddInstanceOneToOneNatRequest) SetInternalAddress(v string) {
	m.InternalAddress = v
}

func (m *AddInstanceOneToOneNatRequest) SetOneToOneNatSpec(v *OneToOneNatSpec) {
	m.OneToOneNatSpec = v
}

func (m *AddInstanceOneToOneNatMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *RemoveInstanceOneToOneNatRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *RemoveInstanceOneToOneNatRequest) SetNetworkInterfaceIndex(v string) {
	m.NetworkInterfaceIndex = v
}

func (m *RemoveInstanceOneToOneNatRequest) SetInternalAddress(v string) {
	m.InternalAddress = v
}

func (m *RemoveInstanceOneToOneNatMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *UpdateInstanceNetworkInterfaceRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *UpdateInstanceNetworkInterfaceRequest) SetNetworkInterfaceIndex(v string) {
	m.NetworkInterfaceIndex = v
}

func (m *UpdateInstanceNetworkInterfaceRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateInstanceNetworkInterfaceRequest) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *UpdateInstanceNetworkInterfaceRequest) SetPrimaryV4AddressSpec(v *PrimaryAddressSpec) {
	m.PrimaryV4AddressSpec = v
}

func (m *UpdateInstanceNetworkInterfaceRequest) SetPrimaryV6AddressSpec(v *PrimaryAddressSpec) {
	m.PrimaryV6AddressSpec = v
}

func (m *UpdateInstanceNetworkInterfaceRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *UpdateInstanceNetworkInterfaceMetadata) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *UpdateInstanceNetworkInterfaceMetadata) SetNetworkInterfaceIndex(v string) {
	m.NetworkInterfaceIndex = v
}

func (m *ListInstanceOperationsRequest) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *ListInstanceOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListInstanceOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListInstanceOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListInstanceOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ResourcesSpec) SetMemory(v int64) {
	m.Memory = v
}

func (m *ResourcesSpec) SetCores(v int64) {
	m.Cores = v
}

func (m *ResourcesSpec) SetCoreFraction(v int64) {
	m.CoreFraction = v
}

func (m *ResourcesSpec) SetGpus(v int64) {
	m.Gpus = v
}

type AttachedDiskSpec_Disk = isAttachedDiskSpec_Disk

func (m *AttachedDiskSpec) SetDisk(v AttachedDiskSpec_Disk) {
	m.Disk = v
}

func (m *AttachedDiskSpec) SetMode(v AttachedDiskSpec_Mode) {
	m.Mode = v
}

func (m *AttachedDiskSpec) SetDeviceName(v string) {
	m.DeviceName = v
}

func (m *AttachedDiskSpec) SetAutoDelete(v bool) {
	m.AutoDelete = v
}

func (m *AttachedDiskSpec) SetDiskSpec(v *AttachedDiskSpec_DiskSpec) {
	m.Disk = &AttachedDiskSpec_DiskSpec_{
		DiskSpec: v,
	}
}

func (m *AttachedDiskSpec) SetDiskId(v string) {
	m.Disk = &AttachedDiskSpec_DiskId{
		DiskId: v,
	}
}

type AttachedDiskSpec_DiskSpec_Source = isAttachedDiskSpec_DiskSpec_Source

func (m *AttachedDiskSpec_DiskSpec) SetSource(v AttachedDiskSpec_DiskSpec_Source) {
	m.Source = v
}

func (m *AttachedDiskSpec_DiskSpec) SetName(v string) {
	m.Name = v
}

func (m *AttachedDiskSpec_DiskSpec) SetDescription(v string) {
	m.Description = v
}

func (m *AttachedDiskSpec_DiskSpec) SetTypeId(v string) {
	m.TypeId = v
}

func (m *AttachedDiskSpec_DiskSpec) SetSize(v int64) {
	m.Size = v
}

func (m *AttachedDiskSpec_DiskSpec) SetImageId(v string) {
	m.Source = &AttachedDiskSpec_DiskSpec_ImageId{
		ImageId: v,
	}
}

func (m *AttachedDiskSpec_DiskSpec) SetSnapshotId(v string) {
	m.Source = &AttachedDiskSpec_DiskSpec_SnapshotId{
		SnapshotId: v,
	}
}

func (m *NetworkInterfaceSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *NetworkInterfaceSpec) SetPrimaryV4AddressSpec(v *PrimaryAddressSpec) {
	m.PrimaryV4AddressSpec = v
}

func (m *NetworkInterfaceSpec) SetPrimaryV6AddressSpec(v *PrimaryAddressSpec) {
	m.PrimaryV6AddressSpec = v
}

func (m *NetworkInterfaceSpec) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *PrimaryAddressSpec) SetAddress(v string) {
	m.Address = v
}

func (m *PrimaryAddressSpec) SetOneToOneNatSpec(v *OneToOneNatSpec) {
	m.OneToOneNatSpec = v
}

func (m *OneToOneNatSpec) SetIpVersion(v IpVersion) {
	m.IpVersion = v
}

func (m *OneToOneNatSpec) SetAddress(v string) {
	m.Address = v
}
