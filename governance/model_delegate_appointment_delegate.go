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

// DelegateAppointmentDelegate The appointed principal who performs delegated duties
type DelegateAppointmentDelegate struct {
	// The Okta user `id`
	ExternalId           string        `json:"externalId" validate:"regexp=00u[0-9a-zA-Z]+"`
	Type                 PrincipalType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _DelegateAppointmentDelegate DelegateAppointmentDelegate

// NewDelegateAppointmentDelegate instantiates a new DelegateAppointmentDelegate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegateAppointmentDelegate(externalId string, type_ PrincipalType) *DelegateAppointmentDelegate {
	this := DelegateAppointmentDelegate{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewDelegateAppointmentDelegateWithDefaults instantiates a new DelegateAppointmentDelegate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegateAppointmentDelegateWithDefaults() *DelegateAppointmentDelegate {
	this := DelegateAppointmentDelegate{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *DelegateAppointmentDelegate) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *DelegateAppointmentDelegate) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *DelegateAppointmentDelegate) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *DelegateAppointmentDelegate) GetType() PrincipalType {
	if o == nil {
		var ret PrincipalType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DelegateAppointmentDelegate) GetTypeOk() (*PrincipalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DelegateAppointmentDelegate) SetType(v PrincipalType) {
	o.Type = v
}

func (o DelegateAppointmentDelegate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["externalId"] = o.ExternalId
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DelegateAppointmentDelegate) UnmarshalJSON(bytes []byte) (err error) {
	varDelegateAppointmentDelegate := _DelegateAppointmentDelegate{}

	err = json.Unmarshal(bytes, &varDelegateAppointmentDelegate)
	if err == nil {
		*o = DelegateAppointmentDelegate(varDelegateAppointmentDelegate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDelegateAppointmentDelegate struct {
	value *DelegateAppointmentDelegate
	isSet bool
}

func (v NullableDelegateAppointmentDelegate) Get() *DelegateAppointmentDelegate {
	return v.value
}

func (v *NullableDelegateAppointmentDelegate) Set(val *DelegateAppointmentDelegate) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegateAppointmentDelegate) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegateAppointmentDelegate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegateAppointmentDelegate(val *DelegateAppointmentDelegate) *NullableDelegateAppointmentDelegate {
	return &NullableDelegateAppointmentDelegate{value: val, isSet: true}
}

func (v NullableDelegateAppointmentDelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegateAppointmentDelegate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
