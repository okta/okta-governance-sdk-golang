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

// EntitlementValueWithParent Attributes related to an entitlement value
type EntitlementValueWithParent struct {
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

type _EntitlementValueWithParent EntitlementValueWithParent

// NewEntitlementValueWithParent instantiates a new EntitlementValueWithParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValueWithParent(id string, name string, externalValue string, entitlementId string, parentResourceOrn string, parent TargetResource) *EntitlementValueWithParent {
	this := EntitlementValueWithParent{}
	this.Id = id
	this.Name = name
	this.ExternalValue = externalValue
	this.EntitlementId = entitlementId
	this.ParentResourceOrn = parentResourceOrn
	this.Parent = parent
	return &this
}

// NewEntitlementValueWithParentWithDefaults instantiates a new EntitlementValueWithParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueWithParentWithDefaults() *EntitlementValueWithParent {
	this := EntitlementValueWithParent{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementValueWithParent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementValueWithParent) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EntitlementValueWithParent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementValueWithParent) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *EntitlementValueWithParent) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *EntitlementValueWithParent) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EntitlementValueWithParent) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EntitlementValueWithParent) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EntitlementValueWithParent) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementValueWithParent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementValueWithParent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementValueWithParent) SetDescription(v string) {
	o.Description = &v
}

// GetEntitlementId returns the EntitlementId field value
func (o *EntitlementValueWithParent) GetEntitlementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetEntitlementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementId, true
}

// SetEntitlementId sets field value
func (o *EntitlementValueWithParent) SetEntitlementId(v string) {
	o.EntitlementId = v
}

// GetParentResourceOrn returns the ParentResourceOrn field value
func (o *EntitlementValueWithParent) GetParentResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetParentResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentResourceOrn, true
}

// SetParentResourceOrn sets field value
func (o *EntitlementValueWithParent) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = v
}

// GetParent returns the Parent field value
func (o *EntitlementValueWithParent) GetParent() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *EntitlementValueWithParent) GetParentOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *EntitlementValueWithParent) SetParent(v TargetResource) {
	o.Parent = v
}

func (o EntitlementValueWithParent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["entitlementId"] = o.EntitlementId
	}
	if true {
		toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	}
	if true {
		toSerialize["parent"] = o.Parent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValueWithParent) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValueWithParent := _EntitlementValueWithParent{}

	err = json.Unmarshal(bytes, &varEntitlementValueWithParent)
	if err == nil {
		*o = EntitlementValueWithParent(varEntitlementValueWithParent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "entitlementId")
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "parent")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValueWithParent struct {
	value *EntitlementValueWithParent
	isSet bool
}

func (v NullableEntitlementValueWithParent) Get() *EntitlementValueWithParent {
	return v.value
}

func (v *NullableEntitlementValueWithParent) Set(val *EntitlementValueWithParent) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValueWithParent) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValueWithParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValueWithParent(val *EntitlementValueWithParent) *NullableEntitlementValueWithParent {
	return &NullableEntitlementValueWithParent{value: val, isSet: true}
}

func (v NullableEntitlementValueWithParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValueWithParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
