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

// checks if the ReviewerAiAgentConnectedResourceSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResourceSecret{}

// ReviewerAiAgentConnectedResourceSecret OPA secret details for an AI agent connection
type ReviewerAiAgentConnectedResourceSecret struct {
	// Okta Resource Name (ORN) of the connected OPA secret
	Orn *string `json:"orn,omitempty"`
	// Name of the connected OPA secret
	Name *string `json:"name,omitempty"`
	// Description of the connected OPA secret
	Description *string `json:"description,omitempty"`
	// Path of the connected OPA secret
	Path                 *string `json:"path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResourceSecret ReviewerAiAgentConnectedResourceSecret

// NewReviewerAiAgentConnectedResourceSecret instantiates a new ReviewerAiAgentConnectedResourceSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResourceSecret() *ReviewerAiAgentConnectedResourceSecret {
	this := ReviewerAiAgentConnectedResourceSecret{}
	return &this
}

// NewReviewerAiAgentConnectedResourceSecretWithDefaults instantiates a new ReviewerAiAgentConnectedResourceSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceSecretWithDefaults() *ReviewerAiAgentConnectedResourceSecret {
	this := ReviewerAiAgentConnectedResourceSecret{}
	return &this
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceSecret) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ReviewerAiAgentConnectedResourceSecret) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceSecret) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewerAiAgentConnectedResourceSecret) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceSecret) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReviewerAiAgentConnectedResourceSecret) SetDescription(v string) {
	o.Description = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceSecret) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceSecret) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ReviewerAiAgentConnectedResourceSecret) SetPath(v string) {
	o.Path = &v
}

func (o ReviewerAiAgentConnectedResourceSecret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResourceSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerAiAgentConnectedResourceSecret) UnmarshalJSON(data []byte) (err error) {
	varReviewerAiAgentConnectedResourceSecret := _ReviewerAiAgentConnectedResourceSecret{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResourceSecret)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResourceSecret(varReviewerAiAgentConnectedResourceSecret)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResourceSecret struct {
	value *ReviewerAiAgentConnectedResourceSecret
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResourceSecret) Get() *ReviewerAiAgentConnectedResourceSecret {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResourceSecret) Set(val *ReviewerAiAgentConnectedResourceSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResourceSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResourceSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResourceSecret(val *ReviewerAiAgentConnectedResourceSecret) *NullableReviewerAiAgentConnectedResourceSecret {
	return &NullableReviewerAiAgentConnectedResourceSecret{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResourceSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResourceSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
