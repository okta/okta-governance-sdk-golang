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

// GetGrant200Response - struct for GetGrant200Response
type GetGrant200Response struct {
	GrantFull                 *GrantFull
	GrantFullWithEntitlements *GrantFullWithEntitlements
}

// GrantFullAsGetGrant200Response is a convenience function that returns GrantFull wrapped in GetGrant200Response
func GrantFullAsGetGrant200Response(v *GrantFull) GetGrant200Response {
	return GetGrant200Response{
		GrantFull: v,
	}
}

// GrantFullWithEntitlementsAsGetGrant200Response is a convenience function that returns GrantFullWithEntitlements wrapped in GetGrant200Response
func GrantFullWithEntitlementsAsGetGrant200Response(v *GrantFullWithEntitlements) GetGrant200Response {
	return GetGrant200Response{
		GrantFullWithEntitlements: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGrant200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GrantFull
	err = json.Unmarshal(data, &dst.GrantFull)
	if err == nil {
		jsonGrantFull, _ := json.Marshal(dst.GrantFull)
		if string(jsonGrantFull) == "{}" { // empty struct
			dst.GrantFull = nil
		} else {
			match++
		}
	} else {
		dst.GrantFull = nil
	}

	// try to unmarshal data into GrantFullWithEntitlements
	err = json.Unmarshal(data, &dst.GrantFullWithEntitlements)
	if err == nil {
		jsonGrantFullWithEntitlements, _ := json.Marshal(dst.GrantFullWithEntitlements)
		if string(jsonGrantFullWithEntitlements) == "{}" { // empty struct
			dst.GrantFullWithEntitlements = nil
		} else {
			match++
		}
	} else {
		dst.GrantFullWithEntitlements = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GrantFull = nil
		dst.GrantFullWithEntitlements = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetGrant200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetGrant200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGrant200Response) MarshalJSON() ([]byte, error) {
	if src.GrantFull != nil {
		return json.Marshal(&src.GrantFull)
	}

	if src.GrantFullWithEntitlements != nil {
		return json.Marshal(&src.GrantFullWithEntitlements)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGrant200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GrantFull != nil {
		return obj.GrantFull
	}

	if obj.GrantFullWithEntitlements != nil {
		return obj.GrantFullWithEntitlements
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetGrant200Response) GetActualInstanceValue() interface{} {
	if obj.GrantFull != nil {
		return *obj.GrantFull
	}

	if obj.GrantFullWithEntitlements != nil {
		return *obj.GrantFullWithEntitlements
	}

	// all schemas are nil
	return nil
}

type NullableGetGrant200Response struct {
	value *GetGrant200Response
	isSet bool
}

func (v NullableGetGrant200Response) Get() *GetGrant200Response {
	return v.value
}

func (v *NullableGetGrant200Response) Set(val *GetGrant200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGrant200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGrant200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGrant200Response(val *GetGrant200Response) *NullableGetGrant200Response {
	return &NullableGetGrant200Response{value: val, isSet: true}
}

func (v NullableGetGrant200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGrant200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
