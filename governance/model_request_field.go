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

// checks if the RequestField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestField{}

// RequestField Fields to be filled out for a request
type RequestField struct {
	// A `read-only` field ID.  Useful for specifying requesterFieldValues when adding a request.
	Id   string           `json:"id"`
	Type RequestFieldType `json:"type"`
	// Valid choices when type is SELECT or MULTISELECT.
	Choices []string `json:"choices,omitempty"`
	// Whether a value to this field is required to advance the request
	Required bool `json:"required"`
	// label of the requester field
	Label *string `json:"label,omitempty"`
	// An admin configured value for this field. Only applies to DURATION fields.
	Value *string `json:"value,omitempty"`
	// The maximum value allowed for this field. Only applies to DURATION fields.
	MaximumValue *string `json:"maximumValue,omitempty"`
	// Indicates this field is immutable.
	ReadOnly             *bool `json:"readOnly,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestField RequestField

// NewRequestField instantiates a new RequestField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestField(id string, type_ RequestFieldType, required bool) *RequestField {
	this := RequestField{}
	this.Id = id
	this.Type = type_
	this.Required = required
	return &this
}

// NewRequestFieldWithDefaults instantiates a new RequestField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFieldWithDefaults() *RequestField {
	this := RequestField{}
	var required bool = true
	this.Required = required
	return &this
}

// GetId returns the Id field value
func (o *RequestField) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestField) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestField) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RequestField) GetType() RequestFieldType {
	if o == nil {
		var ret RequestFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestField) GetTypeOk() (*RequestFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestField) SetType(v RequestFieldType) {
	o.Type = v
}

// GetChoices returns the Choices field value if set, zero value otherwise.
func (o *RequestField) GetChoices() []string {
	if o == nil || IsNil(o.Choices) {
		var ret []string
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestField) GetChoicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *RequestField) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []string and assigns it to the Choices field.
func (o *RequestField) SetChoices(v []string) {
	o.Choices = v
}

// GetRequired returns the Required field value
func (o *RequestField) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *RequestField) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *RequestField) SetRequired(v bool) {
	o.Required = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RequestField) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestField) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RequestField) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RequestField) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RequestField) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestField) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RequestField) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RequestField) SetValue(v string) {
	o.Value = &v
}

// GetMaximumValue returns the MaximumValue field value if set, zero value otherwise.
func (o *RequestField) GetMaximumValue() string {
	if o == nil || IsNil(o.MaximumValue) {
		var ret string
		return ret
	}
	return *o.MaximumValue
}

// GetMaximumValueOk returns a tuple with the MaximumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestField) GetMaximumValueOk() (*string, bool) {
	if o == nil || IsNil(o.MaximumValue) {
		return nil, false
	}
	return o.MaximumValue, true
}

// HasMaximumValue returns a boolean if a field has been set.
func (o *RequestField) HasMaximumValue() bool {
	if o != nil && !IsNil(o.MaximumValue) {
		return true
	}

	return false
}

// SetMaximumValue gets a reference to the given string and assigns it to the MaximumValue field.
func (o *RequestField) SetMaximumValue(v string) {
	o.MaximumValue = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *RequestField) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestField) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *RequestField) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *RequestField) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o RequestField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Choices) {
		toSerialize["choices"] = o.Choices
	}
	toSerialize["required"] = o.Required
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.MaximumValue) {
		toSerialize["maximumValue"] = o.MaximumValue
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"required",
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

	varRequestField := _RequestField{}

	err = json.Unmarshal(data, &varRequestField)

	if err != nil {
		return err
	}

	*o = RequestField(varRequestField)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "choices")
		delete(additionalProperties, "required")
		delete(additionalProperties, "label")
		delete(additionalProperties, "value")
		delete(additionalProperties, "maximumValue")
		delete(additionalProperties, "readOnly")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestField struct {
	value *RequestField
	isSet bool
}

func (v NullableRequestField) Get() *RequestField {
	return v.value
}

func (v *NullableRequestField) Set(val *RequestField) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestField) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestField(val *RequestField) *NullableRequestField {
	return &NullableRequestField{value: val, isSet: true}
}

func (v NullableRequestField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
