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

type DataCatalogConfigObservation struct {
}

type DataCatalogConfigParameters struct {

	// +kubebuilder:validation:Optional
	Catalog *string `json:"catalog,omitempty" tf:"catalog"`

	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database"`

	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name"`
}

type FeatureDefinitionObservation struct {
}

type FeatureDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name"`

	// +kubebuilder:validation:Optional
	FeatureType *string `json:"featureType,omitempty" tf:"feature_type"`
}

type OfflineStoreConfigObservation struct {
}

type OfflineStoreConfigParameters struct {

	// +kubebuilder:validation:Optional
	DataCatalogConfig []DataCatalogConfigParameters `json:"dataCatalogConfig,omitempty" tf:"data_catalog_config"`

	// +kubebuilder:validation:Optional
	DisableGlueTableCreation *bool `json:"disableGlueTableCreation,omitempty" tf:"disable_glue_table_creation"`

	// +kubebuilder:validation:Required
	S3StorageConfig []S3StorageConfigParameters `json:"s3StorageConfig" tf:"s3_storage_config"`
}

type OnlineStoreConfigObservation struct {
}

type OnlineStoreConfigParameters struct {

	// +kubebuilder:validation:Optional
	EnableOnlineStore *bool `json:"enableOnlineStore,omitempty" tf:"enable_online_store"`

	// +kubebuilder:validation:Optional
	SecurityConfig []SecurityConfigParameters `json:"securityConfig,omitempty" tf:"security_config"`
}

type S3StorageConfigObservation struct {
}

type S3StorageConfigParameters struct {

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	// +kubebuilder:validation:Required
	S3URI string `json:"s3Uri" tf:"s3_uri"`
}

type SagemakerFeatureGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SagemakerFeatureGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	EventTimeFeatureName string `json:"eventTimeFeatureName" tf:"event_time_feature_name"`

	// +kubebuilder:validation:Required
	FeatureDefinition []FeatureDefinitionParameters `json:"featureDefinition" tf:"feature_definition"`

	// +kubebuilder:validation:Required
	FeatureGroupName string `json:"featureGroupName" tf:"feature_group_name"`

	// +kubebuilder:validation:Optional
	OfflineStoreConfig []OfflineStoreConfigParameters `json:"offlineStoreConfig,omitempty" tf:"offline_store_config"`

	// +kubebuilder:validation:Optional
	OnlineStoreConfig []OnlineStoreConfigParameters `json:"onlineStoreConfig,omitempty" tf:"online_store_config"`

	// +kubebuilder:validation:Required
	RecordIdentifierFeatureName string `json:"recordIdentifierFeatureName" tf:"record_identifier_feature_name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RoleArn string `json:"roleArn" tf:"role_arn"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type SecurityConfigObservation struct {
}

type SecurityConfigParameters struct {

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`
}

// SagemakerFeatureGroupSpec defines the desired state of SagemakerFeatureGroup
type SagemakerFeatureGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerFeatureGroupParameters `json:"forProvider"`
}

// SagemakerFeatureGroupStatus defines the observed state of SagemakerFeatureGroup.
type SagemakerFeatureGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerFeatureGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerFeatureGroup is the Schema for the SagemakerFeatureGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerFeatureGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerFeatureGroupSpec   `json:"spec"`
	Status            SagemakerFeatureGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerFeatureGroupList contains a list of SagemakerFeatureGroups
type SagemakerFeatureGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerFeatureGroup `json:"items"`
}

// Repository type metadata.
var (
	SagemakerFeatureGroupKind             = "SagemakerFeatureGroup"
	SagemakerFeatureGroupGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerFeatureGroupKind}.String()
	SagemakerFeatureGroupKindAPIVersion   = SagemakerFeatureGroupKind + "." + GroupVersion.String()
	SagemakerFeatureGroupGroupVersionKind = GroupVersion.WithKind(SagemakerFeatureGroupKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerFeatureGroup{}, &SagemakerFeatureGroupList{})
}
