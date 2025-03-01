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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransitGatewayObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AssociationDefaultRouteTableID *string `json:"associationDefaultRouteTableId,omitempty" tf:"association_default_route_table_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	PropagationDefaultRouteTableID *string `json:"propagationDefaultRouteTableId,omitempty" tf:"propagation_default_route_table_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TransitGatewayParameters struct {

	// +kubebuilder:validation:Optional
	AmazonSideAsn *int64 `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// +kubebuilder:validation:Optional
	AutoAcceptSharedAttachments *string `json:"autoAcceptSharedAttachments,omitempty" tf:"auto_accept_shared_attachments,omitempty"`

	// +kubebuilder:validation:Optional
	DNSSupport *string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultRouteTableAssociation *string `json:"defaultRouteTableAssociation,omitempty" tf:"default_route_table_association,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultRouteTablePropagation *string `json:"defaultRouteTablePropagation,omitempty" tf:"default_route_table_propagation,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPNEcmpSupport *string `json:"vpnEcmpSupport,omitempty" tf:"vpn_ecmp_support,omitempty"`
}

// TransitGatewaySpec defines the desired state of TransitGateway
type TransitGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayParameters `json:"forProvider"`
}

// TransitGatewayStatus defines the observed state of TransitGateway.
type TransitGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGateway is the Schema for the TransitGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TransitGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewaySpec   `json:"spec"`
	Status            TransitGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayList contains a list of TransitGateways
type TransitGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGateway `json:"items"`
}

// Repository type metadata.
var (
	TransitGateway_Kind             = "TransitGateway"
	TransitGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGateway_Kind}.String()
	TransitGateway_KindAPIVersion   = TransitGateway_Kind + "." + CRDGroupVersion.String()
	TransitGateway_GroupVersionKind = CRDGroupVersion.WithKind(TransitGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGateway{}, &TransitGatewayList{})
}
