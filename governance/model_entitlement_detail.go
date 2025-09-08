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

// EntitlementDetail A single entitlement entry in the entitlements list of one historical record
type EntitlementDetail struct {
	// Collection of entitlement values.
	Values []EntitlementValueFull `json:"values,omitempty"`
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

type _EntitlementDetail EntitlementDetail

// NewEntitlementDetail instantiates a new EntitlementDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDetail() *EntitlementDetail {
	this := EntitlementDetail{}
	return &this
}

// NewEntitlementDetailWithDefaults instantiates a new EntitlementDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDetailWithDefaults() *EntitlementDetail {
	this := EntitlementDetail{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *EntitlementDetail) GetValues() []EntitlementValueFull {
	if o == nil || o.Values == nil {
		var ret []EntitlementValueFull
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetValuesOk() ([]EntitlementValueFull, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementDetail) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementValueFull and assigns it to the Values field.
func (o *EntitlementDetail) SetValues(v []EntitlementValueFull) {
	o.Values = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementDetail) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementDetail) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementDetail) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementDetail) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementDetail) SetName(v string) {
	o.Name = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *EntitlementDetail) GetExternalValue() string {
	if o == nil || o.ExternalValue == nil {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetExternalValueOk() (*string, bool) {
	if o == nil || o.ExternalValue == nil {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementDetail) HasExternalValue() bool {
	if o != nil && o.ExternalValue != nil {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *EntitlementDetail) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementDetail) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementDetail) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementDetail) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value if set, zero value otherwise.
func (o *EntitlementDetail) GetMultiValue() bool {
	if o == nil || o.MultiValue == nil {
		var ret bool
		return ret
	}
	return *o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetMultiValueOk() (*bool, bool) {
	if o == nil || o.MultiValue == nil {
		return nil, false
	}
	return o.MultiValue, true
}

// HasMultiValue returns a boolean if a field has been set.
func (o *EntitlementDetail) HasMultiValue() bool {
	if o != nil && o.MultiValue != nil {
		return true
	}

	return false
}

// SetMultiValue gets a reference to the given bool and assigns it to the MultiValue field.
func (o *EntitlementDetail) SetMultiValue(v bool) {
	o.MultiValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *EntitlementDetail) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EntitlementDetail) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *EntitlementDetail) SetRequired(v bool) {
	o.Required = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *EntitlementDetail) GetDataType() EntitlementPropertyDatatype {
	if o == nil || o.DataType == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *EntitlementDetail) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given EntitlementPropertyDatatype and assigns it to the DataType field.
func (o *EntitlementDetail) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = &v
}

func (o EntitlementDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
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

func (o *EntitlementDetail) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementDetail := _EntitlementDetail{}

	err = json.Unmarshal(bytes, &varEntitlementDetail)
	if err == nil {
		*o = EntitlementDetail(varEntitlementDetail)
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

type NullableEntitlementDetail struct {
	value *EntitlementDetail
	isSet bool
}

func (v NullableEntitlementDetail) Get() *EntitlementDetail {
	return v.value
}

func (v *NullableEntitlementDetail) Set(val *EntitlementDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDetail(val *EntitlementDetail) *NullableEntitlementDetail {
	return &NullableEntitlementDetail{value: val, isSet: true}
}

func (v NullableEntitlementDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
