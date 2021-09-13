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

type OrganizationsDelegatedAdministratorObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DelegationEnabledDate string `json:"delegationEnabledDate" tf:"delegation_enabled_date"`

	Email string `json:"email" tf:"email"`

	JoinedMethod string `json:"joinedMethod" tf:"joined_method"`

	JoinedTimestamp string `json:"joinedTimestamp" tf:"joined_timestamp"`

	Name string `json:"name" tf:"name"`

	Status string `json:"status" tf:"status"`
}

type OrganizationsDelegatedAdministratorParameters struct {

	// +kubebuilder:validation:Required
	AccountID string `json:"accountId" tf:"account_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServicePrincipal string `json:"servicePrincipal" tf:"service_principal"`
}

// OrganizationsDelegatedAdministratorSpec defines the desired state of OrganizationsDelegatedAdministrator
type OrganizationsDelegatedAdministratorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OrganizationsDelegatedAdministratorParameters `json:"forProvider"`
}

// OrganizationsDelegatedAdministratorStatus defines the observed state of OrganizationsDelegatedAdministrator.
type OrganizationsDelegatedAdministratorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OrganizationsDelegatedAdministratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsDelegatedAdministrator is the Schema for the OrganizationsDelegatedAdministrators API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type OrganizationsDelegatedAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsDelegatedAdministratorSpec   `json:"spec"`
	Status            OrganizationsDelegatedAdministratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsDelegatedAdministratorList contains a list of OrganizationsDelegatedAdministrators
type OrganizationsDelegatedAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsDelegatedAdministrator `json:"items"`
}

// Repository type metadata.
var (
	OrganizationsDelegatedAdministratorKind             = "OrganizationsDelegatedAdministrator"
	OrganizationsDelegatedAdministratorGroupKind        = schema.GroupKind{Group: Group, Kind: OrganizationsDelegatedAdministratorKind}.String()
	OrganizationsDelegatedAdministratorKindAPIVersion   = OrganizationsDelegatedAdministratorKind + "." + GroupVersion.String()
	OrganizationsDelegatedAdministratorGroupVersionKind = GroupVersion.WithKind(OrganizationsDelegatedAdministratorKind)
)

func init() {
	SchemeBuilder.Register(&OrganizationsDelegatedAdministrator{}, &OrganizationsDelegatedAdministratorList{})
}
