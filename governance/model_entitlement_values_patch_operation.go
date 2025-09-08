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

// EntitlementValuesPatchOperation The properties expected in an update entitlement value.
type EntitlementValuesPatchOperation struct {
	Op EntitlementValuePatchOp `json:"op"`
	// The path of the property being updated. ex - `/values/{id}` for `REMOVE`, `REPLACE` and `/values/-` for `ADD` for entitlement value update.
	Path                 string                                `json:"path"`
	Value                *EntitlementValuesPatchOperationValue `json:"value,omitempty"`
	RefType              string                                `json:"refType"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValuesPatchOperation EntitlementValuesPatchOperation

// NewEntitlementValuesPatchOperation instantiates a new EntitlementValuesPatchOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValuesPatchOperation(op EntitlementValuePatchOp, path string, refType string) *EntitlementValuesPatchOperation {
	this := EntitlementValuesPatchOperation{}
	this.Op = op
	this.Path = path
	this.RefType = refType
	return &this
}

// NewEntitlementValuesPatchOperationWithDefaults instantiates a new EntitlementValuesPatchOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValuesPatchOperationWithDefaults() *EntitlementValuesPatchOperation {
	this := EntitlementValuesPatchOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *EntitlementValuesPatchOperation) GetOp() EntitlementValuePatchOp {
	if o == nil {
		var ret EntitlementValuePatchOp
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *EntitlementValuesPatchOperation) GetOpOk() (*EntitlementValuePatchOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *EntitlementValuesPatchOperation) SetOp(v EntitlementValuePatchOp) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *EntitlementValuesPatchOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *EntitlementValuesPatchOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *EntitlementValuesPatchOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementValuesPatchOperation) GetValue() EntitlementValuesPatchOperationValue {
	if o == nil || o.Value == nil {
		var ret EntitlementValuesPatchOperationValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesPatchOperation) GetValueOk() (*EntitlementValuesPatchOperationValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementValuesPatchOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given EntitlementValuesPatchOperationValue and assigns it to the Value field.
func (o *EntitlementValuesPatchOperation) SetValue(v EntitlementValuesPatchOperationValue) {
	o.Value = &v
}

// GetRefType returns the RefType field value
func (o *EntitlementValuesPatchOperation) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *EntitlementValuesPatchOperation) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *EntitlementValuesPatchOperation) SetRefType(v string) {
	o.RefType = v
}

func (o EntitlementValuesPatchOperation) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["refType"] = o.RefType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValuesPatchOperation) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValuesPatchOperation := _EntitlementValuesPatchOperation{}

	err = json.Unmarshal(bytes, &varEntitlementValuesPatchOperation)
	if err == nil {
		*o = EntitlementValuesPatchOperation(varEntitlementValuesPatchOperation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		delete(additionalProperties, "refType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValuesPatchOperation struct {
	value *EntitlementValuesPatchOperation
	isSet bool
}

func (v NullableEntitlementValuesPatchOperation) Get() *EntitlementValuesPatchOperation {
	return v.value
}

func (v *NullableEntitlementValuesPatchOperation) Set(val *EntitlementValuesPatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValuesPatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValuesPatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValuesPatchOperation(val *EntitlementValuesPatchOperation) *NullableEntitlementValuesPatchOperation {
	return &NullableEntitlementValuesPatchOperation{value: val, isSet: true}
}

func (v NullableEntitlementValuesPatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValuesPatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
