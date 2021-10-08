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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ClusterParameters struct {

	// +crossplane:generate:reference:type=CapacityProvider
	// +kubebuilder:validation:Optional
	CapacityProviders []*string `json:"capacityProviders,omitempty" tf:"capacity_providers"`

	// +kubebuilder:validation:Optional
	CapacityProvidersRefs []v1.Reference `json:"capacityProvidersRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CapacityProvidersSelector *v1.Selector `json:"capacityProvidersSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration"`

	// +kubebuilder:validation:Optional
	DefaultCapacityProviderStrategy []DefaultCapacityProviderStrategyParameters `json:"defaultCapacityProviderStrategy,omitempty" tf:"default_capacity_provider_strategy"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Setting []SettingParameters `json:"setting,omitempty" tf:"setting"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags"`
}

type ConfigurationObservation struct {
}

type ConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ExecuteCommandConfiguration []ExecuteCommandConfigurationParameters `json:"executeCommandConfiguration,omitempty" tf:"execute_command_configuration"`
}

type DefaultCapacityProviderStrategyObservation struct {
}

type DefaultCapacityProviderStrategyParameters struct {

	// +kubebuilder:validation:Optional
	Base *int64 `json:"base,omitempty" tf:"base"`

	// +kubebuilder:validation:Required
	CapacityProvider *string `json:"capacityProvider" tf:"capacity_provider"`

	// +kubebuilder:validation:Optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type ExecuteCommandConfigurationObservation struct {
}

type ExecuteCommandConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	// +kubebuilder:validation:Optional
	LogConfiguration []LogConfigurationParameters `json:"logConfiguration,omitempty" tf:"log_configuration"`

	// +kubebuilder:validation:Optional
	Logging *string `json:"logging,omitempty" tf:"logging"`
}

type LogConfigurationObservation struct {
}

type LogConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CloudWatchEncryptionEnabled *bool `json:"cloudWatchEncryptionEnabled,omitempty" tf:"cloud_watch_encryption_enabled"`

	// +kubebuilder:validation:Optional
	CloudWatchLogGroupName *string `json:"cloudWatchLogGroupName,omitempty" tf:"cloud_watch_log_group_name"`

	// +kubebuilder:validation:Optional
	S3BucketEncryptionEnabled *bool `json:"s3BucketEncryptionEnabled,omitempty" tf:"s3_bucket_encryption_enabled"`

	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name"`

	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix"`
}

type SettingObservation struct {
}

type SettingParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	ClusterKind             = "Cluster"
	ClusterGroupKind        = schema.GroupKind{Group: Group, Kind: ClusterKind}.String()
	ClusterKindAPIVersion   = ClusterKind + "." + GroupVersion.String()
	ClusterGroupVersionKind = GroupVersion.WithKind(ClusterKind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
