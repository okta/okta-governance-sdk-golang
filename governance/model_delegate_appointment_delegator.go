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

// checks if the DelegateAppointmentDelegator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegateAppointmentDelegator{}

// DelegateAppointmentDelegator The principal who delegated their governance duties
type DelegateAppointmentDelegator struct {
	// The Okta user `id`
	ExternalId           string        `json:"externalId" validate:"regexp=00u[0-9a-zA-Z]+"`
	Type                 PrincipalType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _DelegateAppointmentDelegator DelegateAppointmentDelegator

// NewDelegateAppointmentDelegator instantiates a new DelegateAppointmentDelegator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegateAppointmentDelegator(externalId string, type_ PrincipalType) *DelegateAppointmentDelegator {
	this := DelegateAppointmentDelegator{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewDelegateAppointmentDelegatorWithDefaults instantiates a new DelegateAppointmentDelegator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegateAppointmentDelegatorWithDefaults() *DelegateAppointmentDelegator {
	this := DelegateAppointmentDelegator{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *DelegateAppointmentDelegator) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *DelegateAppointmentDelegator) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *DelegateAppointmentDelegator) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *DelegateAppointmentDelegator) GetType() PrincipalType {
	if o == nil {
		var ret PrincipalType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DelegateAppointmentDelegator) GetTypeOk() (*PrincipalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DelegateAppointmentDelegator) SetType(v PrincipalType) {
	o.Type = v
}

func (o DelegateAppointmentDelegator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegateAppointmentDelegator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DelegateAppointmentDelegator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
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

	varDelegateAppointmentDelegator := _DelegateAppointmentDelegator{}

	err = json.Unmarshal(data, &varDelegateAppointmentDelegator)

	if err != nil {
		return err
	}

	*o = DelegateAppointmentDelegator(varDelegateAppointmentDelegator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDelegateAppointmentDelegator struct {
	value *DelegateAppointmentDelegator
	isSet bool
}

func (v NullableDelegateAppointmentDelegator) Get() *DelegateAppointmentDelegator {
	return v.value
}

func (v *NullableDelegateAppointmentDelegator) Set(val *DelegateAppointmentDelegator) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegateAppointmentDelegator) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegateAppointmentDelegator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegateAppointmentDelegator(val *DelegateAppointmentDelegator) *NullableDelegateAppointmentDelegator {
	return &NullableDelegateAppointmentDelegator{value: val, isSet: true}
}

func (v NullableDelegateAppointmentDelegator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegateAppointmentDelegator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
