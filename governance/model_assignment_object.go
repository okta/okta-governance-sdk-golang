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

// checks if the AssignmentObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignmentObject{}

// AssignmentObject Assignment object
type AssignmentObject struct {
	// Unique identifier of assignmentObject
	Id *string `json:"id,omitempty"`
	// ORN of the assignmentObject
	Orn *string `json:"orn,omitempty"`
	// The name of the assignmentObject
	Name *string `json:"name,omitempty"`
	// The human-readable description of the assignmentObject
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignmentObject AssignmentObject

// NewAssignmentObject instantiates a new AssignmentObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignmentObject() *AssignmentObject {
	this := AssignmentObject{}
	return &this
}

// NewAssignmentObjectWithDefaults instantiates a new AssignmentObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignmentObjectWithDefaults() *AssignmentObject {
	this := AssignmentObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssignmentObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssignmentObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssignmentObject) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *AssignmentObject) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentObject) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *AssignmentObject) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *AssignmentObject) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssignmentObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssignmentObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssignmentObject) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssignmentObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssignmentObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssignmentObject) SetDescription(v string) {
	o.Description = &v
}

func (o AssignmentObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignmentObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignmentObject) UnmarshalJSON(data []byte) (err error) {
	varAssignmentObject := _AssignmentObject{}

	err = json.Unmarshal(data, &varAssignmentObject)

	if err != nil {
		return err
	}

	*o = AssignmentObject(varAssignmentObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignmentObject struct {
	value *AssignmentObject
	isSet bool
}

func (v NullableAssignmentObject) Get() *AssignmentObject {
	return v.value
}

func (v *NullableAssignmentObject) Set(val *AssignmentObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentObject(val *AssignmentObject) *NullableAssignmentObject {
	return &NullableAssignmentObject{value: val, isSet: true}
}

func (v NullableAssignmentObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
