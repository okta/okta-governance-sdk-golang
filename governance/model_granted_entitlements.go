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

// GrantedEntitlements struct for GrantedEntitlements
type GrantedEntitlements struct {
	Values []GrantedEntitlementValue `json:"values"`
	// The `id` property of an entitlement
	Id string `json:"id"`
	// The display name for an entitlement property
	Name string `json:"name"`
	// The value of an entitlement property
	ExternalValue string `json:"externalValue"`
	// The description of an entitlement property
	Description *string `json:"description,omitempty"`
	// The property that determines if the entitlement property can hold multiple values. If this is set to true, the data type is replaced with an array.
	MultiValue bool `json:"multiValue"`
	// The property that determines if the entitlement property is a required attribute
	Required             bool                        `json:"required"`
	DataType             EntitlementPropertyDatatype `json:"dataType"`
	AdditionalProperties map[string]interface{}
}

type _GrantedEntitlements GrantedEntitlements

// NewGrantedEntitlements instantiates a new GrantedEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantedEntitlements(values []GrantedEntitlementValue, id string, name string, externalValue string, multiValue bool, required bool, dataType EntitlementPropertyDatatype) *GrantedEntitlements {
	this := GrantedEntitlements{}
	this.Id = id
	this.Name = name
	this.ExternalValue = externalValue
	this.MultiValue = multiValue
	this.Required = required
	this.DataType = dataType
	return &this
}

// NewGrantedEntitlementsWithDefaults instantiates a new GrantedEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantedEntitlementsWithDefaults() *GrantedEntitlements {
	this := GrantedEntitlements{}
	return &this
}

// GetValues returns the Values field value
func (o *GrantedEntitlements) GetValues() []GrantedEntitlementValue {
	if o == nil {
		var ret []GrantedEntitlementValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetValuesOk() ([]GrantedEntitlementValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GrantedEntitlements) SetValues(v []GrantedEntitlementValue) {
	o.Values = v
}

// GetId returns the Id field value
func (o *GrantedEntitlements) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GrantedEntitlements) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GrantedEntitlements) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GrantedEntitlements) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *GrantedEntitlements) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *GrantedEntitlements) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GrantedEntitlements) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GrantedEntitlements) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GrantedEntitlements) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value
func (o *GrantedEntitlements) GetMultiValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetMultiValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiValue, true
}

// SetMultiValue sets field value
func (o *GrantedEntitlements) SetMultiValue(v bool) {
	o.MultiValue = v
}

// GetRequired returns the Required field value
func (o *GrantedEntitlements) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *GrantedEntitlements) SetRequired(v bool) {
	o.Required = v
}

// GetDataType returns the DataType field value
func (o *GrantedEntitlements) GetDataType() EntitlementPropertyDatatype {
	if o == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *GrantedEntitlements) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *GrantedEntitlements) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = v
}

func (o GrantedEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["values"] = o.Values
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["multiValue"] = o.MultiValue
	}
	if true {
		toSerialize["required"] = o.Required
	}
	if true {
		toSerialize["dataType"] = o.DataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GrantedEntitlements) UnmarshalJSON(bytes []byte) (err error) {
	varGrantedEntitlements := _GrantedEntitlements{}

	err = json.Unmarshal(bytes, &varGrantedEntitlements)
	if err == nil {
		*o = GrantedEntitlements(varGrantedEntitlements)
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

type NullableGrantedEntitlements struct {
	value *GrantedEntitlements
	isSet bool
}

func (v NullableGrantedEntitlements) Get() *GrantedEntitlements {
	return v.value
}

func (v *NullableGrantedEntitlements) Set(val *GrantedEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantedEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantedEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantedEntitlements(val *GrantedEntitlements) *NullableGrantedEntitlements {
	return &NullableGrantedEntitlements{value: val, isSet: true}
}

func (v NullableGrantedEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantedEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
