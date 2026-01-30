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

// checks if the EntitlementValue2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementValue2{}

// EntitlementValue2 struct for EntitlementValue2
type EntitlementValue2 struct {
	Links EntitlementLink `json:"_links"`
	// The `id` of an entitlement value
	Id string `json:"id"`
	// The display name for an entitlement value
	Name string `json:"name"`
	// The value of an entitlement property value
	ExternalValue string `json:"externalValue"`
	// The read-only `id` of an entitlement property value in the downstream application.
	ExternalId *string `json:"externalId,omitempty"`
	// The description of an entitlement value
	Description *string `json:"description,omitempty"`
	// The `id` property of an entitlement
	EntitlementId string `json:"entitlementId"`
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn    string         `json:"parentResourceOrn"`
	Parent               TargetResource `json:"parent"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValue2 EntitlementValue2

// NewEntitlementValue2 instantiates a new EntitlementValue2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValue2(links EntitlementLink, id string, name string, externalValue string, entitlementId string, parentResourceOrn string, parent TargetResource) *EntitlementValue2 {
	this := EntitlementValue2{}
	this.Id = id
	this.Name = name
	this.ExternalValue = externalValue
	this.EntitlementId = entitlementId
	this.ParentResourceOrn = parentResourceOrn
	this.Parent = parent
	return &this
}

// NewEntitlementValue2WithDefaults instantiates a new EntitlementValue2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValue2WithDefaults() *EntitlementValue2 {
	this := EntitlementValue2{}
	return &this
}

// GetLinks returns the Links field value
func (o *EntitlementValue2) GetLinks() EntitlementLink {
	if o == nil {
		var ret EntitlementLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetLinksOk() (*EntitlementLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EntitlementValue2) SetLinks(v EntitlementLink) {
	o.Links = v
}

// GetId returns the Id field value
func (o *EntitlementValue2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementValue2) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EntitlementValue2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementValue2) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *EntitlementValue2) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *EntitlementValue2) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EntitlementValue2) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EntitlementValue2) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EntitlementValue2) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementValue2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementValue2) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementValue2) SetDescription(v string) {
	o.Description = &v
}

// GetEntitlementId returns the EntitlementId field value
func (o *EntitlementValue2) GetEntitlementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetEntitlementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementId, true
}

// SetEntitlementId sets field value
func (o *EntitlementValue2) SetEntitlementId(v string) {
	o.EntitlementId = v
}

// GetParentResourceOrn returns the ParentResourceOrn field value
func (o *EntitlementValue2) GetParentResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetParentResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentResourceOrn, true
}

// SetParentResourceOrn sets field value
func (o *EntitlementValue2) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = v
}

// GetParent returns the Parent field value
func (o *EntitlementValue2) GetParent() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *EntitlementValue2) GetParentOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *EntitlementValue2) SetParent(v TargetResource) {
	o.Parent = v
}

func (o EntitlementValue2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementValue2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["externalValue"] = o.ExternalValue
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["entitlementId"] = o.EntitlementId
	toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	toSerialize["parent"] = o.Parent

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementValue2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"id",
		"name",
		"externalValue",
		"entitlementId",
		"parentResourceOrn",
		"parent",
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

	varEntitlementValue2 := _EntitlementValue2{}

	err = json.Unmarshal(data, &varEntitlementValue2)

	if err != nil {
		return err
	}

	*o = EntitlementValue2(varEntitlementValue2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "entitlementId")
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "parent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementValue2 struct {
	value *EntitlementValue2
	isSet bool
}

func (v NullableEntitlementValue2) Get() *EntitlementValue2 {
	return v.value
}

func (v *NullableEntitlementValue2) Set(val *EntitlementValue2) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValue2) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValue2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValue2(val *EntitlementValue2) *NullableEntitlementValue2 {
	return &NullableEntitlementValue2{value: val, isSet: true}
}

func (v NullableEntitlementValue2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValue2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
