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

// checks if the GrantPrincipalAccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPrincipalAccessResponse{}

// GrantPrincipalAccessResponse Response for a successful grant principal access request
type GrantPrincipalAccessResponse struct {
	Links                GrantPrincipalAccessResponseLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _GrantPrincipalAccessResponse GrantPrincipalAccessResponse

// NewGrantPrincipalAccessResponse instantiates a new GrantPrincipalAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPrincipalAccessResponse(links GrantPrincipalAccessResponseLinks) *GrantPrincipalAccessResponse {
	this := GrantPrincipalAccessResponse{}
	this.Links = links
	return &this
}

// NewGrantPrincipalAccessResponseWithDefaults instantiates a new GrantPrincipalAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPrincipalAccessResponseWithDefaults() *GrantPrincipalAccessResponse {
	this := GrantPrincipalAccessResponse{}
	return &this
}

// GetLinks returns the Links field value
func (o *GrantPrincipalAccessResponse) GetLinks() GrantPrincipalAccessResponseLinks {
	if o == nil {
		var ret GrantPrincipalAccessResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessResponse) GetLinksOk() (*GrantPrincipalAccessResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GrantPrincipalAccessResponse) SetLinks(v GrantPrincipalAccessResponseLinks) {
	o.Links = v
}

func (o GrantPrincipalAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPrincipalAccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPrincipalAccessResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGrantPrincipalAccessResponse := _GrantPrincipalAccessResponse{}

	err = json.Unmarshal(data, &varGrantPrincipalAccessResponse)

	if err != nil {
		return err
	}

	*o = GrantPrincipalAccessResponse(varGrantPrincipalAccessResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPrincipalAccessResponse struct {
	value *GrantPrincipalAccessResponse
	isSet bool
}

func (v NullableGrantPrincipalAccessResponse) Get() *GrantPrincipalAccessResponse {
	return v.value
}

func (v *NullableGrantPrincipalAccessResponse) Set(val *GrantPrincipalAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalAccessResponse(val *GrantPrincipalAccessResponse) *NullableGrantPrincipalAccessResponse {
	return &NullableGrantPrincipalAccessResponse{value: val, isSet: true}
}

func (v NullableGrantPrincipalAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
