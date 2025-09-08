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

// RequestTypeApprovalUserWritable An approval when the approver must be a particular user.
type RequestTypeApprovalUserWritable struct {
	ApproverType string  `json:"approverType"`
	Description  *string `json:"description,omitempty"`
	// The Okta user `id`
	ApproverUserId string `json:"approverUserId" validate:"regexp=00u[0-9a-zA-Z]+"`
	// A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.
	ApproverFields       []FieldWritable `json:"approverFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalUserWritable RequestTypeApprovalUserWritable

// NewRequestTypeApprovalUserWritable instantiates a new RequestTypeApprovalUserWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalUserWritable(approverType string, approverUserId string) *RequestTypeApprovalUserWritable {
	this := RequestTypeApprovalUserWritable{}
	this.ApproverType = approverType
	this.ApproverUserId = approverUserId
	return &this
}

// NewRequestTypeApprovalUserWritableWithDefaults instantiates a new RequestTypeApprovalUserWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalUserWritableWithDefaults() *RequestTypeApprovalUserWritable {
	this := RequestTypeApprovalUserWritable{}
	return &this
}

// GetApproverType returns the ApproverType field value
func (o *RequestTypeApprovalUserWritable) GetApproverType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverType
}

// GetApproverTypeOk returns a tuple with the ApproverType field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalUserWritable) GetApproverTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverType, true
}

// SetApproverType sets field value
func (o *RequestTypeApprovalUserWritable) SetApproverType(v string) {
	o.ApproverType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestTypeApprovalUserWritable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalUserWritable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeApprovalUserWritable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestTypeApprovalUserWritable) SetDescription(v string) {
	o.Description = &v
}

// GetApproverUserId returns the ApproverUserId field value
func (o *RequestTypeApprovalUserWritable) GetApproverUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverUserId
}

// GetApproverUserIdOk returns a tuple with the ApproverUserId field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalUserWritable) GetApproverUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverUserId, true
}

// SetApproverUserId sets field value
func (o *RequestTypeApprovalUserWritable) SetApproverUserId(v string) {
	o.ApproverUserId = v
}

// GetApproverFields returns the ApproverFields field value if set, zero value otherwise.
func (o *RequestTypeApprovalUserWritable) GetApproverFields() []FieldWritable {
	if o == nil || o.ApproverFields == nil {
		var ret []FieldWritable
		return ret
	}
	return o.ApproverFields
}

// GetApproverFieldsOk returns a tuple with the ApproverFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalUserWritable) GetApproverFieldsOk() ([]FieldWritable, bool) {
	if o == nil || o.ApproverFields == nil {
		return nil, false
	}
	return o.ApproverFields, true
}

// HasApproverFields returns a boolean if a field has been set.
func (o *RequestTypeApprovalUserWritable) HasApproverFields() bool {
	if o != nil && o.ApproverFields != nil {
		return true
	}

	return false
}

// SetApproverFields gets a reference to the given []FieldWritable and assigns it to the ApproverFields field.
func (o *RequestTypeApprovalUserWritable) SetApproverFields(v []FieldWritable) {
	o.ApproverFields = v
}

func (o RequestTypeApprovalUserWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["approverType"] = o.ApproverType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["approverUserId"] = o.ApproverUserId
	}
	if o.ApproverFields != nil {
		toSerialize["approverFields"] = o.ApproverFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeApprovalUserWritable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeApprovalUserWritable := _RequestTypeApprovalUserWritable{}

	err = json.Unmarshal(bytes, &varRequestTypeApprovalUserWritable)
	if err == nil {
		*o = RequestTypeApprovalUserWritable(varRequestTypeApprovalUserWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "approverType")
		delete(additionalProperties, "description")
		delete(additionalProperties, "approverUserId")
		delete(additionalProperties, "approverFields")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeApprovalUserWritable struct {
	value *RequestTypeApprovalUserWritable
	isSet bool
}

func (v NullableRequestTypeApprovalUserWritable) Get() *RequestTypeApprovalUserWritable {
	return v.value
}

func (v *NullableRequestTypeApprovalUserWritable) Set(val *RequestTypeApprovalUserWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalUserWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalUserWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalUserWritable(val *RequestTypeApprovalUserWritable) *NullableRequestTypeApprovalUserWritable {
	return &NullableRequestTypeApprovalUserWritable{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalUserWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalUserWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
