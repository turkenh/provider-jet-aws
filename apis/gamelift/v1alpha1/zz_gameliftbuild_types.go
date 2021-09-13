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

type GameliftBuildObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type GameliftBuildParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	OperatingSystem string `json:"operatingSystem" tf:"operating_system"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	StorageLocation []StorageLocationParameters `json:"storageLocation" tf:"storage_location"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type StorageLocationObservation struct {
}

type StorageLocationParameters struct {

	// +kubebuilder:validation:Required
	Bucket string `json:"bucket" tf:"bucket"`

	// +kubebuilder:validation:Required
	Key string `json:"key" tf:"key"`

	// +kubebuilder:validation:Required
	RoleArn string `json:"roleArn" tf:"role_arn"`
}

// GameliftBuildSpec defines the desired state of GameliftBuild
type GameliftBuildSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GameliftBuildParameters `json:"forProvider"`
}

// GameliftBuildStatus defines the observed state of GameliftBuild.
type GameliftBuildStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GameliftBuildObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftBuild is the Schema for the GameliftBuilds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GameliftBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameliftBuildSpec   `json:"spec"`
	Status            GameliftBuildStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftBuildList contains a list of GameliftBuilds
type GameliftBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameliftBuild `json:"items"`
}

// Repository type metadata.
var (
	GameliftBuildKind             = "GameliftBuild"
	GameliftBuildGroupKind        = schema.GroupKind{Group: Group, Kind: GameliftBuildKind}.String()
	GameliftBuildKindAPIVersion   = GameliftBuildKind + "." + GroupVersion.String()
	GameliftBuildGroupVersionKind = GroupVersion.WithKind(GameliftBuildKind)
)

func init() {
	SchemeBuilder.Register(&GameliftBuild{}, &GameliftBuildList{})
}
