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

// ReconciliationDirectionAction The action Okta takes on a drift in one direction.  * `ACCEPT`: Okta honors the change. An `ADD` becomes a new `IMPORT` grant. A `SUB` from a non-birthright source removes the entitlement from the grant that provided it.  * `REJECT`: Okta rejects the change and pushes the original assignment back to the app.  A drift in this direction that matches one of the direction's `conditions` gets this action. A drift that matches none of them gets the inverse: `REJECT` when this is `ACCEPT`, and `ACCEPT` when this is `REJECT`. When `conditions` is empty every drift in the direction matches, so nothing gets the inverse.  On `subtractive`, the inverse is the destructive one. A `REJECT` whose `conditions` name only some entitlements leaves every other removal accepted, which removes those entitlements from the grants that provided them.  This action never lifts birthright protection. However you configure it, a `SUB` on an entitlement from a `POLICY` source is conflict-rejected, and so is a `SUB` from an `ENTITLEMENT-BUNDLE` source unless the import removes every entitlement that bundle granted. See `affectedGrantSource` on a drift for that exception.
type ReconciliationDirectionAction string

// List of reconciliation-direction-action
const (
	RECONCILIATIONDIRECTIONACTION_ACCEPT ReconciliationDirectionAction = "ACCEPT"
	RECONCILIATIONDIRECTIONACTION_REJECT ReconciliationDirectionAction = "REJECT"
)

// All allowed values of ReconciliationDirectionAction enum
var AllowedReconciliationDirectionActionEnumValues = []ReconciliationDirectionAction{
	"ACCEPT",
	"REJECT",
}

func (v *ReconciliationDirectionAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReconciliationDirectionAction(value)
	for _, existing := range AllowedReconciliationDirectionActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReconciliationDirectionAction", value)
}

// NewReconciliationDirectionActionFromValue returns a pointer to a valid ReconciliationDirectionAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReconciliationDirectionActionFromValue(v string) (*ReconciliationDirectionAction, error) {
	ev := ReconciliationDirectionAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReconciliationDirectionAction: valid values are %v", v, AllowedReconciliationDirectionActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReconciliationDirectionAction) IsValid() bool {
	for _, existing := range AllowedReconciliationDirectionActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reconciliation-direction-action value
func (v ReconciliationDirectionAction) Ptr() *ReconciliationDirectionAction {
	return &v
}

type NullableReconciliationDirectionAction struct {
	value *ReconciliationDirectionAction
	isSet bool
}

func (v NullableReconciliationDirectionAction) Get() *ReconciliationDirectionAction {
	return v.value
}

func (v *NullableReconciliationDirectionAction) Set(val *ReconciliationDirectionAction) {
	v.value = val
	v.isSet = true
}

func (v NullableReconciliationDirectionAction) IsSet() bool {
	return v.isSet
}

func (v *NullableReconciliationDirectionAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReconciliationDirectionAction(val *ReconciliationDirectionAction) *NullableReconciliationDirectionAction {
	return &NullableReconciliationDirectionAction{value: val, isSet: true}
}

func (v NullableReconciliationDirectionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReconciliationDirectionAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
