// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v2alpha/clusters.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/go/apis/envoy/api/v2/core"
	_type "github.com/datawire/ambassador/go/apis/envoy/type"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Admin endpoint uses this wrapper for `/clusters` to display cluster status information.
// See :ref:`/clusters <operations_admin_interface_clusters>` for more information.
type Clusters struct {
	// Mapping from cluster name to each cluster's status.
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{0}
}
func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return m.Size()
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

// Details an individual cluster's current status.
type ClusterStatus struct {
	// Name of the cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Denotes whether this cluster was added via API or configured statically.
	AddedViaApi bool `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	// The success rate threshold used in the last interval. The threshold is used to eject hosts
	// based on their success rate. See
	// :ref:`Cluster outlier detection <arch_overview_outlier_detection>` statistics
	//
	// Note: this field may be omitted in any of the three following cases:
	//
	// 1. There were not enough hosts with enough request volume to proceed with success rate based
	//    outlier ejection.
	// 2. The threshold is computed to be < 0 because a negative value implies that there was no
	//    threshold for that interval.
	// 3. Outlier detection is not enabled for this cluster.
	SuccessRateEjectionThreshold *_type.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	// Mapping from host address to the host's current status.
	HostStatuses         []*HostStatus `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{1}
}
func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return m.Size()
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetAddedViaApi() bool {
	if m != nil {
		return m.AddedViaApi
	}
	return false
}

func (m *ClusterStatus) GetSuccessRateEjectionThreshold() *_type.Percent {
	if m != nil {
		return m.SuccessRateEjectionThreshold
	}
	return nil
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

// Current state of a particular host.
type HostStatus struct {
	// Address of this host.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// List of stats specific to this host.
	Stats []*SimpleMetric `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	// The host's current health status.
	HealthStatus *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	// Request success rate for this host over the last calculated interval.
	//
	// Note: the message will not be present if host did not have enough request volume to calculate
	// success rate or the cluster did not have enough hosts to run through success rate outlier
	// ejection.
	SuccessRate *_type.Percent `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	// The host's weight. If not configured, the value defaults to 1.
	Weight               uint32   `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{2}
}
func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return m.Size()
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *HostStatus) GetStats() []*SimpleMetric {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *HostStatus) GetHealthStatus() *HostHealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return nil
}

func (m *HostStatus) GetSuccessRate() *_type.Percent {
	if m != nil {
		return m.SuccessRate
	}
	return nil
}

func (m *HostStatus) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// Health status for a host.
type HostHealthStatus struct {
	// The host is currently failing active health checks.
	FailedActiveHealthCheck bool `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	// The host is currently considered an outlier and has been ejected.
	FailedOutlierCheck bool `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	// The host is currently being marked as degraded through active health checking.
	FailedActiveDegradedCheck bool `protobuf:"varint,4,opt,name=failed_active_degraded_check,json=failedActiveDegradedCheck,proto3" json:"failed_active_degraded_check,omitempty"`
	// Health status as reported by EDS. Note: only HEALTHY and UNHEALTHY are currently supported
	// here.
	// TODO(mrice32): pipe through remaining EDS health status possibilities.
	EdsHealthStatus      core.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"eds_health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HostHealthStatus) Reset()         { *m = HostHealthStatus{} }
func (m *HostHealthStatus) String() string { return proto.CompactTextString(m) }
func (*HostHealthStatus) ProtoMessage()    {}
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{3}
}
func (m *HostHealthStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostHealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostHealthStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostHealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostHealthStatus.Merge(m, src)
}
func (m *HostHealthStatus) XXX_Size() int {
	return m.Size()
}
func (m *HostHealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostHealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostHealthStatus proto.InternalMessageInfo

func (m *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if m != nil {
		return m.FailedActiveHealthCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedOutlierCheck() bool {
	if m != nil {
		return m.FailedOutlierCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedActiveDegradedCheck() bool {
	if m != nil {
		return m.FailedActiveDegradedCheck
	}
	return false
}

func (m *HostHealthStatus) GetEdsHealthStatus() core.HealthStatus {
	if m != nil {
		return m.EdsHealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*Clusters)(nil), "envoy.admin.v2alpha.Clusters")
	proto.RegisterType((*ClusterStatus)(nil), "envoy.admin.v2alpha.ClusterStatus")
	proto.RegisterType((*HostStatus)(nil), "envoy.admin.v2alpha.HostStatus")
	proto.RegisterType((*HostHealthStatus)(nil), "envoy.admin.v2alpha.HostHealthStatus")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/clusters.proto", fileDescriptor_c6251a3a957f478b) }

var fileDescriptor_c6251a3a957f478b = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xae, 0x1b, 0xc3, 0x5d, 0xd9, 0xf0, 0x10, 0x84, 0x6a, 0x5a, 0xdb, 0x08, 0xa4,
	0x9e, 0x12, 0x14, 0x10, 0x1c, 0x38, 0x40, 0xb7, 0x21, 0x4d, 0xa0, 0x89, 0x2a, 0x43, 0x48, 0x70,
	0x89, 0x8c, 0xfd, 0x58, 0x0c, 0x69, 0x1d, 0xd9, 0x6e, 0xa0, 0x9f, 0x82, 0x23, 0x5f, 0x89, 0x23,
	0x1f, 0x01, 0xf5, 0xc6, 0xb7, 0x40, 0xb1, 0xdd, 0x35, 0x65, 0xdd, 0xcd, 0xf6, 0xfb, 0xfd, 0xed,
	0xf7, 0x7f, 0xef, 0x19, 0x05, 0x30, 0x29, 0xc5, 0x2c, 0x22, 0x6c, 0xcc, 0x27, 0x51, 0x19, 0x93,
	0xbc, 0xc8, 0x48, 0x44, 0xf3, 0xa9, 0xd2, 0x20, 0x55, 0x58, 0x48, 0xa1, 0x05, 0xde, 0x37, 0x4c,
	0x68, 0x98, 0xd0, 0x31, 0x9d, 0xfe, 0x3a, 0xe1, 0x18, 0xb4, 0xe4, 0xd4, 0xe9, 0x3a, 0x5d, 0x87,
	0x14, 0x3c, 0x2a, 0xe3, 0x88, 0x0a, 0x09, 0x11, 0x61, 0x4c, 0x82, 0x5a, 0x00, 0x0f, 0xae, 0x02,
	0x19, 0x90, 0x5c, 0x67, 0x29, 0xcd, 0x80, 0x7e, 0x75, 0x94, 0x6f, 0x29, 0x3d, 0x2b, 0x20, 0x2a,
	0x40, 0x52, 0x98, 0x68, 0x1b, 0x09, 0x3e, 0xa0, 0xed, 0x63, 0x97, 0x2a, 0x3e, 0x43, 0x7b, 0x2e,
	0xed, 0x54, 0x69, 0xa2, 0xa7, 0x0a, 0x94, 0xef, 0xf5, 0x36, 0x06, 0xad, 0x38, 0x08, 0xd7, 0xe4,
	0x1f, 0x3a, 0xe1, 0xb9, 0x61, 0x93, 0x5d, 0x5a, 0xdf, 0x82, 0x0a, 0xfe, 0x7a, 0xa8, 0xbd, 0x82,
	0x60, 0x8c, 0x9a, 0x13, 0x32, 0x06, 0xdf, 0xeb, 0x79, 0x83, 0x9b, 0x89, 0x59, 0xe3, 0x00, 0xb5,
	0x09, 0x63, 0xc0, 0xd2, 0x92, 0x93, 0x94, 0x14, 0xdc, 0x6f, 0xf4, 0xbc, 0xc1, 0x76, 0xd2, 0x32,
	0x87, 0xef, 0x39, 0x19, 0x16, 0x1c, 0x7f, 0x44, 0x5d, 0x35, 0xa5, 0x14, 0x94, 0x4a, 0x25, 0xd1,
	0x90, 0xc2, 0x17, 0xa0, 0x9a, 0x8b, 0x49, 0xaa, 0x33, 0x09, 0x2a, 0x13, 0x39, 0xf3, 0x37, 0x7a,
	0xde, 0xa0, 0x15, 0xef, 0xbb, 0x3c, 0x2b, 0xa3, 0xe1, 0xc8, 0x1a, 0x4d, 0x0e, 0x9c, 0x36, 0x21,
	0x1a, 0x5e, 0x39, 0xe5, 0xbb, 0x85, 0x10, 0x9f, 0xa0, 0x76, 0x26, 0x94, 0x5e, 0x3a, 0x6e, 0x1a,
	0xc7, 0xdd, 0xb5, 0x8e, 0x4f, 0x85, 0xd2, 0xce, 0xee, 0x4e, 0x76, 0xb9, 0x06, 0x15, 0xfc, 0x6c,
	0x20, 0xb4, 0x0c, 0xe2, 0x27, 0xe8, 0x86, 0x6b, 0x93, 0xf1, 0xda, 0x8a, 0x3b, 0x8b, 0xeb, 0x0a,
	0x1e, 0x96, 0x71, 0x58, 0xf5, 0x29, 0x1c, 0x5a, 0x22, 0x59, 0xa0, 0xf8, 0x19, 0xda, 0xac, 0xb2,
	0x50, 0x7e, 0xc3, 0xa4, 0xd0, 0x5f, 0x9b, 0xc2, 0x39, 0x1f, 0x17, 0x39, 0x9c, 0x99, 0x29, 0x49,
	0x2c, 0x8f, 0x5f, 0xa3, 0xb6, 0x6b, 0xba, 0x75, 0xe1, 0xaa, 0xf1, 0xf0, 0x5a, 0x0f, 0xa7, 0x86,
	0xbe, 0x74, 0x52, 0xdb, 0xe1, 0xa7, 0x68, 0xa7, 0x5e, 0x6b, 0xbf, 0x79, 0x7d, 0x61, 0x5b, 0xb5,
	0xc2, 0xe2, 0xbb, 0x68, 0xeb, 0x1b, 0xf0, 0x8b, 0x4c, 0xfb, 0x9b, 0x3d, 0x6f, 0xd0, 0x4e, 0xdc,
	0x2e, 0xf8, 0xd1, 0x40, 0x7b, 0xff, 0x3f, 0x89, 0x9f, 0xa3, 0xce, 0x67, 0xc2, 0x73, 0x60, 0x29,
	0xa1, 0x9a, 0x97, 0x90, 0xd6, 0x67, 0xd6, 0x94, 0x6c, 0x3b, 0xb9, 0x67, 0x89, 0xa1, 0x01, 0xac,
	0xfa, 0xb8, 0x0a, 0xe3, 0x47, 0xe8, 0x8e, 0x13, 0x8b, 0xa9, 0xce, 0x39, 0x48, 0x27, 0xb3, 0x83,
	0x83, 0x6d, 0xec, 0xad, 0x0d, 0x59, 0xc5, 0x0b, 0x74, 0xb0, 0xfa, 0x1c, 0x83, 0x0b, 0x49, 0xaa,
	0xa1, 0xb3, 0xca, 0xa6, 0x51, 0xde, 0xaf, 0x3f, 0x78, 0xe2, 0x08, 0x7b, 0xc1, 0x1b, 0x74, 0x1b,
	0x98, 0x4a, 0xaf, 0x16, 0xf9, 0xd6, 0x72, 0x50, 0x6a, 0x9d, 0x5d, 0x29, 0xef, 0x2e, 0x30, 0x55,
	0x3f, 0x38, 0x7a, 0xf9, 0x6b, 0x7e, 0xe8, 0xfd, 0x9e, 0x1f, 0x7a, 0x7f, 0xe6, 0x87, 0x1e, 0xea,
	0x73, 0x61, 0x6f, 0x28, 0xa4, 0xf8, 0x3e, 0x5b, 0xd7, 0xb1, 0xa3, 0xc5, 0x2f, 0x52, 0xa3, 0xea,
	0xcb, 0x8e, 0xbc, 0x4f, 0x5b, 0xe6, 0xef, 0x3e, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x0e,
	0x28, 0xd9, 0x7a, 0x04, 0x00, 0x00,
}

func (m *Clusters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Clusters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClusterStatuses) > 0 {
		for _, msg := range m.ClusterStatuses {
			dAtA[i] = 0xa
			i++
			i = encodeVarintClusters(dAtA, i, uint64(msg.Size()))
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

func (m *ClusterStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClusters(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.AddedViaApi {
		dAtA[i] = 0x10
		i++
		if m.AddedViaApi {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.SuccessRateEjectionThreshold != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClusters(dAtA, i, uint64(m.SuccessRateEjectionThreshold.Size()))
		n1, err := m.SuccessRateEjectionThreshold.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.HostStatuses) > 0 {
		for _, msg := range m.HostStatuses {
			dAtA[i] = 0x22
			i++
			i = encodeVarintClusters(dAtA, i, uint64(msg.Size()))
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

func (m *HostStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Address != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClusters(dAtA, i, uint64(m.Address.Size()))
		n2, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Stats) > 0 {
		for _, msg := range m.Stats {
			dAtA[i] = 0x12
			i++
			i = encodeVarintClusters(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.HealthStatus != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClusters(dAtA, i, uint64(m.HealthStatus.Size()))
		n3, err := m.HealthStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.SuccessRate != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintClusters(dAtA, i, uint64(m.SuccessRate.Size()))
		n4, err := m.SuccessRate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Weight != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintClusters(dAtA, i, uint64(m.Weight))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HostHealthStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostHealthStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FailedActiveHealthCheck {
		dAtA[i] = 0x8
		i++
		if m.FailedActiveHealthCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.FailedOutlierCheck {
		dAtA[i] = 0x10
		i++
		if m.FailedOutlierCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.EdsHealthStatus != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintClusters(dAtA, i, uint64(m.EdsHealthStatus))
	}
	if m.FailedActiveDegradedCheck {
		dAtA[i] = 0x20
		i++
		if m.FailedActiveDegradedCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintClusters(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Clusters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClusterStatuses) > 0 {
		for _, e := range m.ClusterStatuses {
			l = e.Size()
			n += 1 + l + sovClusters(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovClusters(uint64(l))
	}
	if m.AddedViaApi {
		n += 2
	}
	if m.SuccessRateEjectionThreshold != nil {
		l = m.SuccessRateEjectionThreshold.Size()
		n += 1 + l + sovClusters(uint64(l))
	}
	if len(m.HostStatuses) > 0 {
		for _, e := range m.HostStatuses {
			l = e.Size()
			n += 1 + l + sovClusters(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HostStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovClusters(uint64(l))
	}
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovClusters(uint64(l))
		}
	}
	if m.HealthStatus != nil {
		l = m.HealthStatus.Size()
		n += 1 + l + sovClusters(uint64(l))
	}
	if m.SuccessRate != nil {
		l = m.SuccessRate.Size()
		n += 1 + l + sovClusters(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovClusters(uint64(m.Weight))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HostHealthStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FailedActiveHealthCheck {
		n += 2
	}
	if m.FailedOutlierCheck {
		n += 2
	}
	if m.EdsHealthStatus != 0 {
		n += 1 + sovClusters(uint64(m.EdsHealthStatus))
	}
	if m.FailedActiveDegradedCheck {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovClusters(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClusters(x uint64) (n int) {
	return sovClusters(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Clusters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusters
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
			return fmt.Errorf("proto: Clusters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Clusters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterStatuses = append(m.ClusterStatuses, &ClusterStatus{})
			if err := m.ClusterStatuses[len(m.ClusterStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusters(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusters
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClusters
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
func (m *ClusterStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusters
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
			return fmt.Errorf("proto: ClusterStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedViaApi", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddedViaApi = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRateEjectionThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRateEjectionThreshold == nil {
				m.SuccessRateEjectionThreshold = &_type.Percent{}
			}
			if err := m.SuccessRateEjectionThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostStatuses = append(m.HostStatuses, &HostStatus{})
			if err := m.HostStatuses[len(m.HostStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusters(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusters
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClusters
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
func (m *HostStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusters
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
			return fmt.Errorf("proto: HostStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Address == nil {
				m.Address = &core.Address{}
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stats = append(m.Stats, &SimpleMetric{})
			if err := m.Stats[len(m.Stats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HealthStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HealthStatus == nil {
				m.HealthStatus = &HostHealthStatus{}
			}
			if err := m.HealthStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
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
				return ErrInvalidLengthClusters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRate == nil {
				m.SuccessRate = &_type.Percent{}
			}
			if err := m.SuccessRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClusters(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusters
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClusters
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
func (m *HostHealthStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusters
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
			return fmt.Errorf("proto: HostHealthStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostHealthStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailedActiveHealthCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailedActiveHealthCheck = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailedOutlierCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailedOutlierCheck = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdsHealthStatus", wireType)
			}
			m.EdsHealthStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EdsHealthStatus |= core.HealthStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailedActiveDegradedCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusters
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailedActiveDegradedCheck = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipClusters(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusters
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClusters
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
func skipClusters(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClusters
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
					return 0, ErrIntOverflowClusters
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
					return 0, ErrIntOverflowClusters
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
				return 0, ErrInvalidLengthClusters
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthClusters
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClusters
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
				next, err := skipClusters(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthClusters
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
	ErrInvalidLengthClusters = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClusters   = fmt.Errorf("proto: integer overflow")
)
