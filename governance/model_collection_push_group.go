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

// checks if the CollectionPushGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionPushGroup{}

// CollectionPushGroup A push group represented as a resource in a collection
type CollectionPushGroup struct {
	// Unique Okta Group ID for the push group
	Id *string `json:"id,omitempty"`
	// The name of the push group
	Name *string `json:"name,omitempty"`
	// The description of the push group
	Description *string `json:"description,omitempty"`
	// List of push group logo resources
	Logo []Link `json:"logo,omitempty"`
	// The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as `GET`/`PUT`/`DELETE /v2/collections/{collectionId}/resources/{resourceId}`.
	ResourceId string `json:"resourceId"`
	// The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.
	ResourceOrn          string `json:"resourceOrn"`
	AdditionalProperties map[string]interface{}
}

type _CollectionPushGroup CollectionPushGroup

// NewCollectionPushGroup instantiates a new CollectionPushGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionPushGroup(resourceId string, resourceOrn string) *CollectionPushGroup {
	this := CollectionPushGroup{}
	this.ResourceId = resourceId
	this.ResourceOrn = resourceOrn
	return &this
}

// NewCollectionPushGroupWithDefaults instantiates a new CollectionPushGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionPushGroupWithDefaults() *CollectionPushGroup {
	this := CollectionPushGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CollectionPushGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPushGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CollectionPushGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CollectionPushGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectionPushGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPushGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectionPushGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CollectionPushGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionPushGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPushGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionPushGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionPushGroup) SetDescription(v string) {
	o.Description = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CollectionPushGroup) GetLogo() []Link {
	if o == nil || IsNil(o.Logo) {
		var ret []Link
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPushGroup) GetLogoOk() ([]Link, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CollectionPushGroup) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []Link and assigns it to the Logo field.
func (o *CollectionPushGroup) SetLogo(v []Link) {
	o.Logo = v
}

// GetResourceId returns the ResourceId field value
func (o *CollectionPushGroup) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CollectionPushGroup) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *CollectionPushGroup) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionPushGroup) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionPushGroup) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionPushGroup) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

func (o CollectionPushGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionPushGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceOrn"] = o.ResourceOrn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionPushGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceId",
		"resourceOrn",
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

	varCollectionPushGroup := _CollectionPushGroup{}

	err = json.Unmarshal(data, &varCollectionPushGroup)

	if err != nil {
		return err
	}

	*o = CollectionPushGroup(varCollectionPushGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceOrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionPushGroup struct {
	value *CollectionPushGroup
	isSet bool
}

func (v NullableCollectionPushGroup) Get() *CollectionPushGroup {
	return v.value
}

func (v *NullableCollectionPushGroup) Set(val *CollectionPushGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPushGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPushGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPushGroup(val *CollectionPushGroup) *NullableCollectionPushGroup {
	return &NullableCollectionPushGroup{value: val, isSet: true}
}

func (v NullableCollectionPushGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPushGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
