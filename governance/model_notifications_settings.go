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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NotificationsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettings{}

// NotificationsSettings Notification settings for an integration
type NotificationsSettings struct {
	// Indicates that notifications are enabled for this integration
	Enabled bool `json:"enabled"`
}

type _NotificationsSettings NotificationsSettings

// NewNotificationsSettings instantiates a new NotificationsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettings(enabled bool) *NotificationsSettings {
	this := NotificationsSettings{}
	this.Enabled = enabled
	return &this
}

// NewNotificationsSettingsWithDefaults instantiates a new NotificationsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsWithDefaults() *NotificationsSettings {
	this := NotificationsSettings{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *NotificationsSettings) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *NotificationsSettings) SetEnabled(v bool) {
	o.Enabled = v
}

func (o NotificationsSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *NotificationsSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varNotificationsSettings := _NotificationsSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationsSettings)

	if err != nil {
		return err
	}

	*o = NotificationsSettings(varNotificationsSettings)

	return err
}

type NullableNotificationsSettings struct {
	value *NotificationsSettings
	isSet bool
}

func (v NullableNotificationsSettings) Get() *NotificationsSettings {
	return v.value
}

func (v *NullableNotificationsSettings) Set(val *NotificationsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettings(val *NotificationsSettings) *NullableNotificationsSettings {
	return &NullableNotificationsSettings{value: val, isSet: true}
}

func (v NullableNotificationsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
