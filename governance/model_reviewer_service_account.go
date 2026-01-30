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

// checks if the ReviewerServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerServiceAccount{}

// ReviewerServiceAccount Service account information for reviewer
type ReviewerServiceAccount struct {
	// The service account ID
	Id string `json:"id"`
	// The service account name
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerServiceAccount ReviewerServiceAccount

// NewReviewerServiceAccount instantiates a new ReviewerServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerServiceAccount(id string, name string) *ReviewerServiceAccount {
	this := ReviewerServiceAccount{}
	this.Id = id
	this.Name = name
	return &this
}

// NewReviewerServiceAccountWithDefaults instantiates a new ReviewerServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerServiceAccountWithDefaults() *ReviewerServiceAccount {
	this := ReviewerServiceAccount{}
	return &this
}

// GetId returns the Id field value
func (o *ReviewerServiceAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewerServiceAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewerServiceAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ReviewerServiceAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReviewerServiceAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReviewerServiceAccount) SetName(v string) {
	o.Name = v
}

func (o ReviewerServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerServiceAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varReviewerServiceAccount := _ReviewerServiceAccount{}

	err = json.Unmarshal(data, &varReviewerServiceAccount)

	if err != nil {
		return err
	}

	*o = ReviewerServiceAccount(varReviewerServiceAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerServiceAccount struct {
	value *ReviewerServiceAccount
	isSet bool
}

func (v NullableReviewerServiceAccount) Get() *ReviewerServiceAccount {
	return v.value
}

func (v *NullableReviewerServiceAccount) Set(val *ReviewerServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerServiceAccount(val *ReviewerServiceAccount) *NullableReviewerServiceAccount {
	return &NullableReviewerServiceAccount{value: val, isSet: true}
}

func (v NullableReviewerServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
