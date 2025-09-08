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

// EntitlementPropertyFull Attributes related to Entitlement
type EntitlementPropertyFull struct {
	// The `id` property of an entitlement
	Id *string `json:"id,omitempty"`
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

type _EntitlementPropertyFull EntitlementPropertyFull

// NewEntitlementPropertyFull instantiates a new EntitlementPropertyFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementPropertyFull() *EntitlementPropertyFull {
	this := EntitlementPropertyFull{}
	return &this
}

// NewEntitlementPropertyFullWithDefaults instantiates a new EntitlementPropertyFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementPropertyFullWithDefaults() *EntitlementPropertyFull {
	this := EntitlementPropertyFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementPropertyFull) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPropertyFull) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementPropertyFull) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementPropertyFull) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementPropertyFull) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPropertyFull) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementPropertyFull) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementPropertyFull) SetName(v string) {
	o.Name = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *EntitlementPropertyFull) GetExternalValue() string {
	if o == nil || o.ExternalValue == nil {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPropertyFull) GetExternalValueOk() (*string, bool) {
	if o == nil || o.ExternalValue == nil {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementPropertyFull) HasExternalValue() bool {
	if o != nil && o.ExternalValue != nil {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *EntitlementPropertyFull) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementPropertyFull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPropertyFull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementPropertyFull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementPropertyFull) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value if set, zero value otherwise.
func (o *EntitlementPropertyFull) GetMultiValue() bool {
	if o == nil || o.MultiValue == nil {
		var ret bool
		return ret
	}
	return *o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPropertyFull) GetMultiValueOk() (*bool, bool) {
	if o == nil || o.MultiValue == nil {
		return nil, false
	}
	return o.MultiValue, true
}

// HasMultiValue returns a boolean if a field has been set.
func (o *EntitlementPropertyFull) HasMultiValue() bool {
	if o != nil && o.MultiValue != nil {
		return true
	}

	return false
}

// SetMultiValue gets a reference to the given bool and assigns it to the MultiValue field.
func (o *EntitlementPropertyFull) SetMultiValue(v bool) {
	o.MultiValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *EntitlementPropertyFull) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPropertyFull) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EntitlementPropertyFull) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *EntitlementPropertyFull) SetRequired(v bool) {
	o.Required = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *EntitlementPropertyFull) GetDataType() EntitlementPropertyDatatype {
	if o == nil || o.DataType == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPropertyFull) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *EntitlementPropertyFull) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given EntitlementPropertyDatatype and assigns it to the DataType field.
func (o *EntitlementPropertyFull) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = &v
}

func (o EntitlementPropertyFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
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

func (o *EntitlementPropertyFull) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementPropertyFull := _EntitlementPropertyFull{}

	err = json.Unmarshal(bytes, &varEntitlementPropertyFull)
	if err == nil {
		*o = EntitlementPropertyFull(varEntitlementPropertyFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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

type NullableEntitlementPropertyFull struct {
	value *EntitlementPropertyFull
	isSet bool
}

func (v NullableEntitlementPropertyFull) Get() *EntitlementPropertyFull {
	return v.value
}

func (v *NullableEntitlementPropertyFull) Set(val *EntitlementPropertyFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementPropertyFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementPropertyFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementPropertyFull(val *EntitlementPropertyFull) *NullableEntitlementPropertyFull {
	return &NullableEntitlementPropertyFull{value: val, isSet: true}
}

func (v NullableEntitlementPropertyFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementPropertyFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
