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

type WafregionalRuleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type WafregionalRuleParameters struct {

	// +kubebuilder:validation:Required
	MetricName string `json:"metricName" tf:"metric_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Predicate []WafregionalRulePredicateParameters `json:"predicate,omitempty" tf:"predicate"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type WafregionalRulePredicateObservation struct {
}

type WafregionalRulePredicateParameters struct {

	// +kubebuilder:validation:Required
	DataID string `json:"dataId" tf:"data_id"`

	// +kubebuilder:validation:Required
	Negated bool `json:"negated" tf:"negated"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

// WafregionalRuleSpec defines the desired state of WafregionalRule
type WafregionalRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalRuleParameters `json:"forProvider"`
}

// WafregionalRuleStatus defines the observed state of WafregionalRule.
type WafregionalRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRule is the Schema for the WafregionalRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafregionalRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalRuleSpec   `json:"spec"`
	Status            WafregionalRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRuleList contains a list of WafregionalRules
type WafregionalRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalRule `json:"items"`
}

// Repository type metadata.
var (
	WafregionalRuleKind             = "WafregionalRule"
	WafregionalRuleGroupKind        = schema.GroupKind{Group: Group, Kind: WafregionalRuleKind}.String()
	WafregionalRuleKindAPIVersion   = WafregionalRuleKind + "." + GroupVersion.String()
	WafregionalRuleGroupVersionKind = GroupVersion.WithKind(WafregionalRuleKind)
)

func init() {
	SchemeBuilder.Register(&WafregionalRule{}, &WafregionalRuleList{})
}
