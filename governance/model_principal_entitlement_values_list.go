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

// checks if the PrincipalEntitlementValuesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalEntitlementValuesList{}

// PrincipalEntitlementValuesList Cursor-paginated list of a principal's effective values for one entitlement property (optionally scoped to a single asset). Sorted by value name. The `_links.next` reference is present when more values remain.
type PrincipalEntitlementValuesList struct {
	// The principal's effective entitlement values for the entitlement property on the target
	Data                 []PrincipalEntitlementValue `json:"data"`
	Links                ListLinks                   `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementValuesList PrincipalEntitlementValuesList

// NewPrincipalEntitlementValuesList instantiates a new PrincipalEntitlementValuesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementValuesList(data []PrincipalEntitlementValue, links ListLinks) *PrincipalEntitlementValuesList {
	this := PrincipalEntitlementValuesList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPrincipalEntitlementValuesListWithDefaults instantiates a new PrincipalEntitlementValuesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementValuesListWithDefaults() *PrincipalEntitlementValuesList {
	this := PrincipalEntitlementValuesList{}
	return &this
}

// GetData returns the Data field value
func (o *PrincipalEntitlementValuesList) GetData() []PrincipalEntitlementValue {
	if o == nil {
		var ret []PrincipalEntitlementValue
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValuesList) GetDataOk() ([]PrincipalEntitlementValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PrincipalEntitlementValuesList) SetData(v []PrincipalEntitlementValue) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *PrincipalEntitlementValuesList) GetLinks() ListLinks {
	if o == nil {
		var ret ListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValuesList) GetLinksOk() (*ListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PrincipalEntitlementValuesList) SetLinks(v ListLinks) {
	o.Links = v
}

func (o PrincipalEntitlementValuesList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalEntitlementValuesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalEntitlementValuesList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPrincipalEntitlementValuesList := _PrincipalEntitlementValuesList{}

	err = json.Unmarshal(data, &varPrincipalEntitlementValuesList)

	if err != nil {
		return err
	}

	*o = PrincipalEntitlementValuesList(varPrincipalEntitlementValuesList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalEntitlementValuesList struct {
	value *PrincipalEntitlementValuesList
	isSet bool
}

func (v NullablePrincipalEntitlementValuesList) Get() *PrincipalEntitlementValuesList {
	return v.value
}

func (v *NullablePrincipalEntitlementValuesList) Set(val *PrincipalEntitlementValuesList) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementValuesList) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementValuesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementValuesList(val *PrincipalEntitlementValuesList) *NullablePrincipalEntitlementValuesList {
	return &NullablePrincipalEntitlementValuesList{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementValuesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementValuesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
