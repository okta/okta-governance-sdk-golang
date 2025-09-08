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

// RequestTypeLinks Links available on a single request type representation
type RequestTypeLinks struct {
	// Only available for request type with status of ACTIVE
	CreateRequest        *Link `json:"createRequest,omitempty"`
	Requests             Link  `json:"requests"`
	Self                 Link  `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeLinks RequestTypeLinks

// NewRequestTypeLinks instantiates a new RequestTypeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeLinks(requests Link, self Link) *RequestTypeLinks {
	this := RequestTypeLinks{}
	this.Requests = requests
	this.Self = self
	return &this
}

// NewRequestTypeLinksWithDefaults instantiates a new RequestTypeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeLinksWithDefaults() *RequestTypeLinks {
	this := RequestTypeLinks{}
	return &this
}

// GetCreateRequest returns the CreateRequest field value if set, zero value otherwise.
func (o *RequestTypeLinks) GetCreateRequest() Link {
	if o == nil || o.CreateRequest == nil {
		var ret Link
		return ret
	}
	return *o.CreateRequest
}

// GetCreateRequestOk returns a tuple with the CreateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeLinks) GetCreateRequestOk() (*Link, bool) {
	if o == nil || o.CreateRequest == nil {
		return nil, false
	}
	return o.CreateRequest, true
}

// HasCreateRequest returns a boolean if a field has been set.
func (o *RequestTypeLinks) HasCreateRequest() bool {
	if o != nil && o.CreateRequest != nil {
		return true
	}

	return false
}

// SetCreateRequest gets a reference to the given Link and assigns it to the CreateRequest field.
func (o *RequestTypeLinks) SetCreateRequest(v Link) {
	o.CreateRequest = &v
}

// GetRequests returns the Requests field value
func (o *RequestTypeLinks) GetRequests() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *RequestTypeLinks) GetRequestsOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *RequestTypeLinks) SetRequests(v Link) {
	o.Requests = v
}

// GetSelf returns the Self field value
func (o *RequestTypeLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RequestTypeLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RequestTypeLinks) SetSelf(v Link) {
	o.Self = v
}

func (o RequestTypeLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateRequest != nil {
		toSerialize["createRequest"] = o.CreateRequest
	}
	if true {
		toSerialize["requests"] = o.Requests
	}
	if true {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeLinks) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeLinks := _RequestTypeLinks{}

	err = json.Unmarshal(bytes, &varRequestTypeLinks)
	if err == nil {
		*o = RequestTypeLinks(varRequestTypeLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "createRequest")
		delete(additionalProperties, "requests")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeLinks struct {
	value *RequestTypeLinks
	isSet bool
}

func (v NullableRequestTypeLinks) Get() *RequestTypeLinks {
	return v.value
}

func (v *NullableRequestTypeLinks) Set(val *RequestTypeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeLinks(val *RequestTypeLinks) *NullableRequestTypeLinks {
	return &NullableRequestTypeLinks{value: val, isSet: true}
}

func (v NullableRequestTypeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
