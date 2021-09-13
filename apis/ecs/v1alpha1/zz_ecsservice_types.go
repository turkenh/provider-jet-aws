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

type CapacityProviderStrategyObservation struct {
}

type CapacityProviderStrategyParameters struct {

	// +kubebuilder:validation:Optional
	Base *int64 `json:"base,omitempty" tf:"base"`

	// +kubebuilder:validation:Required
	CapacityProvider string `json:"capacityProvider" tf:"capacity_provider"`

	// +kubebuilder:validation:Optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type DeploymentCircuitBreakerObservation struct {
}

type DeploymentCircuitBreakerParameters struct {

	// +kubebuilder:validation:Required
	Enable bool `json:"enable" tf:"enable"`

	// +kubebuilder:validation:Required
	Rollback bool `json:"rollback" tf:"rollback"`
}

type DeploymentControllerObservation struct {
}

type DeploymentControllerParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type EcsServiceObservation struct {
}

type EcsServiceParameters struct {

	// +kubebuilder:validation:Optional
	CapacityProviderStrategy []CapacityProviderStrategyParameters `json:"capacityProviderStrategy,omitempty" tf:"capacity_provider_strategy"`

	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster"`

	// +kubebuilder:validation:Optional
	DeploymentCircuitBreaker []DeploymentCircuitBreakerParameters `json:"deploymentCircuitBreaker,omitempty" tf:"deployment_circuit_breaker"`

	// +kubebuilder:validation:Optional
	DeploymentController []DeploymentControllerParameters `json:"deploymentController,omitempty" tf:"deployment_controller"`

	// +kubebuilder:validation:Optional
	DeploymentMaximumPercent *int64 `json:"deploymentMaximumPercent,omitempty" tf:"deployment_maximum_percent"`

	// +kubebuilder:validation:Optional
	DeploymentMinimumHealthyPercent *int64 `json:"deploymentMinimumHealthyPercent,omitempty" tf:"deployment_minimum_healthy_percent"`

	// +kubebuilder:validation:Optional
	DesiredCount *int64 `json:"desiredCount,omitempty" tf:"desired_count"`

	// +kubebuilder:validation:Optional
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags"`

	// +kubebuilder:validation:Optional
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command"`

	// +kubebuilder:validation:Optional
	ForceNewDeployment *bool `json:"forceNewDeployment,omitempty" tf:"force_new_deployment"`

	// +kubebuilder:validation:Optional
	HealthCheckGracePeriodSeconds *int64 `json:"healthCheckGracePeriodSeconds,omitempty" tf:"health_check_grace_period_seconds"`

	// +kubebuilder:validation:Optional
	IamRole *string `json:"iamRole,omitempty" tf:"iam_role"`

	// +kubebuilder:validation:Optional
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type"`

	// +kubebuilder:validation:Optional
	LoadBalancer []LoadBalancerParameters `json:"loadBalancer,omitempty" tf:"load_balancer"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration"`

	// +kubebuilder:validation:Optional
	OrderedPlacementStrategy []OrderedPlacementStrategyParameters `json:"orderedPlacementStrategy,omitempty" tf:"ordered_placement_strategy"`

	// +kubebuilder:validation:Optional
	PlacementConstraints []PlacementConstraintsParameters `json:"placementConstraints,omitempty" tf:"placement_constraints"`

	// +kubebuilder:validation:Optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version"`

	// +kubebuilder:validation:Optional
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SchedulingStrategy *string `json:"schedulingStrategy,omitempty" tf:"scheduling_strategy"`

	// +kubebuilder:validation:Optional
	ServiceRegistries []ServiceRegistriesParameters `json:"serviceRegistries,omitempty" tf:"service_registries"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	TaskDefinition *string `json:"taskDefinition,omitempty" tf:"task_definition"`

	// +kubebuilder:validation:Optional
	WaitForSteadyState *bool `json:"waitForSteadyState,omitempty" tf:"wait_for_steady_state"`
}

type LoadBalancerObservation struct {
}

type LoadBalancerParameters struct {

	// +kubebuilder:validation:Required
	ContainerName string `json:"containerName" tf:"container_name"`

	// +kubebuilder:validation:Required
	ContainerPort int64 `json:"containerPort" tf:"container_port"`

	// +kubebuilder:validation:Optional
	ElbName *string `json:"elbName,omitempty" tf:"elb_name"`

	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn"`
}

type NetworkConfigurationObservation struct {
}

type NetworkConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip"`

	// +kubebuilder:validation:Optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	// +kubebuilder:validation:Required
	Subnets []string `json:"subnets" tf:"subnets"`
}

type OrderedPlacementStrategyObservation struct {
}

type OrderedPlacementStrategyParameters struct {

	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type PlacementConstraintsObservation struct {
}

type PlacementConstraintsParameters struct {

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type ServiceRegistriesObservation struct {
}

type ServiceRegistriesParameters struct {

	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name"`

	// +kubebuilder:validation:Optional
	ContainerPort *int64 `json:"containerPort,omitempty" tf:"container_port"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port"`

	// +kubebuilder:validation:Required
	RegistryArn string `json:"registryArn" tf:"registry_arn"`
}

// EcsServiceSpec defines the desired state of EcsService
type EcsServiceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EcsServiceParameters `json:"forProvider"`
}

// EcsServiceStatus defines the observed state of EcsService.
type EcsServiceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EcsServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EcsService is the Schema for the EcsServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EcsService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcsServiceSpec   `json:"spec"`
	Status            EcsServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcsServiceList contains a list of EcsServices
type EcsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcsService `json:"items"`
}

// Repository type metadata.
var (
	EcsServiceKind             = "EcsService"
	EcsServiceGroupKind        = schema.GroupKind{Group: Group, Kind: EcsServiceKind}.String()
	EcsServiceKindAPIVersion   = EcsServiceKind + "." + GroupVersion.String()
	EcsServiceGroupVersionKind = GroupVersion.WithKind(EcsServiceKind)
)

func init() {
	SchemeBuilder.Register(&EcsService{}, &EcsServiceList{})
}
