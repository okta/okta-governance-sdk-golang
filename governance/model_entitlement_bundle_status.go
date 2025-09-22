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

// EntitlementBundleStatus The lifecycle status of an entitlement bundle:  * An entitlement bundle has a status of `ACTIVE` after creation. * When the bundle is deleted, it transitions to a status of `DELETED`.  When the entitlement bundle is purged, it isn't returned in a GET operation.
type EntitlementBundleStatus string

// List of entitlement-bundle-status
const (
	ENTITLEMENTBUNDLESTATUS_ACTIVE  EntitlementBundleStatus = "ACTIVE"
	ENTITLEMENTBUNDLESTATUS_DELETED EntitlementBundleStatus = "DELETED"
)

// All allowed values of EntitlementBundleStatus enum
var AllowedEntitlementBundleStatusEnumValues = []EntitlementBundleStatus{
	"ACTIVE",
	"DELETED",
}

func (v *EntitlementBundleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntitlementBundleStatus(value)
	for _, existing := range AllowedEntitlementBundleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntitlementBundleStatus", value)
}

// NewEntitlementBundleStatusFromValue returns a pointer to a valid EntitlementBundleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntitlementBundleStatusFromValue(v string) (*EntitlementBundleStatus, error) {
	ev := EntitlementBundleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntitlementBundleStatus: valid values are %v", v, AllowedEntitlementBundleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntitlementBundleStatus) IsValid() bool {
	for _, existing := range AllowedEntitlementBundleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to entitlement-bundle-status value
func (v EntitlementBundleStatus) Ptr() *EntitlementBundleStatus {
	return &v
}

type NullableEntitlementBundleStatus struct {
	value *EntitlementBundleStatus
	isSet bool
}

func (v NullableEntitlementBundleStatus) Get() *EntitlementBundleStatus {
	return v.value
}

func (v *NullableEntitlementBundleStatus) Set(val *EntitlementBundleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBundleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBundleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBundleStatus(val *EntitlementBundleStatus) *NullableEntitlementBundleStatus {
	return &NullableEntitlementBundleStatus{value: val, isSet: true}
}

func (v NullableEntitlementBundleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBundleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
