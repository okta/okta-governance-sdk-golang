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

// checks if the MyResourceConnectionCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MyResourceConnectionCommon{}

// MyResourceConnectionCommon struct for MyResourceConnectionCommon
type MyResourceConnectionCommon struct {
	// Unique identifier for the resource connection
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the resource connection
	Orn *string `json:"orn,omitempty"`
	// The status of the connection
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyResourceConnectionCommon MyResourceConnectionCommon

// NewMyResourceConnectionCommon instantiates a new MyResourceConnectionCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyResourceConnectionCommon() *MyResourceConnectionCommon {
	this := MyResourceConnectionCommon{}
	return &this
}

// NewMyResourceConnectionCommonWithDefaults instantiates a new MyResourceConnectionCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyResourceConnectionCommonWithDefaults() *MyResourceConnectionCommon {
	this := MyResourceConnectionCommon{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MyResourceConnectionCommon) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyResourceConnectionCommon) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MyResourceConnectionCommon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MyResourceConnectionCommon) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *MyResourceConnectionCommon) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyResourceConnectionCommon) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *MyResourceConnectionCommon) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *MyResourceConnectionCommon) SetOrn(v string) {
	o.Orn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MyResourceConnectionCommon) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyResourceConnectionCommon) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MyResourceConnectionCommon) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MyResourceConnectionCommon) SetStatus(v string) {
	o.Status = &v
}

func (o MyResourceConnectionCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyResourceConnectionCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MyResourceConnectionCommon) UnmarshalJSON(data []byte) (err error) {
	varMyResourceConnectionCommon := _MyResourceConnectionCommon{}

	err = json.Unmarshal(data, &varMyResourceConnectionCommon)

	if err != nil {
		return err
	}

	*o = MyResourceConnectionCommon(varMyResourceConnectionCommon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyResourceConnectionCommon struct {
	value *MyResourceConnectionCommon
	isSet bool
}

func (v NullableMyResourceConnectionCommon) Get() *MyResourceConnectionCommon {
	return v.value
}

func (v *NullableMyResourceConnectionCommon) Set(val *MyResourceConnectionCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableMyResourceConnectionCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableMyResourceConnectionCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyResourceConnectionCommon(val *MyResourceConnectionCommon) *NullableMyResourceConnectionCommon {
	return &NullableMyResourceConnectionCommon{value: val, isSet: true}
}

func (v NullableMyResourceConnectionCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyResourceConnectionCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
