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

// LabelPatchOp The operation to perform a label update
type LabelPatchOp string

// List of label-patch-op
const (
	LABELPATCHOP_REPLACE LabelPatchOp = "REPLACE"
)

// All allowed values of LabelPatchOp enum
var AllowedLabelPatchOpEnumValues = []LabelPatchOp{
	"REPLACE",
}

func (v *LabelPatchOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LabelPatchOp(value)
	for _, existing := range AllowedLabelPatchOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LabelPatchOp", value)
}

// NewLabelPatchOpFromValue returns a pointer to a valid LabelPatchOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLabelPatchOpFromValue(v string) (*LabelPatchOp, error) {
	ev := LabelPatchOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LabelPatchOp: valid values are %v", v, AllowedLabelPatchOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LabelPatchOp) IsValid() bool {
	for _, existing := range AllowedLabelPatchOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to label-patch-op value
func (v LabelPatchOp) Ptr() *LabelPatchOp {
	return &v
}

type NullableLabelPatchOp struct {
	value *LabelPatchOp
	isSet bool
}

func (v NullableLabelPatchOp) Get() *LabelPatchOp {
	return v.value
}

func (v *NullableLabelPatchOp) Set(val *LabelPatchOp) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelPatchOp) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelPatchOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelPatchOp(val *LabelPatchOp) *NullableLabelPatchOp {
	return &NullableLabelPatchOp{value: val, isSet: true}
}

func (v NullableLabelPatchOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelPatchOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
