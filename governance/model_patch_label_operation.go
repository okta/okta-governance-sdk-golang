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

// checks if the PatchLabelOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchLabelOperation{}

// PatchLabelOperation The properties for updating a label
type PatchLabelOperation struct {
	Op LabelPatchOp `json:"op"`
	// The path of the property being updated. ex - `/name` label category update.
	Path string `json:"path"`
	// The value of the property being updated.
	Value                *string `json:"value,omitempty"`
	RefType              string  `json:"refType"`
	AdditionalProperties map[string]interface{}
}

type _PatchLabelOperation PatchLabelOperation

// NewPatchLabelOperation instantiates a new PatchLabelOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchLabelOperation(op LabelPatchOp, path string, refType string) *PatchLabelOperation {
	this := PatchLabelOperation{}
	this.Op = op
	this.Path = path
	this.RefType = refType
	return &this
}

// NewPatchLabelOperationWithDefaults instantiates a new PatchLabelOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchLabelOperationWithDefaults() *PatchLabelOperation {
	this := PatchLabelOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *PatchLabelOperation) GetOp() LabelPatchOp {
	if o == nil {
		var ret LabelPatchOp
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchLabelOperation) GetOpOk() (*LabelPatchOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchLabelOperation) SetOp(v LabelPatchOp) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchLabelOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchLabelOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchLabelOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchLabelOperation) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLabelOperation) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchLabelOperation) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PatchLabelOperation) SetValue(v string) {
	o.Value = &v
}

// GetRefType returns the RefType field value
func (o *PatchLabelOperation) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *PatchLabelOperation) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *PatchLabelOperation) SetRefType(v string) {
	o.RefType = v
}

func (o PatchLabelOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchLabelOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	toSerialize["refType"] = o.RefType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchLabelOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
		"refType",
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

	varPatchLabelOperation := _PatchLabelOperation{}

	err = json.Unmarshal(data, &varPatchLabelOperation)

	if err != nil {
		return err
	}

	*o = PatchLabelOperation(varPatchLabelOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		delete(additionalProperties, "refType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchLabelOperation struct {
	value *PatchLabelOperation
	isSet bool
}

func (v NullablePatchLabelOperation) Get() *PatchLabelOperation {
	return v.value
}

func (v *NullablePatchLabelOperation) Set(val *PatchLabelOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchLabelOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchLabelOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchLabelOperation(val *PatchLabelOperation) *NullablePatchLabelOperation {
	return &NullablePatchLabelOperation{value: val, isSet: true}
}

func (v NullablePatchLabelOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchLabelOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
