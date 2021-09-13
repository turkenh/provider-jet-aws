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

type MediaConvertQueueObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type MediaConvertQueueParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PricingPlan *string `json:"pricingPlan,omitempty" tf:"pricing_plan"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReservationPlanSettings []ReservationPlanSettingsParameters `json:"reservationPlanSettings,omitempty" tf:"reservation_plan_settings"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ReservationPlanSettingsObservation struct {
}

type ReservationPlanSettingsParameters struct {

	// +kubebuilder:validation:Required
	Commitment string `json:"commitment" tf:"commitment"`

	// +kubebuilder:validation:Required
	RenewalType string `json:"renewalType" tf:"renewal_type"`

	// +kubebuilder:validation:Required
	ReservedSlots int64 `json:"reservedSlots" tf:"reserved_slots"`
}

// MediaConvertQueueSpec defines the desired state of MediaConvertQueue
type MediaConvertQueueSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MediaConvertQueueParameters `json:"forProvider"`
}

// MediaConvertQueueStatus defines the observed state of MediaConvertQueue.
type MediaConvertQueueStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MediaConvertQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MediaConvertQueue is the Schema for the MediaConvertQueues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type MediaConvertQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaConvertQueueSpec   `json:"spec"`
	Status            MediaConvertQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MediaConvertQueueList contains a list of MediaConvertQueues
type MediaConvertQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MediaConvertQueue `json:"items"`
}

// Repository type metadata.
var (
	MediaConvertQueueKind             = "MediaConvertQueue"
	MediaConvertQueueGroupKind        = schema.GroupKind{Group: Group, Kind: MediaConvertQueueKind}.String()
	MediaConvertQueueKindAPIVersion   = MediaConvertQueueKind + "." + GroupVersion.String()
	MediaConvertQueueGroupVersionKind = GroupVersion.WithKind(MediaConvertQueueKind)
)

func init() {
	SchemeBuilder.Register(&MediaConvertQueue{}, &MediaConvertQueueList{})
}
