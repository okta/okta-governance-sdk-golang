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

// checks if the PrincipalEntitlementsChangeLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalEntitlementsChangeLinks{}

// PrincipalEntitlementsChangeLinks Link to current effective entitlements
type PrincipalEntitlementsChangeLinks struct {
	PrincipalEntitlements Link `json:"principal-entitlements"`
	AdditionalProperties  map[string]interface{}
}

type _PrincipalEntitlementsChangeLinks PrincipalEntitlementsChangeLinks

// NewPrincipalEntitlementsChangeLinks instantiates a new PrincipalEntitlementsChangeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementsChangeLinks(principalEntitlements Link) *PrincipalEntitlementsChangeLinks {
	this := PrincipalEntitlementsChangeLinks{}
	this.PrincipalEntitlements = principalEntitlements
	return &this
}

// NewPrincipalEntitlementsChangeLinksWithDefaults instantiates a new PrincipalEntitlementsChangeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementsChangeLinksWithDefaults() *PrincipalEntitlementsChangeLinks {
	this := PrincipalEntitlementsChangeLinks{}
	return &this
}

// GetPrincipalEntitlements returns the PrincipalEntitlements field value
func (o *PrincipalEntitlementsChangeLinks) GetPrincipalEntitlements() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.PrincipalEntitlements
}

// GetPrincipalEntitlementsOk returns a tuple with the PrincipalEntitlements field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsChangeLinks) GetPrincipalEntitlementsOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalEntitlements, true
}

// SetPrincipalEntitlements sets field value
func (o *PrincipalEntitlementsChangeLinks) SetPrincipalEntitlements(v Link) {
	o.PrincipalEntitlements = v
}

func (o PrincipalEntitlementsChangeLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalEntitlementsChangeLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal-entitlements"] = o.PrincipalEntitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalEntitlementsChangeLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal-entitlements",
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

	varPrincipalEntitlementsChangeLinks := _PrincipalEntitlementsChangeLinks{}

	err = json.Unmarshal(data, &varPrincipalEntitlementsChangeLinks)

	if err != nil {
		return err
	}

	*o = PrincipalEntitlementsChangeLinks(varPrincipalEntitlementsChangeLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principal-entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalEntitlementsChangeLinks struct {
	value *PrincipalEntitlementsChangeLinks
	isSet bool
}

func (v NullablePrincipalEntitlementsChangeLinks) Get() *PrincipalEntitlementsChangeLinks {
	return v.value
}

func (v *NullablePrincipalEntitlementsChangeLinks) Set(val *PrincipalEntitlementsChangeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementsChangeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementsChangeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementsChangeLinks(val *PrincipalEntitlementsChangeLinks) *NullablePrincipalEntitlementsChangeLinks {
	return &NullablePrincipalEntitlementsChangeLinks{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementsChangeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementsChangeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
