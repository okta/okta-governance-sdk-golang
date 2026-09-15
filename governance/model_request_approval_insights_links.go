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

// checks if the RequestApprovalInsightsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestApprovalInsightsLinks{}

// RequestApprovalInsightsLinks Links available on a generate request approval insights response
type RequestApprovalInsightsLinks struct {
	Self                 Link `json:"self"`
	Request              Link `json:"request"`
	AdditionalProperties map[string]interface{}
}

type _RequestApprovalInsightsLinks RequestApprovalInsightsLinks

// NewRequestApprovalInsightsLinks instantiates a new RequestApprovalInsightsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestApprovalInsightsLinks(self Link, request Link) *RequestApprovalInsightsLinks {
	this := RequestApprovalInsightsLinks{}
	this.Self = self
	this.Request = request
	return &this
}

// NewRequestApprovalInsightsLinksWithDefaults instantiates a new RequestApprovalInsightsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestApprovalInsightsLinksWithDefaults() *RequestApprovalInsightsLinks {
	this := RequestApprovalInsightsLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *RequestApprovalInsightsLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalInsightsLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RequestApprovalInsightsLinks) SetSelf(v Link) {
	o.Self = v
}

// GetRequest returns the Request field value
func (o *RequestApprovalInsightsLinks) GetRequest() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalInsightsLinks) GetRequestOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *RequestApprovalInsightsLinks) SetRequest(v Link) {
	o.Request = v
}

func (o RequestApprovalInsightsLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestApprovalInsightsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["request"] = o.Request

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestApprovalInsightsLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"request",
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

	varRequestApprovalInsightsLinks := _RequestApprovalInsightsLinks{}

	err = json.Unmarshal(data, &varRequestApprovalInsightsLinks)

	if err != nil {
		return err
	}

	*o = RequestApprovalInsightsLinks(varRequestApprovalInsightsLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "request")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestApprovalInsightsLinks struct {
	value *RequestApprovalInsightsLinks
	isSet bool
}

func (v NullableRequestApprovalInsightsLinks) Get() *RequestApprovalInsightsLinks {
	return v.value
}

func (v *NullableRequestApprovalInsightsLinks) Set(val *RequestApprovalInsightsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprovalInsightsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprovalInsightsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprovalInsightsLinks(val *RequestApprovalInsightsLinks) *NullableRequestApprovalInsightsLinks {
	return &NullableRequestApprovalInsightsLinks{value: val, isSet: true}
}

func (v NullableRequestApprovalInsightsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprovalInsightsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
