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

type SecurityGroupRuleObservation struct {
}

type SecurityGroupRuleParameters struct {

	// +kubebuilder:validation:Optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	FromPort int64 `json:"fromPort" tf:"from_port"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// +kubebuilder:validation:Optional
	PrefixListIds []string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SecurityGroupID string `json:"securityGroupId" tf:"security_group_id"`

	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupID *string `json:"sourceSecurityGroupId,omitempty" tf:"source_security_group_id"`

	// +kubebuilder:validation:Required
	ToPort int64 `json:"toPort" tf:"to_port"`

	// Type of rule, ingress (inbound) or egress (outbound).
	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

// SecurityGroupRuleSpec defines the desired state of SecurityGroupRule
type SecurityGroupRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecurityGroupRuleParameters `json:"forProvider"`
}

// SecurityGroupRuleStatus defines the observed state of SecurityGroupRule.
type SecurityGroupRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecurityGroupRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRule is the Schema for the SecurityGroupRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupRuleSpec   `json:"spec"`
	Status            SecurityGroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRuleList contains a list of SecurityGroupRules
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupRuleKind             = "SecurityGroupRule"
	SecurityGroupRuleGroupKind        = schema.GroupKind{Group: Group, Kind: SecurityGroupRuleKind}.String()
	SecurityGroupRuleKindAPIVersion   = SecurityGroupRuleKind + "." + GroupVersion.String()
	SecurityGroupRuleGroupVersionKind = GroupVersion.WithKind(SecurityGroupRuleKind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupRule{}, &SecurityGroupRuleList{})
}
