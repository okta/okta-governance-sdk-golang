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

// checks if the RequestOnBehalfOfSettingsDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestOnBehalfOfSettingsDetails{}

// RequestOnBehalfOfSettingsDetails Specifies if and for whom a requester may request the resource for.
type RequestOnBehalfOfSettingsDetails struct {
	// Indicates that users who can request this resource could also request for another requester of the same resource
	Allowed bool `json:"allowed"`
	// Which requesters the resource requester can request on behalf of. If onlyFor is not specified then any requester may request a resource on the behalf of any other user
	OnlyFor              []RequestOnBehalfOfType `json:"onlyFor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestOnBehalfOfSettingsDetails RequestOnBehalfOfSettingsDetails

// NewRequestOnBehalfOfSettingsDetails instantiates a new RequestOnBehalfOfSettingsDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestOnBehalfOfSettingsDetails(allowed bool) *RequestOnBehalfOfSettingsDetails {
	this := RequestOnBehalfOfSettingsDetails{}
	this.Allowed = allowed
	return &this
}

// NewRequestOnBehalfOfSettingsDetailsWithDefaults instantiates a new RequestOnBehalfOfSettingsDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestOnBehalfOfSettingsDetailsWithDefaults() *RequestOnBehalfOfSettingsDetails {
	this := RequestOnBehalfOfSettingsDetails{}
	return &this
}

// GetAllowed returns the Allowed field value
func (o *RequestOnBehalfOfSettingsDetails) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *RequestOnBehalfOfSettingsDetails) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *RequestOnBehalfOfSettingsDetails) SetAllowed(v bool) {
	o.Allowed = v
}

// GetOnlyFor returns the OnlyFor field value if set, zero value otherwise.
func (o *RequestOnBehalfOfSettingsDetails) GetOnlyFor() []RequestOnBehalfOfType {
	if o == nil || IsNil(o.OnlyFor) {
		var ret []RequestOnBehalfOfType
		return ret
	}
	return o.OnlyFor
}

// GetOnlyForOk returns a tuple with the OnlyFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestOnBehalfOfSettingsDetails) GetOnlyForOk() ([]RequestOnBehalfOfType, bool) {
	if o == nil || IsNil(o.OnlyFor) {
		return nil, false
	}
	return o.OnlyFor, true
}

// HasOnlyFor returns a boolean if a field has been set.
func (o *RequestOnBehalfOfSettingsDetails) HasOnlyFor() bool {
	if o != nil && !IsNil(o.OnlyFor) {
		return true
	}

	return false
}

// SetOnlyFor gets a reference to the given []RequestOnBehalfOfType and assigns it to the OnlyFor field.
func (o *RequestOnBehalfOfSettingsDetails) SetOnlyFor(v []RequestOnBehalfOfType) {
	o.OnlyFor = v
}

func (o RequestOnBehalfOfSettingsDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestOnBehalfOfSettingsDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed"] = o.Allowed
	if !IsNil(o.OnlyFor) {
		toSerialize["onlyFor"] = o.OnlyFor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestOnBehalfOfSettingsDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowed",
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

	varRequestOnBehalfOfSettingsDetails := _RequestOnBehalfOfSettingsDetails{}

	err = json.Unmarshal(data, &varRequestOnBehalfOfSettingsDetails)

	if err != nil {
		return err
	}

	*o = RequestOnBehalfOfSettingsDetails(varRequestOnBehalfOfSettingsDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowed")
		delete(additionalProperties, "onlyFor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestOnBehalfOfSettingsDetails struct {
	value *RequestOnBehalfOfSettingsDetails
	isSet bool
}

func (v NullableRequestOnBehalfOfSettingsDetails) Get() *RequestOnBehalfOfSettingsDetails {
	return v.value
}

func (v *NullableRequestOnBehalfOfSettingsDetails) Set(val *RequestOnBehalfOfSettingsDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestOnBehalfOfSettingsDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestOnBehalfOfSettingsDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestOnBehalfOfSettingsDetails(val *RequestOnBehalfOfSettingsDetails) *NullableRequestOnBehalfOfSettingsDetails {
	return &NullableRequestOnBehalfOfSettingsDetails{value: val, isSet: true}
}

func (v NullableRequestOnBehalfOfSettingsDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestOnBehalfOfSettingsDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
