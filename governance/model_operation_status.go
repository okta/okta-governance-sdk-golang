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

// OperationStatus The status of the operation
type OperationStatus string

// List of operation-status
const (
	OPERATIONSTATUS_RUNNING   OperationStatus = "RUNNING"
	OPERATIONSTATUS_FAILED    OperationStatus = "FAILED"
	OPERATIONSTATUS_COMPLETED OperationStatus = "COMPLETED"
	OPERATIONSTATUS_CANCELED  OperationStatus = "CANCELED"
)

// All allowed values of OperationStatus enum
var AllowedOperationStatusEnumValues = []OperationStatus{
	"RUNNING",
	"FAILED",
	"COMPLETED",
	"CANCELED",
}

func (v *OperationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationStatus(value)
	for _, existing := range AllowedOperationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationStatus", value)
}

// NewOperationStatusFromValue returns a pointer to a valid OperationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationStatusFromValue(v string) (*OperationStatus, error) {
	ev := OperationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationStatus: valid values are %v", v, AllowedOperationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationStatus) IsValid() bool {
	for _, existing := range AllowedOperationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to operation-status value
func (v OperationStatus) Ptr() *OperationStatus {
	return &v
}

type NullableOperationStatus struct {
	value *OperationStatus
	isSet bool
}

func (v NullableOperationStatus) Get() *OperationStatus {
	return v.value
}

func (v *NullableOperationStatus) Set(val *OperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationStatus(val *OperationStatus) *NullableOperationStatus {
	return &NullableOperationStatus{value: val, isSet: true}
}

func (v NullableOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
