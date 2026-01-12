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

// checks if the RevokePrincipalAccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokePrincipalAccessResponse{}

// RevokePrincipalAccessResponse Revoke principal access response
type RevokePrincipalAccessResponse struct {
	// List of related links for each revoked-access resource
	Data                 []RevokePrincipalAccessResourceLinks `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RevokePrincipalAccessResponse RevokePrincipalAccessResponse

// NewRevokePrincipalAccessResponse instantiates a new RevokePrincipalAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokePrincipalAccessResponse() *RevokePrincipalAccessResponse {
	this := RevokePrincipalAccessResponse{}
	return &this
}

// NewRevokePrincipalAccessResponseWithDefaults instantiates a new RevokePrincipalAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrincipalAccessResponseWithDefaults() *RevokePrincipalAccessResponse {
	this := RevokePrincipalAccessResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RevokePrincipalAccessResponse) GetData() []RevokePrincipalAccessResourceLinks {
	if o == nil || IsNil(o.Data) {
		var ret []RevokePrincipalAccessResourceLinks
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessResponse) GetDataOk() ([]RevokePrincipalAccessResourceLinks, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RevokePrincipalAccessResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RevokePrincipalAccessResourceLinks and assigns it to the Data field.
func (o *RevokePrincipalAccessResponse) SetData(v []RevokePrincipalAccessResourceLinks) {
	o.Data = v
}

func (o RevokePrincipalAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokePrincipalAccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokePrincipalAccessResponse) UnmarshalJSON(data []byte) (err error) {
	varRevokePrincipalAccessResponse := _RevokePrincipalAccessResponse{}

	err = json.Unmarshal(data, &varRevokePrincipalAccessResponse)

	if err != nil {
		return err
	}

	*o = RevokePrincipalAccessResponse(varRevokePrincipalAccessResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokePrincipalAccessResponse struct {
	value *RevokePrincipalAccessResponse
	isSet bool
}

func (v NullableRevokePrincipalAccessResponse) Get() *RevokePrincipalAccessResponse {
	return v.value
}

func (v *NullableRevokePrincipalAccessResponse) Set(val *RevokePrincipalAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokePrincipalAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokePrincipalAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokePrincipalAccessResponse(val *RevokePrincipalAccessResponse) *NullableRevokePrincipalAccessResponse {
	return &NullableRevokePrincipalAccessResponse{value: val, isSet: true}
}

func (v NullableRevokePrincipalAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokePrincipalAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
