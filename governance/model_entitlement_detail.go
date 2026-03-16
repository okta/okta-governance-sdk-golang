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

// checks if the EntitlementDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDetail{}

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
	// Indicate if the entitlement property can hold multiple values. If this property is `true`, then the `dataType` property is set to  `array`.
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
	if o == nil || IsNil(o.Values) {
		var ret []EntitlementValueFull
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetValuesOk() ([]EntitlementValueFull, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *EntitlementDetail) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.ExternalValue) {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetExternalValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalValue) {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementDetail) HasExternalValue() bool {
	if o != nil && !IsNil(o.ExternalValue) {
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.MultiValue) {
		var ret bool
		return ret
	}
	return *o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetMultiValueOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiValue) {
		return nil, false
	}
	return o.MultiValue, true
}

// HasMultiValue returns a boolean if a field has been set.
func (o *EntitlementDetail) HasMultiValue() bool {
	if o != nil && !IsNil(o.MultiValue) {
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
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EntitlementDetail) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
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
	if o == nil || IsNil(o.DataType) {
		var ret EntitlementPropertyDatatype
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDetail) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *EntitlementDetail) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given EntitlementPropertyDatatype and assigns it to the DataType field.
func (o *EntitlementDetail) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = &v
}

func (o EntitlementDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExternalValue) {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MultiValue) {
		toSerialize["multiValue"] = o.MultiValue
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDetail) UnmarshalJSON(data []byte) (err error) {
	varEntitlementDetail := _EntitlementDetail{}

	err = json.Unmarshal(data, &varEntitlementDetail)

	if err != nil {
		return err
	}

	*o = EntitlementDetail(varEntitlementDetail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "values")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "description")
		delete(additionalProperties, "multiValue")
		delete(additionalProperties, "required")
		delete(additionalProperties, "dataType")
		o.AdditionalProperties = additionalProperties
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
