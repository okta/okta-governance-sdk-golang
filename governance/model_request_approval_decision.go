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
	"time"
)

// checks if the RequestApprovalDecision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestApprovalDecision{}

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
	DeciderDelegated *bool `json:"deciderDelegated,omitempty"`
	// Indicates if the decision was made by an escalated decider
	DeciderEscalated     *bool `json:"deciderEscalated,omitempty"`
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
	if o == nil || IsNil(o.DeciderName) {
		var ret string
		return ret
	}
	return *o.DeciderName
}

// GetDeciderNameOk returns a tuple with the DeciderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDeciderNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeciderName) {
		return nil, false
	}
	return o.DeciderName, true
}

// HasDeciderName returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasDeciderName() bool {
	if o != nil && !IsNil(o.DeciderName) {
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
	if o == nil || IsNil(o.OriginalDeciderId) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderId
}

// GetOriginalDeciderIdOk returns a tuple with the OriginalDeciderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetOriginalDeciderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderId) {
		return nil, false
	}
	return o.OriginalDeciderId, true
}

// HasOriginalDeciderId returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasOriginalDeciderId() bool {
	if o != nil && !IsNil(o.OriginalDeciderId) {
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
	if o == nil || IsNil(o.OriginalDeciderFullName) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderFullName
}

// GetOriginalDeciderFullNameOk returns a tuple with the OriginalDeciderFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetOriginalDeciderFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderFullName) {
		return nil, false
	}
	return o.OriginalDeciderFullName, true
}

// HasOriginalDeciderFullName returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasOriginalDeciderFullName() bool {
	if o != nil && !IsNil(o.OriginalDeciderFullName) {
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
	if o == nil || IsNil(o.OriginalDeciderEmail) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderEmail
}

// GetOriginalDeciderEmailOk returns a tuple with the OriginalDeciderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetOriginalDeciderEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderEmail) {
		return nil, false
	}
	return o.OriginalDeciderEmail, true
}

// HasOriginalDeciderEmail returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasOriginalDeciderEmail() bool {
	if o != nil && !IsNil(o.OriginalDeciderEmail) {
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
	if o == nil || IsNil(o.DeciderDelegated) {
		var ret bool
		return ret
	}
	return *o.DeciderDelegated
}

// GetDeciderDelegatedOk returns a tuple with the DeciderDelegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDeciderDelegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeciderDelegated) {
		return nil, false
	}
	return o.DeciderDelegated, true
}

// HasDeciderDelegated returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasDeciderDelegated() bool {
	if o != nil && !IsNil(o.DeciderDelegated) {
		return true
	}

	return false
}

// SetDeciderDelegated gets a reference to the given bool and assigns it to the DeciderDelegated field.
func (o *RequestApprovalDecision) SetDeciderDelegated(v bool) {
	o.DeciderDelegated = &v
}

// GetDeciderEscalated returns the DeciderEscalated field value if set, zero value otherwise.
func (o *RequestApprovalDecision) GetDeciderEscalated() bool {
	if o == nil || IsNil(o.DeciderEscalated) {
		var ret bool
		return ret
	}
	return *o.DeciderEscalated
}

// GetDeciderEscalatedOk returns a tuple with the DeciderEscalated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalDecision) GetDeciderEscalatedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeciderEscalated) {
		return nil, false
	}
	return o.DeciderEscalated, true
}

// HasDeciderEscalated returns a boolean if a field has been set.
func (o *RequestApprovalDecision) HasDeciderEscalated() bool {
	if o != nil && !IsNil(o.DeciderEscalated) {
		return true
	}

	return false
}

// SetDeciderEscalated gets a reference to the given bool and assigns it to the DeciderEscalated field.
func (o *RequestApprovalDecision) SetDeciderEscalated(v bool) {
	o.DeciderEscalated = &v
}

func (o RequestApprovalDecision) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestApprovalDecision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deciderId"] = o.DeciderId
	if !IsNil(o.DeciderName) {
		toSerialize["deciderName"] = o.DeciderName
	}
	toSerialize["decision"] = o.Decision
	toSerialize["decided"] = o.Decided
	if !IsNil(o.OriginalDeciderId) {
		toSerialize["originalDeciderId"] = o.OriginalDeciderId
	}
	if !IsNil(o.OriginalDeciderFullName) {
		toSerialize["originalDeciderFullName"] = o.OriginalDeciderFullName
	}
	if !IsNil(o.OriginalDeciderEmail) {
		toSerialize["originalDeciderEmail"] = o.OriginalDeciderEmail
	}
	if !IsNil(o.DeciderDelegated) {
		toSerialize["deciderDelegated"] = o.DeciderDelegated
	}
	if !IsNil(o.DeciderEscalated) {
		toSerialize["deciderEscalated"] = o.DeciderEscalated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestApprovalDecision) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deciderId",
		"decision",
		"decided",
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

	varRequestApprovalDecision := _RequestApprovalDecision{}

	err = json.Unmarshal(data, &varRequestApprovalDecision)

	if err != nil {
		return err
	}

	*o = RequestApprovalDecision(varRequestApprovalDecision)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deciderId")
		delete(additionalProperties, "deciderName")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "decided")
		delete(additionalProperties, "originalDeciderId")
		delete(additionalProperties, "originalDeciderFullName")
		delete(additionalProperties, "originalDeciderEmail")
		delete(additionalProperties, "deciderDelegated")
		delete(additionalProperties, "deciderEscalated")
		o.AdditionalProperties = additionalProperties
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
