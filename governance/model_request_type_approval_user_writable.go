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

// checks if the RequestTypeApprovalUserWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeApprovalUserWritable{}

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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalUserWritable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeApprovalUserWritable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.ApproverFields) {
		var ret []FieldWritable
		return ret
	}
	return o.ApproverFields
}

// GetApproverFieldsOk returns a tuple with the ApproverFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalUserWritable) GetApproverFieldsOk() ([]FieldWritable, bool) {
	if o == nil || IsNil(o.ApproverFields) {
		return nil, false
	}
	return o.ApproverFields, true
}

// HasApproverFields returns a boolean if a field has been set.
func (o *RequestTypeApprovalUserWritable) HasApproverFields() bool {
	if o != nil && !IsNil(o.ApproverFields) {
		return true
	}

	return false
}

// SetApproverFields gets a reference to the given []FieldWritable and assigns it to the ApproverFields field.
func (o *RequestTypeApprovalUserWritable) SetApproverFields(v []FieldWritable) {
	o.ApproverFields = v
}

func (o RequestTypeApprovalUserWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeApprovalUserWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approverType"] = o.ApproverType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["approverUserId"] = o.ApproverUserId
	if !IsNil(o.ApproverFields) {
		toSerialize["approverFields"] = o.ApproverFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeApprovalUserWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approverType",
		"approverUserId",
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

	varRequestTypeApprovalUserWritable := _RequestTypeApprovalUserWritable{}

	err = json.Unmarshal(data, &varRequestTypeApprovalUserWritable)

	if err != nil {
		return err
	}

	*o = RequestTypeApprovalUserWritable(varRequestTypeApprovalUserWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approverType")
		delete(additionalProperties, "description")
		delete(additionalProperties, "approverUserId")
		delete(additionalProperties, "approverFields")
		o.AdditionalProperties = additionalProperties
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
