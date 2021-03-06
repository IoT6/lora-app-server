// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProfileSettings struct {
	// Existing users in the system can not be assigned to organizations and
	// application and can not be listed by non global admin users.
	DisableAssignExistingUsers bool     `protobuf:"varint,1,opt,name=disable_assign_existing_users,json=disableAssignExistingUsers,proto3" json:"disable_assign_existing_users,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *ProfileSettings) Reset()         { *m = ProfileSettings{} }
func (m *ProfileSettings) String() string { return proto.CompactTextString(m) }
func (*ProfileSettings) ProtoMessage()    {}
func (*ProfileSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{0}
}

func (m *ProfileSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileSettings.Unmarshal(m, b)
}
func (m *ProfileSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileSettings.Marshal(b, m, deterministic)
}
func (m *ProfileSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileSettings.Merge(m, src)
}
func (m *ProfileSettings) XXX_Size() int {
	return xxx_messageInfo_ProfileSettings.Size(m)
}
func (m *ProfileSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileSettings proto.InternalMessageInfo

func (m *ProfileSettings) GetDisableAssignExistingUsers() bool {
	if m != nil {
		return m.DisableAssignExistingUsers
	}
	return false
}

// Defines an organization to which an user is associated.
type OrganizationLink struct {
	// Organization ID.
	OrganizationId int64 `protobuf:"varint,1,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Organization name.
	OrganizationName string `protobuf:"bytes,2,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	// User is admin within the context of this organization.
	IsAdmin bool `protobuf:"varint,3,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrganizationLink) Reset()         { *m = OrganizationLink{} }
func (m *OrganizationLink) String() string { return proto.CompactTextString(m) }
func (*OrganizationLink) ProtoMessage()    {}
func (*OrganizationLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{1}
}

func (m *OrganizationLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationLink.Unmarshal(m, b)
}
func (m *OrganizationLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationLink.Marshal(b, m, deterministic)
}
func (m *OrganizationLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationLink.Merge(m, src)
}
func (m *OrganizationLink) XXX_Size() int {
	return xxx_messageInfo_OrganizationLink.Size(m)
}
func (m *OrganizationLink) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationLink.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationLink proto.InternalMessageInfo

func (m *OrganizationLink) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *OrganizationLink) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *OrganizationLink) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *OrganizationLink) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OrganizationLink) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type LoginRequest struct {
	// Username of the user.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Password of the user.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	// The JWT tag to be used to access lora-app-server interfaces.
	Jwt                  string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type ProfileResponse struct {
	// User object.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Organizations to which the user is associated.
	Organizations []*OrganizationLink `protobuf:"bytes,3,rep,name=organizations,proto3" json:"organizations,omitempty"`
	// Profile settings.
	Settings             *ProfileSettings `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProfileResponse) Reset()         { *m = ProfileResponse{} }
func (m *ProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ProfileResponse) ProtoMessage()    {}
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{4}
}

func (m *ProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileResponse.Unmarshal(m, b)
}
func (m *ProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileResponse.Marshal(b, m, deterministic)
}
func (m *ProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileResponse.Merge(m, src)
}
func (m *ProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ProfileResponse.Size(m)
}
func (m *ProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileResponse proto.InternalMessageInfo

func (m *ProfileResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ProfileResponse) GetOrganizations() []*OrganizationLink {
	if m != nil {
		return m.Organizations
	}
	return nil
}

func (m *ProfileResponse) GetSettings() *ProfileSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type GlobalSearchRequest struct {
	// Search query.
	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	// Max number of results to return.
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset offset of the result-set (for pagination).
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalSearchRequest) Reset()         { *m = GlobalSearchRequest{} }
func (m *GlobalSearchRequest) String() string { return proto.CompactTextString(m) }
func (*GlobalSearchRequest) ProtoMessage()    {}
func (*GlobalSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{5}
}

func (m *GlobalSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalSearchRequest.Unmarshal(m, b)
}
func (m *GlobalSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalSearchRequest.Marshal(b, m, deterministic)
}
func (m *GlobalSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSearchRequest.Merge(m, src)
}
func (m *GlobalSearchRequest) XXX_Size() int {
	return xxx_messageInfo_GlobalSearchRequest.Size(m)
}
func (m *GlobalSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSearchRequest proto.InternalMessageInfo

func (m *GlobalSearchRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *GlobalSearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GlobalSearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type GlobalSearchResponse struct {
	Result               []*GlobalSearchResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GlobalSearchResponse) Reset()         { *m = GlobalSearchResponse{} }
func (m *GlobalSearchResponse) String() string { return proto.CompactTextString(m) }
func (*GlobalSearchResponse) ProtoMessage()    {}
func (*GlobalSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{6}
}

func (m *GlobalSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalSearchResponse.Unmarshal(m, b)
}
func (m *GlobalSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalSearchResponse.Marshal(b, m, deterministic)
}
func (m *GlobalSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSearchResponse.Merge(m, src)
}
func (m *GlobalSearchResponse) XXX_Size() int {
	return xxx_messageInfo_GlobalSearchResponse.Size(m)
}
func (m *GlobalSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSearchResponse proto.InternalMessageInfo

func (m *GlobalSearchResponse) GetResult() []*GlobalSearchResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type GlobalSearchResult struct {
	// Record kind.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Search score.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Organization id.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Organization name.
	OrganizationName string `protobuf:"bytes,4,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	// Application id.
	ApplicationId int64 `protobuf:"varint,5,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	// Application name.
	ApplicationName string `protobuf:"bytes,6,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// Device DevEUI (hex encoded).
	DeviceDevEui string `protobuf:"bytes,7,opt,name=device_dev_eui,json=deviceDevEUI,proto3" json:"device_dev_eui,omitempty"`
	// Device name.
	DeviceName string `protobuf:"bytes,8,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	// Gateway MAC (hex encoded).
	GatewayMac string `protobuf:"bytes,9,opt,name=gateway_mac,json=gatewayMAC,proto3" json:"gateway_mac,omitempty"`
	// Gateway name.
	GatewayName          string   `protobuf:"bytes,10,opt,name=gateway_name,json=gatewayName,proto3" json:"gateway_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalSearchResult) Reset()         { *m = GlobalSearchResult{} }
func (m *GlobalSearchResult) String() string { return proto.CompactTextString(m) }
func (*GlobalSearchResult) ProtoMessage()    {}
func (*GlobalSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{7}
}

func (m *GlobalSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalSearchResult.Unmarshal(m, b)
}
func (m *GlobalSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalSearchResult.Marshal(b, m, deterministic)
}
func (m *GlobalSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSearchResult.Merge(m, src)
}
func (m *GlobalSearchResult) XXX_Size() int {
	return xxx_messageInfo_GlobalSearchResult.Size(m)
}
func (m *GlobalSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSearchResult proto.InternalMessageInfo

func (m *GlobalSearchResult) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *GlobalSearchResult) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *GlobalSearchResult) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *GlobalSearchResult) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *GlobalSearchResult) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *GlobalSearchResult) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *GlobalSearchResult) GetDeviceDevEui() string {
	if m != nil {
		return m.DeviceDevEui
	}
	return ""
}

func (m *GlobalSearchResult) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *GlobalSearchResult) GetGatewayMac() string {
	if m != nil {
		return m.GatewayMac
	}
	return ""
}

func (m *GlobalSearchResult) GetGatewayName() string {
	if m != nil {
		return m.GatewayName
	}
	return ""
}

type BrandingResponse struct {
	// Logo html.
	Logo string `protobuf:"bytes,1,opt,name=logo,proto3" json:"logo,omitempty"`
	// Registration html.
	Registration string `protobuf:"bytes,2,opt,name=registration,proto3" json:"registration,omitempty"`
	// Footer html.
	Footer               string   `protobuf:"bytes,3,opt,name=footer,proto3" json:"footer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrandingResponse) Reset()         { *m = BrandingResponse{} }
func (m *BrandingResponse) String() string { return proto.CompactTextString(m) }
func (*BrandingResponse) ProtoMessage()    {}
func (*BrandingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{8}
}

func (m *BrandingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrandingResponse.Unmarshal(m, b)
}
func (m *BrandingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrandingResponse.Marshal(b, m, deterministic)
}
func (m *BrandingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrandingResponse.Merge(m, src)
}
func (m *BrandingResponse) XXX_Size() int {
	return xxx_messageInfo_BrandingResponse.Size(m)
}
func (m *BrandingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BrandingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BrandingResponse proto.InternalMessageInfo

func (m *BrandingResponse) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *BrandingResponse) GetRegistration() string {
	if m != nil {
		return m.Registration
	}
	return ""
}

func (m *BrandingResponse) GetFooter() string {
	if m != nil {
		return m.Footer
	}
	return ""
}

func init() {
	proto.RegisterType((*ProfileSettings)(nil), "api.ProfileSettings")
	proto.RegisterType((*OrganizationLink)(nil), "api.OrganizationLink")
	proto.RegisterType((*LoginRequest)(nil), "api.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "api.LoginResponse")
	proto.RegisterType((*ProfileResponse)(nil), "api.ProfileResponse")
	proto.RegisterType((*GlobalSearchRequest)(nil), "api.GlobalSearchRequest")
	proto.RegisterType((*GlobalSearchResponse)(nil), "api.GlobalSearchResponse")
	proto.RegisterType((*GlobalSearchResult)(nil), "api.GlobalSearchResult")
	proto.RegisterType((*BrandingResponse)(nil), "api.BrandingResponse")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor_41f4a519b878ee3b) }

var fileDescriptor_41f4a519b878ee3b = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x8e, 0xdb, 0x44,
	0x14, 0x86, 0xe5, 0x38, 0xbb, 0x9b, 0x9c, 0x64, 0x77, 0xd3, 0x69, 0x9a, 0xba, 0xa1, 0xcb, 0xa6,
	0x16, 0x88, 0x00, 0x52, 0x82, 0xc2, 0x15, 0x70, 0x15, 0xd8, 0xa5, 0x8a, 0x54, 0x0a, 0xf2, 0xb6,
	0x48, 0x88, 0x0b, 0x6b, 0x12, 0x4f, 0xcc, 0x50, 0x7b, 0xc6, 0x78, 0x26, 0xbb, 0x94, 0x4b, 0x5e,
	0x81, 0x07, 0xe0, 0xa1, 0x78, 0x05, 0x5e, 0x80, 0x2b, 0xc4, 0x1d, 0x9a, 0xe3, 0xb1, 0x65, 0xa7,
	0xbb, 0x82, 0x3b, 0x9f, 0x33, 0xdf, 0xfc, 0x3e, 0xf3, 0xcf, 0x39, 0x03, 0x27, 0x5c, 0x68, 0x96,
	0x0b, 0x9a, 0xcc, 0xb2, 0x5c, 0x6a, 0x49, 0x5c, 0x9a, 0xf1, 0xf1, 0xe3, 0x58, 0xca, 0x38, 0x61,
	0x73, 0x9a, 0xf1, 0x39, 0x15, 0x42, 0x6a, 0xaa, 0xb9, 0x14, 0xaa, 0x40, 0xc6, 0xe7, 0x76, 0x15,
	0xa3, 0xf5, 0x6e, 0x3b, 0xd7, 0x3c, 0x65, 0x4a, 0xd3, 0x34, 0xb3, 0xc0, 0x5b, 0xfb, 0x00, 0x4b,
	0x33, 0xfd, 0xda, 0x2e, 0xc2, 0x4e, 0xb1, 0xbc, 0xf8, 0xf6, 0x5f, 0xc0, 0xe9, 0x37, 0xb9, 0xdc,
	0xf2, 0x84, 0x5d, 0x31, 0xad, 0xb9, 0x88, 0x15, 0x59, 0xc2, 0x59, 0xc4, 0x15, 0x5d, 0x27, 0x2c,
	0xa4, 0x4a, 0xf1, 0x58, 0x84, 0xec, 0x67, 0xae, 0xcc, 0x5a, 0x68, 0x36, 0x2a, 0xcf, 0x99, 0x38,
	0xd3, 0x4e, 0x30, 0xb6, 0xd0, 0x12, 0x99, 0x4b, 0x8b, 0xbc, 0x34, 0x84, 0xff, 0x8f, 0x03, 0x83,
	0xaf, 0xf3, 0x98, 0x0a, 0xfe, 0x0b, 0xd6, 0xfd, 0x8c, 0x8b, 0x57, 0xe4, 0x3d, 0x38, 0x95, 0xb5,
	0x5c, 0xc8, 0x23, 0x54, 0x72, 0x83, 0x93, 0x7a, 0x7a, 0x75, 0x41, 0x3e, 0x84, 0x7b, 0x0d, 0x50,
	0xd0, 0x94, 0x79, 0xad, 0x89, 0x33, 0xed, 0x06, 0x83, 0xfa, 0xc2, 0x73, 0x9a, 0x32, 0xf2, 0x08,
	0x3a, 0x5c, 0x85, 0x34, 0x4a, 0xb9, 0xf0, 0x5c, 0x2c, 0xec, 0x88, 0xab, 0xa5, 0x09, 0xc9, 0x27,
	0x00, 0x9b, 0x9c, 0x51, 0xcd, 0xa2, 0x90, 0x6a, 0xaf, 0x3d, 0x71, 0xa6, 0xbd, 0xc5, 0x78, 0x56,
	0x38, 0x33, 0x2b, 0x9d, 0x99, 0xbd, 0x28, 0xad, 0x0b, 0xba, 0x96, 0x5e, 0x6a, 0xb3, 0x75, 0x97,
	0x45, 0xe5, 0xd6, 0x83, 0xff, 0xde, 0x6a, 0xe9, 0xa5, 0xf6, 0xbf, 0x84, 0xfe, 0x33, 0x19, 0x73,
	0x11, 0xb0, 0x9f, 0x76, 0x4c, 0x69, 0x32, 0x86, 0x8e, 0xb1, 0x0d, 0x0f, 0xe1, 0xe0, 0x21, 0xaa,
	0xd8, 0xac, 0x65, 0x54, 0xa9, 0x1b, 0x99, 0x47, 0xf6, 0x80, 0x55, 0xec, 0x3f, 0x81, 0x63, 0xab,
	0xa3, 0x32, 0x29, 0x14, 0x23, 0x03, 0x70, 0x7f, 0xbc, 0xd1, 0x56, 0xc3, 0x7c, 0xfa, 0xbf, 0x3b,
	0xd5, 0xed, 0x55, 0xd4, 0x19, 0xb4, 0x8d, 0x3c, 0x62, 0xbd, 0x45, 0x77, 0x46, 0x33, 0x3e, 0x33,
	0x97, 0x12, 0x60, 0x9a, 0x7c, 0x06, 0xc7, 0x75, 0x0b, 0x95, 0xe7, 0x4e, 0xdc, 0x69, 0x6f, 0xf1,
	0x00, 0xb9, 0xfd, 0x2b, 0x0b, 0x9a, 0x2c, 0xf9, 0x08, 0x3a, 0xca, 0x76, 0x89, 0xb5, 0x73, 0x88,
	0xfb, 0xf6, 0x3a, 0x28, 0xa8, 0x28, 0xff, 0x7b, 0xb8, 0xff, 0x34, 0x91, 0x6b, 0x9a, 0x5c, 0x31,
	0x9a, 0x6f, 0x7e, 0x28, 0x3d, 0x19, 0xc1, 0xa1, 0xc2, 0x84, 0x3d, 0x8d, 0x8d, 0xc8, 0x10, 0x0e,
	0x12, 0x9e, 0x72, 0x8d, 0x66, 0xb8, 0x41, 0x11, 0x18, 0x5a, 0x6e, 0xb7, 0x8a, 0x69, 0xbc, 0x60,
	0x37, 0xb0, 0x91, 0xff, 0x14, 0x86, 0x4d, 0x71, 0x6b, 0xc1, 0x1c, 0x0e, 0x73, 0xa6, 0x76, 0x89,
	0xf1, 0xca, 0x1c, 0xee, 0x21, 0x16, 0xb9, 0x87, 0xee, 0x12, 0x1d, 0x58, 0xcc, 0xff, 0xab, 0x05,
	0xe4, 0xcd, 0x65, 0x42, 0xa0, 0xfd, 0x8a, 0x8b, 0xc8, 0xd6, 0x88, 0xdf, 0xa6, 0x42, 0xb5, 0x91,
	0x79, 0xd1, 0x8f, 0xad, 0xa0, 0x08, 0x6e, 0x6b, 0x6d, 0xf7, 0xff, 0xb7, 0x76, 0xfb, 0x8e, 0xd6,
	0x7e, 0x17, 0x4e, 0x68, 0x96, 0x25, 0x7c, 0x53, 0x89, 0x1e, 0xa0, 0xe8, 0x71, 0x2d, 0xbb, 0xba,
	0x20, 0xef, 0xc3, 0xa0, 0x8e, 0xa1, 0xe4, 0x21, 0x4a, 0x9e, 0xd6, 0xf2, 0xa8, 0xf8, 0x0e, 0x9c,
	0x44, 0xec, 0x9a, 0x6f, 0x58, 0x18, 0xb1, 0xeb, 0x90, 0xed, 0xb8, 0x77, 0x84, 0x60, 0xbf, 0xc8,
	0x5e, 0xb0, 0xeb, 0xcb, 0x97, 0x2b, 0x72, 0x0e, 0x3d, 0x4b, 0xa1, 0x56, 0x07, 0x11, 0x28, 0x52,
	0x28, 0x73, 0x0e, 0xbd, 0x98, 0x6a, 0x76, 0x43, 0x5f, 0x87, 0x29, 0xdd, 0x78, 0xdd, 0x02, 0xb0,
	0xa9, 0xaf, 0x96, 0x5f, 0x90, 0x27, 0xd0, 0x2f, 0x01, 0x94, 0x00, 0x24, 0xca, 0x4d, 0x46, 0xc3,
	0x5f, 0xc3, 0xe0, 0xf3, 0x9c, 0x8a, 0x88, 0x8b, 0xb8, 0xba, 0x38, 0x02, 0xed, 0x44, 0xc6, 0xb2,
	0x34, 0xdc, 0x7c, 0x13, 0x1f, 0xfa, 0x39, 0x8b, 0xb9, 0xd2, 0x39, 0x1e, 0xc3, 0x8e, 0x49, 0x23,
	0x67, 0x1a, 0x64, 0x2b, 0xa5, 0x66, 0x39, 0xba, 0xde, 0x0d, 0x6c, 0xb4, 0xf8, 0xbb, 0x05, 0xa7,
	0x2b, 0xfb, 0xb8, 0x5e, 0xb1, 0xdc, 0xd4, 0x4f, 0x9e, 0xc3, 0x01, 0x8e, 0x15, 0xb9, 0x87, 0x5d,
	0x51, 0x1f, 0xd5, 0x31, 0xa9, 0xa7, 0x8a, 0x9a, 0xfc, 0xb7, 0x7f, 0xfd, 0xe3, 0xcf, 0xdf, 0x5a,
	0x9e, 0x7f, 0x1f, 0x9f, 0xe2, 0xf2, 0xa9, 0x9e, 0x27, 0x06, 0xfa, 0xd4, 0xf9, 0x80, 0x7c, 0x0b,
	0x47, 0xb6, 0xfd, 0xc9, 0xe8, 0x8d, 0x07, 0xe2, 0xd2, 0xbc, 0xba, 0xe3, 0xc6, 0x90, 0x54, 0xc2,
	0x67, 0x28, 0xfc, 0x90, 0x3c, 0x68, 0x0a, 0x67, 0x56, 0xec, 0x3b, 0xe8, 0x94, 0xfe, 0xdc, 0x29,
	0x5c, 0x4c, 0xed, 0xbe, 0x8d, 0x65, 0xc9, 0x64, 0xd4, 0x54, 0x5e, 0x97, 0x72, 0x14, 0xfa, 0xf5,
	0x6e, 0x27, 0xde, 0x2d, 0xf3, 0x51, 0x18, 0xf2, 0xe8, 0xb6, 0xc9, 0x29, 0x7e, 0xf2, 0x18, 0x7f,
	0x32, 0x22, 0xc3, 0xe6, 0x4f, 0x8a, 0x41, 0x5e, 0x1f, 0x62, 0xa5, 0x1f, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x59, 0x32, 0x94, 0x5a, 0xdc, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InternalServiceClient is the client API for InternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalServiceClient interface {
	// Log in a user
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Get the current user's profile
	Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileResponse, error)
	// Get the branding for the UI
	Branding(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BrandingResponse, error)
	// Perform a global search.
	GlobalSearch(ctx context.Context, in *GlobalSearchRequest, opts ...grpc.CallOption) (*GlobalSearchResponse, error)
}

type internalServiceClient struct {
	cc *grpc.ClientConn
}

func NewInternalServiceClient(cc *grpc.ClientConn) InternalServiceClient {
	return &internalServiceClient{cc}
}

func (c *internalServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/api.InternalService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/api.InternalService/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Branding(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BrandingResponse, error) {
	out := new(BrandingResponse)
	err := c.cc.Invoke(ctx, "/api.InternalService/Branding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) GlobalSearch(ctx context.Context, in *GlobalSearchRequest, opts ...grpc.CallOption) (*GlobalSearchResponse, error) {
	out := new(GlobalSearchResponse)
	err := c.cc.Invoke(ctx, "/api.InternalService/GlobalSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServiceServer is the server API for InternalService service.
type InternalServiceServer interface {
	// Log in a user
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Get the current user's profile
	Profile(context.Context, *empty.Empty) (*ProfileResponse, error)
	// Get the branding for the UI
	Branding(context.Context, *empty.Empty) (*BrandingResponse, error)
	// Perform a global search.
	GlobalSearch(context.Context, *GlobalSearchRequest) (*GlobalSearchResponse, error)
}

func RegisterInternalServiceServer(s *grpc.Server, srv InternalServiceServer) {
	s.RegisterService(&_InternalService_serviceDesc, srv)
}

func _InternalService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InternalService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InternalService/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Profile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Branding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Branding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InternalService/Branding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Branding(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_GlobalSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).GlobalSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InternalService/GlobalSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).GlobalSearch(ctx, req.(*GlobalSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.InternalService",
	HandlerType: (*InternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _InternalService_Login_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _InternalService_Profile_Handler,
		},
		{
			MethodName: "Branding",
			Handler:    _InternalService_Branding_Handler,
		},
		{
			MethodName: "GlobalSearch",
			Handler:    _InternalService_GlobalSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal.proto",
}
