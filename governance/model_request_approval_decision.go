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
	"time"
)

// RequestApprovalDecision A completed access request approval decision
type RequestApprovalDecision struct {
	// The Okta user `id`
	DeciderId string `json:"deciderId" validate:"regexp=00u[0-9a-zA-Z]+"`
	// Okta username of the decider
	DeciderName *string              `json:"deciderName,omitempty"`
	Decision    ApprovalDecisionEnum `json:"decision"`
	// The date the approval decision is made.
	Decided time.Time `json:"decided"`
	// The Okta user `id`
	OriginalDeciderId *string `json:"originalDeciderId,omitempty" validate:"regexp=00u[0-9a-zA-Z]+"`
	// Full name of the original decider
	OriginalDeciderFullName *string `json:"originalDeciderFullName,omitempty"`
	// Email of the original decider
	OriginalDeciderEmail *string `json:"originalDeciderEmail,omitempty"`
	// Indicates if the decision was made by a delegated decider
	DeciderDelegated     *bool `json:"deciderDelegated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestApprovalDecision RequestApprovalDecision

// NewRequestApprovalDecision instantiates a new RequestApprovalDecision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestApprovalDecision(deciderId string, decision ApprovalDecisionEnum, decided time.Time) *RequestApprovalDecision {
	this := RequestApprovalDecision{}
	this.DeciderId = deciderId
	this.Decision = decision
	this.Decided = decided
	return &this
}

// NewRequestApprovalDecisionWithDefaults instantiates a new RequestApprovalDecision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestApprovalDecisionWithDefaults() *RequestApprovalDecision {
	this := RequestApprovalDecision{}
	return &this
}

// GetDeciderId returns the DeciderId field value
func (o *RequestApprovalDecision) GetDeciderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeciderId
}

// GetDeciderIdOk returns a tuple with the DeciderId field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDeciderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeciderId, true
}

// SetDeciderId sets field value
func (o *RequestApprovalDecision) SetDeciderId(v string) {
	o.DeciderId = v
}

// GetDeciderName returns the DeciderName field value if set, zero value otherwise.
func (o *RequestApprovalDecision) GetDeciderName() string {
	if o == nil || o.DeciderName == nil {
		var ret string
		return ret
	}
	return *o.DeciderName
}

// GetDeciderNameOk returns a tuple with the DeciderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDeciderNameOk() (*string, bool) {
	if o == nil || o.DeciderName == nil {
		return nil, false
	}
	return o.DeciderName, true
}

// HasDeciderName returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasDeciderName() bool {
	if o != nil && o.DeciderName != nil {
		return true
	}

	return false
}

// SetDeciderName gets a reference to the given string and assigns it to the DeciderName field.
func (o *RequestApprovalDecision) SetDeciderName(v string) {
	o.DeciderName = &v
}

// GetDecision returns the Decision field value
func (o *RequestApprovalDecision) GetDecision() ApprovalDecisionEnum {
	if o == nil {
		var ret ApprovalDecisionEnum
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDecisionOk() (*ApprovalDecisionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *RequestApprovalDecision) SetDecision(v ApprovalDecisionEnum) {
	o.Decision = v
}

// GetDecided returns the Decided field value
func (o *RequestApprovalDecision) GetDecided() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Decided
}

// GetDecidedOk returns a tuple with the Decided field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDecidedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decided, true
}

// SetDecided sets field value
func (o *RequestApprovalDecision) SetDecided(v time.Time) {
	o.Decided = v
}

// GetOriginalDeciderId returns the OriginalDeciderId field value if set, zero value otherwise.
func (o *RequestApprovalDecision) GetOriginalDeciderId() string {
	if o == nil || o.OriginalDeciderId == nil {
		var ret string
		return ret
	}
	return *o.OriginalDeciderId
}

// GetOriginalDeciderIdOk returns a tuple with the OriginalDeciderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetOriginalDeciderIdOk() (*string, bool) {
	if o == nil || o.OriginalDeciderId == nil {
		return nil, false
	}
	return o.OriginalDeciderId, true
}

// HasOriginalDeciderId returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasOriginalDeciderId() bool {
	if o != nil && o.OriginalDeciderId != nil {
		return true
	}

	return false
}

// SetOriginalDeciderId gets a reference to the given string and assigns it to the OriginalDeciderId field.
func (o *RequestApprovalDecision) SetOriginalDeciderId(v string) {
	o.OriginalDeciderId = &v
}

// GetOriginalDeciderFullName returns the OriginalDeciderFullName field value if set, zero value otherwise.
func (o *RequestApprovalDecision) GetOriginalDeciderFullName() string {
	if o == nil || o.OriginalDeciderFullName == nil {
		var ret string
		return ret
	}
	return *o.OriginalDeciderFullName
}

// GetOriginalDeciderFullNameOk returns a tuple with the OriginalDeciderFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetOriginalDeciderFullNameOk() (*string, bool) {
	if o == nil || o.OriginalDeciderFullName == nil {
		return nil, false
	}
	return o.OriginalDeciderFullName, true
}

// HasOriginalDeciderFullName returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasOriginalDeciderFullName() bool {
	if o != nil && o.OriginalDeciderFullName != nil {
		return true
	}

	return false
}

// SetOriginalDeciderFullName gets a reference to the given string and assigns it to the OriginalDeciderFullName field.
func (o *RequestApprovalDecision) SetOriginalDeciderFullName(v string) {
	o.OriginalDeciderFullName = &v
}

// GetOriginalDeciderEmail returns the OriginalDeciderEmail field value if set, zero value otherwise.
func (o *RequestApprovalDecision) GetOriginalDeciderEmail() string {
	if o == nil || o.OriginalDeciderEmail == nil {
		var ret string
		return ret
	}
	return *o.OriginalDeciderEmail
}

// GetOriginalDeciderEmailOk returns a tuple with the OriginalDeciderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetOriginalDeciderEmailOk() (*string, bool) {
	if o == nil || o.OriginalDeciderEmail == nil {
		return nil, false
	}
	return o.OriginalDeciderEmail, true
}

// HasOriginalDeciderEmail returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasOriginalDeciderEmail() bool {
	if o != nil && o.OriginalDeciderEmail != nil {
		return true
	}

	return false
}

// SetOriginalDeciderEmail gets a reference to the given string and assigns it to the OriginalDeciderEmail field.
func (o *RequestApprovalDecision) SetOriginalDeciderEmail(v string) {
	o.OriginalDeciderEmail = &v
}

// GetDeciderDelegated returns the DeciderDelegated field value if set, zero value otherwise.
func (o *RequestApprovalDecision) GetDeciderDelegated() bool {
	if o == nil || o.DeciderDelegated == nil {
		var ret bool
		return ret
	}
	return *o.DeciderDelegated
}

// GetDeciderDelegatedOk returns a tuple with the DeciderDelegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDeciderDelegatedOk() (*bool, bool) {
	if o == nil || o.DeciderDelegated == nil {
		return nil, false
	}
	return o.DeciderDelegated, true
}

// HasDeciderDelegated returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasDeciderDelegated() bool {
	if o != nil && o.DeciderDelegated != nil {
		return true
	}

	return false
}

// SetDeciderDelegated gets a reference to the given bool and assigns it to the DeciderDelegated field.
func (o *RequestApprovalDecision) SetDeciderDelegated(v bool) {
	o.DeciderDelegated = &v
}

func (o RequestApprovalDecision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deciderId"] = o.DeciderId
	}
	if o.DeciderName != nil {
		toSerialize["deciderName"] = o.DeciderName
	}
	if true {
		toSerialize["decision"] = o.Decision
	}
	if true {
		toSerialize["decided"] = o.Decided
	}
	if o.OriginalDeciderId != nil {
		toSerialize["originalDeciderId"] = o.OriginalDeciderId
	}
	if o.OriginalDeciderFullName != nil {
		toSerialize["originalDeciderFullName"] = o.OriginalDeciderFullName
	}
	if o.OriginalDeciderEmail != nil {
		toSerialize["originalDeciderEmail"] = o.OriginalDeciderEmail
	}
	if o.DeciderDelegated != nil {
		toSerialize["deciderDelegated"] = o.DeciderDelegated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestApprovalDecision) UnmarshalJSON(bytes []byte) (err error) {
	varRequestApprovalDecision := _RequestApprovalDecision{}

	err = json.Unmarshal(bytes, &varRequestApprovalDecision)
	if err == nil {
		*o = RequestApprovalDecision(varRequestApprovalDecision)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "deciderId")
		delete(additionalProperties, "deciderName")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "decided")
		delete(additionalProperties, "originalDeciderId")
		delete(additionalProperties, "originalDeciderFullName")
		delete(additionalProperties, "originalDeciderEmail")
		delete(additionalProperties, "deciderDelegated")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestApprovalDecision struct {
	value *RequestApprovalDecision
	isSet bool
}

func (v NullableRequestApprovalDecision) Get() *RequestApprovalDecision {
	return v.value
}

func (v *NullableRequestApprovalDecision) Set(val *RequestApprovalDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprovalDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprovalDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprovalDecision(val *RequestApprovalDecision) *NullableRequestApprovalDecision {
	return &NullableRequestApprovalDecision{value: val, isSet: true}
}

func (v NullableRequestApprovalDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprovalDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
