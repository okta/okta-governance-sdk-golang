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

// TaskStatus The status of the task
type TaskStatus string

// List of task-status
const (
	TASKSTATUS_OPEN      TaskStatus = "OPEN"
	TASKSTATUS_COMPLETED TaskStatus = "COMPLETED"
	TASKSTATUS_ASSIGNING TaskStatus = "ASSIGNING"
	TASKSTATUS_LOCKED    TaskStatus = "LOCKED"
)

// All allowed values of TaskStatus enum
var AllowedTaskStatusEnumValues = []TaskStatus{
	"OPEN",
	"COMPLETED",
	"ASSIGNING",
	"LOCKED",
}

func (v *TaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskStatus(value)
	for _, existing := range AllowedTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskStatus", value)
}

// NewTaskStatusFromValue returns a pointer to a valid TaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskStatusFromValue(v string) (*TaskStatus, error) {
	ev := TaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskStatus: valid values are %v", v, AllowedTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskStatus) IsValid() bool {
	for _, existing := range AllowedTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to task-status value
func (v TaskStatus) Ptr() *TaskStatus {
	return &v
}

type NullableTaskStatus struct {
	value *TaskStatus
	isSet bool
}

func (v NullableTaskStatus) Get() *TaskStatus {
	return v.value
}

func (v *NullableTaskStatus) Set(val *TaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStatus(val *TaskStatus) *NullableTaskStatus {
	return &NullableTaskStatus{value: val, isSet: true}
}

func (v NullableTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
