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

// checks if the EntitlementAccessDetailsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementAccessDetailsObject{}

// EntitlementAccessDetailsObject Entitlement access details
type EntitlementAccessDetailsObject struct {
	AssignmentObjectType AssignmentObjectType `json:"assignmentObjectType"`
	AssignmentObject     *AssignmentObject    `json:"assignmentObject,omitempty"`
	AssignmentMethod     AssignmentMethod     `json:"assignmentMethod"`
	// Collection of entitlements with associated values
	Entitlements         []EntitlementFull     `json:"entitlements"`
	AccessDuration       *AssignmentProperties `json:"accessDuration,omitempty"`
	Grant                *GrantObject          `json:"grant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementAccessDetailsObject EntitlementAccessDetailsObject

// NewEntitlementAccessDetailsObject instantiates a new EntitlementAccessDetailsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementAccessDetailsObject(assignmentObjectType AssignmentObjectType, assignmentMethod AssignmentMethod, entitlements []EntitlementFull) *EntitlementAccessDetailsObject {
	this := EntitlementAccessDetailsObject{}
	this.AssignmentObjectType = assignmentObjectType
	this.AssignmentMethod = assignmentMethod
	this.Entitlements = entitlements
	return &this
}

// NewEntitlementAccessDetailsObjectWithDefaults instantiates a new EntitlementAccessDetailsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementAccessDetailsObjectWithDefaults() *EntitlementAccessDetailsObject {
	this := EntitlementAccessDetailsObject{}
	return &this
}

// GetAssignmentObjectType returns the AssignmentObjectType field value
func (o *EntitlementAccessDetailsObject) GetAssignmentObjectType() AssignmentObjectType {
	if o == nil {
		var ret AssignmentObjectType
		return ret
	}

	return o.AssignmentObjectType
}

// GetAssignmentObjectTypeOk returns a tuple with the AssignmentObjectType field value
// and a boolean to check if the value has been set.
func (o *EntitlementAccessDetailsObject) GetAssignmentObjectTypeOk() (*AssignmentObjectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentObjectType, true
}

// SetAssignmentObjectType sets field value
func (o *EntitlementAccessDetailsObject) SetAssignmentObjectType(v AssignmentObjectType) {
	o.AssignmentObjectType = v
}

// GetAssignmentObject returns the AssignmentObject field value if set, zero value otherwise.
func (o *EntitlementAccessDetailsObject) GetAssignmentObject() AssignmentObject {
	if o == nil || IsNil(o.AssignmentObject) {
		var ret AssignmentObject
		return ret
	}
	return *o.AssignmentObject
}

// GetAssignmentObjectOk returns a tuple with the AssignmentObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAccessDetailsObject) GetAssignmentObjectOk() (*AssignmentObject, bool) {
	if o == nil || IsNil(o.AssignmentObject) {
		return nil, false
	}
	return o.AssignmentObject, true
}

// HasAssignmentObject returns a boolean if a field has been set.
func (o *EntitlementAccessDetailsObject) HasAssignmentObject() bool {
	if o != nil && !IsNil(o.AssignmentObject) {
		return true
	}

	return false
}

// SetAssignmentObject gets a reference to the given AssignmentObject and assigns it to the AssignmentObject field.
func (o *EntitlementAccessDetailsObject) SetAssignmentObject(v AssignmentObject) {
	o.AssignmentObject = &v
}

// GetAssignmentMethod returns the AssignmentMethod field value
func (o *EntitlementAccessDetailsObject) GetAssignmentMethod() AssignmentMethod {
	if o == nil {
		var ret AssignmentMethod
		return ret
	}

	return o.AssignmentMethod
}

// GetAssignmentMethodOk returns a tuple with the AssignmentMethod field value
// and a boolean to check if the value has been set.
func (o *EntitlementAccessDetailsObject) GetAssignmentMethodOk() (*AssignmentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentMethod, true
}

// SetAssignmentMethod sets field value
func (o *EntitlementAccessDetailsObject) SetAssignmentMethod(v AssignmentMethod) {
	o.AssignmentMethod = v
}

// GetEntitlements returns the Entitlements field value
func (o *EntitlementAccessDetailsObject) GetEntitlements() []EntitlementFull {
	if o == nil {
		var ret []EntitlementFull
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *EntitlementAccessDetailsObject) GetEntitlementsOk() ([]EntitlementFull, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *EntitlementAccessDetailsObject) SetEntitlements(v []EntitlementFull) {
	o.Entitlements = v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise.
func (o *EntitlementAccessDetailsObject) GetAccessDuration() AssignmentProperties {
	if o == nil || IsNil(o.AccessDuration) {
		var ret AssignmentProperties
		return ret
	}
	return *o.AccessDuration
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAccessDetailsObject) GetAccessDurationOk() (*AssignmentProperties, bool) {
	if o == nil || IsNil(o.AccessDuration) {
		return nil, false
	}
	return o.AccessDuration, true
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *EntitlementAccessDetailsObject) HasAccessDuration() bool {
	if o != nil && !IsNil(o.AccessDuration) {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given AssignmentProperties and assigns it to the AccessDuration field.
func (o *EntitlementAccessDetailsObject) SetAccessDuration(v AssignmentProperties) {
	o.AccessDuration = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *EntitlementAccessDetailsObject) GetGrant() GrantObject {
	if o == nil || IsNil(o.Grant) {
		var ret GrantObject
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAccessDetailsObject) GetGrantOk() (*GrantObject, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *EntitlementAccessDetailsObject) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given GrantObject and assigns it to the Grant field.
func (o *EntitlementAccessDetailsObject) SetGrant(v GrantObject) {
	o.Grant = &v
}

func (o EntitlementAccessDetailsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementAccessDetailsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignmentObjectType"] = o.AssignmentObjectType
	if !IsNil(o.AssignmentObject) {
		toSerialize["assignmentObject"] = o.AssignmentObject
	}
	toSerialize["assignmentMethod"] = o.AssignmentMethod
	toSerialize["entitlements"] = o.Entitlements
	if !IsNil(o.AccessDuration) {
		toSerialize["accessDuration"] = o.AccessDuration
	}
	if !IsNil(o.Grant) {
		toSerialize["grant"] = o.Grant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementAccessDetailsObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assignmentObjectType",
		"assignmentMethod",
		"entitlements",
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

	varEntitlementAccessDetailsObject := _EntitlementAccessDetailsObject{}

	err = json.Unmarshal(data, &varEntitlementAccessDetailsObject)

	if err != nil {
		return err
	}

	*o = EntitlementAccessDetailsObject(varEntitlementAccessDetailsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignmentObjectType")
		delete(additionalProperties, "assignmentObject")
		delete(additionalProperties, "assignmentMethod")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "accessDuration")
		delete(additionalProperties, "grant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementAccessDetailsObject struct {
	value *EntitlementAccessDetailsObject
	isSet bool
}

func (v NullableEntitlementAccessDetailsObject) Get() *EntitlementAccessDetailsObject {
	return v.value
}

func (v *NullableEntitlementAccessDetailsObject) Set(val *EntitlementAccessDetailsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementAccessDetailsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementAccessDetailsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementAccessDetailsObject(val *EntitlementAccessDetailsObject) *NullableEntitlementAccessDetailsObject {
	return &NullableEntitlementAccessDetailsObject{value: val, isSet: true}
}

func (v NullableEntitlementAccessDetailsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementAccessDetailsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
