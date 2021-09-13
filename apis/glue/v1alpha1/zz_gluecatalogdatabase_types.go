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

type GlueCatalogDatabaseObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type GlueCatalogDatabaseParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	LocationURI *string `json:"locationUri,omitempty" tf:"location_uri"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetDatabase []TargetDatabaseParameters `json:"targetDatabase,omitempty" tf:"target_database"`
}

type TargetDatabaseObservation struct {
}

type TargetDatabaseParameters struct {

	// +kubebuilder:validation:Required
	CatalogID string `json:"catalogId" tf:"catalog_id"`

	// +kubebuilder:validation:Required
	DatabaseName string `json:"databaseName" tf:"database_name"`
}

// GlueCatalogDatabaseSpec defines the desired state of GlueCatalogDatabase
type GlueCatalogDatabaseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueCatalogDatabaseParameters `json:"forProvider"`
}

// GlueCatalogDatabaseStatus defines the observed state of GlueCatalogDatabase.
type GlueCatalogDatabaseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueCatalogDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueCatalogDatabase is the Schema for the GlueCatalogDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueCatalogDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueCatalogDatabaseSpec   `json:"spec"`
	Status            GlueCatalogDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueCatalogDatabaseList contains a list of GlueCatalogDatabases
type GlueCatalogDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueCatalogDatabase `json:"items"`
}

// Repository type metadata.
var (
	GlueCatalogDatabaseKind             = "GlueCatalogDatabase"
	GlueCatalogDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: GlueCatalogDatabaseKind}.String()
	GlueCatalogDatabaseKindAPIVersion   = GlueCatalogDatabaseKind + "." + GroupVersion.String()
	GlueCatalogDatabaseGroupVersionKind = GroupVersion.WithKind(GlueCatalogDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&GlueCatalogDatabase{}, &GlueCatalogDatabaseList{})
}
