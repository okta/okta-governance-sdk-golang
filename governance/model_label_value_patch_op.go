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
	"fmt"
)

// LabelValuePatchOp The operation to perform a label value update
type LabelValuePatchOp string

// List of label-value-patch-op
const (
	LABELVALUEPATCHOP_ADD     LabelValuePatchOp = "ADD"
	LABELVALUEPATCHOP_REMOVE  LabelValuePatchOp = "REMOVE"
	LABELVALUEPATCHOP_REPLACE LabelValuePatchOp = "REPLACE"
)

// All allowed values of LabelValuePatchOp enum
var AllowedLabelValuePatchOpEnumValues = []LabelValuePatchOp{
	"ADD",
	"REMOVE",
	"REPLACE",
}

func (v *LabelValuePatchOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LabelValuePatchOp(value)
	for _, existing := range AllowedLabelValuePatchOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LabelValuePatchOp", value)
}

// NewLabelValuePatchOpFromValue returns a pointer to a valid LabelValuePatchOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLabelValuePatchOpFromValue(v string) (*LabelValuePatchOp, error) {
	ev := LabelValuePatchOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LabelValuePatchOp: valid values are %v", v, AllowedLabelValuePatchOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LabelValuePatchOp) IsValid() bool {
	for _, existing := range AllowedLabelValuePatchOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to label-value-patch-op value
func (v LabelValuePatchOp) Ptr() *LabelValuePatchOp {
	return &v
}

type NullableLabelValuePatchOp struct {
	value *LabelValuePatchOp
	isSet bool
}

func (v NullableLabelValuePatchOp) Get() *LabelValuePatchOp {
	return v.value
}

func (v *NullableLabelValuePatchOp) Set(val *LabelValuePatchOp) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelValuePatchOp) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelValuePatchOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelValuePatchOp(val *LabelValuePatchOp) *NullableLabelValuePatchOp {
	return &NullableLabelValuePatchOp{value: val, isSet: true}
}

func (v NullableLabelValuePatchOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelValuePatchOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
