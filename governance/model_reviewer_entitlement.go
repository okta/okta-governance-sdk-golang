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

// checks if the ReviewerEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerEntitlement{}

// ReviewerEntitlement Entitlement information
type ReviewerEntitlement struct {
	// The entitlement `id`
	Id string `json:"id"`
	// The entitlement display name
	Name string `json:"name"`
	// The value of the entitlement property
	ExternalValue        *string `json:"externalValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerEntitlement ReviewerEntitlement

// NewReviewerEntitlement instantiates a new ReviewerEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerEntitlement(id string, name string) *ReviewerEntitlement {
	this := ReviewerEntitlement{}
	this.Id = id
	this.Name = name
	return &this
}

// NewReviewerEntitlementWithDefaults instantiates a new ReviewerEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerEntitlementWithDefaults() *ReviewerEntitlement {
	this := ReviewerEntitlement{}
	return &this
}

// GetId returns the Id field value
func (o *ReviewerEntitlement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewerEntitlement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewerEntitlement) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ReviewerEntitlement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReviewerEntitlement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReviewerEntitlement) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *ReviewerEntitlement) GetExternalValue() string {
	if o == nil || IsNil(o.ExternalValue) {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerEntitlement) GetExternalValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalValue) {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *ReviewerEntitlement) HasExternalValue() bool {
	if o != nil && !IsNil(o.ExternalValue) {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *ReviewerEntitlement) SetExternalValue(v string) {
	o.ExternalValue = &v
}

func (o ReviewerEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.ExternalValue) {
		toSerialize["externalValue"] = o.ExternalValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerEntitlement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varReviewerEntitlement := _ReviewerEntitlement{}

	err = json.Unmarshal(data, &varReviewerEntitlement)

	if err != nil {
		return err
	}

	*o = ReviewerEntitlement(varReviewerEntitlement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerEntitlement struct {
	value *ReviewerEntitlement
	isSet bool
}

func (v NullableReviewerEntitlement) Get() *ReviewerEntitlement {
	return v.value
}

func (v *NullableReviewerEntitlement) Set(val *ReviewerEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerEntitlement(val *ReviewerEntitlement) *NullableReviewerEntitlement {
	return &NullableReviewerEntitlement{value: val, isSet: true}
}

func (v NullableReviewerEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
