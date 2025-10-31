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

// checks if the RequestDecisionCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestDecisionCreatable{}

// RequestDecisionCreatable A completed access request approval decision
type RequestDecisionCreatable struct {
	// The Okta user `id`
	UserId string `json:"userId" validate:"regexp=00u[0-9a-zA-Z]+"`
	// E-mail of the user.
	UserEmail string `json:"userEmail"`
	// Name of the user.
	UserName string               `json:"userName"`
	Decision ApprovalDecisionEnum `json:"decision"`
	// The date the approval decision is made.
	Decided time.Time `json:"decided"`
	// The Okta user `id`
	OriginalDeciderId *string `json:"originalDeciderId,omitempty" validate:"regexp=00u[0-9a-zA-Z]+"`
	// Name of the user.
	OriginalDeciderFullName *string `json:"originalDeciderFullName,omitempty"`
	// E-mail of the user.
	OriginalDeciderEmail *string `json:"originalDeciderEmail,omitempty"`
	// Indicates if the decision was made by a delegated decider
	DeciderDelegated     *bool `json:"deciderDelegated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestDecisionCreatable RequestDecisionCreatable

// NewRequestDecisionCreatable instantiates a new RequestDecisionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestDecisionCreatable(userId string, userEmail string, userName string, decision ApprovalDecisionEnum, decided time.Time) *RequestDecisionCreatable {
	this := RequestDecisionCreatable{}
	this.UserId = userId
	this.UserEmail = userEmail
	this.UserName = userName
	this.Decision = decision
	this.Decided = decided
	return &this
}

// NewRequestDecisionCreatableWithDefaults instantiates a new RequestDecisionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestDecisionCreatableWithDefaults() *RequestDecisionCreatable {
	this := RequestDecisionCreatable{}
	return &this
}

// GetUserId returns the UserId field value
func (o *RequestDecisionCreatable) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *RequestDecisionCreatable) SetUserId(v string) {
	o.UserId = v
}

// GetUserEmail returns the UserEmail field value
func (o *RequestDecisionCreatable) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value
func (o *RequestDecisionCreatable) SetUserEmail(v string) {
	o.UserEmail = v
}

// GetUserName returns the UserName field value
func (o *RequestDecisionCreatable) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *RequestDecisionCreatable) SetUserName(v string) {
	o.UserName = v
}

// GetDecision returns the Decision field value
func (o *RequestDecisionCreatable) GetDecision() ApprovalDecisionEnum {
	if o == nil {
		var ret ApprovalDecisionEnum
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetDecisionOk() (*ApprovalDecisionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *RequestDecisionCreatable) SetDecision(v ApprovalDecisionEnum) {
	o.Decision = v
}

// GetDecided returns the Decided field value
func (o *RequestDecisionCreatable) GetDecided() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Decided
}

// GetDecidedOk returns a tuple with the Decided field value
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetDecidedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decided, true
}

// SetDecided sets field value
func (o *RequestDecisionCreatable) SetDecided(v time.Time) {
	o.Decided = v
}

// GetOriginalDeciderId returns the OriginalDeciderId field value if set, zero value otherwise.
func (o *RequestDecisionCreatable) GetOriginalDeciderId() string {
	if o == nil || IsNil(o.OriginalDeciderId) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderId
}

// GetOriginalDeciderIdOk returns a tuple with the OriginalDeciderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetOriginalDeciderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderId) {
		return nil, false
	}
	return o.OriginalDeciderId, true
}

// HasOriginalDeciderId returns a boolean if a field has been set.
func (o *RequestDecisionCreatable) HasOriginalDeciderId() bool {
	if o != nil && !IsNil(o.OriginalDeciderId) {
		return true
	}

	return false
}

// SetOriginalDeciderId gets a reference to the given string and assigns it to the OriginalDeciderId field.
func (o *RequestDecisionCreatable) SetOriginalDeciderId(v string) {
	o.OriginalDeciderId = &v
}

// GetOriginalDeciderFullName returns the OriginalDeciderFullName field value if set, zero value otherwise.
func (o *RequestDecisionCreatable) GetOriginalDeciderFullName() string {
	if o == nil || IsNil(o.OriginalDeciderFullName) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderFullName
}

// GetOriginalDeciderFullNameOk returns a tuple with the OriginalDeciderFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetOriginalDeciderFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderFullName) {
		return nil, false
	}
	return o.OriginalDeciderFullName, true
}

// HasOriginalDeciderFullName returns a boolean if a field has been set.
func (o *RequestDecisionCreatable) HasOriginalDeciderFullName() bool {
	if o != nil && !IsNil(o.OriginalDeciderFullName) {
		return true
	}

	return false
}

// SetOriginalDeciderFullName gets a reference to the given string and assigns it to the OriginalDeciderFullName field.
func (o *RequestDecisionCreatable) SetOriginalDeciderFullName(v string) {
	o.OriginalDeciderFullName = &v
}

// GetOriginalDeciderEmail returns the OriginalDeciderEmail field value if set, zero value otherwise.
func (o *RequestDecisionCreatable) GetOriginalDeciderEmail() string {
	if o == nil || IsNil(o.OriginalDeciderEmail) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderEmail
}

// GetOriginalDeciderEmailOk returns a tuple with the OriginalDeciderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetOriginalDeciderEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderEmail) {
		return nil, false
	}
	return o.OriginalDeciderEmail, true
}

// HasOriginalDeciderEmail returns a boolean if a field has been set.
func (o *RequestDecisionCreatable) HasOriginalDeciderEmail() bool {
	if o != nil && !IsNil(o.OriginalDeciderEmail) {
		return true
	}

	return false
}

// SetOriginalDeciderEmail gets a reference to the given string and assigns it to the OriginalDeciderEmail field.
func (o *RequestDecisionCreatable) SetOriginalDeciderEmail(v string) {
	o.OriginalDeciderEmail = &v
}

// GetDeciderDelegated returns the DeciderDelegated field value if set, zero value otherwise.
func (o *RequestDecisionCreatable) GetDeciderDelegated() bool {
	if o == nil || IsNil(o.DeciderDelegated) {
		var ret bool
		return ret
	}
	return *o.DeciderDelegated
}

// GetDeciderDelegatedOk returns a tuple with the DeciderDelegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDecisionCreatable) GetDeciderDelegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeciderDelegated) {
		return nil, false
	}
	return o.DeciderDelegated, true
}

// HasDeciderDelegated returns a boolean if a field has been set.
func (o *RequestDecisionCreatable) HasDeciderDelegated() bool {
	if o != nil && !IsNil(o.DeciderDelegated) {
		return true
	}

	return false
}

// SetDeciderDelegated gets a reference to the given bool and assigns it to the DeciderDelegated field.
func (o *RequestDecisionCreatable) SetDeciderDelegated(v bool) {
	o.DeciderDelegated = &v
}

func (o RequestDecisionCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestDecisionCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["userEmail"] = o.UserEmail
	toSerialize["userName"] = o.UserName
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestDecisionCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
		"userEmail",
		"userName",
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

	varRequestDecisionCreatable := _RequestDecisionCreatable{}

	err = json.Unmarshal(data, &varRequestDecisionCreatable)

	if err != nil {
		return err
	}

	*o = RequestDecisionCreatable(varRequestDecisionCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "userEmail")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "decided")
		delete(additionalProperties, "originalDeciderId")
		delete(additionalProperties, "originalDeciderFullName")
		delete(additionalProperties, "originalDeciderEmail")
		delete(additionalProperties, "deciderDelegated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestDecisionCreatable struct {
	value *RequestDecisionCreatable
	isSet bool
}

func (v NullableRequestDecisionCreatable) Get() *RequestDecisionCreatable {
	return v.value
}

func (v *NullableRequestDecisionCreatable) Set(val *RequestDecisionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestDecisionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestDecisionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestDecisionCreatable(val *RequestDecisionCreatable) *NullableRequestDecisionCreatable {
	return &NullableRequestDecisionCreatable{value: val, isSet: true}
}

func (v NullableRequestDecisionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestDecisionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
