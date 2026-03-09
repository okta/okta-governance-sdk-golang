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

// checks if the EntitlementWritableProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementWritableProperties{}

// EntitlementWritableProperties struct for EntitlementWritableProperties
type EntitlementWritableProperties struct {
	// The display name for an entitlement property
	Name string `json:"name"`
	// The value of an entitlement property
	ExternalValue string `json:"externalValue"`
	// The description of an entitlement property
	Description *string `json:"description,omitempty"`
	// Indicate if the entitlement property can hold multiple values. If this property is `true`, then the `dataType` property is set to  `array`.
	MultiValue           bool                        `json:"multiValue"`
	DataType             EntitlementPropertyDatatype `json:"dataType"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementWritableProperties EntitlementWritableProperties

// NewEntitlementWritableProperties instantiates a new EntitlementWritableProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementWritableProperties(name string, externalValue string, multiValue bool, dataType EntitlementPropertyDatatype) *EntitlementWritableProperties {
	this := EntitlementWritableProperties{}
	this.Name = name
	this.ExternalValue = externalValue
	this.MultiValue = multiValue
	this.DataType = dataType
	return &this
}

// NewEntitlementWritablePropertiesWithDefaults instantiates a new EntitlementWritableProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWritablePropertiesWithDefaults() *EntitlementWritableProperties {
	this := EntitlementWritableProperties{}
	return &this
}

// GetName returns the Name field value
func (o *EntitlementWritableProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementWritableProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementWritableProperties) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *EntitlementWritableProperties) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementWritableProperties) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *EntitlementWritableProperties) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementWritableProperties) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWritableProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementWritableProperties) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementWritableProperties) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value
func (o *EntitlementWritableProperties) GetMultiValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementWritableProperties) GetMultiValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiValue, true
}

// SetMultiValue sets field value
func (o *EntitlementWritableProperties) SetMultiValue(v bool) {
	o.MultiValue = v
}

// GetDataType returns the DataType field value
func (o *EntitlementWritableProperties) GetDataType() EntitlementPropertyDatatype {
	if o == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *EntitlementWritableProperties) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *EntitlementWritableProperties) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = v
}

func (o EntitlementWritableProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementWritableProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["externalValue"] = o.ExternalValue
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["multiValue"] = o.MultiValue
	toSerialize["dataType"] = o.DataType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementWritableProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"externalValue",
		"multiValue",
		"dataType",
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

	varEntitlementWritableProperties := _EntitlementWritableProperties{}

	err = json.Unmarshal(data, &varEntitlementWritableProperties)

	if err != nil {
		return err
	}

	*o = EntitlementWritableProperties(varEntitlementWritableProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "description")
		delete(additionalProperties, "multiValue")
		delete(additionalProperties, "dataType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementWritableProperties struct {
	value *EntitlementWritableProperties
	isSet bool
}

func (v NullableEntitlementWritableProperties) Get() *EntitlementWritableProperties {
	return v.value
}

func (v *NullableEntitlementWritableProperties) Set(val *EntitlementWritableProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementWritableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementWritableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementWritableProperties(val *EntitlementWritableProperties) *NullableEntitlementWritableProperties {
	return &NullableEntitlementWritableProperties{value: val, isSet: true}
}

func (v NullableEntitlementWritableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementWritableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
