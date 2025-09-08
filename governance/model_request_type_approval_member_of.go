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

// RequestTypeApprovalMemberOf An approval when the approver must be a member of a particular Okta Group.
type RequestTypeApprovalMemberOf struct {
	ApproverType string  `json:"approverType"`
	Description  *string `json:"description,omitempty"`
	// Okta groups the user persona must be a member of
	ApproverMemberOf []string `json:"approverMemberOf"`
	// A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.
	ApproverFields       []Field `json:"approverFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalMemberOf RequestTypeApprovalMemberOf

// NewRequestTypeApprovalMemberOf instantiates a new RequestTypeApprovalMemberOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalMemberOf(approverType string, approverMemberOf []string) *RequestTypeApprovalMemberOf {
	this := RequestTypeApprovalMemberOf{}
	this.ApproverType = approverType
	this.ApproverMemberOf = approverMemberOf
	return &this
}

// NewRequestTypeApprovalMemberOfWithDefaults instantiates a new RequestTypeApprovalMemberOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalMemberOfWithDefaults() *RequestTypeApprovalMemberOf {
	this := RequestTypeApprovalMemberOf{}
	return &this
}

// GetApproverType returns the ApproverType field value
func (o *RequestTypeApprovalMemberOf) GetApproverType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverType
}

// GetApproverTypeOk returns a tuple with the ApproverType field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalMemberOf) GetApproverTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverType, true
}

// SetApproverType sets field value
func (o *RequestTypeApprovalMemberOf) SetApproverType(v string) {
	o.ApproverType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestTypeApprovalMemberOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalMemberOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeApprovalMemberOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestTypeApprovalMemberOf) SetDescription(v string) {
	o.Description = &v
}

// GetApproverMemberOf returns the ApproverMemberOf field value
func (o *RequestTypeApprovalMemberOf) GetApproverMemberOf() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ApproverMemberOf
}

// GetApproverMemberOfOk returns a tuple with the ApproverMemberOf field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalMemberOf) GetApproverMemberOfOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverMemberOf, true
}

// SetApproverMemberOf sets field value
func (o *RequestTypeApprovalMemberOf) SetApproverMemberOf(v []string) {
	o.ApproverMemberOf = v
}

// GetApproverFields returns the ApproverFields field value if set, zero value otherwise.
func (o *RequestTypeApprovalMemberOf) GetApproverFields() []Field {
	if o == nil || o.ApproverFields == nil {
		var ret []Field
		return ret
	}
	return o.ApproverFields
}

// GetApproverFieldsOk returns a tuple with the ApproverFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalMemberOf) GetApproverFieldsOk() ([]Field, bool) {
	if o == nil || o.ApproverFields == nil {
		return nil, false
	}
	return o.ApproverFields, true
}

// HasApproverFields returns a boolean if a field has been set.
func (o *RequestTypeApprovalMemberOf) HasApproverFields() bool {
	if o != nil && o.ApproverFields != nil {
		return true
	}

	return false
}

// SetApproverFields gets a reference to the given []Field and assigns it to the ApproverFields field.
func (o *RequestTypeApprovalMemberOf) SetApproverFields(v []Field) {
	o.ApproverFields = v
}

func (o RequestTypeApprovalMemberOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["approverType"] = o.ApproverType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["approverMemberOf"] = o.ApproverMemberOf
	}
	if o.ApproverFields != nil {
		toSerialize["approverFields"] = o.ApproverFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeApprovalMemberOf) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeApprovalMemberOf := _RequestTypeApprovalMemberOf{}

	err = json.Unmarshal(bytes, &varRequestTypeApprovalMemberOf)
	if err == nil {
		*o = RequestTypeApprovalMemberOf(varRequestTypeApprovalMemberOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "approverType")
		delete(additionalProperties, "description")
		delete(additionalProperties, "approverMemberOf")
		delete(additionalProperties, "approverFields")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeApprovalMemberOf struct {
	value *RequestTypeApprovalMemberOf
	isSet bool
}

func (v NullableRequestTypeApprovalMemberOf) Get() *RequestTypeApprovalMemberOf {
	return v.value
}

func (v *NullableRequestTypeApprovalMemberOf) Set(val *RequestTypeApprovalMemberOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalMemberOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalMemberOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalMemberOf(val *RequestTypeApprovalMemberOf) *NullableRequestTypeApprovalMemberOf {
	return &NullableRequestTypeApprovalMemberOf{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalMemberOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalMemberOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
