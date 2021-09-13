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

type AttributeObservation struct {
}

type AttributeParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type LbSslNegotiationPolicyObservation struct {
}

type LbSslNegotiationPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Attribute []AttributeParameters `json:"attribute,omitempty" tf:"attribute"`

	// +kubebuilder:validation:Required
	LbPort int64 `json:"lbPort" tf:"lb_port"`

	// +kubebuilder:validation:Required
	LoadBalancer string `json:"loadBalancer" tf:"load_balancer"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// LbSslNegotiationPolicySpec defines the desired state of LbSslNegotiationPolicy
type LbSslNegotiationPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LbSslNegotiationPolicyParameters `json:"forProvider"`
}

// LbSslNegotiationPolicyStatus defines the observed state of LbSslNegotiationPolicy.
type LbSslNegotiationPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LbSslNegotiationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbSslNegotiationPolicy is the Schema for the LbSslNegotiationPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LbSslNegotiationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSslNegotiationPolicySpec   `json:"spec"`
	Status            LbSslNegotiationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbSslNegotiationPolicyList contains a list of LbSslNegotiationPolicys
type LbSslNegotiationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbSslNegotiationPolicy `json:"items"`
}

// Repository type metadata.
var (
	LbSslNegotiationPolicyKind             = "LbSslNegotiationPolicy"
	LbSslNegotiationPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: LbSslNegotiationPolicyKind}.String()
	LbSslNegotiationPolicyKindAPIVersion   = LbSslNegotiationPolicyKind + "." + GroupVersion.String()
	LbSslNegotiationPolicyGroupVersionKind = GroupVersion.WithKind(LbSslNegotiationPolicyKind)
)

func init() {
	SchemeBuilder.Register(&LbSslNegotiationPolicy{}, &LbSslNegotiationPolicyList{})
}
