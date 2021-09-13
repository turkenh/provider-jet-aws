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

type NatGatewayObservation struct {
	NetworkInterfaceID string `json:"networkInterfaceId" tf:"network_interface_id"`

	PrivateIP string `json:"privateIp" tf:"private_ip"`

	PublicIP string `json:"publicIp" tf:"public_ip"`
}

type NatGatewayParameters struct {

	// +kubebuilder:validation:Optional
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id"`

	// +kubebuilder:validation:Optional
	ConnectivityType *string `json:"connectivityType,omitempty" tf:"connectivity_type"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// NatGatewaySpec defines the desired state of NatGateway
type NatGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NatGatewayParameters `json:"forProvider"`
}

// NatGatewayStatus defines the observed state of NatGateway.
type NatGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NatGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NatGateway is the Schema for the NatGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type NatGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatGatewaySpec   `json:"spec"`
	Status            NatGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatGatewayList contains a list of NatGateways
type NatGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatGateway `json:"items"`
}

// Repository type metadata.
var (
	NatGatewayKind             = "NatGateway"
	NatGatewayGroupKind        = schema.GroupKind{Group: Group, Kind: NatGatewayKind}.String()
	NatGatewayKindAPIVersion   = NatGatewayKind + "." + GroupVersion.String()
	NatGatewayGroupVersionKind = GroupVersion.WithKind(NatGatewayKind)
)

func init() {
	SchemeBuilder.Register(&NatGateway{}, &NatGatewayList{})
}
