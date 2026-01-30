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

// checks if the GrantTypeBundleWriteable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantTypeBundleWriteable{}

// GrantTypeBundleWriteable entitlement bundle as the grant source
type GrantTypeBundleWriteable struct {
	// Additive grant type for entitlement bundle.
	GrantType string `json:"grantType"`
	// The entitlement bundle `id`
	EntitlementBundleId  string                     `json:"entitlementBundleId" validate:"regexp=enb[0-9a-zA-Z]+"`
	TargetPrincipal      TargetPrincipal            `json:"targetPrincipal"`
	ScheduleSettings     *ScheduleSettingsWriteable `json:"scheduleSettings,omitempty"`
	Action               *GrantAction               `json:"action,omitempty"`
	Actor                *GrantActor                `json:"actor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantTypeBundleWriteable GrantTypeBundleWriteable

// NewGrantTypeBundleWriteable instantiates a new GrantTypeBundleWriteable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantTypeBundleWriteable(grantType string, entitlementBundleId string, targetPrincipal TargetPrincipal) *GrantTypeBundleWriteable {
	this := GrantTypeBundleWriteable{}
	this.TargetPrincipal = targetPrincipal
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// NewGrantTypeBundleWriteableWithDefaults instantiates a new GrantTypeBundleWriteable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantTypeBundleWriteableWithDefaults() *GrantTypeBundleWriteable {
	this := GrantTypeBundleWriteable{}
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetGrantType returns the GrantType field value
func (o *GrantTypeBundleWriteable) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *GrantTypeBundleWriteable) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *GrantTypeBundleWriteable) SetGrantType(v string) {
	o.GrantType = v
}

// GetEntitlementBundleId returns the EntitlementBundleId field value
func (o *GrantTypeBundleWriteable) GetEntitlementBundleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementBundleId
}

// GetEntitlementBundleIdOk returns a tuple with the EntitlementBundleId field value
// and a boolean to check if the value has been set.
func (o *GrantTypeBundleWriteable) GetEntitlementBundleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementBundleId, true
}

// SetEntitlementBundleId sets field value
func (o *GrantTypeBundleWriteable) SetEntitlementBundleId(v string) {
	o.EntitlementBundleId = v
}

// GetTargetPrincipal returns the TargetPrincipal field value
func (o *GrantTypeBundleWriteable) GetTargetPrincipal() TargetPrincipal {
	if o == nil {
		var ret TargetPrincipal
		return ret
	}

	return o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value
// and a boolean to check if the value has been set.
func (o *GrantTypeBundleWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipal, true
}

// SetTargetPrincipal sets field value
func (o *GrantTypeBundleWriteable) SetTargetPrincipal(v TargetPrincipal) {
	o.TargetPrincipal = v
}

// GetScheduleSettings returns the ScheduleSettings field value if set, zero value otherwise.
func (o *GrantTypeBundleWriteable) GetScheduleSettings() ScheduleSettingsWriteable {
	if o == nil || IsNil(o.ScheduleSettings) {
		var ret ScheduleSettingsWriteable
		return ret
	}
	return *o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypeBundleWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool) {
	if o == nil || IsNil(o.ScheduleSettings) {
		return nil, false
	}
	return o.ScheduleSettings, true
}

// HasScheduleSettings returns a boolean if a field has been set.
func (o *GrantTypeBundleWriteable) HasScheduleSettings() bool {
	if o != nil && !IsNil(o.ScheduleSettings) {
		return true
	}

	return false
}

// SetScheduleSettings gets a reference to the given ScheduleSettingsWriteable and assigns it to the ScheduleSettings field.
func (o *GrantTypeBundleWriteable) SetScheduleSettings(v ScheduleSettingsWriteable) {
	o.ScheduleSettings = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GrantTypeBundleWriteable) GetAction() GrantAction {
	if o == nil || IsNil(o.Action) {
		var ret GrantAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypeBundleWriteable) GetActionOk() (*GrantAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GrantTypeBundleWriteable) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given GrantAction and assigns it to the Action field.
func (o *GrantTypeBundleWriteable) SetAction(v GrantAction) {
	o.Action = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *GrantTypeBundleWriteable) GetActor() GrantActor {
	if o == nil || IsNil(o.Actor) {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypeBundleWriteable) GetActorOk() (*GrantActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *GrantTypeBundleWriteable) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *GrantTypeBundleWriteable) SetActor(v GrantActor) {
	o.Actor = &v
}

func (o GrantTypeBundleWriteable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantTypeBundleWriteable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grantType"] = o.GrantType
	toSerialize["entitlementBundleId"] = o.EntitlementBundleId
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

func (o *GrantTypeBundleWriteable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grantType",
		"entitlementBundleId",
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

	varGrantTypeBundleWriteable := _GrantTypeBundleWriteable{}

	err = json.Unmarshal(data, &varGrantTypeBundleWriteable)

	if err != nil {
		return err
	}

	*o = GrantTypeBundleWriteable(varGrantTypeBundleWriteable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "entitlementBundleId")
		delete(additionalProperties, "targetPrincipal")
		delete(additionalProperties, "scheduleSettings")
		delete(additionalProperties, "action")
		delete(additionalProperties, "actor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantTypeBundleWriteable struct {
	value *GrantTypeBundleWriteable
	isSet bool
}

func (v NullableGrantTypeBundleWriteable) Get() *GrantTypeBundleWriteable {
	return v.value
}

func (v *NullableGrantTypeBundleWriteable) Set(val *GrantTypeBundleWriteable) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantTypeBundleWriteable) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantTypeBundleWriteable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantTypeBundleWriteable(val *GrantTypeBundleWriteable) *NullableGrantTypeBundleWriteable {
	return &NullableGrantTypeBundleWriteable{value: val, isSet: true}
}

func (v NullableGrantTypeBundleWriteable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantTypeBundleWriteable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
