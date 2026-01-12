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

// checks if the RequestConditionCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestConditionCreatable{}

// RequestConditionCreatable A condition defines who can request which resource (with what resource permissions) and what is the approval process if requested.
type RequestConditionCreatable struct {
	RequesterSettings      RequesterSettingsCreatableRequesterSettings     `json:"requesterSettings"`
	AccessScopeSettings    AccessScopeSettingsCreatableAccessScopeSettings `json:"accessScopeSettings"`
	AccessDurationSettings *AccessDurationSettingsCreatable                `json:"accessDurationSettings,omitempty"`
	// The ID of the approval sequence
	ApprovalSequenceId string `json:"approvalSequenceId"`
	// The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority.
	Priority *int32 `json:"priority,omitempty"`
	// Writable unique key on create. Modifiable on update.
	Name string `json:"name"`
	// Human readable description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionCreatable RequestConditionCreatable

// NewRequestConditionCreatable instantiates a new RequestConditionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionCreatable(requesterSettings RequesterSettingsCreatableRequesterSettings, accessScopeSettings AccessScopeSettingsCreatableAccessScopeSettings, approvalSequenceId string, name string) *RequestConditionCreatable {
	this := RequestConditionCreatable{}
	this.Name = name
	return &this
}

// NewRequestConditionCreatableWithDefaults instantiates a new RequestConditionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionCreatableWithDefaults() *RequestConditionCreatable {
	this := RequestConditionCreatable{}
	return &this
}

// GetRequesterSettings returns the RequesterSettings field value
func (o *RequestConditionCreatable) GetRequesterSettings() RequesterSettingsCreatableRequesterSettings {
	if o == nil {
		var ret RequesterSettingsCreatableRequesterSettings
		return ret
	}

	return o.RequesterSettings
}

// GetRequesterSettingsOk returns a tuple with the RequesterSettings field value
// and a boolean to check if the value has been set.
func (o *RequestConditionCreatable) GetRequesterSettingsOk() (*RequesterSettingsCreatableRequesterSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequesterSettings, true
}

// SetRequesterSettings sets field value
func (o *RequestConditionCreatable) SetRequesterSettings(v RequesterSettingsCreatableRequesterSettings) {
	o.RequesterSettings = v
}

// GetAccessScopeSettings returns the AccessScopeSettings field value
func (o *RequestConditionCreatable) GetAccessScopeSettings() AccessScopeSettingsCreatableAccessScopeSettings {
	if o == nil {
		var ret AccessScopeSettingsCreatableAccessScopeSettings
		return ret
	}

	return o.AccessScopeSettings
}

// GetAccessScopeSettingsOk returns a tuple with the AccessScopeSettings field value
// and a boolean to check if the value has been set.
func (o *RequestConditionCreatable) GetAccessScopeSettingsOk() (*AccessScopeSettingsCreatableAccessScopeSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessScopeSettings, true
}

// SetAccessScopeSettings sets field value
func (o *RequestConditionCreatable) SetAccessScopeSettings(v AccessScopeSettingsCreatableAccessScopeSettings) {
	o.AccessScopeSettings = v
}

// GetAccessDurationSettings returns the AccessDurationSettings field value if set, zero value otherwise.
func (o *RequestConditionCreatable) GetAccessDurationSettings() AccessDurationSettingsCreatable {
	if o == nil || IsNil(o.AccessDurationSettings) {
		var ret AccessDurationSettingsCreatable
		return ret
	}
	return *o.AccessDurationSettings
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionCreatable) GetAccessDurationSettingsOk() (*AccessDurationSettingsCreatable, bool) {
	if o == nil || IsNil(o.AccessDurationSettings) {
		return nil, false
	}
	return o.AccessDurationSettings, true
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RequestConditionCreatable) HasAccessDurationSettings() bool {
	if o != nil && !IsNil(o.AccessDurationSettings) {
		return true
	}

	return false
}

// SetAccessDurationSettings gets a reference to the given AccessDurationSettingsCreatable and assigns it to the AccessDurationSettings field.
func (o *RequestConditionCreatable) SetAccessDurationSettings(v AccessDurationSettingsCreatable) {
	o.AccessDurationSettings = &v
}

// GetApprovalSequenceId returns the ApprovalSequenceId field value
func (o *RequestConditionCreatable) GetApprovalSequenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApprovalSequenceId
}

// GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field value
// and a boolean to check if the value has been set.
func (o *RequestConditionCreatable) GetApprovalSequenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalSequenceId, true
}

// SetApprovalSequenceId sets field value
func (o *RequestConditionCreatable) SetApprovalSequenceId(v string) {
	o.ApprovalSequenceId = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RequestConditionCreatable) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionCreatable) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RequestConditionCreatable) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *RequestConditionCreatable) SetPriority(v int32) {
	o.Priority = &v
}

// GetName returns the Name field value
func (o *RequestConditionCreatable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestConditionCreatable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestConditionCreatable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestConditionCreatable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionCreatable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestConditionCreatable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestConditionCreatable) SetDescription(v string) {
	o.Description = &v
}

func (o RequestConditionCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestConditionCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requesterSettings"] = o.RequesterSettings
	toSerialize["accessScopeSettings"] = o.AccessScopeSettings
	if !IsNil(o.AccessDurationSettings) {
		toSerialize["accessDurationSettings"] = o.AccessDurationSettings
	}
	toSerialize["approvalSequenceId"] = o.ApprovalSequenceId
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestConditionCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requesterSettings",
		"accessScopeSettings",
		"approvalSequenceId",
		"name",
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

	varRequestConditionCreatable := _RequestConditionCreatable{}

	err = json.Unmarshal(data, &varRequestConditionCreatable)

	if err != nil {
		return err
	}

	*o = RequestConditionCreatable(varRequestConditionCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requesterSettings")
		delete(additionalProperties, "accessScopeSettings")
		delete(additionalProperties, "accessDurationSettings")
		delete(additionalProperties, "approvalSequenceId")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestConditionCreatable struct {
	value *RequestConditionCreatable
	isSet bool
}

func (v NullableRequestConditionCreatable) Get() *RequestConditionCreatable {
	return v.value
}

func (v *NullableRequestConditionCreatable) Set(val *RequestConditionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionCreatable(val *RequestConditionCreatable) *NullableRequestConditionCreatable {
	return &NullableRequestConditionCreatable{value: val, isSet: true}
}

func (v NullableRequestConditionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
