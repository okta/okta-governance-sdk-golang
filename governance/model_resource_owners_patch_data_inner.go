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
)

// checks if the ResourceOwnersPatchDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOwnersPatchDataInner{}

// ResourceOwnersPatchDataInner struct for ResourceOwnersPatchDataInner
type ResourceOwnersPatchDataInner struct {
	// The operation to be performed for update.
	Op *string `json:"op,omitempty"`
	// The path of the property being updated. Only `/principalOrn` can be updated.
	Path *string `json:"path,omitempty"`
	// The Okta user or group `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an user id, or a group id. See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnersPatchDataInner ResourceOwnersPatchDataInner

// NewResourceOwnersPatchDataInner instantiates a new ResourceOwnersPatchDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnersPatchDataInner() *ResourceOwnersPatchDataInner {
	this := ResourceOwnersPatchDataInner{}
	return &this
}

// NewResourceOwnersPatchDataInnerWithDefaults instantiates a new ResourceOwnersPatchDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnersPatchDataInnerWithDefaults() *ResourceOwnersPatchDataInner {
	this := ResourceOwnersPatchDataInner{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *ResourceOwnersPatchDataInner) GetOp() string {
	if o == nil || IsNil(o.Op) {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersPatchDataInner) GetOpOk() (*string, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *ResourceOwnersPatchDataInner) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *ResourceOwnersPatchDataInner) SetOp(v string) {
	o.Op = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ResourceOwnersPatchDataInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersPatchDataInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ResourceOwnersPatchDataInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ResourceOwnersPatchDataInner) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResourceOwnersPatchDataInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersPatchDataInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResourceOwnersPatchDataInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResourceOwnersPatchDataInner) SetValue(v string) {
	o.Value = &v
}

func (o ResourceOwnersPatchDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOwnersPatchDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceOwnersPatchDataInner) UnmarshalJSON(data []byte) (err error) {
	varResourceOwnersPatchDataInner := _ResourceOwnersPatchDataInner{}

	err = json.Unmarshal(data, &varResourceOwnersPatchDataInner)

	if err != nil {
		return err
	}

	*o = ResourceOwnersPatchDataInner(varResourceOwnersPatchDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceOwnersPatchDataInner struct {
	value *ResourceOwnersPatchDataInner
	isSet bool
}

func (v NullableResourceOwnersPatchDataInner) Get() *ResourceOwnersPatchDataInner {
	return v.value
}

func (v *NullableResourceOwnersPatchDataInner) Set(val *ResourceOwnersPatchDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnersPatchDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnersPatchDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnersPatchDataInner(val *ResourceOwnersPatchDataInner) *NullableResourceOwnersPatchDataInner {
	return &NullableResourceOwnersPatchDataInner{value: val, isSet: true}
}

func (v NullableResourceOwnersPatchDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnersPatchDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
