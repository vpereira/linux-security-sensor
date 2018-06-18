// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fingerprint.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The fingerprinting methods the fingerprinter can be asked to perform.
// If none is given, all are applied.
type FingerprintTuple_Type int32

const (
	FingerprintTuple_FPT_GENERIC FingerprintTuple_Type = 0
	FingerprintTuple_FPT_PE_COFF FingerprintTuple_Type = 1
)

var FingerprintTuple_Type_name = map[int32]string{
	0: "FPT_GENERIC",
	1: "FPT_PE_COFF",
}
var FingerprintTuple_Type_value = map[string]int32{
	"FPT_GENERIC": 0,
	"FPT_PE_COFF": 1,
}

func (x FingerprintTuple_Type) String() string {
	return proto.EnumName(FingerprintTuple_Type_name, int32(x))
}
func (FingerprintTuple_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{2, 0}
}

// The hash functions that a fingerprinting method may employ.
// If none is given, all applicable ones are used.
type FingerprintTuple_HashType int32

const (
	FingerprintTuple_MD5    FingerprintTuple_HashType = 0
	FingerprintTuple_SHA1   FingerprintTuple_HashType = 1
	FingerprintTuple_SHA256 FingerprintTuple_HashType = 2
)

var FingerprintTuple_HashType_name = map[int32]string{
	0: "MD5",
	1: "SHA1",
	2: "SHA256",
}
var FingerprintTuple_HashType_value = map[string]int32{
	"MD5":    0,
	"SHA1":   1,
	"SHA256": 2,
}

func (x FingerprintTuple_HashType) String() string {
	return proto.EnumName(FingerprintTuple_HashType_name, int32(x))
}
func (FingerprintTuple_HashType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{2, 1}
}

type Hash struct {
	Sha256               []byte                    `protobuf:"bytes,1,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Sha1                 []byte                    `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
	Md5                  []byte                    `protobuf:"bytes,3,opt,name=md5,proto3" json:"md5,omitempty"`
	PecoffSha1           []byte                    `protobuf:"bytes,4,opt,name=pecoff_sha1,json=pecoffSha1,proto3" json:"pecoff_sha1,omitempty"`
	PecoffMd5            []byte                    `protobuf:"bytes,5,opt,name=pecoff_md5,json=pecoffMd5,proto3" json:"pecoff_md5,omitempty"`
	PecoffSha256         []byte                    `protobuf:"bytes,7,opt,name=pecoff_sha256,json=pecoffSha256,proto3" json:"pecoff_sha256,omitempty"`
	SignedData           []*AuthenticodeSignedData `protobuf:"bytes,6,rep,name=signed_data,json=signedData,proto3" json:"signed_data,omitempty"`
	NumBytes             uint64                    `protobuf:"varint,8,opt,name=num_bytes,json=numBytes,proto3" json:"num_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{0}
}
func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (dst *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(dst, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetSha256() []byte {
	if m != nil {
		return m.Sha256
	}
	return nil
}

func (m *Hash) GetSha1() []byte {
	if m != nil {
		return m.Sha1
	}
	return nil
}

func (m *Hash) GetMd5() []byte {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *Hash) GetPecoffSha1() []byte {
	if m != nil {
		return m.PecoffSha1
	}
	return nil
}

func (m *Hash) GetPecoffMd5() []byte {
	if m != nil {
		return m.PecoffMd5
	}
	return nil
}

func (m *Hash) GetPecoffSha256() []byte {
	if m != nil {
		return m.PecoffSha256
	}
	return nil
}

func (m *Hash) GetSignedData() []*AuthenticodeSignedData {
	if m != nil {
		return m.SignedData
	}
	return nil
}

func (m *Hash) GetNumBytes() uint64 {
	if m != nil {
		return m.NumBytes
	}
	return 0
}

type AuthenticodeSignedData struct {
	Revision             uint64   `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	CertType             uint64   `protobuf:"varint,2,opt,name=cert_type,json=certType,proto3" json:"cert_type,omitempty"`
	Certificate          []byte   `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticodeSignedData) Reset()         { *m = AuthenticodeSignedData{} }
func (m *AuthenticodeSignedData) String() string { return proto.CompactTextString(m) }
func (*AuthenticodeSignedData) ProtoMessage()    {}
func (*AuthenticodeSignedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{1}
}
func (m *AuthenticodeSignedData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticodeSignedData.Unmarshal(m, b)
}
func (m *AuthenticodeSignedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticodeSignedData.Marshal(b, m, deterministic)
}
func (dst *AuthenticodeSignedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticodeSignedData.Merge(dst, src)
}
func (m *AuthenticodeSignedData) XXX_Size() int {
	return xxx_messageInfo_AuthenticodeSignedData.Size(m)
}
func (m *AuthenticodeSignedData) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticodeSignedData.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticodeSignedData proto.InternalMessageInfo

func (m *AuthenticodeSignedData) GetRevision() uint64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *AuthenticodeSignedData) GetCertType() uint64 {
	if m != nil {
		return m.CertType
	}
	return 0
}

func (m *AuthenticodeSignedData) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type FingerprintTuple struct {
	FpType               FingerprintTuple_Type       `protobuf:"varint,1,opt,name=fp_type,json=fpType,proto3,enum=proto.FingerprintTuple_Type" json:"fp_type,omitempty"`
	Hashers              []FingerprintTuple_HashType `protobuf:"varint,2,rep,packed,name=hashers,proto3,enum=proto.FingerprintTuple_HashType" json:"hashers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FingerprintTuple) Reset()         { *m = FingerprintTuple{} }
func (m *FingerprintTuple) String() string { return proto.CompactTextString(m) }
func (*FingerprintTuple) ProtoMessage()    {}
func (*FingerprintTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{2}
}
func (m *FingerprintTuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FingerprintTuple.Unmarshal(m, b)
}
func (m *FingerprintTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FingerprintTuple.Marshal(b, m, deterministic)
}
func (dst *FingerprintTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FingerprintTuple.Merge(dst, src)
}
func (m *FingerprintTuple) XXX_Size() int {
	return xxx_messageInfo_FingerprintTuple.Size(m)
}
func (m *FingerprintTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_FingerprintTuple.DiscardUnknown(m)
}

var xxx_messageInfo_FingerprintTuple proto.InternalMessageInfo

func (m *FingerprintTuple) GetFpType() FingerprintTuple_Type {
	if m != nil {
		return m.FpType
	}
	return FingerprintTuple_FPT_GENERIC
}

func (m *FingerprintTuple) GetHashers() []FingerprintTuple_HashType {
	if m != nil {
		return m.Hashers
	}
	return nil
}

// Request fingerprints for a file.
type FingerprintRequest struct {
	Pathspec             *PathSpec           `protobuf:"bytes,1,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	Tuples               []*FingerprintTuple `protobuf:"bytes,2,rep,name=tuples,proto3" json:"tuples,omitempty"`
	MaxFilesize          uint64              `protobuf:"varint,3,opt,name=max_filesize,json=maxFilesize,proto3" json:"max_filesize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FingerprintRequest) Reset()         { *m = FingerprintRequest{} }
func (m *FingerprintRequest) String() string { return proto.CompactTextString(m) }
func (*FingerprintRequest) ProtoMessage()    {}
func (*FingerprintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{3}
}
func (m *FingerprintRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FingerprintRequest.Unmarshal(m, b)
}
func (m *FingerprintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FingerprintRequest.Marshal(b, m, deterministic)
}
func (dst *FingerprintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FingerprintRequest.Merge(dst, src)
}
func (m *FingerprintRequest) XXX_Size() int {
	return xxx_messageInfo_FingerprintRequest.Size(m)
}
func (m *FingerprintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FingerprintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FingerprintRequest proto.InternalMessageInfo

func (m *FingerprintRequest) GetPathspec() *PathSpec {
	if m != nil {
		return m.Pathspec
	}
	return nil
}

func (m *FingerprintRequest) GetTuples() []*FingerprintTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

func (m *FingerprintRequest) GetMaxFilesize() uint64 {
	if m != nil {
		return m.MaxFilesize
	}
	return 0
}

// Response data for file hashes and signature blobs.
type FingerprintResponse struct {
	MatchingTypes        []FingerprintTuple_Type `protobuf:"varint,1,rep,packed,name=matching_types,json=matchingTypes,proto3,enum=proto.FingerprintTuple_Type" json:"matching_types,omitempty"`
	Pathspec             *PathSpec               `protobuf:"bytes,3,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	Hash                 *Hash                   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	BytesRead            uint64                  `protobuf:"varint,5,opt,name=bytes_read,json=bytesRead,proto3" json:"bytes_read,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FingerprintResponse) Reset()         { *m = FingerprintResponse{} }
func (m *FingerprintResponse) String() string { return proto.CompactTextString(m) }
func (*FingerprintResponse) ProtoMessage()    {}
func (*FingerprintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{4}
}
func (m *FingerprintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FingerprintResponse.Unmarshal(m, b)
}
func (m *FingerprintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FingerprintResponse.Marshal(b, m, deterministic)
}
func (dst *FingerprintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FingerprintResponse.Merge(dst, src)
}
func (m *FingerprintResponse) XXX_Size() int {
	return xxx_messageInfo_FingerprintResponse.Size(m)
}
func (m *FingerprintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FingerprintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FingerprintResponse proto.InternalMessageInfo

func (m *FingerprintResponse) GetMatchingTypes() []FingerprintTuple_Type {
	if m != nil {
		return m.MatchingTypes
	}
	return nil
}

func (m *FingerprintResponse) GetPathspec() *PathSpec {
	if m != nil {
		return m.Pathspec
	}
	return nil
}

func (m *FingerprintResponse) GetHash() *Hash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *FingerprintResponse) GetBytesRead() uint64 {
	if m != nil {
		return m.BytesRead
	}
	return 0
}

type BufferReference struct {
	Offset               uint64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Length               uint64    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Callback             string    `protobuf:"bytes,3,opt,name=callback,proto3" json:"callback,omitempty"`
	Data                 []byte    `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Pathspec             *PathSpec `protobuf:"bytes,6,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BufferReference) Reset()         { *m = BufferReference{} }
func (m *BufferReference) String() string { return proto.CompactTextString(m) }
func (*BufferReference) ProtoMessage()    {}
func (*BufferReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_fingerprint_627c0050b1e96b56, []int{5}
}
func (m *BufferReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferReference.Unmarshal(m, b)
}
func (m *BufferReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferReference.Marshal(b, m, deterministic)
}
func (dst *BufferReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferReference.Merge(dst, src)
}
func (m *BufferReference) XXX_Size() int {
	return xxx_messageInfo_BufferReference.Size(m)
}
func (m *BufferReference) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferReference.DiscardUnknown(m)
}

var xxx_messageInfo_BufferReference proto.InternalMessageInfo

func (m *BufferReference) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BufferReference) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *BufferReference) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func (m *BufferReference) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BufferReference) GetPathspec() *PathSpec {
	if m != nil {
		return m.Pathspec
	}
	return nil
}

func init() {
	proto.RegisterType((*Hash)(nil), "proto.Hash")
	proto.RegisterType((*AuthenticodeSignedData)(nil), "proto.AuthenticodeSignedData")
	proto.RegisterType((*FingerprintTuple)(nil), "proto.FingerprintTuple")
	proto.RegisterType((*FingerprintRequest)(nil), "proto.FingerprintRequest")
	proto.RegisterType((*FingerprintResponse)(nil), "proto.FingerprintResponse")
	proto.RegisterType((*BufferReference)(nil), "proto.BufferReference")
	proto.RegisterEnum("proto.FingerprintTuple_Type", FingerprintTuple_Type_name, FingerprintTuple_Type_value)
	proto.RegisterEnum("proto.FingerprintTuple_HashType", FingerprintTuple_HashType_name, FingerprintTuple_HashType_value)
}

func init() { proto.RegisterFile("fingerprint.proto", fileDescriptor_fingerprint_627c0050b1e96b56) }

var fileDescriptor_fingerprint_627c0050b1e96b56 = []byte{
	// 915 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0x0e, 0x2d, 0x46, 0x96, 0x87, 0xfe, 0xd0, 0xbb, 0x2f, 0x90, 0x08, 0x6e, 0x83, 0x6c, 0x55,
	0xa4, 0x50, 0x90, 0x96, 0x42, 0x54, 0x28, 0x07, 0x17, 0x29, 0x62, 0xd9, 0x56, 0x1d, 0xb4, 0x4e,
	0x5c, 0x4a, 0x77, 0x61, 0x45, 0x0e, 0xc5, 0x45, 0xc5, 0x8f, 0x72, 0x97, 0xb6, 0xd5, 0x73, 0x7f,
	0x47, 0x7f, 0x4b, 0x0f, 0xfd, 0x15, 0x3d, 0xb6, 0xf7, 0xfe, 0x82, 0x1e, 0x8a, 0x1d, 0x52, 0x8a,
	0x94, 0x3a, 0x4d, 0x4e, 0xe4, 0xcc, 0xce, 0x3c, 0xcf, 0x7c, 0x3c, 0xbb, 0xf0, 0xbf, 0x50, 0x26,
	0x33, 0xcc, 0xb3, 0x5c, 0x26, 0xda, 0xcd, 0xf2, 0x54, 0xa7, 0xec, 0x2e, 0x7d, 0x0e, 0xf7, 0x84,
	0xaf, 0x65, 0x9a, 0xa8, 0xd2, 0x7b, 0x78, 0x74, 0x7d, 0x7d, 0xed, 0x5e, 0xe1, 0x3c, 0xf5, 0x65,
	0x80, 0x37, 0xae, 0x9f, 0xc6, 0xdd, 0x59, 0x3a, 0x17, 0xc9, 0xac, 0x5b, 0x3a, 0x73, 0x91, 0xe9,
	0x34, 0xef, 0x52, 0x70, 0x57, 0x61, 0x2c, 0x12, 0x2d, 0xfd, 0x32, 0xb7, 0xfd, 0x73, 0x1d, 0xec,
	0x73, 0xa1, 0x22, 0x36, 0x84, 0xba, 0x8a, 0x44, 0xaf, 0xff, 0xac, 0x65, 0x71, 0xab, 0xb3, 0x3b,
	0x70, 0xff, 0xf8, 0xfb, 0xcf, 0xdf, 0xac, 0x0e, 0x80, 0x39, 0x3d, 0x95, 0x33, 0x54, 0x9a, 0x1d,
	0x8e, 0xce, 0x8f, 0x7b, 0xfd, 0x67, 0x7c, 0x2a, 0x13, 0x91, 0x2f, 0x78, 0x24, 0x54, 0xc4, 0x03,
	0x3a, 0x72, 0xbd, 0x2a, 0x9b, 0xbd, 0x00, 0x5b, 0x45, 0xe2, 0x69, 0x6b, 0x8b, 0x50, 0x3e, 0x27,
	0x94, 0xcf, 0x36, 0x50, 0x5a, 0xa3, 0xf3, 0xe3, 0xa7, 0xb7, 0x62, 0x50, 0x26, 0x7b, 0x0e, 0xb5,
	0x38, 0xe8, 0xb7, 0x6a, 0x04, 0xf0, 0x84, 0x00, 0x1e, 0x6d, 0x00, 0xdc, 0xbf, 0x38, 0xed, 0xdf,
	0x9a, 0x6f, 0xf2, 0x58, 0x06, 0x4e, 0x86, 0x7e, 0x1a, 0x86, 0x13, 0xaa, 0xc3, 0x26, 0x98, 0xd7,
	0x04, 0xf3, 0x72, 0x03, 0xe6, 0xab, 0xe3, 0x42, 0x47, 0x68, 0x66, 0x91, 0x06, 0xc8, 0xa9, 0xa8,
	0x35, 0x34, 0x1e, 0xa6, 0x39, 0x57, 0x72, 0x96, 0x60, 0xc0, 0x73, 0x9c, 0xc9, 0x34, 0xe1, 0x69,
	0xc8, 0x2f, 0xcf, 0x78, 0x28, 0xe7, 0xe8, 0x7a, 0x50, 0x72, 0x8c, 0x4c, 0xc1, 0x31, 0x54, 0xd6,
	0xc4, 0xd4, 0x7d, 0x97, 0x08, 0x5f, 0x11, 0xe1, 0xf9, 0x06, 0xe1, 0xd1, 0x06, 0xa1, 0x69, 0xe2,
	0xc3, 0xf9, 0x76, 0x4a, 0x86, 0x8b, 0xa0, 0xcf, 0xae, 0x60, 0xef, 0x4d, 0x83, 0x66, 0x61, 0xdb,
	0xc4, 0xf8, 0x3d, 0x31, 0x7e, 0xbb, 0xc1, 0xf8, 0xfc, 0xed, 0x16, 0xcd, 0xf6, 0x3e, 0x9c, 0x74,
	0x77, 0xd5, 0xa4, 0xd9, 0x6c, 0x01, 0x4e, 0x19, 0x3a, 0x09, 0x84, 0x16, 0xad, 0x3a, 0xaf, 0x75,
	0x9c, 0xde, 0x83, 0x52, 0x47, 0xee, 0x3a, 0xd1, 0x88, 0xa2, 0x4e, 0x85, 0x16, 0x83, 0x3e, 0x15,
	0xd5, 0x65, 0x5f, 0x94, 0x3e, 0x6e, 0x32, 0xf9, 0x75, 0x24, 0xfd, 0x88, 0xc7, 0x62, 0xc1, 0xa7,
	0xc8, 0xb3, 0x1c, 0x15, 0x26, 0x9a, 0xcb, 0x64, 0xc9, 0xab, 0x5c, 0x0f, 0xd4, 0x0a, 0x82, 0xf9,
	0xb0, 0x93, 0x14, 0xf1, 0x64, 0xba, 0xd0, 0xa8, 0x5a, 0x0d, 0x6e, 0x75, 0xec, 0xc1, 0x90, 0x50,
	0x5f, 0xb0, 0xaf, 0xc7, 0x11, 0xf2, 0xa4, 0x88, 0xa7, 0x98, 0x9b, 0xba, 0x29, 0x84, 0x17, 0x0a,
	0x03, 0xae, 0x53, 0xee, 0xe7, 0x28, 0x34, 0x72, 0x1d, 0x21, 0x75, 0x8c, 0xca, 0x50, 0xe8, 0x48,
	0x2a, 0x1e, 0xa3, 0x52, 0x62, 0x86, 0xae, 0xd7, 0x48, 0x8a, 0x78, 0x60, 0x92, 0xda, 0x0a, 0xee,
	0xdd, 0xde, 0x01, 0x3b, 0x84, 0x46, 0x8e, 0x57, 0x52, 0xc9, 0x34, 0xa1, 0x9b, 0x61, 0x7b, 0x2b,
	0x9b, 0x7d, 0x04, 0x3b, 0x3e, 0xe6, 0x7a, 0xa2, 0x17, 0x19, 0x92, 0xe0, 0x6d, 0xaf, 0x61, 0x1c,
	0xe3, 0x45, 0x86, 0x8c, 0x83, 0x63, 0xfe, 0x65, 0x28, 0x7d, 0xa1, 0xb1, 0x94, 0xb3, 0xb7, 0xee,
	0x6a, 0xff, 0x6e, 0x41, 0x73, 0xf8, 0xe6, 0x8e, 0x8f, 0x8b, 0x6c, 0x8e, 0xac, 0x0f, 0xdb, 0x61,
	0x56, 0x22, 0x1a, 0xba, 0xfd, 0xde, 0xc7, 0xd5, 0x84, 0xdf, 0x8e, 0x74, 0x0d, 0x8b, 0x57, 0x0f,
	0x33, 0x62, 0x3b, 0x82, 0x6d, 0xea, 0x32, 0x57, 0xad, 0x2d, 0x5e, 0xeb, 0xec, 0xf7, 0xf8, 0xbb,
	0xd2, 0x8c, 0x3c, 0x28, 0x75, 0x99, 0xd0, 0xee, 0x80, 0x4d, 0x18, 0x07, 0xe0, 0x0c, 0x2f, 0xc7,
	0x93, 0x6f, 0xce, 0x5e, 0x9d, 0x79, 0x2f, 0x4f, 0x9a, 0x77, 0x96, 0x8e, 0xcb, 0xb3, 0xc9, 0xc9,
	0xeb, 0xe1, 0xb0, 0x69, 0xb5, 0x1f, 0x43, 0x63, 0x99, 0xce, 0xb6, 0xa1, 0x76, 0x71, 0xda, 0x6f,
	0xde, 0x61, 0x0d, 0xb0, 0xcd, 0xe5, 0x69, 0x5a, 0x0c, 0xa0, 0x5e, 0x6a, 0xac, 0xb9, 0xd5, 0xfe,
	0xd5, 0x02, 0xb6, 0xc6, 0xed, 0xe1, 0x8f, 0x85, 0x51, 0xe4, 0x13, 0x68, 0x64, 0x42, 0x47, 0x2a,
	0x43, 0x9f, 0xfa, 0x73, 0x7a, 0x07, 0x55, 0xa1, 0x97, 0x42, 0x47, 0xa3, 0x0c, 0x7d, 0x6f, 0x15,
	0xc0, 0xba, 0x50, 0xd7, 0xa6, 0xe6, 0xb2, 0x27, 0xa7, 0x77, 0xff, 0x1d, 0x3d, 0x79, 0x55, 0x18,
	0xfb, 0x0e, 0x76, 0x63, 0x71, 0x33, 0x21, 0x15, 0xc9, 0x9f, 0xca, 0xa1, 0xdb, 0x83, 0xc7, 0x24,
	0x97, 0x4f, 0xd9, 0x27, 0x17, 0xe2, 0x46, 0xc6, 0x45, 0x4c, 0x2a, 0xe3, 0x26, 0xc0, 0xe8, 0x64,
	0xfd, 0x95, 0xf5, 0x9c, 0x58, 0xdc, 0x0c, 0xab, 0xec, 0xf6, 0x5f, 0x16, 0xfc, 0x7f, 0xa3, 0x05,
	0x95, 0xa5, 0x89, 0x42, 0x76, 0x02, 0xfb, 0xb1, 0xd0, 0x7e, 0x24, 0x93, 0x19, 0x2d, 0x4a, 0xb5,
	0x2c, 0x1a, 0xf9, 0x7f, 0x6f, 0x6a, 0x6f, 0x99, 0x63, 0x2c, 0xb5, 0x31, 0x88, 0xda, 0xfb, 0x06,
	0xf1, 0x10, 0x6c, 0xb3, 0x2c, 0x7a, 0xcc, 0x9c, 0x9e, 0x53, 0x05, 0x9a, 0x55, 0x78, 0x74, 0xc0,
	0x4e, 0x01, 0x48, 0xfd, 0x93, 0x1c, 0x45, 0x40, 0x4f, 0x90, 0x3d, 0x78, 0x44, 0x6d, 0x3f, 0x64,
	0x0f, 0xc6, 0xa9, 0x16, 0xf3, 0x7f, 0xdd, 0x13, 0xda, 0x7e, 0xe0, 0x7a, 0x3b, 0x64, 0x7a, 0x28,
	0x82, 0xf6, 0x2f, 0x16, 0x1c, 0x0c, 0x8a, 0x30, 0xc4, 0xdc, 0xc3, 0x10, 0x73, 0x4c, 0x7c, 0x64,
	0xf7, 0xa0, 0x9e, 0x86, 0xa1, 0x42, 0x5d, 0xa9, 0xbf, 0xb2, 0x8c, 0x7f, 0x8e, 0xc9, 0x4c, 0x47,
	0x95, 0xf0, 0x2b, 0xcb, 0xdc, 0x17, 0x5f, 0xcc, 0xe7, 0x53, 0xe1, 0xff, 0x40, 0x7d, 0xed, 0x78,
	0x2b, 0x9b, 0x31, 0xb0, 0xe9, 0xe9, 0xa0, 0x37, 0xd9, 0xa3, 0xff, 0x8d, 0x39, 0xd4, 0xdf, 0x33,
	0x87, 0x69, 0x9d, 0x4e, 0xbe, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x8f, 0xa6, 0x14, 0x1b,
	0x07, 0x00, 0x00,
}
