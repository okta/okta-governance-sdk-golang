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

// EntitlementLinks Links available in list response
type EntitlementLinks struct {
	Self                 Link  `json:"self"`
	Values               *Link `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementLinks EntitlementLinks

// NewEntitlementLinks instantiates a new EntitlementLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementLinks(self Link) *EntitlementLinks {
	this := EntitlementLinks{}
	this.Self = self
	return &this
}

// NewEntitlementLinksWithDefaults instantiates a new EntitlementLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementLinksWithDefaults() *EntitlementLinks {
	this := EntitlementLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *EntitlementLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *EntitlementLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *EntitlementLinks) SetSelf(v Link) {
	o.Self = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementLinks) GetValues() Link {
	if o == nil || o.Values == nil {
		var ret Link
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementLinks) GetValuesOk() (*Link, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementLinks) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given Link and assigns it to the Values field.
func (o *EntitlementLinks) SetValues(v Link) {
	o.Values = &v
}

func (o EntitlementLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementLinks) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementLinks := _EntitlementLinks{}

	err = json.Unmarshal(bytes, &varEntitlementLinks)
	if err == nil {
		*o = EntitlementLinks(varEntitlementLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementLinks struct {
	value *EntitlementLinks
	isSet bool
}

func (v NullableEntitlementLinks) Get() *EntitlementLinks {
	return v.value
}

func (v *NullableEntitlementLinks) Set(val *EntitlementLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementLinks(val *EntitlementLinks) *NullableEntitlementLinks {
	return &NullableEntitlementLinks{value: val, isSet: true}
}

func (v NullableEntitlementLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
