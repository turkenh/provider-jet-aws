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

type CodebuildSourceCredentialObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CodebuildSourceCredentialParameters struct {

	// +kubebuilder:validation:Required
	AuthType string `json:"authType" tf:"auth_type"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServerType string `json:"serverType" tf:"server_type"`

	// +kubebuilder:validation:Required
	Token string `json:"token" tf:"token"`

	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name"`
}

// CodebuildSourceCredentialSpec defines the desired state of CodebuildSourceCredential
type CodebuildSourceCredentialSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodebuildSourceCredentialParameters `json:"forProvider"`
}

// CodebuildSourceCredentialStatus defines the observed state of CodebuildSourceCredential.
type CodebuildSourceCredentialStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodebuildSourceCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildSourceCredential is the Schema for the CodebuildSourceCredentials API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CodebuildSourceCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildSourceCredentialSpec   `json:"spec"`
	Status            CodebuildSourceCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildSourceCredentialList contains a list of CodebuildSourceCredentials
type CodebuildSourceCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildSourceCredential `json:"items"`
}

// Repository type metadata.
var (
	CodebuildSourceCredentialKind             = "CodebuildSourceCredential"
	CodebuildSourceCredentialGroupKind        = schema.GroupKind{Group: Group, Kind: CodebuildSourceCredentialKind}.String()
	CodebuildSourceCredentialKindAPIVersion   = CodebuildSourceCredentialKind + "." + GroupVersion.String()
	CodebuildSourceCredentialGroupVersionKind = GroupVersion.WithKind(CodebuildSourceCredentialKind)
)

func init() {
	SchemeBuilder.Register(&CodebuildSourceCredential{}, &CodebuildSourceCredentialList{})
}
