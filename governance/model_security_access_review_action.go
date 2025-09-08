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

// SecurityAccessReviewAction struct for SecurityAccessReviewAction
type SecurityAccessReviewAction struct {
	ActionType           SecurityAccessReviewActionType `json:"actionType"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewAction SecurityAccessReviewAction

// NewSecurityAccessReviewAction instantiates a new SecurityAccessReviewAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewAction(actionType SecurityAccessReviewActionType) *SecurityAccessReviewAction {
	this := SecurityAccessReviewAction{}
	this.ActionType = actionType
	return &this
}

// NewSecurityAccessReviewActionWithDefaults instantiates a new SecurityAccessReviewAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewActionWithDefaults() *SecurityAccessReviewAction {
	this := SecurityAccessReviewAction{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *SecurityAccessReviewAction) GetActionType() SecurityAccessReviewActionType {
	if o == nil {
		var ret SecurityAccessReviewActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAction) GetActionTypeOk() (*SecurityAccessReviewActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *SecurityAccessReviewAction) SetActionType(v SecurityAccessReviewActionType) {
	o.ActionType = v
}

func (o SecurityAccessReviewAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["actionType"] = o.ActionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewAction) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewAction := _SecurityAccessReviewAction{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewAction)
	if err == nil {
		*o = SecurityAccessReviewAction(varSecurityAccessReviewAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actionType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewAction struct {
	value *SecurityAccessReviewAction
	isSet bool
}

func (v NullableSecurityAccessReviewAction) Get() *SecurityAccessReviewAction {
	return v.value
}

func (v *NullableSecurityAccessReviewAction) Set(val *SecurityAccessReviewAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAction(val *SecurityAccessReviewAction) *NullableSecurityAccessReviewAction {
	return &NullableSecurityAccessReviewAction{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
