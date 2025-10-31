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

// checks if the RequestTypeApprovalResourceOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeApprovalResourceOwner{}

// RequestTypeApprovalResourceOwner Specifying the approver must be the resource owner
type RequestTypeApprovalResourceOwner struct {
	ApproverType         string  `json:"approverType"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalResourceOwner RequestTypeApprovalResourceOwner

// NewRequestTypeApprovalResourceOwner instantiates a new RequestTypeApprovalResourceOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalResourceOwner(approverType string) *RequestTypeApprovalResourceOwner {
	this := RequestTypeApprovalResourceOwner{}
	this.ApproverType = approverType
	return &this
}

// NewRequestTypeApprovalResourceOwnerWithDefaults instantiates a new RequestTypeApprovalResourceOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalResourceOwnerWithDefaults() *RequestTypeApprovalResourceOwner {
	this := RequestTypeApprovalResourceOwner{}
	return &this
}

// GetApproverType returns the ApproverType field value
func (o *RequestTypeApprovalResourceOwner) GetApproverType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverType
}

// GetApproverTypeOk returns a tuple with the ApproverType field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalResourceOwner) GetApproverTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverType, true
}

// SetApproverType sets field value
func (o *RequestTypeApprovalResourceOwner) SetApproverType(v string) {
	o.ApproverType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestTypeApprovalResourceOwner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalResourceOwner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeApprovalResourceOwner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestTypeApprovalResourceOwner) SetDescription(v string) {
	o.Description = &v
}

func (o RequestTypeApprovalResourceOwner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeApprovalResourceOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approverType"] = o.ApproverType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeApprovalResourceOwner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approverType",
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

	varRequestTypeApprovalResourceOwner := _RequestTypeApprovalResourceOwner{}

	err = json.Unmarshal(data, &varRequestTypeApprovalResourceOwner)

	if err != nil {
		return err
	}

	*o = RequestTypeApprovalResourceOwner(varRequestTypeApprovalResourceOwner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approverType")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeApprovalResourceOwner struct {
	value *RequestTypeApprovalResourceOwner
	isSet bool
}

func (v NullableRequestTypeApprovalResourceOwner) Get() *RequestTypeApprovalResourceOwner {
	return v.value
}

func (v *NullableRequestTypeApprovalResourceOwner) Set(val *RequestTypeApprovalResourceOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalResourceOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalResourceOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalResourceOwner(val *RequestTypeApprovalResourceOwner) *NullableRequestTypeApprovalResourceOwner {
	return &NullableRequestTypeApprovalResourceOwner{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalResourceOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalResourceOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
