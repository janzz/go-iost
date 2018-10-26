// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/contract/contract.proto

package contract

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Info struct {
	Lang                 string   `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Abi                  []*ABI   `protobuf:"bytes,3,rep,name=abi" json:"abi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_contract_57af315884a7a5bd, []int{0}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Info.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return m.Size()
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Info) GetAbi() []*ABI {
	if m != nil {
		return m.Abi
	}
	return nil
}

type ABI struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payment              int32     `protobuf:"varint,2,opt,name=payment,proto3" json:"payment,omitempty"`
	Limit                *Cost     `protobuf:"bytes,3,opt,name=limit" json:"limit,omitempty"`
	GasPrice             int64     `protobuf:"varint,4,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	Args                 []string  `protobuf:"bytes,5,rep,name=args" json:"args,omitempty"`
	AmountLimit          []*Amount `protobuf:"bytes,6,rep,name=amountLimit" json:"amountLimit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ABI) Reset()         { *m = ABI{} }
func (m *ABI) String() string { return proto.CompactTextString(m) }
func (*ABI) ProtoMessage()    {}
func (*ABI) Descriptor() ([]byte, []int) {
	return fileDescriptor_contract_57af315884a7a5bd, []int{1}
}
func (m *ABI) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ABI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ABI.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ABI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABI.Merge(dst, src)
}
func (m *ABI) XXX_Size() int {
	return m.Size()
}
func (m *ABI) XXX_DiscardUnknown() {
	xxx_messageInfo_ABI.DiscardUnknown(m)
}

var xxx_messageInfo_ABI proto.InternalMessageInfo

func (m *ABI) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ABI) GetPayment() int32 {
	if m != nil {
		return m.Payment
	}
	return 0
}

func (m *ABI) GetLimit() *Cost {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *ABI) GetGasPrice() int64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *ABI) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ABI) GetAmountLimit() []*Amount {
	if m != nil {
		return m.AmountLimit
	}
	return nil
}

type Amount struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Val                  int64    `protobuf:"varint,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Amount) Reset()         { *m = Amount{} }
func (m *Amount) String() string { return proto.CompactTextString(m) }
func (*Amount) ProtoMessage()    {}
func (*Amount) Descriptor() ([]byte, []int) {
	return fileDescriptor_contract_57af315884a7a5bd, []int{2}
}
func (m *Amount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Amount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Amount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Amount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Amount.Merge(dst, src)
}
func (m *Amount) XXX_Size() int {
	return m.Size()
}
func (m *Amount) XXX_DiscardUnknown() {
	xxx_messageInfo_Amount.DiscardUnknown(m)
}

var xxx_messageInfo_Amount proto.InternalMessageInfo

func (m *Amount) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Amount) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

type Cost struct {
	Data                 int64    `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Net                  int64    `protobuf:"varint,2,opt,name=net,proto3" json:"net,omitempty"`
	CPU                  int64    `protobuf:"varint,3,opt,name=CPU,proto3" json:"CPU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cost) Reset()         { *m = Cost{} }
func (m *Cost) String() string { return proto.CompactTextString(m) }
func (*Cost) ProtoMessage()    {}
func (*Cost) Descriptor() ([]byte, []int) {
	return fileDescriptor_contract_57af315884a7a5bd, []int{3}
}
func (m *Cost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Cost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cost.Merge(dst, src)
}
func (m *Cost) XXX_Size() int {
	return m.Size()
}
func (m *Cost) XXX_DiscardUnknown() {
	xxx_messageInfo_Cost.DiscardUnknown(m)
}

var xxx_messageInfo_Cost proto.InternalMessageInfo

func (m *Cost) GetData() int64 {
	if m != nil {
		return m.Data
	}
	return 0
}

func (m *Cost) GetNet() int64 {
	if m != nil {
		return m.Net
	}
	return 0
}

func (m *Cost) GetCPU() int64 {
	if m != nil {
		return m.CPU
	}
	return 0
}

type Contract struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Info                 *Info    `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_contract_57af315884a7a5bd, []int{4}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(dst, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Contract) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Contract) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*Info)(nil), "contract.Info")
	proto.RegisterType((*ABI)(nil), "contract.ABI")
	proto.RegisterType((*Amount)(nil), "contract.Amount")
	proto.RegisterType((*Cost)(nil), "contract.Cost")
	proto.RegisterType((*Contract)(nil), "contract.Contract")
}
func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Lang) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContract(dAtA, i, uint64(len(m.Lang)))
		i += copy(dAtA[i:], m.Lang)
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContract(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if len(m.Abi) > 0 {
		for _, msg := range m.Abi {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintContract(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ABI) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ABI) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContract(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Payment != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.Payment))
	}
	if m.Limit != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.Limit.Size()))
		n1, err := m.Limit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.GasPrice != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.GasPrice))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.AmountLimit) > 0 {
		for _, msg := range m.AmountLimit {
			dAtA[i] = 0x32
			i++
			i = encodeVarintContract(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Amount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Amount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContract(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.Val != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.Val))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Cost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cost) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.Data))
	}
	if m.Net != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.Net))
	}
	if m.CPU != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.CPU))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContract(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Info != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContract(dAtA, i, uint64(m.Info.Size()))
		n2, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Code) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintContract(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Info) Size() (n int) {
	var l int
	_ = l
	l = len(m.Lang)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if len(m.Abi) > 0 {
		for _, e := range m.Abi {
			l = e.Size()
			n += 1 + l + sovContract(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ABI) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.Payment != 0 {
		n += 1 + sovContract(uint64(m.Payment))
	}
	if m.Limit != nil {
		l = m.Limit.Size()
		n += 1 + l + sovContract(uint64(l))
	}
	if m.GasPrice != 0 {
		n += 1 + sovContract(uint64(m.GasPrice))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovContract(uint64(l))
		}
	}
	if len(m.AmountLimit) > 0 {
		for _, e := range m.AmountLimit {
			l = e.Size()
			n += 1 + l + sovContract(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Amount) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.Val != 0 {
		n += 1 + sovContract(uint64(m.Val))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Cost) Size() (n int) {
	var l int
	_ = l
	if m.Data != 0 {
		n += 1 + sovContract(uint64(m.Data))
	}
	if m.Net != 0 {
		n += 1 + sovContract(uint64(m.Net))
	}
	if m.CPU != 0 {
		n += 1 + sovContract(uint64(m.CPU))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Contract) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovContract(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lang = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abi = append(m.Abi, &ABI{})
			if err := m.Abi[len(m.Abi)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContract
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
func (m *ABI) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ABI: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ABI: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payment", wireType)
			}
			m.Payment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Payment |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Limit == nil {
				m.Limit = &Cost{}
			}
			if err := m.Limit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountLimit = append(m.AmountLimit, &Amount{})
			if err := m.AmountLimit[len(m.AmountLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContract
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
func (m *Amount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Amount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Amount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			m.Val = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Val |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContract
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
func (m *Cost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			m.Data = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Data |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Net", wireType)
			}
			m.Net = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Net |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPU", wireType)
			}
			m.CPU = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CPU |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContract
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
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &Info{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthContract
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContract
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
				next, err := skipContract(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthContract = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("core/contract/contract.proto", fileDescriptor_contract_57af315884a7a5bd)
}

var fileDescriptor_contract_57af315884a7a5bd = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4f, 0x6b, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0x75, 0x13, 0x5f, 0x1d, 0x79, 0x45, 0x96, 0xf7, 0x10, 0x68, 0x49, 0x43, 0xe8,
	0x21, 0x27, 0x5b, 0xec, 0xbd, 0xe0, 0x9f, 0x4b, 0xa0, 0x07, 0x59, 0xf0, 0x5c, 0xd6, 0xb8, 0x86,
	0x50, 0xb3, 0x2b, 0xc9, 0x56, 0xe8, 0x37, 0xe9, 0xa7, 0xe9, 0xb9, 0xc7, 0x7e, 0x84, 0x62, 0xbf,
	0x48, 0x99, 0xd1, 0x58, 0x7b, 0x7b, 0x9e, 0x99, 0xe1, 0x37, 0xcf, 0x24, 0x0b, 0x97, 0x99, 0xad,
	0xf4, 0x4d, 0x66, 0x8d, 0xab, 0x54, 0xe6, 0x4e, 0x62, 0xb8, 0xad, 0xac, 0xb3, 0xa2, 0xd3, 0xf8,
	0x78, 0x01, 0x5e, 0x6a, 0xd6, 0x56, 0x08, 0xf0, 0x36, 0xca, 0xe4, 0x01, 0x8b, 0x58, 0xd2, 0x95,
	0xa4, 0x45, 0x00, 0x7f, 0x77, 0xba, 0xaa, 0x0b, 0x6b, 0x82, 0x16, 0x95, 0x1b, 0x2b, 0xae, 0x80,
	0xab, 0x65, 0x11, 0xf0, 0x88, 0x27, 0xbd, 0xd1, 0xbf, 0xe1, 0x89, 0x3e, 0x9e, 0xa4, 0x12, 0x3b,
	0xf1, 0x1b, 0x03, 0x3e, 0x9e, 0xa4, 0x88, 0x35, 0xaa, 0xd4, 0x0d, 0x16, 0x35, 0x62, 0xb7, 0xea,
	0xa5, 0xd4, 0xc6, 0x11, 0xd6, 0x97, 0x8d, 0x15, 0xd7, 0xe0, 0x6f, 0x8a, 0xb2, 0x70, 0x01, 0x8f,
	0x58, 0xd2, 0x1b, 0xf5, 0x7f, 0xc0, 0x53, 0x5b, 0x3b, 0x79, 0x68, 0x8a, 0x0b, 0xe8, 0xe6, 0xaa,
	0x7e, 0xdc, 0x56, 0x45, 0xa6, 0x03, 0x2f, 0x62, 0x09, 0x97, 0x9d, 0x5c, 0xd5, 0x73, 0xf4, 0xb8,
	0x50, 0x55, 0x79, 0x1d, 0xf8, 0x11, 0xc7, 0x85, 0xa8, 0xc5, 0x08, 0x7a, 0xaa, 0xb4, 0xcf, 0xc6,
	0x3d, 0x10, 0xbc, 0x4d, 0xa9, 0x07, 0x67, 0xa9, 0xa9, 0x29, 0xcf, 0x87, 0xe2, 0x5b, 0x68, 0x1f,
	0xca, 0xe2, 0x3f, 0xf8, 0xce, 0x3e, 0x69, 0x73, 0xbc, 0xe1, 0x60, 0xc4, 0x00, 0xf8, 0x4e, 0x6d,
	0xe8, 0x00, 0x2e, 0x51, 0xc6, 0xf7, 0xe0, 0x61, 0x4a, 0x4c, 0xb0, 0x52, 0x4e, 0xd1, 0x38, 0x97,
	0xa4, 0x71, 0xda, 0x68, 0xd7, 0x4c, 0x1b, 0xed, 0xb0, 0x32, 0x9d, 0x2f, 0xe8, 0x50, 0x2e, 0x51,
	0xc6, 0x12, 0x3a, 0xd3, 0x63, 0x22, 0xd1, 0x87, 0x56, 0x3a, 0x3b, 0x2e, 0x6c, 0xa5, 0x33, 0x11,
	0x83, 0x57, 0x98, 0xb5, 0x25, 0xc0, 0xaf, 0xef, 0x82, 0xff, 0x4e, 0x52, 0x0f, 0xf7, 0x66, 0x76,
	0xa5, 0x09, 0xd9, 0x95, 0xa4, 0x27, 0x83, 0xf7, 0x7d, 0xc8, 0x3e, 0xf6, 0x21, 0xfb, 0xdc, 0x87,
	0xec, 0xf5, 0x2b, 0xfc, 0xb3, 0x6c, 0xd3, 0x03, 0xb8, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x1f,
	0xd6, 0x55, 0xc9, 0x20, 0x02, 0x00, 0x00,
}
