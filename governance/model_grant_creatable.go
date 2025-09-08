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
// GrantCreatable - The properties expected in an initial add entitlement bundle
type GrantCreatable struct {
	GrantTypeBundleWriteable *GrantTypeBundleWriteable
	GrantTypeCustomWriteable *GrantTypeCustomWriteable
	GrantTypePolicyWriteable *GrantTypePolicyWriteable
}

// GrantTypeBundleWriteableAsGrantCreatable is a convenience function that returns GrantTypeBundleWriteable wrapped in GrantCreatable
func GrantTypeBundleWriteableAsGrantCreatable(v *GrantTypeBundleWriteable) GrantCreatable {
	return GrantCreatable{
		GrantTypeBundleWriteable: v,
	}
}

// GrantTypeCustomWriteableAsGrantCreatable is a convenience function that returns GrantTypeCustomWriteable wrapped in GrantCreatable
func GrantTypeCustomWriteableAsGrantCreatable(v *GrantTypeCustomWriteable) GrantCreatable {
	return GrantCreatable{
		GrantTypeCustomWriteable: v,
	}
}

// GrantTypePolicyWriteableAsGrantCreatable is a convenience function that returns GrantTypePolicyWriteable wrapped in GrantCreatable
func GrantTypePolicyWriteableAsGrantCreatable(v *GrantTypePolicyWriteable) GrantCreatable {
	return GrantCreatable{
		GrantTypePolicyWriteable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *GrantCreatable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["grantType"] == "CUSTOM" {
		// try to unmarshal JSON data into GrantTypeCustomWriteable
		err = json.Unmarshal(data, &dst.GrantTypeCustomWriteable)
		if err == nil {
			return nil // data stored in dst.GrantTypeCustomWriteable, return on the first match
		} else {
			dst.GrantTypeCustomWriteable = nil
			return fmt.Errorf("Failed to unmarshal GrantCreatable as GrantTypeCustomWriteable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENTITLEMENT-BUNDLE'
	if jsonDict["grantType"] == "ENTITLEMENT-BUNDLE" {
		// try to unmarshal JSON data into GrantTypeBundleWriteable
		err = json.Unmarshal(data, &dst.GrantTypeBundleWriteable)
		if err == nil {
			return nil // data stored in dst.GrantTypeBundleWriteable, return on the first match
		} else {
			dst.GrantTypeBundleWriteable = nil
			return fmt.Errorf("Failed to unmarshal GrantCreatable as GrantTypeBundleWriteable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'POLICY'
	if jsonDict["grantType"] == "POLICY" {
		// try to unmarshal JSON data into GrantTypePolicyWriteable
		err = json.Unmarshal(data, &dst.GrantTypePolicyWriteable)
		if err == nil {
			return nil // data stored in dst.GrantTypePolicyWriteable, return on the first match
		} else {
			dst.GrantTypePolicyWriteable = nil
			return fmt.Errorf("Failed to unmarshal GrantCreatable as GrantTypePolicyWriteable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'grant-type-bundle-writeable'
	if jsonDict["grantType"] == "grant-type-bundle-writeable" {
		// try to unmarshal JSON data into GrantTypeBundleWriteable
		err = json.Unmarshal(data, &dst.GrantTypeBundleWriteable)
		if err == nil {
			return nil // data stored in dst.GrantTypeBundleWriteable, return on the first match
		} else {
			dst.GrantTypeBundleWriteable = nil
			return fmt.Errorf("Failed to unmarshal GrantCreatable as GrantTypeBundleWriteable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'grant-type-custom-writeable'
	if jsonDict["grantType"] == "grant-type-custom-writeable" {
		// try to unmarshal JSON data into GrantTypeCustomWriteable
		err = json.Unmarshal(data, &dst.GrantTypeCustomWriteable)
		if err == nil {
			return nil // data stored in dst.GrantTypeCustomWriteable, return on the first match
		} else {
			dst.GrantTypeCustomWriteable = nil
			return fmt.Errorf("Failed to unmarshal GrantCreatable as GrantTypeCustomWriteable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'grant-type-policy-writeable'
	if jsonDict["grantType"] == "grant-type-policy-writeable" {
		// try to unmarshal JSON data into GrantTypePolicyWriteable
		err = json.Unmarshal(data, &dst.GrantTypePolicyWriteable)
		if err == nil {
			return nil // data stored in dst.GrantTypePolicyWriteable, return on the first match
		} else {
			dst.GrantTypePolicyWriteable = nil
			return fmt.Errorf("Failed to unmarshal GrantCreatable as GrantTypePolicyWriteable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GrantCreatable) MarshalJSON() ([]byte, error) {
	if src.GrantTypeBundleWriteable != nil {
		return json.Marshal(&src.GrantTypeBundleWriteable)
	}

	if src.GrantTypeCustomWriteable != nil {
		return json.Marshal(&src.GrantTypeCustomWriteable)
	}

	if src.GrantTypePolicyWriteable != nil {
		return json.Marshal(&src.GrantTypePolicyWriteable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GrantCreatable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GrantTypeBundleWriteable != nil {
		return obj.GrantTypeBundleWriteable
	}

	if obj.GrantTypeCustomWriteable != nil {
		return obj.GrantTypeCustomWriteable
	}

	if obj.GrantTypePolicyWriteable != nil {
		return obj.GrantTypePolicyWriteable
	}

	// all schemas are nil
	return nil
}

type NullableGrantCreatable struct {
	value *GrantCreatable
	isSet bool
}

func (v NullableGrantCreatable) Get() *GrantCreatable {
	return v.value
}

func (v *NullableGrantCreatable) Set(val *GrantCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantCreatable(val *GrantCreatable) *NullableGrantCreatable {
	return &NullableGrantCreatable{value: val, isSet: true}
}

func (v NullableGrantCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
