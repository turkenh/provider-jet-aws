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

type SecretsmanagerSecretRotationObservation struct {
	RotationEnabled bool `json:"rotationEnabled" tf:"rotation_enabled"`
}

type SecretsmanagerSecretRotationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RotationLambdaArn string `json:"rotationLambdaArn" tf:"rotation_lambda_arn"`

	// +kubebuilder:validation:Required
	RotationRules []SecretsmanagerSecretRotationRotationRulesParameters `json:"rotationRules" tf:"rotation_rules"`

	// +kubebuilder:validation:Required
	SecretID string `json:"secretId" tf:"secret_id"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type SecretsmanagerSecretRotationRotationRulesObservation struct {
}

type SecretsmanagerSecretRotationRotationRulesParameters struct {

	// +kubebuilder:validation:Required
	AutomaticallyAfterDays int64 `json:"automaticallyAfterDays" tf:"automatically_after_days"`
}

// SecretsmanagerSecretRotationSpec defines the desired state of SecretsmanagerSecretRotation
type SecretsmanagerSecretRotationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecretsmanagerSecretRotationParameters `json:"forProvider"`
}

// SecretsmanagerSecretRotationStatus defines the observed state of SecretsmanagerSecretRotation.
type SecretsmanagerSecretRotationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecretsmanagerSecretRotationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecretRotation is the Schema for the SecretsmanagerSecretRotations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SecretsmanagerSecretRotation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretsmanagerSecretRotationSpec   `json:"spec"`
	Status            SecretsmanagerSecretRotationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecretRotationList contains a list of SecretsmanagerSecretRotations
type SecretsmanagerSecretRotationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretsmanagerSecretRotation `json:"items"`
}

// Repository type metadata.
var (
	SecretsmanagerSecretRotationKind             = "SecretsmanagerSecretRotation"
	SecretsmanagerSecretRotationGroupKind        = schema.GroupKind{Group: Group, Kind: SecretsmanagerSecretRotationKind}.String()
	SecretsmanagerSecretRotationKindAPIVersion   = SecretsmanagerSecretRotationKind + "." + GroupVersion.String()
	SecretsmanagerSecretRotationGroupVersionKind = GroupVersion.WithKind(SecretsmanagerSecretRotationKind)
)

func init() {
	SchemeBuilder.Register(&SecretsmanagerSecretRotation{}, &SecretsmanagerSecretRotationList{})
}
