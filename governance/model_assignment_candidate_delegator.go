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

// checks if the AssignmentCandidateDelegator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignmentCandidateDelegator{}

// AssignmentCandidateDelegator The original principal, if this candidate is a delegate. Only present when `skipDelegateReplacement=false` in the request query.
type AssignmentCandidateDelegator struct {
	// The Okta user `id`
	ExternalId string        `json:"externalId" validate:"regexp=00u[0-9a-zA-Z]+"`
	Type       PrincipalType `json:"type"`
	// The delegator `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.
	Orn                  *string `json:"orn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignmentCandidateDelegator AssignmentCandidateDelegator

// NewAssignmentCandidateDelegator instantiates a new AssignmentCandidateDelegator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignmentCandidateDelegator(externalId string, type_ PrincipalType) *AssignmentCandidateDelegator {
	this := AssignmentCandidateDelegator{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewAssignmentCandidateDelegatorWithDefaults instantiates a new AssignmentCandidateDelegator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignmentCandidateDelegatorWithDefaults() *AssignmentCandidateDelegator {
	this := AssignmentCandidateDelegator{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *AssignmentCandidateDelegator) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AssignmentCandidateDelegator) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *AssignmentCandidateDelegator) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *AssignmentCandidateDelegator) GetType() PrincipalType {
	if o == nil {
		var ret PrincipalType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssignmentCandidateDelegator) GetTypeOk() (*PrincipalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssignmentCandidateDelegator) SetType(v PrincipalType) {
	o.Type = v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *AssignmentCandidateDelegator) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignmentCandidateDelegator) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *AssignmentCandidateDelegator) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *AssignmentCandidateDelegator) SetOrn(v string) {
	o.Orn = &v
}

func (o AssignmentCandidateDelegator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignmentCandidateDelegator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignmentCandidateDelegator) UnmarshalJSON(data []byte) (err error) {
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

	varAssignmentCandidateDelegator := _AssignmentCandidateDelegator{}

	err = json.Unmarshal(data, &varAssignmentCandidateDelegator)

	if err != nil {
		return err
	}

	*o = AssignmentCandidateDelegator(varAssignmentCandidateDelegator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "orn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignmentCandidateDelegator struct {
	value *AssignmentCandidateDelegator
	isSet bool
}

func (v NullableAssignmentCandidateDelegator) Get() *AssignmentCandidateDelegator {
	return v.value
}

func (v *NullableAssignmentCandidateDelegator) Set(val *AssignmentCandidateDelegator) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignmentCandidateDelegator) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignmentCandidateDelegator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignmentCandidateDelegator(val *AssignmentCandidateDelegator) *NullableAssignmentCandidateDelegator {
	return &NullableAssignmentCandidateDelegator{value: val, isSet: true}
}

func (v NullableAssignmentCandidateDelegator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignmentCandidateDelegator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
