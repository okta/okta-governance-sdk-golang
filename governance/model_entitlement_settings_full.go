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

// checks if the EntitlementSettingsFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementSettingsFull{}

// EntitlementSettingsFull Entitlement settings for a resource
type EntitlementSettingsFull struct {
	Status               EntitlementSettingsResponseStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementSettingsFull EntitlementSettingsFull

// NewEntitlementSettingsFull instantiates a new EntitlementSettingsFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementSettingsFull(status EntitlementSettingsResponseStatus) *EntitlementSettingsFull {
	this := EntitlementSettingsFull{}
	this.Status = status
	return &this
}

// NewEntitlementSettingsFullWithDefaults instantiates a new EntitlementSettingsFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementSettingsFullWithDefaults() *EntitlementSettingsFull {
	this := EntitlementSettingsFull{}
	return &this
}

// GetStatus returns the Status field value
func (o *EntitlementSettingsFull) GetStatus() EntitlementSettingsResponseStatus {
	if o == nil {
		var ret EntitlementSettingsResponseStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EntitlementSettingsFull) GetStatusOk() (*EntitlementSettingsResponseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EntitlementSettingsFull) SetStatus(v EntitlementSettingsResponseStatus) {
	o.Status = v
}

func (o EntitlementSettingsFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementSettingsFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementSettingsFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varEntitlementSettingsFull := _EntitlementSettingsFull{}

	err = json.Unmarshal(data, &varEntitlementSettingsFull)

	if err != nil {
		return err
	}

	*o = EntitlementSettingsFull(varEntitlementSettingsFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementSettingsFull struct {
	value *EntitlementSettingsFull
	isSet bool
}

func (v NullableEntitlementSettingsFull) Get() *EntitlementSettingsFull {
	return v.value
}

func (v *NullableEntitlementSettingsFull) Set(val *EntitlementSettingsFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementSettingsFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementSettingsFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementSettingsFull(val *EntitlementSettingsFull) *NullableEntitlementSettingsFull {
	return &NullableEntitlementSettingsFull{value: val, isSet: true}
}

func (v NullableEntitlementSettingsFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementSettingsFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
