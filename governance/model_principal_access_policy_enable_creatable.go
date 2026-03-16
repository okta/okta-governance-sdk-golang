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

// checks if the PrincipalAccessPolicyEnableCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalAccessPolicyEnableCreatable{}

// PrincipalAccessPolicyEnableCreatable struct for PrincipalAccessPolicyEnableCreatable
type PrincipalAccessPolicyEnableCreatable struct {
	// The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.
	PrincipalOrn string `json:"principalOrn"`
	// The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn string `json:"resourceOrn"`
	// If true, would skip removing bundle grant and create an entitlement grant with delta. By default it is false.
	MergeExistingEntitlements *bool `json:"mergeExistingEntitlements,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _PrincipalAccessPolicyEnableCreatable PrincipalAccessPolicyEnableCreatable

// NewPrincipalAccessPolicyEnableCreatable instantiates a new PrincipalAccessPolicyEnableCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalAccessPolicyEnableCreatable(principalOrn string, resourceOrn string) *PrincipalAccessPolicyEnableCreatable {
	this := PrincipalAccessPolicyEnableCreatable{}
	this.PrincipalOrn = principalOrn
	this.ResourceOrn = resourceOrn
	return &this
}

// NewPrincipalAccessPolicyEnableCreatableWithDefaults instantiates a new PrincipalAccessPolicyEnableCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalAccessPolicyEnableCreatableWithDefaults() *PrincipalAccessPolicyEnableCreatable {
	this := PrincipalAccessPolicyEnableCreatable{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *PrincipalAccessPolicyEnableCreatable) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableCreatable) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *PrincipalAccessPolicyEnableCreatable) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *PrincipalAccessPolicyEnableCreatable) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableCreatable) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *PrincipalAccessPolicyEnableCreatable) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetMergeExistingEntitlements returns the MergeExistingEntitlements field value if set, zero value otherwise.
func (o *PrincipalAccessPolicyEnableCreatable) GetMergeExistingEntitlements() bool {
	if o == nil || IsNil(o.MergeExistingEntitlements) {
		var ret bool
		return ret
	}
	return *o.MergeExistingEntitlements
}

// GetMergeExistingEntitlementsOk returns a tuple with the MergeExistingEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableCreatable) GetMergeExistingEntitlementsOk() (*bool, bool) {
	if o == nil || IsNil(o.MergeExistingEntitlements) {
		return nil, false
	}
	return o.MergeExistingEntitlements, true
}

// HasMergeExistingEntitlements returns a boolean if a field has been set.
func (o *PrincipalAccessPolicyEnableCreatable) HasMergeExistingEntitlements() bool {
	if o != nil && !IsNil(o.MergeExistingEntitlements) {
		return true
	}

	return false
}

// SetMergeExistingEntitlements gets a reference to the given bool and assigns it to the MergeExistingEntitlements field.
func (o *PrincipalAccessPolicyEnableCreatable) SetMergeExistingEntitlements(v bool) {
	o.MergeExistingEntitlements = &v
}

func (o PrincipalAccessPolicyEnableCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalAccessPolicyEnableCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalOrn"] = o.PrincipalOrn
	toSerialize["resourceOrn"] = o.ResourceOrn
	if !IsNil(o.MergeExistingEntitlements) {
		toSerialize["mergeExistingEntitlements"] = o.MergeExistingEntitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalAccessPolicyEnableCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalOrn",
		"resourceOrn",
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

	varPrincipalAccessPolicyEnableCreatable := _PrincipalAccessPolicyEnableCreatable{}

	err = json.Unmarshal(data, &varPrincipalAccessPolicyEnableCreatable)

	if err != nil {
		return err
	}

	*o = PrincipalAccessPolicyEnableCreatable(varPrincipalAccessPolicyEnableCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "mergeExistingEntitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalAccessPolicyEnableCreatable struct {
	value *PrincipalAccessPolicyEnableCreatable
	isSet bool
}

func (v NullablePrincipalAccessPolicyEnableCreatable) Get() *PrincipalAccessPolicyEnableCreatable {
	return v.value
}

func (v *NullablePrincipalAccessPolicyEnableCreatable) Set(val *PrincipalAccessPolicyEnableCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalAccessPolicyEnableCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalAccessPolicyEnableCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalAccessPolicyEnableCreatable(val *PrincipalAccessPolicyEnableCreatable) *NullablePrincipalAccessPolicyEnableCreatable {
	return &NullablePrincipalAccessPolicyEnableCreatable{value: val, isSet: true}
}

func (v NullablePrincipalAccessPolicyEnableCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalAccessPolicyEnableCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
