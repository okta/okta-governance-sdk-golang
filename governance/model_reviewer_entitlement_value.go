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

// ReviewerEntitlementValue Entitlement value. Only applicable if `resourceType = APPLICATION` and Entitlement Management is enabled.
type ReviewerEntitlementValue struct {
	// The entitlement value id
	Id string `json:"id"`
	// The entitlement value name
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerEntitlementValue ReviewerEntitlementValue

// NewReviewerEntitlementValue instantiates a new ReviewerEntitlementValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerEntitlementValue(id string, name string) *ReviewerEntitlementValue {
	this := ReviewerEntitlementValue{}
	this.Id = id
	this.Name = name
	return &this
}

// NewReviewerEntitlementValueWithDefaults instantiates a new ReviewerEntitlementValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerEntitlementValueWithDefaults() *ReviewerEntitlementValue {
	this := ReviewerEntitlementValue{}
	return &this
}

// GetId returns the Id field value
func (o *ReviewerEntitlementValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewerEntitlementValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewerEntitlementValue) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ReviewerEntitlementValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReviewerEntitlementValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReviewerEntitlementValue) SetName(v string) {
	o.Name = v
}

func (o ReviewerEntitlementValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewerEntitlementValue) UnmarshalJSON(bytes []byte) (err error) {
	varReviewerEntitlementValue := _ReviewerEntitlementValue{}

	err = json.Unmarshal(bytes, &varReviewerEntitlementValue)
	if err == nil {
		*o = ReviewerEntitlementValue(varReviewerEntitlementValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewerEntitlementValue struct {
	value *ReviewerEntitlementValue
	isSet bool
}

func (v NullableReviewerEntitlementValue) Get() *ReviewerEntitlementValue {
	return v.value
}

func (v *NullableReviewerEntitlementValue) Set(val *ReviewerEntitlementValue) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerEntitlementValue) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerEntitlementValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerEntitlementValue(val *ReviewerEntitlementValue) *NullableReviewerEntitlementValue {
	return &NullableReviewerEntitlementValue{value: val, isSet: true}
}

func (v NullableReviewerEntitlementValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerEntitlementValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
