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

// checks if the EntitlementUpdatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementUpdatable{}

// EntitlementUpdatable struct for EntitlementUpdatable
type EntitlementUpdatable struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn string         `json:"parentResourceOrn"`
	Parent            TargetResource `json:"parent"`
	// Collection of entitlement values for update operations.
	Values []EntitlementValueUpdatable `json:"values"`
	// The display name for an entitlement property
	Name string `json:"name"`
	// The value of an entitlement property
	ExternalValue string `json:"externalValue"`
	// The description of an entitlement property
	Description *string `json:"description,omitempty"`
	// Indicates if the entitlement property can hold multiple values. If this property is `true`, then the `dataType` property is set to `array`.
	MultiValue           bool                        `json:"multiValue"`
	DataType             EntitlementPropertyDatatype `json:"dataType"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementUpdatable EntitlementUpdatable

// NewEntitlementUpdatable instantiates a new EntitlementUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementUpdatable(id string, parentResourceOrn string, parent TargetResource, values []EntitlementValueUpdatable, name string, externalValue string, multiValue bool, dataType EntitlementPropertyDatatype) *EntitlementUpdatable {
	this := EntitlementUpdatable{}
	this.Name = name
	this.ExternalValue = externalValue
	this.MultiValue = multiValue
	this.DataType = dataType
	return &this
}

// NewEntitlementUpdatableWithDefaults instantiates a new EntitlementUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementUpdatableWithDefaults() *EntitlementUpdatable {
	this := EntitlementUpdatable{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementUpdatable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementUpdatable) SetId(v string) {
	o.Id = v
}

// GetParentResourceOrn returns the ParentResourceOrn field value
func (o *EntitlementUpdatable) GetParentResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetParentResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentResourceOrn, true
}

// SetParentResourceOrn sets field value
func (o *EntitlementUpdatable) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = v
}

// GetParent returns the Parent field value
func (o *EntitlementUpdatable) GetParent() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetParentOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *EntitlementUpdatable) SetParent(v TargetResource) {
	o.Parent = v
}

// GetValues returns the Values field value
func (o *EntitlementUpdatable) GetValues() []EntitlementValueUpdatable {
	if o == nil {
		var ret []EntitlementValueUpdatable
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetValuesOk() ([]EntitlementValueUpdatable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *EntitlementUpdatable) SetValues(v []EntitlementValueUpdatable) {
	o.Values = v
}

// GetName returns the Name field value
func (o *EntitlementUpdatable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementUpdatable) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *EntitlementUpdatable) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *EntitlementUpdatable) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementUpdatable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementUpdatable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementUpdatable) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value
func (o *EntitlementUpdatable) GetMultiValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetMultiValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiValue, true
}

// SetMultiValue sets field value
func (o *EntitlementUpdatable) SetMultiValue(v bool) {
	o.MultiValue = v
}

// GetDataType returns the DataType field value
func (o *EntitlementUpdatable) GetDataType() EntitlementPropertyDatatype {
	if o == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *EntitlementUpdatable) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *EntitlementUpdatable) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = v
}

func (o EntitlementUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementUpdatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	toSerialize["parent"] = o.Parent
	toSerialize["values"] = o.Values
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

func (o *EntitlementUpdatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"parentResourceOrn",
		"parent",
		"values",
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

	varEntitlementUpdatable := _EntitlementUpdatable{}

	err = json.Unmarshal(data, &varEntitlementUpdatable)

	if err != nil {
		return err
	}

	*o = EntitlementUpdatable(varEntitlementUpdatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "parentResourceOrn")
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

type NullableEntitlementUpdatable struct {
	value *EntitlementUpdatable
	isSet bool
}

func (v NullableEntitlementUpdatable) Get() *EntitlementUpdatable {
	return v.value
}

func (v *NullableEntitlementUpdatable) Set(val *EntitlementUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementUpdatable(val *EntitlementUpdatable) *NullableEntitlementUpdatable {
	return &NullableEntitlementUpdatable{value: val, isSet: true}
}

func (v NullableEntitlementUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
