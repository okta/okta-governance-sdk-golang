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

// checks if the SlackIntegrationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackIntegrationSettings{}

// SlackIntegrationSettings Slack integration settings for the org
type SlackIntegrationSettings struct {
	// The type of integration
	Type string `json:"type"`
	// The integration ID
	IntegrationId string                     `json:"integrationId" validate:"regexp=goi[0-9a-zA-Z]+"`
	Notifications *SlackNotificationSettings `json:"notifications,omitempty"`
	// Indicates whether requests can be initiated from this Slack integration
	CanInitiateRequest *bool `json:"canInitiateRequest,omitempty"`
	// Indicates whether requests can be approved from this Slack integration
	CanApproveRequest *bool `json:"canApproveRequest,omitempty"`
}

type _SlackIntegrationSettings SlackIntegrationSettings

// NewSlackIntegrationSettings instantiates a new SlackIntegrationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackIntegrationSettings(type_ string, integrationId string) *SlackIntegrationSettings {
	this := SlackIntegrationSettings{}
	this.Type = type_
	this.IntegrationId = integrationId
	return &this
}

// NewSlackIntegrationSettingsWithDefaults instantiates a new SlackIntegrationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackIntegrationSettingsWithDefaults() *SlackIntegrationSettings {
	this := SlackIntegrationSettings{}
	return &this
}

// GetType returns the Type field value
func (o *SlackIntegrationSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SlackIntegrationSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SlackIntegrationSettings) SetType(v string) {
	o.Type = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *SlackIntegrationSettings) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *SlackIntegrationSettings) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *SlackIntegrationSettings) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *SlackIntegrationSettings) GetNotifications() SlackNotificationSettings {
	if o == nil || IsNil(o.Notifications) {
		var ret SlackNotificationSettings
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationSettings) GetNotificationsOk() (*SlackNotificationSettings, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *SlackIntegrationSettings) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given SlackNotificationSettings and assigns it to the Notifications field.
func (o *SlackIntegrationSettings) SetNotifications(v SlackNotificationSettings) {
	o.Notifications = &v
}

// GetCanInitiateRequest returns the CanInitiateRequest field value if set, zero value otherwise.
func (o *SlackIntegrationSettings) GetCanInitiateRequest() bool {
	if o == nil || IsNil(o.CanInitiateRequest) {
		var ret bool
		return ret
	}
	return *o.CanInitiateRequest
}

// GetCanInitiateRequestOk returns a tuple with the CanInitiateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationSettings) GetCanInitiateRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.CanInitiateRequest) {
		return nil, false
	}
	return o.CanInitiateRequest, true
}

// HasCanInitiateRequest returns a boolean if a field has been set.
func (o *SlackIntegrationSettings) HasCanInitiateRequest() bool {
	if o != nil && !IsNil(o.CanInitiateRequest) {
		return true
	}

	return false
}

// SetCanInitiateRequest gets a reference to the given bool and assigns it to the CanInitiateRequest field.
func (o *SlackIntegrationSettings) SetCanInitiateRequest(v bool) {
	o.CanInitiateRequest = &v
}

// GetCanApproveRequest returns the CanApproveRequest field value if set, zero value otherwise.
func (o *SlackIntegrationSettings) GetCanApproveRequest() bool {
	if o == nil || IsNil(o.CanApproveRequest) {
		var ret bool
		return ret
	}
	return *o.CanApproveRequest
}

// GetCanApproveRequestOk returns a tuple with the CanApproveRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationSettings) GetCanApproveRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.CanApproveRequest) {
		return nil, false
	}
	return o.CanApproveRequest, true
}

// HasCanApproveRequest returns a boolean if a field has been set.
func (o *SlackIntegrationSettings) HasCanApproveRequest() bool {
	if o != nil && !IsNil(o.CanApproveRequest) {
		return true
	}

	return false
}

// SetCanApproveRequest gets a reference to the given bool and assigns it to the CanApproveRequest field.
func (o *SlackIntegrationSettings) SetCanApproveRequest(v bool) {
	o.CanApproveRequest = &v
}

func (o SlackIntegrationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackIntegrationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.CanInitiateRequest) {
		toSerialize["canInitiateRequest"] = o.CanInitiateRequest
	}
	if !IsNil(o.CanApproveRequest) {
		toSerialize["canApproveRequest"] = o.CanApproveRequest
	}
	return toSerialize, nil
}

func (o *SlackIntegrationSettings) UnmarshalJSON(data []byte) (err error) {
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

	varSlackIntegrationSettings := _SlackIntegrationSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlackIntegrationSettings)

	if err != nil {
		return err
	}

	*o = SlackIntegrationSettings(varSlackIntegrationSettings)

	return err
}

type NullableSlackIntegrationSettings struct {
	value *SlackIntegrationSettings
	isSet bool
}

func (v NullableSlackIntegrationSettings) Get() *SlackIntegrationSettings {
	return v.value
}

func (v *NullableSlackIntegrationSettings) Set(val *SlackIntegrationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackIntegrationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackIntegrationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackIntegrationSettings(val *SlackIntegrationSettings) *NullableSlackIntegrationSettings {
	return &NullableSlackIntegrationSettings{value: val, isSet: true}
}

func (v NullableSlackIntegrationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackIntegrationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
