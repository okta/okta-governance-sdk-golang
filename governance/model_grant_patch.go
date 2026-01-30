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

// checks if the GrantPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPatch{}

// GrantPatch struct for GrantPatch
type GrantPatch struct {
	// Unique identifier for the object
	Id                   string                    `json:"id"`
	ScheduleSettings     ScheduleSettingsWriteable `json:"scheduleSettings"`
	AdditionalProperties map[string]interface{}
}

type _GrantPatch GrantPatch

// NewGrantPatch instantiates a new GrantPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPatch(id string, scheduleSettings ScheduleSettingsWriteable) *GrantPatch {
	this := GrantPatch{}
	this.Id = id
	this.ScheduleSettings = scheduleSettings
	return &this
}

// NewGrantPatchWithDefaults instantiates a new GrantPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPatchWithDefaults() *GrantPatch {
	this := GrantPatch{}
	return &this
}

// GetId returns the Id field value
func (o *GrantPatch) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GrantPatch) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GrantPatch) SetId(v string) {
	o.Id = v
}

// GetScheduleSettings returns the ScheduleSettings field value
func (o *GrantPatch) GetScheduleSettings() ScheduleSettingsWriteable {
	if o == nil {
		var ret ScheduleSettingsWriteable
		return ret
	}

	return o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value
// and a boolean to check if the value has been set.
func (o *GrantPatch) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleSettings, true
}

// SetScheduleSettings sets field value
func (o *GrantPatch) SetScheduleSettings(v ScheduleSettingsWriteable) {
	o.ScheduleSettings = v
}

func (o GrantPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["scheduleSettings"] = o.ScheduleSettings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"scheduleSettings",
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

	varGrantPatch := _GrantPatch{}

	err = json.Unmarshal(data, &varGrantPatch)

	if err != nil {
		return err
	}

	*o = GrantPatch(varGrantPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "scheduleSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPatch struct {
	value *GrantPatch
	isSet bool
}

func (v NullableGrantPatch) Get() *GrantPatch {
	return v.value
}

func (v *NullableGrantPatch) Set(val *GrantPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPatch(val *GrantPatch) *NullableGrantPatch {
	return &NullableGrantPatch{value: val, isSet: true}
}

func (v NullableGrantPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
