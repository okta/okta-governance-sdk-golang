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

// checks if the EntitlementsFullWithParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementsFullWithParent{}

// EntitlementsFullWithParent Representation of all entitlement. Entitlement values are optional
type EntitlementsFullWithParent struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn string         `json:"parentResourceOrn"`
	Parent            TargetResource `json:"parent"`
	// Collection of entitlement values.
	Values   []EntitlementValueFull `json:"values"`
	Links    *EntitlementLinks      `json:"_links,omitempty"`
	Metadata *ListMetadata          `json:"metadata,omitempty"`
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
	Required             *bool                       `json:"required,omitempty"`
	DataType             EntitlementPropertyDatatype `json:"dataType"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementsFullWithParent EntitlementsFullWithParent

// NewEntitlementsFullWithParent instantiates a new EntitlementsFullWithParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementsFullWithParent(parentResourceOrn string, parent TargetResource, values []EntitlementValueFull, id string, name string, externalValue string, multiValue bool, dataType EntitlementPropertyDatatype) *EntitlementsFullWithParent {
	this := EntitlementsFullWithParent{}
	this.Id = id
	this.Name = name
	this.ExternalValue = externalValue
	this.MultiValue = multiValue
	this.DataType = dataType
	return &this
}

// NewEntitlementsFullWithParentWithDefaults instantiates a new EntitlementsFullWithParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsFullWithParentWithDefaults() *EntitlementsFullWithParent {
	this := EntitlementsFullWithParent{}
	return &this
}

// GetParentResourceOrn returns the ParentResourceOrn field value
func (o *EntitlementsFullWithParent) GetParentResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetParentResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentResourceOrn, true
}

// SetParentResourceOrn sets field value
func (o *EntitlementsFullWithParent) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = v
}

// GetParent returns the Parent field value
func (o *EntitlementsFullWithParent) GetParent() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetParentOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *EntitlementsFullWithParent) SetParent(v TargetResource) {
	o.Parent = v
}

// GetValues returns the Values field value
func (o *EntitlementsFullWithParent) GetValues() []EntitlementValueFull {
	if o == nil {
		var ret []EntitlementValueFull
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetValuesOk() ([]EntitlementValueFull, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *EntitlementsFullWithParent) SetValues(v []EntitlementValueFull) {
	o.Values = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EntitlementsFullWithParent) GetLinks() EntitlementLinks {
	if o == nil || IsNil(o.Links) {
		var ret EntitlementLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetLinksOk() (*EntitlementLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EntitlementsFullWithParent) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EntitlementLinks and assigns it to the Links field.
func (o *EntitlementsFullWithParent) SetLinks(v EntitlementLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EntitlementsFullWithParent) GetMetadata() ListMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EntitlementsFullWithParent) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *EntitlementsFullWithParent) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

// GetId returns the Id field value
func (o *EntitlementsFullWithParent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementsFullWithParent) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EntitlementsFullWithParent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementsFullWithParent) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *EntitlementsFullWithParent) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *EntitlementsFullWithParent) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementsFullWithParent) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementsFullWithParent) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementsFullWithParent) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value
func (o *EntitlementsFullWithParent) GetMultiValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetMultiValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiValue, true
}

// SetMultiValue sets field value
func (o *EntitlementsFullWithParent) SetMultiValue(v bool) {
	o.MultiValue = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *EntitlementsFullWithParent) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EntitlementsFullWithParent) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *EntitlementsFullWithParent) SetRequired(v bool) {
	o.Required = &v
}

// GetDataType returns the DataType field value
func (o *EntitlementsFullWithParent) GetDataType() EntitlementPropertyDatatype {
	if o == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *EntitlementsFullWithParent) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *EntitlementsFullWithParent) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = v
}

func (o EntitlementsFullWithParent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementsFullWithParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	toSerialize["parent"] = o.Parent
	toSerialize["values"] = o.Values
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["externalValue"] = o.ExternalValue
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["multiValue"] = o.MultiValue
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	toSerialize["dataType"] = o.DataType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementsFullWithParent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parentResourceOrn",
		"parent",
		"values",
		"id",
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

	varEntitlementsFullWithParent := _EntitlementsFullWithParent{}

	err = json.Unmarshal(data, &varEntitlementsFullWithParent)

	if err != nil {
		return err
	}

	*o = EntitlementsFullWithParent(varEntitlementsFullWithParent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "values")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
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

type NullableEntitlementsFullWithParent struct {
	value *EntitlementsFullWithParent
	isSet bool
}

func (v NullableEntitlementsFullWithParent) Get() *EntitlementsFullWithParent {
	return v.value
}

func (v *NullableEntitlementsFullWithParent) Set(val *EntitlementsFullWithParent) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementsFullWithParent) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementsFullWithParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementsFullWithParent(val *EntitlementsFullWithParent) *NullableEntitlementsFullWithParent {
	return &NullableEntitlementsFullWithParent{value: val, isSet: true}
}

func (v NullableEntitlementsFullWithParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementsFullWithParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
