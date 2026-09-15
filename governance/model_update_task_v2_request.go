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

// checks if the UpdateTaskV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTaskV2Request{}

// UpdateTaskV2Request struct for UpdateTaskV2Request
type UpdateTaskV2Request struct {
	// List of task assignees to perform actions on the task. A maximum of 10 assignees can be assigned to a task.
	Assignees            []TaskAssignees `json:"assignees,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateTaskV2Request UpdateTaskV2Request

// NewUpdateTaskV2Request instantiates a new UpdateTaskV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTaskV2Request() *UpdateTaskV2Request {
	this := UpdateTaskV2Request{}
	return &this
}

// NewUpdateTaskV2RequestWithDefaults instantiates a new UpdateTaskV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTaskV2RequestWithDefaults() *UpdateTaskV2Request {
	this := UpdateTaskV2Request{}
	return &this
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *UpdateTaskV2Request) GetAssignees() []TaskAssignees {
	if o == nil || IsNil(o.Assignees) {
		var ret []TaskAssignees
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskV2Request) GetAssigneesOk() ([]TaskAssignees, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *UpdateTaskV2Request) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []TaskAssignees and assigns it to the Assignees field.
func (o *UpdateTaskV2Request) SetAssignees(v []TaskAssignees) {
	o.Assignees = v
}

func (o UpdateTaskV2Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTaskV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignees) {
		toSerialize["assignees"] = o.Assignees
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateTaskV2Request) UnmarshalJSON(data []byte) (err error) {
	varUpdateTaskV2Request := _UpdateTaskV2Request{}

	err = json.Unmarshal(data, &varUpdateTaskV2Request)

	if err != nil {
		return err
	}

	*o = UpdateTaskV2Request(varUpdateTaskV2Request)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignees")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateTaskV2Request struct {
	value *UpdateTaskV2Request
	isSet bool
}

func (v NullableUpdateTaskV2Request) Get() *UpdateTaskV2Request {
	return v.value
}

func (v *NullableUpdateTaskV2Request) Set(val *UpdateTaskV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTaskV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTaskV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTaskV2Request(val *UpdateTaskV2Request) *NullableUpdateTaskV2Request {
	return &NullableUpdateTaskV2Request{value: val, isSet: true}
}

func (v NullableUpdateTaskV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTaskV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
