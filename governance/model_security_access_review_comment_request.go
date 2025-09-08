/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 3.2.0
Contact: devex-public@okta.com
*/

package governance

import (
	"encoding/json"
)

// SecurityAccessReviewCommentRequest struct for SecurityAccessReviewCommentRequest
type SecurityAccessReviewCommentRequest struct {
	// A comment for the security access review
	Comment              string `json:"comment"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewCommentRequest SecurityAccessReviewCommentRequest

// NewSecurityAccessReviewCommentRequest instantiates a new SecurityAccessReviewCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewCommentRequest(comment string) *SecurityAccessReviewCommentRequest {
	this := SecurityAccessReviewCommentRequest{}
	this.Comment = comment
	return &this
}

// NewSecurityAccessReviewCommentRequestWithDefaults instantiates a new SecurityAccessReviewCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewCommentRequestWithDefaults() *SecurityAccessReviewCommentRequest {
	this := SecurityAccessReviewCommentRequest{}
	return &this
}

// GetComment returns the Comment field value
func (o *SecurityAccessReviewCommentRequest) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewCommentRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *SecurityAccessReviewCommentRequest) SetComment(v string) {
	o.Comment = v
}

func (o SecurityAccessReviewCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewCommentRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewCommentRequest := _SecurityAccessReviewCommentRequest{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewCommentRequest)
	if err == nil {
		*o = SecurityAccessReviewCommentRequest(varSecurityAccessReviewCommentRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewCommentRequest struct {
	value *SecurityAccessReviewCommentRequest
	isSet bool
}

func (v NullableSecurityAccessReviewCommentRequest) Get() *SecurityAccessReviewCommentRequest {
	return v.value
}

func (v *NullableSecurityAccessReviewCommentRequest) Set(val *SecurityAccessReviewCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewCommentRequest(val *SecurityAccessReviewCommentRequest) *NullableSecurityAccessReviewCommentRequest {
	return &NullableSecurityAccessReviewCommentRequest{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
