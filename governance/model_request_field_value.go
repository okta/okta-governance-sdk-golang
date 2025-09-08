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

// RequestFieldValue struct for RequestFieldValue
type RequestFieldValue struct {
	// The ID of a `requesterField`.  For `OKTA` request approval, this is typically the ID of an input field defined in an approval sequence.  For `EXTERNAL` request approval, this is the ID of an input field relevant to the external approval system.  The following are reserved identifiers: - `ACCESS_DURATION`: An identifier used to specify the duration of the access being requested. (**Note:**   This identifier is to be grandfathered. The new name will be `OKTA_ACCESS_DURATION`.) - `OKTA_REQUESTED_FOR`: An identifier used to specify who the resource is being requested for - `OKTA_ACCESS_DURATION`: An identifier used to specify the duration of the requested access. If a request    field is included in this ID, the field value must be a `DURATION` expression in ISO format. The maximum    length of this value is 5 characters. For example, `P1D` represents a duration of one day. - Any identifiers with `OKTA_` prefix are reserved for future use.
	Id string `json:"id"`
	// A human-readable description of `requesterField`. It's used for display purposes and is optional
	Label *string           `json:"label,omitempty"`
	Type  *RequestFieldType `json:"type,omitempty"`
	// The value of `requesterField`, which depends on the type of the field: - `TEXT`: A multi-line string - `SELECT`: An option predefined by the approval system - `ISO_DATE`: An ISO formatted date string. for example, `2022-05-05T14:15:22Z` - `DURATION`: An ISO format duration expression with a maximum length of 5 characters. For example,    `P1D` indicates a duration of one day. - `OKTA_USER_ID`: Okta User ID  If the field type is one of the previously listed, this property is required.
	Value *string `json:"value,omitempty"`
	// The values of `requesterField` with the type `MULTISELECT`.  If the field type is `MULTISELECT`, this property is required.
	Values               []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestFieldValue RequestFieldValue

// NewRequestFieldValue instantiates a new RequestFieldValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestFieldValue(id string) *RequestFieldValue {
	this := RequestFieldValue{}
	this.Id = id
	return &this
}

// NewRequestFieldValueWithDefaults instantiates a new RequestFieldValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFieldValueWithDefaults() *RequestFieldValue {
	this := RequestFieldValue{}
	return &this
}

// GetId returns the Id field value
func (o *RequestFieldValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestFieldValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestFieldValue) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RequestFieldValue) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldValue) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RequestFieldValue) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RequestFieldValue) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RequestFieldValue) GetType() RequestFieldType {
	if o == nil || o.Type == nil {
		var ret RequestFieldType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldValue) GetTypeOk() (*RequestFieldType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RequestFieldValue) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given RequestFieldType and assigns it to the Type field.
func (o *RequestFieldValue) SetType(v RequestFieldType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RequestFieldValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RequestFieldValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RequestFieldValue) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *RequestFieldValue) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFieldValue) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *RequestFieldValue) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *RequestFieldValue) SetValues(v []string) {
	o.Values = v
}

func (o RequestFieldValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestFieldValue) UnmarshalJSON(bytes []byte) (err error) {
	varRequestFieldValue := _RequestFieldValue{}

	err = json.Unmarshal(bytes, &varRequestFieldValue)
	if err == nil {
		*o = RequestFieldValue(varRequestFieldValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestFieldValue struct {
	value *RequestFieldValue
	isSet bool
}

func (v NullableRequestFieldValue) Get() *RequestFieldValue {
	return v.value
}

func (v *NullableRequestFieldValue) Set(val *RequestFieldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFieldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFieldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFieldValue(val *RequestFieldValue) *NullableRequestFieldValue {
	return &NullableRequestFieldValue{value: val, isSet: true}
}

func (v NullableRequestFieldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFieldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
