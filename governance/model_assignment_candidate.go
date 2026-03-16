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

// checks if the AssignmentCandidate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignmentCandidate{}

// AssignmentCandidate A candidate selected for assignment
type AssignmentCandidate struct {
	// ID of the candidate
	ExternalId string        `json:"externalId"`
	Type       CandidateType `json:"type"`
	// The Okta user or group `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.
	Orn                  *string                       `json:"orn,omitempty"`
	Delegator            *AssignmentCandidateDelegator `json:"delegator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignmentCandidate AssignmentCandidate

// NewAssignmentCandidate instantiates a new AssignmentCandidate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignmentCandidate(externalId string, type_ CandidateType) *AssignmentCandidate {
	this := AssignmentCandidate{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewAssignmentCandidateWithDefaults instantiates a new AssignmentCandidate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignmentCandidateWithDefaults() *AssignmentCandidate {
	this := AssignmentCandidate{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *AssignmentCandidate) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AssignmentCandidate) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *AssignmentCandidate) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *AssignmentCandidate) GetType() CandidateType {
	if o == nil {
		var ret CandidateType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssignmentCandidate) GetTypeOk() (*CandidateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssignmentCandidate) SetType(v CandidateType) {
	o.Type = v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *AssignmentCandidate) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentCandidate) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *AssignmentCandidate) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *AssignmentCandidate) SetOrn(v string) {
	o.Orn = &v
}

// GetDelegator returns the Delegator field value if set, zero value otherwise.
func (o *AssignmentCandidate) GetDelegator() AssignmentCandidateDelegator {
	if o == nil || IsNil(o.Delegator) {
		var ret AssignmentCandidateDelegator
		return ret
	}
	return *o.Delegator
}

// GetDelegatorOk returns a tuple with the Delegator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentCandidate) GetDelegatorOk() (*AssignmentCandidateDelegator, bool) {
	if o == nil || IsNil(o.Delegator) {
		return nil, false
	}
	return o.Delegator, true
}

// HasDelegator returns a boolean if a field has been set.
func (o *AssignmentCandidate) HasDelegator() bool {
	if o != nil && !IsNil(o.Delegator) {
		return true
	}

	return false
}

// SetDelegator gets a reference to the given AssignmentCandidateDelegator and assigns it to the Delegator field.
func (o *AssignmentCandidate) SetDelegator(v AssignmentCandidateDelegator) {
	o.Delegator = &v
}

func (o AssignmentCandidate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignmentCandidate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Delegator) {
		toSerialize["delegator"] = o.Delegator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignmentCandidate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
		"type",
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

	varAssignmentCandidate := _AssignmentCandidate{}

	err = json.Unmarshal(data, &varAssignmentCandidate)

	if err != nil {
		return err
	}

	*o = AssignmentCandidate(varAssignmentCandidate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "delegator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignmentCandidate struct {
	value *AssignmentCandidate
	isSet bool
}

func (v NullableAssignmentCandidate) Get() *AssignmentCandidate {
	return v.value
}

func (v *NullableAssignmentCandidate) Set(val *AssignmentCandidate) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentCandidate) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentCandidate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentCandidate(val *AssignmentCandidate) *NullableAssignmentCandidate {
	return &NullableAssignmentCandidate{value: val, isSet: true}
}

func (v NullableAssignmentCandidate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentCandidate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
