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

// RemediationSettings Specify the action to be taken after a reviewer makes a decision to `APPROVE` or `REVOKE` the access, or if the campaign was `CLOSED` and there was no response from the reviewer.
type RemediationSettings struct {
	AccessApproved          ApprovedRemediationAction   `json:"accessApproved"`
	AccessRevoked           RevokedRemediationAction    `json:"accessRevoked"`
	NoResponse              NoResponseRemediationAction `json:"noResponse"`
	AutoRemediationSettings *AutoRemediationSettings    `json:"autoRemediationSettings,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _RemediationSettings RemediationSettings

// NewRemediationSettings instantiates a new RemediationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediationSettings(accessApproved ApprovedRemediationAction, accessRevoked RevokedRemediationAction, noResponse NoResponseRemediationAction) *RemediationSettings {
	this := RemediationSettings{}
	this.AccessApproved = accessApproved
	this.AccessRevoked = accessRevoked
	this.NoResponse = noResponse
	return &this
}

// NewRemediationSettingsWithDefaults instantiates a new RemediationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediationSettingsWithDefaults() *RemediationSettings {
	this := RemediationSettings{}
	var accessApproved ApprovedRemediationAction = APPROVEDREMEDIATIONACTION_NO_ACTION
	this.AccessApproved = accessApproved
	var accessRevoked RevokedRemediationAction = REVOKEDREMEDIATIONACTION_NO_ACTION
	this.AccessRevoked = accessRevoked
	var noResponse NoResponseRemediationAction = NORESPONSEREMEDIATIONACTION_NO_ACTION
	this.NoResponse = noResponse
	return &this
}

// GetAccessApproved returns the AccessApproved field value
func (o *RemediationSettings) GetAccessApproved() ApprovedRemediationAction {
	if o == nil {
		var ret ApprovedRemediationAction
		return ret
	}

	return o.AccessApproved
}

// GetAccessApprovedOk returns a tuple with the AccessApproved field value
// and a boolean to check if the value has been set.
func (o *RemediationSettings) GetAccessApprovedOk() (*ApprovedRemediationAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessApproved, true
}

// SetAccessApproved sets field value
func (o *RemediationSettings) SetAccessApproved(v ApprovedRemediationAction) {
	o.AccessApproved = v
}

// GetAccessRevoked returns the AccessRevoked field value
func (o *RemediationSettings) GetAccessRevoked() RevokedRemediationAction {
	if o == nil {
		var ret RevokedRemediationAction
		return ret
	}

	return o.AccessRevoked
}

// GetAccessRevokedOk returns a tuple with the AccessRevoked field value
// and a boolean to check if the value has been set.
func (o *RemediationSettings) GetAccessRevokedOk() (*RevokedRemediationAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessRevoked, true
}

// SetAccessRevoked sets field value
func (o *RemediationSettings) SetAccessRevoked(v RevokedRemediationAction) {
	o.AccessRevoked = v
}

// GetNoResponse returns the NoResponse field value
func (o *RemediationSettings) GetNoResponse() NoResponseRemediationAction {
	if o == nil {
		var ret NoResponseRemediationAction
		return ret
	}

	return o.NoResponse
}

// GetNoResponseOk returns a tuple with the NoResponse field value
// and a boolean to check if the value has been set.
func (o *RemediationSettings) GetNoResponseOk() (*NoResponseRemediationAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoResponse, true
}

// SetNoResponse sets field value
func (o *RemediationSettings) SetNoResponse(v NoResponseRemediationAction) {
	o.NoResponse = v
}

// GetAutoRemediationSettings returns the AutoRemediationSettings field value if set, zero value otherwise.
func (o *RemediationSettings) GetAutoRemediationSettings() AutoRemediationSettings {
	if o == nil || o.AutoRemediationSettings == nil {
		var ret AutoRemediationSettings
		return ret
	}
	return *o.AutoRemediationSettings
}

// GetAutoRemediationSettingsOk returns a tuple with the AutoRemediationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationSettings) GetAutoRemediationSettingsOk() (*AutoRemediationSettings, bool) {
	if o == nil || o.AutoRemediationSettings == nil {
		return nil, false
	}
	return o.AutoRemediationSettings, true
}

// HasAutoRemediationSettings returns a boolean if a field has been set.
func (o *RemediationSettings) HasAutoRemediationSettings() bool {
	if o != nil && o.AutoRemediationSettings != nil {
		return true
	}

	return false
}

// SetAutoRemediationSettings gets a reference to the given AutoRemediationSettings and assigns it to the AutoRemediationSettings field.
func (o *RemediationSettings) SetAutoRemediationSettings(v AutoRemediationSettings) {
	o.AutoRemediationSettings = &v
}

func (o RemediationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessApproved"] = o.AccessApproved
	}
	if true {
		toSerialize["accessRevoked"] = o.AccessRevoked
	}
	if true {
		toSerialize["noResponse"] = o.NoResponse
	}
	if o.AutoRemediationSettings != nil {
		toSerialize["autoRemediationSettings"] = o.AutoRemediationSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RemediationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varRemediationSettings := _RemediationSettings{}

	err = json.Unmarshal(bytes, &varRemediationSettings)
	if err == nil {
		*o = RemediationSettings(varRemediationSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accessApproved")
		delete(additionalProperties, "accessRevoked")
		delete(additionalProperties, "noResponse")
		delete(additionalProperties, "autoRemediationSettings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRemediationSettings struct {
	value *RemediationSettings
	isSet bool
}

func (v NullableRemediationSettings) Get() *RemediationSettings {
	return v.value
}

func (v *NullableRemediationSettings) Set(val *RemediationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediationSettings(val *RemediationSettings) *NullableRemediationSettings {
	return &NullableRemediationSettings{value: val, isSet: true}
}

func (v NullableRemediationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
