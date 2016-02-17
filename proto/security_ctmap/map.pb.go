// Code generated by protoc-gen-go.
// source: proto/security_ctmap/map.proto
// DO NOT EDIT!

/*
Package security_ctmap is a generated protocol buffer package.

It is generated from these files:
	proto/security_ctmap/map.proto

It has these top-level messages:
	EpochHead
	Step
	Entry
	PublicKey
	SignedEpochHead
	SignedEntryUpdate
*/
package security_ctmap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/google/e2e-key-server/proto/google_protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// EpochHead is the head node of the Merkle Tree as well as additional metadata
// for the tree.
type EpochHead struct {
	// Realm is the domain...
	Realm string `protobuf:"bytes,1,opt,name=realm" json:"realm,omitempty"`
	// Epoch number
	Epoch int64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	// Root is the value of the root node of the merkle tree.
	Root []byte `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// IssueTime is the time when this epoch was released. All epochs for the
	// same keyserver MUST have non-decreasing IssueTimes.
	IssueTime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_time,json=issueTime" json:"issue_time,omitempty"`
	// Hash of previous SEH. SHA512_256.
	PreviousHash []byte `protobuf:"bytes,5,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
}

func (m *EpochHead) Reset()                    { *m = EpochHead{} }
func (m *EpochHead) String() string            { return proto.CompactTextString(m) }
func (*EpochHead) ProtoMessage()               {}
func (*EpochHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EpochHead) GetIssueTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.IssueTime
	}
	return nil
}

// Step is a combined, ordered list of SignedEntryUpdates and SignedEpochHeads
// which are made available to verifiers.
type Step struct {
	// Types that are valid to be assigned to Type:
	//	*Step_EntryChanged
	//	*Step_SignedEpochHead
	Type isStep_Type `protobuf_oneof:"type"`
	// epoch of this udpate.
	Epoch int64 `protobuf:"varint,3,opt,name=epoch" json:"epoch,omitempty"`
	// commitment_timestamp is the ordered commitment_timestamp of this step.
	CommitmentTimestamp int64 `protobuf:"varint,4,opt,name=commitment_timestamp,json=commitmentTimestamp" json:"commitment_timestamp,omitempty"`
}

func (m *Step) Reset()                    { *m = Step{} }
func (m *Step) String() string            { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()               {}
func (*Step) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isStep_Type interface {
	isStep_Type()
}

type Step_EntryChanged struct {
	EntryChanged []byte `protobuf:"bytes,1,opt,name=entry_changed,json=entryChanged,proto3,oneof"`
}
type Step_SignedEpochHead struct {
	SignedEpochHead *SignedEpochHead `protobuf:"bytes,2,opt,name=signed_epoch_head,json=signedEpochHead,oneof"`
}

func (*Step_EntryChanged) isStep_Type()    {}
func (*Step_SignedEpochHead) isStep_Type() {}

func (m *Step) GetType() isStep_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Step) GetEntryChanged() []byte {
	if x, ok := m.GetType().(*Step_EntryChanged); ok {
		return x.EntryChanged
	}
	return nil
}

func (m *Step) GetSignedEpochHead() *SignedEpochHead {
	if x, ok := m.GetType().(*Step_SignedEpochHead); ok {
		return x.SignedEpochHead
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Step) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Step_OneofMarshaler, _Step_OneofUnmarshaler, _Step_OneofSizer, []interface{}{
		(*Step_EntryChanged)(nil),
		(*Step_SignedEpochHead)(nil),
	}
}

func _Step_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Step)
	// type
	switch x := m.Type.(type) {
	case *Step_EntryChanged:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EntryChanged)
	case *Step_SignedEpochHead:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SignedEpochHead); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Step.Type has unexpected type %T", x)
	}
	return nil
}

func _Step_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Step)
	switch tag {
	case 1: // type.entry_changed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Type = &Step_EntryChanged{x}
		return true, err
	case 2: // type.signed_epoch_head
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignedEpochHead)
		err := b.DecodeMessage(msg)
		m.Type = &Step_SignedEpochHead{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Step_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Step)
	// type
	switch x := m.Type.(type) {
	case *Step_EntryChanged:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EntryChanged)))
		n += len(x.EntryChanged)
	case *Step_SignedEpochHead:
		s := proto.Size(x.SignedEpochHead)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Entry is the leaf node object for the Merkle Tree. Its unique index in the
// tree is identified by a hash of an verifiable unpredictable function on the
// user_id.
type Entry struct {
	// Index is the location in the merkle tree for this entry.
	// If signing keys are not unique per user, we need to tie updates to a
	// particular profile.
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// UpdateCount prevents replaying old signed EntryUpdates.
	// not nessesarilly incremented by only one each update.
	UpdateCount uint64 `protobuf:"varint,2,opt,name=update_count,json=updateCount" json:"update_count,omitempty"`
	// EntryKey allows verifiers to validate updates to Entry.
	EntryKey []*PublicKey `protobuf:"bytes,3,rep,name=entry_key,json=entryKey" json:"entry_key,omitempty"`
	// profile_commitment is a cryptographic commitment to the Profile of the form
	// HMAC(profile_commitment_key, serialized_profile)
	ProfileCommitment []byte `protobuf:"bytes,4,opt,name=profile_commitment,json=profileCommitment,proto3" json:"profile_commitment,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetEntryKey() []*PublicKey {
	if m != nil {
		return m.EntryKey
	}
	return nil
}

// PublicKey defines a key this domain uses to sign EpochHeads with.
type PublicKey struct {
	// KeyFormats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_2048
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_2048 struct {
	RsaVerifyingSha256_2048 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_2048,json=rsaVerifyingSha2562048,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_2048) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_2048() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_2048); ok {
		return x.RsaVerifyingSha256_2048
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_2048)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_2048
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_2048{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_2048)))
		n += len(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// SignedEpochHead represents a signed state of the Merkel tree.
type SignedEpochHead struct {
	// Serialized EpochHead.
	EpochHead []byte `protobuf:"bytes,1,opt,name=epoch_head,json=epochHead,proto3" json:"epoch_head,omitempty"`
	// Signature of head, using the signature type of the key.
	// keyed by the first 64 bits bytes of the hash of the key.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SignedEpochHead) Reset()                    { *m = SignedEpochHead{} }
func (m *SignedEpochHead) String() string            { return proto.CompactTextString(m) }
func (*SignedEpochHead) ProtoMessage()               {}
func (*SignedEpochHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SignedEpochHead) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// SignedEntryUpdate is what users provide to update their profiles.
// A serialized SignedEntryUpdate is used as the leaf value in the MerkleTree.
type SignedEntryUpdate struct {
	// NewEntry is the serialized protobuf Entry.
	NewEntry []byte `protobuf:"bytes,1,opt,name=new_entry,json=newEntry,proto3" json:"new_entry,omitempty"`
	// Signature of entry, by the entry_key inside entry AND the old key from the
	// previous epoch. The first proves ownership of new epoch key, and the
	// second proves the the correct owner is making this change.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SignedEntryUpdate) Reset()                    { *m = SignedEntryUpdate{} }
func (m *SignedEntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*SignedEntryUpdate) ProtoMessage()               {}
func (*SignedEntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SignedEntryUpdate) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func init() {
	proto.RegisterType((*EpochHead)(nil), "security_ctmap.EpochHead")
	proto.RegisterType((*Step)(nil), "security_ctmap.Step")
	proto.RegisterType((*Entry)(nil), "security_ctmap.Entry")
	proto.RegisterType((*PublicKey)(nil), "security_ctmap.PublicKey")
	proto.RegisterType((*SignedEpochHead)(nil), "security_ctmap.SignedEpochHead")
	proto.RegisterType((*SignedEntryUpdate)(nil), "security_ctmap.SignedEntryUpdate")
}

var fileDescriptor0 = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xb9, 0x49, 0xfb, 0x6b, 0xa6, 0x29, 0xa5, 0x4b, 0x85, 0xd2, 0x20, 0xfe, 0x05, 0x21,
	0x71, 0xa9, 0x4d, 0x0d, 0xad, 0x5a, 0x10, 0x97, 0x56, 0x45, 0x95, 0x10, 0xa2, 0x6c, 0x81, 0xeb,
	0xca, 0xb1, 0xa7, 0xb1, 0x45, 0xfc, 0x47, 0xde, 0x75, 0x21, 0xdf, 0x07, 0x89, 0x03, 0x5f, 0x02,
	0x71, 0xe7, 0x3b, 0x31, 0xbb, 0x1b, 0x3b, 0x49, 0x05, 0xe2, 0xc4, 0x21, 0x92, 0xe7, 0xbd, 0xd9,
	0xd9, 0xf7, 0xde, 0x6c, 0xe0, 0x4e, 0x51, 0xe6, 0x2a, 0xf7, 0x24, 0x86, 0x55, 0x99, 0xa8, 0x89,
	0x08, 0x55, 0x1a, 0x14, 0x1e, 0xfd, 0x5c, 0x43, 0xb0, 0x6b, 0x8b, 0x4c, 0xff, 0xe5, 0x28, 0x51,
	0x71, 0x35, 0x74, 0xc3, 0x3c, 0xf5, 0x46, 0x79, 0x3e, 0x1a, 0xa3, 0x87, 0x3e, 0xee, 0x7c, 0xc4,
	0xc9, 0x8e, 0xc4, 0xf2, 0x12, 0x4b, 0xcf, 0x0e, 0xb4, 0x9c, 0x30, 0xc5, 0xb0, 0xba, 0xf0, 0x54,
	0x92, 0xa2, 0x54, 0x41, 0x3a, 0x9d, 0x3b, 0xf8, 0xe6, 0x40, 0xe7, 0xa4, 0xc8, 0xc3, 0xf8, 0x14,
	0x83, 0x88, 0x6d, 0xc1, 0x72, 0x89, 0xc1, 0x38, 0xed, 0x39, 0xf7, 0x9c, 0x47, 0x1d, 0x6e, 0x0b,
	0x8d, 0xa2, 0x6e, 0xe9, 0x2d, 0x11, 0xda, 0xe2, 0xb6, 0x60, 0x0c, 0xda, 0x65, 0x9e, 0xab, 0x5e,
	0x8b, 0xc0, 0x2e, 0x37, 0xdf, 0xec, 0x10, 0x20, 0x91, 0xb2, 0x42, 0xa1, 0xaf, 0xe9, 0xb5, 0x89,
	0x59, 0xf3, 0xfb, 0xae, 0xd5, 0xe0, 0xd6, 0x1a, 0xdc, 0x77, 0xb5, 0x06, 0xde, 0x31, 0xdd, 0xba,
	0x66, 0x0f, 0x60, 0xbd, 0x28, 0xf1, 0x32, 0xc9, 0x2b, 0x29, 0xe2, 0x40, 0xc6, 0xbd, 0x65, 0x33,
	0xb7, 0x5b, 0x83, 0xa7, 0x84, 0x0d, 0x7e, 0x3a, 0xd0, 0x3e, 0x57, 0x58, 0xb0, 0x87, 0xb0, 0x8e,
	0x99, 0x2a, 0x29, 0x8d, 0x38, 0xc8, 0x46, 0x18, 0x19, 0xc1, 0xdd, 0xd3, 0xff, 0x78, 0xd7, 0xc0,
	0xc7, 0x16, 0x65, 0xaf, 0x61, 0x53, 0x26, 0xa3, 0x0c, 0x23, 0x61, 0x34, 0x8b, 0x98, 0x4c, 0x1a,
	0x17, 0x6b, 0xfe, 0x5d, 0x77, 0x31, 0x51, 0xf7, 0xdc, 0x34, 0x36, 0x59, 0xd0, 0xac, 0x0d, 0xb9,
	0x08, 0xcd, 0x82, 0x68, 0xcd, 0x07, 0xb1, 0x0b, 0x5b, 0xb4, 0x85, 0x34, 0x51, 0x29, 0xdd, 0x2d,
	0x9a, 0x80, 0x8d, 0xfd, 0x16, 0xbf, 0x31, 0xe3, 0x1a, 0xdf, 0x47, 0x2b, 0xd0, 0x56, 0x93, 0x02,
	0x07, 0x5f, 0x1d, 0x58, 0x3e, 0xd1, 0x82, 0xf5, 0xe8, 0x24, 0x8b, 0xf0, 0xb3, 0x35, 0xc2, 0x6d,
	0xc1, 0xee, 0x43, 0xb7, 0x2a, 0xa2, 0x40, 0xa1, 0x08, 0xf3, 0x2a, 0x53, 0x46, 0x7a, 0x9b, 0xaf,
	0x59, 0xec, 0x58, 0x43, 0x6c, 0x1f, 0x3a, 0x36, 0x09, 0x5a, 0x3d, 0xe9, 0x6a, 0x91, 0xb5, 0xed,
	0xab, 0xd6, 0xce, 0xaa, 0xe1, 0x38, 0x09, 0x5f, 0xe1, 0x84, 0xaf, 0x9a, 0x5e, 0xfa, 0x62, 0x3b,
	0xc0, 0x68, 0x21, 0x17, 0xc9, 0x58, 0xcf, 0xae, 0x15, 0x1a, 0xcd, 0x5d, 0xbe, 0x39, 0x65, 0x8e,
	0x1b, 0x62, 0xf0, 0x85, 0xde, 0x49, 0x33, 0x86, 0xf5, 0xe1, 0x7f, 0x8c, 0xfc, 0xbd, 0xbd, 0xdd,
	0xc3, 0x26, 0xf8, 0x1a, 0x60, 0xcf, 0x61, 0xbb, 0x94, 0x81, 0xa0, 0x27, 0x98, 0x5c, 0x4c, 0x92,
	0x6c, 0x24, 0x64, 0x1c, 0xf8, 0x7b, 0xfb, 0xc2, 0x7f, 0xfc, 0xf4, 0xc0, 0x18, 0xd0, 0xdd, 0x37,
	0xa9, 0xe5, 0x43, 0xdd, 0x71, 0x6e, 0x1a, 0x34, 0xcf, 0x7c, 0xd8, 0xc2, 0x30, 0x5a, 0x38, 0x5e,
	0x10, 0x67, 0x1f, 0x19, 0x9d, 0x63, 0x86, 0x6d, 0x4e, 0x9e, 0x11, 0x77, 0x04, 0xb0, 0x4a, 0xde,
	0x85, 0x09, 0xf4, 0xbb, 0x03, 0x1b, 0x57, 0x16, 0xc9, 0x6e, 0x03, 0xcc, 0x6d, 0xdf, 0xe6, 0xdb,
	0xc1, 0x86, 0x7e, 0x03, 0xa0, 0xf7, 0x1c, 0xa8, 0xaa, 0x44, 0x49, 0x02, 0x75, 0x82, 0xde, 0x5f,
	0x1e, 0x87, 0xa9, 0xed, 0x09, 0xb3, 0x3e, 0x3e, 0x37, 0xa2, 0xff, 0xc2, 0x4a, 0x98, 0xa3, 0xd9,
	0x75, 0x68, 0xe9, 0xf5, 0xe8, 0xbb, 0x57, 0xb8, 0xfe, 0xd4, 0xfb, 0xbe, 0x0c, 0xc6, 0x15, 0xda,
	0x44, 0xb8, 0x2d, 0x9e, 0x2d, 0x1d, 0x38, 0x83, 0x1f, 0x0e, 0x6c, 0x4e, 0xaf, 0xd3, 0x67, 0xdf,
	0x9b, 0x5d, 0xb3, 0x5b, 0xd0, 0xc9, 0xf0, 0x93, 0x30, 0xeb, 0x9b, 0x7a, 0x58, 0x25, 0xc0, 0x8e,
	0x7f, 0xfb, 0x1b, 0x0b, 0xbb, 0x7f, 0xb0, 0x30, 0x9b, 0xf9, 0x0f, 0x4d, 0x0c, 0x57, 0xcc, 0x9f,
	0xfd, 0xc9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x9b, 0x1b, 0x05, 0xd7, 0x04, 0x00, 0x00,
}