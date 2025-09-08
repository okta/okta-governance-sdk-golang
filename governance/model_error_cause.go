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
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrorCause) HasReason() bool {
	if o != nil && o.Reason != nil {
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
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ErrorCause) HasLocation() bool {
	if o != nil && o.Location != nil {
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
	if o == nil || o.LocationType == nil {
		var ret string
		return ret
	}
	return *o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetLocationTypeOk() (*string, bool) {
	if o == nil || o.LocationType == nil {
		return nil, false
	}
	return o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *ErrorCause) HasLocationType() bool {
	if o != nil && o.LocationType != nil {
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
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ErrorCause) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ErrorCause) SetDomain(v string) {
	o.Domain = &v
}

func (o ErrorCause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errorSummary"] = o.ErrorSummary
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.LocationType != nil {
		toSerialize["locationType"] = o.LocationType
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ErrorCause) UnmarshalJSON(bytes []byte) (err error) {
	varErrorCause := _ErrorCause{}

	err = json.Unmarshal(bytes, &varErrorCause)
	if err == nil {
		*o = ErrorCause(varErrorCause)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "location")
		delete(additionalProperties, "locationType")
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
