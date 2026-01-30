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

// checks if the CollectionResourcePropertiesReadOnly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourcePropertiesReadOnly{}

// CollectionResourcePropertiesReadOnly struct for CollectionResourcePropertiesReadOnly
type CollectionResourcePropertiesReadOnly struct {
	// The Okta `app.id` of the resource  See the [List applications](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Application/#tag/Application/operation/listApplications) endpoint to retrieve application IDs.
	ResourceId           *string                  `json:"resourceId,omitempty"`
	Links                *CollectionResourceLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourcePropertiesReadOnly CollectionResourcePropertiesReadOnly

// NewCollectionResourcePropertiesReadOnly instantiates a new CollectionResourcePropertiesReadOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourcePropertiesReadOnly() *CollectionResourcePropertiesReadOnly {
	this := CollectionResourcePropertiesReadOnly{}
	return &this
}

// NewCollectionResourcePropertiesReadOnlyWithDefaults instantiates a new CollectionResourcePropertiesReadOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourcePropertiesReadOnlyWithDefaults() *CollectionResourcePropertiesReadOnly {
	this := CollectionResourcePropertiesReadOnly{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *CollectionResourcePropertiesReadOnly) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourcePropertiesReadOnly) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CollectionResourcePropertiesReadOnly) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *CollectionResourcePropertiesReadOnly) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionResourcePropertiesReadOnly) GetLinks() CollectionResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret CollectionResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourcePropertiesReadOnly) GetLinksOk() (*CollectionResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CollectionResourcePropertiesReadOnly) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionResourceLinks and assigns it to the Links field.
func (o *CollectionResourcePropertiesReadOnly) SetLinks(v CollectionResourceLinks) {
	o.Links = &v
}

func (o CollectionResourcePropertiesReadOnly) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourcePropertiesReadOnly) ToMap() (map[string]interface{}, error) {
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

func (o *CollectionResourcePropertiesReadOnly) UnmarshalJSON(data []byte) (err error) {
	varCollectionResourcePropertiesReadOnly := _CollectionResourcePropertiesReadOnly{}

	err = json.Unmarshal(data, &varCollectionResourcePropertiesReadOnly)

	if err != nil {
		return err
	}

	*o = CollectionResourcePropertiesReadOnly(varCollectionResourcePropertiesReadOnly)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourcePropertiesReadOnly struct {
	value *CollectionResourcePropertiesReadOnly
	isSet bool
}

func (v NullableCollectionResourcePropertiesReadOnly) Get() *CollectionResourcePropertiesReadOnly {
	return v.value
}

func (v *NullableCollectionResourcePropertiesReadOnly) Set(val *CollectionResourcePropertiesReadOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourcePropertiesReadOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourcePropertiesReadOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourcePropertiesReadOnly(val *CollectionResourcePropertiesReadOnly) *NullableCollectionResourcePropertiesReadOnly {
	return &NullableCollectionResourcePropertiesReadOnly{value: val, isSet: true}
}

func (v NullableCollectionResourcePropertiesReadOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourcePropertiesReadOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
