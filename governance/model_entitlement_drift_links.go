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

// checks if the EntitlementDriftLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDriftLinks{}

// EntitlementDriftLinks Links to this drift and its related resources
type EntitlementDriftLinks struct {
	Self                 Link  `json:"self"`
	Revert               *Link `json:"revert,omitempty"`
	Resource             *Link `json:"resource,omitempty"`
	Principal            *Link `json:"principal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDriftLinks EntitlementDriftLinks

// NewEntitlementDriftLinks instantiates a new EntitlementDriftLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDriftLinks(self Link) *EntitlementDriftLinks {
	this := EntitlementDriftLinks{}
	this.Self = self
	return &this
}

// NewEntitlementDriftLinksWithDefaults instantiates a new EntitlementDriftLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDriftLinksWithDefaults() *EntitlementDriftLinks {
	this := EntitlementDriftLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *EntitlementDriftLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *EntitlementDriftLinks) SetSelf(v Link) {
	o.Self = v
}

// GetRevert returns the Revert field value if set, zero value otherwise.
func (o *EntitlementDriftLinks) GetRevert() Link {
	if o == nil || IsNil(o.Revert) {
		var ret Link
		return ret
	}
	return *o.Revert
}

// GetRevertOk returns a tuple with the Revert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftLinks) GetRevertOk() (*Link, bool) {
	if o == nil || IsNil(o.Revert) {
		return nil, false
	}
	return o.Revert, true
}

// HasRevert returns a boolean if a field has been set.
func (o *EntitlementDriftLinks) HasRevert() bool {
	if o != nil && !IsNil(o.Revert) {
		return true
	}

	return false
}

// SetRevert gets a reference to the given Link and assigns it to the Revert field.
func (o *EntitlementDriftLinks) SetRevert(v Link) {
	o.Revert = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *EntitlementDriftLinks) GetResource() Link {
	if o == nil || IsNil(o.Resource) {
		var ret Link
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftLinks) GetResourceOk() (*Link, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *EntitlementDriftLinks) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given Link and assigns it to the Resource field.
func (o *EntitlementDriftLinks) SetResource(v Link) {
	o.Resource = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *EntitlementDriftLinks) GetPrincipal() Link {
	if o == nil || IsNil(o.Principal) {
		var ret Link
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftLinks) GetPrincipalOk() (*Link, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *EntitlementDriftLinks) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given Link and assigns it to the Principal field.
func (o *EntitlementDriftLinks) SetPrincipal(v Link) {
	o.Principal = &v
}

func (o EntitlementDriftLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDriftLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Revert) {
		toSerialize["revert"] = o.Revert
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDriftLinks) UnmarshalJSON(data []byte) (err error) {
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

	varEntitlementDriftLinks := _EntitlementDriftLinks{}

	err = json.Unmarshal(data, &varEntitlementDriftLinks)

	if err != nil {
		return err
	}

	*o = EntitlementDriftLinks(varEntitlementDriftLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "revert")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "principal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDriftLinks struct {
	value *EntitlementDriftLinks
	isSet bool
}

func (v NullableEntitlementDriftLinks) Get() *EntitlementDriftLinks {
	return v.value
}

func (v *NullableEntitlementDriftLinks) Set(val *EntitlementDriftLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftLinks(val *EntitlementDriftLinks) *NullableEntitlementDriftLinks {
	return &NullableEntitlementDriftLinks{value: val, isSet: true}
}

func (v NullableEntitlementDriftLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
