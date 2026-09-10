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

// checks if the RequestApprovalInsightsGenerateSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestApprovalInsightsGenerateSuccess{}

// RequestApprovalInsightsGenerateSuccess The result of generating approval insights for a request
type RequestApprovalInsightsGenerateSuccess struct {
	Links                RequestApprovalInsightsLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _RequestApprovalInsightsGenerateSuccess RequestApprovalInsightsGenerateSuccess

// NewRequestApprovalInsightsGenerateSuccess instantiates a new RequestApprovalInsightsGenerateSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestApprovalInsightsGenerateSuccess(links RequestApprovalInsightsLinks) *RequestApprovalInsightsGenerateSuccess {
	this := RequestApprovalInsightsGenerateSuccess{}
	this.Links = links
	return &this
}

// NewRequestApprovalInsightsGenerateSuccessWithDefaults instantiates a new RequestApprovalInsightsGenerateSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestApprovalInsightsGenerateSuccessWithDefaults() *RequestApprovalInsightsGenerateSuccess {
	this := RequestApprovalInsightsGenerateSuccess{}
	return &this
}

// GetLinks returns the Links field value
func (o *RequestApprovalInsightsGenerateSuccess) GetLinks() RequestApprovalInsightsLinks {
	if o == nil {
		var ret RequestApprovalInsightsLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalInsightsGenerateSuccess) GetLinksOk() (*RequestApprovalInsightsLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestApprovalInsightsGenerateSuccess) SetLinks(v RequestApprovalInsightsLinks) {
	o.Links = v
}

func (o RequestApprovalInsightsGenerateSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestApprovalInsightsGenerateSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestApprovalInsightsGenerateSuccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRequestApprovalInsightsGenerateSuccess := _RequestApprovalInsightsGenerateSuccess{}

	err = json.Unmarshal(data, &varRequestApprovalInsightsGenerateSuccess)

	if err != nil {
		return err
	}

	*o = RequestApprovalInsightsGenerateSuccess(varRequestApprovalInsightsGenerateSuccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestApprovalInsightsGenerateSuccess struct {
	value *RequestApprovalInsightsGenerateSuccess
	isSet bool
}

func (v NullableRequestApprovalInsightsGenerateSuccess) Get() *RequestApprovalInsightsGenerateSuccess {
	return v.value
}

func (v *NullableRequestApprovalInsightsGenerateSuccess) Set(val *RequestApprovalInsightsGenerateSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprovalInsightsGenerateSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprovalInsightsGenerateSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprovalInsightsGenerateSuccess(val *RequestApprovalInsightsGenerateSuccess) *NullableRequestApprovalInsightsGenerateSuccess {
	return &NullableRequestApprovalInsightsGenerateSuccess{value: val, isSet: true}
}

func (v NullableRequestApprovalInsightsGenerateSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprovalInsightsGenerateSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
