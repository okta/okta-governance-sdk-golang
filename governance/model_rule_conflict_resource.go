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

// RuleConflictResource struct for RuleConflictResource
type RuleConflictResource struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn          *string `json:"resourceOrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RuleConflictResource RuleConflictResource

// NewRuleConflictResource instantiates a new RuleConflictResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleConflictResource() *RuleConflictResource {
	this := RuleConflictResource{}
	return &this
}

// NewRuleConflictResourceWithDefaults instantiates a new RuleConflictResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConflictResourceWithDefaults() *RuleConflictResource {
	this := RuleConflictResource{}
	return &this
}

// GetResourceOrn returns the ResourceOrn field value if set, zero value otherwise.
func (o *RuleConflictResource) GetResourceOrn() string {
	if o == nil || o.ResourceOrn == nil {
		var ret string
		return ret
	}
	return *o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConflictResource) GetResourceOrnOk() (*string, bool) {
	if o == nil || o.ResourceOrn == nil {
		return nil, false
	}
	return o.ResourceOrn, true
}

// HasResourceOrn returns a boolean if a field has been set.
func (o *RuleConflictResource) HasResourceOrn() bool {
	if o != nil && o.ResourceOrn != nil {
		return true
	}

	return false
}

// SetResourceOrn gets a reference to the given string and assigns it to the ResourceOrn field.
func (o *RuleConflictResource) SetResourceOrn(v string) {
	o.ResourceOrn = &v
}

func (o RuleConflictResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceOrn != nil {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RuleConflictResource) UnmarshalJSON(bytes []byte) (err error) {
	varRuleConflictResource := _RuleConflictResource{}

	err = json.Unmarshal(bytes, &varRuleConflictResource)
	if err == nil {
		*o = RuleConflictResource(varRuleConflictResource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceOrn")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRuleConflictResource struct {
	value *RuleConflictResource
	isSet bool
}

func (v NullableRuleConflictResource) Get() *RuleConflictResource {
	return v.value
}

func (v *NullableRuleConflictResource) Set(val *RuleConflictResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleConflictResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleConflictResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleConflictResource(val *RuleConflictResource) *NullableRuleConflictResource {
	return &NullableRuleConflictResource{value: val, isSet: true}
}

func (v NullableRuleConflictResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleConflictResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
