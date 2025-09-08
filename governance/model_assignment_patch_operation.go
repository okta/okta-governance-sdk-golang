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

// AssignmentPatchOperation The properties expected in an update collection assignment.
type AssignmentPatchOperation struct {
	Op AssignmentPatchOp `json:"op"`
	// The path of the property that's updated. For example, `/expirationTime` and `/timeZone` for the assignment update.
	Path string `json:"path"`
	// The value of the property that's updated
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignmentPatchOperation AssignmentPatchOperation

// NewAssignmentPatchOperation instantiates a new AssignmentPatchOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignmentPatchOperation(op AssignmentPatchOp, path string) *AssignmentPatchOperation {
	this := AssignmentPatchOperation{}
	this.Op = op
	this.Path = path
	return &this
}

// NewAssignmentPatchOperationWithDefaults instantiates a new AssignmentPatchOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignmentPatchOperationWithDefaults() *AssignmentPatchOperation {
	this := AssignmentPatchOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *AssignmentPatchOperation) GetOp() AssignmentPatchOp {
	if o == nil {
		var ret AssignmentPatchOp
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *AssignmentPatchOperation) GetOpOk() (*AssignmentPatchOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *AssignmentPatchOperation) SetOp(v AssignmentPatchOp) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *AssignmentPatchOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AssignmentPatchOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *AssignmentPatchOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AssignmentPatchOperation) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentPatchOperation) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AssignmentPatchOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AssignmentPatchOperation) SetValue(v string) {
	o.Value = &v
}

func (o AssignmentPatchOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssignmentPatchOperation) UnmarshalJSON(bytes []byte) (err error) {
	varAssignmentPatchOperation := _AssignmentPatchOperation{}

	err = json.Unmarshal(bytes, &varAssignmentPatchOperation)
	if err == nil {
		*o = AssignmentPatchOperation(varAssignmentPatchOperation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAssignmentPatchOperation struct {
	value *AssignmentPatchOperation
	isSet bool
}

func (v NullableAssignmentPatchOperation) Get() *AssignmentPatchOperation {
	return v.value
}

func (v *NullableAssignmentPatchOperation) Set(val *AssignmentPatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentPatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentPatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentPatchOperation(val *AssignmentPatchOperation) *NullableAssignmentPatchOperation {
	return &NullableAssignmentPatchOperation{value: val, isSet: true}
}

func (v NullableAssignmentPatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentPatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
