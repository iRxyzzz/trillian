// Code generated by protoc-gen-go.
// source: trillian.proto
// DO NOT EDIT!

package trillian

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import sigpb "github.com/google/trillian/crypto/sigpb"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines the way empty / node / leaf hashes are constructed incorporating
// preimage protection, which can be application specific.
type HashStrategy int32

const (
	// Hash strategy cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	HashStrategy_UNKNOWN_HASH_STRATEGY HashStrategy = 0
	// Certificate Transparency strategy: leaf hash prefix = 0x00, node prefix =
	// 0x01, empty hash is digest([]byte{}), as defined in the specification.
	HashStrategy_RFC_6962 HashStrategy = 1
)

var HashStrategy_name = map[int32]string{
	0: "UNKNOWN_HASH_STRATEGY",
	1: "RFC_6962",
}
var HashStrategy_value = map[string]int32{
	"UNKNOWN_HASH_STRATEGY": 0,
	"RFC_6962":              1,
}

func (x HashStrategy) String() string {
	return proto.EnumName(HashStrategy_name, int32(x))
}
func (HashStrategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// State of the tree.
type TreeState int32

const (
	// Tree state cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeState_UNKNOWN_TREE_STATE TreeState = 0
	// Active trees are able to respond to both read and write requests.
	TreeState_ACTIVE TreeState = 1
	// Frozen trees are only able to respond to read requests, writing to a frozen
	// tree is forbidden.
	TreeState_FROZEN TreeState = 2
	// Tree was been deleted, therefore is invisible and acts similarly to a
	// non-existing tree for all requests.
	// A soft deleted tree may be undeleted while the soft-deletion period has not
	// passed.
	TreeState_SOFT_DELETED TreeState = 3
	// A hard deleted tree was been definitely deleted and cannot be recovered.
	// Acts an a non-existing tree for all read and write requests, but blocks the
	// tree ID from ever being reused.
	TreeState_HARD_DELETED TreeState = 4
)

var TreeState_name = map[int32]string{
	0: "UNKNOWN_TREE_STATE",
	1: "ACTIVE",
	2: "FROZEN",
	3: "SOFT_DELETED",
	4: "HARD_DELETED",
}
var TreeState_value = map[string]int32{
	"UNKNOWN_TREE_STATE": 0,
	"ACTIVE":             1,
	"FROZEN":             2,
	"SOFT_DELETED":       3,
	"HARD_DELETED":       4,
}

func (x TreeState) String() string {
	return proto.EnumName(TreeState_name, int32(x))
}
func (TreeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

// Type of the tree.
type TreeType int32

const (
	// Tree type cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeType_UNKNOWN_TREE_TYPE TreeType = 0
	// Tree represents a verifiable log.
	TreeType_LOG TreeType = 1
	// Tree represents a verifiable map.
	TreeType_MAP TreeType = 2
)

var TreeType_name = map[int32]string{
	0: "UNKNOWN_TREE_TYPE",
	1: "LOG",
	2: "MAP",
}
var TreeType_value = map[string]int32{
	"UNKNOWN_TREE_TYPE": 0,
	"LOG":               1,
	"MAP":               2,
}

func (x TreeType) String() string {
	return proto.EnumName(TreeType_name, int32(x))
}
func (TreeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

// Represents a tree, which may be either a verifiable log or map.
// Readonly attributes are assigned at tree creation, after which they may not
// be modified.
type Tree struct {
	// ID of the tree.
	// Readonly.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	// State of the tree.
	// Trees are active after creation. At any point the tree may transition
	// between ACTIVE and FROZEN.
	// Deleted trees are set as SOFT_DELETED for a certain time period, after
	// which they'll automatically transition to HARD_DELETED.
	TreeState TreeState `protobuf:"varint,2,opt,name=tree_state,json=treeState,enum=trillian.TreeState" json:"tree_state,omitempty"`
	// Type of the tree.
	// Readonly.
	TreeType TreeType `protobuf:"varint,3,opt,name=tree_type,json=treeType,enum=trillian.TreeType" json:"tree_type,omitempty"`
	// Hash strategy to be used by the tree.
	// Readonly.
	HashStrategy HashStrategy `protobuf:"varint,4,opt,name=hash_strategy,json=hashStrategy,enum=trillian.HashStrategy" json:"hash_strategy,omitempty"`
	// Hash algorithm to be used by the tree.
	// Readonly.
	HashAlgorithm sigpb.DigitallySigned_HashAlgorithm `protobuf:"varint,5,opt,name=hash_algorithm,json=hashAlgorithm,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// Signature algorithm to be used by the tree.
	// Readonly.
	SignatureAlgorithm sigpb.DigitallySigned_SignatureAlgorithm `protobuf:"varint,6,opt,name=signature_algorithm,json=signatureAlgorithm,enum=sigpb.DigitallySigned_SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// Display name of the tree.
	// Optional.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Description of the tree,
	// Optional.
	Description string `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	// Timestamp of tree creation.
	// Readonly.
	// TODO(codingllama): Migrate public APIs to google.protobuf.Timestamp.
	CreateTimeMillisSinceEpoch int64 `protobuf:"varint,10,opt,name=create_time_millis_since_epoch,json=createTimeMillisSinceEpoch" json:"create_time_millis_since_epoch,omitempty"`
	// Timestamp of last tree update.
	// Readonly (automatically assigned on updates).
	UpdateTimeMillisSinceEpoch int64 `protobuf:"varint,11,opt,name=update_time_millis_since_epoch,json=updateTimeMillisSinceEpoch" json:"update_time_millis_since_epoch,omitempty"`
	// Identifies the private key used for signing tree heads and entry
	// timestamps.
	// This can be any type of message to accommodate different key management
	// systems, e.g. PEM files, HSMs, etc.
	// Private keys are write-only: they're never returned by RPCs.
	// TODO(RJPercival): Implement sufficient validation to allow this field to be
	// mutable. It should be mutable in the sense that the key can be migrated to
	// a different key management system, but the key itself should never change.
	PrivateKey *google_protobuf.Any `protobuf:"bytes,12,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// Storage-specific settings.
	// Varies according to the storage implementation backing Trillian.
	StorageSettings *google_protobuf.Any `protobuf:"bytes,13,opt,name=storage_settings,json=storageSettings" json:"storage_settings,omitempty"`
	// The public key used for verifying tree heads and entry timestamps.
	// Readonly.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,14,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
}

func (m *Tree) Reset()                    { *m = Tree{} }
func (m *Tree) String() string            { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()               {}
func (*Tree) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Tree) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *Tree) GetTreeState() TreeState {
	if m != nil {
		return m.TreeState
	}
	return TreeState_UNKNOWN_TREE_STATE
}

func (m *Tree) GetTreeType() TreeType {
	if m != nil {
		return m.TreeType
	}
	return TreeType_UNKNOWN_TREE_TYPE
}

func (m *Tree) GetHashStrategy() HashStrategy {
	if m != nil {
		return m.HashStrategy
	}
	return HashStrategy_UNKNOWN_HASH_STRATEGY
}

func (m *Tree) GetHashAlgorithm() sigpb.DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return sigpb.DigitallySigned_NONE
}

func (m *Tree) GetSignatureAlgorithm() sigpb.DigitallySigned_SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return sigpb.DigitallySigned_ANONYMOUS
}

func (m *Tree) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Tree) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tree) GetCreateTimeMillisSinceEpoch() int64 {
	if m != nil {
		return m.CreateTimeMillisSinceEpoch
	}
	return 0
}

func (m *Tree) GetUpdateTimeMillisSinceEpoch() int64 {
	if m != nil {
		return m.UpdateTimeMillisSinceEpoch
	}
	return 0
}

func (m *Tree) GetPrivateKey() *google_protobuf.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *Tree) GetStorageSettings() *google_protobuf.Any {
	if m != nil {
		return m.StorageSettings
	}
	return nil
}

func (m *Tree) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type SignedEntryTimestamp struct {
	TimestampNanos int64                  `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	LogId          int64                  `protobuf:"varint,2,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	Signature      *sigpb.DigitallySigned `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *SignedEntryTimestamp) Reset()                    { *m = SignedEntryTimestamp{} }
func (m *SignedEntryTimestamp) String() string            { return proto.CompactTextString(m) }
func (*SignedEntryTimestamp) ProtoMessage()               {}
func (*SignedEntryTimestamp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SignedEntryTimestamp) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedEntryTimestamp) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *SignedEntryTimestamp) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

// SignedLogRoot represents a commitment by a Log to a particular tree.
type SignedLogRoot struct {
	// epoch nanoseconds, good until 2500ish
	TimestampNanos int64  `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	RootHash       []byte `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// TreeSize is the number of entries in the tree.
	TreeSize int64 `protobuf:"varint,3,opt,name=tree_size,json=treeSize" json:"tree_size,omitempty"`
	// TODO(al): define serialized format for the signature scheme.
	Signature    *sigpb.DigitallySigned `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	LogId        int64                  `protobuf:"varint,5,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	TreeRevision int64                  `protobuf:"varint,6,opt,name=tree_revision,json=treeRevision" json:"tree_revision,omitempty"`
}

func (m *SignedLogRoot) Reset()                    { *m = SignedLogRoot{} }
func (m *SignedLogRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedLogRoot) ProtoMessage()               {}
func (*SignedLogRoot) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *SignedLogRoot) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedLogRoot) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *SignedLogRoot) GetTreeSize() int64 {
	if m != nil {
		return m.TreeSize
	}
	return 0
}

func (m *SignedLogRoot) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedLogRoot) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *SignedLogRoot) GetTreeRevision() int64 {
	if m != nil {
		return m.TreeRevision
	}
	return 0
}

type MapperMetadata struct {
	SourceLogId                  []byte `protobuf:"bytes,1,opt,name=source_log_id,json=sourceLogId,proto3" json:"source_log_id,omitempty"`
	HighestFullyCompletedSeq     int64  `protobuf:"varint,2,opt,name=highest_fully_completed_seq,json=highestFullyCompletedSeq" json:"highest_fully_completed_seq,omitempty"`
	HighestPartiallyCompletedSeq int64  `protobuf:"varint,3,opt,name=highest_partially_completed_seq,json=highestPartiallyCompletedSeq" json:"highest_partially_completed_seq,omitempty"`
}

func (m *MapperMetadata) Reset()                    { *m = MapperMetadata{} }
func (m *MapperMetadata) String() string            { return proto.CompactTextString(m) }
func (*MapperMetadata) ProtoMessage()               {}
func (*MapperMetadata) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *MapperMetadata) GetSourceLogId() []byte {
	if m != nil {
		return m.SourceLogId
	}
	return nil
}

func (m *MapperMetadata) GetHighestFullyCompletedSeq() int64 {
	if m != nil {
		return m.HighestFullyCompletedSeq
	}
	return 0
}

func (m *MapperMetadata) GetHighestPartiallyCompletedSeq() int64 {
	if m != nil {
		return m.HighestPartiallyCompletedSeq
	}
	return 0
}

// SignedMapRoot represents a commitment by a Map to a particular tree.
type SignedMapRoot struct {
	TimestampNanos int64           `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	RootHash       []byte          `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	Metadata       *MapperMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	// TODO(al): define serialized format for the signature scheme.
	Signature   *sigpb.DigitallySigned `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	MapId       int64                  `protobuf:"varint,5,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	MapRevision int64                  `protobuf:"varint,6,opt,name=map_revision,json=mapRevision" json:"map_revision,omitempty"`
}

func (m *SignedMapRoot) Reset()                    { *m = SignedMapRoot{} }
func (m *SignedMapRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedMapRoot) ProtoMessage()               {}
func (*SignedMapRoot) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *SignedMapRoot) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedMapRoot) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *SignedMapRoot) GetMetadata() *MapperMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SignedMapRoot) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedMapRoot) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *SignedMapRoot) GetMapRevision() int64 {
	if m != nil {
		return m.MapRevision
	}
	return 0
}

func init() {
	proto.RegisterType((*Tree)(nil), "trillian.Tree")
	proto.RegisterType((*SignedEntryTimestamp)(nil), "trillian.SignedEntryTimestamp")
	proto.RegisterType((*SignedLogRoot)(nil), "trillian.SignedLogRoot")
	proto.RegisterType((*MapperMetadata)(nil), "trillian.MapperMetadata")
	proto.RegisterType((*SignedMapRoot)(nil), "trillian.SignedMapRoot")
	proto.RegisterEnum("trillian.HashStrategy", HashStrategy_name, HashStrategy_value)
	proto.RegisterEnum("trillian.TreeState", TreeState_name, TreeState_value)
	proto.RegisterEnum("trillian.TreeType", TreeType_name, TreeType_value)
}

func init() { proto.RegisterFile("trillian.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 916 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0xf2, 0xe3, 0xd8, 0xc7, 0x3f, 0x55, 0xd9, 0x26, 0x53, 0xd3, 0x61, 0xf3, 0xbc, 0x01,
	0xcb, 0x72, 0x61, 0x0f, 0x4e, 0xdb, 0x61, 0x18, 0x86, 0xc1, 0x4b, 0x94, 0x26, 0x4b, 0xec, 0x04,
	0x92, 0xb6, 0xa1, 0xbd, 0x21, 0x68, 0x8b, 0x95, 0x89, 0x4a, 0x22, 0x2b, 0xd2, 0x05, 0xd4, 0x67,
	0xd8, 0x13, 0xed, 0x79, 0x06, 0xec, 0x21, 0x76, 0x33, 0x90, 0x92, 0x6c, 0x67, 0xed, 0xba, 0x62,
	0xe8, 0x8d, 0xcd, 0xf3, 0x9d, 0xef, 0xfb, 0xf8, 0x73, 0x0e, 0x29, 0xe8, 0xa8, 0x8c, 0xc5, 0x31,
	0x23, 0x69, 0x5f, 0x64, 0x5c, 0x71, 0x54, 0xaf, 0xe2, 0xfd, 0x47, 0x11, 0x53, 0xf3, 0xc5, 0xb4,
	0x3f, 0xe3, 0xc9, 0x20, 0xe2, 0x3c, 0x8a, 0xe9, 0xa0, 0xca, 0x0d, 0x66, 0x59, 0x2e, 0x14, 0x1f,
	0xbc, 0xa0, 0xb9, 0x14, 0xd3, 0xf2, 0xaf, 0x30, 0xd8, 0x3f, 0xfa, 0x6f, 0x99, 0x64, 0x91, 0x98,
	0x16, 0xbf, 0xa5, 0xe8, 0x7e, 0xc9, 0x34, 0xd1, 0x74, 0xf1, 0x7c, 0x40, 0xd2, 0xbc, 0x48, 0xf5,
	0xfe, 0xdc, 0x86, 0xad, 0x20, 0xa3, 0x14, 0x7d, 0x04, 0x3b, 0x2a, 0xa3, 0x14, 0xb3, 0xd0, 0xb1,
	0xba, 0xd6, 0xc1, 0xa6, 0x57, 0xd3, 0xe1, 0x79, 0x88, 0x86, 0x00, 0x26, 0x21, 0x15, 0x51, 0xd4,
	0xd9, 0xe8, 0x5a, 0x07, 0x9d, 0xe1, 0xdd, 0xfe, 0x72, 0x5f, 0x5a, 0xec, 0xeb, 0x94, 0xd7, 0x50,
	0xd5, 0x10, 0x0d, 0xc0, 0x04, 0x58, 0xe5, 0x82, 0x3a, 0x9b, 0x46, 0x82, 0x6e, 0x4a, 0x82, 0x5c,
	0x50, 0xaf, 0xae, 0xca, 0x11, 0xfa, 0x0e, 0xda, 0x73, 0x22, 0xe7, 0x58, 0xaa, 0x8c, 0x28, 0x1a,
	0xe5, 0xce, 0x96, 0x11, 0xed, 0xad, 0x44, 0x67, 0x44, 0xce, 0xfd, 0x32, 0xeb, 0xb5, 0xe6, 0x6b,
	0x11, 0xba, 0x80, 0x8e, 0x11, 0x93, 0x38, 0xe2, 0x19, 0x53, 0xf3, 0xc4, 0xd9, 0x36, 0xea, 0x2f,
	0xfa, 0xc5, 0x21, 0x9c, 0xb0, 0x88, 0x29, 0x12, 0xc7, 0xb9, 0xcf, 0xa2, 0x94, 0x86, 0xc6, 0x6a,
	0x54, 0x71, 0x3d, 0x33, 0xf1, 0x32, 0x44, 0xcf, 0xe0, 0xae, 0x64, 0x51, 0x4a, 0xd4, 0x22, 0xa3,
	0x6b, 0x8e, 0x35, 0xe3, 0xf8, 0xd5, 0xbf, 0x38, 0xfa, 0x95, 0x62, 0x65, 0x8b, 0xe4, 0x1b, 0x18,
	0xfa, 0x0c, 0x5a, 0x21, 0x93, 0x22, 0x26, 0x39, 0x4e, 0x49, 0x42, 0x9d, 0x7a, 0xd7, 0x3a, 0x68,
	0x78, 0xcd, 0x12, 0x9b, 0x90, 0x84, 0xa2, 0x2e, 0x34, 0x43, 0x2a, 0x67, 0x19, 0x13, 0x8a, 0xf1,
	0xd4, 0x69, 0x94, 0x8c, 0x15, 0x84, 0x7e, 0x84, 0x4f, 0x66, 0x19, 0x25, 0x8a, 0x62, 0xc5, 0x12,
	0x8a, 0x13, 0x7d, 0x3e, 0x12, 0x4b, 0x96, 0xce, 0x28, 0xa6, 0x82, 0xcf, 0xe6, 0x0e, 0x98, 0xfa,
	0xed, 0x17, 0xac, 0x80, 0x25, 0x74, 0x6c, 0x38, 0xbe, 0xa6, 0xb8, 0x9a, 0xa1, 0x3d, 0x16, 0x22,
	0x7c, 0x97, 0x47, 0xb3, 0xf0, 0x28, 0x58, 0x6f, 0xf5, 0x78, 0x04, 0x4d, 0x91, 0xb1, 0x57, 0xda,
	0xe4, 0x05, 0xcd, 0x9d, 0x56, 0xd7, 0x3a, 0x68, 0x0e, 0xef, 0xf5, 0x8b, 0x56, 0xeb, 0x57, 0xad,
	0xd6, 0x1f, 0xa5, 0xb9, 0x07, 0x25, 0xf1, 0x82, 0xe6, 0xe8, 0x07, 0xb0, 0xa5, 0xe2, 0x19, 0x89,
	0x28, 0x96, 0x54, 0x29, 0x96, 0x46, 0xd2, 0x69, 0xbf, 0x43, 0x7b, 0xbb, 0x64, 0xfb, 0x25, 0x19,
	0x7d, 0x0d, 0x20, 0x16, 0xd3, 0x98, 0xcd, 0xcc, 0xb4, 0x1d, 0x23, 0xbd, 0xd3, 0x2f, 0x2f, 0xc9,
	0xb5, 0xc9, 0x5c, 0xd0, 0xdc, 0x6b, 0x88, 0x6a, 0xf8, 0xd3, 0x56, 0x7d, 0xc7, 0xae, 0xf7, 0x7e,
	0xb3, 0xe0, 0x5e, 0x51, 0x2e, 0x37, 0x55, 0x59, 0xae, 0xf7, 0x24, 0x15, 0x49, 0x04, 0xfa, 0x12,
	0x6e, 0xab, 0x2a, 0xc0, 0x29, 0x49, 0xb9, 0x2c, 0x6f, 0x40, 0x67, 0x09, 0x4f, 0x34, 0x8a, 0x76,
	0xa1, 0x16, 0xf3, 0x48, 0xdf, 0x90, 0x0d, 0x93, 0xdf, 0x8e, 0x79, 0x74, 0x1e, 0xa2, 0x87, 0xd0,
	0x58, 0xd6, 0xda, 0x34, 0x7b, 0x73, 0xb8, 0xf7, 0xf6, 0x3e, 0xf1, 0x56, 0xc4, 0xde, 0x1f, 0x16,
	0xb4, 0x0b, 0xf4, 0x92, 0x47, 0x1e, 0xe7, 0xea, 0xfd, 0xd7, 0xf1, 0x00, 0x1a, 0x19, 0xe7, 0x0a,
	0xeb, 0xc6, 0x35, 0x4b, 0x69, 0x79, 0x75, 0x0d, 0xe8, 0xbe, 0xd6, 0xc9, 0xe2, 0xba, 0xb2, 0xd7,
	0xc5, 0x6a, 0x36, 0x8b, 0x6b, 0xe6, 0xb3, 0xd7, 0xf4, 0xe6, 0x52, 0xb7, 0xde, 0x73, 0xa9, 0x6b,
	0xfb, 0xde, 0x5e, 0xdf, 0xf7, 0xe7, 0xd0, 0x36, 0x33, 0x65, 0xf4, 0x15, 0x93, 0xba, 0x59, 0x6b,
	0x26, 0xdb, 0xd2, 0xa0, 0x57, 0x62, 0xbd, 0xdf, 0x2d, 0xe8, 0x8c, 0x89, 0x10, 0x34, 0x1b, 0x53,
	0x45, 0x42, 0xa2, 0x08, 0xea, 0x41, 0x5b, 0xf2, 0x45, 0x36, 0xa3, 0xb8, 0x74, 0xb5, 0xcc, 0x16,
	0x9a, 0x05, 0x78, 0x69, 0xbc, 0xbf, 0x87, 0x07, 0x73, 0x16, 0xcd, 0xa9, 0x54, 0xf8, 0xf9, 0x22,
	0x8e, 0x73, 0x3c, 0xe3, 0x89, 0x88, 0xa9, 0xa2, 0x21, 0x96, 0xf4, 0x65, 0x79, 0xfe, 0x4e, 0x49,
	0x39, 0xd5, 0x8c, 0xe3, 0x8a, 0xe0, 0xd3, 0x97, 0xc8, 0x85, 0x4f, 0x2b, 0xb9, 0x20, 0x99, 0x62,
	0xe4, 0x4d, 0x8b, 0xe2, 0x68, 0x3e, 0x2e, 0x69, 0xd7, 0x15, 0x6b, 0xdd, 0xa6, 0xf7, 0xd7, 0xb2,
	0x46, 0x63, 0x22, 0x3e, 0x60, 0x8d, 0x1e, 0x42, 0x3d, 0x29, 0x4f, 0xa3, 0x6c, 0x18, 0x67, 0xf5,
	0xd0, 0xdd, 0x3c, 0x2d, 0x6f, 0xc9, 0xfc, 0xff, 0xc5, 0x4b, 0x88, 0x58, 0x2b, 0x5e, 0x42, 0xc4,
	0x79, 0xa8, 0x9f, 0x22, 0x0d, 0xff, 0xa3, 0x76, 0xcd, 0x84, 0x88, 0xaa, 0x74, 0x87, 0xdf, 0x40,
	0x6b, 0xfd, 0xd1, 0x45, 0xf7, 0x61, 0xf7, 0xe7, 0xc9, 0xc5, 0xe4, 0xea, 0xd7, 0x09, 0x3e, 0x1b,
	0xf9, 0x67, 0xd8, 0x0f, 0xbc, 0x51, 0xe0, 0x3e, 0x79, 0x6a, 0xdf, 0x42, 0x2d, 0xa8, 0x7b, 0xa7,
	0xc7, 0xf8, 0xf1, 0xb7, 0x8f, 0x87, 0xb6, 0x75, 0x88, 0xa1, 0xb1, 0xfc, 0x2a, 0xa0, 0x3d, 0x40,
	0x95, 0x2a, 0xf0, 0x5c, 0x17, 0xfb, 0xc1, 0x28, 0x70, 0xed, 0x5b, 0x08, 0xa0, 0x36, 0x3a, 0x0e,
	0xce, 0x7f, 0x71, 0x6d, 0x4b, 0x8f, 0x4f, 0xbd, 0xab, 0x67, 0xee, 0xc4, 0xde, 0x40, 0x36, 0xb4,
	0xfc, 0xab, 0xd3, 0x00, 0x9f, 0xb8, 0x97, 0x6e, 0xe0, 0x9e, 0xd8, 0x9b, 0x1a, 0x39, 0x1b, 0x79,
	0x27, 0x4b, 0x64, 0xeb, 0xf0, 0x08, 0xea, 0xd5, 0x37, 0x04, 0xed, 0xc2, 0x9d, 0x1b, 0xfe, 0xc1,
	0xd3, 0x6b, 0x6d, 0xbf, 0x03, 0x9b, 0x97, 0x57, 0x4f, 0x6c, 0x4b, 0x0f, 0xc6, 0xa3, 0x6b, 0x7b,
	0x63, 0x5a, 0x33, 0xcf, 0xca, 0xd1, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x52, 0x45, 0xb4,
	0x93, 0x07, 0x00, 0x00,
}
