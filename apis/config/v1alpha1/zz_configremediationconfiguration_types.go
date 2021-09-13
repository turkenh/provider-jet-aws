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

type ConfigRemediationConfigurationObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type ConfigRemediationConfigurationParameters struct {

	// +kubebuilder:validation:Required
	ConfigRuleName string `json:"configRuleName" tf:"config_rule_name"`

	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`

	// +kubebuilder:validation:Required
	TargetID string `json:"targetId" tf:"target_id"`

	// +kubebuilder:validation:Required
	TargetType string `json:"targetType" tf:"target_type"`

	// +kubebuilder:validation:Optional
	TargetVersion *string `json:"targetVersion,omitempty" tf:"target_version"`
}

type ParameterObservation struct {
}

type ParameterParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	ResourceValue *string `json:"resourceValue,omitempty" tf:"resource_value"`

	// +kubebuilder:validation:Optional
	StaticValue *string `json:"staticValue,omitempty" tf:"static_value"`
}

// ConfigRemediationConfigurationSpec defines the desired state of ConfigRemediationConfiguration
type ConfigRemediationConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConfigRemediationConfigurationParameters `json:"forProvider"`
}

// ConfigRemediationConfigurationStatus defines the observed state of ConfigRemediationConfiguration.
type ConfigRemediationConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConfigRemediationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigRemediationConfiguration is the Schema for the ConfigRemediationConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ConfigRemediationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigRemediationConfigurationSpec   `json:"spec"`
	Status            ConfigRemediationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigRemediationConfigurationList contains a list of ConfigRemediationConfigurations
type ConfigRemediationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigRemediationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	ConfigRemediationConfigurationKind             = "ConfigRemediationConfiguration"
	ConfigRemediationConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigRemediationConfigurationKind}.String()
	ConfigRemediationConfigurationKindAPIVersion   = ConfigRemediationConfigurationKind + "." + GroupVersion.String()
	ConfigRemediationConfigurationGroupVersionKind = GroupVersion.WithKind(ConfigRemediationConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&ConfigRemediationConfiguration{}, &ConfigRemediationConfigurationList{})
}
