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

// RequestApproval2 A completed access request approval
type RequestApproval2 struct {
	Type ApprovalProviderType `json:"type"`
	// A human readable name of the request approval system, for example, Okta Access Requests or ServiceNow
	ProviderName string `json:"providerName"`
	// A description of the request approval system
	ProviderDescription *string `json:"providerDescription,omitempty"`
	// The external request `id` from a request approval system, for example, ServiceNow or JIRA
	ExternalRequestId *string                `json:"externalRequestId,omitempty"`
	Status            *RequestApprovalStatus `json:"status,omitempty"`
	// The date the approval decision is made.
	Decided *time.Time `json:"decided,omitempty"`
	// The approval decisions
	Decisions            []RequestApprovalDecision `json:"decisions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestApproval2 RequestApproval2

// NewRequestApproval2 instantiates a new RequestApproval2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestApproval2(type_ ApprovalProviderType, providerName string) *RequestApproval2 {
	this := RequestApproval2{}
	this.Type = type_
	this.ProviderName = providerName
	return &this
}

// NewRequestApproval2WithDefaults instantiates a new RequestApproval2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestApproval2WithDefaults() *RequestApproval2 {
	this := RequestApproval2{}
	return &this
}

// GetType returns the Type field value
func (o *RequestApproval2) GetType() ApprovalProviderType {
	if o == nil {
		var ret ApprovalProviderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestApproval2) GetTypeOk() (*ApprovalProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestApproval2) SetType(v ApprovalProviderType) {
	o.Type = v
}

// GetProviderName returns the ProviderName field value
func (o *RequestApproval2) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *RequestApproval2) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *RequestApproval2) SetProviderName(v string) {
	o.ProviderName = v
}

// GetProviderDescription returns the ProviderDescription field value if set, zero value otherwise.
func (o *RequestApproval2) GetProviderDescription() string {
	if o == nil || o.ProviderDescription == nil {
		var ret string
		return ret
	}
	return *o.ProviderDescription
}

// GetProviderDescriptionOk returns a tuple with the ProviderDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApproval2) GetProviderDescriptionOk() (*string, bool) {
	if o == nil || o.ProviderDescription == nil {
		return nil, false
	}
	return o.ProviderDescription, true
}

// HasProviderDescription returns a boolean if a field has been set.
func (o *RequestApproval2) HasProviderDescription() bool {
	if o != nil && o.ProviderDescription != nil {
		return true
	}

	return false
}

// SetProviderDescription gets a reference to the given string and assigns it to the ProviderDescription field.
func (o *RequestApproval2) SetProviderDescription(v string) {
	o.ProviderDescription = &v
}

// GetExternalRequestId returns the ExternalRequestId field value if set, zero value otherwise.
func (o *RequestApproval2) GetExternalRequestId() string {
	if o == nil || o.ExternalRequestId == nil {
		var ret string
		return ret
	}
	return *o.ExternalRequestId
}

// GetExternalRequestIdOk returns a tuple with the ExternalRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApproval2) GetExternalRequestIdOk() (*string, bool) {
	if o == nil || o.ExternalRequestId == nil {
		return nil, false
	}
	return o.ExternalRequestId, true
}

// HasExternalRequestId returns a boolean if a field has been set.
func (o *RequestApproval2) HasExternalRequestId() bool {
	if o != nil && o.ExternalRequestId != nil {
		return true
	}

	return false
}

// SetExternalRequestId gets a reference to the given string and assigns it to the ExternalRequestId field.
func (o *RequestApproval2) SetExternalRequestId(v string) {
	o.ExternalRequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestApproval2) GetStatus() RequestApprovalStatus {
	if o == nil || o.Status == nil {
		var ret RequestApprovalStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApproval2) GetStatusOk() (*RequestApprovalStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestApproval2) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RequestApprovalStatus and assigns it to the Status field.
func (o *RequestApproval2) SetStatus(v RequestApprovalStatus) {
	o.Status = &v
}

// GetDecided returns the Decided field value if set, zero value otherwise.
func (o *RequestApproval2) GetDecided() time.Time {
	if o == nil || o.Decided == nil {
		var ret time.Time
		return ret
	}
	return *o.Decided
}

// GetDecidedOk returns a tuple with the Decided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApproval2) GetDecidedOk() (*time.Time, bool) {
	if o == nil || o.Decided == nil {
		return nil, false
	}
	return o.Decided, true
}

// HasDecided returns a boolean if a field has been set.
func (o *RequestApproval2) HasDecided() bool {
	if o != nil && o.Decided != nil {
		return true
	}

	return false
}

// SetDecided gets a reference to the given time.Time and assigns it to the Decided field.
func (o *RequestApproval2) SetDecided(v time.Time) {
	o.Decided = &v
}

// GetDecisions returns the Decisions field value if set, zero value otherwise.
func (o *RequestApproval2) GetDecisions() []RequestApprovalDecision {
	if o == nil || o.Decisions == nil {
		var ret []RequestApprovalDecision
		return ret
	}
	return o.Decisions
}

// GetDecisionsOk returns a tuple with the Decisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestApproval2) GetDecisionsOk() ([]RequestApprovalDecision, bool) {
	if o == nil || o.Decisions == nil {
		return nil, false
	}
	return o.Decisions, true
}

// HasDecisions returns a boolean if a field has been set.
func (o *RequestApproval2) HasDecisions() bool {
	if o != nil && o.Decisions != nil {
		return true
	}

	return false
}

// SetDecisions gets a reference to the given []RequestApprovalDecision and assigns it to the Decisions field.
func (o *RequestApproval2) SetDecisions(v []RequestApprovalDecision) {
	o.Decisions = v
}

func (o RequestApproval2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if o.ProviderDescription != nil {
		toSerialize["providerDescription"] = o.ProviderDescription
	}
	if o.ExternalRequestId != nil {
		toSerialize["externalRequestId"] = o.ExternalRequestId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Decided != nil {
		toSerialize["decided"] = o.Decided
	}
	if o.Decisions != nil {
		toSerialize["decisions"] = o.Decisions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestApproval2) UnmarshalJSON(bytes []byte) (err error) {
	varRequestApproval2 := _RequestApproval2{}

	err = json.Unmarshal(bytes, &varRequestApproval2)
	if err == nil {
		*o = RequestApproval2(varRequestApproval2)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "providerName")
		delete(additionalProperties, "providerDescription")
		delete(additionalProperties, "externalRequestId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "decided")
		delete(additionalProperties, "decisions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestApproval2 struct {
	value *RequestApproval2
	isSet bool
}

func (v NullableRequestApproval2) Get() *RequestApproval2 {
	return v.value
}

func (v *NullableRequestApproval2) Set(val *RequestApproval2) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApproval2) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApproval2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApproval2(val *RequestApproval2) *NullableRequestApproval2 {
	return &NullableRequestApproval2{value: val, isSet: true}
}

func (v NullableRequestApproval2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApproval2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
