// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/staking/v1beta1/authz.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// AuthorizationType defines the type of staking module authorization type
type AuthorizationType int32

const (
	// AUTHORIZATION_TYPE_UNSPECIFIED specifies an unknown authorization type
	AuthorizationType_AUTHORIZATION_TYPE_UNSPECIFIED AuthorizationType = 0
	// AUTHORIZATION_TYPE_DELEGATE defines an authorization type for Msg/Delegate
	AuthorizationType_AUTHORIZATION_TYPE_DELEGATE AuthorizationType = 1
	// AUTHORIZATION_TYPE_UNDELEGATE defines an authorization type for Msg/Undelegate
	AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE AuthorizationType = 2
	// AUTHORIZATION_TYPE_REDELEGATE defines an authorization type for Msg/BeginRedelegate
	AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE AuthorizationType = 3
)

var AuthorizationType_name = map[int32]string{
	0: "AUTHORIZATION_TYPE_UNSPECIFIED",
	1: "AUTHORIZATION_TYPE_DELEGATE",
	2: "AUTHORIZATION_TYPE_UNDELEGATE",
	3: "AUTHORIZATION_TYPE_REDELEGATE",
}

var AuthorizationType_value = map[string]int32{
	"AUTHORIZATION_TYPE_UNSPECIFIED": 0,
	"AUTHORIZATION_TYPE_DELEGATE":    1,
	"AUTHORIZATION_TYPE_UNDELEGATE":  2,
	"AUTHORIZATION_TYPE_REDELEGATE":  3,
}

func (x AuthorizationType) String() string {
	return proto.EnumName(AuthorizationType_name, int32(x))
}

func (AuthorizationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6d8cdbc6f4432f0, []int{0}
}

// StakeAuthorization defines authorization for delegate/undelegate/redelegate.
type StakeAuthorization struct {
	// max_tokens specifies the maximum amount of tokens can be delegate to a validator. If it is
	// empty, there is no spend limit and any amount of coins can be delegated.
	MaxTokens *types.Coin `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"max_tokens,omitempty"`
	// validators is the oneof that represents either allow_list or deny_list
	//
	// Types that are valid to be assigned to Validators:
	//	*StakeAuthorization_AllowList
	//	*StakeAuthorization_DenyList
	Validators isStakeAuthorization_Validators `protobuf_oneof:"validators"`
	// authorization_type defines one of AuthorizationType.
	AuthorizationType AuthorizationType `protobuf:"varint,4,opt,name=authorization_type,json=authorizationType,proto3,enum=cosmos.staking.v1beta1.AuthorizationType" json:"authorization_type,omitempty"`
}

func (m *StakeAuthorization) Reset()         { *m = StakeAuthorization{} }
func (m *StakeAuthorization) String() string { return proto.CompactTextString(m) }
func (*StakeAuthorization) ProtoMessage()    {}
func (*StakeAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d8cdbc6f4432f0, []int{0}
}
func (m *StakeAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeAuthorization.Merge(m, src)
}
func (m *StakeAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *StakeAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_StakeAuthorization proto.InternalMessageInfo

type isStakeAuthorization_Validators interface {
	isStakeAuthorization_Validators()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StakeAuthorization_AllowList struct {
	AllowList *StakeAuthorization_Validators `protobuf:"bytes,2,opt,name=allow_list,json=allowList,proto3,oneof" json:"allow_list,omitempty"`
}
type StakeAuthorization_DenyList struct {
	DenyList *StakeAuthorization_Validators `protobuf:"bytes,3,opt,name=deny_list,json=denyList,proto3,oneof" json:"deny_list,omitempty"`
}

func (*StakeAuthorization_AllowList) isStakeAuthorization_Validators() {}
func (*StakeAuthorization_DenyList) isStakeAuthorization_Validators()  {}

func (m *StakeAuthorization) GetValidators() isStakeAuthorization_Validators {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *StakeAuthorization) GetMaxTokens() *types.Coin {
	if m != nil {
		return m.MaxTokens
	}
	return nil
}

func (m *StakeAuthorization) GetAllowList() *StakeAuthorization_Validators {
	if x, ok := m.GetValidators().(*StakeAuthorization_AllowList); ok {
		return x.AllowList
	}
	return nil
}

func (m *StakeAuthorization) GetDenyList() *StakeAuthorization_Validators {
	if x, ok := m.GetValidators().(*StakeAuthorization_DenyList); ok {
		return x.DenyList
	}
	return nil
}

func (m *StakeAuthorization) GetAuthorizationType() AuthorizationType {
	if m != nil {
		return m.AuthorizationType
	}
	return AuthorizationType_AUTHORIZATION_TYPE_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StakeAuthorization) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StakeAuthorization_AllowList)(nil),
		(*StakeAuthorization_DenyList)(nil),
	}
}

// Validators defines list of validator addresses.
type StakeAuthorization_Validators struct {
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (m *StakeAuthorization_Validators) Reset()         { *m = StakeAuthorization_Validators{} }
func (m *StakeAuthorization_Validators) String() string { return proto.CompactTextString(m) }
func (*StakeAuthorization_Validators) ProtoMessage()    {}
func (*StakeAuthorization_Validators) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d8cdbc6f4432f0, []int{0, 0}
}
func (m *StakeAuthorization_Validators) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeAuthorization_Validators) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeAuthorization_Validators.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeAuthorization_Validators) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeAuthorization_Validators.Merge(m, src)
}
func (m *StakeAuthorization_Validators) XXX_Size() int {
	return m.Size()
}
func (m *StakeAuthorization_Validators) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeAuthorization_Validators.DiscardUnknown(m)
}

var xxx_messageInfo_StakeAuthorization_Validators proto.InternalMessageInfo

func (m *StakeAuthorization_Validators) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterEnum("cosmos.staking.v1beta1.AuthorizationType", AuthorizationType_name, AuthorizationType_value)
	proto.RegisterType((*StakeAuthorization)(nil), "cosmos.staking.v1beta1.StakeAuthorization")
	proto.RegisterType((*StakeAuthorization_Validators)(nil), "cosmos.staking.v1beta1.StakeAuthorization.Validators")
}

func init() {
	proto.RegisterFile("cosmos/staking/v1beta1/authz.proto", fileDescriptor_d6d8cdbc6f4432f0)
}

var fileDescriptor_d6d8cdbc6f4432f0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4d, 0x6e, 0xd3, 0x40,
	0x18, 0x86, 0xed, 0x06, 0x01, 0xf9, 0xf8, 0x51, 0x33, 0xaa, 0x50, 0x1a, 0x84, 0x5b, 0xb2, 0x21,
	0xfc, 0x64, 0xac, 0x06, 0xb1, 0x61, 0x85, 0xd3, 0xba, 0x34, 0x52, 0xd5, 0x56, 0x8e, 0x5b, 0x41,
	0x37, 0xd6, 0x24, 0x1e, 0x39, 0xa3, 0xc4, 0x9e, 0xc8, 0x33, 0x29, 0x49, 0x4f, 0xc1, 0x09, 0x38,
	0x00, 0xeb, 0x1e, 0x02, 0xb1, 0xaa, 0x58, 0xb1, 0x03, 0x25, 0x12, 0xe7, 0x40, 0xb6, 0xa7, 0x86,
	0xd2, 0x84, 0x4d, 0x57, 0xb6, 0xfc, 0x3d, 0x7e, 0xde, 0x77, 0xac, 0xcf, 0x50, 0xed, 0x72, 0x11,
	0x72, 0x61, 0x0a, 0x49, 0xfa, 0x2c, 0x0a, 0xcc, 0x93, 0x8d, 0x0e, 0x95, 0x64, 0xc3, 0x24, 0x23,
	0xd9, 0x3b, 0xc5, 0xc3, 0x98, 0x4b, 0x8e, 0x1e, 0x64, 0x0c, 0x56, 0x0c, 0x56, 0x4c, 0x65, 0x25,
	0xe0, 0x01, 0x4f, 0x11, 0x33, 0xb9, 0xcb, 0xe8, 0xca, 0x6a, 0x46, 0x7b, 0xd9, 0x40, 0xbd, 0x9a,
	0x8d, 0x0c, 0x15, 0xd6, 0x21, 0x82, 0xe6, 0x49, 0x5d, 0xce, 0xa2, 0x6c, 0x5e, 0xfd, 0x55, 0x00,
	0xd4, 0x96, 0xa4, 0x4f, 0xad, 0x91, 0xec, 0xf1, 0x98, 0x9d, 0x12, 0xc9, 0x78, 0x84, 0x28, 0x40,
	0x48, 0xc6, 0x9e, 0xe4, 0x7d, 0x1a, 0x89, 0xb2, 0xbe, 0xae, 0xd7, 0xee, 0x34, 0x56, 0xb1, 0x32,
	0x27, 0xae, 0x8b, 0x46, 0x78, 0x93, 0xb3, 0xa8, 0xf9, 0xfc, 0xf3, 0x8f, 0xb5, 0x27, 0x01, 0x93,
	0xbd, 0x51, 0x07, 0x77, 0x79, 0xa8, 0x2a, 0xa8, 0x4b, 0x5d, 0xf8, 0x7d, 0x53, 0x4e, 0x86, 0x54,
	0xa4, 0xb0, 0x53, 0x0c, 0xc9, 0xd8, 0x4d, 0xc5, 0xe8, 0x08, 0x80, 0x0c, 0x06, 0xfc, 0x83, 0x37,
	0x60, 0x42, 0x96, 0x97, 0xd2, 0x98, 0x57, 0x78, 0xfe, 0xd9, 0xf1, 0xd5, 0x9a, 0xf8, 0x88, 0x0c,
	0x98, 0x4f, 0x24, 0x8f, 0xc5, 0x8e, 0xe6, 0x14, 0x53, 0xd5, 0x2e, 0x13, 0x12, 0xb9, 0x50, 0xf4,
	0x69, 0x34, 0xc9, 0xb4, 0x85, 0xeb, 0x69, 0x6f, 0x27, 0xa6, 0xd4, 0xfa, 0x0e, 0x10, 0xf9, 0x9b,
	0xf3, 0x92, 0x43, 0x95, 0x6f, 0xac, 0xeb, 0xb5, 0xfb, 0x8d, 0xa7, 0x8b, 0xf4, 0x97, 0xcc, 0xee,
	0x64, 0x48, 0x9d, 0x12, 0xf9, 0xf7, 0x51, 0xe5, 0x0d, 0xc0, 0x9f, 0x4c, 0xd4, 0x80, 0x5b, 0xc4,
	0xf7, 0x63, 0x2a, 0x92, 0x2f, 0x5f, 0xa8, 0x15, 0x9b, 0xe5, 0x6f, 0x67, 0xf5, 0x15, 0xe5, 0xb7,
	0xb2, 0x49, 0x5b, 0xc6, 0x2c, 0x0a, 0x9c, 0x0b, 0xf0, 0x75, 0xe9, 0xeb, 0x59, 0xfd, 0xde, 0xa5,
	0xac, 0xe6, 0x5d, 0x80, 0x93, 0x5c, 0xfa, 0xec, 0x93, 0x0e, 0xa5, 0x2b, 0x5d, 0x50, 0x15, 0x0c,
	0xeb, 0xd0, 0xdd, 0xd9, 0x77, 0x5a, 0xc7, 0x96, 0xdb, 0xda, 0xdf, 0xf3, 0xdc, 0xf7, 0x07, 0xb6,
	0x77, 0xb8, 0xd7, 0x3e, 0xb0, 0x37, 0x5b, 0xdb, 0x2d, 0x7b, 0x6b, 0x59, 0x43, 0x6b, 0xf0, 0x70,
	0x0e, 0xb3, 0x65, 0xef, 0xda, 0x6f, 0x2d, 0xd7, 0x5e, 0xd6, 0xd1, 0x63, 0x78, 0x34, 0x57, 0x92,
	0x23, 0x4b, 0x0b, 0x10, 0xc7, 0xce, 0x91, 0x42, 0x73, 0xfb, 0xcb, 0xd4, 0xd0, 0xcf, 0xa7, 0x86,
	0xfe, 0x73, 0x6a, 0xe8, 0x1f, 0x67, 0x86, 0x76, 0x3e, 0x33, 0xb4, 0xef, 0x33, 0x43, 0x3b, 0x7e,
	0xf1, 0xdf, 0xcd, 0x1a, 0xe7, 0x3f, 0x52, 0xba, 0x63, 0x9d, 0x9b, 0xe9, 0x62, 0xbf, 0xfc, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0xd5, 0x4d, 0xf3, 0x67, 0x03, 0x00, 0x00,
}

func (m *StakeAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuthorizationType != 0 {
		i = encodeVarintAuthz(dAtA, i, uint64(m.AuthorizationType))
		i--
		dAtA[i] = 0x20
	}
	if m.Validators != nil {
		{
			size := m.Validators.Size()
			i -= size
			if _, err := m.Validators.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.MaxTokens != nil {
		{
			size, err := m.MaxTokens.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StakeAuthorization_AllowList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization_AllowList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AllowList != nil {
		{
			size, err := m.AllowList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *StakeAuthorization_DenyList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization_DenyList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DenyList != nil {
		{
			size, err := m.DenyList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *StakeAuthorization_Validators) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeAuthorization_Validators) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeAuthorization_Validators) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		for iNdEx := len(m.Address) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Address[iNdEx])
			copy(dAtA[i:], m.Address[iNdEx])
			i = encodeVarintAuthz(dAtA, i, uint64(len(m.Address[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakeAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTokens != nil {
		l = m.MaxTokens.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	if m.Validators != nil {
		n += m.Validators.Size()
	}
	if m.AuthorizationType != 0 {
		n += 1 + sovAuthz(uint64(m.AuthorizationType))
	}
	return n
}

func (m *StakeAuthorization_AllowList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AllowList != nil {
		l = m.AllowList.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}
func (m *StakeAuthorization_DenyList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DenyList != nil {
		l = m.DenyList.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}
func (m *StakeAuthorization_Validators) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Address) > 0 {
		for _, s := range m.Address {
			l = len(s)
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakeAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StakeAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakeAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxTokens == nil {
				m.MaxTokens = &types.Coin{}
			}
			if err := m.MaxTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StakeAuthorization_Validators{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Validators = &StakeAuthorization_AllowList{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenyList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StakeAuthorization_Validators{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Validators = &StakeAuthorization_DenyList{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationType", wireType)
			}
			m.AuthorizationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthorizationType |= AuthorizationType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StakeAuthorization_Validators) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Validators: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validators: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthz
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
