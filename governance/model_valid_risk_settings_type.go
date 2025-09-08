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

// ValidRiskSettingsType struct for ValidRiskSettingsType
type ValidRiskSettingsType struct {
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidRiskSettingsType ValidRiskSettingsType

// NewValidRiskSettingsType instantiates a new ValidRiskSettingsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidRiskSettingsType() *ValidRiskSettingsType {
	this := ValidRiskSettingsType{}
	return &this
}

// NewValidRiskSettingsTypeWithDefaults instantiates a new ValidRiskSettingsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidRiskSettingsTypeWithDefaults() *ValidRiskSettingsType {
	this := ValidRiskSettingsType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ValidRiskSettingsType) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidRiskSettingsType) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ValidRiskSettingsType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ValidRiskSettingsType) SetType(v string) {
	o.Type = &v
}

func (o ValidRiskSettingsType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ValidRiskSettingsType) UnmarshalJSON(bytes []byte) (err error) {
	varValidRiskSettingsType := _ValidRiskSettingsType{}

	err = json.Unmarshal(bytes, &varValidRiskSettingsType)
	if err == nil {
		*o = ValidRiskSettingsType(varValidRiskSettingsType)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableValidRiskSettingsType struct {
	value *ValidRiskSettingsType
	isSet bool
}

func (v NullableValidRiskSettingsType) Get() *ValidRiskSettingsType {
	return v.value
}

func (v *NullableValidRiskSettingsType) Set(val *ValidRiskSettingsType) {
	v.value = val
	v.isSet = true
}

func (v NullableValidRiskSettingsType) IsSet() bool {
	return v.isSet
}

func (v *NullableValidRiskSettingsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidRiskSettingsType(val *ValidRiskSettingsType) *NullableValidRiskSettingsType {
	return &NullableValidRiskSettingsType{value: val, isSet: true}
}

func (v NullableValidRiskSettingsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidRiskSettingsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
