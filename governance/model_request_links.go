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

// checks if the RequestLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestLinks{}

// RequestLinks Links available on a single request representation
type RequestLinks struct {
	RequestType          Link `json:"requestType"`
	Self                 Link `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _RequestLinks RequestLinks

// NewRequestLinks instantiates a new RequestLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestLinks(requestType Link, self Link) *RequestLinks {
	this := RequestLinks{}
	this.RequestType = requestType
	this.Self = self
	return &this
}

// NewRequestLinksWithDefaults instantiates a new RequestLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestLinksWithDefaults() *RequestLinks {
	this := RequestLinks{}
	return &this
}

// GetRequestType returns the RequestType field value
func (o *RequestLinks) GetRequestType() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *RequestLinks) GetRequestTypeOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *RequestLinks) SetRequestType(v Link) {
	o.RequestType = v
}

// GetSelf returns the Self field value
func (o *RequestLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RequestLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RequestLinks) SetSelf(v Link) {
	o.Self = v
}

func (o RequestLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestType"] = o.RequestType
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestType",
		"self",
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

	varRequestLinks := _RequestLinks{}

	err = json.Unmarshal(data, &varRequestLinks)

	if err != nil {
		return err
	}

	*o = RequestLinks(varRequestLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestType")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestLinks struct {
	value *RequestLinks
	isSet bool
}

func (v NullableRequestLinks) Get() *RequestLinks {
	return v.value
}

func (v *NullableRequestLinks) Set(val *RequestLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestLinks(val *RequestLinks) *NullableRequestLinks {
	return &NullableRequestLinks{value: val, isSet: true}
}

func (v NullableRequestLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
