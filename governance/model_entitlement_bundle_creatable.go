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

// EntitlementBundleCreatable The properties expected in an initial add entitlement bundle
type EntitlementBundleCreatable struct {
	Target TargetResource `json:"target"`
	// Collection of entitlements and associated value identifiers
	Entitlements []EntitlementCreatable `json:"entitlements"`
	// The unique name of the entitlement bundle
	Name string `json:"name"`
	// The human-readable description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementBundleCreatable EntitlementBundleCreatable

// NewEntitlementBundleCreatable instantiates a new EntitlementBundleCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementBundleCreatable(target TargetResource, entitlements []EntitlementCreatable, name string) *EntitlementBundleCreatable {
	this := EntitlementBundleCreatable{}
	this.Name = name
	return &this
}

// NewEntitlementBundleCreatableWithDefaults instantiates a new EntitlementBundleCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementBundleCreatableWithDefaults() *EntitlementBundleCreatable {
	this := EntitlementBundleCreatable{}
	return &this
}

// GetTarget returns the Target field value
func (o *EntitlementBundleCreatable) GetTarget() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleCreatable) GetTargetOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *EntitlementBundleCreatable) SetTarget(v TargetResource) {
	o.Target = v
}

// GetEntitlements returns the Entitlements field value
func (o *EntitlementBundleCreatable) GetEntitlements() []EntitlementCreatable {
	if o == nil {
		var ret []EntitlementCreatable
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleCreatable) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *EntitlementBundleCreatable) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetName returns the Name field value
func (o *EntitlementBundleCreatable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleCreatable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementBundleCreatable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementBundleCreatable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleCreatable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementBundleCreatable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementBundleCreatable) SetDescription(v string) {
	o.Description = &v
}

func (o EntitlementBundleCreatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["entitlements"] = o.Entitlements
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementBundleCreatable) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementBundleCreatable := _EntitlementBundleCreatable{}

	err = json.Unmarshal(bytes, &varEntitlementBundleCreatable)
	if err == nil {
		*o = EntitlementBundleCreatable(varEntitlementBundleCreatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "target")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementBundleCreatable struct {
	value *EntitlementBundleCreatable
	isSet bool
}

func (v NullableEntitlementBundleCreatable) Get() *EntitlementBundleCreatable {
	return v.value
}

func (v *NullableEntitlementBundleCreatable) Set(val *EntitlementBundleCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBundleCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBundleCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBundleCreatable(val *EntitlementBundleCreatable) *NullableEntitlementBundleCreatable {
	return &NullableEntitlementBundleCreatable{value: val, isSet: true}
}

func (v NullableEntitlementBundleCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBundleCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
