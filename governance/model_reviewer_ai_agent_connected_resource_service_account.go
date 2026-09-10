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

// checks if the ReviewerAiAgentConnectedResourceServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResourceServiceAccount{}

// ReviewerAiAgentConnectedResourceServiceAccount Service account details for an AI agent connection
type ReviewerAiAgentConnectedResourceServiceAccount struct {
	// Okta Resource Name (ORN) of the connected service account
	Orn *string `json:"orn,omitempty"`
	// Name of the connected service account
	Name                 *string                                            `json:"name,omitempty"`
	App                  *ReviewerAiAgentConnectedResourceServiceAccountApp `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResourceServiceAccount ReviewerAiAgentConnectedResourceServiceAccount

// NewReviewerAiAgentConnectedResourceServiceAccount instantiates a new ReviewerAiAgentConnectedResourceServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResourceServiceAccount() *ReviewerAiAgentConnectedResourceServiceAccount {
	this := ReviewerAiAgentConnectedResourceServiceAccount{}
	return &this
}

// NewReviewerAiAgentConnectedResourceServiceAccountWithDefaults instantiates a new ReviewerAiAgentConnectedResourceServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceServiceAccountWithDefaults() *ReviewerAiAgentConnectedResourceServiceAccount {
	this := ReviewerAiAgentConnectedResourceServiceAccount{}
	return &this
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) SetName(v string) {
	o.Name = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetApp() ReviewerAiAgentConnectedResourceServiceAccountApp {
	if o == nil || IsNil(o.App) {
		var ret ReviewerAiAgentConnectedResourceServiceAccountApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) GetAppOk() (*ReviewerAiAgentConnectedResourceServiceAccountApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given ReviewerAiAgentConnectedResourceServiceAccountApp and assigns it to the App field.
func (o *ReviewerAiAgentConnectedResourceServiceAccount) SetApp(v ReviewerAiAgentConnectedResourceServiceAccountApp) {
	o.App = &v
}

func (o ReviewerAiAgentConnectedResourceServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResourceServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerAiAgentConnectedResourceServiceAccount) UnmarshalJSON(data []byte) (err error) {
	varReviewerAiAgentConnectedResourceServiceAccount := _ReviewerAiAgentConnectedResourceServiceAccount{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResourceServiceAccount)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResourceServiceAccount(varReviewerAiAgentConnectedResourceServiceAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResourceServiceAccount struct {
	value *ReviewerAiAgentConnectedResourceServiceAccount
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResourceServiceAccount) Get() *ReviewerAiAgentConnectedResourceServiceAccount {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResourceServiceAccount) Set(val *ReviewerAiAgentConnectedResourceServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResourceServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResourceServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResourceServiceAccount(val *ReviewerAiAgentConnectedResourceServiceAccount) *NullableReviewerAiAgentConnectedResourceServiceAccount {
	return &NullableReviewerAiAgentConnectedResourceServiceAccount{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResourceServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResourceServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
