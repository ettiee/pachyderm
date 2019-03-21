// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/pkg/config/config.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config specifies the pachyderm config that is read and interpreted by the
// pachctl command-line tool. Right now, this is stored at
// $HOME/.pachyderm/config.
//
// Different versions of the pachyderm config are specified as subfields of this
// message (this allows us to make significant changes to the config structure
// without breaking existing users by defining a new config version).
//
// DO NOT change or remove field numbers from this proto, otherwise ALL
// pachyderm user configs will fail to parse.
type Config struct {
	UserID string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Configuration options. Exactly one of these fields should be set
	// (depending on which version of the config is being used)
	V1                   *ConfigV1 `protobuf:"bytes,2,opt,name=v1,proto3" json:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f651abce1dcdf3, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Config) GetV1() *ConfigV1 {
	if m != nil {
		return m.V1
	}
	return nil
}

// ConfigV1 specifies v1 of the pachyderm config (June 30 2017 - present)
// DO NOT change or remove field numbers from this proto, as if you do, v1 user
// configs will become unparseable.
type ConfigV1 struct {
	// A host:port pointing pachd at a pachyderm cluster. Similar to the
	// PACHD_ADDRESS environment variable, though PACHD_ADDRESS overrides
	// this.
	PachdAddress string `protobuf:"bytes,2,opt,name=pachd_address,json=pachdAddress,proto3" json:"pachd_address,omitempty"`
	// Trusted root certificates (overrides installed certificates), formatted
	// as base64-encoded PEM
	ServerCAs string `protobuf:"bytes,3,opt,name=server_cas,json=serverCas,proto3" json:"server_cas,omitempty"`
	// A secret token identifying the current pachctl user within their
	// pachyderm cluster. This is included in all RPCs sent by pachctl, and used
	// to determine if pachctl actions are authorized.
	SessionToken         string   `protobuf:"bytes,1,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigV1) Reset()         { *m = ConfigV1{} }
func (m *ConfigV1) String() string { return proto.CompactTextString(m) }
func (*ConfigV1) ProtoMessage()    {}
func (*ConfigV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f651abce1dcdf3, []int{1}
}
func (m *ConfigV1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigV1.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigV1.Merge(m, src)
}
func (m *ConfigV1) XXX_Size() int {
	return m.Size()
}
func (m *ConfigV1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigV1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigV1 proto.InternalMessageInfo

func (m *ConfigV1) GetPachdAddress() string {
	if m != nil {
		return m.PachdAddress
	}
	return ""
}

func (m *ConfigV1) GetServerCAs() string {
	if m != nil {
		return m.ServerCAs
	}
	return ""
}

func (m *ConfigV1) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "config.Config")
	proto.RegisterType((*ConfigV1)(nil), "config.ConfigV1")
}

func init() { proto.RegisterFile("client/pkg/config/config.proto", fileDescriptor_60f651abce1dcdf3) }

var fileDescriptor_60f651abce1dcdf3 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x84, 0x51, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x48,
	0x1f, 0xc4, 0x82, 0xc8, 0x2a, 0xf9, 0x73, 0xb1, 0x39, 0x83, 0xe5, 0x85, 0x94, 0xb9, 0xd8, 0x4b,
	0x8b, 0x53, 0x8b, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xb8, 0x1e, 0xdd,
	0x93, 0x67, 0x0b, 0x2d, 0x4e, 0x2d, 0xf2, 0x74, 0x09, 0x62, 0x03, 0x49, 0x79, 0xa6, 0x08, 0x29,
	0x70, 0x31, 0x95, 0x19, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xe8, 0x41, 0xed, 0x81,
	0x18, 0x10, 0x66, 0x18, 0xc4, 0x54, 0x66, 0xa8, 0xd4, 0xc4, 0xc8, 0xc5, 0x01, 0x13, 0x10, 0x52,
	0xe6, 0xe2, 0x2d, 0x48, 0x4c, 0xce, 0x48, 0x89, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x06,
	0xeb, 0xe4, 0x0c, 0xe2, 0x01, 0x0b, 0x3a, 0x42, 0xc4, 0x84, 0x74, 0xb8, 0xb8, 0x8a, 0x53, 0x8b,
	0xca, 0x52, 0x8b, 0xe2, 0x93, 0x13, 0x8b, 0x25, 0x98, 0xc1, 0x76, 0xf3, 0x3e, 0xba, 0x27, 0xcf,
	0x19, 0x0c, 0x16, 0x75, 0x76, 0x2c, 0x0e, 0xe2, 0x84, 0x28, 0x70, 0x4e, 0x2c, 0x06, 0x19, 0x59,
	0x9c, 0x5a, 0x5c, 0x9c, 0x99, 0x9f, 0x17, 0x5f, 0x92, 0x9f, 0x9d, 0x9a, 0x07, 0x71, 0x6c, 0x10,
	0x0f, 0x54, 0x30, 0x04, 0x24, 0xe6, 0xe4, 0x76, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xb2, 0xbd, 0x32, 0x25, 0xb5, 0x08, 0x99, 0x55, 0x5c, 0x94,
	0xac, 0x8f, 0x11, 0x90, 0x49, 0x6c, 0xe0, 0x40, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x39,
	0x15, 0x7f, 0xf9, 0x64, 0x01, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	if m.V1 != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.V1.Size()))
		n1, err := m.V1.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConfigV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigV1) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SessionToken) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.SessionToken)))
		i += copy(dAtA[i:], m.SessionToken)
	}
	if len(m.PachdAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.PachdAddress)))
		i += copy(dAtA[i:], m.PachdAddress)
	}
	if len(m.ServerCAs) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ServerCAs)))
		i += copy(dAtA[i:], m.ServerCAs)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.V1 != nil {
		l = m.V1.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigV1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionToken)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.PachdAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ServerCAs)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.V1 == nil {
				m.V1 = &ConfigV1{}
			}
			if err := m.V1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ConfigV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PachdAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PachdAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerCAs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerCAs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)
