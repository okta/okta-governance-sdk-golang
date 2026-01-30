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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AccessDurationSettingsRequesterSpecifiedDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessDurationSettingsRequesterSpecifiedDuration{}

// AccessDurationSettingsRequesterSpecifiedDuration Setting to specify the maximum duration that an end user can request access for when making a request.
type AccessDurationSettingsRequesterSpecifiedDuration struct {
	Type string `json:"type"`
	// The maximum duration set by the requester for access durations. Use ISO8061 notation for duration values, see https://tc39.es/proposal-temporal/docs/duration.html. The admin sets the maximum duration that can be requested and can't exceed 72 hours (`PT72H`), 90 days (`90D`), or 12 weeks (`P12W`). For example:    - 24 hours (`PT24H`)   - 7 days (`P7D`)   - 2 weeks (`P2W`)
	MaximumDuration string `json:"maximumDuration"`
}

type _AccessDurationSettingsRequesterSpecifiedDuration AccessDurationSettingsRequesterSpecifiedDuration

// NewAccessDurationSettingsRequesterSpecifiedDuration instantiates a new AccessDurationSettingsRequesterSpecifiedDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessDurationSettingsRequesterSpecifiedDuration(type_ string, maximumDuration string) *AccessDurationSettingsRequesterSpecifiedDuration {
	this := AccessDurationSettingsRequesterSpecifiedDuration{}
	this.Type = type_
	this.MaximumDuration = maximumDuration
	return &this
}

// NewAccessDurationSettingsRequesterSpecifiedDurationWithDefaults instantiates a new AccessDurationSettingsRequesterSpecifiedDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessDurationSettingsRequesterSpecifiedDurationWithDefaults() *AccessDurationSettingsRequesterSpecifiedDuration {
	this := AccessDurationSettingsRequesterSpecifiedDuration{}
	return &this
}

// GetType returns the Type field value
func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessDurationSettingsRequesterSpecifiedDuration) SetType(v string) {
	o.Type = v
}

// GetMaximumDuration returns the MaximumDuration field value
func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetMaximumDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaximumDuration
}

// GetMaximumDurationOk returns a tuple with the MaximumDuration field value
// and a boolean to check if the value has been set.
func (o *AccessDurationSettingsRequesterSpecifiedDuration) GetMaximumDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumDuration, true
}

// SetMaximumDuration sets field value
func (o *AccessDurationSettingsRequesterSpecifiedDuration) SetMaximumDuration(v string) {
	o.MaximumDuration = v
}

func (o AccessDurationSettingsRequesterSpecifiedDuration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessDurationSettingsRequesterSpecifiedDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["maximumDuration"] = o.MaximumDuration
	return toSerialize, nil
}

func (o *AccessDurationSettingsRequesterSpecifiedDuration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"maximumDuration",
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

	varAccessDurationSettingsRequesterSpecifiedDuration := _AccessDurationSettingsRequesterSpecifiedDuration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessDurationSettingsRequesterSpecifiedDuration)

	if err != nil {
		return err
	}

	*o = AccessDurationSettingsRequesterSpecifiedDuration(varAccessDurationSettingsRequesterSpecifiedDuration)

	return err
}

type NullableAccessDurationSettingsRequesterSpecifiedDuration struct {
	value *AccessDurationSettingsRequesterSpecifiedDuration
	isSet bool
}

func (v NullableAccessDurationSettingsRequesterSpecifiedDuration) Get() *AccessDurationSettingsRequesterSpecifiedDuration {
	return v.value
}

func (v *NullableAccessDurationSettingsRequesterSpecifiedDuration) Set(val *AccessDurationSettingsRequesterSpecifiedDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessDurationSettingsRequesterSpecifiedDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessDurationSettingsRequesterSpecifiedDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessDurationSettingsRequesterSpecifiedDuration(val *AccessDurationSettingsRequesterSpecifiedDuration) *NullableAccessDurationSettingsRequesterSpecifiedDuration {
	return &NullableAccessDurationSettingsRequesterSpecifiedDuration{value: val, isSet: true}
}

func (v NullableAccessDurationSettingsRequesterSpecifiedDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessDurationSettingsRequesterSpecifiedDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
