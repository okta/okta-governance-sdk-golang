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

// EntitlementPatchOperation The properties expected in an update entitlement.
type EntitlementPatchOperation struct {
	Op EntitlementPatchOp `json:"op"`
	// The path of the property being updated. ex - `/name` and `/description` for the entitlement update.
	Path string `json:"path"`
	// The value of the property being updated.
	Value                *string `json:"value,omitempty"`
	RefType              string  `json:"refType"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementPatchOperation EntitlementPatchOperation

// NewEntitlementPatchOperation instantiates a new EntitlementPatchOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementPatchOperation(op EntitlementPatchOp, path string, refType string) *EntitlementPatchOperation {
	this := EntitlementPatchOperation{}
	this.Op = op
	this.Path = path
	this.RefType = refType
	return &this
}

// NewEntitlementPatchOperationWithDefaults instantiates a new EntitlementPatchOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementPatchOperationWithDefaults() *EntitlementPatchOperation {
	this := EntitlementPatchOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *EntitlementPatchOperation) GetOp() EntitlementPatchOp {
	if o == nil {
		var ret EntitlementPatchOp
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *EntitlementPatchOperation) GetOpOk() (*EntitlementPatchOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *EntitlementPatchOperation) SetOp(v EntitlementPatchOp) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *EntitlementPatchOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *EntitlementPatchOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *EntitlementPatchOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementPatchOperation) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementPatchOperation) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementPatchOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementPatchOperation) SetValue(v string) {
	o.Value = &v
}

// GetRefType returns the RefType field value
func (o *EntitlementPatchOperation) GetRefType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *EntitlementPatchOperation) GetRefTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *EntitlementPatchOperation) SetRefType(v string) {
	o.RefType = v
}

func (o EntitlementPatchOperation) MarshalJSON() ([]byte, error) {
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

func (o *EntitlementPatchOperation) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementPatchOperation := _EntitlementPatchOperation{}

	err = json.Unmarshal(bytes, &varEntitlementPatchOperation)
	if err == nil {
		*o = EntitlementPatchOperation(varEntitlementPatchOperation)
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

type NullableEntitlementPatchOperation struct {
	value *EntitlementPatchOperation
	isSet bool
}

func (v NullableEntitlementPatchOperation) Get() *EntitlementPatchOperation {
	return v.value
}

func (v *NullableEntitlementPatchOperation) Set(val *EntitlementPatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementPatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementPatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementPatchOperation(val *EntitlementPatchOperation) *NullableEntitlementPatchOperation {
	return &NullableEntitlementPatchOperation{value: val, isSet: true}
}

func (v NullableEntitlementPatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementPatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
