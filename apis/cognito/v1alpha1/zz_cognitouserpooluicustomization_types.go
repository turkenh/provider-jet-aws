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

type CognitoUserPoolUiCustomizationObservation struct {
	CSSVersion string `json:"cssVersion" tf:"css_version"`

	CreationDate string `json:"creationDate" tf:"creation_date"`

	ImageURL string `json:"imageUrl" tf:"image_url"`

	LastModifiedDate string `json:"lastModifiedDate" tf:"last_modified_date"`
}

type CognitoUserPoolUiCustomizationParameters struct {

	// +kubebuilder:validation:Optional
	CSS *string `json:"css,omitempty" tf:"css"`

	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id"`

	// +kubebuilder:validation:Optional
	ImageFile *string `json:"imageFile,omitempty" tf:"image_file"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	UserPoolID string `json:"userPoolId" tf:"user_pool_id"`
}

// CognitoUserPoolUiCustomizationSpec defines the desired state of CognitoUserPoolUiCustomization
type CognitoUserPoolUiCustomizationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoUserPoolUiCustomizationParameters `json:"forProvider"`
}

// CognitoUserPoolUiCustomizationStatus defines the observed state of CognitoUserPoolUiCustomization.
type CognitoUserPoolUiCustomizationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoUserPoolUiCustomizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolUiCustomization is the Schema for the CognitoUserPoolUiCustomizations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CognitoUserPoolUiCustomization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolUiCustomizationSpec   `json:"spec"`
	Status            CognitoUserPoolUiCustomizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolUiCustomizationList contains a list of CognitoUserPoolUiCustomizations
type CognitoUserPoolUiCustomizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPoolUiCustomization `json:"items"`
}

// Repository type metadata.
var (
	CognitoUserPoolUiCustomizationKind             = "CognitoUserPoolUiCustomization"
	CognitoUserPoolUiCustomizationGroupKind        = schema.GroupKind{Group: Group, Kind: CognitoUserPoolUiCustomizationKind}.String()
	CognitoUserPoolUiCustomizationKindAPIVersion   = CognitoUserPoolUiCustomizationKind + "." + GroupVersion.String()
	CognitoUserPoolUiCustomizationGroupVersionKind = GroupVersion.WithKind(CognitoUserPoolUiCustomizationKind)
)

func init() {
	SchemeBuilder.Register(&CognitoUserPoolUiCustomization{}, &CognitoUserPoolUiCustomizationList{})
}
