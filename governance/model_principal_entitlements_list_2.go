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

// checks if the PrincipalEntitlementsList2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalEntitlementsList2{}

// PrincipalEntitlementsList2 A principal's effective entitlements on a resource. `data[]` is a uniform list of rows keyed by `targetOrn`. For resources that don't support assets, it's a single row for the resource itself. For resources that support assets, it's the resource assets that the principal can reach. Each row carries the principal's complete effective entitlements on that target. Sorted by name. The `_links.next` reference is present when more rows remain.
type PrincipalEntitlementsList2 struct {
	// The Okta user in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format
	PrincipalOrn string `json:"principalOrn"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn string `json:"resourceOrn"`
	// The principal's effective-entitlement rows. Each row is the resource itself or one of its resource assets.
	Data                 []PrincipalEntitlementRow `json:"data"`
	Links                ListLinks                 `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementsList2 PrincipalEntitlementsList2

// NewPrincipalEntitlementsList2 instantiates a new PrincipalEntitlementsList2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementsList2(principalOrn string, resourceOrn string, data []PrincipalEntitlementRow, links ListLinks) *PrincipalEntitlementsList2 {
	this := PrincipalEntitlementsList2{}
	this.PrincipalOrn = principalOrn
	this.ResourceOrn = resourceOrn
	this.Data = data
	this.Links = links
	return &this
}

// NewPrincipalEntitlementsList2WithDefaults instantiates a new PrincipalEntitlementsList2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementsList2WithDefaults() *PrincipalEntitlementsList2 {
	this := PrincipalEntitlementsList2{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *PrincipalEntitlementsList2) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsList2) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *PrincipalEntitlementsList2) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *PrincipalEntitlementsList2) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsList2) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *PrincipalEntitlementsList2) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetData returns the Data field value
func (o *PrincipalEntitlementsList2) GetData() []PrincipalEntitlementRow {
	if o == nil {
		var ret []PrincipalEntitlementRow
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsList2) GetDataOk() ([]PrincipalEntitlementRow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PrincipalEntitlementsList2) SetData(v []PrincipalEntitlementRow) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *PrincipalEntitlementsList2) GetLinks() ListLinks {
	if o == nil {
		var ret ListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsList2) GetLinksOk() (*ListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PrincipalEntitlementsList2) SetLinks(v ListLinks) {
	o.Links = v
}

func (o PrincipalEntitlementsList2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalEntitlementsList2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalOrn"] = o.PrincipalOrn
	toSerialize["resourceOrn"] = o.ResourceOrn
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalEntitlementsList2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalOrn",
		"resourceOrn",
		"data",
		"_links",
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

	varPrincipalEntitlementsList2 := _PrincipalEntitlementsList2{}

	err = json.Unmarshal(data, &varPrincipalEntitlementsList2)

	if err != nil {
		return err
	}

	*o = PrincipalEntitlementsList2(varPrincipalEntitlementsList2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalEntitlementsList2 struct {
	value *PrincipalEntitlementsList2
	isSet bool
}

func (v NullablePrincipalEntitlementsList2) Get() *PrincipalEntitlementsList2 {
	return v.value
}

func (v *NullablePrincipalEntitlementsList2) Set(val *PrincipalEntitlementsList2) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementsList2) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementsList2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementsList2(val *PrincipalEntitlementsList2) *NullablePrincipalEntitlementsList2 {
	return &NullablePrincipalEntitlementsList2{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementsList2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementsList2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
