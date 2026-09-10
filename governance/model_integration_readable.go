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

// checks if the IntegrationReadable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationReadable{}

// IntegrationReadable struct for IntegrationReadable
type IntegrationReadable struct {
	// The integration ID
	Id                   *string            `json:"id,omitempty" validate:"regexp=goi[0-9a-zA-Z]+"`
	Type                 *IntegrationType   `json:"type,omitempty"`
	Status               *IntegrationStatus `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IntegrationReadable IntegrationReadable

// NewIntegrationReadable instantiates a new IntegrationReadable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationReadable() *IntegrationReadable {
	this := IntegrationReadable{}
	return &this
}

// NewIntegrationReadableWithDefaults instantiates a new IntegrationReadable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationReadableWithDefaults() *IntegrationReadable {
	this := IntegrationReadable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationReadable) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationReadable) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationReadable) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationReadable) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegrationReadable) GetType() IntegrationType {
	if o == nil || IsNil(o.Type) {
		var ret IntegrationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationReadable) GetTypeOk() (*IntegrationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationReadable) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IntegrationType and assigns it to the Type field.
func (o *IntegrationReadable) SetType(v IntegrationType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntegrationReadable) GetStatus() IntegrationStatus {
	if o == nil || IsNil(o.Status) {
		var ret IntegrationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationReadable) GetStatusOk() (*IntegrationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationReadable) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IntegrationStatus and assigns it to the Status field.
func (o *IntegrationReadable) SetStatus(v IntegrationStatus) {
	o.Status = &v
}

func (o IntegrationReadable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationReadable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntegrationReadable) UnmarshalJSON(data []byte) (err error) {
	varIntegrationReadable := _IntegrationReadable{}

	err = json.Unmarshal(data, &varIntegrationReadable)

	if err != nil {
		return err
	}

	*o = IntegrationReadable(varIntegrationReadable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegrationReadable struct {
	value *IntegrationReadable
	isSet bool
}

func (v NullableIntegrationReadable) Get() *IntegrationReadable {
	return v.value
}

func (v *NullableIntegrationReadable) Set(val *IntegrationReadable) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationReadable) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationReadable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationReadable(val *IntegrationReadable) *NullableIntegrationReadable {
	return &NullableIntegrationReadable{value: val, isSet: true}
}

func (v NullableIntegrationReadable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationReadable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
