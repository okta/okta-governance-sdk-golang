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

// checks if the CollectionFullWithFilterContextV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionFullWithFilterContextV2{}

// CollectionFullWithFilterContextV2 A collection with filter-context fields, populated only on LIST responses when filtering by `resourceOrn` or `principalOrn`.
type CollectionFullWithFilterContextV2 struct {
	// The name of a resource collection
	Name string `json:"name"`
	// The human-readable description
	Description *string `json:"description,omitempty"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string          `json:"lastUpdatedBy"`
	Links         CollectionLinks `json:"_links"`
	// The `id` of the collection in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).
	Orn                  string                          `json:"orn"`
	Counts               *CollectionCountsV2             `json:"counts,omitempty"`
	ResourceRelationship *CollectionResourceRelationship `json:"resourceRelationship,omitempty"`
	PrincipalAssignment  *CollectionPrincipalAssignment  `json:"principalAssignment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionFullWithFilterContextV2 CollectionFullWithFilterContextV2

// NewCollectionFullWithFilterContextV2 instantiates a new CollectionFullWithFilterContextV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionFullWithFilterContextV2(name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links CollectionLinks, orn string) *CollectionFullWithFilterContextV2 {
	this := CollectionFullWithFilterContextV2{}
	this.Name = name
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.Orn = orn
	return &this
}

// NewCollectionFullWithFilterContextV2WithDefaults instantiates a new CollectionFullWithFilterContextV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionFullWithFilterContextV2WithDefaults() *CollectionFullWithFilterContextV2 {
	this := CollectionFullWithFilterContextV2{}
	return &this
}

// GetName returns the Name field value
func (o *CollectionFullWithFilterContextV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CollectionFullWithFilterContextV2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionFullWithFilterContextV2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionFullWithFilterContextV2) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionFullWithFilterContextV2) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *CollectionFullWithFilterContextV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CollectionFullWithFilterContextV2) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CollectionFullWithFilterContextV2) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *CollectionFullWithFilterContextV2) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *CollectionFullWithFilterContextV2) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CollectionFullWithFilterContextV2) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *CollectionFullWithFilterContextV2) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *CollectionFullWithFilterContextV2) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *CollectionFullWithFilterContextV2) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *CollectionFullWithFilterContextV2) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *CollectionFullWithFilterContextV2) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CollectionFullWithFilterContextV2) SetLinks(v CollectionLinks) {
	o.Links = v
}

// GetOrn returns the Orn field value
func (o *CollectionFullWithFilterContextV2) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *CollectionFullWithFilterContextV2) SetOrn(v string) {
	o.Orn = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *CollectionFullWithFilterContextV2) GetCounts() CollectionCountsV2 {
	if o == nil || IsNil(o.Counts) {
		var ret CollectionCountsV2
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetCountsOk() (*CollectionCountsV2, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *CollectionFullWithFilterContextV2) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given CollectionCountsV2 and assigns it to the Counts field.
func (o *CollectionFullWithFilterContextV2) SetCounts(v CollectionCountsV2) {
	o.Counts = &v
}

// GetResourceRelationship returns the ResourceRelationship field value if set, zero value otherwise.
func (o *CollectionFullWithFilterContextV2) GetResourceRelationship() CollectionResourceRelationship {
	if o == nil || IsNil(o.ResourceRelationship) {
		var ret CollectionResourceRelationship
		return ret
	}
	return *o.ResourceRelationship
}

// GetResourceRelationshipOk returns a tuple with the ResourceRelationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetResourceRelationshipOk() (*CollectionResourceRelationship, bool) {
	if o == nil || IsNil(o.ResourceRelationship) {
		return nil, false
	}
	return o.ResourceRelationship, true
}

// HasResourceRelationship returns a boolean if a field has been set.
func (o *CollectionFullWithFilterContextV2) HasResourceRelationship() bool {
	if o != nil && !IsNil(o.ResourceRelationship) {
		return true
	}

	return false
}

// SetResourceRelationship gets a reference to the given CollectionResourceRelationship and assigns it to the ResourceRelationship field.
func (o *CollectionFullWithFilterContextV2) SetResourceRelationship(v CollectionResourceRelationship) {
	o.ResourceRelationship = &v
}

// GetPrincipalAssignment returns the PrincipalAssignment field value if set, zero value otherwise.
func (o *CollectionFullWithFilterContextV2) GetPrincipalAssignment() CollectionPrincipalAssignment {
	if o == nil || IsNil(o.PrincipalAssignment) {
		var ret CollectionPrincipalAssignment
		return ret
	}
	return *o.PrincipalAssignment
}

// GetPrincipalAssignmentOk returns a tuple with the PrincipalAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFullWithFilterContextV2) GetPrincipalAssignmentOk() (*CollectionPrincipalAssignment, bool) {
	if o == nil || IsNil(o.PrincipalAssignment) {
		return nil, false
	}
	return o.PrincipalAssignment, true
}

// HasPrincipalAssignment returns a boolean if a field has been set.
func (o *CollectionFullWithFilterContextV2) HasPrincipalAssignment() bool {
	if o != nil && !IsNil(o.PrincipalAssignment) {
		return true
	}

	return false
}

// SetPrincipalAssignment gets a reference to the given CollectionPrincipalAssignment and assigns it to the PrincipalAssignment field.
func (o *CollectionFullWithFilterContextV2) SetPrincipalAssignment(v CollectionPrincipalAssignment) {
	o.PrincipalAssignment = &v
}

func (o CollectionFullWithFilterContextV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionFullWithFilterContextV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	toSerialize["_links"] = o.Links
	toSerialize["orn"] = o.Orn
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !IsNil(o.ResourceRelationship) {
		toSerialize["resourceRelationship"] = o.ResourceRelationship
	}
	if !IsNil(o.PrincipalAssignment) {
		toSerialize["principalAssignment"] = o.PrincipalAssignment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionFullWithFilterContextV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"_links",
		"orn",
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

	varCollectionFullWithFilterContextV2 := _CollectionFullWithFilterContextV2{}

	err = json.Unmarshal(data, &varCollectionFullWithFilterContextV2)

	if err != nil {
		return err
	}

	*o = CollectionFullWithFilterContextV2(varCollectionFullWithFilterContextV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "counts")
		delete(additionalProperties, "resourceRelationship")
		delete(additionalProperties, "principalAssignment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionFullWithFilterContextV2 struct {
	value *CollectionFullWithFilterContextV2
	isSet bool
}

func (v NullableCollectionFullWithFilterContextV2) Get() *CollectionFullWithFilterContextV2 {
	return v.value
}

func (v *NullableCollectionFullWithFilterContextV2) Set(val *CollectionFullWithFilterContextV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionFullWithFilterContextV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionFullWithFilterContextV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionFullWithFilterContextV2(val *CollectionFullWithFilterContextV2) *NullableCollectionFullWithFilterContextV2 {
	return &NullableCollectionFullWithFilterContextV2{value: val, isSet: true}
}

func (v NullableCollectionFullWithFilterContextV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionFullWithFilterContextV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
