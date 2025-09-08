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
)

// RiskRule struct for RiskRule
type RiskRule struct {
	// The name of a resource rule causing a conflict
	Name string `json:"name"`
	// The human readable description
	Description *string `json:"description,omitempty"`
	// Human readable name of the resource
	ResourceName         *string `json:"resourceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskRule RiskRule

// NewRiskRule instantiates a new RiskRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRule(name string) *RiskRule {
	this := RiskRule{}
	this.Name = name
	return &this
}

// NewRiskRuleWithDefaults instantiates a new RiskRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRuleWithDefaults() *RiskRule {
	this := RiskRule{}
	return &this
}

// GetName returns the Name field value
func (o *RiskRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskRule) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskRule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskRule) SetDescription(v string) {
	o.Description = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *RiskRule) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRule) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *RiskRule) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *RiskRule) SetResourceName(v string) {
	o.ResourceName = &v
}

func (o RiskRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ResourceName != nil {
		toSerialize["resourceName"] = o.ResourceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskRule) UnmarshalJSON(bytes []byte) (err error) {
	varRiskRule := _RiskRule{}

	err = json.Unmarshal(bytes, &varRiskRule)
	if err == nil {
		*o = RiskRule(varRiskRule)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "resourceName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskRule struct {
	value *RiskRule
	isSet bool
}

func (v NullableRiskRule) Get() *RiskRule {
	return v.value
}

func (v *NullableRiskRule) Set(val *RiskRule) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRule(val *RiskRule) *NullableRiskRule {
	return &NullableRiskRule{value: val, isSet: true}
}

func (v NullableRiskRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
