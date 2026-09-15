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

// checks if the IntegrationFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationFull{}

// IntegrationFull struct for IntegrationFull
type IntegrationFull struct {
	// The integration ID
	Id                   *string            `json:"id,omitempty" validate:"regexp=goi[0-9a-zA-Z]+"`
	Type                 *IntegrationType   `json:"type,omitempty"`
	Status               *IntegrationStatus `json:"status,omitempty"`
	Links                *IntegrationLinks  `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IntegrationFull IntegrationFull

// NewIntegrationFull instantiates a new IntegrationFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationFull() *IntegrationFull {
	this := IntegrationFull{}
	return &this
}

// NewIntegrationFullWithDefaults instantiates a new IntegrationFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationFullWithDefaults() *IntegrationFull {
	this := IntegrationFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationFull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationFull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationFull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationFull) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegrationFull) GetType() IntegrationType {
	if o == nil || IsNil(o.Type) {
		var ret IntegrationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationFull) GetTypeOk() (*IntegrationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationFull) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IntegrationType and assigns it to the Type field.
func (o *IntegrationFull) SetType(v IntegrationType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntegrationFull) GetStatus() IntegrationStatus {
	if o == nil || IsNil(o.Status) {
		var ret IntegrationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationFull) GetStatusOk() (*IntegrationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationFull) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IntegrationStatus and assigns it to the Status field.
func (o *IntegrationFull) SetStatus(v IntegrationStatus) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IntegrationFull) GetLinks() IntegrationLinks {
	if o == nil || IsNil(o.Links) {
		var ret IntegrationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationFull) GetLinksOk() (*IntegrationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IntegrationFull) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given IntegrationLinks and assigns it to the Links field.
func (o *IntegrationFull) SetLinks(v IntegrationLinks) {
	o.Links = &v
}

func (o IntegrationFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationFull) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntegrationFull) UnmarshalJSON(data []byte) (err error) {
	varIntegrationFull := _IntegrationFull{}

	err = json.Unmarshal(data, &varIntegrationFull)

	if err != nil {
		return err
	}

	*o = IntegrationFull(varIntegrationFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegrationFull struct {
	value *IntegrationFull
	isSet bool
}

func (v NullableIntegrationFull) Get() *IntegrationFull {
	return v.value
}

func (v *NullableIntegrationFull) Set(val *IntegrationFull) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationFull) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationFull(val *IntegrationFull) *NullableIntegrationFull {
	return &NullableIntegrationFull{value: val, isSet: true}
}

func (v NullableIntegrationFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
