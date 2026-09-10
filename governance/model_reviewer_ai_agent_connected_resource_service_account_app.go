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

// checks if the ReviewerAiAgentConnectedResourceServiceAccountApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResourceServiceAccountApp{}

// ReviewerAiAgentConnectedResourceServiceAccountApp App details of the connected service account
type ReviewerAiAgentConnectedResourceServiceAccountApp struct {
	// Okta Resource Name (ORN) of the app
	Orn *string `json:"orn,omitempty"`
	// Name of the app
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResourceServiceAccountApp ReviewerAiAgentConnectedResourceServiceAccountApp

// NewReviewerAiAgentConnectedResourceServiceAccountApp instantiates a new ReviewerAiAgentConnectedResourceServiceAccountApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResourceServiceAccountApp() *ReviewerAiAgentConnectedResourceServiceAccountApp {
	this := ReviewerAiAgentConnectedResourceServiceAccountApp{}
	return &this
}

// NewReviewerAiAgentConnectedResourceServiceAccountAppWithDefaults instantiates a new ReviewerAiAgentConnectedResourceServiceAccountApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceServiceAccountAppWithDefaults() *ReviewerAiAgentConnectedResourceServiceAccountApp {
	this := ReviewerAiAgentConnectedResourceServiceAccountApp{}
	return &this
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) SetName(v string) {
	o.Name = &v
}

func (o ReviewerAiAgentConnectedResourceServiceAccountApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResourceServiceAccountApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *ReviewerAiAgentConnectedResourceServiceAccountApp) UnmarshalJSON(data []byte) (err error) {
	varReviewerAiAgentConnectedResourceServiceAccountApp := _ReviewerAiAgentConnectedResourceServiceAccountApp{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResourceServiceAccountApp)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResourceServiceAccountApp(varReviewerAiAgentConnectedResourceServiceAccountApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResourceServiceAccountApp struct {
	value *ReviewerAiAgentConnectedResourceServiceAccountApp
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResourceServiceAccountApp) Get() *ReviewerAiAgentConnectedResourceServiceAccountApp {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResourceServiceAccountApp) Set(val *ReviewerAiAgentConnectedResourceServiceAccountApp) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResourceServiceAccountApp) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResourceServiceAccountApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResourceServiceAccountApp(val *ReviewerAiAgentConnectedResourceServiceAccountApp) *NullableReviewerAiAgentConnectedResourceServiceAccountApp {
	return &NullableReviewerAiAgentConnectedResourceServiceAccountApp{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResourceServiceAccountApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResourceServiceAccountApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
