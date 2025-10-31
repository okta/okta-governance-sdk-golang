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

// checks if the Criteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Criteria{}

// Criteria struct for Criteria
type Criteria struct {
	// Name of the criteria
	Name *string `json:"name,omitempty"`
	// Attribute that the criteria applies to
	Attribute *string `json:"attribute,omitempty"`
	// Operation performed on the criteria value
	Operation            *string        `json:"operation,omitempty"`
	Value                *CriteriaValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Criteria Criteria

// NewCriteria instantiates a new Criteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCriteria() *Criteria {
	this := Criteria{}
	return &this
}

// NewCriteriaWithDefaults instantiates a new Criteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCriteriaWithDefaults() *Criteria {
	this := Criteria{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Criteria) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Criteria) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Criteria) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Criteria) SetName(v string) {
	o.Name = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *Criteria) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Criteria) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *Criteria) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *Criteria) SetAttribute(v string) {
	o.Attribute = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *Criteria) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Criteria) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *Criteria) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *Criteria) SetOperation(v string) {
	o.Operation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Criteria) GetValue() CriteriaValue {
	if o == nil || IsNil(o.Value) {
		var ret CriteriaValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Criteria) GetValueOk() (*CriteriaValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Criteria) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CriteriaValue and assigns it to the Value field.
func (o *Criteria) SetValue(v CriteriaValue) {
	o.Value = &v
}

func (o Criteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Criteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Criteria) UnmarshalJSON(data []byte) (err error) {
	varCriteria := _Criteria{}

	err = json.Unmarshal(data, &varCriteria)

	if err != nil {
		return err
	}

	*o = Criteria(varCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCriteria struct {
	value *Criteria
	isSet bool
}

func (v NullableCriteria) Get() *Criteria {
	return v.value
}

func (v *NullableCriteria) Set(val *Criteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCriteria(val *Criteria) *NullableCriteria {
	return &NullableCriteria{value: val, isSet: true}
}

func (v NullableCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
