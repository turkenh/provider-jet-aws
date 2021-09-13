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

type AppsyncApiKeyObservation struct {
	Key string `json:"key" tf:"key"`
}

type AppsyncApiKeyParameters struct {

	// +kubebuilder:validation:Required
	APIID string `json:"apiId" tf:"api_id"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	Expires *string `json:"expires,omitempty" tf:"expires"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// AppsyncApiKeySpec defines the desired state of AppsyncApiKey
type AppsyncApiKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppsyncApiKeyParameters `json:"forProvider"`
}

// AppsyncApiKeyStatus defines the observed state of AppsyncApiKey.
type AppsyncApiKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppsyncApiKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncApiKey is the Schema for the AppsyncApiKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppsyncApiKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncApiKeySpec   `json:"spec"`
	Status            AppsyncApiKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncApiKeyList contains a list of AppsyncApiKeys
type AppsyncApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncApiKey `json:"items"`
}

// Repository type metadata.
var (
	AppsyncApiKeyKind             = "AppsyncApiKey"
	AppsyncApiKeyGroupKind        = schema.GroupKind{Group: Group, Kind: AppsyncApiKeyKind}.String()
	AppsyncApiKeyKindAPIVersion   = AppsyncApiKeyKind + "." + GroupVersion.String()
	AppsyncApiKeyGroupVersionKind = GroupVersion.WithKind(AppsyncApiKeyKind)
)

func init() {
	SchemeBuilder.Register(&AppsyncApiKey{}, &AppsyncApiKeyList{})
}
