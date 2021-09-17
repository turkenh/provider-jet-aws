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
	"github.com/crossplane-contrib/terrajet/pkg/json"
)

// GetTerraformResourceType returns Terraform resource type for this Ec2ClientVpnAuthorizationRule
func (mg *Ec2ClientVpnAuthorizationRule) GetTerraformResourceType() string {
	return "aws_ec2_client_vpn_authorization_rule"
}

// GetTerraformResourceIdField returns Terraform identifier field for this Ec2ClientVpnAuthorizationRule
func (tr *Ec2ClientVpnAuthorizationRule) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this Ec2ClientVpnAuthorizationRule
func (tr *Ec2ClientVpnAuthorizationRule) GetObservation() ([]byte, error) {
	return json.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this Ec2ClientVpnAuthorizationRule
func (tr *Ec2ClientVpnAuthorizationRule) SetObservation(data []byte) error {
	return json.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this Ec2ClientVpnAuthorizationRule
func (tr *Ec2ClientVpnAuthorizationRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.JSParser.Unmarshal(p, &base)
}

// SetParameters for this Ec2ClientVpnAuthorizationRule
func (tr *Ec2ClientVpnAuthorizationRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}
