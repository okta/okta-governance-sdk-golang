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

// ListGrants200Response - struct for ListGrants200Response
type ListGrants200Response struct {
	GrantsList                 *GrantsList
	GrantsListWithEntitlements *GrantsListWithEntitlements
}

// GrantsListAsListGrants200Response is a convenience function that returns GrantsList wrapped in ListGrants200Response
func GrantsListAsListGrants200Response(v *GrantsList) ListGrants200Response {
	return ListGrants200Response{
		GrantsList: v,
	}
}

// GrantsListWithEntitlementsAsListGrants200Response is a convenience function that returns GrantsListWithEntitlements wrapped in ListGrants200Response
func GrantsListWithEntitlementsAsListGrants200Response(v *GrantsListWithEntitlements) ListGrants200Response {
	return ListGrants200Response{
		GrantsListWithEntitlements: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListGrants200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GrantsList
	err = json.Unmarshal(data, &dst.GrantsList)
	if err == nil {
		jsonGrantsList, _ := json.Marshal(dst.GrantsList)
		if string(jsonGrantsList) == "{}" { // empty struct
			dst.GrantsList = nil
		} else {
			match++
		}
	} else {
		dst.GrantsList = nil
	}

	// try to unmarshal data into GrantsListWithEntitlements
	err = json.Unmarshal(data, &dst.GrantsListWithEntitlements)
	if err == nil {
		jsonGrantsListWithEntitlements, _ := json.Marshal(dst.GrantsListWithEntitlements)
		if string(jsonGrantsListWithEntitlements) == "{}" { // empty struct
			dst.GrantsListWithEntitlements = nil
		} else {
			match++
		}
	} else {
		dst.GrantsListWithEntitlements = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GrantsList = nil
		dst.GrantsListWithEntitlements = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListGrants200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListGrants200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListGrants200Response) MarshalJSON() ([]byte, error) {
	if src.GrantsList != nil {
		return json.Marshal(&src.GrantsList)
	}

	if src.GrantsListWithEntitlements != nil {
		return json.Marshal(&src.GrantsListWithEntitlements)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListGrants200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GrantsList != nil {
		return obj.GrantsList
	}

	if obj.GrantsListWithEntitlements != nil {
		return obj.GrantsListWithEntitlements
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ListGrants200Response) GetActualInstanceValue() interface{} {
	if obj.GrantsList != nil {
		return *obj.GrantsList
	}

	if obj.GrantsListWithEntitlements != nil {
		return *obj.GrantsListWithEntitlements
	}

	// all schemas are nil
	return nil
}

type NullableListGrants200Response struct {
	value *ListGrants200Response
	isSet bool
}

func (v NullableListGrants200Response) Get() *ListGrants200Response {
	return v.value
}

func (v *NullableListGrants200Response) Set(val *ListGrants200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListGrants200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListGrants200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGrants200Response(val *ListGrants200Response) *NullableListGrants200Response {
	return &NullableListGrants200Response{value: val, isSet: true}
}

func (v NullableListGrants200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGrants200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
