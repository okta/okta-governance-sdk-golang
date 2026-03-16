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

// checks if the AiMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMessage{}

// AiMessage An AI-generated summary is presented as a long text string, and sometimes multiple paragraphs If an AI summary fails to generate, an array of errors is returned with possible steps to correct the failure.
type AiMessage struct {
	// Generated summary message
	Message string `json:"message"`
	// Generated summary message as a delta (for streaming)
	DeltaMessage *string `json:"deltaMessage,omitempty"`
	// Whenever summary generation has resulted in an error or blocked, the array of errors will detail the reasons.
	Errors               []AiMessageErrorsInner `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiMessage AiMessage

// NewAiMessage instantiates a new AiMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMessage(message string) *AiMessage {
	this := AiMessage{}
	this.Message = message
	return &this
}

// NewAiMessageWithDefaults instantiates a new AiMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMessageWithDefaults() *AiMessage {
	this := AiMessage{}
	return &this
}

// GetMessage returns the Message field value
func (o *AiMessage) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiMessage) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AiMessage) SetMessage(v string) {
	o.Message = v
}

// GetDeltaMessage returns the DeltaMessage field value if set, zero value otherwise.
func (o *AiMessage) GetDeltaMessage() string {
	if o == nil || IsNil(o.DeltaMessage) {
		var ret string
		return ret
	}
	return *o.DeltaMessage
}

// GetDeltaMessageOk returns a tuple with the DeltaMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetDeltaMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DeltaMessage) {
		return nil, false
	}
	return o.DeltaMessage, true
}

// HasDeltaMessage returns a boolean if a field has been set.
func (o *AiMessage) HasDeltaMessage() bool {
	if o != nil && !IsNil(o.DeltaMessage) {
		return true
	}

	return false
}

// SetDeltaMessage gets a reference to the given string and assigns it to the DeltaMessage field.
func (o *AiMessage) SetDeltaMessage(v string) {
	o.DeltaMessage = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AiMessage) GetErrors() []AiMessageErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []AiMessageErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMessage) GetErrorsOk() ([]AiMessageErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AiMessage) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AiMessageErrorsInner and assigns it to the Errors field.
func (o *AiMessage) SetErrors(v []AiMessageErrorsInner) {
	o.Errors = v
}

func (o AiMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.DeltaMessage) {
		toSerialize["deltaMessage"] = o.DeltaMessage
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varAiMessage := _AiMessage{}

	err = json.Unmarshal(data, &varAiMessage)

	if err != nil {
		return err
	}

	*o = AiMessage(varAiMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "deltaMessage")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiMessage struct {
	value *AiMessage
	isSet bool
}

func (v NullableAiMessage) Get() *AiMessage {
	return v.value
}

func (v *NullableAiMessage) Set(val *AiMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMessage(val *AiMessage) *NullableAiMessage {
	return &NullableAiMessage{value: val, isSet: true}
}

func (v NullableAiMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
