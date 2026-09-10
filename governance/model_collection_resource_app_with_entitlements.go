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

// checks if the CollectionResourceAppWithEntitlements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceAppWithEntitlements{}

// CollectionResourceAppWithEntitlements Representation of an EM-enabled app resource after updating entitlements
type CollectionResourceAppWithEntitlements struct {
	ResourceConfiguration *CollectionResourceConfiguration `json:"resourceConfiguration,omitempty"`
	// Collection of entitlements with associated values
	Entitlements []EntitlementFull `json:"entitlements,omitempty"`
	// The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.
	ResourceOrn string `json:"resourceOrn"`
	// The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as `GET`/`PUT`/`DELETE /v2/collections/{collectionId}/resources/{resourceId}`.
	ResourceId           *string                   `json:"resourceId,omitempty"`
	Links                CollectionResourceLinksV2 `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceAppWithEntitlements CollectionResourceAppWithEntitlements

// NewCollectionResourceAppWithEntitlements instantiates a new CollectionResourceAppWithEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceAppWithEntitlements(resourceOrn string, links CollectionResourceLinksV2) *CollectionResourceAppWithEntitlements {
	this := CollectionResourceAppWithEntitlements{}
	this.ResourceOrn = resourceOrn
	this.Links = links
	return &this
}

// NewCollectionResourceAppWithEntitlementsWithDefaults instantiates a new CollectionResourceAppWithEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceAppWithEntitlementsWithDefaults() *CollectionResourceAppWithEntitlements {
	this := CollectionResourceAppWithEntitlements{}
	return &this
}

// GetResourceConfiguration returns the ResourceConfiguration field value if set, zero value otherwise.
func (o *CollectionResourceAppWithEntitlements) GetResourceConfiguration() CollectionResourceConfiguration {
	if o == nil || IsNil(o.ResourceConfiguration) {
		var ret CollectionResourceConfiguration
		return ret
	}
	return *o.ResourceConfiguration
}

// GetResourceConfigurationOk returns a tuple with the ResourceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceAppWithEntitlements) GetResourceConfigurationOk() (*CollectionResourceConfiguration, bool) {
	if o == nil || IsNil(o.ResourceConfiguration) {
		return nil, false
	}
	return o.ResourceConfiguration, true
}

// HasResourceConfiguration returns a boolean if a field has been set.
func (o *CollectionResourceAppWithEntitlements) HasResourceConfiguration() bool {
	if o != nil && !IsNil(o.ResourceConfiguration) {
		return true
	}

	return false
}

// SetResourceConfiguration gets a reference to the given CollectionResourceConfiguration and assigns it to the ResourceConfiguration field.
func (o *CollectionResourceAppWithEntitlements) SetResourceConfiguration(v CollectionResourceConfiguration) {
	o.ResourceConfiguration = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceAppWithEntitlements) GetEntitlements() []EntitlementFull {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementFull
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceAppWithEntitlements) GetEntitlementsOk() ([]EntitlementFull, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceAppWithEntitlements) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementFull and assigns it to the Entitlements field.
func (o *CollectionResourceAppWithEntitlements) SetEntitlements(v []EntitlementFull) {
	o.Entitlements = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionResourceAppWithEntitlements) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceAppWithEntitlements) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionResourceAppWithEntitlements) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *CollectionResourceAppWithEntitlements) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceAppWithEntitlements) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CollectionResourceAppWithEntitlements) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *CollectionResourceAppWithEntitlements) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetLinks returns the Links field value
func (o *CollectionResourceAppWithEntitlements) GetLinks() CollectionResourceLinksV2 {
	if o == nil {
		var ret CollectionResourceLinksV2
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceAppWithEntitlements) GetLinksOk() (*CollectionResourceLinksV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CollectionResourceAppWithEntitlements) SetLinks(v CollectionResourceLinksV2) {
	o.Links = v
}

func (o CollectionResourceAppWithEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceAppWithEntitlements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceConfiguration) {
		toSerialize["resourceConfiguration"] = o.ResourceConfiguration
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
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

func (o *CollectionResourceAppWithEntitlements) UnmarshalJSON(data []byte) (err error) {
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

	varCollectionResourceAppWithEntitlements := _CollectionResourceAppWithEntitlements{}

	err = json.Unmarshal(data, &varCollectionResourceAppWithEntitlements)

	if err != nil {
		return err
	}

	*o = CollectionResourceAppWithEntitlements(varCollectionResourceAppWithEntitlements)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceConfiguration")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceAppWithEntitlements struct {
	value *CollectionResourceAppWithEntitlements
	isSet bool
}

func (v NullableCollectionResourceAppWithEntitlements) Get() *CollectionResourceAppWithEntitlements {
	return v.value
}

func (v *NullableCollectionResourceAppWithEntitlements) Set(val *CollectionResourceAppWithEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceAppWithEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceAppWithEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceAppWithEntitlements(val *CollectionResourceAppWithEntitlements) *NullableCollectionResourceAppWithEntitlements {
	return &NullableCollectionResourceAppWithEntitlements{value: val, isSet: true}
}

func (v NullableCollectionResourceAppWithEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceAppWithEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
