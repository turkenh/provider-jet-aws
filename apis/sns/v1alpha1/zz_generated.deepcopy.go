// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplication) DeepCopyInto(out *SnsPlatformApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplication.
func (in *SnsPlatformApplication) DeepCopy() *SnsPlatformApplication {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsPlatformApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationList) DeepCopyInto(out *SnsPlatformApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsPlatformApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationList.
func (in *SnsPlatformApplicationList) DeepCopy() *SnsPlatformApplicationList {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsPlatformApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationObservation) DeepCopyInto(out *SnsPlatformApplicationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationObservation.
func (in *SnsPlatformApplicationObservation) DeepCopy() *SnsPlatformApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationParameters) DeepCopyInto(out *SnsPlatformApplicationParameters) {
	*out = *in
	if in.EventDeliveryFailureTopicArn != nil {
		in, out := &in.EventDeliveryFailureTopicArn, &out.EventDeliveryFailureTopicArn
		*out = new(string)
		**out = **in
	}
	if in.EventEndpointCreatedTopicArn != nil {
		in, out := &in.EventEndpointCreatedTopicArn, &out.EventEndpointCreatedTopicArn
		*out = new(string)
		**out = **in
	}
	if in.EventEndpointDeletedTopicArn != nil {
		in, out := &in.EventEndpointDeletedTopicArn, &out.EventEndpointDeletedTopicArn
		*out = new(string)
		**out = **in
	}
	if in.EventEndpointUpdatedTopicArn != nil {
		in, out := &in.EventEndpointUpdatedTopicArn, &out.EventEndpointUpdatedTopicArn
		*out = new(string)
		**out = **in
	}
	if in.FailureFeedbackRoleArn != nil {
		in, out := &in.FailureFeedbackRoleArn, &out.FailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.PlatformPrincipal != nil {
		in, out := &in.PlatformPrincipal, &out.PlatformPrincipal
		*out = new(string)
		**out = **in
	}
	if in.SuccessFeedbackRoleArn != nil {
		in, out := &in.SuccessFeedbackRoleArn, &out.SuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.SuccessFeedbackSampleRate != nil {
		in, out := &in.SuccessFeedbackSampleRate, &out.SuccessFeedbackSampleRate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationParameters.
func (in *SnsPlatformApplicationParameters) DeepCopy() *SnsPlatformApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationSpec) DeepCopyInto(out *SnsPlatformApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationSpec.
func (in *SnsPlatformApplicationSpec) DeepCopy() *SnsPlatformApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsPlatformApplicationStatus) DeepCopyInto(out *SnsPlatformApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsPlatformApplicationStatus.
func (in *SnsPlatformApplicationStatus) DeepCopy() *SnsPlatformApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(SnsPlatformApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsSmsPreferences) DeepCopyInto(out *SnsSmsPreferences) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsSmsPreferences.
func (in *SnsSmsPreferences) DeepCopy() *SnsSmsPreferences {
	if in == nil {
		return nil
	}
	out := new(SnsSmsPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsSmsPreferences) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsSmsPreferencesList) DeepCopyInto(out *SnsSmsPreferencesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsSmsPreferences, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsSmsPreferencesList.
func (in *SnsSmsPreferencesList) DeepCopy() *SnsSmsPreferencesList {
	if in == nil {
		return nil
	}
	out := new(SnsSmsPreferencesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsSmsPreferencesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsSmsPreferencesObservation) DeepCopyInto(out *SnsSmsPreferencesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsSmsPreferencesObservation.
func (in *SnsSmsPreferencesObservation) DeepCopy() *SnsSmsPreferencesObservation {
	if in == nil {
		return nil
	}
	out := new(SnsSmsPreferencesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsSmsPreferencesParameters) DeepCopyInto(out *SnsSmsPreferencesParameters) {
	*out = *in
	if in.DefaultSenderID != nil {
		in, out := &in.DefaultSenderID, &out.DefaultSenderID
		*out = new(string)
		**out = **in
	}
	if in.DefaultSmsType != nil {
		in, out := &in.DefaultSmsType, &out.DefaultSmsType
		*out = new(string)
		**out = **in
	}
	if in.DeliveryStatusIamRoleArn != nil {
		in, out := &in.DeliveryStatusIamRoleArn, &out.DeliveryStatusIamRoleArn
		*out = new(string)
		**out = **in
	}
	if in.DeliveryStatusSuccessSamplingRate != nil {
		in, out := &in.DeliveryStatusSuccessSamplingRate, &out.DeliveryStatusSuccessSamplingRate
		*out = new(string)
		**out = **in
	}
	if in.MonthlySpendLimit != nil {
		in, out := &in.MonthlySpendLimit, &out.MonthlySpendLimit
		*out = new(string)
		**out = **in
	}
	if in.UsageReportS3Bucket != nil {
		in, out := &in.UsageReportS3Bucket, &out.UsageReportS3Bucket
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsSmsPreferencesParameters.
func (in *SnsSmsPreferencesParameters) DeepCopy() *SnsSmsPreferencesParameters {
	if in == nil {
		return nil
	}
	out := new(SnsSmsPreferencesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsSmsPreferencesSpec) DeepCopyInto(out *SnsSmsPreferencesSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsSmsPreferencesSpec.
func (in *SnsSmsPreferencesSpec) DeepCopy() *SnsSmsPreferencesSpec {
	if in == nil {
		return nil
	}
	out := new(SnsSmsPreferencesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsSmsPreferencesStatus) DeepCopyInto(out *SnsSmsPreferencesStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsSmsPreferencesStatus.
func (in *SnsSmsPreferencesStatus) DeepCopy() *SnsSmsPreferencesStatus {
	if in == nil {
		return nil
	}
	out := new(SnsSmsPreferencesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopic) DeepCopyInto(out *SnsTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopic.
func (in *SnsTopic) DeepCopy() *SnsTopic {
	if in == nil {
		return nil
	}
	out := new(SnsTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicList) DeepCopyInto(out *SnsTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicList.
func (in *SnsTopicList) DeepCopy() *SnsTopicList {
	if in == nil {
		return nil
	}
	out := new(SnsTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicObservation) DeepCopyInto(out *SnsTopicObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicObservation.
func (in *SnsTopicObservation) DeepCopy() *SnsTopicObservation {
	if in == nil {
		return nil
	}
	out := new(SnsTopicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicParameters) DeepCopyInto(out *SnsTopicParameters) {
	*out = *in
	if in.ApplicationFailureFeedbackRoleArn != nil {
		in, out := &in.ApplicationFailureFeedbackRoleArn, &out.ApplicationFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationSuccessFeedbackRoleArn != nil {
		in, out := &in.ApplicationSuccessFeedbackRoleArn, &out.ApplicationSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.ApplicationSuccessFeedbackSampleRate != nil {
		in, out := &in.ApplicationSuccessFeedbackSampleRate, &out.ApplicationSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.ContentBasedDeduplication != nil {
		in, out := &in.ContentBasedDeduplication, &out.ContentBasedDeduplication
		*out = new(bool)
		**out = **in
	}
	if in.DeliveryPolicy != nil {
		in, out := &in.DeliveryPolicy, &out.DeliveryPolicy
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FifoTopic != nil {
		in, out := &in.FifoTopic, &out.FifoTopic
		*out = new(bool)
		**out = **in
	}
	if in.FirehoseFailureFeedbackRoleArn != nil {
		in, out := &in.FirehoseFailureFeedbackRoleArn, &out.FirehoseFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.FirehoseSuccessFeedbackRoleArn != nil {
		in, out := &in.FirehoseSuccessFeedbackRoleArn, &out.FirehoseSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.FirehoseSuccessFeedbackSampleRate != nil {
		in, out := &in.FirehoseSuccessFeedbackSampleRate, &out.FirehoseSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.HTTPFailureFeedbackRoleArn != nil {
		in, out := &in.HTTPFailureFeedbackRoleArn, &out.HTTPFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.HTTPSuccessFeedbackRoleArn != nil {
		in, out := &in.HTTPSuccessFeedbackRoleArn, &out.HTTPSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.HTTPSuccessFeedbackSampleRate != nil {
		in, out := &in.HTTPSuccessFeedbackSampleRate, &out.HTTPSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.KmsMasterKeyID != nil {
		in, out := &in.KmsMasterKeyID, &out.KmsMasterKeyID
		*out = new(string)
		**out = **in
	}
	if in.LambdaFailureFeedbackRoleArn != nil {
		in, out := &in.LambdaFailureFeedbackRoleArn, &out.LambdaFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.LambdaSuccessFeedbackRoleArn != nil {
		in, out := &in.LambdaSuccessFeedbackRoleArn, &out.LambdaSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.LambdaSuccessFeedbackSampleRate != nil {
		in, out := &in.LambdaSuccessFeedbackSampleRate, &out.LambdaSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.SqsFailureFeedbackRoleArn != nil {
		in, out := &in.SqsFailureFeedbackRoleArn, &out.SqsFailureFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.SqsSuccessFeedbackRoleArn != nil {
		in, out := &in.SqsSuccessFeedbackRoleArn, &out.SqsSuccessFeedbackRoleArn
		*out = new(string)
		**out = **in
	}
	if in.SqsSuccessFeedbackSampleRate != nil {
		in, out := &in.SqsSuccessFeedbackSampleRate, &out.SqsSuccessFeedbackSampleRate
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicParameters.
func (in *SnsTopicParameters) DeepCopy() *SnsTopicParameters {
	if in == nil {
		return nil
	}
	out := new(SnsTopicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicPolicy) DeepCopyInto(out *SnsTopicPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicPolicy.
func (in *SnsTopicPolicy) DeepCopy() *SnsTopicPolicy {
	if in == nil {
		return nil
	}
	out := new(SnsTopicPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicPolicyList) DeepCopyInto(out *SnsTopicPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsTopicPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicPolicyList.
func (in *SnsTopicPolicyList) DeepCopy() *SnsTopicPolicyList {
	if in == nil {
		return nil
	}
	out := new(SnsTopicPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicPolicyObservation) DeepCopyInto(out *SnsTopicPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicPolicyObservation.
func (in *SnsTopicPolicyObservation) DeepCopy() *SnsTopicPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SnsTopicPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicPolicyParameters) DeepCopyInto(out *SnsTopicPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicPolicyParameters.
func (in *SnsTopicPolicyParameters) DeepCopy() *SnsTopicPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SnsTopicPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicPolicySpec) DeepCopyInto(out *SnsTopicPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicPolicySpec.
func (in *SnsTopicPolicySpec) DeepCopy() *SnsTopicPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SnsTopicPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicPolicyStatus) DeepCopyInto(out *SnsTopicPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicPolicyStatus.
func (in *SnsTopicPolicyStatus) DeepCopy() *SnsTopicPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SnsTopicPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSpec) DeepCopyInto(out *SnsTopicSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSpec.
func (in *SnsTopicSpec) DeepCopy() *SnsTopicSpec {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicStatus) DeepCopyInto(out *SnsTopicStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicStatus.
func (in *SnsTopicStatus) DeepCopy() *SnsTopicStatus {
	if in == nil {
		return nil
	}
	out := new(SnsTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscription) DeepCopyInto(out *SnsTopicSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscription.
func (in *SnsTopicSubscription) DeepCopy() *SnsTopicSubscription {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionList) DeepCopyInto(out *SnsTopicSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnsTopicSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionList.
func (in *SnsTopicSubscriptionList) DeepCopy() *SnsTopicSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnsTopicSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionObservation) DeepCopyInto(out *SnsTopicSubscriptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionObservation.
func (in *SnsTopicSubscriptionObservation) DeepCopy() *SnsTopicSubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionParameters) DeepCopyInto(out *SnsTopicSubscriptionParameters) {
	*out = *in
	if in.ConfirmationTimeoutInMinutes != nil {
		in, out := &in.ConfirmationTimeoutInMinutes, &out.ConfirmationTimeoutInMinutes
		*out = new(int64)
		**out = **in
	}
	if in.DeliveryPolicy != nil {
		in, out := &in.DeliveryPolicy, &out.DeliveryPolicy
		*out = new(string)
		**out = **in
	}
	if in.EndpointAutoConfirms != nil {
		in, out := &in.EndpointAutoConfirms, &out.EndpointAutoConfirms
		*out = new(bool)
		**out = **in
	}
	if in.FilterPolicy != nil {
		in, out := &in.FilterPolicy, &out.FilterPolicy
		*out = new(string)
		**out = **in
	}
	if in.RawMessageDelivery != nil {
		in, out := &in.RawMessageDelivery, &out.RawMessageDelivery
		*out = new(bool)
		**out = **in
	}
	if in.RedrivePolicy != nil {
		in, out := &in.RedrivePolicy, &out.RedrivePolicy
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionRoleArn != nil {
		in, out := &in.SubscriptionRoleArn, &out.SubscriptionRoleArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionParameters.
func (in *SnsTopicSubscriptionParameters) DeepCopy() *SnsTopicSubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionSpec) DeepCopyInto(out *SnsTopicSubscriptionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionSpec.
func (in *SnsTopicSubscriptionSpec) DeepCopy() *SnsTopicSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnsTopicSubscriptionStatus) DeepCopyInto(out *SnsTopicSubscriptionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnsTopicSubscriptionStatus.
func (in *SnsTopicSubscriptionStatus) DeepCopy() *SnsTopicSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SnsTopicSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
