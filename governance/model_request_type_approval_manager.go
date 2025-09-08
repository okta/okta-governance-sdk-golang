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

// RequestTypeApprovalManager Specifying the approver must be the requester's manager
type RequestTypeApprovalManager struct {
	ApproverType string  `json:"approverType"`
	Description  *string `json:"description,omitempty"`
	// A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.
	ApproverFields       []Field `json:"approverFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalManager RequestTypeApprovalManager

// NewRequestTypeApprovalManager instantiates a new RequestTypeApprovalManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalManager(approverType string) *RequestTypeApprovalManager {
	this := RequestTypeApprovalManager{}
	this.ApproverType = approverType
	return &this
}

// NewRequestTypeApprovalManagerWithDefaults instantiates a new RequestTypeApprovalManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalManagerWithDefaults() *RequestTypeApprovalManager {
	this := RequestTypeApprovalManager{}
	return &this
}

// GetApproverType returns the ApproverType field value
func (o *RequestTypeApprovalManager) GetApproverType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverType
}

// GetApproverTypeOk returns a tuple with the ApproverType field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalManager) GetApproverTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverType, true
}

// SetApproverType sets field value
func (o *RequestTypeApprovalManager) SetApproverType(v string) {
	o.ApproverType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestTypeApprovalManager) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalManager) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeApprovalManager) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestTypeApprovalManager) SetDescription(v string) {
	o.Description = &v
}

// GetApproverFields returns the ApproverFields field value if set, zero value otherwise.
func (o *RequestTypeApprovalManager) GetApproverFields() []Field {
	if o == nil || o.ApproverFields == nil {
		var ret []Field
		return ret
	}
	return o.ApproverFields
}

// GetApproverFieldsOk returns a tuple with the ApproverFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalManager) GetApproverFieldsOk() ([]Field, bool) {
	if o == nil || o.ApproverFields == nil {
		return nil, false
	}
	return o.ApproverFields, true
}

// HasApproverFields returns a boolean if a field has been set.
func (o *RequestTypeApprovalManager) HasApproverFields() bool {
	if o != nil && o.ApproverFields != nil {
		return true
	}

	return false
}

// SetApproverFields gets a reference to the given []Field and assigns it to the ApproverFields field.
func (o *RequestTypeApprovalManager) SetApproverFields(v []Field) {
	o.ApproverFields = v
}

func (o RequestTypeApprovalManager) MarshalJSON() ([]byte, error) {
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

func (o *RequestTypeApprovalManager) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeApprovalManager := _RequestTypeApprovalManager{}

	err = json.Unmarshal(bytes, &varRequestTypeApprovalManager)
	if err == nil {
		*o = RequestTypeApprovalManager(varRequestTypeApprovalManager)
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

type NullableRequestTypeApprovalManager struct {
	value *RequestTypeApprovalManager
	isSet bool
}

func (v NullableRequestTypeApprovalManager) Get() *RequestTypeApprovalManager {
	return v.value
}

func (v *NullableRequestTypeApprovalManager) Set(val *RequestTypeApprovalManager) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalManager) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalManager(val *RequestTypeApprovalManager) *NullableRequestTypeApprovalManager {
	return &NullableRequestTypeApprovalManager{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
