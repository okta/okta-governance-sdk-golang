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

// EntitlementDriftAffectedGrantSource Where the affected entitlement's grant came from when the drift was detected. `POLICY` and `ENTITLEMENT-BUNDLE` are birthright sources: a user gets them automatically, and reconciliation never removes them.  Because Okta never removes a birthright entitlement, a `SUB` from a `POLICY` source is always `CONFLICT_REJECTED`. A `SUB` from an `ENTITLEMENT-BUNDLE` source is too, with one exception: if the import removes every entitlement the bundle granted, Okta accepts the removals and drops the whole bundle grant.
type EntitlementDriftAffectedGrantSource string

// List of entitlement-drift-affected-grant-source
const (
	ENTITLEMENTDRIFTAFFECTEDGRANTSOURCE_POLICY             EntitlementDriftAffectedGrantSource = "POLICY"
	ENTITLEMENTDRIFTAFFECTEDGRANTSOURCE_CUSTOM             EntitlementDriftAffectedGrantSource = "CUSTOM"
	ENTITLEMENTDRIFTAFFECTEDGRANTSOURCE_ENTITLEMENT_BUNDLE EntitlementDriftAffectedGrantSource = "ENTITLEMENT-BUNDLE"
	ENTITLEMENTDRIFTAFFECTEDGRANTSOURCE_ENTITLEMENT        EntitlementDriftAffectedGrantSource = "ENTITLEMENT"
	ENTITLEMENTDRIFTAFFECTEDGRANTSOURCE_IMPORT             EntitlementDriftAffectedGrantSource = "IMPORT"
)

// All allowed values of EntitlementDriftAffectedGrantSource enum
var AllowedEntitlementDriftAffectedGrantSourceEnumValues = []EntitlementDriftAffectedGrantSource{
	"POLICY",
	"CUSTOM",
	"ENTITLEMENT-BUNDLE",
	"ENTITLEMENT",
	"IMPORT",
}

func (v *EntitlementDriftAffectedGrantSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementDriftAffectedGrantSource(value)
	for _, existing := range AllowedEntitlementDriftAffectedGrantSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementDriftAffectedGrantSource", value)
}

// NewEntitlementDriftAffectedGrantSourceFromValue returns a pointer to a valid EntitlementDriftAffectedGrantSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementDriftAffectedGrantSourceFromValue(v string) (*EntitlementDriftAffectedGrantSource, error) {
	ev := EntitlementDriftAffectedGrantSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementDriftAffectedGrantSource: valid values are %v", v, AllowedEntitlementDriftAffectedGrantSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementDriftAffectedGrantSource) IsValid() bool {
	for _, existing := range AllowedEntitlementDriftAffectedGrantSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-drift-affected-grant-source value
func (v EntitlementDriftAffectedGrantSource) Ptr() *EntitlementDriftAffectedGrantSource {
	return &v
}

type NullableEntitlementDriftAffectedGrantSource struct {
	value *EntitlementDriftAffectedGrantSource
	isSet bool
}

func (v NullableEntitlementDriftAffectedGrantSource) Get() *EntitlementDriftAffectedGrantSource {
	return v.value
}

func (v *NullableEntitlementDriftAffectedGrantSource) Set(val *EntitlementDriftAffectedGrantSource) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftAffectedGrantSource) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftAffectedGrantSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftAffectedGrantSource(val *EntitlementDriftAffectedGrantSource) *NullableEntitlementDriftAffectedGrantSource {
	return &NullableEntitlementDriftAffectedGrantSource{value: val, isSet: true}
}

func (v NullableEntitlementDriftAffectedGrantSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftAffectedGrantSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
