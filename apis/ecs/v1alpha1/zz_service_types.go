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

type CapacityProviderStrategyObservation struct {
}

type CapacityProviderStrategyParameters struct {

	// +kubebuilder:validation:Optional
	Base *int64 `json:"base,omitempty" tf:"base,omitempty"`

	// +kubebuilder:validation:Required
	CapacityProvider *string `json:"capacityProvider" tf:"capacity_provider,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type DeploymentCircuitBreakerObservation struct {
}

type DeploymentCircuitBreakerParameters struct {

	// +kubebuilder:validation:Required
	Enable *bool `json:"enable" tf:"enable,omitempty"`

	// +kubebuilder:validation:Required
	Rollback *bool `json:"rollback" tf:"rollback,omitempty"`
}

type DeploymentControllerObservation struct {
}

type DeploymentControllerParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LoadBalancerObservation struct {
}

type LoadBalancerParameters struct {

	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Required
	ContainerPort *int64 `json:"containerPort" tf:"container_port,omitempty"`

	// +kubebuilder:validation:Optional
	ElbName *string `json:"elbName,omitempty" tf:"elb_name,omitempty"`

	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`
}

type NetworkConfigurationObservation struct {
}

type NetworkConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupsRefs []v1.Reference `json:"securityGroupsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupsSelector *v1.Selector `json:"securityGroupsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetsRefs []v1.Reference `json:"subnetsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetsSelector *v1.Selector `json:"subnetsSelector,omitempty" tf:"-"`
}

type OrderedPlacementStrategyObservation struct {
}

type OrderedPlacementStrategyParameters struct {

	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PlacementConstraintsObservation struct {
}

type PlacementConstraintsParameters struct {

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ServiceObservation struct {
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ServiceParameters struct {

	// +kubebuilder:validation:Optional
	CapacityProviderStrategy []CapacityProviderStrategyParameters `json:"capacityProviderStrategy,omitempty" tf:"capacity_provider_strategy,omitempty"`

	// +crossplane:generate:reference:type=Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRef *v1.Reference `json:"clusterRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterSelector *v1.Selector `json:"clusterSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeploymentCircuitBreaker []DeploymentCircuitBreakerParameters `json:"deploymentCircuitBreaker,omitempty" tf:"deployment_circuit_breaker,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentController []DeploymentControllerParameters `json:"deploymentController,omitempty" tf:"deployment_controller,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentMaximumPercent *int64 `json:"deploymentMaximumPercent,omitempty" tf:"deployment_maximum_percent,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentMinimumHealthyPercent *int64 `json:"deploymentMinimumHealthyPercent,omitempty" tf:"deployment_minimum_healthy_percent,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredCount *int64 `json:"desiredCount,omitempty" tf:"desired_count,omitempty"`

	// +kubebuilder:validation:Optional
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags,omitempty"`

	// +kubebuilder:validation:Optional
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command,omitempty"`

	// +kubebuilder:validation:Optional
	ForceNewDeployment *bool `json:"forceNewDeployment,omitempty" tf:"force_new_deployment,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckGracePeriodSeconds *int64 `json:"healthCheckGracePeriodSeconds,omitempty" tf:"health_check_grace_period_seconds,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	IamRole *string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`

	// +kubebuilder:validation:Optional
	IamRoleRef *v1.Reference `json:"iamRoleRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IamRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancer []LoadBalancerParameters `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	OrderedPlacementStrategy []OrderedPlacementStrategyParameters `json:"orderedPlacementStrategy,omitempty" tf:"ordered_placement_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementConstraints []PlacementConstraintsParameters `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`

	// +kubebuilder:validation:Optional
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Optional
	SchedulingStrategy *string `json:"schedulingStrategy,omitempty" tf:"scheduling_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceRegistries []ServiceRegistriesParameters `json:"serviceRegistries,omitempty" tf:"service_registries,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TaskDefinition *string `json:"taskDefinition,omitempty" tf:"task_definition,omitempty"`

	// +kubebuilder:validation:Optional
	WaitForSteadyState *bool `json:"waitForSteadyState,omitempty" tf:"wait_for_steady_state,omitempty"`
}

type ServiceRegistriesObservation struct {
}

type ServiceRegistriesParameters struct {

	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerPort *int64 `json:"containerPort,omitempty" tf:"container_port,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	RegistryArn *string `json:"registryArn" tf:"registry_arn,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	ServiceKind             = "Service"
	ServiceGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceKind}.String()
	ServiceKindAPIVersion   = ServiceKind + "." + GroupVersion.String()
	ServiceGroupVersionKind = GroupVersion.WithKind(ServiceKind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
