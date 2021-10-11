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

type FargateProfileObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FargateProfileParameters struct {

	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PodExecutionRoleArn *string `json:"podExecutionRoleArn,omitempty" tf:"pod_execution_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	PodExecutionRoleArnRef *v1.Reference `json:"podExecutionRoleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PodExecutionRoleArnSelector *v1.Selector `json:"podExecutionRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Required
	Selector []SelectorParameters `json:"selector" tf:"selector,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SelectorObservation struct {
}

type SelectorParameters struct {

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

// FargateProfileSpec defines the desired state of FargateProfile
type FargateProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FargateProfileParameters `json:"forProvider"`
}

// FargateProfileStatus defines the observed state of FargateProfile.
type FargateProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FargateProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FargateProfile is the Schema for the FargateProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type FargateProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FargateProfileSpec   `json:"spec"`
	Status            FargateProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FargateProfileList contains a list of FargateProfiles
type FargateProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FargateProfile `json:"items"`
}

// Repository type metadata.
var (
	FargateProfileKind             = "FargateProfile"
	FargateProfileGroupKind        = schema.GroupKind{Group: Group, Kind: FargateProfileKind}.String()
	FargateProfileKindAPIVersion   = FargateProfileKind + "." + GroupVersion.String()
	FargateProfileGroupVersionKind = GroupVersion.WithKind(FargateProfileKind)
)

func init() {
	SchemeBuilder.Register(&FargateProfile{}, &FargateProfileList{})
}
