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

// PrincipalEntitlement struct for PrincipalEntitlement
type PrincipalEntitlement struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn *string         `json:"parentResourceOrn,omitempty"`
	Parent            *TargetResource `json:"parent,omitempty"`
	// Collection of entitlement values.
	Values []EntitlementValueFull `json:"values,omitempty"`
	// The Okta user `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetPrincipalOrn *string              `json:"targetPrincipalOrn,omitempty"`
	TargetPrincipal    *TargetPrincipalFull `json:"targetPrincipal,omitempty"`
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

type _PrincipalEntitlement PrincipalEntitlement

// NewPrincipalEntitlement instantiates a new PrincipalEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlement() *PrincipalEntitlement {
	this := PrincipalEntitlement{}
	return &this
}

// NewPrincipalEntitlementWithDefaults instantiates a new PrincipalEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementWithDefaults() *PrincipalEntitlement {
	this := PrincipalEntitlement{}
	return &this
}

// GetParentResourceOrn returns the ParentResourceOrn field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetParentResourceOrn() string {
	if o == nil || o.ParentResourceOrn == nil {
		var ret string
		return ret
	}
	return *o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetParentResourceOrnOk() (*string, bool) {
	if o == nil || o.ParentResourceOrn == nil {
		return nil, false
	}
	return o.ParentResourceOrn, true
}

// HasParentResourceOrn returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasParentResourceOrn() bool {
	if o != nil && o.ParentResourceOrn != nil {
		return true
	}

	return false
}

// SetParentResourceOrn gets a reference to the given string and assigns it to the ParentResourceOrn field.
func (o *PrincipalEntitlement) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetParent() TargetResource {
	if o == nil || o.Parent == nil {
		var ret TargetResource
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetParentOk() (*TargetResource, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given TargetResource and assigns it to the Parent field.
func (o *PrincipalEntitlement) SetParent(v TargetResource) {
	o.Parent = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetValues() []EntitlementValueFull {
	if o == nil || o.Values == nil {
		var ret []EntitlementValueFull
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetValuesOk() ([]EntitlementValueFull, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntitlementValueFull and assigns it to the Values field.
func (o *PrincipalEntitlement) SetValues(v []EntitlementValueFull) {
	o.Values = v
}

// GetTargetPrincipalOrn returns the TargetPrincipalOrn field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetTargetPrincipalOrn() string {
	if o == nil || o.TargetPrincipalOrn == nil {
		var ret string
		return ret
	}
	return *o.TargetPrincipalOrn
}

// GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetTargetPrincipalOrnOk() (*string, bool) {
	if o == nil || o.TargetPrincipalOrn == nil {
		return nil, false
	}
	return o.TargetPrincipalOrn, true
}

// HasTargetPrincipalOrn returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasTargetPrincipalOrn() bool {
	if o != nil && o.TargetPrincipalOrn != nil {
		return true
	}

	return false
}

// SetTargetPrincipalOrn gets a reference to the given string and assigns it to the TargetPrincipalOrn field.
func (o *PrincipalEntitlement) SetTargetPrincipalOrn(v string) {
	o.TargetPrincipalOrn = &v
}

// GetTargetPrincipal returns the TargetPrincipal field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetTargetPrincipal() TargetPrincipalFull {
	if o == nil || o.TargetPrincipal == nil {
		var ret TargetPrincipalFull
		return ret
	}
	return *o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetTargetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil || o.TargetPrincipal == nil {
		return nil, false
	}
	return o.TargetPrincipal, true
}

// HasTargetPrincipal returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasTargetPrincipal() bool {
	if o != nil && o.TargetPrincipal != nil {
		return true
	}

	return false
}

// SetTargetPrincipal gets a reference to the given TargetPrincipalFull and assigns it to the TargetPrincipal field.
func (o *PrincipalEntitlement) SetTargetPrincipal(v TargetPrincipalFull) {
	o.TargetPrincipal = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PrincipalEntitlement) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrincipalEntitlement) SetName(v string) {
	o.Name = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetExternalValue() string {
	if o == nil || o.ExternalValue == nil {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetExternalValueOk() (*string, bool) {
	if o == nil || o.ExternalValue == nil {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasExternalValue() bool {
	if o != nil && o.ExternalValue != nil {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *PrincipalEntitlement) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PrincipalEntitlement) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValue returns the MultiValue field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetMultiValue() bool {
	if o == nil || o.MultiValue == nil {
		var ret bool
		return ret
	}
	return *o.MultiValue
}

// GetMultiValueOk returns a tuple with the MultiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetMultiValueOk() (*bool, bool) {
	if o == nil || o.MultiValue == nil {
		return nil, false
	}
	return o.MultiValue, true
}

// HasMultiValue returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasMultiValue() bool {
	if o != nil && o.MultiValue != nil {
		return true
	}

	return false
}

// SetMultiValue gets a reference to the given bool and assigns it to the MultiValue field.
func (o *PrincipalEntitlement) SetMultiValue(v bool) {
	o.MultiValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PrincipalEntitlement) SetRequired(v bool) {
	o.Required = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *PrincipalEntitlement) GetDataType() EntitlementPropertyDatatype {
	if o == nil || o.DataType == nil {
		var ret EntitlementPropertyDatatype
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlement) GetDataTypeOk() (*EntitlementPropertyDatatype, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *PrincipalEntitlement) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given EntitlementPropertyDatatype and assigns it to the DataType field.
func (o *PrincipalEntitlement) SetDataType(v EntitlementPropertyDatatype) {
	o.DataType = &v
}

func (o PrincipalEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentResourceOrn != nil {
		toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.TargetPrincipalOrn != nil {
		toSerialize["targetPrincipalOrn"] = o.TargetPrincipalOrn
	}
	if o.TargetPrincipal != nil {
		toSerialize["targetPrincipal"] = o.TargetPrincipal
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

func (o *PrincipalEntitlement) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalEntitlement := _PrincipalEntitlement{}

	err = json.Unmarshal(bytes, &varPrincipalEntitlement)
	if err == nil {
		*o = PrincipalEntitlement(varPrincipalEntitlement)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "values")
		delete(additionalProperties, "targetPrincipalOrn")
		delete(additionalProperties, "targetPrincipal")
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

type NullablePrincipalEntitlement struct {
	value *PrincipalEntitlement
	isSet bool
}

func (v NullablePrincipalEntitlement) Get() *PrincipalEntitlement {
	return v.value
}

func (v *NullablePrincipalEntitlement) Set(val *PrincipalEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlement(val *PrincipalEntitlement) *NullablePrincipalEntitlement {
	return &NullablePrincipalEntitlement{value: val, isSet: true}
}

func (v NullablePrincipalEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
