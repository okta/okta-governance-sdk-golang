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
)

// checks if the ExternalPrincipalProfileMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalPrincipalProfileMetadata{}

// ExternalPrincipalProfileMetadata Additional metadata about the resource owner, if applicable.
type ExternalPrincipalProfileMetadata struct {
	// Total number of users assigned to the group.
	UserAssignmentCount  *int32 `json:"userAssignmentCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExternalPrincipalProfileMetadata ExternalPrincipalProfileMetadata

// NewExternalPrincipalProfileMetadata instantiates a new ExternalPrincipalProfileMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPrincipalProfileMetadata() *ExternalPrincipalProfileMetadata {
	this := ExternalPrincipalProfileMetadata{}
	return &this
}

// NewExternalPrincipalProfileMetadataWithDefaults instantiates a new ExternalPrincipalProfileMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPrincipalProfileMetadataWithDefaults() *ExternalPrincipalProfileMetadata {
	this := ExternalPrincipalProfileMetadata{}
	return &this
}

// GetUserAssignmentCount returns the UserAssignmentCount field value if set, zero value otherwise.
func (o *ExternalPrincipalProfileMetadata) GetUserAssignmentCount() int32 {
	if o == nil || IsNil(o.UserAssignmentCount) {
		var ret int32
		return ret
	}
	return *o.UserAssignmentCount
}

// GetUserAssignmentCountOk returns a tuple with the UserAssignmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPrincipalProfileMetadata) GetUserAssignmentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UserAssignmentCount) {
		return nil, false
	}
	return o.UserAssignmentCount, true
}

// HasUserAssignmentCount returns a boolean if a field has been set.
func (o *ExternalPrincipalProfileMetadata) HasUserAssignmentCount() bool {
	if o != nil && !IsNil(o.UserAssignmentCount) {
		return true
	}

	return false
}

// SetUserAssignmentCount gets a reference to the given int32 and assigns it to the UserAssignmentCount field.
func (o *ExternalPrincipalProfileMetadata) SetUserAssignmentCount(v int32) {
	o.UserAssignmentCount = &v
}

func (o ExternalPrincipalProfileMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalPrincipalProfileMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserAssignmentCount) {
		toSerialize["userAssignmentCount"] = o.UserAssignmentCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalPrincipalProfileMetadata) UnmarshalJSON(data []byte) (err error) {
	varExternalPrincipalProfileMetadata := _ExternalPrincipalProfileMetadata{}

	err = json.Unmarshal(data, &varExternalPrincipalProfileMetadata)

	if err != nil {
		return err
	}

	*o = ExternalPrincipalProfileMetadata(varExternalPrincipalProfileMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userAssignmentCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalPrincipalProfileMetadata struct {
	value *ExternalPrincipalProfileMetadata
	isSet bool
}

func (v NullableExternalPrincipalProfileMetadata) Get() *ExternalPrincipalProfileMetadata {
	return v.value
}

func (v *NullableExternalPrincipalProfileMetadata) Set(val *ExternalPrincipalProfileMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPrincipalProfileMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPrincipalProfileMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPrincipalProfileMetadata(val *ExternalPrincipalProfileMetadata) *NullableExternalPrincipalProfileMetadata {
	return &NullableExternalPrincipalProfileMetadata{value: val, isSet: true}
}

func (v NullableExternalPrincipalProfileMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPrincipalProfileMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
