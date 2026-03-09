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

// checks if the AiSummaryStreamSseMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSummaryStreamSseMessage{}

// AiSummaryStreamSseMessage The server-sent events (SSE) format: * The `data` property contains a JSON string. * The `data` property contains chunks of the AI-generated summary.
type AiSummaryStreamSseMessage struct {
	// The event ID
	Id *string `json:"id,omitempty"`
	// The event type
	Event *string `json:"event,omitempty"`
	// The event data, in JSON string format
	Data                 string `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AiSummaryStreamSseMessage AiSummaryStreamSseMessage

// NewAiSummaryStreamSseMessage instantiates a new AiSummaryStreamSseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSummaryStreamSseMessage(data string) *AiSummaryStreamSseMessage {
	this := AiSummaryStreamSseMessage{}
	this.Data = data
	return &this
}

// NewAiSummaryStreamSseMessageWithDefaults instantiates a new AiSummaryStreamSseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSummaryStreamSseMessageWithDefaults() *AiSummaryStreamSseMessage {
	this := AiSummaryStreamSseMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiSummaryStreamSseMessage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSummaryStreamSseMessage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiSummaryStreamSseMessage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AiSummaryStreamSseMessage) SetId(v string) {
	o.Id = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AiSummaryStreamSseMessage) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSummaryStreamSseMessage) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AiSummaryStreamSseMessage) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *AiSummaryStreamSseMessage) SetEvent(v string) {
	o.Event = &v
}

// GetData returns the Data field value
func (o *AiSummaryStreamSseMessage) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AiSummaryStreamSseMessage) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AiSummaryStreamSseMessage) SetData(v string) {
	o.Data = v
}

func (o AiSummaryStreamSseMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSummaryStreamSseMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiSummaryStreamSseMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varAiSummaryStreamSseMessage := _AiSummaryStreamSseMessage{}

	err = json.Unmarshal(data, &varAiSummaryStreamSseMessage)

	if err != nil {
		return err
	}

	*o = AiSummaryStreamSseMessage(varAiSummaryStreamSseMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "event")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiSummaryStreamSseMessage struct {
	value *AiSummaryStreamSseMessage
	isSet bool
}

func (v NullableAiSummaryStreamSseMessage) Get() *AiSummaryStreamSseMessage {
	return v.value
}

func (v *NullableAiSummaryStreamSseMessage) Set(val *AiSummaryStreamSseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSummaryStreamSseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSummaryStreamSseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSummaryStreamSseMessage(val *AiSummaryStreamSseMessage) *NullableAiSummaryStreamSseMessage {
	return &NullableAiSummaryStreamSseMessage{value: val, isSet: true}
}

func (v NullableAiSummaryStreamSseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSummaryStreamSseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
