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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["principal-entitlements"] = o.PrincipalEntitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrincipalEntitlementsChangeLinks) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalEntitlementsChangeLinks := _PrincipalEntitlementsChangeLinks{}

	err = json.Unmarshal(bytes, &varPrincipalEntitlementsChangeLinks)
	if err == nil {
		*o = PrincipalEntitlementsChangeLinks(varPrincipalEntitlementsChangeLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "principal-entitlements")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
