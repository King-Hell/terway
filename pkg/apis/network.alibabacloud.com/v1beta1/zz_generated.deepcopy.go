//go:build !ignore_autogenerated

/*
Copyright 2021 Terway Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Allocation) DeepCopyInto(out *Allocation) {
	*out = *in
	out.AllocationType = in.AllocationType
	in.ENI.DeepCopyInto(&out.ENI)
	if in.ExtraRoutes != nil {
		in, out := &in.ExtraRoutes, &out.ExtraRoutes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Allocation.
func (in *Allocation) DeepCopy() *Allocation {
	if in == nil {
		return nil
	}
	out := new(Allocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationType) DeepCopyInto(out *AllocationType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationType.
func (in *AllocationType) DeepCopy() *AllocationType {
	if in == nil {
		return nil
	}
	out := new(AllocationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentOptions) DeepCopyInto(out *AttachmentOptions) {
	*out = *in
	if in.Trunk != nil {
		in, out := &in.Trunk, &out.Trunk
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentOptions.
func (in *AttachmentOptions) DeepCopy() *AttachmentOptions {
	if in == nil {
		return nil
	}
	out := new(AttachmentOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNIStatusInfo) DeepCopyInto(out *CNIStatusInfo) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNIStatusInfo.
func (in *CNIStatusInfo) DeepCopy() *CNIStatusInfo {
	if in == nil {
		return nil
	}
	out := new(CNIStatusInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.ObservedTime.DeepCopyInto(&out.ObservedTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENI) DeepCopyInto(out *ENI) {
	*out = *in
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AttachmentOptions.DeepCopyInto(&out.AttachmentOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENI.
func (in *ENI) DeepCopy() *ENI {
	if in == nil {
		return nil
	}
	out := new(ENI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENIInfo) DeepCopyInto(out *ENIInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENIInfo.
func (in *ENIInfo) DeepCopy() *ENIInfo {
	if in == nil {
		return nil
	}
	out := new(ENIInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENIOptions) DeepCopyInto(out *ENIOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENIOptions.
func (in *ENIOptions) DeepCopy() *ENIOptions {
	if in == nil {
		return nil
	}
	out := new(ENIOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENISpec) DeepCopyInto(out *ENISpec) {
	*out = *in
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagFilter != nil {
		in, out := &in.TagFilter, &out.TagFilter
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VSwitchOptions != nil {
		in, out := &in.VSwitchOptions, &out.VSwitchOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENISpec.
func (in *ENISpec) DeepCopy() *ENISpec {
	if in == nil {
		return nil
	}
	out := new(ENISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flavor) DeepCopyInto(out *Flavor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flavor.
func (in *Flavor) DeepCopy() *Flavor {
	if in == nil {
		return nil
	}
	out := new(Flavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IP) DeepCopyInto(out *IP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IP.
func (in *IP) DeepCopy() *IP {
	if in == nil {
		return nil
	}
	out := new(IP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in IPMap) DeepCopyInto(out *IPMap) {
	{
		in := &in
		*out = make(IPMap, len(*in))
		for key, val := range *in {
			var outVal *IP
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(IP)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPMap.
func (in IPMap) DeepCopy() IPMap {
	if in == nil {
		return nil
	}
	out := new(IPMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterface) DeepCopyInto(out *NetworkInterface) {
	*out = *in
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = make(map[string]*IP, len(*in))
		for key, val := range *in {
			var outVal *IP
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(IP)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = make(map[string]*IP, len(*in))
		for key, val := range *in {
			var outVal *IP
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(IP)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]Condition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterface.
func (in *NetworkInterface) DeepCopy() *NetworkInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Node) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeCap) DeepCopyInto(out *NodeCap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeCap.
func (in *NodeCap) DeepCopy() *NodeCap {
	if in == nil {
		return nil
	}
	out := new(NodeCap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeList) DeepCopyInto(out *NodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeList.
func (in *NodeList) DeepCopy() *NodeList {
	if in == nil {
		return nil
	}
	out := new(NodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMetadata) DeepCopyInto(out *NodeMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMetadata.
func (in *NodeMetadata) DeepCopy() *NodeMetadata {
	if in == nil {
		return nil
	}
	out := new(NodeMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRuntime) DeepCopyInto(out *NodeRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRuntime.
func (in *NodeRuntime) DeepCopy() *NodeRuntime {
	if in == nil {
		return nil
	}
	out := new(NodeRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRuntimeList) DeepCopyInto(out *NodeRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRuntimeList.
func (in *NodeRuntimeList) DeepCopy() *NodeRuntimeList {
	if in == nil {
		return nil
	}
	out := new(NodeRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRuntimeSpec) DeepCopyInto(out *NodeRuntimeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRuntimeSpec.
func (in *NodeRuntimeSpec) DeepCopy() *NodeRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRuntimeStatus) DeepCopyInto(out *NodeRuntimeStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(map[string]*RuntimePodStatus, len(*in))
		for key, val := range *in {
			var outVal *RuntimePodStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RuntimePodStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRuntimeStatus.
func (in *NodeRuntimeStatus) DeepCopy() *NodeRuntimeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeRuntimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	out.NodeMetadata = in.NodeMetadata
	out.NodeCap = in.NodeCap
	if in.ENISpec != nil {
		in, out := &in.ENISpec, &out.ENISpec
		*out = new(ENISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(PoolSpec)
		**out = **in
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = make([]Flavor, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	in.NextSyncOpenAPITime.DeepCopyInto(&out.NextSyncOpenAPITime)
	in.LastSyncOpenAPITime.DeepCopyInto(&out.LastSyncOpenAPITime)
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make(map[string]*NetworkInterface, len(*in))
		for key, val := range *in {
			var outVal *NetworkInterface
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(NetworkInterface)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodENI) DeepCopyInto(out *PodENI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodENI.
func (in *PodENI) DeepCopy() *PodENI {
	if in == nil {
		return nil
	}
	out := new(PodENI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodENI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodENIList) DeepCopyInto(out *PodENIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodENI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodENIList.
func (in *PodENIList) DeepCopy() *PodENIList {
	if in == nil {
		return nil
	}
	out := new(PodENIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodENIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodENISpec) DeepCopyInto(out *PodENISpec) {
	*out = *in
	if in.Allocations != nil {
		in, out := &in.Allocations, &out.Allocations
		*out = make([]Allocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodENISpec.
func (in *PodENISpec) DeepCopy() *PodENISpec {
	if in == nil {
		return nil
	}
	out := new(PodENISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodENIStatus) DeepCopyInto(out *PodENIStatus) {
	*out = *in
	in.PodLastSeen.DeepCopyInto(&out.PodLastSeen)
	if in.ENIInfos != nil {
		in, out := &in.ENIInfos, &out.ENIInfos
		*out = make(map[string]ENIInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodENIStatus.
func (in *PodENIStatus) DeepCopy() *PodENIStatus {
	if in == nil {
		return nil
	}
	out := new(PodENIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworking) DeepCopyInto(out *PodNetworking) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworking.
func (in *PodNetworking) DeepCopy() *PodNetworking {
	if in == nil {
		return nil
	}
	out := new(PodNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodNetworking) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkingList) DeepCopyInto(out *PodNetworkingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodNetworking, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkingList.
func (in *PodNetworkingList) DeepCopy() *PodNetworkingList {
	if in == nil {
		return nil
	}
	out := new(PodNetworkingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodNetworkingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkingSpec) DeepCopyInto(out *PodNetworkingSpec) {
	*out = *in
	out.ENIOptions = in.ENIOptions
	out.AllocationType = in.AllocationType
	in.Selector.DeepCopyInto(&out.Selector)
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VSwitchOptions != nil {
		in, out := &in.VSwitchOptions, &out.VSwitchOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.VSwitchSelectOptions = in.VSwitchSelectOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkingSpec.
func (in *PodNetworkingSpec) DeepCopy() *PodNetworkingSpec {
	if in == nil {
		return nil
	}
	out := new(PodNetworkingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkingStatus) DeepCopyInto(out *PodNetworkingStatus) {
	*out = *in
	if in.VSwitches != nil {
		in, out := &in.VSwitches, &out.VSwitches
		*out = make([]VSwitch, len(*in))
		copy(*out, *in)
	}
	in.UpdateAt.DeepCopyInto(&out.UpdateAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkingStatus.
func (in *PodNetworkingStatus) DeepCopy() *PodNetworkingStatus {
	if in == nil {
		return nil
	}
	out := new(PodNetworkingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpec) DeepCopyInto(out *PoolSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpec.
func (in *PoolSpec) DeepCopy() *PoolSpec {
	if in == nil {
		return nil
	}
	out := new(PoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimePodSpec) DeepCopyInto(out *RuntimePodSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimePodSpec.
func (in *RuntimePodSpec) DeepCopy() *RuntimePodSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimePodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimePodStatus) DeepCopyInto(out *RuntimePodStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make(map[CNIStatus]*CNIStatusInfo, len(*in))
		for key, val := range *in {
			var outVal *CNIStatusInfo
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(CNIStatusInfo)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimePodStatus.
func (in *RuntimePodStatus) DeepCopy() *RuntimePodStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimePodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSwitch) DeepCopyInto(out *VSwitch) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSwitch.
func (in *VSwitch) DeepCopy() *VSwitch {
	if in == nil {
		return nil
	}
	out := new(VSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSwitchSelectOptions) DeepCopyInto(out *VSwitchSelectOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSwitchSelectOptions.
func (in *VSwitchSelectOptions) DeepCopy() *VSwitchSelectOptions {
	if in == nil {
		return nil
	}
	out := new(VSwitchSelectOptions)
	in.DeepCopyInto(out)
	return out
}
