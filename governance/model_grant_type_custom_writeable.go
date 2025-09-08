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

// GrantTypeCustomWriteable Custom resource source object
type GrantTypeCustomWriteable struct {
	// Base grant type for assigning specific entitlement and respective value(s).
	GrantType string         `json:"grantType"`
	Target    TargetResource `json:"target"`
	// Collection of entitlements and associated value identifiers
	Entitlements         []EntitlementCreatable     `json:"entitlements,omitempty"`
	TargetPrincipal      TargetPrincipal            `json:"targetPrincipal"`
	ScheduleSettings     *ScheduleSettingsWriteable `json:"scheduleSettings,omitempty"`
	Action               *GrantAction               `json:"action,omitempty"`
	Actor                *GrantActor                `json:"actor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantTypeCustomWriteable GrantTypeCustomWriteable

// NewGrantTypeCustomWriteable instantiates a new GrantTypeCustomWriteable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantTypeCustomWriteable(grantType string, target TargetResource, targetPrincipal TargetPrincipal) *GrantTypeCustomWriteable {
	this := GrantTypeCustomWriteable{}
	this.TargetPrincipal = targetPrincipal
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// NewGrantTypeCustomWriteableWithDefaults instantiates a new GrantTypeCustomWriteable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantTypeCustomWriteableWithDefaults() *GrantTypeCustomWriteable {
	this := GrantTypeCustomWriteable{}
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = &action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetGrantType returns the GrantType field value
func (o *GrantTypeCustomWriteable) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *GrantTypeCustomWriteable) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *GrantTypeCustomWriteable) SetGrantType(v string) {
	o.GrantType = v
}

// GetTarget returns the Target field value
func (o *GrantTypeCustomWriteable) GetTarget() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *GrantTypeCustomWriteable) GetTargetOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *GrantTypeCustomWriteable) SetTarget(v TargetResource) {
	o.Target = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *GrantTypeCustomWriteable) GetEntitlements() []EntitlementCreatable {
	if o == nil || o.Entitlements == nil {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypeCustomWriteable) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *GrantTypeCustomWriteable) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *GrantTypeCustomWriteable) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetTargetPrincipal returns the TargetPrincipal field value
func (o *GrantTypeCustomWriteable) GetTargetPrincipal() TargetPrincipal {
	if o == nil {
		var ret TargetPrincipal
		return ret
	}

	return o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value
// and a boolean to check if the value has been set.
func (o *GrantTypeCustomWriteable) GetTargetPrincipalOk() (*TargetPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipal, true
}

// SetTargetPrincipal sets field value
func (o *GrantTypeCustomWriteable) SetTargetPrincipal(v TargetPrincipal) {
	o.TargetPrincipal = v
}

// GetScheduleSettings returns the ScheduleSettings field value if set, zero value otherwise.
func (o *GrantTypeCustomWriteable) GetScheduleSettings() ScheduleSettingsWriteable {
	if o == nil || o.ScheduleSettings == nil {
		var ret ScheduleSettingsWriteable
		return ret
	}
	return *o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypeCustomWriteable) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool) {
	if o == nil || o.ScheduleSettings == nil {
		return nil, false
	}
	return o.ScheduleSettings, true
}

// HasScheduleSettings returns a boolean if a field has been set.
func (o *GrantTypeCustomWriteable) HasScheduleSettings() bool {
	if o != nil && o.ScheduleSettings != nil {
		return true
	}

	return false
}

// SetScheduleSettings gets a reference to the given ScheduleSettingsWriteable and assigns it to the ScheduleSettings field.
func (o *GrantTypeCustomWriteable) SetScheduleSettings(v ScheduleSettingsWriteable) {
	o.ScheduleSettings = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GrantTypeCustomWriteable) GetAction() GrantAction {
	if o == nil || o.Action == nil {
		var ret GrantAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypeCustomWriteable) GetActionOk() (*GrantAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GrantTypeCustomWriteable) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given GrantAction and assigns it to the Action field.
func (o *GrantTypeCustomWriteable) SetAction(v GrantAction) {
	o.Action = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *GrantTypeCustomWriteable) GetActor() GrantActor {
	if o == nil || o.Actor == nil {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantTypeCustomWriteable) GetActorOk() (*GrantActor, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *GrantTypeCustomWriteable) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *GrantTypeCustomWriteable) SetActor(v GrantActor) {
	o.Actor = &v
}

func (o GrantTypeCustomWriteable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["grantType"] = o.GrantType
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
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

func (o *GrantTypeCustomWriteable) UnmarshalJSON(bytes []byte) (err error) {
	varGrantTypeCustomWriteable := _GrantTypeCustomWriteable{}

	err = json.Unmarshal(bytes, &varGrantTypeCustomWriteable)
	if err == nil {
		*o = GrantTypeCustomWriteable(varGrantTypeCustomWriteable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "target")
		delete(additionalProperties, "entitlements")
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

type NullableGrantTypeCustomWriteable struct {
	value *GrantTypeCustomWriteable
	isSet bool
}

func (v NullableGrantTypeCustomWriteable) Get() *GrantTypeCustomWriteable {
	return v.value
}

func (v *NullableGrantTypeCustomWriteable) Set(val *GrantTypeCustomWriteable) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantTypeCustomWriteable) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantTypeCustomWriteable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantTypeCustomWriteable(val *GrantTypeCustomWriteable) *NullableGrantTypeCustomWriteable {
	return &NullableGrantTypeCustomWriteable{value: val, isSet: true}
}

func (v NullableGrantTypeCustomWriteable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantTypeCustomWriteable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
