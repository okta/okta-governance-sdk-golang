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

// checks if the ResourceOwnersCatalogResourcesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOwnersCatalogResourcesResponse{}

// ResourceOwnersCatalogResourcesResponse struct for ResourceOwnersCatalogResourcesResponse
type ResourceOwnersCatalogResourcesResponse struct {
	// The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn string `json:"parentResourceOrn"`
	// Resource owner details
	Data                 []ResourceOwnerResource `json:"data,omitempty"`
	Links                ResourceOwnersListLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnersCatalogResourcesResponse ResourceOwnersCatalogResourcesResponse

// NewResourceOwnersCatalogResourcesResponse instantiates a new ResourceOwnersCatalogResourcesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnersCatalogResourcesResponse(parentResourceOrn string, links ResourceOwnersListLinks) *ResourceOwnersCatalogResourcesResponse {
	this := ResourceOwnersCatalogResourcesResponse{}
	this.ParentResourceOrn = parentResourceOrn
	this.Links = links
	return &this
}

// NewResourceOwnersCatalogResourcesResponseWithDefaults instantiates a new ResourceOwnersCatalogResourcesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnersCatalogResourcesResponseWithDefaults() *ResourceOwnersCatalogResourcesResponse {
	this := ResourceOwnersCatalogResourcesResponse{}
	return &this
}

// GetParentResourceOrn returns the ParentResourceOrn field value
func (o *ResourceOwnersCatalogResourcesResponse) GetParentResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnersCatalogResourcesResponse) GetParentResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentResourceOrn, true
}

// SetParentResourceOrn sets field value
func (o *ResourceOwnersCatalogResourcesResponse) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResourceOwnersCatalogResourcesResponse) GetData() []ResourceOwnerResource {
	if o == nil || IsNil(o.Data) {
		var ret []ResourceOwnerResource
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersCatalogResourcesResponse) GetDataOk() ([]ResourceOwnerResource, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResourceOwnersCatalogResourcesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResourceOwnerResource and assigns it to the Data field.
func (o *ResourceOwnersCatalogResourcesResponse) SetData(v []ResourceOwnerResource) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ResourceOwnersCatalogResourcesResponse) GetLinks() ResourceOwnersListLinks {
	if o == nil {
		var ret ResourceOwnersListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnersCatalogResourcesResponse) GetLinksOk() (*ResourceOwnersListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ResourceOwnersCatalogResourcesResponse) SetLinks(v ResourceOwnersListLinks) {
	o.Links = v
}

func (o ResourceOwnersCatalogResourcesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOwnersCatalogResourcesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceOwnersCatalogResourcesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parentResourceOrn",
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

	varResourceOwnersCatalogResourcesResponse := _ResourceOwnersCatalogResourcesResponse{}

	err = json.Unmarshal(data, &varResourceOwnersCatalogResourcesResponse)

	if err != nil {
		return err
	}

	*o = ResourceOwnersCatalogResourcesResponse(varResourceOwnersCatalogResourcesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceOwnersCatalogResourcesResponse struct {
	value *ResourceOwnersCatalogResourcesResponse
	isSet bool
}

func (v NullableResourceOwnersCatalogResourcesResponse) Get() *ResourceOwnersCatalogResourcesResponse {
	return v.value
}

func (v *NullableResourceOwnersCatalogResourcesResponse) Set(val *ResourceOwnersCatalogResourcesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnersCatalogResourcesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnersCatalogResourcesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnersCatalogResourcesResponse(val *ResourceOwnersCatalogResourcesResponse) *NullableResourceOwnersCatalogResourcesResponse {
	return &NullableResourceOwnersCatalogResourcesResponse{value: val, isSet: true}
}

func (v NullableResourceOwnersCatalogResourcesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnersCatalogResourcesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
