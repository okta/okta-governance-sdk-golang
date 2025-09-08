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

// Decision the model 'Decision'
type Decision string

// List of decision
const (
	DECISION_UNREVIEWED Decision = "UNREVIEWED"
	DECISION_APPROVE    Decision = "APPROVE"
	DECISION_REVOKE     Decision = "REVOKE"
)

// All allowed values of Decision enum
var AllowedDecisionEnumValues = []Decision{
	"UNREVIEWED",
	"APPROVE",
	"REVOKE",
}

func (v *Decision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Decision(value)
	for _, existing := range AllowedDecisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Decision", value)
}

// NewDecisionFromValue returns a pointer to a valid Decision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDecisionFromValue(v string) (*Decision, error) {
	ev := Decision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Decision: valid values are %v", v, AllowedDecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Decision) IsValid() bool {
	for _, existing := range AllowedDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to decision value
func (v Decision) Ptr() *Decision {
	return &v
}

type NullableDecision struct {
	value *Decision
	isSet bool
}

func (v NullableDecision) Get() *Decision {
	return v.value
}

func (v *NullableDecision) Set(val *Decision) {
	v.value = val
	v.isSet = true
}

func (v NullableDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecision(val *Decision) *NullableDecision {
	return &NullableDecision{value: val, isSet: true}
}

func (v NullableDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
