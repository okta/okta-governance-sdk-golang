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
)

// checks if the TaskList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskList{}

// TaskList struct for TaskList
type TaskList struct {
	// All tasks on the current page
	Data                 []TaskSparse   `json:"data,omitempty"`
	Links                *TaskListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskList TaskList

// NewTaskList instantiates a new TaskList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskList() *TaskList {
	this := TaskList{}
	return &this
}

// NewTaskListWithDefaults instantiates a new TaskList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskListWithDefaults() *TaskList {
	this := TaskList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TaskList) GetData() []TaskSparse {
	if o == nil || IsNil(o.Data) {
		var ret []TaskSparse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskList) GetDataOk() ([]TaskSparse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TaskList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TaskSparse and assigns it to the Data field.
func (o *TaskList) SetData(v []TaskSparse) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TaskList) GetLinks() TaskListLinks {
	if o == nil || IsNil(o.Links) {
		var ret TaskListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskList) GetLinksOk() (*TaskListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TaskList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TaskListLinks and assigns it to the Links field.
func (o *TaskList) SetLinks(v TaskListLinks) {
	o.Links = &v
}

func (o TaskList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaskList) UnmarshalJSON(data []byte) (err error) {
	varTaskList := _TaskList{}

	err = json.Unmarshal(data, &varTaskList)

	if err != nil {
		return err
	}

	*o = TaskList(varTaskList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaskList struct {
	value *TaskList
	isSet bool
}

func (v NullableTaskList) Get() *TaskList {
	return v.value
}

func (v *NullableTaskList) Set(val *TaskList) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskList) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskList(val *TaskList) *NullableTaskList {
	return &NullableTaskList{value: val, isSet: true}
}

func (v NullableTaskList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
