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
)

// checks if the GrantDetailsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantDetailsLinks{}

// GrantDetailsLinks struct for GrantDetailsLinks
type GrantDetailsLinks struct {
	Self                 *Link `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantDetailsLinks GrantDetailsLinks

// NewGrantDetailsLinks instantiates a new GrantDetailsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantDetailsLinks() *GrantDetailsLinks {
	this := GrantDetailsLinks{}
	return &this
}

// NewGrantDetailsLinksWithDefaults instantiates a new GrantDetailsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantDetailsLinksWithDefaults() *GrantDetailsLinks {
	this := GrantDetailsLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *GrantDetailsLinks) GetSelf() Link {
	if o == nil || IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantDetailsLinks) GetSelfOk() (*Link, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *GrantDetailsLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *GrantDetailsLinks) SetSelf(v Link) {
	o.Self = &v
}

func (o GrantDetailsLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantDetailsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantDetailsLinks) UnmarshalJSON(data []byte) (err error) {
	varGrantDetailsLinks := _GrantDetailsLinks{}

	err = json.Unmarshal(data, &varGrantDetailsLinks)

	if err != nil {
		return err
	}

	*o = GrantDetailsLinks(varGrantDetailsLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantDetailsLinks struct {
	value *GrantDetailsLinks
	isSet bool
}

func (v NullableGrantDetailsLinks) Get() *GrantDetailsLinks {
	return v.value
}

func (v *NullableGrantDetailsLinks) Set(val *GrantDetailsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantDetailsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantDetailsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantDetailsLinks(val *GrantDetailsLinks) *NullableGrantDetailsLinks {
	return &NullableGrantDetailsLinks{value: val, isSet: true}
}

func (v NullableGrantDetailsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantDetailsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
