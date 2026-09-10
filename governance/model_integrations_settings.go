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

// checks if the IntegrationsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationsSettings{}

// IntegrationsSettings Integration settings for a specific integration
type IntegrationsSettings struct {
	// Integration type
	Type string `json:"type"`
	// The integration ID
	IntegrationId string                 `json:"integrationId" validate:"regexp=goi[0-9a-zA-Z]+"`
	Notifications *NotificationsSettings `json:"notifications,omitempty"`
}

type _IntegrationsSettings IntegrationsSettings

// NewIntegrationsSettings instantiates a new IntegrationsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationsSettings(type_ string, integrationId string) *IntegrationsSettings {
	this := IntegrationsSettings{}
	this.Type = type_
	this.IntegrationId = integrationId
	return &this
}

// NewIntegrationsSettingsWithDefaults instantiates a new IntegrationsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationsSettingsWithDefaults() *IntegrationsSettings {
	this := IntegrationsSettings{}
	return &this
}

// GetType returns the Type field value
func (o *IntegrationsSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationsSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationsSettings) SetType(v string) {
	o.Type = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *IntegrationsSettings) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *IntegrationsSettings) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *IntegrationsSettings) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *IntegrationsSettings) GetNotifications() NotificationsSettings {
	if o == nil || IsNil(o.Notifications) {
		var ret NotificationsSettings
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationsSettings) GetNotificationsOk() (*NotificationsSettings, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *IntegrationsSettings) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given NotificationsSettings and assigns it to the Notifications field.
func (o *IntegrationsSettings) SetNotifications(v NotificationsSettings) {
	o.Notifications = &v
}

func (o IntegrationsSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

func (o *IntegrationsSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"integrationId",
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

	varIntegrationsSettings := _IntegrationsSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIntegrationsSettings)

	if err != nil {
		return err
	}

	*o = IntegrationsSettings(varIntegrationsSettings)

	return err
}

type NullableIntegrationsSettings struct {
	value *IntegrationsSettings
	isSet bool
}

func (v NullableIntegrationsSettings) Get() *IntegrationsSettings {
	return v.value
}

func (v *NullableIntegrationsSettings) Set(val *IntegrationsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationsSettings(val *IntegrationsSettings) *NullableIntegrationsSettings {
	return &NullableIntegrationsSettings{value: val, isSet: true}
}

func (v NullableIntegrationsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
