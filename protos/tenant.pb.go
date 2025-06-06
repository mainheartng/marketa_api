// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: protos/tenant.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tenant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subdomain     string                 `protobuf:"bytes,3,opt,name=subdomain,proto3" json:"subdomain,omitempty"`
	Status        Status                 `protobuf:"varint,4,opt,name=status,proto3,enum=types.Status" json:"status,omitempty"`
	Settings      *TenantSettings        `protobuf:"bytes,5,opt,name=settings,proto3" json:"settings,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	mi := &file_protos_tenant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetSubdomain() string {
	if x != nil {
		return x.Subdomain
	}
	return ""
}

func (x *Tenant) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *Tenant) GetSettings() *TenantSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Tenant) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Tenant) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type TenantSettings struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Currency       string                 `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Timezone       string                 `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Locale         string                 `protobuf:"bytes,3,opt,name=locale,proto3" json:"locale,omitempty"`
	CustomSettings map[string]string      `protobuf:"bytes,4,rep,name=custom_settings,json=customSettings,proto3" json:"custom_settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TenantSettings) Reset() {
	*x = TenantSettings{}
	mi := &file_protos_tenant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantSettings) ProtoMessage() {}

func (x *TenantSettings) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantSettings.ProtoReflect.Descriptor instead.
func (*TenantSettings) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *TenantSettings) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TenantSettings) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *TenantSettings) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *TenantSettings) GetCustomSettings() map[string]string {
	if x != nil {
		return x.CustomSettings
	}
	return nil
}

// Tenant Service Requests/Responses
type CreateTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Subdomain     string                 `protobuf:"bytes,2,opt,name=subdomain,proto3" json:"subdomain,omitempty"`
	Settings      *TenantSettings        `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	AdminEmail    string                 `protobuf:"bytes,4,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	mi := &file_protos_tenant_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTenantRequest) GetSubdomain() string {
	if x != nil {
		return x.Subdomain
	}
	return ""
}

func (x *CreateTenantRequest) GetSettings() *TenantSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *CreateTenantRequest) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

type CreateTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tenant        *Tenant                `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTenantResponse) Reset() {
	*x = CreateTenantResponse{}
	mi := &file_protos_tenant_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantResponse) ProtoMessage() {}

func (x *CreateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantResponse.ProtoReflect.Descriptor instead.
func (*CreateTenantResponse) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTenantResponse) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

func (x *CreateTenantResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateTenantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	mi := &file_protos_tenant_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *GetTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type GetTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tenant        *Tenant                `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantResponse) Reset() {
	*x = GetTenantResponse{}
	mi := &file_protos_tenant_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantResponse) ProtoMessage() {}

func (x *GetTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantResponse.ProtoReflect.Descriptor instead.
func (*GetTenantResponse) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *GetTenantResponse) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

func (x *GetTenantResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetTenantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Settings      *TenantSettings        `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	mi := &file_protos_tenant_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *UpdateTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTenantRequest) GetSettings() *TenantSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tenant        *Tenant                `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTenantResponse) Reset() {
	*x = UpdateTenantResponse{}
	mi := &file_protos_tenant_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantResponse) ProtoMessage() {}

func (x *UpdateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantResponse.ProtoReflect.Descriptor instead.
func (*UpdateTenantResponse) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTenantResponse) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

func (x *UpdateTenantResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateTenantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListTenantsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pagination    *PageRequest           `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	SearchQuery   string                 `protobuf:"bytes,2,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty"`
	StatusFilter  Status                 `protobuf:"varint,3,opt,name=status_filter,json=statusFilter,proto3,enum=types.Status" json:"status_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantsRequest) Reset() {
	*x = ListTenantsRequest{}
	mi := &file_protos_tenant_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsRequest) ProtoMessage() {}

func (x *ListTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsRequest.ProtoReflect.Descriptor instead.
func (*ListTenantsRequest) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{8}
}

func (x *ListTenantsRequest) GetPagination() *PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListTenantsRequest) GetSearchQuery() string {
	if x != nil {
		return x.SearchQuery
	}
	return ""
}

func (x *ListTenantsRequest) GetStatusFilter() Status {
	if x != nil {
		return x.StatusFilter
	}
	return Status_STATUS_UNSPECIFIED
}

type ListTenantsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tenants       []*Tenant              `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	Pagination    *PageResponse          `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Success       bool                   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantsResponse) Reset() {
	*x = ListTenantsResponse{}
	mi := &file_protos_tenant_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsResponse) ProtoMessage() {}

func (x *ListTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsResponse.ProtoReflect.Descriptor instead.
func (*ListTenantsResponse) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{9}
}

func (x *ListTenantsResponse) GetTenants() []*Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *ListTenantsResponse) GetPagination() *PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListTenantsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ListTenantsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ActivateTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivateTenantRequest) Reset() {
	*x = ActivateTenantRequest{}
	mi := &file_protos_tenant_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateTenantRequest) ProtoMessage() {}

func (x *ActivateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateTenantRequest.ProtoReflect.Descriptor instead.
func (*ActivateTenantRequest) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{10}
}

func (x *ActivateTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type ActivateTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivateTenantResponse) Reset() {
	*x = ActivateTenantResponse{}
	mi := &file_protos_tenant_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateTenantResponse) ProtoMessage() {}

func (x *ActivateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateTenantResponse.ProtoReflect.Descriptor instead.
func (*ActivateTenantResponse) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{11}
}

func (x *ActivateTenantResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ActivateTenantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeactivateTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Reason        string                 `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeactivateTenantRequest) Reset() {
	*x = DeactivateTenantRequest{}
	mi := &file_protos_tenant_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateTenantRequest) ProtoMessage() {}

func (x *DeactivateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateTenantRequest.ProtoReflect.Descriptor instead.
func (*DeactivateTenantRequest) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{12}
}

func (x *DeactivateTenantRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *DeactivateTenantRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type DeactivateTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeactivateTenantResponse) Reset() {
	*x = DeactivateTenantResponse{}
	mi := &file_protos_tenant_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateTenantResponse) ProtoMessage() {}

func (x *DeactivateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tenant_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateTenantResponse.ProtoReflect.Descriptor instead.
func (*DeactivateTenantResponse) Descriptor() ([]byte, []int) {
	return file_protos_tenant_proto_rawDescGZIP(), []int{13}
}

func (x *DeactivateTenantResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeactivateTenantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_tenant_proto protoreflect.FileDescriptor

const file_protos_tenant_proto_rawDesc = "" +
	"\n" +
	"\x13protos/tenant.proto\x12\atenants\x1a\x12protos/types.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa9\x02\n" +
	"\x06Tenant\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1c\n" +
	"\tsubdomain\x18\x03 \x01(\tR\tsubdomain\x12%\n" +
	"\x06status\x18\x04 \x01(\x0e2\r.types.StatusR\x06status\x123\n" +
	"\bsettings\x18\x05 \x01(\v2\x17.tenants.TenantSettingsR\bsettings\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\xf9\x01\n" +
	"\x0eTenantSettings\x12\x1a\n" +
	"\bcurrency\x18\x01 \x01(\tR\bcurrency\x12\x1a\n" +
	"\btimezone\x18\x02 \x01(\tR\btimezone\x12\x16\n" +
	"\x06locale\x18\x03 \x01(\tR\x06locale\x12T\n" +
	"\x0fcustom_settings\x18\x04 \x03(\v2+.tenants.TenantSettings.CustomSettingsEntryR\x0ecustomSettings\x1aA\n" +
	"\x13CustomSettingsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x9d\x01\n" +
	"\x13CreateTenantRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1c\n" +
	"\tsubdomain\x18\x02 \x01(\tR\tsubdomain\x123\n" +
	"\bsettings\x18\x03 \x01(\v2\x17.tenants.TenantSettingsR\bsettings\x12\x1f\n" +
	"\vadmin_email\x18\x04 \x01(\tR\n" +
	"adminEmail\"s\n" +
	"\x14CreateTenantResponse\x12'\n" +
	"\x06tenant\x18\x01 \x01(\v2\x0f.tenants.TenantR\x06tenant\x12\x18\n" +
	"\asuccess\x18\x02 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\"/\n" +
	"\x10GetTenantRequest\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\"p\n" +
	"\x11GetTenantResponse\x12'\n" +
	"\x06tenant\x18\x01 \x01(\v2\x0f.tenants.TenantR\x06tenant\x12\x18\n" +
	"\asuccess\x18\x02 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\"{\n" +
	"\x13UpdateTenantRequest\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x123\n" +
	"\bsettings\x18\x03 \x01(\v2\x17.tenants.TenantSettingsR\bsettings\"s\n" +
	"\x14UpdateTenantResponse\x12'\n" +
	"\x06tenant\x18\x01 \x01(\v2\x0f.tenants.TenantR\x06tenant\x12\x18\n" +
	"\asuccess\x18\x02 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\"\x9f\x01\n" +
	"\x12ListTenantsRequest\x122\n" +
	"\n" +
	"pagination\x18\x01 \x01(\v2\x12.types.PageRequestR\n" +
	"pagination\x12!\n" +
	"\fsearch_query\x18\x02 \x01(\tR\vsearchQuery\x122\n" +
	"\rstatus_filter\x18\x03 \x01(\x0e2\r.types.StatusR\fstatusFilter\"\xa9\x01\n" +
	"\x13ListTenantsResponse\x12)\n" +
	"\atenants\x18\x01 \x03(\v2\x0f.tenants.TenantR\atenants\x123\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x13.types.PageResponseR\n" +
	"pagination\x12\x18\n" +
	"\asuccess\x18\x03 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x04 \x01(\tR\amessage\"4\n" +
	"\x15ActivateTenantRequest\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\"L\n" +
	"\x16ActivateTenantResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"N\n" +
	"\x17DeactivateTenantRequest\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\x12\x16\n" +
	"\x06reason\x18\x02 \x01(\tR\x06reason\"N\n" +
	"\x18DeactivateTenantResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xe3\x03\n" +
	"\rTenantService\x12K\n" +
	"\fCreateTenant\x12\x1c.tenants.CreateTenantRequest\x1a\x1d.tenants.CreateTenantResponse\x12B\n" +
	"\tGetTenant\x12\x19.tenants.GetTenantRequest\x1a\x1a.tenants.GetTenantResponse\x12K\n" +
	"\fUpdateTenant\x12\x1c.tenants.UpdateTenantRequest\x1a\x1d.tenants.UpdateTenantResponse\x12H\n" +
	"\vListTenants\x12\x1b.tenants.ListTenantsRequest\x1a\x1c.tenants.ListTenantsResponse\x12Q\n" +
	"\x0eActivateTenant\x12\x1e.tenants.ActivateTenantRequest\x1a\x1f.tenants.ActivateTenantResponse\x12W\n" +
	"\x10DeactivateTenant\x12 .tenants.DeactivateTenantRequest\x1a!.tenants.DeactivateTenantResponseB2Z0github.com/mainheartng/marketa_api/protos;protosb\x06proto3"

var (
	file_protos_tenant_proto_rawDescOnce sync.Once
	file_protos_tenant_proto_rawDescData []byte
)

func file_protos_tenant_proto_rawDescGZIP() []byte {
	file_protos_tenant_proto_rawDescOnce.Do(func() {
		file_protos_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_tenant_proto_rawDesc), len(file_protos_tenant_proto_rawDesc)))
	})
	return file_protos_tenant_proto_rawDescData
}

var file_protos_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_protos_tenant_proto_goTypes = []any{
	(*Tenant)(nil),                   // 0: tenants.Tenant
	(*TenantSettings)(nil),           // 1: tenants.TenantSettings
	(*CreateTenantRequest)(nil),      // 2: tenants.CreateTenantRequest
	(*CreateTenantResponse)(nil),     // 3: tenants.CreateTenantResponse
	(*GetTenantRequest)(nil),         // 4: tenants.GetTenantRequest
	(*GetTenantResponse)(nil),        // 5: tenants.GetTenantResponse
	(*UpdateTenantRequest)(nil),      // 6: tenants.UpdateTenantRequest
	(*UpdateTenantResponse)(nil),     // 7: tenants.UpdateTenantResponse
	(*ListTenantsRequest)(nil),       // 8: tenants.ListTenantsRequest
	(*ListTenantsResponse)(nil),      // 9: tenants.ListTenantsResponse
	(*ActivateTenantRequest)(nil),    // 10: tenants.ActivateTenantRequest
	(*ActivateTenantResponse)(nil),   // 11: tenants.ActivateTenantResponse
	(*DeactivateTenantRequest)(nil),  // 12: tenants.DeactivateTenantRequest
	(*DeactivateTenantResponse)(nil), // 13: tenants.DeactivateTenantResponse
	nil,                              // 14: tenants.TenantSettings.CustomSettingsEntry
	(Status)(0),                      // 15: types.Status
	(*timestamppb.Timestamp)(nil),    // 16: google.protobuf.Timestamp
	(*PageRequest)(nil),              // 17: types.PageRequest
	(*PageResponse)(nil),             // 18: types.PageResponse
}
var file_protos_tenant_proto_depIdxs = []int32{
	15, // 0: tenants.Tenant.status:type_name -> types.Status
	1,  // 1: tenants.Tenant.settings:type_name -> tenants.TenantSettings
	16, // 2: tenants.Tenant.created_at:type_name -> google.protobuf.Timestamp
	16, // 3: tenants.Tenant.updated_at:type_name -> google.protobuf.Timestamp
	14, // 4: tenants.TenantSettings.custom_settings:type_name -> tenants.TenantSettings.CustomSettingsEntry
	1,  // 5: tenants.CreateTenantRequest.settings:type_name -> tenants.TenantSettings
	0,  // 6: tenants.CreateTenantResponse.tenant:type_name -> tenants.Tenant
	0,  // 7: tenants.GetTenantResponse.tenant:type_name -> tenants.Tenant
	1,  // 8: tenants.UpdateTenantRequest.settings:type_name -> tenants.TenantSettings
	0,  // 9: tenants.UpdateTenantResponse.tenant:type_name -> tenants.Tenant
	17, // 10: tenants.ListTenantsRequest.pagination:type_name -> types.PageRequest
	15, // 11: tenants.ListTenantsRequest.status_filter:type_name -> types.Status
	0,  // 12: tenants.ListTenantsResponse.tenants:type_name -> tenants.Tenant
	18, // 13: tenants.ListTenantsResponse.pagination:type_name -> types.PageResponse
	2,  // 14: tenants.TenantService.CreateTenant:input_type -> tenants.CreateTenantRequest
	4,  // 15: tenants.TenantService.GetTenant:input_type -> tenants.GetTenantRequest
	6,  // 16: tenants.TenantService.UpdateTenant:input_type -> tenants.UpdateTenantRequest
	8,  // 17: tenants.TenantService.ListTenants:input_type -> tenants.ListTenantsRequest
	10, // 18: tenants.TenantService.ActivateTenant:input_type -> tenants.ActivateTenantRequest
	12, // 19: tenants.TenantService.DeactivateTenant:input_type -> tenants.DeactivateTenantRequest
	3,  // 20: tenants.TenantService.CreateTenant:output_type -> tenants.CreateTenantResponse
	5,  // 21: tenants.TenantService.GetTenant:output_type -> tenants.GetTenantResponse
	7,  // 22: tenants.TenantService.UpdateTenant:output_type -> tenants.UpdateTenantResponse
	9,  // 23: tenants.TenantService.ListTenants:output_type -> tenants.ListTenantsResponse
	11, // 24: tenants.TenantService.ActivateTenant:output_type -> tenants.ActivateTenantResponse
	13, // 25: tenants.TenantService.DeactivateTenant:output_type -> tenants.DeactivateTenantResponse
	20, // [20:26] is the sub-list for method output_type
	14, // [14:20] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_protos_tenant_proto_init() }
func file_protos_tenant_proto_init() {
	if File_protos_tenant_proto != nil {
		return
	}
	file_protos_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_tenant_proto_rawDesc), len(file_protos_tenant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_tenant_proto_goTypes,
		DependencyIndexes: file_protos_tenant_proto_depIdxs,
		MessageInfos:      file_protos_tenant_proto_msgTypes,
	}.Build()
	File_protos_tenant_proto = out.File
	file_protos_tenant_proto_goTypes = nil
	file_protos_tenant_proto_depIdxs = nil
}
