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

// checks if the RequestActionPending type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestActionPending{}

// RequestActionPending A pending action, approvals still required
type RequestActionPending struct {
	// The status of a pending action
	Status string            `json:"status"`
	Action RequestActionEnum `json:"action"`
	// A unique identifier of the action taken in the request
	ActionId             string `json:"actionId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	AdditionalProperties map[string]interface{}
}

type _RequestActionPending RequestActionPending

// NewRequestActionPending instantiates a new RequestActionPending object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestActionPending(status string, action RequestActionEnum, actionId string) *RequestActionPending {
	this := RequestActionPending{}
	this.Status = status
	this.Action = action
	this.ActionId = actionId
	return &this
}

// NewRequestActionPendingWithDefaults instantiates a new RequestActionPending object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestActionPendingWithDefaults() *RequestActionPending {
	this := RequestActionPending{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestActionPending) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestActionPending) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestActionPending) SetStatus(v string) {
	o.Status = v
}

// GetAction returns the Action field value
func (o *RequestActionPending) GetAction() RequestActionEnum {
	if o == nil {
		var ret RequestActionEnum
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RequestActionPending) GetActionOk() (*RequestActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RequestActionPending) SetAction(v RequestActionEnum) {
	o.Action = v
}

// GetActionId returns the ActionId field value
func (o *RequestActionPending) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *RequestActionPending) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *RequestActionPending) SetActionId(v string) {
	o.ActionId = v
}

func (o RequestActionPending) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestActionPending) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["action"] = o.Action
	toSerialize["actionId"] = o.ActionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestActionPending) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"action",
		"actionId",
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

	varRequestActionPending := _RequestActionPending{}

	err = json.Unmarshal(data, &varRequestActionPending)

	if err != nil {
		return err
	}

	*o = RequestActionPending(varRequestActionPending)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "action")
		delete(additionalProperties, "actionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestActionPending struct {
	value *RequestActionPending
	isSet bool
}

func (v NullableRequestActionPending) Get() *RequestActionPending {
	return v.value
}

func (v *NullableRequestActionPending) Set(val *RequestActionPending) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestActionPending) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestActionPending) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestActionPending(val *RequestActionPending) *NullableRequestActionPending {
	return &NullableRequestActionPending{value: val, isSet: true}
}

func (v NullableRequestActionPending) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestActionPending) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
