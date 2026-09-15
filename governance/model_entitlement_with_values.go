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

// checks if the EntitlementWithValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementWithValues{}

// EntitlementWithValues An entitlement property that the principal effectively has, with the first page of its effective values
type EntitlementWithValues struct {
	// The `id` property of an entitlement
	Id string `json:"id"`
	// The display name for an entitlement property
	Name string `json:"name"`
	// The value of an entitlement property
	ExternalValue *string `json:"externalValue,omitempty"`
	// The description of an entitlement property
	Description *string `json:"description,omitempty"`
	// Indicates if the entitlement property can hold multiple values. If this property is `true`, then the `dataType` property is set to `array`.
	MultiValue *bool `json:"multiValue,omitempty"`
	// The property that determines if the entitlement property is a required attribute
	Required *bool                        `json:"required,omitempty"`
	DataType *EntitlementPropertyDatatype `json:"dataType,omitempty"`
	// The first page of the principal's effective values for this entitlement property. Up to 20 values are returned inline, ordered by name in ascending order. If more values remain, the `_links.next` reference points to `GET /governance/api/v2/principal-entitlements/values` with the query parameters for the next page. When the `_links.next` reference is absent, `values` is the complete set.
	Values               []PrincipalEntitlementValue      `json:"values"`
	Links                *EntitlementWithValuesAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementWithValues EntitlementWithValues

// NewEntitlementWithValues instantiates a new EntitlementWithValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementWithValues(id string, name string, values []PrincipalEntitlementValue) *EntitlementWithValues {
	this := EntitlementWithValues{}
	this.Id = id
	this.Name = name
	this.Values = values
	return &this
}

// NewEntitlementWithValuesWithDefaults instantiates a new EntitlementWithValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWithValuesWithDefaults() *EntitlementWithValues {
	this := EntitlementWithValues{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementWithValues) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementWithValues) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EntitlementWithValues) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementWithValues) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *EntitlementWithValues) GetExternalValue() string {
	if o == nil || IsNil(o.ExternalValue) {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetExternalValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalValue) {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementWithValues) HasExternalValue() bool {
	if o != nil && !IsNil(o.ExternalValue) {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *EntitlementWithValues) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementWithValues) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementWithValues) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementWithValues) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value if set, zero value otherwise.
func (o *EntitlementWithValues) GetMultiValue() bool {
	if o == nil || IsNil(o.MultiValue) {
		var ret bool
		return ret
	}
	return *o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetMultiValueOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiValue) {
		return nil, false
	}
	return o.MultiValue, true
}

// HasMultiValue returns a boolean if a field has been set.
func (o *EntitlementWithValues) HasMultiValue() bool {
	if o != nil && !IsNil(o.MultiValue) {
		return true
	}

	return false
}

// SetMultiValue gets a reference to the given bool and assigns it to the MultiValue field.
func (o *EntitlementWithValues) SetMultiValue(v bool) {
	o.MultiValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *EntitlementWithValues) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EntitlementWithValues) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *EntitlementWithValues) SetRequired(v bool) {
	o.Required = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *EntitlementWithValues) GetDataType() EntitlementPropertyDatatype {
	if o == nil || IsNil(o.DataType) {
		var ret EntitlementPropertyDatatype
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *EntitlementWithValues) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given EntitlementPropertyDatatype and assigns it to the DataType field.
func (o *EntitlementWithValues) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = &v
}

// GetValues returns the Values field value
func (o *EntitlementWithValues) GetValues() []PrincipalEntitlementValue {
	if o == nil {
		var ret []PrincipalEntitlementValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetValuesOk() ([]PrincipalEntitlementValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *EntitlementWithValues) SetValues(v []PrincipalEntitlementValue) {
	o.Values = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EntitlementWithValues) GetLinks() EntitlementWithValuesAllOfLinks {
	if o == nil || IsNil(o.Links) {
		var ret EntitlementWithValuesAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementWithValues) GetLinksOk() (*EntitlementWithValuesAllOfLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EntitlementWithValues) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EntitlementWithValuesAllOfLinks and assigns it to the Links field.
func (o *EntitlementWithValues) SetLinks(v EntitlementWithValuesAllOfLinks) {
	o.Links = &v
}

func (o EntitlementWithValues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementWithValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
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
	toSerialize["values"] = o.Values
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementWithValues) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"values",
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

	varEntitlementWithValues := _EntitlementWithValues{}

	err = json.Unmarshal(data, &varEntitlementWithValues)

	if err != nil {
		return err
	}

	*o = EntitlementWithValues(varEntitlementWithValues)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "description")
		delete(additionalProperties, "multiValue")
		delete(additionalProperties, "required")
		delete(additionalProperties, "dataType")
		delete(additionalProperties, "values")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementWithValues struct {
	value *EntitlementWithValues
	isSet bool
}

func (v NullableEntitlementWithValues) Get() *EntitlementWithValues {
	return v.value
}

func (v *NullableEntitlementWithValues) Set(val *EntitlementWithValues) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementWithValues) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementWithValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementWithValues(val *EntitlementWithValues) *NullableEntitlementWithValues {
	return &NullableEntitlementWithValues{value: val, isSet: true}
}

func (v NullableEntitlementWithValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementWithValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
