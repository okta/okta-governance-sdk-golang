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

// EntitlementValueCreatable Entitlement value id
type EntitlementValueCreatable struct {
	// The `id` of an entitlement value
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValueCreatable EntitlementValueCreatable

// NewEntitlementValueCreatable instantiates a new EntitlementValueCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValueCreatable() *EntitlementValueCreatable {
	this := EntitlementValueCreatable{}
	return &this
}

// NewEntitlementValueCreatableWithDefaults instantiates a new EntitlementValueCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueCreatableWithDefaults() *EntitlementValueCreatable {
	this := EntitlementValueCreatable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementValueCreatable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueCreatable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementValueCreatable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementValueCreatable) SetId(v string) {
	o.Id = &v
}

func (o EntitlementValueCreatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValueCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValueCreatable := _EntitlementValueCreatable{}

	err = json.Unmarshal(bytes, &varEntitlementValueCreatable)
	if err == nil {
		*o = EntitlementValueCreatable(varEntitlementValueCreatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValueCreatable struct {
	value *EntitlementValueCreatable
	isSet bool
}

func (v NullableEntitlementValueCreatable) Get() *EntitlementValueCreatable {
	return v.value
}

func (v *NullableEntitlementValueCreatable) Set(val *EntitlementValueCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValueCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValueCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValueCreatable(val *EntitlementValueCreatable) *NullableEntitlementValueCreatable {
	return &NullableEntitlementValueCreatable{value: val, isSet: true}
}

func (v NullableEntitlementValueCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValueCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
