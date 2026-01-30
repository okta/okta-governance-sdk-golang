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

// checks if the GrantTypePolicyWriteable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantTypePolicyWriteable{}

// GrantTypePolicyWriteable Configured policy as the source for grant request
type GrantTypePolicyWriteable struct {
	// Base grant type for assigning entitlements driven by policy
	GrantType            string                     `json:"grantType"`
	Target               TargetResource             `json:"target"`
	TargetPrincipal      TargetPrincipal            `json:"targetPrincipal"`
	ScheduleSettings     *ScheduleSettingsWriteable `json:"scheduleSettings,omitempty"`
	Action               *GrantAction               `json:"action,omitempty"`
	Actor                *GrantActor                `json:"actor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantTypePolicyWriteable GrantTypePolicyWriteable

// NewGrantTypePolicyWriteable instantiates a new GrantTypePolicyWriteable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantTypePolicyWriteable(grantType string, target TargetResource, targetPrincipal TargetPrincipal) *GrantTypePolicyWriteable {
	this := GrantTypePolicyWriteable{}
	this.TargetPrincipal = targetPrincipal
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// NewGrantTypePolicyWriteableWithDefaults instantiates a new GrantTypePolicyWriteable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantTypePolicyWriteableWithDefaults() *GrantTypePolicyWriteable {
	this := GrantTypePolicyWriteable{}
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetGrantType returns the GrantType field value
func (o *GrantTypePolicyWriteable) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *GrantTypePolicyWriteable) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *GrantTypePolicyWriteable) SetGrantType(v string) {
	o.GrantType = v
}

// GetTarget returns the Target field value
func (o *GrantTypePolicyWriteable) GetTarget() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *GrantTypePolicyWriteable) GetTargetOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *GrantTypePolicyWriteable) SetTarget(v TargetResource) {
	o.Target = v
}

// GetTargetPrincipal returns the TargetPrincipal field value
func (o *GrantTypePolicyWriteable) GetTargetPrincipal() TargetPrincipal {
	if o == nil {
		var ret TargetPrincipal
		return ret
	}

	return o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value
// and a boolean to check if the value has been set.
func (o *GrantTypePolicyWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipal, true
}

// SetTargetPrincipal sets field value
func (o *GrantTypePolicyWriteable) SetTargetPrincipal(v TargetPrincipal) {
	o.TargetPrincipal = v
}

// GetScheduleSettings returns the ScheduleSettings field value if set, zero value otherwise.
func (o *GrantTypePolicyWriteable) GetScheduleSettings() ScheduleSettingsWriteable {
	if o == nil || IsNil(o.ScheduleSettings) {
		var ret ScheduleSettingsWriteable
		return ret
	}
	return *o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypePolicyWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool) {
	if o == nil || IsNil(o.ScheduleSettings) {
		return nil, false
	}
	return o.ScheduleSettings, true
}

// HasScheduleSettings returns a boolean if a field has been set.
func (o *GrantTypePolicyWriteable) HasScheduleSettings() bool {
	if o != nil && !IsNil(o.ScheduleSettings) {
		return true
	}

	return false
}

// SetScheduleSettings gets a reference to the given ScheduleSettingsWriteable and assigns it to the ScheduleSettings field.
func (o *GrantTypePolicyWriteable) SetScheduleSettings(v ScheduleSettingsWriteable) {
	o.ScheduleSettings = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GrantTypePolicyWriteable) GetAction() GrantAction {
	if o == nil || IsNil(o.Action) {
		var ret GrantAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypePolicyWriteable) GetActionOk() (*GrantAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GrantTypePolicyWriteable) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given GrantAction and assigns it to the Action field.
func (o *GrantTypePolicyWriteable) SetAction(v GrantAction) {
	o.Action = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *GrantTypePolicyWriteable) GetActor() GrantActor {
	if o == nil || IsNil(o.Actor) {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypePolicyWriteable) GetActorOk() (*GrantActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *GrantTypePolicyWriteable) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *GrantTypePolicyWriteable) SetActor(v GrantActor) {
	o.Actor = &v
}

func (o GrantTypePolicyWriteable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantTypePolicyWriteable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grantType"] = o.GrantType
	toSerialize["target"] = o.Target
	toSerialize["targetPrincipal"] = o.TargetPrincipal
	if !IsNil(o.ScheduleSettings) {
		toSerialize["scheduleSettings"] = o.ScheduleSettings
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantTypePolicyWriteable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grantType",
		"target",
		"targetPrincipal",
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

	varGrantTypePolicyWriteable := _GrantTypePolicyWriteable{}

	err = json.Unmarshal(data, &varGrantTypePolicyWriteable)

	if err != nil {
		return err
	}

	*o = GrantTypePolicyWriteable(varGrantTypePolicyWriteable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "target")
		delete(additionalProperties, "targetPrincipal")
		delete(additionalProperties, "scheduleSettings")
		delete(additionalProperties, "action")
		delete(additionalProperties, "actor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantTypePolicyWriteable struct {
	value *GrantTypePolicyWriteable
	isSet bool
}

func (v NullableGrantTypePolicyWriteable) Get() *GrantTypePolicyWriteable {
	return v.value
}

func (v *NullableGrantTypePolicyWriteable) Set(val *GrantTypePolicyWriteable) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantTypePolicyWriteable) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantTypePolicyWriteable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantTypePolicyWriteable(val *GrantTypePolicyWriteable) *NullableGrantTypePolicyWriteable {
	return &NullableGrantTypePolicyWriteable{value: val, isSet: true}
}

func (v NullableGrantTypePolicyWriteable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantTypePolicyWriteable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
