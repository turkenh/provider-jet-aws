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

type DeadLetterConfigObservation struct {
}

type DeadLetterConfigParameters struct {

	// +kubebuilder:validation:Required
	TargetArn string `json:"targetArn" tf:"target_arn"`
}

type EnvironmentObservation struct {
}

type EnvironmentParameters struct {

	// +kubebuilder:validation:Optional
	Variables map[string]string `json:"variables,omitempty" tf:"variables"`
}

type FileSystemConfigObservation struct {
}

type FileSystemConfigParameters struct {

	// +kubebuilder:validation:Required
	Arn string `json:"arn" tf:"arn"`

	// +kubebuilder:validation:Required
	LocalMountPath string `json:"localMountPath" tf:"local_mount_path"`
}

type ImageConfigObservation struct {
}

type ImageConfigParameters struct {

	// +kubebuilder:validation:Optional
	Command []string `json:"command,omitempty" tf:"command"`

	// +kubebuilder:validation:Optional
	EntryPoint []string `json:"entryPoint,omitempty" tf:"entry_point"`

	// +kubebuilder:validation:Optional
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory"`
}

type LambdaFunctionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	InvokeArn string `json:"invokeArn" tf:"invoke_arn"`

	LastModified string `json:"lastModified" tf:"last_modified"`

	QualifiedArn string `json:"qualifiedArn" tf:"qualified_arn"`

	SigningJobArn string `json:"signingJobArn" tf:"signing_job_arn"`

	SigningProfileVersionArn string `json:"signingProfileVersionArn" tf:"signing_profile_version_arn"`

	SourceCodeSize int64 `json:"sourceCodeSize" tf:"source_code_size"`

	Version string `json:"version" tf:"version"`
}

type LambdaFunctionParameters struct {

	// +kubebuilder:validation:Optional
	CodeSigningConfigArn *string `json:"codeSigningConfigArn,omitempty" tf:"code_signing_config_arn"`

	// +kubebuilder:validation:Optional
	DeadLetterConfig []DeadLetterConfigParameters `json:"deadLetterConfig,omitempty" tf:"dead_letter_config"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	Environment []EnvironmentParameters `json:"environment,omitempty" tf:"environment"`

	// +kubebuilder:validation:Optional
	FileSystemConfig []FileSystemConfigParameters `json:"fileSystemConfig,omitempty" tf:"file_system_config"`

	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename"`

	// +kubebuilder:validation:Required
	FunctionName string `json:"functionName" tf:"function_name"`

	// +kubebuilder:validation:Optional
	Handler *string `json:"handler,omitempty" tf:"handler"`

	// +kubebuilder:validation:Optional
	ImageConfig []ImageConfigParameters `json:"imageConfig,omitempty" tf:"image_config"`

	// +kubebuilder:validation:Optional
	ImageURI *string `json:"imageUri,omitempty" tf:"image_uri"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`

	// +kubebuilder:validation:Optional
	Layers []string `json:"layers,omitempty" tf:"layers"`

	// +kubebuilder:validation:Optional
	MemorySize *int64 `json:"memorySize,omitempty" tf:"memory_size"`

	// +kubebuilder:validation:Optional
	PackageType *string `json:"packageType,omitempty" tf:"package_type"`

	// +kubebuilder:validation:Optional
	Publish *bool `json:"publish,omitempty" tf:"publish"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReservedConcurrentExecutions *int64 `json:"reservedConcurrentExecutions,omitempty" tf:"reserved_concurrent_executions"`

	// +kubebuilder:validation:Required
	Role string `json:"role" tf:"role"`

	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime"`

	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket"`

	// +kubebuilder:validation:Optional
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key"`

	// +kubebuilder:validation:Optional
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version"`

	// +kubebuilder:validation:Optional
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`

	// +kubebuilder:validation:Optional
	TracingConfig []TracingConfigParameters `json:"tracingConfig,omitempty" tf:"tracing_config"`

	// +kubebuilder:validation:Optional
	VpcConfig []VpcConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config"`
}

type TracingConfigObservation struct {
}

type TracingConfigParameters struct {

	// +kubebuilder:validation:Required
	Mode string `json:"mode" tf:"mode"`
}

type VpcConfigObservation struct {
	VpcID string `json:"vpcId" tf:"vpc_id"`
}

type VpcConfigParameters struct {

	// +kubebuilder:validation:Required
	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	// +kubebuilder:validation:Required
	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`
}

// LambdaFunctionSpec defines the desired state of LambdaFunction
type LambdaFunctionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LambdaFunctionParameters `json:"forProvider"`
}

// LambdaFunctionStatus defines the observed state of LambdaFunction.
type LambdaFunctionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LambdaFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaFunction is the Schema for the LambdaFunctions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LambdaFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaFunctionSpec   `json:"spec"`
	Status            LambdaFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaFunctionList contains a list of LambdaFunctions
type LambdaFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaFunction `json:"items"`
}

// Repository type metadata.
var (
	LambdaFunctionKind             = "LambdaFunction"
	LambdaFunctionGroupKind        = schema.GroupKind{Group: Group, Kind: LambdaFunctionKind}.String()
	LambdaFunctionKindAPIVersion   = LambdaFunctionKind + "." + GroupVersion.String()
	LambdaFunctionGroupVersionKind = GroupVersion.WithKind(LambdaFunctionKind)
)

func init() {
	SchemeBuilder.Register(&LambdaFunction{}, &LambdaFunctionList{})
}
