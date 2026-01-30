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

// checks if the ErrorCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorCause{}

// ErrorCause A specific cause of an error
type ErrorCause struct {
	// A more specific summary for the error containing the error cause.
	ErrorSummary string `json:"errorSummary"`
	// An enumerated value to represent the reason why the error occurred. This enumeration allows codes to adapt to different conditions in which an error code can occur.
	Reason *string `json:"reason,omitempty"`
	// A value that represents the key where the error cause occurred. This is used with `locationType` to give a holistic view of where the error cause occurred. For example, if `locationType` is body and the location is username and the reason was UNIQUE_CONSTRAINT, you can derive that the username was already taken.
	Location *string `json:"location,omitempty"`
	// A value that represents where the error cause occurred. For example, in the body or header of the request. This value is not required for cases where the request is correct, but there was another reason why the error occurred (server-side state conflict, rate limit violation, and so on)
	LocationType *string `json:"locationType,omitempty"`
	// A value that represents the domain of the service in which the error occurs. This value is used to isolate the error cause reason.
	Domain               *string `json:"domain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorCause ErrorCause

// NewErrorCause instantiates a new ErrorCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorCause(errorSummary string) *ErrorCause {
	this := ErrorCause{}
	this.ErrorSummary = errorSummary
	return &this
}

// NewErrorCauseWithDefaults instantiates a new ErrorCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorCauseWithDefaults() *ErrorCause {
	this := ErrorCause{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value
func (o *ErrorCause) GetErrorSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetErrorSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorSummary, true
}

// SetErrorSummary sets field value
func (o *ErrorCause) SetErrorSummary(v string) {
	o.ErrorSummary = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrorCause) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorCause) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrorCause) SetReason(v string) {
	o.Reason = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ErrorCause) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ErrorCause) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ErrorCause) SetLocation(v string) {
	o.Location = &v
}

// GetLocationType returns the LocationType field value if set, zero value otherwise.
func (o *ErrorCause) GetLocationType() string {
	if o == nil || IsNil(o.LocationType) {
		var ret string
		return ret
	}
	return *o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetLocationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LocationType) {
		return nil, false
	}
	return o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *ErrorCause) HasLocationType() bool {
	if o != nil && !IsNil(o.LocationType) {
		return true
	}

	return false
}

// SetLocationType gets a reference to the given string and assigns it to the LocationType field.
func (o *ErrorCause) SetLocationType(v string) {
	o.LocationType = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ErrorCause) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ErrorCause) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ErrorCause) SetDomain(v string) {
	o.Domain = &v
}

func (o ErrorCause) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorSummary"] = o.ErrorSummary
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.LocationType) {
		toSerialize["locationType"] = o.LocationType
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorCause) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"errorSummary",
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

	varErrorCause := _ErrorCause{}

	err = json.Unmarshal(data, &varErrorCause)

	if err != nil {
		return err
	}

	*o = ErrorCause(varErrorCause)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "location")
		delete(additionalProperties, "locationType")
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorCause struct {
	value *ErrorCause
	isSet bool
}

func (v NullableErrorCause) Get() *ErrorCause {
	return v.value
}

func (v *NullableErrorCause) Set(val *ErrorCause) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCause) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCause(val *ErrorCause) *NullableErrorCause {
	return &NullableErrorCause{value: val, isSet: true}
}

func (v NullableErrorCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
