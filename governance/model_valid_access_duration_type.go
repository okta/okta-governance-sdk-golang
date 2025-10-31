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
)

// checks if the ValidAccessDurationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidAccessDurationType{}

// ValidAccessDurationType struct for ValidAccessDurationType
type ValidAccessDurationType struct {
	Type                 *AccessDurationType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidAccessDurationType ValidAccessDurationType

// NewValidAccessDurationType instantiates a new ValidAccessDurationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidAccessDurationType() *ValidAccessDurationType {
	this := ValidAccessDurationType{}
	return &this
}

// NewValidAccessDurationTypeWithDefaults instantiates a new ValidAccessDurationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidAccessDurationTypeWithDefaults() *ValidAccessDurationType {
	this := ValidAccessDurationType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ValidAccessDurationType) GetType() AccessDurationType {
	if o == nil || IsNil(o.Type) {
		var ret AccessDurationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidAccessDurationType) GetTypeOk() (*AccessDurationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ValidAccessDurationType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AccessDurationType and assigns it to the Type field.
func (o *ValidAccessDurationType) SetType(v AccessDurationType) {
	o.Type = &v
}

func (o ValidAccessDurationType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidAccessDurationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidAccessDurationType) UnmarshalJSON(data []byte) (err error) {
	varValidAccessDurationType := _ValidAccessDurationType{}

	err = json.Unmarshal(data, &varValidAccessDurationType)

	if err != nil {
		return err
	}

	*o = ValidAccessDurationType(varValidAccessDurationType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidAccessDurationType struct {
	value *ValidAccessDurationType
	isSet bool
}

func (v NullableValidAccessDurationType) Get() *ValidAccessDurationType {
	return v.value
}

func (v *NullableValidAccessDurationType) Set(val *ValidAccessDurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableValidAccessDurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableValidAccessDurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidAccessDurationType(val *ValidAccessDurationType) *NullableValidAccessDurationType {
	return &NullableValidAccessDurationType{value: val, isSet: true}
}

func (v NullableValidAccessDurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidAccessDurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
