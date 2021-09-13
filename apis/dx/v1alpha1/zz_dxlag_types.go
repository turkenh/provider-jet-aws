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

type DxLagObservation struct {
	Arn string `json:"arn" tf:"arn"`

	HasLogicalRedundancy string `json:"hasLogicalRedundancy" tf:"has_logical_redundancy"`

	JumboFrameCapable bool `json:"jumboFrameCapable" tf:"jumbo_frame_capable"`
}

type DxLagParameters struct {

	// +kubebuilder:validation:Required
	ConnectionsBandwidth string `json:"connectionsBandwidth" tf:"connections_bandwidth"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

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
}

// DxLagSpec defines the desired state of DxLag
type DxLagSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxLagParameters `json:"forProvider"`
}

// DxLagStatus defines the observed state of DxLag.
type DxLagStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxLagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxLag is the Schema for the DxLags API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DxLag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxLagSpec   `json:"spec"`
	Status            DxLagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxLagList contains a list of DxLags
type DxLagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxLag `json:"items"`
}

// Repository type metadata.
var (
	DxLagKind             = "DxLag"
	DxLagGroupKind        = schema.GroupKind{Group: Group, Kind: DxLagKind}.String()
	DxLagKindAPIVersion   = DxLagKind + "." + GroupVersion.String()
	DxLagGroupVersionKind = GroupVersion.WithKind(DxLagKind)
)

func init() {
	SchemeBuilder.Register(&DxLag{}, &DxLagList{})
}
