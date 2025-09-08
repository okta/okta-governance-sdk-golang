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

// RequestConditionPatchable A condition defines who can request which resource (with what resource permissions) and what is the approval process if requested.
type RequestConditionPatchable struct {
	RequesterSettings      *RequesterSettingsCreatableRequesterSettings     `json:"requesterSettings,omitempty"`
	AccessScopeSettings    *AccessScopeSettingsCreatableAccessScopeSettings `json:"accessScopeSettings,omitempty"`
	AccessDurationSettings NullableAccessDurationSettingsPatchable          `json:"accessDurationSettings,omitempty"`
	// The ID of the approval sequence
	ApprovalSequenceId *string `json:"approvalSequenceId,omitempty"`
	// The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority.
	Priority *int32 `json:"priority,omitempty"`
	// Writable unique key on Create. Modifiable on update.
	Name *string `json:"name,omitempty"`
	// Human readable description.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionPatchable RequestConditionPatchable

// NewRequestConditionPatchable instantiates a new RequestConditionPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionPatchable() *RequestConditionPatchable {
	this := RequestConditionPatchable{}
	return &this
}

// NewRequestConditionPatchableWithDefaults instantiates a new RequestConditionPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionPatchableWithDefaults() *RequestConditionPatchable {
	this := RequestConditionPatchable{}
	return &this
}

// GetRequesterSettings returns the RequesterSettings field value if set, zero value otherwise.
func (o *RequestConditionPatchable) GetRequesterSettings() RequesterSettingsCreatableRequesterSettings {
	if o == nil || o.RequesterSettings == nil {
		var ret RequesterSettingsCreatableRequesterSettings
		return ret
	}
	return *o.RequesterSettings
}

// GetRequesterSettingsOk returns a tuple with the RequesterSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionPatchable) GetRequesterSettingsOk() (*RequesterSettingsCreatableRequesterSettings, bool) {
	if o == nil || o.RequesterSettings == nil {
		return nil, false
	}
	return o.RequesterSettings, true
}

// HasRequesterSettings returns a boolean if a field has been set.
func (o *RequestConditionPatchable) HasRequesterSettings() bool {
	if o != nil && o.RequesterSettings != nil {
		return true
	}

	return false
}

// SetRequesterSettings gets a reference to the given RequesterSettingsCreatableRequesterSettings and assigns it to the RequesterSettings field.
func (o *RequestConditionPatchable) SetRequesterSettings(v RequesterSettingsCreatableRequesterSettings) {
	o.RequesterSettings = &v
}

// GetAccessScopeSettings returns the AccessScopeSettings field value if set, zero value otherwise.
func (o *RequestConditionPatchable) GetAccessScopeSettings() AccessScopeSettingsCreatableAccessScopeSettings {
	if o == nil || o.AccessScopeSettings == nil {
		var ret AccessScopeSettingsCreatableAccessScopeSettings
		return ret
	}
	return *o.AccessScopeSettings
}

// GetAccessScopeSettingsOk returns a tuple with the AccessScopeSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionPatchable) GetAccessScopeSettingsOk() (*AccessScopeSettingsCreatableAccessScopeSettings, bool) {
	if o == nil || o.AccessScopeSettings == nil {
		return nil, false
	}
	return o.AccessScopeSettings, true
}

// HasAccessScopeSettings returns a boolean if a field has been set.
func (o *RequestConditionPatchable) HasAccessScopeSettings() bool {
	if o != nil && o.AccessScopeSettings != nil {
		return true
	}

	return false
}

// SetAccessScopeSettings gets a reference to the given AccessScopeSettingsCreatableAccessScopeSettings and assigns it to the AccessScopeSettings field.
func (o *RequestConditionPatchable) SetAccessScopeSettings(v AccessScopeSettingsCreatableAccessScopeSettings) {
	o.AccessScopeSettings = &v
}

// GetAccessDurationSettings returns the AccessDurationSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestConditionPatchable) GetAccessDurationSettings() AccessDurationSettingsPatchable {
	if o == nil || o.AccessDurationSettings.Get() == nil {
		var ret AccessDurationSettingsPatchable
		return ret
	}
	return *o.AccessDurationSettings.Get()
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestConditionPatchable) GetAccessDurationSettingsOk() (*AccessDurationSettingsPatchable, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessDurationSettings.Get(), o.AccessDurationSettings.IsSet()
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RequestConditionPatchable) HasAccessDurationSettings() bool {
	if o != nil && o.AccessDurationSettings.IsSet() {
		return true
	}

	return false
}

// SetAccessDurationSettings gets a reference to the given NullableAccessDurationSettingsPatchable and assigns it to the AccessDurationSettings field.
func (o *RequestConditionPatchable) SetAccessDurationSettings(v AccessDurationSettingsPatchable) {
	o.AccessDurationSettings.Set(&v)
}

// SetAccessDurationSettingsNil sets the value for AccessDurationSettings to be an explicit nil
func (o *RequestConditionPatchable) SetAccessDurationSettingsNil() {
	o.AccessDurationSettings.Set(nil)
}

// UnsetAccessDurationSettings ensures that no value is present for AccessDurationSettings, not even an explicit nil
func (o *RequestConditionPatchable) UnsetAccessDurationSettings() {
	o.AccessDurationSettings.Unset()
}

// GetApprovalSequenceId returns the ApprovalSequenceId field value if set, zero value otherwise.
func (o *RequestConditionPatchable) GetApprovalSequenceId() string {
	if o == nil || o.ApprovalSequenceId == nil {
		var ret string
		return ret
	}
	return *o.ApprovalSequenceId
}

// GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionPatchable) GetApprovalSequenceIdOk() (*string, bool) {
	if o == nil || o.ApprovalSequenceId == nil {
		return nil, false
	}
	return o.ApprovalSequenceId, true
}

// HasApprovalSequenceId returns a boolean if a field has been set.
func (o *RequestConditionPatchable) HasApprovalSequenceId() bool {
	if o != nil && o.ApprovalSequenceId != nil {
		return true
	}

	return false
}

// SetApprovalSequenceId gets a reference to the given string and assigns it to the ApprovalSequenceId field.
func (o *RequestConditionPatchable) SetApprovalSequenceId(v string) {
	o.ApprovalSequenceId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RequestConditionPatchable) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionPatchable) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RequestConditionPatchable) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *RequestConditionPatchable) SetPriority(v int32) {
	o.Priority = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestConditionPatchable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionPatchable) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestConditionPatchable) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestConditionPatchable) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestConditionPatchable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionPatchable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestConditionPatchable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestConditionPatchable) SetDescription(v string) {
	o.Description = &v
}

func (o RequestConditionPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequesterSettings != nil {
		toSerialize["requesterSettings"] = o.RequesterSettings
	}
	if o.AccessScopeSettings != nil {
		toSerialize["accessScopeSettings"] = o.AccessScopeSettings
	}
	if o.AccessDurationSettings.IsSet() {
		toSerialize["accessDurationSettings"] = o.AccessDurationSettings.Get()
	}
	if o.ApprovalSequenceId != nil {
		toSerialize["approvalSequenceId"] = o.ApprovalSequenceId
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestConditionPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestConditionPatchable := _RequestConditionPatchable{}

	err = json.Unmarshal(bytes, &varRequestConditionPatchable)
	if err == nil {
		*o = RequestConditionPatchable(varRequestConditionPatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requesterSettings")
		delete(additionalProperties, "accessScopeSettings")
		delete(additionalProperties, "accessDurationSettings")
		delete(additionalProperties, "approvalSequenceId")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestConditionPatchable struct {
	value *RequestConditionPatchable
	isSet bool
}

func (v NullableRequestConditionPatchable) Get() *RequestConditionPatchable {
	return v.value
}

func (v *NullableRequestConditionPatchable) Set(val *RequestConditionPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionPatchable(val *RequestConditionPatchable) *NullableRequestConditionPatchable {
	return &NullableRequestConditionPatchable{value: val, isSet: true}
}

func (v NullableRequestConditionPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
