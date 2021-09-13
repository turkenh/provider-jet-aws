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

type RouteTableAssociationObservation struct {
}

type RouteTableAssociationParameters struct {

	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RouteTableID string `json:"routeTableId" tf:"route_table_id"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`
}

// RouteTableAssociationSpec defines the desired state of RouteTableAssociation
type RouteTableAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RouteTableAssociationParameters `json:"forProvider"`
}

// RouteTableAssociationStatus defines the observed state of RouteTableAssociation.
type RouteTableAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAssociation is the Schema for the RouteTableAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type RouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableAssociationSpec   `json:"spec"`
	Status            RouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAssociationList contains a list of RouteTableAssociations
type RouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	RouteTableAssociationKind             = "RouteTableAssociation"
	RouteTableAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: RouteTableAssociationKind}.String()
	RouteTableAssociationKindAPIVersion   = RouteTableAssociationKind + "." + GroupVersion.String()
	RouteTableAssociationGroupVersionKind = GroupVersion.WithKind(RouteTableAssociationKind)
)

func init() {
	SchemeBuilder.Register(&RouteTableAssociation{}, &RouteTableAssociationList{})
}
