//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 James Riley O'Donnell.

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Allowlist) DeepCopyInto(out *Allowlist) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Allowlist.
func (in *Allowlist) DeepCopy() *Allowlist {
	if in == nil {
		return nil
	}
	out := new(Allowlist)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Allowlist) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowlistList) DeepCopyInto(out *AllowlistList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Allowlist, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowlistList.
func (in *AllowlistList) DeepCopy() *AllowlistList {
	if in == nil {
		return nil
	}
	out := new(AllowlistList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AllowlistList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowlistSpec) DeepCopyInto(out *AllowlistSpec) {
	*out = *in
	if in.Logins != nil {
		in, out := &in.Logins, &out.Logins
		*out = make([]G8sTargets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SelfSignedTLSBundles != nil {
		in, out := &in.SelfSignedTLSBundles, &out.SelfSignedTLSBundles
		*out = make([]G8sTargets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSHKeyPairs != nil {
		in, out := &in.SSHKeyPairs, &out.SSHKeyPairs
		*out = make([]G8sTargets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowlistSpec.
func (in *AllowlistSpec) DeepCopy() *AllowlistSpec {
	if in == nil {
		return nil
	}
	out := new(AllowlistSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowlistStatus) DeepCopyInto(out *AllowlistStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowlistStatus.
func (in *AllowlistStatus) DeepCopy() *AllowlistStatus {
	if in == nil {
		return nil
	}
	out := new(AllowlistStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in G8s) DeepCopyInto(out *G8s) {
	{
		in := &in
		*out = make(G8s, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8s.
func (in G8s) DeepCopy() G8s {
	if in == nil {
		return nil
	}
	out := new(G8s)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sTargets) DeepCopyInto(out *G8sTargets) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]Target, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sTargets.
func (in *G8sTargets) DeepCopy() *G8sTargets {
	if in == nil {
		return nil
	}
	out := new(G8sTargets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Login) DeepCopyInto(out *Login) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Login.
func (in *Login) DeepCopy() *Login {
	if in == nil {
		return nil
	}
	out := new(Login)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Login) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginList) DeepCopyInto(out *LoginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Login, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginList.
func (in *LoginList) DeepCopy() *LoginList {
	if in == nil {
		return nil
	}
	out := new(LoginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginSpec) DeepCopyInto(out *LoginSpec) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(PasswordSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginSpec.
func (in *LoginSpec) DeepCopy() *LoginSpec {
	if in == nil {
		return nil
	}
	out := new(LoginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginStatus) DeepCopyInto(out *LoginStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginStatus.
func (in *LoginStatus) DeepCopy() *LoginStatus {
	if in == nil {
		return nil
	}
	out := new(LoginStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSpec) DeepCopyInto(out *PasswordSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSpec.
func (in *PasswordSpec) DeepCopy() *PasswordSpec {
	if in == nil {
		return nil
	}
	out := new(PasswordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyPair) DeepCopyInto(out *SSHKeyPair) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyPair.
func (in *SSHKeyPair) DeepCopy() *SSHKeyPair {
	if in == nil {
		return nil
	}
	out := new(SSHKeyPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHKeyPair) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyPairList) DeepCopyInto(out *SSHKeyPairList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSHKeyPair, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyPairList.
func (in *SSHKeyPairList) DeepCopy() *SSHKeyPairList {
	if in == nil {
		return nil
	}
	out := new(SSHKeyPairList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHKeyPairList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyPairSpec) DeepCopyInto(out *SSHKeyPairSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyPairSpec.
func (in *SSHKeyPairSpec) DeepCopy() *SSHKeyPairSpec {
	if in == nil {
		return nil
	}
	out := new(SSHKeyPairSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHKeyPairStatus) DeepCopyInto(out *SSHKeyPairStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHKeyPairStatus.
func (in *SSHKeyPairStatus) DeepCopy() *SSHKeyPairStatus {
	if in == nil {
		return nil
	}
	out := new(SSHKeyPairStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedTLSBundle) DeepCopyInto(out *SelfSignedTLSBundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedTLSBundle.
func (in *SelfSignedTLSBundle) DeepCopy() *SelfSignedTLSBundle {
	if in == nil {
		return nil
	}
	out := new(SelfSignedTLSBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSignedTLSBundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedTLSBundleList) DeepCopyInto(out *SelfSignedTLSBundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SelfSignedTLSBundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedTLSBundleList.
func (in *SelfSignedTLSBundleList) DeepCopy() *SelfSignedTLSBundleList {
	if in == nil {
		return nil
	}
	out := new(SelfSignedTLSBundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSignedTLSBundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedTLSBundleSpec) DeepCopyInto(out *SelfSignedTLSBundleSpec) {
	*out = *in
	if in.SANs != nil {
		in, out := &in.SANs, &out.SANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedTLSBundleSpec.
func (in *SelfSignedTLSBundleSpec) DeepCopy() *SelfSignedTLSBundleSpec {
	if in == nil {
		return nil
	}
	out := new(SelfSignedTLSBundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedTLSBundleStatus) DeepCopyInto(out *SelfSignedTLSBundleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedTLSBundleStatus.
func (in *SelfSignedTLSBundleStatus) DeepCopy() *SelfSignedTLSBundleStatus {
	if in == nil {
		return nil
	}
	out := new(SelfSignedTLSBundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}
