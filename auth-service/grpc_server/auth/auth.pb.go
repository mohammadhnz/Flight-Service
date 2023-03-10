// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package authorization

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SignupData_GENDER_TYPES int32

const (
	SignupData_MEN   SignupData_GENDER_TYPES = 0
	SignupData_WOMEN SignupData_GENDER_TYPES = 1
	SignupData_SO_SO SignupData_GENDER_TYPES = 2
)

var SignupData_GENDER_TYPES_name = map[int32]string{
	0: "MEN",
	1: "WOMEN",
	2: "SO_SO",
}

var SignupData_GENDER_TYPES_value = map[string]int32{
	"MEN":   0,
	"WOMEN": 1,
	"SO_SO": 2,
}

func (x SignupData_GENDER_TYPES) String() string {
	return proto.EnumName(SignupData_GENDER_TYPES_name, int32(x))
}

func (SignupData_GENDER_TYPES) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0, 0}
}

type SignupResponse_SIGNUP_STATUS int32

const (
	SignupResponse_DUPLICATE_EMAIL_OR_PHONE_NUMBER_403 SignupResponse_SIGNUP_STATUS = 0
	SignupResponse_FAILED_TO_HASH_PASSWORD_400         SignupResponse_SIGNUP_STATUS = 1
	SignupResponse_FAILED_TO_CREATE_USER_400           SignupResponse_SIGNUP_STATUS = 2
	SignupResponse_OK_201                              SignupResponse_SIGNUP_STATUS = 3
)

var SignupResponse_SIGNUP_STATUS_name = map[int32]string{
	0: "DUPLICATE_EMAIL_OR_PHONE_NUMBER_403",
	1: "FAILED_TO_HASH_PASSWORD_400",
	2: "FAILED_TO_CREATE_USER_400",
	3: "OK_201",
}

var SignupResponse_SIGNUP_STATUS_value = map[string]int32{
	"DUPLICATE_EMAIL_OR_PHONE_NUMBER_403": 0,
	"FAILED_TO_HASH_PASSWORD_400":         1,
	"FAILED_TO_CREATE_USER_400":           2,
	"OK_201":                              3,
}

func (x SignupResponse_SIGNUP_STATUS) String() string {
	return proto.EnumName(SignupResponse_SIGNUP_STATUS_name, int32(x))
}

func (SignupResponse_SIGNUP_STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1, 0}
}

type FreshTokensResponse_SINGIN_STATUS int32

const (
	FreshTokensResponse_WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD FreshTokensResponse_SINGIN_STATUS = 0
	FreshTokensResponse_OK_200                                  FreshTokensResponse_SINGIN_STATUS = 1
	FreshTokensResponse_FAILED_TO_CREATE_NEW_TOKEN              FreshTokensResponse_SINGIN_STATUS = 2
)

var FreshTokensResponse_SINGIN_STATUS_name = map[int32]string{
	0: "WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD",
	1: "OK_200",
	2: "FAILED_TO_CREATE_NEW_TOKEN",
}

var FreshTokensResponse_SINGIN_STATUS_value = map[string]int32{
	"WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD": 0,
	"OK_200":                     1,
	"FAILED_TO_CREATE_NEW_TOKEN": 2,
}

func (x FreshTokensResponse_SINGIN_STATUS) String() string {
	return proto.EnumName(FreshTokensResponse_SINGIN_STATUS_name, int32(x))
}

func (FreshTokensResponse_SINGIN_STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3, 0}
}

type SignOutResponse_SIGN_OUT_STATUS int32

const (
	SignOutResponse_AUTHORIZATION_ERROR_401 SignOutResponse_SIGN_OUT_STATUS = 0
	SignOutResponse_FAILED_TO_SIGN_OUT      SignOutResponse_SIGN_OUT_STATUS = 1
	SignOutResponse_OK_201                  SignOutResponse_SIGN_OUT_STATUS = 2
)

var SignOutResponse_SIGN_OUT_STATUS_name = map[int32]string{
	0: "AUTHORIZATION_ERROR_401",
	1: "FAILED_TO_SIGN_OUT",
	2: "OK_201",
}

var SignOutResponse_SIGN_OUT_STATUS_value = map[string]int32{
	"AUTHORIZATION_ERROR_401": 0,
	"FAILED_TO_SIGN_OUT":      1,
	"OK_201":                  2,
}

func (x SignOutResponse_SIGN_OUT_STATUS) String() string {
	return proto.EnumName(SignOutResponse_SIGN_OUT_STATUS_name, int32(x))
}

func (SignOutResponse_SIGN_OUT_STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5, 0}
}

type UserInfoResponse_USERINFO_STATUS int32

const (
	UserInfoResponse_AUTHORIZATION_ERROR_401 UserInfoResponse_USERINFO_STATUS = 0
	UserInfoResponse_FAILED_TO_SIGN_OUT      UserInfoResponse_USERINFO_STATUS = 1
	UserInfoResponse_OK_200                  UserInfoResponse_USERINFO_STATUS = 2
)

var UserInfoResponse_USERINFO_STATUS_name = map[int32]string{
	0: "AUTHORIZATION_ERROR_401",
	1: "FAILED_TO_SIGN_OUT",
	2: "OK_200",
}

var UserInfoResponse_USERINFO_STATUS_value = map[string]int32{
	"AUTHORIZATION_ERROR_401": 0,
	"FAILED_TO_SIGN_OUT":      1,
	"OK_200":                  2,
}

func (x UserInfoResponse_USERINFO_STATUS) String() string {
	return proto.EnumName(UserInfoResponse_USERINFO_STATUS_name, int32(x))
}

func (UserInfoResponse_USERINFO_STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6, 0}
}

type SignupData struct {
	Email                string                  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string                  `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Gender               SignupData_GENDER_TYPES `protobuf:"varint,5,opt,name=gender,proto3,enum=authorization.SignupData_GENDER_TYPES" json:"gender,omitempty"`
	FirstName            string                  `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string                  `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Password             string                  `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SignupData) Reset()         { *m = SignupData{} }
func (m *SignupData) String() string { return proto.CompactTextString(m) }
func (*SignupData) ProtoMessage()    {}
func (*SignupData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *SignupData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupData.Unmarshal(m, b)
}
func (m *SignupData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupData.Marshal(b, m, deterministic)
}
func (m *SignupData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupData.Merge(m, src)
}
func (m *SignupData) XXX_Size() int {
	return xxx_messageInfo_SignupData.Size(m)
}
func (m *SignupData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupData.DiscardUnknown(m)
}

var xxx_messageInfo_SignupData proto.InternalMessageInfo

func (m *SignupData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupData) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SignupData) GetGender() SignupData_GENDER_TYPES {
	if m != nil {
		return m.Gender
	}
	return SignupData_MEN
}

func (m *SignupData) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *SignupData) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *SignupData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignupResponse struct {
	Status               SignupResponse_SIGNUP_STATUS `protobuf:"varint,7,opt,name=status,proto3,enum=authorization.SignupResponse_SIGNUP_STATUS" json:"status,omitempty"`
	UserId               int64                        `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SignupResponse) Reset()         { *m = SignupResponse{} }
func (m *SignupResponse) String() string { return proto.CompactTextString(m) }
func (*SignupResponse) ProtoMessage()    {}
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *SignupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupResponse.Unmarshal(m, b)
}
func (m *SignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupResponse.Marshal(b, m, deterministic)
}
func (m *SignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupResponse.Merge(m, src)
}
func (m *SignupResponse) XXX_Size() int {
	return xxx_messageInfo_SignupResponse.Size(m)
}
func (m *SignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignupResponse proto.InternalMessageInfo

func (m *SignupResponse) GetStatus() SignupResponse_SIGNUP_STATUS {
	if m != nil {
		return m.Status
	}
	return SignupResponse_DUPLICATE_EMAIL_OR_PHONE_NUMBER_403
}

func (m *SignupResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type SignInData struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInData) Reset()         { *m = SignInData{} }
func (m *SignInData) String() string { return proto.CompactTextString(m) }
func (*SignInData) ProtoMessage()    {}
func (*SignInData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *SignInData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInData.Unmarshal(m, b)
}
func (m *SignInData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInData.Marshal(b, m, deterministic)
}
func (m *SignInData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInData.Merge(m, src)
}
func (m *SignInData) XXX_Size() int {
	return xxx_messageInfo_SignInData.Size(m)
}
func (m *SignInData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInData.DiscardUnknown(m)
}

var xxx_messageInfo_SignInData proto.InternalMessageInfo

func (m *SignInData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignInData) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SignInData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type FreshTokensResponse struct {
	AccessToken          string                            `protobuf:"bytes,3,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken         string                            `protobuf:"bytes,4,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	Status               FreshTokensResponse_SINGIN_STATUS `protobuf:"varint,7,opt,name=status,proto3,enum=authorization.FreshTokensResponse_SINGIN_STATUS" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FreshTokensResponse) Reset()         { *m = FreshTokensResponse{} }
func (m *FreshTokensResponse) String() string { return proto.CompactTextString(m) }
func (*FreshTokensResponse) ProtoMessage()    {}
func (*FreshTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *FreshTokensResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreshTokensResponse.Unmarshal(m, b)
}
func (m *FreshTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreshTokensResponse.Marshal(b, m, deterministic)
}
func (m *FreshTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreshTokensResponse.Merge(m, src)
}
func (m *FreshTokensResponse) XXX_Size() int {
	return xxx_messageInfo_FreshTokensResponse.Size(m)
}
func (m *FreshTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FreshTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FreshTokensResponse proto.InternalMessageInfo

func (m *FreshTokensResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *FreshTokensResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *FreshTokensResponse) GetStatus() FreshTokensResponse_SINGIN_STATUS {
	if m != nil {
		return m.Status
	}
	return FreshTokensResponse_WRONG_EMAIL_OR_PHONE_NUMBER_OR_PASSWORD
}

type TokenData struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenData) Reset()         { *m = TokenData{} }
func (m *TokenData) String() string { return proto.CompactTextString(m) }
func (*TokenData) ProtoMessage()    {}
func (*TokenData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *TokenData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenData.Unmarshal(m, b)
}
func (m *TokenData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenData.Marshal(b, m, deterministic)
}
func (m *TokenData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenData.Merge(m, src)
}
func (m *TokenData) XXX_Size() int {
	return xxx_messageInfo_TokenData.Size(m)
}
func (m *TokenData) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenData.DiscardUnknown(m)
}

var xxx_messageInfo_TokenData proto.InternalMessageInfo

func (m *TokenData) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *TokenData) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type SignOutResponse struct {
	Status               SignOutResponse_SIGN_OUT_STATUS `protobuf:"varint,4,opt,name=status,proto3,enum=authorization.SignOutResponse_SIGN_OUT_STATUS" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SignOutResponse) Reset()         { *m = SignOutResponse{} }
func (m *SignOutResponse) String() string { return proto.CompactTextString(m) }
func (*SignOutResponse) ProtoMessage()    {}
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *SignOutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutResponse.Unmarshal(m, b)
}
func (m *SignOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutResponse.Marshal(b, m, deterministic)
}
func (m *SignOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutResponse.Merge(m, src)
}
func (m *SignOutResponse) XXX_Size() int {
	return xxx_messageInfo_SignOutResponse.Size(m)
}
func (m *SignOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutResponse proto.InternalMessageInfo

func (m *SignOutResponse) GetStatus() SignOutResponse_SIGN_OUT_STATUS {
	if m != nil {
		return m.Status
	}
	return SignOutResponse_AUTHORIZATION_ERROR_401
}

type UserInfoResponse struct {
	User                 *UserData                        `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Status               UserInfoResponse_USERINFO_STATUS `protobuf:"varint,9,opt,name=status,proto3,enum=authorization.UserInfoResponse_USERINFO_STATUS" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetUser() *UserData {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserInfoResponse) GetStatus() UserInfoResponse_USERINFO_STATUS {
	if m != nil {
		return m.Status
	}
	return UserInfoResponse_AUTHORIZATION_ERROR_401
}

type UserData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Gender               string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	FirstName            string   `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserData) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UserData) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserData) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserData) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterEnum("authorization.SignupData_GENDER_TYPES", SignupData_GENDER_TYPES_name, SignupData_GENDER_TYPES_value)
	proto.RegisterEnum("authorization.SignupResponse_SIGNUP_STATUS", SignupResponse_SIGNUP_STATUS_name, SignupResponse_SIGNUP_STATUS_value)
	proto.RegisterEnum("authorization.FreshTokensResponse_SINGIN_STATUS", FreshTokensResponse_SINGIN_STATUS_name, FreshTokensResponse_SINGIN_STATUS_value)
	proto.RegisterEnum("authorization.SignOutResponse_SIGN_OUT_STATUS", SignOutResponse_SIGN_OUT_STATUS_name, SignOutResponse_SIGN_OUT_STATUS_value)
	proto.RegisterEnum("authorization.UserInfoResponse_USERINFO_STATUS", UserInfoResponse_USERINFO_STATUS_name, UserInfoResponse_USERINFO_STATUS_value)
	proto.RegisterType((*SignupData)(nil), "authorization.SignupData")
	proto.RegisterType((*SignupResponse)(nil), "authorization.SignupResponse")
	proto.RegisterType((*SignInData)(nil), "authorization.SignInData")
	proto.RegisterType((*FreshTokensResponse)(nil), "authorization.FreshTokensResponse")
	proto.RegisterType((*TokenData)(nil), "authorization.TokenData")
	proto.RegisterType((*SignOutResponse)(nil), "authorization.SignOutResponse")
	proto.RegisterType((*UserInfoResponse)(nil), "authorization.UserInfoResponse")
	proto.RegisterType((*UserData)(nil), "authorization.UserData")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xfa, 0x46,
	0x10, 0xc7, 0x26, 0x18, 0x98, 0xe4, 0x4f, 0xac, 0x6d, 0x95, 0x10, 0xa2, 0x7c, 0xd4, 0x95, 0x9a,
	0x48, 0x51, 0x5d, 0x42, 0x72, 0xae, 0xe4, 0x06, 0x03, 0x56, 0x12, 0x2f, 0x5d, 0xdb, 0x42, 0xed,
	0x65, 0xe5, 0x04, 0x27, 0x58, 0x0d, 0x36, 0xb2, 0x8d, 0x2a, 0xf5, 0xdc, 0x63, 0xa5, 0x3e, 0x46,
	0x4f, 0x7d, 0xa3, 0xbe, 0x46, 0x6f, 0x3d, 0x54, 0xbb, 0x06, 0x03, 0x86, 0x90, 0x4a, 0xbd, 0x79,
	0xe7, 0x7b, 0x7e, 0xf3, 0x9b, 0x31, 0x80, 0x3b, 0x4d, 0x46, 0xea, 0x24, 0x0a, 0x93, 0x10, 0x7d,
	0x62, 0xdf, 0x61, 0xe4, 0xff, 0xe2, 0x26, 0x7e, 0x18, 0x28, 0xbf, 0x89, 0x00, 0x96, 0xff, 0x1a,
	0x4c, 0x27, 0x6d, 0x37, 0x71, 0xd1, 0xe7, 0x50, 0xf2, 0xc6, 0xae, 0xff, 0x56, 0x2f, 0x9e, 0x0b,
	0x97, 0x55, 0x92, 0x3e, 0xd0, 0x17, 0xb0, 0x37, 0x19, 0x85, 0x81, 0x47, 0x83, 0xe9, 0xf8, 0xc9,
	0x8b, 0xea, 0x3b, 0x5c, 0xb9, 0xcb, 0x65, 0x26, 0x17, 0xa1, 0x6f, 0x41, 0x7a, 0xf5, 0x82, 0xa1,
	0x17, 0xd5, 0x4b, 0xe7, 0xc2, 0x65, 0xad, 0xf5, 0x95, 0xba, 0x92, 0x47, 0x5d, 0xe4, 0x50, 0xbb,
	0xba, 0xd9, 0xd6, 0x09, 0xb5, 0x7f, 0xe8, 0xeb, 0x16, 0x99, 0x79, 0xa1, 0x13, 0x80, 0x17, 0x3f,
	0x8a, 0x13, 0x1a, 0xb8, 0x63, 0xaf, 0x2e, 0xf1, 0x04, 0x55, 0x2e, 0x31, 0xdd, 0xb1, 0x87, 0x8e,
	0xa1, 0xfa, 0xe6, 0xce, 0xb5, 0x65, 0xae, 0xad, 0x30, 0x01, 0x57, 0x36, 0xa0, 0x32, 0x71, 0xe3,
	0xf8, 0xe7, 0x30, 0x1a, 0xd6, 0x2b, 0xa9, 0x6e, 0xfe, 0x56, 0xbe, 0x86, 0xbd, 0xe5, 0x7c, 0xa8,
	0x0c, 0xc5, 0x47, 0xdd, 0x94, 0x0b, 0xa8, 0x0a, 0xa5, 0x01, 0x66, 0x9f, 0x02, 0xfb, 0xb4, 0x30,
	0xb5, 0xb0, 0x2c, 0x2a, 0x7f, 0x0b, 0x50, 0x4b, 0x4b, 0x25, 0x5e, 0x3c, 0x09, 0x83, 0xd8, 0x43,
	0x77, 0x20, 0xc5, 0x89, 0x9b, 0x4c, 0x63, 0x9e, 0xb7, 0xd6, 0xba, 0xda, 0xd8, 0xd9, 0xdc, 0x5c,
	0xb5, 0x8c, 0xae, 0xe9, 0xf4, 0xa9, 0x65, 0x6b, 0xb6, 0x63, 0x91, 0x99, 0x2b, 0x3a, 0x84, 0xf2,
	0x34, 0xf6, 0x22, 0xea, 0xa7, 0x15, 0x16, 0x89, 0xc4, 0x9e, 0xc6, 0x50, 0xf9, 0x55, 0x80, 0x4f,
	0x2b, 0x2e, 0xe8, 0x02, 0xbe, 0x6c, 0x3b, 0xfd, 0x07, 0xe3, 0x4e, 0xb3, 0x75, 0xaa, 0x3f, 0x6a,
	0xc6, 0x03, 0xc5, 0x84, 0xf6, 0x7b, 0xd8, 0xd4, 0xa9, 0xe9, 0x3c, 0x7e, 0xa7, 0x13, 0x7a, 0xdb,
	0xbc, 0x91, 0x0b, 0xe8, 0x0c, 0x8e, 0x3b, 0x9a, 0xf1, 0xa0, 0xb7, 0xa9, 0x8d, 0x69, 0x4f, 0xb3,
	0x7a, 0xb4, 0xaf, 0x59, 0xd6, 0x00, 0x93, 0x36, 0xbd, 0x6d, 0x36, 0x65, 0x01, 0x9d, 0xc0, 0xd1,
	0xc2, 0xe0, 0x8e, 0xe8, 0x2c, 0xa0, 0x63, 0x71, 0xff, 0xa6, 0x2c, 0x22, 0x00, 0x09, 0xdf, 0xd3,
	0x56, 0xf3, 0x5a, 0x2e, 0x2a, 0x6e, 0xca, 0x02, 0x23, 0x58, 0x65, 0x81, 0xb0, 0x8d, 0x05, 0xe2,
	0x3a, 0x0b, 0x96, 0x27, 0x51, 0xcc, 0x4d, 0xe2, 0x77, 0x11, 0x3e, 0xeb, 0x44, 0x5e, 0x3c, 0xb2,
	0xc3, 0x9f, 0xbc, 0x20, 0xce, 0xf0, 0x3d, 0x87, 0x5d, 0xf7, 0xf9, 0xd9, 0x8b, 0x63, 0x2e, 0x9f,
	0xb9, 0x2d, 0x8b, 0x90, 0x02, 0x7b, 0x91, 0xf7, 0x92, 0xb9, 0xce, 0xe8, 0xb7, 0x22, 0x43, 0xbd,
	0xdc, 0x94, 0x9a, 0xb9, 0x29, 0x6d, 0xc8, 0xac, 0x5a, 0x86, 0xd9, 0x35, 0xcc, 0xdc, 0xa8, 0x94,
	0x11, 0x1b, 0xc8, 0x92, 0x02, 0x5d, 0xc1, 0xc5, 0x80, 0x60, 0xb3, 0xfb, 0xce, 0x30, 0xd8, 0x7b,
	0x86, 0xbb, 0x5c, 0xc8, 0x40, 0x65, 0xf8, 0x9f, 0x42, 0x63, 0x0d, 0x7f, 0x53, 0x1f, 0x50, 0x1b,
	0xdf, 0xeb, 0xa6, 0x2c, 0x2a, 0xdf, 0x43, 0x95, 0x57, 0xc4, 0x31, 0xcf, 0xc1, 0x20, 0x7c, 0x0c,
	0x83, 0xb8, 0x0e, 0x83, 0xf2, 0xa7, 0x00, 0xfb, 0x6c, 0x90, 0x78, 0x9a, 0x64, 0x00, 0x77, 0x32,
	0x68, 0x76, 0x38, 0x34, 0xea, 0x06, 0x02, 0x2f, 0xd9, 0x73, 0x06, 0x53, 0xec, 0xd8, 0x79, 0x60,
	0x08, 0xec, 0xe7, 0x54, 0xe8, 0x18, 0x0e, 0x35, 0xc7, 0xee, 0x61, 0x62, 0xfc, 0xa8, 0xd9, 0x06,
	0x36, 0xa9, 0x4e, 0x08, 0x66, 0xfc, 0xba, 0x96, 0x0b, 0xe8, 0x00, 0xd0, 0xa2, 0xfd, 0xb9, 0xa7,
	0x2c, 0x2c, 0xf1, 0x4e, 0x54, 0xfe, 0x12, 0x40, 0x76, 0xd8, 0x26, 0x04, 0x2f, 0x61, 0x56, 0xf0,
	0x15, 0xec, 0xb0, 0xed, 0xe0, 0x0d, 0xee, 0xb6, 0x0e, 0x73, 0xe5, 0x32, 0x73, 0x86, 0x18, 0xe1,
	0x46, 0xa8, 0x9b, 0x75, 0x57, 0xe5, 0xdd, 0x7d, 0xb3, 0xc1, 0x7c, 0x39, 0xba, 0xca, 0x36, 0xc0,
	0x30, 0x3b, 0x78, 0x43, 0x7b, 0x39, 0xd5, 0xff, 0x6b, 0xaf, 0x29, 0x8b, 0xca, 0x1f, 0x02, 0x54,
	0xe6, 0xf5, 0xa2, 0x1a, 0x88, 0xfe, 0x90, 0x0f, 0xb6, 0x48, 0x44, 0x7f, 0xb8, 0xd8, 0x32, 0x71,
	0xdb, 0x96, 0x15, 0xd7, 0xb7, 0xec, 0x20, 0xbb, 0xb5, 0xe9, 0x26, 0x6c, 0xbe, 0xa1, 0xa5, 0xad,
	0x37, 0x54, 0x5a, 0xbd, 0xa1, 0xad, 0x7f, 0x44, 0xa8, 0x69, 0xd3, 0x64, 0xe4, 0x05, 0x89, 0xff,
	0xcc, 0x91, 0x43, 0x6d, 0x90, 0xd2, 0xdb, 0x86, 0x8e, 0xde, 0x3d, 0xe6, 0x8d, 0x93, 0xad, 0xd7,
	0x50, 0x29, 0x20, 0x23, 0x8d, 0x62, 0x04, 0x1b, 0xa3, 0xa4, 0x07, 0xa7, 0xa1, 0x7c, 0xbc, 0xad,
	0x4a, 0x01, 0xe9, 0x50, 0x9e, 0x71, 0x15, 0xd5, 0x73, 0x0e, 0xd9, 0x1e, 0x35, 0x4e, 0xb7, 0xb3,
	0x5b, 0x29, 0xa0, 0x6e, 0x3a, 0x13, 0x46, 0x8a, 0x2d, 0x71, 0xce, 0x3e, 0xe0, 0x11, 0x6f, 0xad,
	0x4c, 0xd2, 0xe5, 0xdb, 0x12, 0xe7, 0x3f, 0xb5, 0xf6, 0x24, 0xf1, 0x9f, 0xf3, 0xcd, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc3, 0x25, 0xdb, 0xe6, 0xaa, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	Signup(ctx context.Context, in *SignupData, opts ...grpc.CallOption) (*SignupResponse, error)
	SignIn(ctx context.Context, in *SignInData, opts ...grpc.CallOption) (*FreshTokensResponse, error)
	SignOut(ctx context.Context, in *TokenData, opts ...grpc.CallOption) (*SignOutResponse, error)
	UserInfo(ctx context.Context, in *TokenData, opts ...grpc.CallOption) (*UserInfoResponse, error)
	Refresh(ctx context.Context, in *TokenData, opts ...grpc.CallOption) (*FreshTokensResponse, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Signup(ctx context.Context, in *SignupData, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/authorization.Authentication/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) SignIn(ctx context.Context, in *SignInData, opts ...grpc.CallOption) (*FreshTokensResponse, error) {
	out := new(FreshTokensResponse)
	err := c.cc.Invoke(ctx, "/authorization.Authentication/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) SignOut(ctx context.Context, in *TokenData, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/authorization.Authentication/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) UserInfo(ctx context.Context, in *TokenData, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/authorization.Authentication/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Refresh(ctx context.Context, in *TokenData, opts ...grpc.CallOption) (*FreshTokensResponse, error) {
	out := new(FreshTokensResponse)
	err := c.cc.Invoke(ctx, "/authorization.Authentication/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	Signup(context.Context, *SignupData) (*SignupResponse, error)
	SignIn(context.Context, *SignInData) (*FreshTokensResponse, error)
	SignOut(context.Context, *TokenData) (*SignOutResponse, error)
	UserInfo(context.Context, *TokenData) (*UserInfoResponse, error)
	Refresh(context.Context, *TokenData) (*FreshTokensResponse, error)
}

// UnimplementedAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (*UnimplementedAuthenticationServer) Signup(ctx context.Context, req *SignupData) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (*UnimplementedAuthenticationServer) SignIn(ctx context.Context, req *SignInData) (*FreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAuthenticationServer) SignOut(ctx context.Context, req *TokenData) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (*UnimplementedAuthenticationServer) UserInfo(ctx context.Context, req *TokenData) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (*UnimplementedAuthenticationServer) Refresh(ctx context.Context, req *TokenData) (*FreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authentication/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Signup(ctx, req.(*SignupData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authentication/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).SignIn(ctx, req.(*SignInData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authentication/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).SignOut(ctx, req.(*TokenData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authentication/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).UserInfo(ctx, req.(*TokenData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authentication/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Refresh(ctx, req.(*TokenData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _Authentication_Signup_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Authentication_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _Authentication_SignOut_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Authentication_UserInfo_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Authentication_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
