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

// PrincipalEntitlementsList struct for PrincipalEntitlementsList
type PrincipalEntitlementsList struct {
	// Principal entitlements list
	Data                 []PrincipalEntitlement          `json:"data,omitempty"`
	Links                *PrincipalEntitlementsListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementsList PrincipalEntitlementsList

// NewPrincipalEntitlementsList instantiates a new PrincipalEntitlementsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementsList() *PrincipalEntitlementsList {
	this := PrincipalEntitlementsList{}
	return &this
}

// NewPrincipalEntitlementsListWithDefaults instantiates a new PrincipalEntitlementsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementsListWithDefaults() *PrincipalEntitlementsList {
	this := PrincipalEntitlementsList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PrincipalEntitlementsList) GetData() []PrincipalEntitlement {
	if o == nil || o.Data == nil {
		var ret []PrincipalEntitlement
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsList) GetDataOk() ([]PrincipalEntitlement, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PrincipalEntitlementsList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []PrincipalEntitlement and assigns it to the Data field.
func (o *PrincipalEntitlementsList) SetData(v []PrincipalEntitlement) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PrincipalEntitlementsList) GetLinks() PrincipalEntitlementsListLinks {
	if o == nil || o.Links == nil {
		var ret PrincipalEntitlementsListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsList) GetLinksOk() (*PrincipalEntitlementsListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PrincipalEntitlementsList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PrincipalEntitlementsListLinks and assigns it to the Links field.
func (o *PrincipalEntitlementsList) SetLinks(v PrincipalEntitlementsListLinks) {
	o.Links = &v
}

func (o PrincipalEntitlementsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrincipalEntitlementsList) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalEntitlementsList := _PrincipalEntitlementsList{}

	err = json.Unmarshal(bytes, &varPrincipalEntitlementsList)
	if err == nil {
		*o = PrincipalEntitlementsList(varPrincipalEntitlementsList)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePrincipalEntitlementsList struct {
	value *PrincipalEntitlementsList
	isSet bool
}

func (v NullablePrincipalEntitlementsList) Get() *PrincipalEntitlementsList {
	return v.value
}

func (v *NullablePrincipalEntitlementsList) Set(val *PrincipalEntitlementsList) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementsList) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementsList(val *PrincipalEntitlementsList) *NullablePrincipalEntitlementsList {
	return &NullablePrincipalEntitlementsList{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
