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

// checks if the ResourceOwnersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOwnersResponse{}

// ResourceOwnersResponse struct for ResourceOwnersResponse
type ResourceOwnersResponse struct {
	// Resource owner details.
	Data                 []ResourceOwner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnersResponse ResourceOwnersResponse

// NewResourceOwnersResponse instantiates a new ResourceOwnersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnersResponse() *ResourceOwnersResponse {
	this := ResourceOwnersResponse{}
	return &this
}

// NewResourceOwnersResponseWithDefaults instantiates a new ResourceOwnersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnersResponseWithDefaults() *ResourceOwnersResponse {
	this := ResourceOwnersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResourceOwnersResponse) GetData() []ResourceOwner {
	if o == nil || IsNil(o.Data) {
		var ret []ResourceOwner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersResponse) GetDataOk() ([]ResourceOwner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResourceOwnersResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResourceOwner and assigns it to the Data field.
func (o *ResourceOwnersResponse) SetData(v []ResourceOwner) {
	o.Data = v
}

func (o ResourceOwnersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOwnersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceOwnersResponse) UnmarshalJSON(data []byte) (err error) {
	varResourceOwnersResponse := _ResourceOwnersResponse{}

	err = json.Unmarshal(data, &varResourceOwnersResponse)

	if err != nil {
		return err
	}

	*o = ResourceOwnersResponse(varResourceOwnersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceOwnersResponse struct {
	value *ResourceOwnersResponse
	isSet bool
}

func (v NullableResourceOwnersResponse) Get() *ResourceOwnersResponse {
	return v.value
}

func (v *NullableResourceOwnersResponse) Set(val *ResourceOwnersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnersResponse(val *ResourceOwnersResponse) *NullableResourceOwnersResponse {
	return &NullableResourceOwnersResponse{value: val, isSet: true}
}

func (v NullableResourceOwnersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
