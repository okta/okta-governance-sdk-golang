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

// checks if the EntitlementValuesPatchOperationValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementValuesPatchOperationValue{}

// EntitlementValuesPatchOperationValue struct for EntitlementValuesPatchOperationValue
type EntitlementValuesPatchOperationValue struct {
	// The display name for an entitlement value
	Name *string `json:"name,omitempty"`
	// The value of an entitlement property value
	ExternalValue *string `json:"externalValue,omitempty"`
	// The description of an entitlement value
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValuesPatchOperationValue EntitlementValuesPatchOperationValue

// NewEntitlementValuesPatchOperationValue instantiates a new EntitlementValuesPatchOperationValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValuesPatchOperationValue() *EntitlementValuesPatchOperationValue {
	this := EntitlementValuesPatchOperationValue{}
	return &this
}

// NewEntitlementValuesPatchOperationValueWithDefaults instantiates a new EntitlementValuesPatchOperationValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValuesPatchOperationValueWithDefaults() *EntitlementValuesPatchOperationValue {
	this := EntitlementValuesPatchOperationValue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementValuesPatchOperationValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesPatchOperationValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementValuesPatchOperationValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementValuesPatchOperationValue) SetName(v string) {
	o.Name = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *EntitlementValuesPatchOperationValue) GetExternalValue() string {
	if o == nil || IsNil(o.ExternalValue) {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesPatchOperationValue) GetExternalValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalValue) {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementValuesPatchOperationValue) HasExternalValue() bool {
	if o != nil && !IsNil(o.ExternalValue) {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *EntitlementValuesPatchOperationValue) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementValuesPatchOperationValue) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesPatchOperationValue) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementValuesPatchOperationValue) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementValuesPatchOperationValue) SetDescription(v string) {
	o.Description = &v
}

func (o EntitlementValuesPatchOperationValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementValuesPatchOperationValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExternalValue) {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementValuesPatchOperationValue) UnmarshalJSON(data []byte) (err error) {
	varEntitlementValuesPatchOperationValue := _EntitlementValuesPatchOperationValue{}

	err = json.Unmarshal(data, &varEntitlementValuesPatchOperationValue)

	if err != nil {
		return err
	}

	*o = EntitlementValuesPatchOperationValue(varEntitlementValuesPatchOperationValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementValuesPatchOperationValue struct {
	value *EntitlementValuesPatchOperationValue
	isSet bool
}

func (v NullableEntitlementValuesPatchOperationValue) Get() *EntitlementValuesPatchOperationValue {
	return v.value
}

func (v *NullableEntitlementValuesPatchOperationValue) Set(val *EntitlementValuesPatchOperationValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValuesPatchOperationValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValuesPatchOperationValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValuesPatchOperationValue(val *EntitlementValuesPatchOperationValue) *NullableEntitlementValuesPatchOperationValue {
	return &NullableEntitlementValuesPatchOperationValue{value: val, isSet: true}
}

func (v NullableEntitlementValuesPatchOperationValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValuesPatchOperationValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
