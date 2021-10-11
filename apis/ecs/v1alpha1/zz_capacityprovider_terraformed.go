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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/pkg/errors"

	common "github.com/crossplane-contrib/provider-tf-aws/config/common"
	"github.com/crossplane-contrib/terrajet/pkg/resource"
	"github.com/crossplane-contrib/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this CapacityProvider
func (mg *CapacityProvider) GetTerraformResourceType() string {
	return "aws_ecs_capacity_provider"
}

// GetTerraformResourceIDField returns Terraform identifier field for this CapacityProvider
func (tr *CapacityProvider) GetTerraformResourceIDField() string {
	return "id"
}

// GetConnectionDetailsMapping for this CapacityProvider
func (tr *CapacityProvider) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CapacityProvider
func (tr *CapacityProvider) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CapacityProvider
func (tr *CapacityProvider) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this CapacityProvider
func (tr *CapacityProvider) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	common.ExternalNameAsName(base, meta.GetExternalName(tr))
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CapacityProvider
func (tr *CapacityProvider) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CapacityProvider using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CapacityProvider) LateInitialize(attrs []byte) (bool, error) {
	params := &CapacityProviderParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}
