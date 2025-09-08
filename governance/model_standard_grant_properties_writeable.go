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

// StandardGrantPropertiesWriteable struct for StandardGrantPropertiesWriteable
type StandardGrantPropertiesWriteable struct {
	TargetPrincipal      TargetPrincipal            `json:"targetPrincipal"`
	ScheduleSettings     *ScheduleSettingsWriteable `json:"scheduleSettings,omitempty"`
	Action               *GrantAction               `json:"action,omitempty"`
	Actor                *GrantActor                `json:"actor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StandardGrantPropertiesWriteable StandardGrantPropertiesWriteable

// NewStandardGrantPropertiesWriteable instantiates a new StandardGrantPropertiesWriteable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardGrantPropertiesWriteable(targetPrincipal TargetPrincipal) *StandardGrantPropertiesWriteable {
	this := StandardGrantPropertiesWriteable{}
	this.TargetPrincipal = targetPrincipal
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// NewStandardGrantPropertiesWriteableWithDefaults instantiates a new StandardGrantPropertiesWriteable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardGrantPropertiesWriteableWithDefaults() *StandardGrantPropertiesWriteable {
	this := StandardGrantPropertiesWriteable{}
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetTargetPrincipal returns the TargetPrincipal field value
func (o *StandardGrantPropertiesWriteable) GetTargetPrincipal() TargetPrincipal {
	if o == nil {
		var ret TargetPrincipal
		return ret
	}

	return o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value
// and a boolean to check if the value has been set.
func (o *StandardGrantPropertiesWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipal, true
}

// SetTargetPrincipal sets field value
func (o *StandardGrantPropertiesWriteable) SetTargetPrincipal(v TargetPrincipal) {
	o.TargetPrincipal = v
}

// GetScheduleSettings returns the ScheduleSettings field value if set, zero value otherwise.
func (o *StandardGrantPropertiesWriteable) GetScheduleSettings() ScheduleSettingsWriteable {
	if o == nil || o.ScheduleSettings == nil {
		var ret ScheduleSettingsWriteable
		return ret
	}
	return *o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardGrantPropertiesWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool) {
	if o == nil || o.ScheduleSettings == nil {
		return nil, false
	}
	return o.ScheduleSettings, true
}

// HasScheduleSettings returns a boolean if a field has been set.
func (o *StandardGrantPropertiesWriteable) HasScheduleSettings() bool {
	if o != nil && o.ScheduleSettings != nil {
		return true
	}

	return false
}

// SetScheduleSettings gets a reference to the given ScheduleSettingsWriteable and assigns it to the ScheduleSettings field.
func (o *StandardGrantPropertiesWriteable) SetScheduleSettings(v ScheduleSettingsWriteable) {
	o.ScheduleSettings = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *StandardGrantPropertiesWriteable) GetAction() GrantAction {
	if o == nil || o.Action == nil {
		var ret GrantAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardGrantPropertiesWriteable) GetActionOk() (*GrantAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *StandardGrantPropertiesWriteable) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given GrantAction and assigns it to the Action field.
func (o *StandardGrantPropertiesWriteable) SetAction(v GrantAction) {
	o.Action = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *StandardGrantPropertiesWriteable) GetActor() GrantActor {
	if o == nil || o.Actor == nil {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardGrantPropertiesWriteable) GetActorOk() (*GrantActor, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *StandardGrantPropertiesWriteable) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *StandardGrantPropertiesWriteable) SetActor(v GrantActor) {
	o.Actor = &v
}

func (o StandardGrantPropertiesWriteable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["targetPrincipal"] = o.TargetPrincipal
	}
	if o.ScheduleSettings != nil {
		toSerialize["scheduleSettings"] = o.ScheduleSettings
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StandardGrantPropertiesWriteable) UnmarshalJSON(bytes []byte) (err error) {
	varStandardGrantPropertiesWriteable := _StandardGrantPropertiesWriteable{}

	err = json.Unmarshal(bytes, &varStandardGrantPropertiesWriteable)
	if err == nil {
		*o = StandardGrantPropertiesWriteable(varStandardGrantPropertiesWriteable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "targetPrincipal")
		delete(additionalProperties, "scheduleSettings")
		delete(additionalProperties, "action")
		delete(additionalProperties, "actor")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableStandardGrantPropertiesWriteable struct {
	value *StandardGrantPropertiesWriteable
	isSet bool
}

func (v NullableStandardGrantPropertiesWriteable) Get() *StandardGrantPropertiesWriteable {
	return v.value
}

func (v *NullableStandardGrantPropertiesWriteable) Set(val *StandardGrantPropertiesWriteable) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardGrantPropertiesWriteable) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardGrantPropertiesWriteable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardGrantPropertiesWriteable(val *StandardGrantPropertiesWriteable) *NullableStandardGrantPropertiesWriteable {
	return &NullableStandardGrantPropertiesWriteable{value: val, isSet: true}
}

func (v NullableStandardGrantPropertiesWriteable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardGrantPropertiesWriteable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
