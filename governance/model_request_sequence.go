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

// checks if the RequestSequence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestSequence{}

// RequestSequence A sequence that decides what actions must be taken on a request
type RequestSequence struct {
	// Unique idenfitifer of the request sequence
	Id string `json:"id"`
	// Name of the request sequence
	Name string `json:"name"`
	// Description of the request sequence
	Description             string                    `json:"description"`
	CompatibleResourceTypes []CompatibleResourceTypes `json:"compatibleResourceTypes,omitempty"`
	// Link to edit the request sequence
	Link                 string `json:"link"`
	AdditionalProperties map[string]interface{}
}

type _RequestSequence RequestSequence

// NewRequestSequence instantiates a new RequestSequence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSequence(id string, name string, description string, link string) *RequestSequence {
	this := RequestSequence{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Link = link
	return &this
}

// NewRequestSequenceWithDefaults instantiates a new RequestSequence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSequenceWithDefaults() *RequestSequence {
	this := RequestSequence{}
	return &this
}

// GetId returns the Id field value
func (o *RequestSequence) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestSequence) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestSequence) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RequestSequence) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestSequence) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestSequence) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RequestSequence) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RequestSequence) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RequestSequence) SetDescription(v string) {
	o.Description = v
}

// GetCompatibleResourceTypes returns the CompatibleResourceTypes field value if set, zero value otherwise.
func (o *RequestSequence) GetCompatibleResourceTypes() []CompatibleResourceTypes {
	if o == nil || IsNil(o.CompatibleResourceTypes) {
		var ret []CompatibleResourceTypes
		return ret
	}
	return o.CompatibleResourceTypes
}

// GetCompatibleResourceTypesOk returns a tuple with the CompatibleResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSequence) GetCompatibleResourceTypesOk() ([]CompatibleResourceTypes, bool) {
	if o == nil || IsNil(o.CompatibleResourceTypes) {
		return nil, false
	}
	return o.CompatibleResourceTypes, true
}

// HasCompatibleResourceTypes returns a boolean if a field has been set.
func (o *RequestSequence) HasCompatibleResourceTypes() bool {
	if o != nil && !IsNil(o.CompatibleResourceTypes) {
		return true
	}

	return false
}

// SetCompatibleResourceTypes gets a reference to the given []CompatibleResourceTypes and assigns it to the CompatibleResourceTypes field.
func (o *RequestSequence) SetCompatibleResourceTypes(v []CompatibleResourceTypes) {
	o.CompatibleResourceTypes = v
}

// GetLink returns the Link field value
func (o *RequestSequence) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *RequestSequence) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *RequestSequence) SetLink(v string) {
	o.Link = v
}

func (o RequestSequence) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSequence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.CompatibleResourceTypes) {
		toSerialize["compatibleResourceTypes"] = o.CompatibleResourceTypes
	}
	toSerialize["link"] = o.Link

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestSequence) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"link",
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

	varRequestSequence := _RequestSequence{}

	err = json.Unmarshal(data, &varRequestSequence)

	if err != nil {
		return err
	}

	*o = RequestSequence(varRequestSequence)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "compatibleResourceTypes")
		delete(additionalProperties, "link")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestSequence struct {
	value *RequestSequence
	isSet bool
}

func (v NullableRequestSequence) Get() *RequestSequence {
	return v.value
}

func (v *NullableRequestSequence) Set(val *RequestSequence) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSequence) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSequence(val *RequestSequence) *NullableRequestSequence {
	return &NullableRequestSequence{value: val, isSet: true}
}

func (v NullableRequestSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
