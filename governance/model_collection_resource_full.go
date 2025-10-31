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

// checks if the CollectionResourceFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceFull{}

// CollectionResourceFull Full representation of a resource within a resource collection
type CollectionResourceFull struct {
	ResourceProfile *ResourceProfile `json:"resourceProfile,omitempty"`
	// Collection of entitlements with associated values
	Entitlements []EntitlementFull `json:"entitlements,omitempty"`
	// The number of entitlements associated with this resource in the collection. Use the `include` query parameter to return this count in the response.
	EntitlementValueCount *int32 `json:"entitlementValueCount,omitempty"`
	// The ORN identifier for a specific app. Other resource types aren't supported.  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint for reference.
	ResourceOrn string `json:"resourceOrn"`
	// The Okta `app.id` of the resource  See the [List applications](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Application/#tag/Application/operation/listApplications) endpoint to retrieve application IDs.
	ResourceId           *string                 `json:"resourceId,omitempty"`
	Links                CollectionResourceLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceFull CollectionResourceFull

// NewCollectionResourceFull instantiates a new CollectionResourceFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceFull(resourceOrn string, links CollectionResourceLinks) *CollectionResourceFull {
	this := CollectionResourceFull{}
	this.ResourceOrn = resourceOrn
	this.Links = links
	return &this
}

// NewCollectionResourceFullWithDefaults instantiates a new CollectionResourceFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceFullWithDefaults() *CollectionResourceFull {
	this := CollectionResourceFull{}
	return &this
}

// GetResourceProfile returns the ResourceProfile field value if set, zero value otherwise.
func (o *CollectionResourceFull) GetResourceProfile() ResourceProfile {
	if o == nil || IsNil(o.ResourceProfile) {
		var ret ResourceProfile
		return ret
	}
	return *o.ResourceProfile
}

// GetResourceProfileOk returns a tuple with the ResourceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetResourceProfileOk() (*ResourceProfile, bool) {
	if o == nil || IsNil(o.ResourceProfile) {
		return nil, false
	}
	return o.ResourceProfile, true
}

// HasResourceProfile returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasResourceProfile() bool {
	if o != nil && !IsNil(o.ResourceProfile) {
		return true
	}

	return false
}

// SetResourceProfile gets a reference to the given ResourceProfile and assigns it to the ResourceProfile field.
func (o *CollectionResourceFull) SetResourceProfile(v ResourceProfile) {
	o.ResourceProfile = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceFull) GetEntitlements() []EntitlementFull {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementFull
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetEntitlementsOk() ([]EntitlementFull, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementFull and assigns it to the Entitlements field.
func (o *CollectionResourceFull) SetEntitlements(v []EntitlementFull) {
	o.Entitlements = v
}

// GetEntitlementValueCount returns the EntitlementValueCount field value if set, zero value otherwise.
func (o *CollectionResourceFull) GetEntitlementValueCount() int32 {
	if o == nil || IsNil(o.EntitlementValueCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementValueCount
}

// GetEntitlementValueCountOk returns a tuple with the EntitlementValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetEntitlementValueCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitlementValueCount) {
		return nil, false
	}
	return o.EntitlementValueCount, true
}

// HasEntitlementValueCount returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasEntitlementValueCount() bool {
	if o != nil && !IsNil(o.EntitlementValueCount) {
		return true
	}

	return false
}

// SetEntitlementValueCount gets a reference to the given int32 and assigns it to the EntitlementValueCount field.
func (o *CollectionResourceFull) SetEntitlementValueCount(v int32) {
	o.EntitlementValueCount = &v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionResourceFull) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionResourceFull) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *CollectionResourceFull) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *CollectionResourceFull) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetLinks returns the Links field value
func (o *CollectionResourceFull) GetLinks() CollectionResourceLinks {
	if o == nil {
		var ret CollectionResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetLinksOk() (*CollectionResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CollectionResourceFull) SetLinks(v CollectionResourceLinks) {
	o.Links = v
}

func (o CollectionResourceFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceProfile) {
		toSerialize["resourceProfile"] = o.ResourceProfile
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.EntitlementValueCount) {
		toSerialize["entitlementValueCount"] = o.EntitlementValueCount
	}
	toSerialize["resourceOrn"] = o.ResourceOrn
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceOrn",
		"_links",
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

	varCollectionResourceFull := _CollectionResourceFull{}

	err = json.Unmarshal(data, &varCollectionResourceFull)

	if err != nil {
		return err
	}

	*o = CollectionResourceFull(varCollectionResourceFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceProfile")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "entitlementValueCount")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceFull struct {
	value *CollectionResourceFull
	isSet bool
}

func (v NullableCollectionResourceFull) Get() *CollectionResourceFull {
	return v.value
}

func (v *NullableCollectionResourceFull) Set(val *CollectionResourceFull) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceFull) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceFull(val *CollectionResourceFull) *NullableCollectionResourceFull {
	return &NullableCollectionResourceFull{value: val, isSet: true}
}

func (v NullableCollectionResourceFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
