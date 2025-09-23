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
// RequestResourceCreatable - A representation of a resource that can be requested for access.  **Note:** The resources available for the request are subject to their permission settings. For example, the Okta Admin App can only be requested by Super Admin users.
type RequestResourceCreatable struct {
	RequestResourceCatalogEntryCreatable *RequestResourceCatalogEntryCreatable
}

// RequestResourceCatalogEntryCreatableAsRequestResourceCreatable is a convenience function that returns RequestResourceCatalogEntryCreatable wrapped in RequestResourceCreatable
func RequestResourceCatalogEntryCreatableAsRequestResourceCreatable(v *RequestResourceCatalogEntryCreatable) RequestResourceCreatable {
	return RequestResourceCreatable{
		RequestResourceCatalogEntryCreatable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequestResourceCreatable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'CATALOG_ENTRY'
	if jsonDict["type"] == "CATALOG_ENTRY" {
		// try to unmarshal JSON data into RequestResourceCatalogEntryCreatable
		err = json.Unmarshal(data, &dst.RequestResourceCatalogEntryCreatable)
		if err == nil {
			return nil // data stored in dst.RequestResourceCatalogEntryCreatable, return on the first match
		} else {
			dst.RequestResourceCatalogEntryCreatable = nil
			return fmt.Errorf("Failed to unmarshal RequestResourceCreatable as RequestResourceCatalogEntryCreatable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-resource-catalog-entry-creatable'
	if jsonDict["type"] == "request-resource-catalog-entry-creatable" {
		// try to unmarshal JSON data into RequestResourceCatalogEntryCreatable
		err = json.Unmarshal(data, &dst.RequestResourceCatalogEntryCreatable)
		if err == nil {
			return nil // data stored in dst.RequestResourceCatalogEntryCreatable, return on the first match
		} else {
			dst.RequestResourceCatalogEntryCreatable = nil
			return fmt.Errorf("Failed to unmarshal RequestResourceCreatable as RequestResourceCatalogEntryCreatable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestResourceCreatable) MarshalJSON() ([]byte, error) {
	if src.RequestResourceCatalogEntryCreatable != nil {
		return json.Marshal(&src.RequestResourceCatalogEntryCreatable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestResourceCreatable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestResourceCatalogEntryCreatable != nil {
		return obj.RequestResourceCatalogEntryCreatable
	}

	// all schemas are nil
	return nil
}

type NullableRequestResourceCreatable struct {
	value *RequestResourceCreatable
	isSet bool
}

func (v NullableRequestResourceCreatable) Get() *RequestResourceCreatable {
	return v.value
}

func (v *NullableRequestResourceCreatable) Set(val *RequestResourceCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResourceCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResourceCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResourceCreatable(val *RequestResourceCreatable) *NullableRequestResourceCreatable {
	return &NullableRequestResourceCreatable{value: val, isSet: true}
}

func (v NullableRequestResourceCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResourceCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
