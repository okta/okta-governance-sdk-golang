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

// checks if the Note type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Note{}

// Note struct for Note
type Note struct {
	Id                   *string `json:"id,omitempty"`
	Note                 *string `json:"note,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Note Note

// NewNote instantiates a new Note object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNote() *Note {
	this := Note{}
	return &this
}

// NewNoteWithDefaults instantiates a new Note object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteWithDefaults() *Note {
	this := Note{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Note) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Note) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Note) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Note) SetId(v string) {
	o.Id = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *Note) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Note) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *Note) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *Note) SetNote(v string) {
	o.Note = &v
}

func (o Note) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Note) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Note) UnmarshalJSON(data []byte) (err error) {
	varNote := _Note{}

	err = json.Unmarshal(data, &varNote)

	if err != nil {
		return err
	}

	*o = Note(varNote)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "note")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNote struct {
	value *Note
	isSet bool
}

func (v NullableNote) Get() *Note {
	return v.value
}

func (v *NullableNote) Set(val *Note) {
	v.value = val
	v.isSet = true
}

func (v NullableNote) IsSet() bool {
	return v.isSet
}

func (v *NullableNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNote(val *Note) *NullableNote {
	return &NullableNote{value: val, isSet: true}
}

func (v NullableNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
