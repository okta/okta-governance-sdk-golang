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

// NoResponseRemediationAction Specifies the action if the reviewer doesn't respond to the request or if the campaign is closed before an action is taken.` NO_ACTION` indicates there is no remediation action taken and the user retains the same access.` DENY` indicates the user will have their access revoked when the campaign ends as long as they are not assigned to a group through Group Rules. >**Note:** If a user is assigned to a resource by way of group and the campaign is targeted to users assigned directly to the application, the user won't be revoked and marked with “Manual Intervention required”. If the same campaign is configured to use the Enhanced Group Remediation, and the group the user is assigned to is included in the campaign the user is revoked. If the user is assigned through Group Rules and a campaign attempts to remove the user from the group, the user won't be removed and the review marked as “Manual Intervention required”.
type NoResponseRemediationAction string

// List of no-response-remediation-action
const (
	NORESPONSEREMEDIATIONACTION_NO_ACTION NoResponseRemediationAction = "NO_ACTION"
	NORESPONSEREMEDIATIONACTION_DENY      NoResponseRemediationAction = "DENY"
)

// All allowed values of NoResponseRemediationAction enum
var AllowedNoResponseRemediationActionEnumValues = []NoResponseRemediationAction{
	"NO_ACTION",
	"DENY",
}

func (v *NoResponseRemediationAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NoResponseRemediationAction(value)
	for _, existing := range AllowedNoResponseRemediationActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NoResponseRemediationAction", value)
}

// NewNoResponseRemediationActionFromValue returns a pointer to a valid NoResponseRemediationAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNoResponseRemediationActionFromValue(v string) (*NoResponseRemediationAction, error) {
	ev := NoResponseRemediationAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NoResponseRemediationAction: valid values are %v", v, AllowedNoResponseRemediationActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NoResponseRemediationAction) IsValid() bool {
	for _, existing := range AllowedNoResponseRemediationActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to no-response-remediation-action value
func (v NoResponseRemediationAction) Ptr() *NoResponseRemediationAction {
	return &v
}

type NullableNoResponseRemediationAction struct {
	value *NoResponseRemediationAction
	isSet bool
}

func (v NullableNoResponseRemediationAction) Get() *NoResponseRemediationAction {
	return v.value
}

func (v *NullableNoResponseRemediationAction) Set(val *NoResponseRemediationAction) {
	v.value = val
	v.isSet = true
}

func (v NullableNoResponseRemediationAction) IsSet() bool {
	return v.isSet
}

func (v *NullableNoResponseRemediationAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoResponseRemediationAction(val *NoResponseRemediationAction) *NullableNoResponseRemediationAction {
	return &NullableNoResponseRemediationAction{value: val, isSet: true}
}

func (v NullableNoResponseRemediationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoResponseRemediationAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
