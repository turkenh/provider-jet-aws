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

type HomeDirectoryMappingsObservation struct {
}

type HomeDirectoryMappingsParameters struct {

	// +kubebuilder:validation:Required
	Entry string `json:"entry" tf:"entry"`

	// +kubebuilder:validation:Required
	Target string `json:"target" tf:"target"`
}

type PosixProfileObservation struct {
}

type PosixProfileParameters struct {

	// +kubebuilder:validation:Required
	GID int64 `json:"gid" tf:"gid"`

	// +kubebuilder:validation:Optional
	SecondaryGids []int64 `json:"secondaryGids,omitempty" tf:"secondary_gids"`

	// +kubebuilder:validation:Required
	UID int64 `json:"uid" tf:"uid"`
}

type TransferUserObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type TransferUserParameters struct {

	// +kubebuilder:validation:Optional
	HomeDirectory *string `json:"homeDirectory,omitempty" tf:"home_directory"`

	// +kubebuilder:validation:Optional
	HomeDirectoryMappings []HomeDirectoryMappingsParameters `json:"homeDirectoryMappings,omitempty" tf:"home_directory_mappings"`

	// +kubebuilder:validation:Optional
	HomeDirectoryType *string `json:"homeDirectoryType,omitempty" tf:"home_directory_type"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy"`

	// +kubebuilder:validation:Optional
	PosixProfile []PosixProfileParameters `json:"posixProfile,omitempty" tf:"posix_profile"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Role string `json:"role" tf:"role"`

	// +kubebuilder:validation:Required
	ServerID string `json:"serverId" tf:"server_id"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Required
	UserName string `json:"userName" tf:"user_name"`
}

// TransferUserSpec defines the desired state of TransferUser
type TransferUserSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TransferUserParameters `json:"forProvider"`
}

// TransferUserStatus defines the observed state of TransferUser.
type TransferUserStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TransferUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransferUser is the Schema for the TransferUsers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type TransferUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferUserSpec   `json:"spec"`
	Status            TransferUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferUserList contains a list of TransferUsers
type TransferUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferUser `json:"items"`
}

// Repository type metadata.
var (
	TransferUserKind             = "TransferUser"
	TransferUserGroupKind        = schema.GroupKind{Group: Group, Kind: TransferUserKind}.String()
	TransferUserKindAPIVersion   = TransferUserKind + "." + GroupVersion.String()
	TransferUserGroupVersionKind = GroupVersion.WithKind(TransferUserKind)
)

func init() {
	SchemeBuilder.Register(&TransferUser{}, &TransferUserList{})
}
