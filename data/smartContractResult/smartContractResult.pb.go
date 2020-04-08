// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smartContractResult.proto

package smartContractResult

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_data "github.com/ElrondNetwork/elrond-go/data"
	github_com_ElrondNetwork_elrond_vm_common "github.com/ElrondNetwork/elrond-vm-common"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Transaction holds all the data needed for a value transfer or SC call
type SmartContractResult struct {
	Nonce        uint64                                             `protobuf:"varint,1,opt,name=Nonce,proto3" json:"nonce"`
	Value        *math_big.Int                                      `protobuf:"bytes,2,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"value"`
	RcvAddr      []byte                                             `protobuf:"bytes,3,opt,name=RcvAddr,proto3" json:"receiver"`
	SndAddr      []byte                                             `protobuf:"bytes,4,opt,name=SndAddr,proto3" json:"sender"`
	Code         []byte                                             `protobuf:"bytes,5,opt,name=Code,proto3" json:"code,omitempty"`
	Data         []byte                                             `protobuf:"bytes,6,opt,name=Data,proto3" json:"data,omitempty"`
	TxHash       []byte                                             `protobuf:"bytes,7,opt,name=TxHash,proto3" json:"txHash"`
	GasLimit     uint64                                             `protobuf:"varint,8,opt,name=GasLimit,proto3" json:"gasLimit"`
	GasPrice     uint64                                             `protobuf:"varint,9,opt,name=GasPrice,proto3" json:"gasPrice"`
	CallType     github_com_ElrondNetwork_elrond_vm_common.CallType `protobuf:"varint,10,opt,name=CallType,proto3,casttype=github.com/ElrondNetwork/elrond-vm-common.CallType" json:"callType"`
	CodeMetadata []byte                                             `protobuf:"bytes,11,opt,name=CodeMetadata,proto3" json:"codeMetadata,omitempty"`
}

func (m *SmartContractResult) Reset()      { *m = SmartContractResult{} }
func (*SmartContractResult) ProtoMessage() {}
func (*SmartContractResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1605de0d3d805, []int{0}
}
func (m *SmartContractResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartContractResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SmartContractResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartContractResult.Merge(m, src)
}
func (m *SmartContractResult) XXX_Size() int {
	return m.Size()
}
func (m *SmartContractResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartContractResult.DiscardUnknown(m)
}

var xxx_messageInfo_SmartContractResult proto.InternalMessageInfo

func (m *SmartContractResult) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *SmartContractResult) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SmartContractResult) GetRcvAddr() []byte {
	if m != nil {
		return m.RcvAddr
	}
	return nil
}

func (m *SmartContractResult) GetSndAddr() []byte {
	if m != nil {
		return m.SndAddr
	}
	return nil
}

func (m *SmartContractResult) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *SmartContractResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SmartContractResult) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *SmartContractResult) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *SmartContractResult) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *SmartContractResult) GetCallType() github_com_ElrondNetwork_elrond_vm_common.CallType {
	if m != nil {
		return m.CallType
	}
	return 0
}

func (m *SmartContractResult) GetCodeMetadata() []byte {
	if m != nil {
		return m.CodeMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*SmartContractResult)(nil), "proto.SmartContractResult")
}

func init() { proto.RegisterFile("smartContractResult.proto", fileDescriptor_edc1605de0d3d805) }

var fileDescriptor_edc1605de0d3d805 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x63, 0xd6, 0x76, 0x9d, 0xa9, 0x38, 0x78, 0x12, 0x32, 0x3b, 0x38, 0xd5, 0x84, 0xa6,
	0x1e, 0x68, 0x22, 0xc1, 0x11, 0x09, 0x69, 0x29, 0x08, 0x26, 0xb1, 0x0a, 0x79, 0x13, 0x07, 0x0e,
	0x48, 0xae, 0x63, 0xd2, 0x88, 0x24, 0xae, 0x1c, 0xb7, 0xb0, 0x1b, 0x1f, 0x81, 0x8f, 0x81, 0xf8,
	0x24, 0x1c, 0x2b, 0x4e, 0x3d, 0x05, 0x9a, 0x5e, 0x50, 0x4e, 0x3b, 0x73, 0x42, 0x76, 0xb3, 0xae,
	0x13, 0x48, 0x9c, 0xda, 0xf7, 0xfb, 0xff, 0xf2, 0xec, 0xf7, 0x64, 0x78, 0x2f, 0x4f, 0x99, 0xd2,
	0x03, 0x99, 0x69, 0xc5, 0xb8, 0xa6, 0x22, 0x9f, 0x26, 0xda, 0x9b, 0x28, 0xa9, 0x25, 0x6a, 0xda,
	0x9f, 0x83, 0x7e, 0x14, 0xeb, 0xf1, 0x74, 0xe4, 0x71, 0x99, 0xfa, 0x91, 0x8c, 0xa4, 0x6f, 0xf1,
	0x68, 0xfa, 0xce, 0x56, 0xb6, 0xb0, 0xff, 0xd6, 0x5f, 0x1d, 0x7e, 0x6f, 0xc0, 0xfd, 0xb3, 0xbf,
	0x7b, 0x22, 0x17, 0x36, 0x87, 0x32, 0xe3, 0x02, 0x83, 0x2e, 0xe8, 0x35, 0x82, 0xbd, 0xaa, 0x70,
	0x9b, 0x99, 0x01, 0x74, 0xcd, 0x51, 0x08, 0x9b, 0xaf, 0x59, 0x32, 0x15, 0xf8, 0x56, 0x17, 0xf4,
	0x3a, 0xc1, 0xd0, 0x08, 0x33, 0x03, 0xbe, 0xfe, 0x70, 0x8f, 0x53, 0xa6, 0xc7, 0xfe, 0x28, 0x8e,
	0xbc, 0x93, 0x4c, 0x3f, 0xde, 0xba, 0xd0, 0xb3, 0x44, 0xc9, 0x2c, 0x1c, 0x0a, 0xfd, 0x41, 0xaa,
	0xf7, 0xbe, 0xb0, 0x55, 0x3f, 0x92, 0x7e, 0xc8, 0x34, 0xf3, 0x82, 0x38, 0x3a, 0xc9, 0xf4, 0x80,
	0xe5, 0x5a, 0x28, 0xba, 0x6e, 0x8e, 0x8e, 0xe0, 0x2e, 0xe5, 0xb3, 0xe3, 0x30, 0x54, 0x78, 0xc7,
	0x9e, 0xd3, 0xa9, 0x0a, 0xb7, 0xad, 0x04, 0x17, 0xf1, 0x4c, 0x28, 0x7a, 0x15, 0xa2, 0xfb, 0x70,
	0xf7, 0x2c, 0x0b, 0xad, 0xd7, 0xb0, 0x1e, 0xac, 0x0a, 0xb7, 0x95, 0x8b, 0x2c, 0x34, 0x56, 0x1d,
	0xa1, 0x23, 0xd8, 0x18, 0xc8, 0x50, 0xe0, 0xa6, 0x55, 0x50, 0x55, 0xb8, 0x77, 0xb8, 0x0c, 0xc5,
	0x03, 0x99, 0xc6, 0x5a, 0xa4, 0x13, 0x7d, 0x41, 0x6d, 0x6e, 0xbc, 0xa7, 0x4c, 0x33, 0xdc, 0xba,
	0xf6, 0xcc, 0x0d, 0xb7, 0x3d, 0x93, 0xa3, 0x43, 0xd8, 0x3a, 0xff, 0xf8, 0x82, 0xe5, 0x63, 0xbc,
	0x7b, 0x7d, 0xa8, 0xb6, 0x84, 0xd6, 0x09, 0xea, 0xc1, 0xf6, 0x73, 0x96, 0xbf, 0x8c, 0xd3, 0x58,
	0xe3, 0xb6, 0xdd, 0xa5, 0x1d, 0x21, 0xaa, 0x19, 0xdd, 0xa4, 0xb5, 0xf9, 0x4a, 0xc5, 0x5c, 0xe0,
	0xbd, 0x1b, 0xa6, 0x65, 0x74, 0x93, 0xa2, 0xb7, 0xb0, 0x3d, 0x60, 0x49, 0x72, 0x7e, 0x31, 0x11,
	0x18, 0x76, 0x41, 0x6f, 0x27, 0x08, 0x8c, 0xc9, 0x6b, 0xf6, 0xbb, 0x70, 0x1f, 0xfe, 0x6f, 0xe9,
	0xb3, 0xb4, 0xcf, 0x65, 0x9a, 0xca, 0xcc, 0xbb, 0xea, 0x44, 0x37, 0x3d, 0xd1, 0x13, 0xd8, 0x31,
	0x7b, 0x38, 0x15, 0x9a, 0x99, 0xb9, 0xf1, 0x6d, 0x3b, 0xdd, 0x41, 0x55, 0xb8, 0x77, 0xf9, 0x16,
	0xdf, 0xda, 0xc7, 0x0d, 0x3f, 0x38, 0x9d, 0x2f, 0x89, 0xb3, 0x58, 0x12, 0xe7, 0x72, 0x49, 0xc0,
	0xa7, 0x92, 0x80, 0x2f, 0x25, 0x01, 0xdf, 0x4a, 0x02, 0xe6, 0x25, 0x01, 0x8b, 0x92, 0x80, 0x9f,
	0x25, 0x01, 0xbf, 0x4a, 0xe2, 0x5c, 0x96, 0x04, 0x7c, 0x5e, 0x11, 0x67, 0xbe, 0x22, 0xce, 0x62,
	0x45, 0x9c, 0x37, 0xfb, 0xff, 0x78, 0xdf, 0xa3, 0x96, 0x7d, 0xaa, 0x8f, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x29, 0x16, 0x1c, 0x99, 0xfd, 0x02, 0x00, 0x00,
}

func (this *SmartContractResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SmartContractResult)
	if !ok {
		that2, ok := that.(SmartContractResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.RcvAddr, that1.RcvAddr) {
		return false
	}
	if !bytes.Equal(this.SndAddr, that1.SndAddr) {
		return false
	}
	if !bytes.Equal(this.Code, that1.Code) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.TxHash, that1.TxHash) {
		return false
	}
	if this.GasLimit != that1.GasLimit {
		return false
	}
	if this.GasPrice != that1.GasPrice {
		return false
	}
	if this.CallType != that1.CallType {
		return false
	}
	if !bytes.Equal(this.CodeMetadata, that1.CodeMetadata) {
		return false
	}
	return true
}
func (this *SmartContractResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 15)
	s = append(s, "&smartContractResult.SmartContractResult{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "RcvAddr: "+fmt.Sprintf("%#v", this.RcvAddr)+",\n")
	s = append(s, "SndAddr: "+fmt.Sprintf("%#v", this.SndAddr)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "TxHash: "+fmt.Sprintf("%#v", this.TxHash)+",\n")
	s = append(s, "GasLimit: "+fmt.Sprintf("%#v", this.GasLimit)+",\n")
	s = append(s, "GasPrice: "+fmt.Sprintf("%#v", this.GasPrice)+",\n")
	s = append(s, "CallType: "+fmt.Sprintf("%#v", this.CallType)+",\n")
	s = append(s, "CodeMetadata: "+fmt.Sprintf("%#v", this.CodeMetadata)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSmartContractResult(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SmartContractResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartContractResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartContractResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CodeMetadata) > 0 {
		i -= len(m.CodeMetadata)
		copy(dAtA[i:], m.CodeMetadata)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.CodeMetadata)))
		i--
		dAtA[i] = 0x5a
	}
	if m.CallType != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.CallType))
		i--
		dAtA[i] = 0x50
	}
	if m.GasPrice != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x48
	}
	if m.GasLimit != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x40
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SndAddr) > 0 {
		i -= len(m.SndAddr)
		copy(dAtA[i:], m.SndAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.SndAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RcvAddr) > 0 {
		i -= len(m.RcvAddr)
		copy(dAtA[i:], m.RcvAddr)
		i = encodeVarintSmartContractResult(dAtA, i, uint64(len(m.RcvAddr)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSmartContractResult(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Nonce != 0 {
		i = encodeVarintSmartContractResult(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSmartContractResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovSmartContractResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SmartContractResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovSmartContractResult(uint64(m.Nonce))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.RcvAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.SndAddr)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovSmartContractResult(uint64(m.GasLimit))
	}
	if m.GasPrice != 0 {
		n += 1 + sovSmartContractResult(uint64(m.GasPrice))
	}
	if m.CallType != 0 {
		n += 1 + sovSmartContractResult(uint64(m.CallType))
	}
	l = len(m.CodeMetadata)
	if l > 0 {
		n += 1 + l + sovSmartContractResult(uint64(l))
	}
	return n
}

func sovSmartContractResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSmartContractResult(x uint64) (n int) {
	return sovSmartContractResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SmartContractResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SmartContractResult{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`RcvAddr:` + fmt.Sprintf("%v", this.RcvAddr) + `,`,
		`SndAddr:` + fmt.Sprintf("%v", this.SndAddr) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`TxHash:` + fmt.Sprintf("%v", this.TxHash) + `,`,
		`GasLimit:` + fmt.Sprintf("%v", this.GasLimit) + `,`,
		`GasPrice:` + fmt.Sprintf("%v", this.GasPrice) + `,`,
		`CallType:` + fmt.Sprintf("%v", this.CallType) + `,`,
		`CodeMetadata:` + fmt.Sprintf("%v", this.CodeMetadata) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSmartContractResult(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SmartContractResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSmartContractResult
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
			return fmt.Errorf("proto: SmartContractResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartContractResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RcvAddr = append(m.RcvAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RcvAddr == nil {
				m.RcvAddr = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SndAddr = append(m.SndAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SndAddr == nil {
				m.SndAddr = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallType", wireType)
			}
			m.CallType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CallType |= github_com_ElrondNetwork_elrond_vm_common.CallType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeMetadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSmartContractResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeMetadata = append(m.CodeMetadata[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeMetadata == nil {
				m.CodeMetadata = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSmartContractResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSmartContractResult
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSmartContractResult
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
func skipSmartContractResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSmartContractResult
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
					return 0, ErrIntOverflowSmartContractResult
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
					return 0, ErrIntOverflowSmartContractResult
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
				return 0, ErrInvalidLengthSmartContractResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSmartContractResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSmartContractResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSmartContractResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSmartContractResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSmartContractResult = fmt.Errorf("proto: unexpected end of group")
)
