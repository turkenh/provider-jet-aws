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

// GetTerraformResourceType returns Terraform resource type for this S3BucketNotification
func (mg *S3BucketNotification) GetTerraformResourceType() string {
	return "aws_s3_bucket_notification"
}

// GetTerraformResourceIdField returns Terraform identifier field for this S3BucketNotification
func (tr *S3BucketNotification) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this S3BucketNotification
func (tr *S3BucketNotification) GetObservation() ([]byte, error) {
	return json.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this S3BucketNotification
func (tr *S3BucketNotification) SetObservation(data []byte) error {
	return json.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this S3BucketNotification
func (tr *S3BucketNotification) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.JSParser.Unmarshal(p, &base)
}

// SetParameters for this S3BucketNotification
func (tr *S3BucketNotification) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}
