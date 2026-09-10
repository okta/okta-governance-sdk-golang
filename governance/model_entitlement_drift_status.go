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

// EntitlementDriftStatus Where the drift is in its lifecycle.  Okta has decided how to resolve the drift and is still applying it:   * `AUTO_ACCEPTED`: Okta accepted the change and is applying it.   * `AUTO_REJECTED`: Okta rejected the change and is restoring the original assignment in the app.   * `CONFLICT_REJECTED`: The change tried to remove a birthright entitlement (`POLICY` or `ENTITLEMENT-BUNDLE`), so Okta is restoring the original assignment.  Finished:   * `APPLIED`: The drift is resolved and any change to the app is complete.   * `REVERTED`: An admin undid an `APPLIED` + `AUTO_ACCEPTED` drift. You can't revert rejected or conflict-rejected drifts.   * `SUPERSEDED`: A later reconciliation cycle for the same resource, principal, and entitlement replaced this drift, so it no longer reflects the current state.  Failed:   * `ERROR`: Okta couldn't finish resolving the drift. This isn't final. The next reconciliation cycle for the same resource, principal, and entitlement supersedes it.
type EntitlementDriftStatus string

// List of entitlement-drift-status
const (
	ENTITLEMENTDRIFTSTATUS_AUTO_ACCEPTED     EntitlementDriftStatus = "AUTO_ACCEPTED"
	ENTITLEMENTDRIFTSTATUS_AUTO_REJECTED     EntitlementDriftStatus = "AUTO_REJECTED"
	ENTITLEMENTDRIFTSTATUS_CONFLICT_REJECTED EntitlementDriftStatus = "CONFLICT_REJECTED"
	ENTITLEMENTDRIFTSTATUS_APPLIED           EntitlementDriftStatus = "APPLIED"
	ENTITLEMENTDRIFTSTATUS_REVERTED          EntitlementDriftStatus = "REVERTED"
	ENTITLEMENTDRIFTSTATUS_SUPERSEDED        EntitlementDriftStatus = "SUPERSEDED"
	ENTITLEMENTDRIFTSTATUS_ERROR             EntitlementDriftStatus = "ERROR"
)

// All allowed values of EntitlementDriftStatus enum
var AllowedEntitlementDriftStatusEnumValues = []EntitlementDriftStatus{
	"AUTO_ACCEPTED",
	"AUTO_REJECTED",
	"CONFLICT_REJECTED",
	"APPLIED",
	"REVERTED",
	"SUPERSEDED",
	"ERROR",
}

func (v *EntitlementDriftStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementDriftStatus(value)
	for _, existing := range AllowedEntitlementDriftStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementDriftStatus", value)
}

// NewEntitlementDriftStatusFromValue returns a pointer to a valid EntitlementDriftStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementDriftStatusFromValue(v string) (*EntitlementDriftStatus, error) {
	ev := EntitlementDriftStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementDriftStatus: valid values are %v", v, AllowedEntitlementDriftStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementDriftStatus) IsValid() bool {
	for _, existing := range AllowedEntitlementDriftStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-drift-status value
func (v EntitlementDriftStatus) Ptr() *EntitlementDriftStatus {
	return &v
}

type NullableEntitlementDriftStatus struct {
	value *EntitlementDriftStatus
	isSet bool
}

func (v NullableEntitlementDriftStatus) Get() *EntitlementDriftStatus {
	return v.value
}

func (v *NullableEntitlementDriftStatus) Set(val *EntitlementDriftStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftStatus(val *EntitlementDriftStatus) *NullableEntitlementDriftStatus {
	return &NullableEntitlementDriftStatus{value: val, isSet: true}
}

func (v NullableEntitlementDriftStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
