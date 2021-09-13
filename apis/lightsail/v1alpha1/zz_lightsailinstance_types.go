/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type LightsailInstanceObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CPUCount int64 `json:"cpuCount" tf:"cpu_count"`

	CreatedAt string `json:"createdAt" tf:"created_at"`

	IPv6Address string `json:"ipv6Address" tf:"ipv6_address"`

	IPv6Addresses []string `json:"ipv6Addresses" tf:"ipv6_addresses"`

	IsStaticIP bool `json:"isStaticIp" tf:"is_static_ip"`

	PrivateIPAddress string `json:"privateIpAddress" tf:"private_ip_address"`

	PublicIPAddress string `json:"publicIpAddress" tf:"public_ip_address"`

	RAMSize float64 `json:"ramSize" tf:"ram_size"`

	Username string `json:"username" tf:"username"`
}

type LightsailInstanceParameters struct {

	// +kubebuilder:validation:Required
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	// +kubebuilder:validation:Required
	BlueprintID string `json:"blueprintId" tf:"blueprint_id"`

	// +kubebuilder:validation:Required
	BundleID string `json:"bundleId" tf:"bundle_id"`

	// +kubebuilder:validation:Optional
	KeyPairName *string `json:"keyPairName,omitempty" tf:"key_pair_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
}

// LightsailInstanceSpec defines the desired state of LightsailInstance
type LightsailInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LightsailInstanceParameters `json:"forProvider"`
}

// LightsailInstanceStatus defines the observed state of LightsailInstance.
type LightsailInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LightsailInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailInstance is the Schema for the LightsailInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LightsailInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailInstanceSpec   `json:"spec"`
	Status            LightsailInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailInstanceList contains a list of LightsailInstances
type LightsailInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LightsailInstance `json:"items"`
}

// Repository type metadata.
var (
	LightsailInstanceKind             = "LightsailInstance"
	LightsailInstanceGroupKind        = schema.GroupKind{Group: Group, Kind: LightsailInstanceKind}.String()
	LightsailInstanceKindAPIVersion   = LightsailInstanceKind + "." + GroupVersion.String()
	LightsailInstanceGroupVersionKind = GroupVersion.WithKind(LightsailInstanceKind)
)

func init() {
	SchemeBuilder.Register(&LightsailInstance{}, &LightsailInstanceList{})
}
