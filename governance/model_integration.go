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

// checks if the Integration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Integration{}

// Integration struct for Integration
type Integration struct {
	// Integration ID
	Id                   *string            `json:"id,omitempty"`
	Type                 *IntegrationType   `json:"type,omitempty"`
	Status               *IntegrationStatus `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Integration Integration

// NewIntegration instantiates a new Integration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegration() *Integration {
	this := Integration{}
	return &this
}

// NewIntegrationWithDefaults instantiates a new Integration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationWithDefaults() *Integration {
	this := Integration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Integration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Integration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Integration) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Integration) GetType() IntegrationType {
	if o == nil || IsNil(o.Type) {
		var ret IntegrationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetTypeOk() (*IntegrationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Integration) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IntegrationType and assigns it to the Type field.
func (o *Integration) SetType(v IntegrationType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Integration) GetStatus() IntegrationStatus {
	if o == nil || IsNil(o.Status) {
		var ret IntegrationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetStatusOk() (*IntegrationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Integration) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IntegrationStatus and assigns it to the Status field.
func (o *Integration) SetStatus(v IntegrationStatus) {
	o.Status = &v
}

func (o Integration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Integration) ToMap() (map[string]interface{}, error) {
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

func (o *Integration) UnmarshalJSON(data []byte) (err error) {
	varIntegration := _Integration{}

	err = json.Unmarshal(data, &varIntegration)

	if err != nil {
		return err
	}

	*o = Integration(varIntegration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegration struct {
	value *Integration
	isSet bool
}

func (v NullableIntegration) Get() *Integration {
	return v.value
}

func (v *NullableIntegration) Set(val *Integration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegration(val *Integration) *NullableIntegration {
	return &NullableIntegration{value: val, isSet: true}
}

func (v NullableIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
