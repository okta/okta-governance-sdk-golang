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

// SecurityAccessReviewReviewers struct for SecurityAccessReviewReviewers
type SecurityAccessReviewReviewers struct {
	// Security access review reviewers
	Data                 []PrincipalProfile `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewReviewers SecurityAccessReviewReviewers

// NewSecurityAccessReviewReviewers instantiates a new SecurityAccessReviewReviewers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewReviewers(data []PrincipalProfile) *SecurityAccessReviewReviewers {
	this := SecurityAccessReviewReviewers{}
	this.Data = data
	return &this
}

// NewSecurityAccessReviewReviewersWithDefaults instantiates a new SecurityAccessReviewReviewers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewReviewersWithDefaults() *SecurityAccessReviewReviewers {
	this := SecurityAccessReviewReviewers{}
	return &this
}

// GetData returns the Data field value
func (o *SecurityAccessReviewReviewers) GetData() []PrincipalProfile {
	if o == nil {
		var ret []PrincipalProfile
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewReviewers) GetDataOk() ([]PrincipalProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SecurityAccessReviewReviewers) SetData(v []PrincipalProfile) {
	o.Data = v
}

func (o SecurityAccessReviewReviewers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewReviewers) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewReviewers := _SecurityAccessReviewReviewers{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewReviewers)
	if err == nil {
		*o = SecurityAccessReviewReviewers(varSecurityAccessReviewReviewers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewReviewers struct {
	value *SecurityAccessReviewReviewers
	isSet bool
}

func (v NullableSecurityAccessReviewReviewers) Get() *SecurityAccessReviewReviewers {
	return v.value
}

func (v *NullableSecurityAccessReviewReviewers) Set(val *SecurityAccessReviewReviewers) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewReviewers) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewReviewers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewReviewers(val *SecurityAccessReviewReviewers) *NullableSecurityAccessReviewReviewers {
	return &NullableSecurityAccessReviewReviewers{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewReviewers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewReviewers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
