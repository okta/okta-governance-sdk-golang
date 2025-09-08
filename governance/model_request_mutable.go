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

// RequestMutable Properties which are mutable in mutation operations
type RequestMutable struct {
	// The request type `id`.
	RequestTypeId *string `json:"requestTypeId,omitempty" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	// The subject of the request
	Subject *string `json:"subject,omitempty"`
	// A list of requester Okta user `id`s.
	RequesterUserIds     []string `json:"requesterUserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestMutable RequestMutable

// NewRequestMutable instantiates a new RequestMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMutable() *RequestMutable {
	this := RequestMutable{}
	return &this
}

// NewRequestMutableWithDefaults instantiates a new RequestMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMutableWithDefaults() *RequestMutable {
	this := RequestMutable{}
	return &this
}

// GetRequestTypeId returns the RequestTypeId field value if set, zero value otherwise.
func (o *RequestMutable) GetRequestTypeId() string {
	if o == nil || o.RequestTypeId == nil {
		var ret string
		return ret
	}
	return *o.RequestTypeId
}

// GetRequestTypeIdOk returns a tuple with the RequestTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMutable) GetRequestTypeIdOk() (*string, bool) {
	if o == nil || o.RequestTypeId == nil {
		return nil, false
	}
	return o.RequestTypeId, true
}

// HasRequestTypeId returns a boolean if a field has been set.
func (o *RequestMutable) HasRequestTypeId() bool {
	if o != nil && o.RequestTypeId != nil {
		return true
	}

	return false
}

// SetRequestTypeId gets a reference to the given string and assigns it to the RequestTypeId field.
func (o *RequestMutable) SetRequestTypeId(v string) {
	o.RequestTypeId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *RequestMutable) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMutable) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *RequestMutable) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *RequestMutable) SetSubject(v string) {
	o.Subject = &v
}

// GetRequesterUserIds returns the RequesterUserIds field value if set, zero value otherwise.
func (o *RequestMutable) GetRequesterUserIds() []string {
	if o == nil || o.RequesterUserIds == nil {
		var ret []string
		return ret
	}
	return o.RequesterUserIds
}

// GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMutable) GetRequesterUserIdsOk() ([]string, bool) {
	if o == nil || o.RequesterUserIds == nil {
		return nil, false
	}
	return o.RequesterUserIds, true
}

// HasRequesterUserIds returns a boolean if a field has been set.
func (o *RequestMutable) HasRequesterUserIds() bool {
	if o != nil && o.RequesterUserIds != nil {
		return true
	}

	return false
}

// SetRequesterUserIds gets a reference to the given []string and assigns it to the RequesterUserIds field.
func (o *RequestMutable) SetRequesterUserIds(v []string) {
	o.RequesterUserIds = v
}

func (o RequestMutable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestTypeId != nil {
		toSerialize["requestTypeId"] = o.RequestTypeId
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.RequesterUserIds != nil {
		toSerialize["requesterUserIds"] = o.RequesterUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestMutable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestMutable := _RequestMutable{}

	err = json.Unmarshal(bytes, &varRequestMutable)
	if err == nil {
		*o = RequestMutable(varRequestMutable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestTypeId")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "requesterUserIds")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestMutable struct {
	value *RequestMutable
	isSet bool
}

func (v NullableRequestMutable) Get() *RequestMutable {
	return v.value
}

func (v *NullableRequestMutable) Set(val *RequestMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMutable(val *RequestMutable) *NullableRequestMutable {
	return &NullableRequestMutable{value: val, isSet: true}
}

func (v NullableRequestMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
