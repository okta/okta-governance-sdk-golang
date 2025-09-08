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

// EntitlementChangedFull All entitlements with the property values
type EntitlementChangedFull struct {
	Values []EntitlementValueChanged `json:"values,omitempty"`
	// The `id` property of an entitlement
	Id string `json:"id"`
	// The display name for an entitlement property
	Name *string `json:"name,omitempty"`
	// The value of an entitlement property
	ExternalValue *string `json:"externalValue,omitempty"`
	// The description of an entitlement property
	Description *string `json:"description,omitempty"`
	// The property that determines if the entitlement property can hold multiple values. If this is set to true, the data type is replaced with an array.
	MultiValue *bool `json:"multiValue,omitempty"`
	// The property that determines if the entitlement property is a required attribute
	Required             *bool                        `json:"required,omitempty"`
	DataType             *EntitlementPropertyDatatype `json:"dataType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementChangedFull EntitlementChangedFull

// NewEntitlementChangedFull instantiates a new EntitlementChangedFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementChangedFull(id string) *EntitlementChangedFull {
	this := EntitlementChangedFull{}
	this.Id = id
	return &this
}

// NewEntitlementChangedFullWithDefaults instantiates a new EntitlementChangedFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementChangedFullWithDefaults() *EntitlementChangedFull {
	this := EntitlementChangedFull{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementChangedFull) GetValues() []EntitlementValueChanged {
	if o == nil || o.Values == nil {
		var ret []EntitlementValueChanged
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetValuesOk() ([]EntitlementValueChanged, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementChangedFull) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementValueChanged and assigns it to the Values field.
func (o *EntitlementChangedFull) SetValues(v []EntitlementValueChanged) {
	o.Values = v
}

// GetId returns the Id field value
func (o *EntitlementChangedFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementChangedFull) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementChangedFull) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementChangedFull) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementChangedFull) SetName(v string) {
	o.Name = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *EntitlementChangedFull) GetExternalValue() string {
	if o == nil || o.ExternalValue == nil {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetExternalValueOk() (*string, bool) {
	if o == nil || o.ExternalValue == nil {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementChangedFull) HasExternalValue() bool {
	if o != nil && o.ExternalValue != nil {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *EntitlementChangedFull) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementChangedFull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementChangedFull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementChangedFull) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value if set, zero value otherwise.
func (o *EntitlementChangedFull) GetMultiValue() bool {
	if o == nil || o.MultiValue == nil {
		var ret bool
		return ret
	}
	return *o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetMultiValueOk() (*bool, bool) {
	if o == nil || o.MultiValue == nil {
		return nil, false
	}
	return o.MultiValue, true
}

// HasMultiValue returns a boolean if a field has been set.
func (o *EntitlementChangedFull) HasMultiValue() bool {
	if o != nil && o.MultiValue != nil {
		return true
	}

	return false
}

// SetMultiValue gets a reference to the given bool and assigns it to the MultiValue field.
func (o *EntitlementChangedFull) SetMultiValue(v bool) {
	o.MultiValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *EntitlementChangedFull) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EntitlementChangedFull) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *EntitlementChangedFull) SetRequired(v bool) {
	o.Required = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *EntitlementChangedFull) GetDataType() EntitlementPropertyDatatype {
	if o == nil || o.DataType == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementChangedFull) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *EntitlementChangedFull) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given EntitlementPropertyDatatype and assigns it to the DataType field.
func (o *EntitlementChangedFull) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = &v
}

func (o EntitlementChangedFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExternalValue != nil {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.MultiValue != nil {
		toSerialize["multiValue"] = o.MultiValue
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementChangedFull) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementChangedFull := _EntitlementChangedFull{}

	err = json.Unmarshal(bytes, &varEntitlementChangedFull)
	if err == nil {
		*o = EntitlementChangedFull(varEntitlementChangedFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "values")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "description")
		delete(additionalProperties, "multiValue")
		delete(additionalProperties, "required")
		delete(additionalProperties, "dataType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementChangedFull struct {
	value *EntitlementChangedFull
	isSet bool
}

func (v NullableEntitlementChangedFull) Get() *EntitlementChangedFull {
	return v.value
}

func (v *NullableEntitlementChangedFull) Set(val *EntitlementChangedFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementChangedFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementChangedFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementChangedFull(val *EntitlementChangedFull) *NullableEntitlementChangedFull {
	return &NullableEntitlementChangedFull{value: val, isSet: true}
}

func (v NullableEntitlementChangedFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementChangedFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
