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

// RequestTypeApprovalManagerWritable Specifying the approver must be the requester's manager
type RequestTypeApprovalManagerWritable struct {
	ApproverType string  `json:"approverType"`
	Description  *string `json:"description,omitempty"`
	// A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.
	ApproverFields       []FieldWritable `json:"approverFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalManagerWritable RequestTypeApprovalManagerWritable

// NewRequestTypeApprovalManagerWritable instantiates a new RequestTypeApprovalManagerWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalManagerWritable(approverType string) *RequestTypeApprovalManagerWritable {
	this := RequestTypeApprovalManagerWritable{}
	this.ApproverType = approverType
	return &this
}

// NewRequestTypeApprovalManagerWritableWithDefaults instantiates a new RequestTypeApprovalManagerWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalManagerWritableWithDefaults() *RequestTypeApprovalManagerWritable {
	this := RequestTypeApprovalManagerWritable{}
	return &this
}

// GetApproverType returns the ApproverType field value
func (o *RequestTypeApprovalManagerWritable) GetApproverType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverType
}

// GetApproverTypeOk returns a tuple with the ApproverType field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalManagerWritable) GetApproverTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverType, true
}

// SetApproverType sets field value
func (o *RequestTypeApprovalManagerWritable) SetApproverType(v string) {
	o.ApproverType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestTypeApprovalManagerWritable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalManagerWritable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeApprovalManagerWritable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestTypeApprovalManagerWritable) SetDescription(v string) {
	o.Description = &v
}

// GetApproverFields returns the ApproverFields field value if set, zero value otherwise.
func (o *RequestTypeApprovalManagerWritable) GetApproverFields() []FieldWritable {
	if o == nil || o.ApproverFields == nil {
		var ret []FieldWritable
		return ret
	}
	return o.ApproverFields
}

// GetApproverFieldsOk returns a tuple with the ApproverFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalManagerWritable) GetApproverFieldsOk() ([]FieldWritable, bool) {
	if o == nil || o.ApproverFields == nil {
		return nil, false
	}
	return o.ApproverFields, true
}

// HasApproverFields returns a boolean if a field has been set.
func (o *RequestTypeApprovalManagerWritable) HasApproverFields() bool {
	if o != nil && o.ApproverFields != nil {
		return true
	}

	return false
}

// SetApproverFields gets a reference to the given []FieldWritable and assigns it to the ApproverFields field.
func (o *RequestTypeApprovalManagerWritable) SetApproverFields(v []FieldWritable) {
	o.ApproverFields = v
}

func (o RequestTypeApprovalManagerWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["approverType"] = o.ApproverType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ApproverFields != nil {
		toSerialize["approverFields"] = o.ApproverFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeApprovalManagerWritable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeApprovalManagerWritable := _RequestTypeApprovalManagerWritable{}

	err = json.Unmarshal(bytes, &varRequestTypeApprovalManagerWritable)
	if err == nil {
		*o = RequestTypeApprovalManagerWritable(varRequestTypeApprovalManagerWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "approverType")
		delete(additionalProperties, "description")
		delete(additionalProperties, "approverFields")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeApprovalManagerWritable struct {
	value *RequestTypeApprovalManagerWritable
	isSet bool
}

func (v NullableRequestTypeApprovalManagerWritable) Get() *RequestTypeApprovalManagerWritable {
	return v.value
}

func (v *NullableRequestTypeApprovalManagerWritable) Set(val *RequestTypeApprovalManagerWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalManagerWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalManagerWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalManagerWritable(val *RequestTypeApprovalManagerWritable) *NullableRequestTypeApprovalManagerWritable {
	return &NullableRequestTypeApprovalManagerWritable{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalManagerWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalManagerWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
