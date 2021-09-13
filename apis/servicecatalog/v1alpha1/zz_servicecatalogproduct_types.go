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

type ProvisioningArtifactParametersObservation struct {
}

type ProvisioningArtifactParametersParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	TemplatePhysicalID *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id"`

	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServicecatalogProductObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedTime string `json:"createdTime" tf:"created_time"`

	HasDefaultPath bool `json:"hasDefaultPath" tf:"has_default_path"`

	Status string `json:"status" tf:"status"`
}

type ServicecatalogProductParameters struct {

	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	Distributor *string `json:"distributor,omitempty" tf:"distributor"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Owner string `json:"owner" tf:"owner"`

	// +kubebuilder:validation:Required
	ProvisioningArtifactParameters []ProvisioningArtifactParametersParameters `json:"provisioningArtifactParameters" tf:"provisioning_artifact_parameters"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SupportDescription *string `json:"supportDescription,omitempty" tf:"support_description"`

	// +kubebuilder:validation:Optional
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email"`

	// +kubebuilder:validation:Optional
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

// ServicecatalogProductSpec defines the desired state of ServicecatalogProduct
type ServicecatalogProductSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogProductParameters `json:"forProvider"`
}

// ServicecatalogProductStatus defines the observed state of ServicecatalogProduct.
type ServicecatalogProductStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogProductObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogProduct is the Schema for the ServicecatalogProducts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ServicecatalogProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogProductSpec   `json:"spec"`
	Status            ServicecatalogProductStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogProductList contains a list of ServicecatalogProducts
type ServicecatalogProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogProduct `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogProductKind             = "ServicecatalogProduct"
	ServicecatalogProductGroupKind        = schema.GroupKind{Group: Group, Kind: ServicecatalogProductKind}.String()
	ServicecatalogProductKindAPIVersion   = ServicecatalogProductKind + "." + GroupVersion.String()
	ServicecatalogProductGroupVersionKind = GroupVersion.WithKind(ServicecatalogProductKind)
)

func init() {
	SchemeBuilder.Register(&ServicecatalogProduct{}, &ServicecatalogProductList{})
}
