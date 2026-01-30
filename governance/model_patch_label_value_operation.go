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

// checks if the PatchLabelValueOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchLabelValueOperation{}

// PatchLabelValueOperation The properties for updating a label value
type PatchLabelValueOperation struct {
	Op LabelValuePatchOp `json:"op"`
	// The path of the property to update.  Use the following path format:  * `/values/-` for the `ADD` operation * `/values/{id}` for the `REMOVE` operation * `/values/{id}/...` for the `REPLACE` operation
	Path  string            `json:"path"`
	Value *LabelValueUpdate `json:"value,omitempty"`
	// The label property for the update operation
	RefType              string `json:"refType"`
	AdditionalProperties map[string]interface{}
}

type _PatchLabelValueOperation PatchLabelValueOperation

// NewPatchLabelValueOperation instantiates a new PatchLabelValueOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchLabelValueOperation(op LabelValuePatchOp, path string, refType string) *PatchLabelValueOperation {
	this := PatchLabelValueOperation{}
	this.Op = op
	this.Path = path
	this.RefType = refType
	return &this
}

// NewPatchLabelValueOperationWithDefaults instantiates a new PatchLabelValueOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchLabelValueOperationWithDefaults() *PatchLabelValueOperation {
	this := PatchLabelValueOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *PatchLabelValueOperation) GetOp() LabelValuePatchOp {
	if o == nil {
		var ret LabelValuePatchOp
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchLabelValueOperation) GetOpOk() (*LabelValuePatchOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchLabelValueOperation) SetOp(v LabelValuePatchOp) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchLabelValueOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchLabelValueOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchLabelValueOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchLabelValueOperation) GetValue() LabelValueUpdate {
	if o == nil || IsNil(o.Value) {
		var ret LabelValueUpdate
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLabelValueOperation) GetValueOk() (*LabelValueUpdate, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchLabelValueOperation) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given LabelValueUpdate and assigns it to the Value field.
func (o *PatchLabelValueOperation) SetValue(v LabelValueUpdate) {
	o.Value = &v
}

// GetRefType returns the RefType field value
func (o *PatchLabelValueOperation) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *PatchLabelValueOperation) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *PatchLabelValueOperation) SetRefType(v string) {
	o.RefType = v
}

func (o PatchLabelValueOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchLabelValueOperation) ToMap() (map[string]interface{}, error) {
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

func (o *PatchLabelValueOperation) UnmarshalJSON(data []byte) (err error) {
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

	varPatchLabelValueOperation := _PatchLabelValueOperation{}

	err = json.Unmarshal(data, &varPatchLabelValueOperation)

	if err != nil {
		return err
	}

	*o = PatchLabelValueOperation(varPatchLabelValueOperation)

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

type NullablePatchLabelValueOperation struct {
	value *PatchLabelValueOperation
	isSet bool
}

func (v NullablePatchLabelValueOperation) Get() *PatchLabelValueOperation {
	return v.value
}

func (v *NullablePatchLabelValueOperation) Set(val *PatchLabelValueOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchLabelValueOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchLabelValueOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchLabelValueOperation(val *PatchLabelValueOperation) *NullablePatchLabelValueOperation {
	return &NullablePatchLabelValueOperation{value: val, isSet: true}
}

func (v NullablePatchLabelValueOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchLabelValueOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
