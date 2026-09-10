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

// EntitlementDriftResolution How Okta resolved the drift.  * `AUTO_ACCEPTED`: Okta honored the change. An `ADD` becomes a new `IMPORT` grant. A `SUB` removes the entitlement from the grant that provided it.  * `AUTO_REJECTED`: Okta rejected the change, based on the action that was configured for that direction, and pushed the original assignment back to the app.  * `CONFLICT_REJECTED`: Okta rejected a `SUB` because the entitlement came from a birthright source (`POLICY` or `ENTITLEMENT-BUNDLE`), which it never removes, and pushed the original assignment back to the app.
type EntitlementDriftResolution string

// List of entitlement-drift-resolution
const (
	ENTITLEMENTDRIFTRESOLUTION_AUTO_ACCEPTED     EntitlementDriftResolution = "AUTO_ACCEPTED"
	ENTITLEMENTDRIFTRESOLUTION_AUTO_REJECTED     EntitlementDriftResolution = "AUTO_REJECTED"
	ENTITLEMENTDRIFTRESOLUTION_CONFLICT_REJECTED EntitlementDriftResolution = "CONFLICT_REJECTED"
)

// All allowed values of EntitlementDriftResolution enum
var AllowedEntitlementDriftResolutionEnumValues = []EntitlementDriftResolution{
	"AUTO_ACCEPTED",
	"AUTO_REJECTED",
	"CONFLICT_REJECTED",
}

func (v *EntitlementDriftResolution) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementDriftResolution(value)
	for _, existing := range AllowedEntitlementDriftResolutionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementDriftResolution", value)
}

// NewEntitlementDriftResolutionFromValue returns a pointer to a valid EntitlementDriftResolution
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementDriftResolutionFromValue(v string) (*EntitlementDriftResolution, error) {
	ev := EntitlementDriftResolution(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementDriftResolution: valid values are %v", v, AllowedEntitlementDriftResolutionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementDriftResolution) IsValid() bool {
	for _, existing := range AllowedEntitlementDriftResolutionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-drift-resolution value
func (v EntitlementDriftResolution) Ptr() *EntitlementDriftResolution {
	return &v
}

type NullableEntitlementDriftResolution struct {
	value *EntitlementDriftResolution
	isSet bool
}

func (v NullableEntitlementDriftResolution) Get() *EntitlementDriftResolution {
	return v.value
}

func (v *NullableEntitlementDriftResolution) Set(val *EntitlementDriftResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftResolution(val *EntitlementDriftResolution) *NullableEntitlementDriftResolution {
	return &NullableEntitlementDriftResolution{value: val, isSet: true}
}

func (v NullableEntitlementDriftResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
