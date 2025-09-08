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

// ApprovedRemediationAction Specifies the action by default if the reviewer approves access. `NO_ACTION` indicates there is no remediation action and the user retains access.
type ApprovedRemediationAction string

// List of approved-remediation-action
const (
	APPROVEDREMEDIATIONACTION_NO_ACTION ApprovedRemediationAction = "NO_ACTION"
)

// All allowed values of ApprovedRemediationAction enum
var AllowedApprovedRemediationActionEnumValues = []ApprovedRemediationAction{
	"NO_ACTION",
}

func (v *ApprovedRemediationAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApprovedRemediationAction(value)
	for _, existing := range AllowedApprovedRemediationActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApprovedRemediationAction", value)
}

// NewApprovedRemediationActionFromValue returns a pointer to a valid ApprovedRemediationAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApprovedRemediationActionFromValue(v string) (*ApprovedRemediationAction, error) {
	ev := ApprovedRemediationAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApprovedRemediationAction: valid values are %v", v, AllowedApprovedRemediationActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApprovedRemediationAction) IsValid() bool {
	for _, existing := range AllowedApprovedRemediationActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to approved-remediation-action value
func (v ApprovedRemediationAction) Ptr() *ApprovedRemediationAction {
	return &v
}

type NullableApprovedRemediationAction struct {
	value *ApprovedRemediationAction
	isSet bool
}

func (v NullableApprovedRemediationAction) Get() *ApprovedRemediationAction {
	return v.value
}

func (v *NullableApprovedRemediationAction) Set(val *ApprovedRemediationAction) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovedRemediationAction) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovedRemediationAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovedRemediationAction(val *ApprovedRemediationAction) *NullableApprovedRemediationAction {
	return &NullableApprovedRemediationAction{value: val, isSet: true}
}

func (v NullableApprovedRemediationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovedRemediationAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
