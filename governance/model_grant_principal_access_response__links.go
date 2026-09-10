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

// checks if the GrantPrincipalAccessResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPrincipalAccessResponseLinks{}

// GrantPrincipalAccessResponseLinks Links to related resources
type GrantPrincipalAccessResponseLinks struct {
	PrincipalAccess      GrantPrincipalAccessResponseLinksPrincipalAccess `json:"principalAccess"`
	AdditionalProperties map[string]interface{}
}

type _GrantPrincipalAccessResponseLinks GrantPrincipalAccessResponseLinks

// NewGrantPrincipalAccessResponseLinks instantiates a new GrantPrincipalAccessResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPrincipalAccessResponseLinks(principalAccess GrantPrincipalAccessResponseLinksPrincipalAccess) *GrantPrincipalAccessResponseLinks {
	this := GrantPrincipalAccessResponseLinks{}
	this.PrincipalAccess = principalAccess
	return &this
}

// NewGrantPrincipalAccessResponseLinksWithDefaults instantiates a new GrantPrincipalAccessResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPrincipalAccessResponseLinksWithDefaults() *GrantPrincipalAccessResponseLinks {
	this := GrantPrincipalAccessResponseLinks{}
	return &this
}

// GetPrincipalAccess returns the PrincipalAccess field value
func (o *GrantPrincipalAccessResponseLinks) GetPrincipalAccess() GrantPrincipalAccessResponseLinksPrincipalAccess {
	if o == nil {
		var ret GrantPrincipalAccessResponseLinksPrincipalAccess
		return ret
	}

	return o.PrincipalAccess
}

// GetPrincipalAccessOk returns a tuple with the PrincipalAccess field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessResponseLinks) GetPrincipalAccessOk() (*GrantPrincipalAccessResponseLinksPrincipalAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalAccess, true
}

// SetPrincipalAccess sets field value
func (o *GrantPrincipalAccessResponseLinks) SetPrincipalAccess(v GrantPrincipalAccessResponseLinksPrincipalAccess) {
	o.PrincipalAccess = v
}

func (o GrantPrincipalAccessResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPrincipalAccessResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalAccess"] = o.PrincipalAccess

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPrincipalAccessResponseLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalAccess",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGrantPrincipalAccessResponseLinks := _GrantPrincipalAccessResponseLinks{}

	err = json.Unmarshal(data, &varGrantPrincipalAccessResponseLinks)

	if err != nil {
		return err
	}

	*o = GrantPrincipalAccessResponseLinks(varGrantPrincipalAccessResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalAccess")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPrincipalAccessResponseLinks struct {
	value *GrantPrincipalAccessResponseLinks
	isSet bool
}

func (v NullableGrantPrincipalAccessResponseLinks) Get() *GrantPrincipalAccessResponseLinks {
	return v.value
}

func (v *NullableGrantPrincipalAccessResponseLinks) Set(val *GrantPrincipalAccessResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalAccessResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalAccessResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalAccessResponseLinks(val *GrantPrincipalAccessResponseLinks) *NullableGrantPrincipalAccessResponseLinks {
	return &NullableGrantPrincipalAccessResponseLinks{value: val, isSet: true}
}

func (v NullableGrantPrincipalAccessResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalAccessResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
