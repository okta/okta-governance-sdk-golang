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

// RequestCreatable The properties expected in an initial create request
type RequestCreatable struct {
	// The request type `id`.
	RequestTypeId string `json:"requestTypeId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	// The subject of the request
	Subject string `json:"subject"`
	// A list of requester Okta user `id`s.
	RequesterUserIds []string `json:"requesterUserIds,omitempty"`
	// Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.
	RequesterFieldValues []FieldValueWritable `json:"requesterFieldValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestCreatable RequestCreatable

// NewRequestCreatable instantiates a new RequestCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestCreatable(requestTypeId string, subject string) *RequestCreatable {
	this := RequestCreatable{}
	this.RequestTypeId = requestTypeId
	this.Subject = subject
	return &this
}

// NewRequestCreatableWithDefaults instantiates a new RequestCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestCreatableWithDefaults() *RequestCreatable {
	this := RequestCreatable{}
	return &this
}

// GetRequestTypeId returns the RequestTypeId field value
func (o *RequestCreatable) GetRequestTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestTypeId
}

// GetRequestTypeIdOk returns a tuple with the RequestTypeId field value
// and a boolean to check if the value has been set.
func (o *RequestCreatable) GetRequestTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTypeId, true
}

// SetRequestTypeId sets field value
func (o *RequestCreatable) SetRequestTypeId(v string) {
	o.RequestTypeId = v
}

// GetSubject returns the Subject field value
func (o *RequestCreatable) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *RequestCreatable) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *RequestCreatable) SetSubject(v string) {
	o.Subject = v
}

// GetRequesterUserIds returns the RequesterUserIds field value if set, zero value otherwise.
func (o *RequestCreatable) GetRequesterUserIds() []string {
	if o == nil || o.RequesterUserIds == nil {
		var ret []string
		return ret
	}
	return o.RequesterUserIds
}

// GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreatable) GetRequesterUserIdsOk() ([]string, bool) {
	if o == nil || o.RequesterUserIds == nil {
		return nil, false
	}
	return o.RequesterUserIds, true
}

// HasRequesterUserIds returns a boolean if a field has been set.
func (o *RequestCreatable) HasRequesterUserIds() bool {
	if o != nil && o.RequesterUserIds != nil {
		return true
	}

	return false
}

// SetRequesterUserIds gets a reference to the given []string and assigns it to the RequesterUserIds field.
func (o *RequestCreatable) SetRequesterUserIds(v []string) {
	o.RequesterUserIds = v
}

// GetRequesterFieldValues returns the RequesterFieldValues field value if set, zero value otherwise.
func (o *RequestCreatable) GetRequesterFieldValues() []FieldValueWritable {
	if o == nil || o.RequesterFieldValues == nil {
		var ret []FieldValueWritable
		return ret
	}
	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreatable) GetRequesterFieldValuesOk() ([]FieldValueWritable, bool) {
	if o == nil || o.RequesterFieldValues == nil {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// HasRequesterFieldValues returns a boolean if a field has been set.
func (o *RequestCreatable) HasRequesterFieldValues() bool {
	if o != nil && o.RequesterFieldValues != nil {
		return true
	}

	return false
}

// SetRequesterFieldValues gets a reference to the given []FieldValueWritable and assigns it to the RequesterFieldValues field.
func (o *RequestCreatable) SetRequesterFieldValues(v []FieldValueWritable) {
	o.RequesterFieldValues = v
}

func (o RequestCreatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestTypeId"] = o.RequestTypeId
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if o.RequesterUserIds != nil {
		toSerialize["requesterUserIds"] = o.RequesterUserIds
	}
	if o.RequesterFieldValues != nil {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestCreatable := _RequestCreatable{}

	err = json.Unmarshal(bytes, &varRequestCreatable)
	if err == nil {
		*o = RequestCreatable(varRequestCreatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestTypeId")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "requesterUserIds")
		delete(additionalProperties, "requesterFieldValues")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestCreatable struct {
	value *RequestCreatable
	isSet bool
}

func (v NullableRequestCreatable) Get() *RequestCreatable {
	return v.value
}

func (v *NullableRequestCreatable) Set(val *RequestCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCreatable(val *RequestCreatable) *NullableRequestCreatable {
	return &NullableRequestCreatable{value: val, isSet: true}
}

func (v NullableRequestCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
