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

// ModelError An Error Object
type ModelError struct {
	// An error code unique to the error
	ErrorCode string `json:"errorCode"`
	// An error identifier useful for troubleshooting an error with support
	ErrorId string `json:"errorId"`
	// An error code description detailing the error
	ErrorSummary string `json:"errorSummary"`
	// An indicator where to look out to troubleshoot the error
	ErrorLink *string `json:"errorLink,omitempty"`
	// An optional array of string values that explains possible reasons for the error.
	ErrorCauses          []ErrorCause `json:"errorCauses,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelError ModelError

// NewModelError instantiates a new ModelError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelError(errorCode string, errorId string, errorSummary string) *ModelError {
	this := ModelError{}
	this.ErrorCode = errorCode
	this.ErrorId = errorId
	this.ErrorSummary = errorSummary
	return &this
}

// NewModelErrorWithDefaults instantiates a new ModelError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelErrorWithDefaults() *ModelError {
	this := ModelError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *ModelError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ModelError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorId returns the ErrorId field value
func (o *ModelError) GetErrorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorId, true
}

// SetErrorId sets field value
func (o *ModelError) SetErrorId(v string) {
	o.ErrorId = v
}

// GetErrorSummary returns the ErrorSummary field value
func (o *ModelError) GetErrorSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorSummary, true
}

// SetErrorSummary sets field value
func (o *ModelError) SetErrorSummary(v string) {
	o.ErrorSummary = v
}

// GetErrorLink returns the ErrorLink field value if set, zero value otherwise.
func (o *ModelError) GetErrorLink() string {
	if o == nil || o.ErrorLink == nil {
		var ret string
		return ret
	}
	return *o.ErrorLink
}

// GetErrorLinkOk returns a tuple with the ErrorLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorLinkOk() (*string, bool) {
	if o == nil || o.ErrorLink == nil {
		return nil, false
	}
	return o.ErrorLink, true
}

// HasErrorLink returns a boolean if a field has been set.
func (o *ModelError) HasErrorLink() bool {
	if o != nil && o.ErrorLink != nil {
		return true
	}

	return false
}

// SetErrorLink gets a reference to the given string and assigns it to the ErrorLink field.
func (o *ModelError) SetErrorLink(v string) {
	o.ErrorLink = &v
}

// GetErrorCauses returns the ErrorCauses field value if set, zero value otherwise.
func (o *ModelError) GetErrorCauses() []ErrorCause {
	if o == nil || o.ErrorCauses == nil {
		var ret []ErrorCause
		return ret
	}
	return o.ErrorCauses
}

// GetErrorCausesOk returns a tuple with the ErrorCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelError) GetErrorCausesOk() ([]ErrorCause, bool) {
	if o == nil || o.ErrorCauses == nil {
		return nil, false
	}
	return o.ErrorCauses, true
}

// HasErrorCauses returns a boolean if a field has been set.
func (o *ModelError) HasErrorCauses() bool {
	if o != nil && o.ErrorCauses != nil {
		return true
	}

	return false
}

// SetErrorCauses gets a reference to the given []ErrorCause and assigns it to the ErrorCauses field.
func (o *ModelError) SetErrorCauses(v []ErrorCause) {
	o.ErrorCauses = v
}

func (o ModelError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if true {
		toSerialize["errorId"] = o.ErrorId
	}
	if true {
		toSerialize["errorSummary"] = o.ErrorSummary
	}
	if o.ErrorLink != nil {
		toSerialize["errorLink"] = o.ErrorLink
	}
	if o.ErrorCauses != nil {
		toSerialize["errorCauses"] = o.ErrorCauses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ModelError) UnmarshalJSON(bytes []byte) (err error) {
	varModelError := _ModelError{}

	err = json.Unmarshal(bytes, &varModelError)
	if err == nil {
		*o = ModelError(varModelError)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorId")
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "errorLink")
		delete(additionalProperties, "errorCauses")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableModelError struct {
	value *ModelError
	isSet bool
}

func (v NullableModelError) Get() *ModelError {
	return v.value
}

func (v *NullableModelError) Set(val *ModelError) {
	v.value = val
	v.isSet = true
}

func (v NullableModelError) IsSet() bool {
	return v.isSet
}

func (v *NullableModelError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelError(val *ModelError) *NullableModelError {
	return &NullableModelError{value: val, isSet: true}
}

func (v NullableModelError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
