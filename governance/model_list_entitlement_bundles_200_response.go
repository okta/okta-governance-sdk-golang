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

// model_oneof.mustache
// ListEntitlementBundles200Response - struct for ListEntitlementBundles200Response
type ListEntitlementBundles200Response struct {
	EntitlementBundlesList                 *EntitlementBundlesList
	EntitlementBundlesListWithEntitlements *EntitlementBundlesListWithEntitlements
}

// EntitlementBundlesListAsListEntitlementBundles200Response is a convenience function that returns EntitlementBundlesList wrapped in ListEntitlementBundles200Response
func EntitlementBundlesListAsListEntitlementBundles200Response(v *EntitlementBundlesList) ListEntitlementBundles200Response {
	return ListEntitlementBundles200Response{
		EntitlementBundlesList: v,
	}
}

// EntitlementBundlesListWithEntitlementsAsListEntitlementBundles200Response is a convenience function that returns EntitlementBundlesListWithEntitlements wrapped in ListEntitlementBundles200Response
func EntitlementBundlesListWithEntitlementsAsListEntitlementBundles200Response(v *EntitlementBundlesListWithEntitlements) ListEntitlementBundles200Response {
	return ListEntitlementBundles200Response{
		EntitlementBundlesListWithEntitlements: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListEntitlementBundles200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EntitlementBundlesList
	err = json.Unmarshal(data, &dst.EntitlementBundlesList)
	if err == nil {
		jsonEntitlementBundlesList, _ := json.Marshal(dst.EntitlementBundlesList)
		if string(jsonEntitlementBundlesList) == "{}" { // empty struct
			dst.EntitlementBundlesList = nil
		} else {
			match++
		}
	} else {
		dst.EntitlementBundlesList = nil
	}

	// try to unmarshal data into EntitlementBundlesListWithEntitlements
	err = json.Unmarshal(data, &dst.EntitlementBundlesListWithEntitlements)
	if err == nil {
		jsonEntitlementBundlesListWithEntitlements, _ := json.Marshal(dst.EntitlementBundlesListWithEntitlements)
		if string(jsonEntitlementBundlesListWithEntitlements) == "{}" { // empty struct
			dst.EntitlementBundlesListWithEntitlements = nil
		} else {
			match++
		}
	} else {
		dst.EntitlementBundlesListWithEntitlements = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EntitlementBundlesList = nil
		dst.EntitlementBundlesListWithEntitlements = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListEntitlementBundles200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListEntitlementBundles200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListEntitlementBundles200Response) MarshalJSON() ([]byte, error) {
	if src.EntitlementBundlesList != nil {
		return json.Marshal(&src.EntitlementBundlesList)
	}

	if src.EntitlementBundlesListWithEntitlements != nil {
		return json.Marshal(&src.EntitlementBundlesListWithEntitlements)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListEntitlementBundles200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EntitlementBundlesList != nil {
		return obj.EntitlementBundlesList
	}

	if obj.EntitlementBundlesListWithEntitlements != nil {
		return obj.EntitlementBundlesListWithEntitlements
	}

	// all schemas are nil
	return nil
}

type NullableListEntitlementBundles200Response struct {
	value *ListEntitlementBundles200Response
	isSet bool
}

func (v NullableListEntitlementBundles200Response) Get() *ListEntitlementBundles200Response {
	return v.value
}

func (v *NullableListEntitlementBundles200Response) Set(val *ListEntitlementBundles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListEntitlementBundles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListEntitlementBundles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEntitlementBundles200Response(val *ListEntitlementBundles200Response) *NullableListEntitlementBundles200Response {
	return &NullableListEntitlementBundles200Response{value: val, isSet: true}
}

func (v NullableListEntitlementBundles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEntitlementBundles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
