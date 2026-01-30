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

// checks if the EntitlementValueWritableProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementValueWritableProperties{}

// EntitlementValueWritableProperties struct for EntitlementValueWritableProperties
type EntitlementValueWritableProperties struct {
	// The display name for an entitlement value
	Name string `json:"name"`
	// The value of an entitlement property value
	ExternalValue string `json:"externalValue"`
	// The description of an entitlement value
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValueWritableProperties EntitlementValueWritableProperties

// NewEntitlementValueWritableProperties instantiates a new EntitlementValueWritableProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValueWritableProperties(name string, externalValue string) *EntitlementValueWritableProperties {
	this := EntitlementValueWritableProperties{}
	this.Name = name
	this.ExternalValue = externalValue
	return &this
}

// NewEntitlementValueWritablePropertiesWithDefaults instantiates a new EntitlementValueWritableProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueWritablePropertiesWithDefaults() *EntitlementValueWritableProperties {
	this := EntitlementValueWritableProperties{}
	return &this
}

// GetName returns the Name field value
func (o *EntitlementValueWritableProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWritableProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementValueWritableProperties) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *EntitlementValueWritableProperties) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWritableProperties) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *EntitlementValueWritableProperties) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementValueWritableProperties) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueWritableProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementValueWritableProperties) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementValueWritableProperties) SetDescription(v string) {
	o.Description = &v
}

func (o EntitlementValueWritableProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementValueWritableProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["externalValue"] = o.ExternalValue
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementValueWritableProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"externalValue",
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

	varEntitlementValueWritableProperties := _EntitlementValueWritableProperties{}

	err = json.Unmarshal(data, &varEntitlementValueWritableProperties)

	if err != nil {
		return err
	}

	*o = EntitlementValueWritableProperties(varEntitlementValueWritableProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementValueWritableProperties struct {
	value *EntitlementValueWritableProperties
	isSet bool
}

func (v NullableEntitlementValueWritableProperties) Get() *EntitlementValueWritableProperties {
	return v.value
}

func (v *NullableEntitlementValueWritableProperties) Set(val *EntitlementValueWritableProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValueWritableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValueWritableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValueWritableProperties(val *EntitlementValueWritableProperties) *NullableEntitlementValueWritableProperties {
	return &NullableEntitlementValueWritableProperties{value: val, isSet: true}
}

func (v NullableEntitlementValueWritableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValueWritableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
