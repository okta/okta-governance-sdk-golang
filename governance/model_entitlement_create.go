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

// checks if the EntitlementCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementCreate{}

// EntitlementCreate The properties expected in an initial add entitlement
type EntitlementCreate struct {
	Parent *TargetResource                      `json:"parent,omitempty"`
	Values []EntitlementValueWritableProperties `json:"values,omitempty"`
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

type _EntitlementCreate EntitlementCreate

// NewEntitlementCreate instantiates a new EntitlementCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementCreate(name string, externalValue string, multiValue bool, dataType EntitlementPropertyDatatype) *EntitlementCreate {
	this := EntitlementCreate{}
	this.Name = name
	this.ExternalValue = externalValue
	this.MultiValue = multiValue
	this.DataType = dataType
	return &this
}

// NewEntitlementCreateWithDefaults instantiates a new EntitlementCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementCreateWithDefaults() *EntitlementCreate {
	this := EntitlementCreate{}
	return &this
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *EntitlementCreate) GetParent() TargetResource {
	if o == nil || IsNil(o.Parent) {
		var ret TargetResource
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementCreate) GetParentOk() (*TargetResource, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *EntitlementCreate) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given TargetResource and assigns it to the Parent field.
func (o *EntitlementCreate) SetParent(v TargetResource) {
	o.Parent = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementCreate) GetValues() []EntitlementValueWritableProperties {
	if o == nil || IsNil(o.Values) {
		var ret []EntitlementValueWritableProperties
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementCreate) GetValuesOk() ([]EntitlementValueWritableProperties, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementCreate) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementValueWritableProperties and assigns it to the Values field.
func (o *EntitlementCreate) SetValues(v []EntitlementValueWritableProperties) {
	o.Values = v
}

// GetName returns the Name field value
func (o *EntitlementCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementCreate) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *EntitlementCreate) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementCreate) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *EntitlementCreate) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementCreate) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value
func (o *EntitlementCreate) GetMultiValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementCreate) GetMultiValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiValue, true
}

// SetMultiValue sets field value
func (o *EntitlementCreate) SetMultiValue(v bool) {
	o.MultiValue = v
}

// GetDataType returns the DataType field value
func (o *EntitlementCreate) GetDataType() EntitlementPropertyDatatype {
	if o == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *EntitlementCreate) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *EntitlementCreate) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = v
}

func (o EntitlementCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
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

func (o *EntitlementCreate) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementCreate := _EntitlementCreate{}

	err = json.Unmarshal(data, &varEntitlementCreate)

	if err != nil {
		return err
	}

	*o = EntitlementCreate(varEntitlementCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parent")
		delete(additionalProperties, "values")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "description")
		delete(additionalProperties, "multiValue")
		delete(additionalProperties, "dataType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementCreate struct {
	value *EntitlementCreate
	isSet bool
}

func (v NullableEntitlementCreate) Get() *EntitlementCreate {
	return v.value
}

func (v *NullableEntitlementCreate) Set(val *EntitlementCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementCreate(val *EntitlementCreate) *NullableEntitlementCreate {
	return &NullableEntitlementCreate{value: val, isSet: true}
}

func (v NullableEntitlementCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
