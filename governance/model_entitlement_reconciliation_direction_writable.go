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

// checks if the EntitlementReconciliationDirectionWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationDirectionWritable{}

// EntitlementReconciliationDirectionWritable The action Okta takes on drifts in one direction, and the conditions that pick out which drifts get it.  `action` and `conditions` are both required on any direction object you send, whatever `mode` is. Separately, the configuration requires both direction objects when `mode` is `ENABLED`.
type EntitlementReconciliationDirectionWritable struct {
	Action ReconciliationDirectionAction `json:"action"`
	// The drifts that a direction's `action` applies to.  Conditions combine with `OR`, so a drift that matches any one of them gets the action. A drift that matches none of them gets the inverse action.  An empty list means the action applies to every drift in the direction.
	Conditions           []EntitlementReconciliationConditionsWritableInner `json:"conditions"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationDirectionWritable EntitlementReconciliationDirectionWritable

// NewEntitlementReconciliationDirectionWritable instantiates a new EntitlementReconciliationDirectionWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationDirectionWritable(action ReconciliationDirectionAction, conditions []EntitlementReconciliationConditionsWritableInner) *EntitlementReconciliationDirectionWritable {
	this := EntitlementReconciliationDirectionWritable{}
	this.Action = action
	this.Conditions = conditions
	return &this
}

// NewEntitlementReconciliationDirectionWritableWithDefaults instantiates a new EntitlementReconciliationDirectionWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationDirectionWritableWithDefaults() *EntitlementReconciliationDirectionWritable {
	this := EntitlementReconciliationDirectionWritable{}
	return &this
}

// GetAction returns the Action field value
func (o *EntitlementReconciliationDirectionWritable) GetAction() ReconciliationDirectionAction {
	if o == nil {
		var ret ReconciliationDirectionAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationDirectionWritable) GetActionOk() (*ReconciliationDirectionAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *EntitlementReconciliationDirectionWritable) SetAction(v ReconciliationDirectionAction) {
	o.Action = v
}

// GetConditions returns the Conditions field value
func (o *EntitlementReconciliationDirectionWritable) GetConditions() []EntitlementReconciliationConditionsWritableInner {
	if o == nil {
		var ret []EntitlementReconciliationConditionsWritableInner
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationDirectionWritable) GetConditionsOk() ([]EntitlementReconciliationConditionsWritableInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *EntitlementReconciliationDirectionWritable) SetConditions(v []EntitlementReconciliationConditionsWritableInner) {
	o.Conditions = v
}

func (o EntitlementReconciliationDirectionWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationDirectionWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["conditions"] = o.Conditions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationDirectionWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"conditions",
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

	varEntitlementReconciliationDirectionWritable := _EntitlementReconciliationDirectionWritable{}

	err = json.Unmarshal(data, &varEntitlementReconciliationDirectionWritable)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationDirectionWritable(varEntitlementReconciliationDirectionWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationDirectionWritable struct {
	value *EntitlementReconciliationDirectionWritable
	isSet bool
}

func (v NullableEntitlementReconciliationDirectionWritable) Get() *EntitlementReconciliationDirectionWritable {
	return v.value
}

func (v *NullableEntitlementReconciliationDirectionWritable) Set(val *EntitlementReconciliationDirectionWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationDirectionWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationDirectionWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationDirectionWritable(val *EntitlementReconciliationDirectionWritable) *NullableEntitlementReconciliationDirectionWritable {
	return &NullableEntitlementReconciliationDirectionWritable{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationDirectionWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationDirectionWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
