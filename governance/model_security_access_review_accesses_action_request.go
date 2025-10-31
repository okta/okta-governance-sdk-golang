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

// checks if the SecurityAccessReviewAccessesActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewAccessesActionRequest{}

// SecurityAccessReviewAccessesActionRequest struct for SecurityAccessReviewAccessesActionRequest
type SecurityAccessReviewAccessesActionRequest struct {
	Type                 SecurityAccessReviewAccessItemSupportedAction `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewAccessesActionRequest SecurityAccessReviewAccessesActionRequest

// NewSecurityAccessReviewAccessesActionRequest instantiates a new SecurityAccessReviewAccessesActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewAccessesActionRequest(type_ SecurityAccessReviewAccessItemSupportedAction) *SecurityAccessReviewAccessesActionRequest {
	this := SecurityAccessReviewAccessesActionRequest{}
	this.Type = type_
	return &this
}

// NewSecurityAccessReviewAccessesActionRequestWithDefaults instantiates a new SecurityAccessReviewAccessesActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewAccessesActionRequestWithDefaults() *SecurityAccessReviewAccessesActionRequest {
	this := SecurityAccessReviewAccessesActionRequest{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityAccessReviewAccessesActionRequest) GetType() SecurityAccessReviewAccessItemSupportedAction {
	if o == nil {
		var ret SecurityAccessReviewAccessItemSupportedAction
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessesActionRequest) GetTypeOk() (*SecurityAccessReviewAccessItemSupportedAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityAccessReviewAccessesActionRequest) SetType(v SecurityAccessReviewAccessItemSupportedAction) {
	o.Type = v
}

func (o SecurityAccessReviewAccessesActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewAccessesActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewAccessesActionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSecurityAccessReviewAccessesActionRequest := _SecurityAccessReviewAccessesActionRequest{}

	err = json.Unmarshal(data, &varSecurityAccessReviewAccessesActionRequest)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewAccessesActionRequest(varSecurityAccessReviewAccessesActionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewAccessesActionRequest struct {
	value *SecurityAccessReviewAccessesActionRequest
	isSet bool
}

func (v NullableSecurityAccessReviewAccessesActionRequest) Get() *SecurityAccessReviewAccessesActionRequest {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessesActionRequest) Set(val *SecurityAccessReviewAccessesActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessesActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessesActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessesActionRequest(val *SecurityAccessReviewAccessesActionRequest) *NullableSecurityAccessReviewAccessesActionRequest {
	return &NullableSecurityAccessReviewAccessesActionRequest{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessesActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessesActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
