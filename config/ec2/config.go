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

package ec2

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

func init() {
	config.Store.SetForResource("aws_instance", config.Resource{
		ExternalName: config.ExternalName{
			// Set to true explicitly since the value is calculated by AWS.
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"subnet_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
			"vpc_security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"security_groups": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"root_block_device[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"network_interface[*].network_interface_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"ebs_block_device[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
		},
	})

	config.Store.SetForResource("aws_eip", config.Resource{
		Kind: "ElasticIP",
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"instance": {
				Type: "Instance",
			},
		},

		UseAsync: true,
	})

	config.Store.SetForResource("aws_ec2_transit_gateway", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
	})

	config.Store.SetForResource("aws_ec2_transit_gateway_route", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_attachment_id": {
				Type: "TransitGatewayVpcAttachment",
			},
			"transit_gateway_route_table_id": {
				Type: "TransitGatewayRouteTable",
			},
		},
	})

	config.Store.SetForResource("aws_ec2_transit_gateway_route_table", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_id": {
				Type: "TransitGateway",
			},
		},
	})
	config.Store.SetForResource("aws_ec2_transit_gateway_route_table_association", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_attachment_id": {
				Type: "TransitGatewayVpcAttachment",
			},
			"transit_gateway_route_table_id": {
				Type: "TransitGatewayRouteTable",
			},
		},
	})
	config.Store.SetForResource("aws_ec2_transit_gateway_vpc_attachment", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"subnet_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
			"transit_gateway_id": {
				Type: "TransitGateway",
			},
			"vpc_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.VPC",
			},
		},
	})
	config.Store.SetForResource("aws_ec2_transit_gateway_vpc_attachment_accepter", config.Resource{
		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},
		References: map[string]config.Reference{
			"transit_gateway_attachment_id": {
				Type: "TransitGatewayVpcAttachment",
			},
		},
	})

	config.Store.SetForResource("aws_launch_template", config.Resource{
		// Note(turkenh): Kind as "LaunchTemplate" fails with:
		// panic: cannot generate crd: cannot build types for LaunchTemplate:
		//  cannot build the types: cannot generate parameters type name of
		//  LaunchTemplate: could not generate a unique name for
		//  LaunchTemplateParameters
		Kind: "EC2LaunchTemplate",

		ExternalName: config.ExternalName{
			DisableNameInitializer: true,
		},

		References: map[string]config.Reference{
			"security_group_names": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"vpc_security_group_ids": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"block_device_mappings[*].ebs[*].kms_key_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
			},
			"iam_instance_profile[*].arn": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.InstanceProfile",
				Extractor: common.PathARNExtractor,
			},
			"iam_instance_profile[*].name": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.InstanceProfile",
			},
			"network_interfaces[*].network_interface_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.NetworkInterface",
			},
			"network_interfaces[*].security_groups": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup",
			},
			"network_interfaces[*].subnet_id": {
				Type: "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet",
			},
		},
	})

}
