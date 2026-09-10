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

// checks if the EntitlementReconciliationConfigLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationConfigLinks{}

// EntitlementReconciliationConfigLinks Links to this configuration and its resource
type EntitlementReconciliationConfigLinks struct {
	Self                 Link  `json:"self"`
	Resource             *Link `json:"resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationConfigLinks EntitlementReconciliationConfigLinks

// NewEntitlementReconciliationConfigLinks instantiates a new EntitlementReconciliationConfigLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationConfigLinks(self Link) *EntitlementReconciliationConfigLinks {
	this := EntitlementReconciliationConfigLinks{}
	this.Self = self
	return &this
}

// NewEntitlementReconciliationConfigLinksWithDefaults instantiates a new EntitlementReconciliationConfigLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationConfigLinksWithDefaults() *EntitlementReconciliationConfigLinks {
	this := EntitlementReconciliationConfigLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *EntitlementReconciliationConfigLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *EntitlementReconciliationConfigLinks) SetSelf(v Link) {
	o.Self = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigLinks) GetResource() Link {
	if o == nil || IsNil(o.Resource) {
		var ret Link
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigLinks) GetResourceOk() (*Link, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigLinks) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given Link and assigns it to the Resource field.
func (o *EntitlementReconciliationConfigLinks) SetResource(v Link) {
	o.Resource = &v
}

func (o EntitlementReconciliationConfigLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationConfigLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationConfigLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varEntitlementReconciliationConfigLinks := _EntitlementReconciliationConfigLinks{}

	err = json.Unmarshal(data, &varEntitlementReconciliationConfigLinks)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationConfigLinks(varEntitlementReconciliationConfigLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "resource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationConfigLinks struct {
	value *EntitlementReconciliationConfigLinks
	isSet bool
}

func (v NullableEntitlementReconciliationConfigLinks) Get() *EntitlementReconciliationConfigLinks {
	return v.value
}

func (v *NullableEntitlementReconciliationConfigLinks) Set(val *EntitlementReconciliationConfigLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationConfigLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationConfigLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationConfigLinks(val *EntitlementReconciliationConfigLinks) *NullableEntitlementReconciliationConfigLinks {
	return &NullableEntitlementReconciliationConfigLinks{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationConfigLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationConfigLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
