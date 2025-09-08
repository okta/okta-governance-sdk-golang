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

// RequestTypeApprovalResourceOwnerWritable Specifying the approver must be the resource owner
type RequestTypeApprovalResourceOwnerWritable struct {
	ApproverType         string  `json:"approverType"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalResourceOwnerWritable RequestTypeApprovalResourceOwnerWritable

// NewRequestTypeApprovalResourceOwnerWritable instantiates a new RequestTypeApprovalResourceOwnerWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalResourceOwnerWritable(approverType string) *RequestTypeApprovalResourceOwnerWritable {
	this := RequestTypeApprovalResourceOwnerWritable{}
	this.ApproverType = approverType
	return &this
}

// NewRequestTypeApprovalResourceOwnerWritableWithDefaults instantiates a new RequestTypeApprovalResourceOwnerWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalResourceOwnerWritableWithDefaults() *RequestTypeApprovalResourceOwnerWritable {
	this := RequestTypeApprovalResourceOwnerWritable{}
	return &this
}

// GetApproverType returns the ApproverType field value
func (o *RequestTypeApprovalResourceOwnerWritable) GetApproverType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverType
}

// GetApproverTypeOk returns a tuple with the ApproverType field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalResourceOwnerWritable) GetApproverTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverType, true
}

// SetApproverType sets field value
func (o *RequestTypeApprovalResourceOwnerWritable) SetApproverType(v string) {
	o.ApproverType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestTypeApprovalResourceOwnerWritable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalResourceOwnerWritable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeApprovalResourceOwnerWritable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestTypeApprovalResourceOwnerWritable) SetDescription(v string) {
	o.Description = &v
}

func (o RequestTypeApprovalResourceOwnerWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["approverType"] = o.ApproverType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeApprovalResourceOwnerWritable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeApprovalResourceOwnerWritable := _RequestTypeApprovalResourceOwnerWritable{}

	err = json.Unmarshal(bytes, &varRequestTypeApprovalResourceOwnerWritable)
	if err == nil {
		*o = RequestTypeApprovalResourceOwnerWritable(varRequestTypeApprovalResourceOwnerWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "approverType")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeApprovalResourceOwnerWritable struct {
	value *RequestTypeApprovalResourceOwnerWritable
	isSet bool
}

func (v NullableRequestTypeApprovalResourceOwnerWritable) Get() *RequestTypeApprovalResourceOwnerWritable {
	return v.value
}

func (v *NullableRequestTypeApprovalResourceOwnerWritable) Set(val *RequestTypeApprovalResourceOwnerWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalResourceOwnerWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalResourceOwnerWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalResourceOwnerWritable(val *RequestTypeApprovalResourceOwnerWritable) *NullableRequestTypeApprovalResourceOwnerWritable {
	return &NullableRequestTypeApprovalResourceOwnerWritable{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalResourceOwnerWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalResourceOwnerWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
