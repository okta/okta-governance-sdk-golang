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
)

// checks if the SecurityAccessReviewActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewActions{}

// SecurityAccessReviewActions struct for SecurityAccessReviewActions
type SecurityAccessReviewActions struct {
	// Returns all possible actions for the given security access review
	Data                 []SecurityAccessReviewAction `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewActions SecurityAccessReviewActions

// NewSecurityAccessReviewActions instantiates a new SecurityAccessReviewActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewActions() *SecurityAccessReviewActions {
	this := SecurityAccessReviewActions{}
	return &this
}

// NewSecurityAccessReviewActionsWithDefaults instantiates a new SecurityAccessReviewActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewActionsWithDefaults() *SecurityAccessReviewActions {
	this := SecurityAccessReviewActions{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SecurityAccessReviewActions) GetData() []SecurityAccessReviewAction {
	if o == nil || IsNil(o.Data) {
		var ret []SecurityAccessReviewAction
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewActions) GetDataOk() ([]SecurityAccessReviewAction, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SecurityAccessReviewActions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SecurityAccessReviewAction and assigns it to the Data field.
func (o *SecurityAccessReviewActions) SetData(v []SecurityAccessReviewAction) {
	o.Data = v
}

func (o SecurityAccessReviewActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewActions) UnmarshalJSON(data []byte) (err error) {
	varSecurityAccessReviewActions := _SecurityAccessReviewActions{}

	err = json.Unmarshal(data, &varSecurityAccessReviewActions)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewActions(varSecurityAccessReviewActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewActions struct {
	value *SecurityAccessReviewActions
	isSet bool
}

func (v NullableSecurityAccessReviewActions) Get() *SecurityAccessReviewActions {
	return v.value
}

func (v *NullableSecurityAccessReviewActions) Set(val *SecurityAccessReviewActions) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewActions) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewActions(val *SecurityAccessReviewActions) *NullableSecurityAccessReviewActions {
	return &NullableSecurityAccessReviewActions{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
