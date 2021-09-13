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

type WafregionalRegexPatternSetObservation struct {
}

type WafregionalRegexPatternSetParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	RegexPatternStrings []string `json:"regexPatternStrings,omitempty" tf:"regex_pattern_strings"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// WafregionalRegexPatternSetSpec defines the desired state of WafregionalRegexPatternSet
type WafregionalRegexPatternSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalRegexPatternSetParameters `json:"forProvider"`
}

// WafregionalRegexPatternSetStatus defines the observed state of WafregionalRegexPatternSet.
type WafregionalRegexPatternSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalRegexPatternSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRegexPatternSet is the Schema for the WafregionalRegexPatternSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafregionalRegexPatternSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalRegexPatternSetSpec   `json:"spec"`
	Status            WafregionalRegexPatternSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRegexPatternSetList contains a list of WafregionalRegexPatternSets
type WafregionalRegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalRegexPatternSet `json:"items"`
}

// Repository type metadata.
var (
	WafregionalRegexPatternSetKind             = "WafregionalRegexPatternSet"
	WafregionalRegexPatternSetGroupKind        = schema.GroupKind{Group: Group, Kind: WafregionalRegexPatternSetKind}.String()
	WafregionalRegexPatternSetKindAPIVersion   = WafregionalRegexPatternSetKind + "." + GroupVersion.String()
	WafregionalRegexPatternSetGroupVersionKind = GroupVersion.WithKind(WafregionalRegexPatternSetKind)
)

func init() {
	SchemeBuilder.Register(&WafregionalRegexPatternSet{}, &WafregionalRegexPatternSetList{})
}
