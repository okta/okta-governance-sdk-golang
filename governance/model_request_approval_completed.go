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

// checks if the RequestApprovalCompleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestApprovalCompleted{}

// RequestApprovalCompleted A completed access request approval
type RequestApprovalCompleted struct {
	// The status of a completed approval
	Status string `json:"status"`
	// Okta username of the user which approved the request
	ApproverName string `json:"approverName"`
	// A unique identifier of the approval in the request
	ApprovalId string `json:"approvalId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	// Okta User.id of the user which approved the request
	ApproverId string `json:"approverId" validate:"regexp=00u[0-9a-zA-Z]+"`
	Decision   string `json:"decision"`
	// The date the approver made their decision.
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
	DeciderEscalated *bool `json:"deciderEscalated,omitempty"`
	// Values to field prompts provided by the approver at the time of approval.  All approval fields specified in the related request type are represented in the same order as defined in the request type.
	FieldValues          []FieldValue `json:"fieldValues"`
	AdditionalProperties map[string]interface{}
}

type _RequestApprovalCompleted RequestApprovalCompleted

// NewRequestApprovalCompleted instantiates a new RequestApprovalCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestApprovalCompleted(status string, approverName string, approvalId string, approverId string, decision string, decided time.Time, fieldValues []FieldValue) *RequestApprovalCompleted {
	this := RequestApprovalCompleted{}
	this.Status = status
	this.ApproverName = approverName
	this.ApprovalId = approvalId
	this.ApproverId = approverId
	this.Decision = decision
	this.Decided = decided
	this.FieldValues = fieldValues
	return &this
}

// NewRequestApprovalCompletedWithDefaults instantiates a new RequestApprovalCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestApprovalCompletedWithDefaults() *RequestApprovalCompleted {
	this := RequestApprovalCompleted{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestApprovalCompleted) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestApprovalCompleted) SetStatus(v string) {
	o.Status = v
}

// GetApproverName returns the ApproverName field value
func (o *RequestApprovalCompleted) GetApproverName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverName
}

// GetApproverNameOk returns a tuple with the ApproverName field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetApproverNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverName, true
}

// SetApproverName sets field value
func (o *RequestApprovalCompleted) SetApproverName(v string) {
	o.ApproverName = v
}

// GetApprovalId returns the ApprovalId field value
func (o *RequestApprovalCompleted) GetApprovalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApprovalId
}

// GetApprovalIdOk returns a tuple with the ApprovalId field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetApprovalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalId, true
}

// SetApprovalId sets field value
func (o *RequestApprovalCompleted) SetApprovalId(v string) {
	o.ApprovalId = v
}

// GetApproverId returns the ApproverId field value
func (o *RequestApprovalCompleted) GetApproverId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetApproverIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverId, true
}

// SetApproverId sets field value
func (o *RequestApprovalCompleted) SetApproverId(v string) {
	o.ApproverId = v
}

// GetDecision returns the Decision field value
func (o *RequestApprovalCompleted) GetDecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetDecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *RequestApprovalCompleted) SetDecision(v string) {
	o.Decision = v
}

// GetDecided returns the Decided field value
func (o *RequestApprovalCompleted) GetDecided() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Decided
}

// GetDecidedOk returns a tuple with the Decided field value
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetDecidedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decided, true
}

// SetDecided sets field value
func (o *RequestApprovalCompleted) SetDecided(v time.Time) {
	o.Decided = v
}

// GetOriginalDeciderId returns the OriginalDeciderId field value if set, zero value otherwise.
func (o *RequestApprovalCompleted) GetOriginalDeciderId() string {
	if o == nil || IsNil(o.OriginalDeciderId) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderId
}

// GetOriginalDeciderIdOk returns a tuple with the OriginalDeciderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetOriginalDeciderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderId) {
		return nil, false
	}
	return o.OriginalDeciderId, true
}

// HasOriginalDeciderId returns a boolean if a field has been set.
func (o *RequestApprovalCompleted) HasOriginalDeciderId() bool {
	if o != nil && !IsNil(o.OriginalDeciderId) {
		return true
	}

	return false
}

// SetOriginalDeciderId gets a reference to the given string and assigns it to the OriginalDeciderId field.
func (o *RequestApprovalCompleted) SetOriginalDeciderId(v string) {
	o.OriginalDeciderId = &v
}

// GetOriginalDeciderFullName returns the OriginalDeciderFullName field value if set, zero value otherwise.
func (o *RequestApprovalCompleted) GetOriginalDeciderFullName() string {
	if o == nil || IsNil(o.OriginalDeciderFullName) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderFullName
}

// GetOriginalDeciderFullNameOk returns a tuple with the OriginalDeciderFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetOriginalDeciderFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderFullName) {
		return nil, false
	}
	return o.OriginalDeciderFullName, true
}

// HasOriginalDeciderFullName returns a boolean if a field has been set.
func (o *RequestApprovalCompleted) HasOriginalDeciderFullName() bool {
	if o != nil && !IsNil(o.OriginalDeciderFullName) {
		return true
	}

	return false
}

// SetOriginalDeciderFullName gets a reference to the given string and assigns it to the OriginalDeciderFullName field.
func (o *RequestApprovalCompleted) SetOriginalDeciderFullName(v string) {
	o.OriginalDeciderFullName = &v
}

// GetOriginalDeciderEmail returns the OriginalDeciderEmail field value if set, zero value otherwise.
func (o *RequestApprovalCompleted) GetOriginalDeciderEmail() string {
	if o == nil || IsNil(o.OriginalDeciderEmail) {
		var ret string
		return ret
	}
	return *o.OriginalDeciderEmail
}

// GetOriginalDeciderEmailOk returns a tuple with the OriginalDeciderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetOriginalDeciderEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalDeciderEmail) {
		return nil, false
	}
	return o.OriginalDeciderEmail, true
}

// HasOriginalDeciderEmail returns a boolean if a field has been set.
func (o *RequestApprovalCompleted) HasOriginalDeciderEmail() bool {
	if o != nil && !IsNil(o.OriginalDeciderEmail) {
		return true
	}

	return false
}

// SetOriginalDeciderEmail gets a reference to the given string and assigns it to the OriginalDeciderEmail field.
func (o *RequestApprovalCompleted) SetOriginalDeciderEmail(v string) {
	o.OriginalDeciderEmail = &v
}

// GetDeciderDelegated returns the DeciderDelegated field value if set, zero value otherwise.
func (o *RequestApprovalCompleted) GetDeciderDelegated() bool {
	if o == nil || IsNil(o.DeciderDelegated) {
		var ret bool
		return ret
	}
	return *o.DeciderDelegated
}

// GetDeciderDelegatedOk returns a tuple with the DeciderDelegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetDeciderDelegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeciderDelegated) {
		return nil, false
	}
	return o.DeciderDelegated, true
}

// HasDeciderDelegated returns a boolean if a field has been set.
func (o *RequestApprovalCompleted) HasDeciderDelegated() bool {
	if o != nil && !IsNil(o.DeciderDelegated) {
		return true
	}

	return false
}

// SetDeciderDelegated gets a reference to the given bool and assigns it to the DeciderDelegated field.
func (o *RequestApprovalCompleted) SetDeciderDelegated(v bool) {
	o.DeciderDelegated = &v
}

// GetDeciderEscalated returns the DeciderEscalated field value if set, zero value otherwise.
func (o *RequestApprovalCompleted) GetDeciderEscalated() bool {
	if o == nil || IsNil(o.DeciderEscalated) {
		var ret bool
		return ret
	}
	return *o.DeciderEscalated
}

// GetDeciderEscalatedOk returns a tuple with the DeciderEscalated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApprovalCompleted) GetDeciderEscalatedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeciderEscalated) {
		return nil, false
	}
	return o.DeciderEscalated, true
}

// HasDeciderEscalated returns a boolean if a field has been set.
func (o *RequestApprovalCompleted) HasDeciderEscalated() bool {
	if o != nil && !IsNil(o.DeciderEscalated) {
		return true
	}

	return false
}

// SetDeciderEscalated gets a reference to the given bool and assigns it to the DeciderEscalated field.
func (o *RequestApprovalCompleted) SetDeciderEscalated(v bool) {
	o.DeciderEscalated = &v
}

// GetFieldValues returns the FieldValues field value
// If the value is explicit nil, the zero value for []FieldValue will be returned
func (o *RequestApprovalCompleted) GetFieldValues() []FieldValue {
	if o == nil {
		var ret []FieldValue
		return ret
	}

	return o.FieldValues
}

// GetFieldValuesOk returns a tuple with the FieldValues field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestApprovalCompleted) GetFieldValuesOk() ([]FieldValue, bool) {
	if o == nil || IsNil(o.FieldValues) {
		return nil, false
	}
	return o.FieldValues, true
}

// SetFieldValues sets field value
func (o *RequestApprovalCompleted) SetFieldValues(v []FieldValue) {
	o.FieldValues = v
}

func (o RequestApprovalCompleted) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestApprovalCompleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["approverName"] = o.ApproverName
	toSerialize["approvalId"] = o.ApprovalId
	toSerialize["approverId"] = o.ApproverId
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
	if o.FieldValues != nil {
		toSerialize["fieldValues"] = o.FieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestApprovalCompleted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"approverName",
		"approvalId",
		"approverId",
		"decision",
		"decided",
		"fieldValues",
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

	varRequestApprovalCompleted := _RequestApprovalCompleted{}

	err = json.Unmarshal(data, &varRequestApprovalCompleted)

	if err != nil {
		return err
	}

	*o = RequestApprovalCompleted(varRequestApprovalCompleted)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "approverName")
		delete(additionalProperties, "approvalId")
		delete(additionalProperties, "approverId")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "decided")
		delete(additionalProperties, "originalDeciderId")
		delete(additionalProperties, "originalDeciderFullName")
		delete(additionalProperties, "originalDeciderEmail")
		delete(additionalProperties, "deciderDelegated")
		delete(additionalProperties, "deciderEscalated")
		delete(additionalProperties, "fieldValues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestApprovalCompleted struct {
	value *RequestApprovalCompleted
	isSet bool
}

func (v NullableRequestApprovalCompleted) Get() *RequestApprovalCompleted {
	return v.value
}

func (v *NullableRequestApprovalCompleted) Set(val *RequestApprovalCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApprovalCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApprovalCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApprovalCompleted(val *RequestApprovalCompleted) *NullableRequestApprovalCompleted {
	return &NullableRequestApprovalCompleted{value: val, isSet: true}
}

func (v NullableRequestApprovalCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApprovalCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
