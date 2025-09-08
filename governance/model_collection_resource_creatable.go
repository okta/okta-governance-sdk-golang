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
)

// CollectionResourceCreatable The properties expected when adding a new resource to a collection
type CollectionResourceCreatable struct {
	// Collection of entitlements and associated value identifiers
	Entitlements []EntitlementCreatable `json:"entitlements,omitempty"`
	// The ORN identifier for a specific app. Other resource types aren't supported.  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint for reference.
	ResourceOrn          string `json:"resourceOrn"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceCreatable CollectionResourceCreatable

// NewCollectionResourceCreatable instantiates a new CollectionResourceCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceCreatable(resourceOrn string) *CollectionResourceCreatable {
	this := CollectionResourceCreatable{}
	this.ResourceOrn = resourceOrn
	return &this
}

// NewCollectionResourceCreatableWithDefaults instantiates a new CollectionResourceCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceCreatableWithDefaults() *CollectionResourceCreatable {
	this := CollectionResourceCreatable{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceCreatable) GetEntitlements() []EntitlementCreatable {
	if o == nil || o.Entitlements == nil {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceCreatable) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceCreatable) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *CollectionResourceCreatable) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionResourceCreatable) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceCreatable) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionResourceCreatable) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

func (o CollectionResourceCreatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if true {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionResourceCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionResourceCreatable := _CollectionResourceCreatable{}

	err = json.Unmarshal(bytes, &varCollectionResourceCreatable)
	if err == nil {
		*o = CollectionResourceCreatable(varCollectionResourceCreatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "resourceOrn")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCollectionResourceCreatable struct {
	value *CollectionResourceCreatable
	isSet bool
}

func (v NullableCollectionResourceCreatable) Get() *CollectionResourceCreatable {
	return v.value
}

func (v *NullableCollectionResourceCreatable) Set(val *CollectionResourceCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceCreatable(val *CollectionResourceCreatable) *NullableCollectionResourceCreatable {
	return &NullableCollectionResourceCreatable{value: val, isSet: true}
}

func (v NullableCollectionResourceCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
