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

// checks if the RequestFieldChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestFieldChoice{}

// RequestFieldChoice A selectable choice for a request field
type RequestFieldChoice struct {
	// Unique identifier for the choice. Stable across calls and used as the answer value.
	Id string `json:"id"`
	// Display label for the choice
	Label string `json:"label"`
	// Optional description providing additional context for the choice
	Description *string `json:"description,omitempty"`
	// Optional category label for the choice.
	Category             *string `json:"category,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestFieldChoice RequestFieldChoice

// NewRequestFieldChoice instantiates a new RequestFieldChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestFieldChoice(id string, label string) *RequestFieldChoice {
	this := RequestFieldChoice{}
	this.Id = id
	this.Label = label
	return &this
}

// NewRequestFieldChoiceWithDefaults instantiates a new RequestFieldChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFieldChoiceWithDefaults() *RequestFieldChoice {
	this := RequestFieldChoice{}
	return &this
}

// GetId returns the Id field value
func (o *RequestFieldChoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestFieldChoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestFieldChoice) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *RequestFieldChoice) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *RequestFieldChoice) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *RequestFieldChoice) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestFieldChoice) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldChoice) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestFieldChoice) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestFieldChoice) SetDescription(v string) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *RequestFieldChoice) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldChoice) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *RequestFieldChoice) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *RequestFieldChoice) SetCategory(v string) {
	o.Category = &v
}

func (o RequestFieldChoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestFieldChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestFieldChoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"label",
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

	varRequestFieldChoice := _RequestFieldChoice{}

	err = json.Unmarshal(data, &varRequestFieldChoice)

	if err != nil {
		return err
	}

	*o = RequestFieldChoice(varRequestFieldChoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "category")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestFieldChoice struct {
	value *RequestFieldChoice
	isSet bool
}

func (v NullableRequestFieldChoice) Get() *RequestFieldChoice {
	return v.value
}

func (v *NullableRequestFieldChoice) Set(val *RequestFieldChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFieldChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFieldChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFieldChoice(val *RequestFieldChoice) *NullableRequestFieldChoice {
	return &NullableRequestFieldChoice{value: val, isSet: true}
}

func (v NullableRequestFieldChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFieldChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
