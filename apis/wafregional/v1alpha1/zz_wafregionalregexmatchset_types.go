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

type RegexMatchTupleFieldToMatchObservation struct {
}

type RegexMatchTupleFieldToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type RegexMatchTupleObservation struct {
}

type RegexMatchTupleParameters struct {

	// +kubebuilder:validation:Required
	FieldToMatch []RegexMatchTupleFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match"`

	// +kubebuilder:validation:Required
	RegexPatternSetID string `json:"regexPatternSetId" tf:"regex_pattern_set_id"`

	// +kubebuilder:validation:Required
	TextTransformation string `json:"textTransformation" tf:"text_transformation"`
}

type WafregionalRegexMatchSetObservation struct {
}

type WafregionalRegexMatchSetParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	RegexMatchTuple []RegexMatchTupleParameters `json:"regexMatchTuple,omitempty" tf:"regex_match_tuple"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// WafregionalRegexMatchSetSpec defines the desired state of WafregionalRegexMatchSet
type WafregionalRegexMatchSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalRegexMatchSetParameters `json:"forProvider"`
}

// WafregionalRegexMatchSetStatus defines the observed state of WafregionalRegexMatchSet.
type WafregionalRegexMatchSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalRegexMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRegexMatchSet is the Schema for the WafregionalRegexMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafregionalRegexMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalRegexMatchSetSpec   `json:"spec"`
	Status            WafregionalRegexMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRegexMatchSetList contains a list of WafregionalRegexMatchSets
type WafregionalRegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalRegexMatchSet `json:"items"`
}

// Repository type metadata.
var (
	WafregionalRegexMatchSetKind             = "WafregionalRegexMatchSet"
	WafregionalRegexMatchSetGroupKind        = schema.GroupKind{Group: Group, Kind: WafregionalRegexMatchSetKind}.String()
	WafregionalRegexMatchSetKindAPIVersion   = WafregionalRegexMatchSetKind + "." + GroupVersion.String()
	WafregionalRegexMatchSetGroupVersionKind = GroupVersion.WithKind(WafregionalRegexMatchSetKind)
)

func init() {
	SchemeBuilder.Register(&WafregionalRegexMatchSet{}, &WafregionalRegexMatchSetList{})
}
