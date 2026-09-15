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

// checks if the ReviewerAiAgentConnectedResourceApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResourceApp{}

// ReviewerAiAgentConnectedResourceApp App details for an AI agent connection
type ReviewerAiAgentConnectedResourceApp struct {
	// ID of the connected app
	Id *string `json:"id,omitempty"`
	// Okta Resource Name (ORN) of the connected app
	Orn *string `json:"orn,omitempty"`
	// Name of the connected app
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResourceApp ReviewerAiAgentConnectedResourceApp

// NewReviewerAiAgentConnectedResourceApp instantiates a new ReviewerAiAgentConnectedResourceApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResourceApp() *ReviewerAiAgentConnectedResourceApp {
	this := ReviewerAiAgentConnectedResourceApp{}
	return &this
}

// NewReviewerAiAgentConnectedResourceAppWithDefaults instantiates a new ReviewerAiAgentConnectedResourceApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceAppWithDefaults() *ReviewerAiAgentConnectedResourceApp {
	this := ReviewerAiAgentConnectedResourceApp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReviewerAiAgentConnectedResourceApp) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApp) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApp) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApp) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ReviewerAiAgentConnectedResourceApp) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewerAiAgentConnectedResourceApp) SetName(v string) {
	o.Name = &v
}

func (o ReviewerAiAgentConnectedResourceApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResourceApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerAiAgentConnectedResourceApp) UnmarshalJSON(data []byte) (err error) {
	varReviewerAiAgentConnectedResourceApp := _ReviewerAiAgentConnectedResourceApp{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResourceApp)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResourceApp(varReviewerAiAgentConnectedResourceApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResourceApp struct {
	value *ReviewerAiAgentConnectedResourceApp
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResourceApp) Get() *ReviewerAiAgentConnectedResourceApp {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResourceApp) Set(val *ReviewerAiAgentConnectedResourceApp) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResourceApp) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResourceApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResourceApp(val *ReviewerAiAgentConnectedResourceApp) *NullableReviewerAiAgentConnectedResourceApp {
	return &NullableReviewerAiAgentConnectedResourceApp{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResourceApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResourceApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
