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

// checks if the CollectionAssignmentOriginsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionAssignmentOriginsResponse{}

// CollectionAssignmentOriginsResponse Collection origin for each requested assignment
type CollectionAssignmentOriginsResponse struct {
	Results              []CollectionAssignmentOrigin `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionAssignmentOriginsResponse CollectionAssignmentOriginsResponse

// NewCollectionAssignmentOriginsResponse instantiates a new CollectionAssignmentOriginsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionAssignmentOriginsResponse() *CollectionAssignmentOriginsResponse {
	this := CollectionAssignmentOriginsResponse{}
	return &this
}

// NewCollectionAssignmentOriginsResponseWithDefaults instantiates a new CollectionAssignmentOriginsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionAssignmentOriginsResponseWithDefaults() *CollectionAssignmentOriginsResponse {
	this := CollectionAssignmentOriginsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CollectionAssignmentOriginsResponse) GetResults() []CollectionAssignmentOrigin {
	if o == nil || IsNil(o.Results) {
		var ret []CollectionAssignmentOrigin
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOriginsResponse) GetResultsOk() ([]CollectionAssignmentOrigin, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CollectionAssignmentOriginsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CollectionAssignmentOrigin and assigns it to the Results field.
func (o *CollectionAssignmentOriginsResponse) SetResults(v []CollectionAssignmentOrigin) {
	o.Results = v
}

func (o CollectionAssignmentOriginsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionAssignmentOriginsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionAssignmentOriginsResponse) UnmarshalJSON(data []byte) (err error) {
	varCollectionAssignmentOriginsResponse := _CollectionAssignmentOriginsResponse{}

	err = json.Unmarshal(data, &varCollectionAssignmentOriginsResponse)

	if err != nil {
		return err
	}

	*o = CollectionAssignmentOriginsResponse(varCollectionAssignmentOriginsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionAssignmentOriginsResponse struct {
	value *CollectionAssignmentOriginsResponse
	isSet bool
}

func (v NullableCollectionAssignmentOriginsResponse) Get() *CollectionAssignmentOriginsResponse {
	return v.value
}

func (v *NullableCollectionAssignmentOriginsResponse) Set(val *CollectionAssignmentOriginsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionAssignmentOriginsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionAssignmentOriginsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionAssignmentOriginsResponse(val *CollectionAssignmentOriginsResponse) *NullableCollectionAssignmentOriginsResponse {
	return &NullableCollectionAssignmentOriginsResponse{value: val, isSet: true}
}

func (v NullableCollectionAssignmentOriginsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionAssignmentOriginsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
