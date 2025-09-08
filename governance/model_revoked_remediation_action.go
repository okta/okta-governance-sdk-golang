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

// RevokedRemediationAction Specifies the action if the reviewer revokes access. `NO_ACTION` indicates the user retains the same access. `DENY` indicates the user will have their access revoked as long as they are not assigned to a group through Group Rules. >**Note:** If a user is assigned to a resource by way of group and the campaign is targeted to users assigned directly to the application, the user won't be revoked and marked with “Manual Intervention required”. If the same campaign is configured to use the Enhanced Group Remediation, and the group the user is assigned to is included in the campaign the user is revoked. If the user is assigned through Group Rules and a campaign attempts to remove the user from the group, the user won't be removed and the review marked as “Manual Intervention required”.
type RevokedRemediationAction string

// List of revoked-remediation-action
const (
	REVOKEDREMEDIATIONACTION_NO_ACTION RevokedRemediationAction = "NO_ACTION"
	REVOKEDREMEDIATIONACTION_DENY      RevokedRemediationAction = "DENY"
)

// All allowed values of RevokedRemediationAction enum
var AllowedRevokedRemediationActionEnumValues = []RevokedRemediationAction{
	"NO_ACTION",
	"DENY",
}

func (v *RevokedRemediationAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RevokedRemediationAction(value)
	for _, existing := range AllowedRevokedRemediationActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RevokedRemediationAction", value)
}

// NewRevokedRemediationActionFromValue returns a pointer to a valid RevokedRemediationAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRevokedRemediationActionFromValue(v string) (*RevokedRemediationAction, error) {
	ev := RevokedRemediationAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RevokedRemediationAction: valid values are %v", v, AllowedRevokedRemediationActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RevokedRemediationAction) IsValid() bool {
	for _, existing := range AllowedRevokedRemediationActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to revoked-remediation-action value
func (v RevokedRemediationAction) Ptr() *RevokedRemediationAction {
	return &v
}

type NullableRevokedRemediationAction struct {
	value *RevokedRemediationAction
	isSet bool
}

func (v NullableRevokedRemediationAction) Get() *RevokedRemediationAction {
	return v.value
}

func (v *NullableRevokedRemediationAction) Set(val *RevokedRemediationAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokedRemediationAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokedRemediationAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokedRemediationAction(val *RevokedRemediationAction) *NullableRevokedRemediationAction {
	return &NullableRevokedRemediationAction{value: val, isSet: true}
}

func (v NullableRevokedRemediationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokedRemediationAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
