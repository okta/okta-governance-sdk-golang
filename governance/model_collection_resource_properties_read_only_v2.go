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
)

// checks if the CollectionResourcePropertiesReadOnlyV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourcePropertiesReadOnlyV2{}

// CollectionResourcePropertiesReadOnlyV2 struct for CollectionResourcePropertiesReadOnlyV2
type CollectionResourcePropertiesReadOnlyV2 struct {
	// The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as `GET`/`PUT`/`DELETE /v2/collections/{collectionId}/resources/{resourceId}`.
	ResourceId           *string                    `json:"resourceId,omitempty"`
	Links                *CollectionResourceLinksV2 `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourcePropertiesReadOnlyV2 CollectionResourcePropertiesReadOnlyV2

// NewCollectionResourcePropertiesReadOnlyV2 instantiates a new CollectionResourcePropertiesReadOnlyV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourcePropertiesReadOnlyV2() *CollectionResourcePropertiesReadOnlyV2 {
	this := CollectionResourcePropertiesReadOnlyV2{}
	return &this
}

// NewCollectionResourcePropertiesReadOnlyV2WithDefaults instantiates a new CollectionResourcePropertiesReadOnlyV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourcePropertiesReadOnlyV2WithDefaults() *CollectionResourcePropertiesReadOnlyV2 {
	this := CollectionResourcePropertiesReadOnlyV2{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *CollectionResourcePropertiesReadOnlyV2) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourcePropertiesReadOnlyV2) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CollectionResourcePropertiesReadOnlyV2) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *CollectionResourcePropertiesReadOnlyV2) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionResourcePropertiesReadOnlyV2) GetLinks() CollectionResourceLinksV2 {
	if o == nil || IsNil(o.Links) {
		var ret CollectionResourceLinksV2
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourcePropertiesReadOnlyV2) GetLinksOk() (*CollectionResourceLinksV2, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CollectionResourcePropertiesReadOnlyV2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionResourceLinksV2 and assigns it to the Links field.
func (o *CollectionResourcePropertiesReadOnlyV2) SetLinks(v CollectionResourceLinksV2) {
	o.Links = &v
}

func (o CollectionResourcePropertiesReadOnlyV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourcePropertiesReadOnlyV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourcePropertiesReadOnlyV2) UnmarshalJSON(data []byte) (err error) {
	varCollectionResourcePropertiesReadOnlyV2 := _CollectionResourcePropertiesReadOnlyV2{}

	err = json.Unmarshal(data, &varCollectionResourcePropertiesReadOnlyV2)

	if err != nil {
		return err
	}

	*o = CollectionResourcePropertiesReadOnlyV2(varCollectionResourcePropertiesReadOnlyV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourcePropertiesReadOnlyV2 struct {
	value *CollectionResourcePropertiesReadOnlyV2
	isSet bool
}

func (v NullableCollectionResourcePropertiesReadOnlyV2) Get() *CollectionResourcePropertiesReadOnlyV2 {
	return v.value
}

func (v *NullableCollectionResourcePropertiesReadOnlyV2) Set(val *CollectionResourcePropertiesReadOnlyV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourcePropertiesReadOnlyV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourcePropertiesReadOnlyV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourcePropertiesReadOnlyV2(val *CollectionResourcePropertiesReadOnlyV2) *NullableCollectionResourcePropertiesReadOnlyV2 {
	return &NullableCollectionResourcePropertiesReadOnlyV2{value: val, isSet: true}
}

func (v NullableCollectionResourcePropertiesReadOnlyV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourcePropertiesReadOnlyV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
