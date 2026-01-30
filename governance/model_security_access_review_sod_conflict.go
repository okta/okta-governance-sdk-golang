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

// checks if the SecurityAccessReviewSodConflict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewSodConflict{}

// SecurityAccessReviewSodConflict struct for SecurityAccessReviewSodConflict
type SecurityAccessReviewSodConflict struct {
	// ID of SOD conflict rule
	Id string `json:"id"`
	// Name of SOD conflict rule
	Name string `json:"name"`
	// Description of SOD conflict rule
	Description *string                                 `json:"description,omitempty"`
	Operator    SecurityAccessReviewSodConflictOperator `json:"operator"`
	// List of conflicted entitlements that caused this SOD conflict
	ConflictedEntitlements []SecurityAccessReviewSodConflictConflictedEntitlementInfo `json:"conflictedEntitlements"`
	AdditionalProperties   map[string]interface{}
}

type _SecurityAccessReviewSodConflict SecurityAccessReviewSodConflict

// NewSecurityAccessReviewSodConflict instantiates a new SecurityAccessReviewSodConflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewSodConflict(id string, name string, operator SecurityAccessReviewSodConflictOperator, conflictedEntitlements []SecurityAccessReviewSodConflictConflictedEntitlementInfo) *SecurityAccessReviewSodConflict {
	this := SecurityAccessReviewSodConflict{}
	this.Id = id
	this.Name = name
	this.Operator = operator
	this.ConflictedEntitlements = conflictedEntitlements
	return &this
}

// NewSecurityAccessReviewSodConflictWithDefaults instantiates a new SecurityAccessReviewSodConflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewSodConflictWithDefaults() *SecurityAccessReviewSodConflict {
	this := SecurityAccessReviewSodConflict{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewSodConflict) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflict) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewSodConflict) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SecurityAccessReviewSodConflict) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflict) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReviewSodConflict) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityAccessReviewSodConflict) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflict) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityAccessReviewSodConflict) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityAccessReviewSodConflict) SetDescription(v string) {
	o.Description = &v
}

// GetOperator returns the Operator field value
func (o *SecurityAccessReviewSodConflict) GetOperator() SecurityAccessReviewSodConflictOperator {
	if o == nil {
		var ret SecurityAccessReviewSodConflictOperator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflict) GetOperatorOk() (*SecurityAccessReviewSodConflictOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SecurityAccessReviewSodConflict) SetOperator(v SecurityAccessReviewSodConflictOperator) {
	o.Operator = v
}

// GetConflictedEntitlements returns the ConflictedEntitlements field value
func (o *SecurityAccessReviewSodConflict) GetConflictedEntitlements() []SecurityAccessReviewSodConflictConflictedEntitlementInfo {
	if o == nil {
		var ret []SecurityAccessReviewSodConflictConflictedEntitlementInfo
		return ret
	}

	return o.ConflictedEntitlements
}

// GetConflictedEntitlementsOk returns a tuple with the ConflictedEntitlements field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSodConflict) GetConflictedEntitlementsOk() ([]SecurityAccessReviewSodConflictConflictedEntitlementInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConflictedEntitlements, true
}

// SetConflictedEntitlements sets field value
func (o *SecurityAccessReviewSodConflict) SetConflictedEntitlements(v []SecurityAccessReviewSodConflictConflictedEntitlementInfo) {
	o.ConflictedEntitlements = v
}

func (o SecurityAccessReviewSodConflict) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewSodConflict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["operator"] = o.Operator
	toSerialize["conflictedEntitlements"] = o.ConflictedEntitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewSodConflict) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"operator",
		"conflictedEntitlements",
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

	varSecurityAccessReviewSodConflict := _SecurityAccessReviewSodConflict{}

	err = json.Unmarshal(data, &varSecurityAccessReviewSodConflict)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewSodConflict(varSecurityAccessReviewSodConflict)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "conflictedEntitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewSodConflict struct {
	value *SecurityAccessReviewSodConflict
	isSet bool
}

func (v NullableSecurityAccessReviewSodConflict) Get() *SecurityAccessReviewSodConflict {
	return v.value
}

func (v *NullableSecurityAccessReviewSodConflict) Set(val *SecurityAccessReviewSodConflict) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSodConflict) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSodConflict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSodConflict(val *SecurityAccessReviewSodConflict) *NullableSecurityAccessReviewSodConflict {
	return &NullableSecurityAccessReviewSodConflict{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSodConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSodConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
