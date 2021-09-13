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

type RunConfigObservation struct {
}

type RunConfigParameters struct {

	// +kubebuilder:validation:Optional
	ActiveTracing *bool `json:"activeTracing,omitempty" tf:"active_tracing"`

	// +kubebuilder:validation:Optional
	MemoryInMb *int64 `json:"memoryInMb,omitempty" tf:"memory_in_mb"`

	// +kubebuilder:validation:Optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Optional
	DurationInSeconds *int64 `json:"durationInSeconds,omitempty" tf:"duration_in_seconds"`

	// +kubebuilder:validation:Required
	Expression string `json:"expression" tf:"expression"`
}

type SyntheticsCanaryObservation struct {
	Arn string `json:"arn" tf:"arn"`

	EngineArn string `json:"engineArn" tf:"engine_arn"`

	SourceLocationArn string `json:"sourceLocationArn" tf:"source_location_arn"`

	Status string `json:"status" tf:"status"`

	Timeline []TimelineObservation `json:"timeline" tf:"timeline"`
}

type SyntheticsCanaryParameters struct {

	// +kubebuilder:validation:Required
	ArtifactS3Location string `json:"artifactS3Location" tf:"artifact_s3_location"`

	// +kubebuilder:validation:Required
	ExecutionRoleArn string `json:"executionRoleArn" tf:"execution_role_arn"`

	// +kubebuilder:validation:Optional
	FailureRetentionPeriod *int64 `json:"failureRetentionPeriod,omitempty" tf:"failure_retention_period"`

	// +kubebuilder:validation:Required
	Handler string `json:"handler" tf:"handler"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RunConfig []RunConfigParameters `json:"runConfig,omitempty" tf:"run_config"`

	// +kubebuilder:validation:Required
	RuntimeVersion string `json:"runtimeVersion" tf:"runtime_version"`

	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket"`

	// +kubebuilder:validation:Optional
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key"`

	// +kubebuilder:validation:Optional
	S3Version *string `json:"s3Version,omitempty" tf:"s3_version"`

	// +kubebuilder:validation:Required
	Schedule []ScheduleParameters `json:"schedule" tf:"schedule"`

	// +kubebuilder:validation:Optional
	StartCanary *bool `json:"startCanary,omitempty" tf:"start_canary"`

	// +kubebuilder:validation:Optional
	SuccessRetentionPeriod *int64 `json:"successRetentionPeriod,omitempty" tf:"success_retention_period"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	VpcConfig []VpcConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config"`

	// +kubebuilder:validation:Optional
	ZipFile *string `json:"zipFile,omitempty" tf:"zip_file"`
}

type TimelineObservation struct {
	Created string `json:"created" tf:"created"`

	LastModified string `json:"lastModified" tf:"last_modified"`

	LastStarted string `json:"lastStarted" tf:"last_started"`

	LastStopped string `json:"lastStopped" tf:"last_stopped"`
}

type TimelineParameters struct {
}

type VpcConfigObservation struct {
	VpcID string `json:"vpcId" tf:"vpc_id"`
}

type VpcConfigParameters struct {

	// +kubebuilder:validation:Optional
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	// +kubebuilder:validation:Optional
	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`
}

// SyntheticsCanarySpec defines the desired state of SyntheticsCanary
type SyntheticsCanarySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SyntheticsCanaryParameters `json:"forProvider"`
}

// SyntheticsCanaryStatus defines the observed state of SyntheticsCanary.
type SyntheticsCanaryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SyntheticsCanaryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SyntheticsCanary is the Schema for the SyntheticsCanarys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SyntheticsCanary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SyntheticsCanarySpec   `json:"spec"`
	Status            SyntheticsCanaryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SyntheticsCanaryList contains a list of SyntheticsCanarys
type SyntheticsCanaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SyntheticsCanary `json:"items"`
}

// Repository type metadata.
var (
	SyntheticsCanaryKind             = "SyntheticsCanary"
	SyntheticsCanaryGroupKind        = schema.GroupKind{Group: Group, Kind: SyntheticsCanaryKind}.String()
	SyntheticsCanaryKindAPIVersion   = SyntheticsCanaryKind + "." + GroupVersion.String()
	SyntheticsCanaryGroupVersionKind = GroupVersion.WithKind(SyntheticsCanaryKind)
)

func init() {
	SchemeBuilder.Register(&SyntheticsCanary{}, &SyntheticsCanaryList{})
}
