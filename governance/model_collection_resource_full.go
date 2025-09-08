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
	if o == nil || o.ResourceProfile == nil {
		var ret ResourceProfile
		return ret
	}
	return *o.ResourceProfile
}

// GetResourceProfileOk returns a tuple with the ResourceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetResourceProfileOk() (*ResourceProfile, bool) {
	if o == nil || o.ResourceProfile == nil {
		return nil, false
	}
	return o.ResourceProfile, true
}

// HasResourceProfile returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasResourceProfile() bool {
	if o != nil && o.ResourceProfile != nil {
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
	if o == nil || o.Entitlements == nil {
		var ret []EntitlementFull
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetEntitlementsOk() ([]EntitlementFull, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
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
	if o == nil || o.EntitlementValueCount == nil {
		var ret int32
		return ret
	}
	return *o.EntitlementValueCount
}

// GetEntitlementValueCountOk returns a tuple with the EntitlementValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetEntitlementValueCountOk() (*int32, bool) {
	if o == nil || o.EntitlementValueCount == nil {
		return nil, false
	}
	return o.EntitlementValueCount, true
}

// HasEntitlementValueCount returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasEntitlementValueCount() bool {
	if o != nil && o.EntitlementValueCount != nil {
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
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFull) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CollectionResourceFull) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
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
	toSerialize := map[string]interface{}{}
	if o.ResourceProfile != nil {
		toSerialize["resourceProfile"] = o.ResourceProfile
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if o.EntitlementValueCount != nil {
		toSerialize["entitlementValueCount"] = o.EntitlementValueCount
	}
	if true {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CollectionResourceFull) UnmarshalJSON(bytes []byte) (err error) {
	varCollectionResourceFull := _CollectionResourceFull{}

	err = json.Unmarshal(bytes, &varCollectionResourceFull)
	if err == nil {
		*o = CollectionResourceFull(varCollectionResourceFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceProfile")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "entitlementValueCount")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
