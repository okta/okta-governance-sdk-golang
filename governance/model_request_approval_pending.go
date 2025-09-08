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

// RequestApprovalPending A pending access request approval
type RequestApprovalPending struct {
	// The status of a pending approval
	Status string `json:"status"`
	// A unique identifier of the approval in the request
	ApprovalId           string `json:"approvalId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	AdditionalProperties map[string]interface{}
}

type _RequestApprovalPending RequestApprovalPending

// NewRequestApprovalPending instantiates a new RequestApprovalPending object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestApprovalPending(status string, approvalId string) *RequestApprovalPending {
	this := RequestApprovalPending{}
	this.Status = status
	this.ApprovalId = approvalId
	return &this
}

// NewRequestApprovalPendingWithDefaults instantiates a new RequestApprovalPending object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestApprovalPendingWithDefaults() *RequestApprovalPending {
	this := RequestApprovalPending{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestApprovalPending) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalPending) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestApprovalPending) SetStatus(v string) {
	o.Status = v
}

// GetApprovalId returns the ApprovalId field value
func (o *RequestApprovalPending) GetApprovalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApprovalId
}

// GetApprovalIdOk returns a tuple with the ApprovalId field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalPending) GetApprovalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalId, true
}

// SetApprovalId sets field value
func (o *RequestApprovalPending) SetApprovalId(v string) {
	o.ApprovalId = v
}

func (o RequestApprovalPending) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["approvalId"] = o.ApprovalId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestApprovalPending) UnmarshalJSON(bytes []byte) (err error) {
	varRequestApprovalPending := _RequestApprovalPending{}

	err = json.Unmarshal(bytes, &varRequestApprovalPending)
	if err == nil {
		*o = RequestApprovalPending(varRequestApprovalPending)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "approvalId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestApprovalPending struct {
	value *RequestApprovalPending
	isSet bool
}

func (v NullableRequestApprovalPending) Get() *RequestApprovalPending {
	return v.value
}

func (v *NullableRequestApprovalPending) Set(val *RequestApprovalPending) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprovalPending) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprovalPending) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprovalPending(val *RequestApprovalPending) *NullableRequestApprovalPending {
	return &NullableRequestApprovalPending{value: val, isSet: true}
}

func (v NullableRequestApprovalPending) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprovalPending) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
