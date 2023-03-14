//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOverride) DeepCopyInto(out *ClusterOverride) {
	*out = *in
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]OverridePatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOverride.
func (in *ClusterOverride) DeepCopy() *ClusterOverride {
	if in == nil {
		return nil
	}
	out := new(ClusterOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerOverride) DeepCopyInto(out *ControllerOverride) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ClusterOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerOverride.
func (in *ControllerOverride) DeepCopy() *ControllerOverride {
	if in == nil {
		return nil
	}
	out := new(ControllerOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedDeployment) DeepCopyInto(out *FederatedDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedDeployment.
func (in *FederatedDeployment) DeepCopy() *FederatedDeployment {
	if in == nil {
		return nil
	}
	out := new(FederatedDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedDeploymentList) DeepCopyInto(out *FederatedDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedDeploymentList.
func (in *FederatedDeploymentList) DeepCopy() *FederatedDeploymentList {
	if in == nil {
		return nil
	}
	out := new(FederatedDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedDeploymentSpec) DeepCopyInto(out *FederatedDeploymentSpec) {
	*out = *in
	in.GenericSpecWithPlacements.DeepCopyInto(&out.GenericSpecWithPlacements)
	in.GenericSpecWithOverrides.DeepCopyInto(&out.GenericSpecWithOverrides)
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedDeploymentSpec.
func (in *FederatedDeploymentSpec) DeepCopy() *FederatedDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericClusterReference) DeepCopyInto(out *GenericClusterReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterReference.
func (in *GenericClusterReference) DeepCopy() *GenericClusterReference {
	if in == nil {
		return nil
	}
	out := new(GenericClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericClusterStatus) DeepCopyInto(out *GenericClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterStatus.
func (in *GenericClusterStatus) DeepCopy() *GenericClusterStatus {
	if in == nil {
		return nil
	}
	out := new(GenericClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericCondition) DeepCopyInto(out *GenericCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericCondition.
func (in *GenericCondition) DeepCopy() *GenericCondition {
	if in == nil {
		return nil
	}
	out := new(GenericCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericFederatedFollower) DeepCopyInto(out *GenericFederatedFollower) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericFederatedFollower.
func (in *GenericFederatedFollower) DeepCopy() *GenericFederatedFollower {
	if in == nil {
		return nil
	}
	out := new(GenericFederatedFollower)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericFederatedStatus) DeepCopyInto(out *GenericFederatedStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*GenericCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(GenericCondition)
				**out = **in
			}
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]GenericClusterStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericFederatedStatus.
func (in *GenericFederatedStatus) DeepCopy() *GenericFederatedStatus {
	if in == nil {
		return nil
	}
	out := new(GenericFederatedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericFollowerSpec) DeepCopyInto(out *GenericFollowerSpec) {
	*out = *in
	in.GenericSpecWithPlacements.DeepCopyInto(&out.GenericSpecWithPlacements)
	in.GenericSpecWithFollows.DeepCopyInto(&out.GenericSpecWithFollows)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericFollowerSpec.
func (in *GenericFollowerSpec) DeepCopy() *GenericFollowerSpec {
	if in == nil {
		return nil
	}
	out := new(GenericFollowerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericObjectWithOverrides) DeepCopyInto(out *GenericObjectWithOverrides) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(GenericSpecWithOverrides)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericObjectWithOverrides.
func (in *GenericObjectWithOverrides) DeepCopy() *GenericObjectWithOverrides {
	if in == nil {
		return nil
	}
	out := new(GenericObjectWithOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericObjectWithPlacements) DeepCopyInto(out *GenericObjectWithPlacements) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericObjectWithPlacements.
func (in *GenericObjectWithPlacements) DeepCopy() *GenericObjectWithPlacements {
	if in == nil {
		return nil
	}
	out := new(GenericObjectWithPlacements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericObjectWithStatus) DeepCopyInto(out *GenericObjectWithStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(GenericFederatedStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericObjectWithStatus.
func (in *GenericObjectWithStatus) DeepCopy() *GenericObjectWithStatus {
	if in == nil {
		return nil
	}
	out := new(GenericObjectWithStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericSpecWithFollows) DeepCopyInto(out *GenericSpecWithFollows) {
	*out = *in
	if in.Follows != nil {
		in, out := &in.Follows, &out.Follows
		*out = make([]LeaderReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericSpecWithFollows.
func (in *GenericSpecWithFollows) DeepCopy() *GenericSpecWithFollows {
	if in == nil {
		return nil
	}
	out := new(GenericSpecWithFollows)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericSpecWithOverrides) DeepCopyInto(out *GenericSpecWithOverrides) {
	*out = *in
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]ControllerOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericSpecWithOverrides.
func (in *GenericSpecWithOverrides) DeepCopy() *GenericSpecWithOverrides {
	if in == nil {
		return nil
	}
	out := new(GenericSpecWithOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericSpecWithPlacements) DeepCopyInto(out *GenericSpecWithPlacements) {
	*out = *in
	if in.Placements != nil {
		in, out := &in.Placements, &out.Placements
		*out = make([]PlacementWithController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericSpecWithPlacements.
func (in *GenericSpecWithPlacements) DeepCopy() *GenericSpecWithPlacements {
	if in == nil {
		return nil
	}
	out := new(GenericSpecWithPlacements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderReference) DeepCopyInto(out *LeaderReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderReference.
func (in *LeaderReference) DeepCopy() *LeaderReference {
	if in == nil {
		return nil
	}
	out := new(LeaderReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in OverridePatches) DeepCopyInto(out *OverridePatches) {
	{
		in := &in
		*out = make(OverridePatches, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverridePatches.
func (in OverridePatches) DeepCopy() OverridePatches {
	if in == nil {
		return nil
	}
	out := new(OverridePatches)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]GenericClusterReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementWithController) DeepCopyInto(out *PlacementWithController) {
	*out = *in
	in.Placement.DeepCopyInto(&out.Placement)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementWithController.
func (in *PlacementWithController) DeepCopy() *PlacementWithController {
	if in == nil {
		return nil
	}
	out := new(PlacementWithController)
	in.DeepCopyInto(out)
	return out
}
