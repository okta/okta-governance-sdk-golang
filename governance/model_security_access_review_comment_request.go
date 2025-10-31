/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2025 - Present Okta, Inc.

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
	"fmt"
)

// checks if the SecurityAccessReviewCommentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewCommentRequest{}

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewCommentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comment"] = o.Comment

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewCommentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"comment",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSecurityAccessReviewCommentRequest := _SecurityAccessReviewCommentRequest{}

	err = json.Unmarshal(data, &varSecurityAccessReviewCommentRequest)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewCommentRequest(varSecurityAccessReviewCommentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
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
