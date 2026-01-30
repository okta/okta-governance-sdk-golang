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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ResourceDefaultAccessScopeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceDefaultAccessScopeSettings{}

// ResourceDefaultAccessScopeSettings Default Access scope settings associated with the condition's resource.
type ResourceDefaultAccessScopeSettings struct {
	Type string `json:"type"`
}

type _ResourceDefaultAccessScopeSettings ResourceDefaultAccessScopeSettings

// NewResourceDefaultAccessScopeSettings instantiates a new ResourceDefaultAccessScopeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceDefaultAccessScopeSettings(type_ string) *ResourceDefaultAccessScopeSettings {
	this := ResourceDefaultAccessScopeSettings{}
	this.Type = type_
	return &this
}

// NewResourceDefaultAccessScopeSettingsWithDefaults instantiates a new ResourceDefaultAccessScopeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceDefaultAccessScopeSettingsWithDefaults() *ResourceDefaultAccessScopeSettings {
	this := ResourceDefaultAccessScopeSettings{}
	return &this
}

// GetType returns the Type field value
func (o *ResourceDefaultAccessScopeSettings) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceDefaultAccessScopeSettings) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceDefaultAccessScopeSettings) SetType(v string) {
	o.Type = v
}

func (o ResourceDefaultAccessScopeSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceDefaultAccessScopeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ResourceDefaultAccessScopeSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varResourceDefaultAccessScopeSettings := _ResourceDefaultAccessScopeSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceDefaultAccessScopeSettings)

	if err != nil {
		return err
	}

	*o = ResourceDefaultAccessScopeSettings(varResourceDefaultAccessScopeSettings)

	return err
}

type NullableResourceDefaultAccessScopeSettings struct {
	value *ResourceDefaultAccessScopeSettings
	isSet bool
}

func (v NullableResourceDefaultAccessScopeSettings) Get() *ResourceDefaultAccessScopeSettings {
	return v.value
}

func (v *NullableResourceDefaultAccessScopeSettings) Set(val *ResourceDefaultAccessScopeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDefaultAccessScopeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDefaultAccessScopeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDefaultAccessScopeSettings(val *ResourceDefaultAccessScopeSettings) *NullableResourceDefaultAccessScopeSettings {
	return &NullableResourceDefaultAccessScopeSettings{value: val, isSet: true}
}

func (v NullableResourceDefaultAccessScopeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDefaultAccessScopeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
