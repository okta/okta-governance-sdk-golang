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
	"time"
)

// checks if the TaskFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskFull{}

// TaskFull Full representation of a task.
type TaskFull struct {
	// The unique identifier for the task
	Id string `json:"id" validate:"regexp=^[0-9a-f]{24}$"`
	// List of task assignees that perform actions on the task
	Assignees []TaskAssignees `json:"assignees"`
	Status    TaskStatus      `json:"status"`
	// Human readable label for the task
	Label string `json:"label"`
	// The request ID associated with the task
	RequestId string `json:"requestId"`
	// The ISO 8601 formatted date and time when the object was created
	CreatedAt time.Time `json:"createdAt"`
	// The ISO 8601 formatted date and time when the object was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	Type      TaskType  `json:"type"`
	// The value of the task completion that's determined by the assignee
	Value NullableString `json:"value,omitempty"`
	// Indicates whether a task is escalated to another user. See [Escalate Tasks](https://help.okta.com/okta_help.htm?type=oie&id=csh-escl-task) in the product documentation.
	IsEscalated *bool `json:"isEscalated,omitempty"`
	// Indicates whether a task is delegated to another user. See [Governance delegates](https://help.okta.com/okta_help.htm?type=oie&id=csh-governance-delegates) in the product documentation.
	IsDelegated *bool `json:"isDelegated,omitempty"`
	// ID of the original assignee before delegation or escalation
	OriginalAssigneeId   NullableString   `json:"originalAssigneeId,omitempty"`
	CompletedBy          *TaskCompletedBy `json:"completedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskFull TaskFull

// NewTaskFull instantiates a new TaskFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskFull(id string, assignees []TaskAssignees, status TaskStatus, label string, requestId string, createdAt time.Time, updatedAt time.Time, type_ TaskType) *TaskFull {
	this := TaskFull{}
	this.Id = id
	this.Assignees = assignees
	this.Status = status
	this.Label = label
	this.RequestId = requestId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Type = type_
	return &this
}

// NewTaskFullWithDefaults instantiates a new TaskFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskFullWithDefaults() *TaskFull {
	this := TaskFull{}
	return &this
}

// GetId returns the Id field value
func (o *TaskFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskFull) SetId(v string) {
	o.Id = v
}

// GetAssignees returns the Assignees field value
func (o *TaskFull) GetAssignees() []TaskAssignees {
	if o == nil {
		var ret []TaskAssignees
		return ret
	}

	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetAssigneesOk() ([]TaskAssignees, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignees, true
}

// SetAssignees sets field value
func (o *TaskFull) SetAssignees(v []TaskAssignees) {
	o.Assignees = v
}

// GetStatus returns the Status field value
func (o *TaskFull) GetStatus() TaskStatus {
	if o == nil {
		var ret TaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetStatusOk() (*TaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TaskFull) SetStatus(v TaskStatus) {
	o.Status = v
}

// GetLabel returns the Label field value
func (o *TaskFull) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TaskFull) SetLabel(v string) {
	o.Label = v
}

// GetRequestId returns the RequestId field value
func (o *TaskFull) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TaskFull) SetRequestId(v string) {
	o.RequestId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaskFull) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaskFull) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TaskFull) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TaskFull) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetType returns the Type field value
func (o *TaskFull) GetType() TaskType {
	if o == nil {
		var ret TaskType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaskFull) GetTypeOk() (*TaskType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaskFull) SetType(v TaskType) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskFull) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskFull) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *TaskFull) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *TaskFull) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *TaskFull) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *TaskFull) UnsetValue() {
	o.Value.Unset()
}

// GetIsEscalated returns the IsEscalated field value if set, zero value otherwise.
func (o *TaskFull) GetIsEscalated() bool {
	if o == nil || IsNil(o.IsEscalated) {
		var ret bool
		return ret
	}
	return *o.IsEscalated
}

// GetIsEscalatedOk returns a tuple with the IsEscalated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFull) GetIsEscalatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEscalated) {
		return nil, false
	}
	return o.IsEscalated, true
}

// HasIsEscalated returns a boolean if a field has been set.
func (o *TaskFull) HasIsEscalated() bool {
	if o != nil && !IsNil(o.IsEscalated) {
		return true
	}

	return false
}

// SetIsEscalated gets a reference to the given bool and assigns it to the IsEscalated field.
func (o *TaskFull) SetIsEscalated(v bool) {
	o.IsEscalated = &v
}

// GetIsDelegated returns the IsDelegated field value if set, zero value otherwise.
func (o *TaskFull) GetIsDelegated() bool {
	if o == nil || IsNil(o.IsDelegated) {
		var ret bool
		return ret
	}
	return *o.IsDelegated
}

// GetIsDelegatedOk returns a tuple with the IsDelegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFull) GetIsDelegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDelegated) {
		return nil, false
	}
	return o.IsDelegated, true
}

// HasIsDelegated returns a boolean if a field has been set.
func (o *TaskFull) HasIsDelegated() bool {
	if o != nil && !IsNil(o.IsDelegated) {
		return true
	}

	return false
}

// SetIsDelegated gets a reference to the given bool and assigns it to the IsDelegated field.
func (o *TaskFull) SetIsDelegated(v bool) {
	o.IsDelegated = &v
}

// GetOriginalAssigneeId returns the OriginalAssigneeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskFull) GetOriginalAssigneeId() string {
	if o == nil || IsNil(o.OriginalAssigneeId.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalAssigneeId.Get()
}

// GetOriginalAssigneeIdOk returns a tuple with the OriginalAssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskFull) GetOriginalAssigneeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalAssigneeId.Get(), o.OriginalAssigneeId.IsSet()
}

// HasOriginalAssigneeId returns a boolean if a field has been set.
func (o *TaskFull) HasOriginalAssigneeId() bool {
	if o != nil && o.OriginalAssigneeId.IsSet() {
		return true
	}

	return false
}

// SetOriginalAssigneeId gets a reference to the given NullableString and assigns it to the OriginalAssigneeId field.
func (o *TaskFull) SetOriginalAssigneeId(v string) {
	o.OriginalAssigneeId.Set(&v)
}

// SetOriginalAssigneeIdNil sets the value for OriginalAssigneeId to be an explicit nil
func (o *TaskFull) SetOriginalAssigneeIdNil() {
	o.OriginalAssigneeId.Set(nil)
}

// UnsetOriginalAssigneeId ensures that no value is present for OriginalAssigneeId, not even an explicit nil
func (o *TaskFull) UnsetOriginalAssigneeId() {
	o.OriginalAssigneeId.Unset()
}

// GetCompletedBy returns the CompletedBy field value if set, zero value otherwise.
func (o *TaskFull) GetCompletedBy() TaskCompletedBy {
	if o == nil || IsNil(o.CompletedBy) {
		var ret TaskCompletedBy
		return ret
	}
	return *o.CompletedBy
}

// GetCompletedByOk returns a tuple with the CompletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskFull) GetCompletedByOk() (*TaskCompletedBy, bool) {
	if o == nil || IsNil(o.CompletedBy) {
		return nil, false
	}
	return o.CompletedBy, true
}

// HasCompletedBy returns a boolean if a field has been set.
func (o *TaskFull) HasCompletedBy() bool {
	if o != nil && !IsNil(o.CompletedBy) {
		return true
	}

	return false
}

// SetCompletedBy gets a reference to the given TaskCompletedBy and assigns it to the CompletedBy field.
func (o *TaskFull) SetCompletedBy(v TaskCompletedBy) {
	o.CompletedBy = &v
}

func (o TaskFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["assignees"] = o.Assignees
	toSerialize["status"] = o.Status
	toSerialize["label"] = o.Label
	toSerialize["requestId"] = o.RequestId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["type"] = o.Type
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if !IsNil(o.IsEscalated) {
		toSerialize["isEscalated"] = o.IsEscalated
	}
	if !IsNil(o.IsDelegated) {
		toSerialize["isDelegated"] = o.IsDelegated
	}
	if o.OriginalAssigneeId.IsSet() {
		toSerialize["originalAssigneeId"] = o.OriginalAssigneeId.Get()
	}
	if !IsNil(o.CompletedBy) {
		toSerialize["completedBy"] = o.CompletedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"assignees",
		"status",
		"label",
		"requestId",
		"createdAt",
		"updatedAt",
		"type",
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

	varTaskFull := _TaskFull{}

	err = json.Unmarshal(data, &varTaskFull)

	if err != nil {
		return err
	}

	*o = TaskFull(varTaskFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "assignees")
		delete(additionalProperties, "status")
		delete(additionalProperties, "label")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "isEscalated")
		delete(additionalProperties, "isDelegated")
		delete(additionalProperties, "originalAssigneeId")
		delete(additionalProperties, "completedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskFull struct {
	value *TaskFull
	isSet bool
}

func (v NullableTaskFull) Get() *TaskFull {
	return v.value
}

func (v *NullableTaskFull) Set(val *TaskFull) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskFull) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskFull(val *TaskFull) *NullableTaskFull {
	return &NullableTaskFull{value: val, isSet: true}
}

func (v NullableTaskFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
