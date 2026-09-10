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

// checks if the GrantPrincipalAccessResponseLinksPrincipalAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPrincipalAccessResponseLinksPrincipalAccess{}

// GrantPrincipalAccessResponseLinksPrincipalAccess Link to retrieve the principal's access
type GrantPrincipalAccessResponseLinksPrincipalAccess struct {
	// The URL to the principal access GET endpoint
	Href                 string `json:"href"`
	AdditionalProperties map[string]interface{}
}

type _GrantPrincipalAccessResponseLinksPrincipalAccess GrantPrincipalAccessResponseLinksPrincipalAccess

// NewGrantPrincipalAccessResponseLinksPrincipalAccess instantiates a new GrantPrincipalAccessResponseLinksPrincipalAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPrincipalAccessResponseLinksPrincipalAccess(href string) *GrantPrincipalAccessResponseLinksPrincipalAccess {
	this := GrantPrincipalAccessResponseLinksPrincipalAccess{}
	this.Href = href
	return &this
}

// NewGrantPrincipalAccessResponseLinksPrincipalAccessWithDefaults instantiates a new GrantPrincipalAccessResponseLinksPrincipalAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPrincipalAccessResponseLinksPrincipalAccessWithDefaults() *GrantPrincipalAccessResponseLinksPrincipalAccess {
	this := GrantPrincipalAccessResponseLinksPrincipalAccess{}
	return &this
}

// GetHref returns the Href field value
func (o *GrantPrincipalAccessResponseLinksPrincipalAccess) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessResponseLinksPrincipalAccess) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *GrantPrincipalAccessResponseLinksPrincipalAccess) SetHref(v string) {
	o.Href = v
}

func (o GrantPrincipalAccessResponseLinksPrincipalAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPrincipalAccessResponseLinksPrincipalAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPrincipalAccessResponseLinksPrincipalAccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
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

	varGrantPrincipalAccessResponseLinksPrincipalAccess := _GrantPrincipalAccessResponseLinksPrincipalAccess{}

	err = json.Unmarshal(data, &varGrantPrincipalAccessResponseLinksPrincipalAccess)

	if err != nil {
		return err
	}

	*o = GrantPrincipalAccessResponseLinksPrincipalAccess(varGrantPrincipalAccessResponseLinksPrincipalAccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPrincipalAccessResponseLinksPrincipalAccess struct {
	value *GrantPrincipalAccessResponseLinksPrincipalAccess
	isSet bool
}

func (v NullableGrantPrincipalAccessResponseLinksPrincipalAccess) Get() *GrantPrincipalAccessResponseLinksPrincipalAccess {
	return v.value
}

func (v *NullableGrantPrincipalAccessResponseLinksPrincipalAccess) Set(val *GrantPrincipalAccessResponseLinksPrincipalAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalAccessResponseLinksPrincipalAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalAccessResponseLinksPrincipalAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalAccessResponseLinksPrincipalAccess(val *GrantPrincipalAccessResponseLinksPrincipalAccess) *NullableGrantPrincipalAccessResponseLinksPrincipalAccess {
	return &NullableGrantPrincipalAccessResponseLinksPrincipalAccess{value: val, isSet: true}
}

func (v NullableGrantPrincipalAccessResponseLinksPrincipalAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalAccessResponseLinksPrincipalAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
