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

// ResourceOwnersCatalogResourcesResponse struct for ResourceOwnersCatalogResourcesResponse
type ResourceOwnersCatalogResourcesResponse struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn string `json:"parentResourceOrn"`
	// Resource owner details.
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
	if o == nil || o.Data == nil {
		var ret []ResourceOwnerResource
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersCatalogResourcesResponse) GetDataOk() ([]ResourceOwnerResource, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResourceOwnersCatalogResourcesResponse) HasData() bool {
	if o != nil && o.Data != nil {
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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceOwnersCatalogResourcesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varResourceOwnersCatalogResourcesResponse := _ResourceOwnersCatalogResourcesResponse{}

	err = json.Unmarshal(bytes, &varResourceOwnersCatalogResourcesResponse)
	if err == nil {
		*o = ResourceOwnersCatalogResourcesResponse(varResourceOwnersCatalogResourcesResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
