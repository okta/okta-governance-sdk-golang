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

// TaskType The type of the task
type TaskType string

// List of task-type
const (
	TASKTYPE_QUESTION TaskType = "QUESTION"
	TASKTYPE_APPROVAL TaskType = "APPROVAL"
	TASKTYPE_TODO     TaskType = "TODO"
)

// All allowed values of TaskType enum
var AllowedTaskTypeEnumValues = []TaskType{
	"QUESTION",
	"APPROVAL",
	"TODO",
}

func (v *TaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskType(value)
	for _, existing := range AllowedTaskTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskType", value)
}

// NewTaskTypeFromValue returns a pointer to a valid TaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskTypeFromValue(v string) (*TaskType, error) {
	ev := TaskType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskType: valid values are %v", v, AllowedTaskTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskType) IsValid() bool {
	for _, existing := range AllowedTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to task-type value
func (v TaskType) Ptr() *TaskType {
	return &v
}

type NullableTaskType struct {
	value *TaskType
	isSet bool
}

func (v NullableTaskType) Get() *TaskType {
	return v.value
}

func (v *NullableTaskType) Set(val *TaskType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskType(val *TaskType) *NullableTaskType {
	return &NullableTaskType{value: val, isSet: true}
}

func (v NullableTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
