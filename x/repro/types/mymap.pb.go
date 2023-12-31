// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: maprepro/repro/mymap.proto

package types

import (
	fmt "fmt"
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

type Mymap struct {
	Index    string    `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator  string    `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	InnerMap *InnerMap `protobuf:"bytes,3,opt,name=innerMap,proto3" json:"innerMap,omitempty"`
}

func (m *Mymap) Reset()         { *m = Mymap{} }
func (m *Mymap) String() string { return proto.CompactTextString(m) }
func (*Mymap) ProtoMessage()    {}
func (*Mymap) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d829411b55fc07a, []int{0}
}
func (m *Mymap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mymap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mymap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mymap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mymap.Merge(m, src)
}
func (m *Mymap) XXX_Size() int {
	return m.Size()
}
func (m *Mymap) XXX_DiscardUnknown() {
	xxx_messageInfo_Mymap.DiscardUnknown(m)
}

var xxx_messageInfo_Mymap proto.InternalMessageInfo

func (m *Mymap) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Mymap) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Mymap) GetInnerMap() *InnerMap {
	if m != nil {
		return m.InnerMap
	}
	return nil
}

type ThirdMap struct {
	ThirdMap map[string]int32 `protobuf:"bytes,1,rep,name=thirdMap,proto3" json:"thirdMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *ThirdMap) Reset()         { *m = ThirdMap{} }
func (m *ThirdMap) String() string { return proto.CompactTextString(m) }
func (*ThirdMap) ProtoMessage()    {}
func (*ThirdMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d829411b55fc07a, []int{1}
}
func (m *ThirdMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThirdMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThirdMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ThirdMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThirdMap.Merge(m, src)
}
func (m *ThirdMap) XXX_Size() int {
	return m.Size()
}
func (m *ThirdMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ThirdMap.DiscardUnknown(m)
}

var xxx_messageInfo_ThirdMap proto.InternalMessageInfo

func (m *ThirdMap) GetThirdMap() map[string]int32 {
	if m != nil {
		return m.ThirdMap
	}
	return nil
}

type InnerMap struct {
	InnerMap map[string]SecMap `protobuf:"bytes,1,rep,name=innerMap,proto3" json:"innerMap" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *InnerMap) Reset()         { *m = InnerMap{} }
func (m *InnerMap) String() string { return proto.CompactTextString(m) }
func (*InnerMap) ProtoMessage()    {}
func (*InnerMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d829411b55fc07a, []int{2}
}
func (m *InnerMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InnerMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InnerMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InnerMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerMap.Merge(m, src)
}
func (m *InnerMap) XXX_Size() int {
	return m.Size()
}
func (m *InnerMap) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerMap.DiscardUnknown(m)
}

var xxx_messageInfo_InnerMap proto.InternalMessageInfo

func (m *InnerMap) GetInnerMap() map[string]SecMap {
	if m != nil {
		return m.InnerMap
	}
	return nil
}

type SecMap struct {
	ThirdMaps map[string]ThirdMapList `protobuf:"bytes,1,rep,name=thirdMaps,proto3" json:"thirdMaps" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SecMap) Reset()         { *m = SecMap{} }
func (m *SecMap) String() string { return proto.CompactTextString(m) }
func (*SecMap) ProtoMessage()    {}
func (*SecMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d829411b55fc07a, []int{3}
}
func (m *SecMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecMap.Merge(m, src)
}
func (m *SecMap) XXX_Size() int {
	return m.Size()
}
func (m *SecMap) XXX_DiscardUnknown() {
	xxx_messageInfo_SecMap.DiscardUnknown(m)
}

var xxx_messageInfo_SecMap proto.InternalMessageInfo

func (m *SecMap) GetThirdMaps() map[string]ThirdMapList {
	if m != nil {
		return m.ThirdMaps
	}
	return nil
}

type ThirdMapList struct {
	ThirdMap []ThirdMap `protobuf:"bytes,1,rep,name=thirdMap,proto3" json:"thirdMap"`
}

func (m *ThirdMapList) Reset()         { *m = ThirdMapList{} }
func (m *ThirdMapList) String() string { return proto.CompactTextString(m) }
func (*ThirdMapList) ProtoMessage()    {}
func (*ThirdMapList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d829411b55fc07a, []int{4}
}
func (m *ThirdMapList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThirdMapList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThirdMapList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ThirdMapList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThirdMapList.Merge(m, src)
}
func (m *ThirdMapList) XXX_Size() int {
	return m.Size()
}
func (m *ThirdMapList) XXX_DiscardUnknown() {
	xxx_messageInfo_ThirdMapList.DiscardUnknown(m)
}

var xxx_messageInfo_ThirdMapList proto.InternalMessageInfo

func (m *ThirdMapList) GetThirdMap() []ThirdMap {
	if m != nil {
		return m.ThirdMap
	}
	return nil
}

func init() {
	proto.RegisterType((*Mymap)(nil), "maprepro.repro.Mymap")
	proto.RegisterType((*ThirdMap)(nil), "maprepro.repro.ThirdMap")
	proto.RegisterMapType((map[string]int32)(nil), "maprepro.repro.ThirdMap.ThirdMapEntry")
	proto.RegisterType((*InnerMap)(nil), "maprepro.repro.InnerMap")
	proto.RegisterMapType((map[string]SecMap)(nil), "maprepro.repro.InnerMap.InnerMapEntry")
	proto.RegisterType((*SecMap)(nil), "maprepro.repro.SecMap")
	proto.RegisterMapType((map[string]ThirdMapList)(nil), "maprepro.repro.SecMap.ThirdMapsEntry")
	proto.RegisterType((*ThirdMapList)(nil), "maprepro.repro.ThirdMapList")
}

func init() { proto.RegisterFile("maprepro/repro/mymap.proto", fileDescriptor_2d829411b55fc07a) }

var fileDescriptor_2d829411b55fc07a = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4d, 0x2c, 0x28,
	0x4a, 0x2d, 0x28, 0xca, 0xd7, 0x87, 0x90, 0xb9, 0x95, 0xb9, 0x89, 0x05, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x7c, 0x30, 0x39, 0x3d, 0x30, 0x29, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x96,
	0xd2, 0x07, 0xb1, 0x20, 0xaa, 0x94, 0x72, 0xb9, 0x58, 0x7d, 0x41, 0x9a, 0x84, 0x44, 0xb8, 0x58,
	0x33, 0xf3, 0x52, 0x52, 0x2b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x09,
	0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x26, 0xb0, 0x38, 0x8c, 0x2b, 0x64,
	0xc2, 0xc5, 0x91, 0x99, 0x97, 0x97, 0x5a, 0xe4, 0x9b, 0x58, 0x20, 0xc1, 0xac, 0xc0, 0xa8, 0xc1,
	0x6d, 0x24, 0xa1, 0x87, 0x6a, 0xa3, 0x9e, 0x27, 0x54, 0x3e, 0x08, 0xae, 0x52, 0xa9, 0x9b, 0x91,
	0x8b, 0x23, 0x24, 0x23, 0xb3, 0x28, 0xc5, 0x37, 0xb1, 0x40, 0xc8, 0x89, 0x8b, 0xa3, 0x04, 0xca,
	0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0x43, 0x37, 0x02, 0xa6, 0x16, 0xce, 0x70, 0xcd,
	0x2b, 0x29, 0xaa, 0x0c, 0x82, 0xeb, 0x93, 0xb2, 0xe6, 0xe2, 0x45, 0x91, 0x12, 0x12, 0xe0, 0x62,
	0xce, 0x4e, 0xad, 0x84, 0xfa, 0x02, 0xc4, 0x04, 0xf9, 0xac, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0xec,
	0x03, 0xd6, 0x20, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0x51, 0x69, 0x25, 0x23, 0x17, 0x07, 0xcc, 0x91,
	0x42, 0x1e, 0x48, 0x1e, 0xc2, 0xe1, 0x1a, 0x98, 0x5a, 0x38, 0x03, 0x6c, 0xa5, 0x13, 0xcb, 0x89,
	0x7b, 0xf2, 0x0c, 0x08, 0x4f, 0x4a, 0x05, 0x73, 0xf1, 0xa2, 0x28, 0xc0, 0xe2, 0x26, 0x1d, 0x64,
	0x37, 0x71, 0x1b, 0x89, 0xa1, 0xdb, 0x14, 0x9c, 0x9a, 0x0c, 0x0a, 0x38, 0x24, 0xb7, 0xae, 0x67,
	0xe4, 0x62, 0x83, 0x88, 0x0a, 0x79, 0x72, 0x71, 0xc2, 0xfc, 0x5f, 0x0c, 0x75, 0xaa, 0x2a, 0x76,
	0x03, 0xe0, 0xc1, 0x56, 0x8c, 0xec, 0x52, 0x84, 0x6e, 0xa9, 0x28, 0x2e, 0x3e, 0x54, 0x25, 0x58,
	0xdc, 0x6a, 0x84, 0xea, 0x56, 0x19, 0x5c, 0x71, 0xe4, 0x93, 0x59, 0x5c, 0x82, 0xec, 0x62, 0x2f,
	0x2e, 0x1e, 0x64, 0x29, 0x21, 0x2b, 0x8c, 0xe8, 0x96, 0xc0, 0x65, 0x14, 0x2c, 0x48, 0x61, 0xea,
	0x9d, 0x0c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x0c, 0x9e, 0x05,
	0x2a, 0xa0, 0x99, 0xa0, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0xbe, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x74, 0xe9, 0x60, 0xe4, 0x23, 0x03, 0x00, 0x00,
}

func (m *Mymap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mymap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mymap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InnerMap != nil {
		{
			size, err := m.InnerMap.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMymap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMymap(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintMymap(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ThirdMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThirdMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThirdMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ThirdMap) > 0 {
		for k := range m.ThirdMap {
			v := m.ThirdMap[k]
			baseI := i
			i = encodeVarintMymap(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMymap(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMymap(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *InnerMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InnerMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InnerMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InnerMap) > 0 {
		for k := range m.InnerMap {
			v := m.InnerMap[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMymap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMymap(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMymap(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SecMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ThirdMaps) > 0 {
		for k := range m.ThirdMaps {
			v := m.ThirdMaps[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMymap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMymap(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMymap(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ThirdMapList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThirdMapList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThirdMapList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ThirdMap) > 0 {
		for iNdEx := len(m.ThirdMap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ThirdMap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMymap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMymap(dAtA []byte, offset int, v uint64) int {
	offset -= sovMymap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mymap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovMymap(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMymap(uint64(l))
	}
	if m.InnerMap != nil {
		l = m.InnerMap.Size()
		n += 1 + l + sovMymap(uint64(l))
	}
	return n
}

func (m *ThirdMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ThirdMap) > 0 {
		for k, v := range m.ThirdMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMymap(uint64(len(k))) + 1 + sovMymap(uint64(v))
			n += mapEntrySize + 1 + sovMymap(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *InnerMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.InnerMap) > 0 {
		for k, v := range m.InnerMap {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovMymap(uint64(len(k))) + 1 + l + sovMymap(uint64(l))
			n += mapEntrySize + 1 + sovMymap(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *SecMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ThirdMaps) > 0 {
		for k, v := range m.ThirdMaps {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovMymap(uint64(len(k))) + 1 + l + sovMymap(uint64(l))
			n += mapEntrySize + 1 + sovMymap(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ThirdMapList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ThirdMap) > 0 {
		for _, e := range m.ThirdMap {
			l = e.Size()
			n += 1 + l + sovMymap(uint64(l))
		}
	}
	return n
}

func sovMymap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMymap(x uint64) (n int) {
	return sovMymap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mymap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMymap
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
			return fmt.Errorf("proto: Mymap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mymap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymap
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
				return ErrInvalidLengthMymap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMymap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymap
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
				return ErrInvalidLengthMymap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMymap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymap
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
				return ErrInvalidLengthMymap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMymap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InnerMap == nil {
				m.InnerMap = &InnerMap{}
			}
			if err := m.InnerMap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMymap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMymap
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
func (m *ThirdMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMymap
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
			return fmt.Errorf("proto: ThirdMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThirdMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThirdMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymap
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
				return ErrInvalidLengthMymap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMymap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ThirdMap == nil {
				m.ThirdMap = make(map[string]int32)
			}
			var mapkey string
			var mapvalue int32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMymap
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMymap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMymap
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMymap
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMymap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMymap(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMymap
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ThirdMap[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMymap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMymap
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
func (m *InnerMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMymap
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
			return fmt.Errorf("proto: InnerMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InnerMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymap
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
				return ErrInvalidLengthMymap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMymap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InnerMap == nil {
				m.InnerMap = make(map[string]SecMap)
			}
			var mapkey string
			mapvalue := &SecMap{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMymap
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMymap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMymap
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMymap
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMymap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthMymap
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthMymap
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &SecMap{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMymap(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMymap
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.InnerMap[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMymap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMymap
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
func (m *SecMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMymap
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
			return fmt.Errorf("proto: SecMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThirdMaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymap
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
				return ErrInvalidLengthMymap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMymap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ThirdMaps == nil {
				m.ThirdMaps = make(map[string]ThirdMapList)
			}
			var mapkey string
			mapvalue := &ThirdMapList{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMymap
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMymap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMymap
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMymap
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMymap
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthMymap
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthMymap
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ThirdMapList{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMymap(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMymap
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ThirdMaps[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMymap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMymap
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
func (m *ThirdMapList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMymap
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
			return fmt.Errorf("proto: ThirdMapList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThirdMapList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThirdMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMymap
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
				return ErrInvalidLengthMymap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMymap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThirdMap = append(m.ThirdMap, ThirdMap{})
			if err := m.ThirdMap[len(m.ThirdMap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMymap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMymap
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
func skipMymap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMymap
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
					return 0, ErrIntOverflowMymap
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
					return 0, ErrIntOverflowMymap
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
				return 0, ErrInvalidLengthMymap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMymap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMymap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMymap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMymap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMymap = fmt.Errorf("proto: unexpected end of group")
)
