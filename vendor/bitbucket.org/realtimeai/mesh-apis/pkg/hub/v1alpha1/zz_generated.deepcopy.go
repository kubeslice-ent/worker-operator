//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProperty) DeepCopyInto(out *ClusterProperty) {
	*out = *in
	out.Telemetry = in.Telemetry
	out.GeoLocation = in.GeoLocation
	out.Monitoring = in.Monitoring
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProperty.
func (in *ClusterProperty) DeepCopy() *ClusterProperty {
	if in == nil {
		return nil
	}
	out := new(ClusterProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	out.ClusterProperty = in.ClusterProperty
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalGatewayConfig) DeepCopyInto(out *ExternalGatewayConfig) {
	*out = *in
	out.Ingress = in.Ingress
	out.Egress = in.Egress
	out.NsIngress = in.NsIngress
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalGatewayConfig.
func (in *ExternalGatewayConfig) DeepCopy() *ExternalGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalGatewayConfigOptions) DeepCopyInto(out *ExternalGatewayConfigOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalGatewayConfigOptions.
func (in *ExternalGatewayConfigOptions) DeepCopy() *ExternalGatewayConfigOptions {
	if in == nil {
		return nil
	}
	out := new(ExternalGatewayConfigOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoLocation) DeepCopyInto(out *GeoLocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoLocation.
func (in *GeoLocation) DeepCopy() *GeoLocation {
	if in == nil {
		return nil
	}
	out := new(GeoLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesDashboard) DeepCopyInto(out *KubernetesDashboard) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesDashboard.
func (in *KubernetesDashboard) DeepCopy() *KubernetesDashboard {
	if in == nil {
		return nil
	}
	out := new(KubernetesDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	out.KubernetesDashboard = in.KubernetesDashboard
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceIsolationProfile) DeepCopyInto(out *NamespaceIsolationProfile) {
	*out = *in
	if in.ApplicationNamespaces != nil {
		in, out := &in.ApplicationNamespaces, &out.ApplicationNamespaces
		*out = make([]SliceNamespaceSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedNamespaces != nil {
		in, out := &in.AllowedNamespaces, &out.AllowedNamespaces
		*out = make([]SliceNamespaceSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceIsolationProfile.
func (in *NamespaceIsolationProfile) DeepCopy() *NamespaceIsolationProfile {
	if in == nil {
		return nil
	}
	out := new(NamespaceIsolationProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QOSProfile) DeepCopyInto(out *QOSProfile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QOSProfile.
func (in *QOSProfile) DeepCopy() *QOSProfile {
	if in == nil {
		return nil
	}
	out := new(QOSProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReadWrite != nil {
		in, out := &in.ReadWrite, &out.ReadWrite
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryEndpoint) DeepCopyInto(out *ServiceDiscoveryEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryEndpoint.
func (in *ServiceDiscoveryEndpoint) DeepCopy() *ServiceDiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPort) DeepCopyInto(out *ServiceDiscoveryPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPort.
func (in *ServiceDiscoveryPort) DeepCopy() *ServiceDiscoveryPort {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfig) DeepCopyInto(out *ServiceExportConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfig.
func (in *ServiceExportConfig) DeepCopy() *ServiceExportConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExportConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfigList) DeepCopyInto(out *ServiceExportConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceExportConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfigList.
func (in *ServiceExportConfigList) DeepCopy() *ServiceExportConfigList {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExportConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfigSpec) DeepCopyInto(out *ServiceExportConfigSpec) {
	*out = *in
	if in.ServiceDiscoveryEndpoints != nil {
		in, out := &in.ServiceDiscoveryEndpoints, &out.ServiceDiscoveryEndpoints
		*out = make([]ServiceDiscoveryEndpoint, len(*in))
		copy(*out, *in)
	}
	if in.ServiceDiscoveryPorts != nil {
		in, out := &in.ServiceDiscoveryPorts, &out.ServiceDiscoveryPorts
		*out = make([]ServiceDiscoveryPort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfigSpec.
func (in *ServiceExportConfigSpec) DeepCopy() *ServiceExportConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfigStatus) DeepCopyInto(out *ServiceExportConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfigStatus.
func (in *ServiceExportConfigStatus) DeepCopy() *ServiceExportConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfig) DeepCopyInto(out *SliceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfig.
func (in *SliceConfig) DeepCopy() *SliceConfig {
	if in == nil {
		return nil
	}
	out := new(SliceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SliceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfigList) DeepCopyInto(out *SliceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SliceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfigList.
func (in *SliceConfigList) DeepCopy() *SliceConfigList {
	if in == nil {
		return nil
	}
	out := new(SliceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SliceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfigSpec) DeepCopyInto(out *SliceConfigSpec) {
	*out = *in
	out.SliceGatewayProvider = in.SliceGatewayProvider
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.QosProfileDetails = in.QosProfileDetails
	in.NamespaceIsolationProfile.DeepCopyInto(&out.NamespaceIsolationProfile)
	if in.ExternalGatewayConfig != nil {
		in, out := &in.ExternalGatewayConfig, &out.ExternalGatewayConfig
		*out = make([]ExternalGatewayConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfigSpec.
func (in *SliceConfigSpec) DeepCopy() *SliceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SliceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfigStatus) DeepCopyInto(out *SliceConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfigStatus.
func (in *SliceConfigStatus) DeepCopy() *SliceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SliceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceNamespaceSelection) DeepCopyInto(out *SliceNamespaceSelection) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceNamespaceSelection.
func (in *SliceNamespaceSelection) DeepCopy() *SliceNamespaceSelection {
	if in == nil {
		return nil
	}
	out := new(SliceNamespaceSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpokeSliceGatewayProvider) DeepCopyInto(out *SpokeSliceGatewayProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpokeSliceGatewayProvider.
func (in *SpokeSliceGatewayProvider) DeepCopy() *SpokeSliceGatewayProvider {
	if in == nil {
		return nil
	}
	out := new(SpokeSliceGatewayProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Telemetry) DeepCopyInto(out *Telemetry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Telemetry.
func (in *Telemetry) DeepCopy() *Telemetry {
	if in == nil {
		return nil
	}
	out := new(Telemetry)
	in.DeepCopyInto(out)
	return out
}