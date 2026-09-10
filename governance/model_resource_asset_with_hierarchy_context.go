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
	"time"
)

// checks if the ResourceAssetWithHierarchyContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetWithHierarchyContext{}

// ResourceAssetWithHierarchyContext Representation of a resource asset with context about its hierarchy.
type ResourceAssetWithHierarchyContext struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string              `json:"lastUpdatedBy"`
	Links         *ResourceAssetLinks `json:"_links,omitempty"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Orn string `json:"orn"`
	// The display name for a resource asset
	Name string `json:"name"`
	// The description of a resource asset
	Description *string                 `json:"description,omitempty"`
	Type        ResourceAssetTypeSparse `json:"type"`
	// The external ID of a resource asset
	ExternalId string `json:"externalId"`
	// Whether this asset has any children in the hierarchy. Present only if ?include=children_info is specified
	HasChildren *bool `json:"hasChildren,omitempty"`
	// The ID of the direct parent asset, or null if this is a root asset. Present only if ?include=ancestors is specified
	ParentId             NullableString `json:"parentId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetWithHierarchyContext ResourceAssetWithHierarchyContext

// NewResourceAssetWithHierarchyContext instantiates a new ResourceAssetWithHierarchyContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetWithHierarchyContext(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, orn string, name string, type_ ResourceAssetTypeSparse, externalId string) *ResourceAssetWithHierarchyContext {
	this := ResourceAssetWithHierarchyContext{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Orn = orn
	this.Name = name
	this.Type = type_
	this.ExternalId = externalId
	return &this
}

// NewResourceAssetWithHierarchyContextWithDefaults instantiates a new ResourceAssetWithHierarchyContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetWithHierarchyContextWithDefaults() *ResourceAssetWithHierarchyContext {
	this := ResourceAssetWithHierarchyContext{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceAssetWithHierarchyContext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceAssetWithHierarchyContext) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ResourceAssetWithHierarchyContext) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ResourceAssetWithHierarchyContext) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *ResourceAssetWithHierarchyContext) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ResourceAssetWithHierarchyContext) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *ResourceAssetWithHierarchyContext) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *ResourceAssetWithHierarchyContext) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *ResourceAssetWithHierarchyContext) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *ResourceAssetWithHierarchyContext) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceAssetWithHierarchyContext) GetLinks() ResourceAssetLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceAssetLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetLinksOk() (*ResourceAssetLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceAssetWithHierarchyContext) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceAssetLinks and assigns it to the Links field.
func (o *ResourceAssetWithHierarchyContext) SetLinks(v ResourceAssetLinks) {
	o.Links = &v
}

// GetOrn returns the Orn field value
func (o *ResourceAssetWithHierarchyContext) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *ResourceAssetWithHierarchyContext) SetOrn(v string) {
	o.Orn = v
}

// GetName returns the Name field value
func (o *ResourceAssetWithHierarchyContext) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceAssetWithHierarchyContext) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceAssetWithHierarchyContext) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceAssetWithHierarchyContext) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceAssetWithHierarchyContext) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *ResourceAssetWithHierarchyContext) GetType() ResourceAssetTypeSparse {
	if o == nil {
		var ret ResourceAssetTypeSparse
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetTypeOk() (*ResourceAssetTypeSparse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceAssetWithHierarchyContext) SetType(v ResourceAssetTypeSparse) {
	o.Type = v
}

// GetExternalId returns the ExternalId field value
func (o *ResourceAssetWithHierarchyContext) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *ResourceAssetWithHierarchyContext) SetExternalId(v string) {
	o.ExternalId = v
}

// GetHasChildren returns the HasChildren field value if set, zero value otherwise.
func (o *ResourceAssetWithHierarchyContext) GetHasChildren() bool {
	if o == nil || IsNil(o.HasChildren) {
		var ret bool
		return ret
	}
	return *o.HasChildren
}

// GetHasChildrenOk returns a tuple with the HasChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetWithHierarchyContext) GetHasChildrenOk() (*bool, bool) {
	if o == nil || IsNil(o.HasChildren) {
		return nil, false
	}
	return o.HasChildren, true
}

// HasHasChildren returns a boolean if a field has been set.
func (o *ResourceAssetWithHierarchyContext) HasHasChildren() bool {
	if o != nil && !IsNil(o.HasChildren) {
		return true
	}

	return false
}

// SetHasChildren gets a reference to the given bool and assigns it to the HasChildren field.
func (o *ResourceAssetWithHierarchyContext) SetHasChildren(v bool) {
	o.HasChildren = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceAssetWithHierarchyContext) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceAssetWithHierarchyContext) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ResourceAssetWithHierarchyContext) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *ResourceAssetWithHierarchyContext) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *ResourceAssetWithHierarchyContext) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *ResourceAssetWithHierarchyContext) UnsetParentId() {
	o.ParentId.Unset()
}

func (o ResourceAssetWithHierarchyContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetWithHierarchyContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["orn"] = o.Orn
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	toSerialize["externalId"] = o.ExternalId
	if !IsNil(o.HasChildren) {
		toSerialize["hasChildren"] = o.HasChildren
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetWithHierarchyContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"orn",
		"name",
		"type",
		"externalId",
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

	varResourceAssetWithHierarchyContext := _ResourceAssetWithHierarchyContext{}

	err = json.Unmarshal(data, &varResourceAssetWithHierarchyContext)

	if err != nil {
		return err
	}

	*o = ResourceAssetWithHierarchyContext(varResourceAssetWithHierarchyContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "hasChildren")
		delete(additionalProperties, "parentId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetWithHierarchyContext struct {
	value *ResourceAssetWithHierarchyContext
	isSet bool
}

func (v NullableResourceAssetWithHierarchyContext) Get() *ResourceAssetWithHierarchyContext {
	return v.value
}

func (v *NullableResourceAssetWithHierarchyContext) Set(val *ResourceAssetWithHierarchyContext) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetWithHierarchyContext) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetWithHierarchyContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetWithHierarchyContext(val *ResourceAssetWithHierarchyContext) *NullableResourceAssetWithHierarchyContext {
	return &NullableResourceAssetWithHierarchyContext{value: val, isSet: true}
}

func (v NullableResourceAssetWithHierarchyContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetWithHierarchyContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
