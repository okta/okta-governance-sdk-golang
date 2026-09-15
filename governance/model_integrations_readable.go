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

// checks if the IntegrationsReadable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationsReadable{}

// IntegrationsReadable struct for IntegrationsReadable
type IntegrationsReadable struct {
	// List of integration settings
	Data                 []IntegrationReadable `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IntegrationsReadable IntegrationsReadable

// NewIntegrationsReadable instantiates a new IntegrationsReadable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationsReadable() *IntegrationsReadable {
	this := IntegrationsReadable{}
	return &this
}

// NewIntegrationsReadableWithDefaults instantiates a new IntegrationsReadable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationsReadableWithDefaults() *IntegrationsReadable {
	this := IntegrationsReadable{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IntegrationsReadable) GetData() []IntegrationReadable {
	if o == nil || IsNil(o.Data) {
		var ret []IntegrationReadable
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationsReadable) GetDataOk() ([]IntegrationReadable, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IntegrationsReadable) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []IntegrationReadable and assigns it to the Data field.
func (o *IntegrationsReadable) SetData(v []IntegrationReadable) {
	o.Data = v
}

func (o IntegrationsReadable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationsReadable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntegrationsReadable) UnmarshalJSON(data []byte) (err error) {
	varIntegrationsReadable := _IntegrationsReadable{}

	err = json.Unmarshal(data, &varIntegrationsReadable)

	if err != nil {
		return err
	}

	*o = IntegrationsReadable(varIntegrationsReadable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegrationsReadable struct {
	value *IntegrationsReadable
	isSet bool
}

func (v NullableIntegrationsReadable) Get() *IntegrationsReadable {
	return v.value
}

func (v *NullableIntegrationsReadable) Set(val *IntegrationsReadable) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationsReadable) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationsReadable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationsReadable(val *IntegrationsReadable) *NullableIntegrationsReadable {
	return &NullableIntegrationsReadable{value: val, isSet: true}
}

func (v NullableIntegrationsReadable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationsReadable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
