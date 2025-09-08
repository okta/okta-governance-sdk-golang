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

// GrantedEntitlementValue struct for GrantedEntitlementValue
type GrantedEntitlementValue struct {
	// The `id` of an entitlement value
	Id string `json:"id"`
	// The display name for an entitlement value
	Name string `json:"name"`
	// The value of an entitlement property value
	ExternalValue string `json:"externalValue"`
	// A granted entitlement may not be effective if the same entitlement is granted by a higher priority additional grant
	Effective            *bool `json:"effective,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantedEntitlementValue GrantedEntitlementValue

// NewGrantedEntitlementValue instantiates a new GrantedEntitlementValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantedEntitlementValue(id string, name string, externalValue string) *GrantedEntitlementValue {
	this := GrantedEntitlementValue{}
	this.Id = id
	this.Name = name
	this.ExternalValue = externalValue
	return &this
}

// NewGrantedEntitlementValueWithDefaults instantiates a new GrantedEntitlementValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantedEntitlementValueWithDefaults() *GrantedEntitlementValue {
	this := GrantedEntitlementValue{}
	return &this
}

// GetId returns the Id field value
func (o *GrantedEntitlementValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlementValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GrantedEntitlementValue) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GrantedEntitlementValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlementValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GrantedEntitlementValue) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *GrantedEntitlementValue) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlementValue) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *GrantedEntitlementValue) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetEffective returns the Effective field value if set, zero value otherwise.
func (o *GrantedEntitlementValue) GetEffective() bool {
	if o == nil || o.Effective == nil {
		var ret bool
		return ret
	}
	return *o.Effective
}

// GetEffectiveOk returns a tuple with the Effective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantedEntitlementValue) GetEffectiveOk() (*bool, bool) {
	if o == nil || o.Effective == nil {
		return nil, false
	}
	return o.Effective, true
}

// HasEffective returns a boolean if a field has been set.
func (o *GrantedEntitlementValue) HasEffective() bool {
	if o != nil && o.Effective != nil {
		return true
	}

	return false
}

// SetEffective gets a reference to the given bool and assigns it to the Effective field.
func (o *GrantedEntitlementValue) SetEffective(v bool) {
	o.Effective = &v
}

func (o GrantedEntitlementValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if o.Effective != nil {
		toSerialize["effective"] = o.Effective
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GrantedEntitlementValue) UnmarshalJSON(bytes []byte) (err error) {
	varGrantedEntitlementValue := _GrantedEntitlementValue{}

	err = json.Unmarshal(bytes, &varGrantedEntitlementValue)
	if err == nil {
		*o = GrantedEntitlementValue(varGrantedEntitlementValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "effective")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGrantedEntitlementValue struct {
	value *GrantedEntitlementValue
	isSet bool
}

func (v NullableGrantedEntitlementValue) Get() *GrantedEntitlementValue {
	return v.value
}

func (v *NullableGrantedEntitlementValue) Set(val *GrantedEntitlementValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantedEntitlementValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantedEntitlementValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantedEntitlementValue(val *GrantedEntitlementValue) *NullableGrantedEntitlementValue {
	return &NullableGrantedEntitlementValue{value: val, isSet: true}
}

func (v NullableGrantedEntitlementValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantedEntitlementValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
