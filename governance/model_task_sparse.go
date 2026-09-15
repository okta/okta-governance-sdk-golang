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

// checks if the TaskSparse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskSparse{}

// TaskSparse Sparse representation of a Task resource
type TaskSparse struct {
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
	Value                NullableString `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskSparse TaskSparse

// NewTaskSparse instantiates a new TaskSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskSparse(id string, assignees []TaskAssignees, status TaskStatus, label string, requestId string, createdAt time.Time, updatedAt time.Time, type_ TaskType) *TaskSparse {
	this := TaskSparse{}
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

// NewTaskSparseWithDefaults instantiates a new TaskSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskSparseWithDefaults() *TaskSparse {
	this := TaskSparse{}
	return &this
}

// GetId returns the Id field value
func (o *TaskSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskSparse) SetId(v string) {
	o.Id = v
}

// GetAssignees returns the Assignees field value
func (o *TaskSparse) GetAssignees() []TaskAssignees {
	if o == nil {
		var ret []TaskAssignees
		return ret
	}

	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetAssigneesOk() ([]TaskAssignees, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignees, true
}

// SetAssignees sets field value
func (o *TaskSparse) SetAssignees(v []TaskAssignees) {
	o.Assignees = v
}

// GetStatus returns the Status field value
func (o *TaskSparse) GetStatus() TaskStatus {
	if o == nil {
		var ret TaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetStatusOk() (*TaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TaskSparse) SetStatus(v TaskStatus) {
	o.Status = v
}

// GetLabel returns the Label field value
func (o *TaskSparse) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TaskSparse) SetLabel(v string) {
	o.Label = v
}

// GetRequestId returns the RequestId field value
func (o *TaskSparse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TaskSparse) SetRequestId(v string) {
	o.RequestId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaskSparse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaskSparse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TaskSparse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TaskSparse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetType returns the Type field value
func (o *TaskSparse) GetType() TaskType {
	if o == nil {
		var ret TaskType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaskSparse) GetTypeOk() (*TaskType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaskSparse) SetType(v TaskType) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSparse) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSparse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *TaskSparse) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *TaskSparse) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *TaskSparse) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *TaskSparse) UnsetValue() {
	o.Value.Unset()
}

func (o TaskSparse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskSparse) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskSparse) UnmarshalJSON(data []byte) (err error) {
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

	varTaskSparse := _TaskSparse{}

	err = json.Unmarshal(data, &varTaskSparse)

	if err != nil {
		return err
	}

	*o = TaskSparse(varTaskSparse)

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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskSparse struct {
	value *TaskSparse
	isSet bool
}

func (v NullableTaskSparse) Get() *TaskSparse {
	return v.value
}

func (v *NullableTaskSparse) Set(val *TaskSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskSparse(val *TaskSparse) *NullableTaskSparse {
	return &NullableTaskSparse{value: val, isSet: true}
}

func (v NullableTaskSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
