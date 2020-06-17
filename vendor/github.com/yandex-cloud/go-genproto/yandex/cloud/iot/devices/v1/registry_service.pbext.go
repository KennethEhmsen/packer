// Code generated by protoc-gen-goext. DO NOT EDIT.

package devices

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetRegistryRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListRegistriesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListRegistriesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRegistriesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRegistriesResponse) SetRegistries(v []*Registry) {
	m.Registries = v
}

func (m *ListRegistriesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateRegistryRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateRegistryRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateRegistryRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateRegistryRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateRegistryRequest) SetCertificates(v []*CreateRegistryRequest_Certificate) {
	m.Certificates = v
}

func (m *CreateRegistryRequest) SetPassword(v string) {
	m.Password = v
}

func (m *CreateRegistryRequest_Certificate) SetCertificateData(v string) {
	m.CertificateData = v
}

func (m *CreateRegistryMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *UpdateRegistryRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *UpdateRegistryRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateRegistryRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateRegistryRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateRegistryRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateRegistryMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListRegistryCertificatesRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListRegistryCertificatesResponse) SetCertificates(v []*RegistryCertificate) {
	m.Certificates = v
}

func (m *AddRegistryCertificateRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *AddRegistryCertificateRequest) SetCertificateData(v string) {
	m.CertificateData = v
}

func (m *AddRegistryCertificateMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *AddRegistryCertificateMetadata) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *DeleteRegistryCertificateRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryCertificateRequest) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *DeleteRegistryCertificateMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryCertificateMetadata) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *ListRegistryPasswordsRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListRegistryPasswordsResponse) SetPasswords(v []*RegistryPassword) {
	m.Passwords = v
}

func (m *AddRegistryPasswordRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *AddRegistryPasswordRequest) SetPassword(v string) {
	m.Password = v
}

func (m *AddRegistryPasswordMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *AddRegistryPasswordMetadata) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *DeleteRegistryPasswordRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryPasswordRequest) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *DeleteRegistryPasswordMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryPasswordMetadata) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *ListDeviceTopicAliasesRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListDeviceTopicAliasesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDeviceTopicAliasesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDeviceTopicAliasesResponse) SetAliases(v []*DeviceAlias) {
	m.Aliases = v
}

func (m *ListDeviceTopicAliasesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListRegistryOperationsRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListRegistryOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRegistryOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRegistryOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListRegistryOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListRegistryOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
