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

type Ec2CapacityReservationObservation struct {
	Arn string `json:"arn" tf:"arn"`

	OwnerID string `json:"ownerId" tf:"owner_id"`
}

type Ec2CapacityReservationParameters struct {

	// +kubebuilder:validation:Required
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	// +kubebuilder:validation:Optional
	EbsOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized"`

	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date"`

	// +kubebuilder:validation:Optional
	EndDateType *string `json:"endDateType,omitempty" tf:"end_date_type"`

	// +kubebuilder:validation:Optional
	EphemeralStorage *bool `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage"`

	// +kubebuilder:validation:Required
	InstanceCount int64 `json:"instanceCount" tf:"instance_count"`

	// +kubebuilder:validation:Optional
	InstanceMatchCriteria *string `json:"instanceMatchCriteria,omitempty" tf:"instance_match_criteria"`

	// +kubebuilder:validation:Required
	InstancePlatform string `json:"instancePlatform" tf:"instance_platform"`

	// +kubebuilder:validation:Required
	InstanceType string `json:"instanceType" tf:"instance_type"`

	// +kubebuilder:validation:Optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	Tenancy *string `json:"tenancy,omitempty" tf:"tenancy"`
}

// Ec2CapacityReservationSpec defines the desired state of Ec2CapacityReservation
type Ec2CapacityReservationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2CapacityReservationParameters `json:"forProvider"`
}

// Ec2CapacityReservationStatus defines the observed state of Ec2CapacityReservation.
type Ec2CapacityReservationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2CapacityReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2CapacityReservation is the Schema for the Ec2CapacityReservations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2CapacityReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2CapacityReservationSpec   `json:"spec"`
	Status            Ec2CapacityReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2CapacityReservationList contains a list of Ec2CapacityReservations
type Ec2CapacityReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2CapacityReservation `json:"items"`
}

// Repository type metadata.
var (
	Ec2CapacityReservationKind             = "Ec2CapacityReservation"
	Ec2CapacityReservationGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2CapacityReservationKind}.String()
	Ec2CapacityReservationKindAPIVersion   = Ec2CapacityReservationKind + "." + GroupVersion.String()
	Ec2CapacityReservationGroupVersionKind = GroupVersion.WithKind(Ec2CapacityReservationKind)
)

func init() {
	SchemeBuilder.Register(&Ec2CapacityReservation{}, &Ec2CapacityReservationList{})
}
