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

// checks if the ResourceOwnersListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOwnersListResponse{}

// ResourceOwnersListResponse struct for ResourceOwnersListResponse
type ResourceOwnersListResponse struct {
	// Resource owner details.
	Data                 []ResourceOwner          `json:"data,omitempty"`
	Links                *ResourceOwnersListLinks `json:"_links,omitempty"`
	Metadata             *ListMetadata            `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnersListResponse ResourceOwnersListResponse

// NewResourceOwnersListResponse instantiates a new ResourceOwnersListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnersListResponse() *ResourceOwnersListResponse {
	this := ResourceOwnersListResponse{}
	return &this
}

// NewResourceOwnersListResponseWithDefaults instantiates a new ResourceOwnersListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnersListResponseWithDefaults() *ResourceOwnersListResponse {
	this := ResourceOwnersListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResourceOwnersListResponse) GetData() []ResourceOwner {
	if o == nil || IsNil(o.Data) {
		var ret []ResourceOwner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersListResponse) GetDataOk() ([]ResourceOwner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResourceOwnersListResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResourceOwner and assigns it to the Data field.
func (o *ResourceOwnersListResponse) SetData(v []ResourceOwner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceOwnersListResponse) GetLinks() ResourceOwnersListLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceOwnersListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersListResponse) GetLinksOk() (*ResourceOwnersListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceOwnersListResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceOwnersListLinks and assigns it to the Links field.
func (o *ResourceOwnersListResponse) SetLinks(v ResourceOwnersListLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ResourceOwnersListResponse) GetMetadata() ListMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersListResponse) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ResourceOwnersListResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *ResourceOwnersListResponse) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

func (o ResourceOwnersListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOwnersListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceOwnersListResponse) UnmarshalJSON(data []byte) (err error) {
	varResourceOwnersListResponse := _ResourceOwnersListResponse{}

	err = json.Unmarshal(data, &varResourceOwnersListResponse)

	if err != nil {
		return err
	}

	*o = ResourceOwnersListResponse(varResourceOwnersListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceOwnersListResponse struct {
	value *ResourceOwnersListResponse
	isSet bool
}

func (v NullableResourceOwnersListResponse) Get() *ResourceOwnersListResponse {
	return v.value
}

func (v *NullableResourceOwnersListResponse) Set(val *ResourceOwnersListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnersListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnersListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnersListResponse(val *ResourceOwnersListResponse) *NullableResourceOwnersListResponse {
	return &NullableResourceOwnersListResponse{value: val, isSet: true}
}

func (v NullableResourceOwnersListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnersListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
