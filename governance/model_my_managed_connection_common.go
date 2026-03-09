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

// checks if the MyManagedConnectionCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MyManagedConnectionCommon{}

// MyManagedConnectionCommon struct for MyManagedConnectionCommon
type MyManagedConnectionCommon struct {
	// Unique identifier for the managed connection
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection
	Orn *string `json:"orn,omitempty"`
	// The status of the connection
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyManagedConnectionCommon MyManagedConnectionCommon

// NewMyManagedConnectionCommon instantiates a new MyManagedConnectionCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyManagedConnectionCommon() *MyManagedConnectionCommon {
	this := MyManagedConnectionCommon{}
	return &this
}

// NewMyManagedConnectionCommonWithDefaults instantiates a new MyManagedConnectionCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyManagedConnectionCommonWithDefaults() *MyManagedConnectionCommon {
	this := MyManagedConnectionCommon{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MyManagedConnectionCommon) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyManagedConnectionCommon) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MyManagedConnectionCommon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MyManagedConnectionCommon) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *MyManagedConnectionCommon) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyManagedConnectionCommon) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *MyManagedConnectionCommon) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *MyManagedConnectionCommon) SetOrn(v string) {
	o.Orn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MyManagedConnectionCommon) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyManagedConnectionCommon) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MyManagedConnectionCommon) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MyManagedConnectionCommon) SetStatus(v string) {
	o.Status = &v
}

func (o MyManagedConnectionCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyManagedConnectionCommon) ToMap() (map[string]interface{}, error) {
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

func (o *MyManagedConnectionCommon) UnmarshalJSON(data []byte) (err error) {
	varMyManagedConnectionCommon := _MyManagedConnectionCommon{}

	err = json.Unmarshal(data, &varMyManagedConnectionCommon)

	if err != nil {
		return err
	}

	*o = MyManagedConnectionCommon(varMyManagedConnectionCommon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyManagedConnectionCommon struct {
	value *MyManagedConnectionCommon
	isSet bool
}

func (v NullableMyManagedConnectionCommon) Get() *MyManagedConnectionCommon {
	return v.value
}

func (v *NullableMyManagedConnectionCommon) Set(val *MyManagedConnectionCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableMyManagedConnectionCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableMyManagedConnectionCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyManagedConnectionCommon(val *MyManagedConnectionCommon) *NullableMyManagedConnectionCommon {
	return &NullableMyManagedConnectionCommon{value: val, isSet: true}
}

func (v NullableMyManagedConnectionCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyManagedConnectionCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
