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

// ResourceOwnersUpdatable The properties expected in create or update request for a resource owner.
type ResourceOwnersUpdatable struct {
	// Owners for the resource. If this is not provided, all current owners will be removed.
	PrincipalOrns []string `json:"principalOrns,omitempty"`
	// Resources that are being owned
	ResourceOrns         []string `json:"resourceOrns"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnersUpdatable ResourceOwnersUpdatable

// NewResourceOwnersUpdatable instantiates a new ResourceOwnersUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnersUpdatable(resourceOrns []string) *ResourceOwnersUpdatable {
	this := ResourceOwnersUpdatable{}
	this.ResourceOrns = resourceOrns
	return &this
}

// NewResourceOwnersUpdatableWithDefaults instantiates a new ResourceOwnersUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnersUpdatableWithDefaults() *ResourceOwnersUpdatable {
	this := ResourceOwnersUpdatable{}
	return &this
}

// GetPrincipalOrns returns the PrincipalOrns field value if set, zero value otherwise.
func (o *ResourceOwnersUpdatable) GetPrincipalOrns() []string {
	if o == nil || o.PrincipalOrns == nil {
		var ret []string
		return ret
	}
	return o.PrincipalOrns
}

// GetPrincipalOrnsOk returns a tuple with the PrincipalOrns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnersUpdatable) GetPrincipalOrnsOk() ([]string, bool) {
	if o == nil || o.PrincipalOrns == nil {
		return nil, false
	}
	return o.PrincipalOrns, true
}

// HasPrincipalOrns returns a boolean if a field has been set.
func (o *ResourceOwnersUpdatable) HasPrincipalOrns() bool {
	if o != nil && o.PrincipalOrns != nil {
		return true
	}

	return false
}

// SetPrincipalOrns gets a reference to the given []string and assigns it to the PrincipalOrns field.
func (o *ResourceOwnersUpdatable) SetPrincipalOrns(v []string) {
	o.PrincipalOrns = v
}

// GetResourceOrns returns the ResourceOrns field value
func (o *ResourceOwnersUpdatable) GetResourceOrns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceOrns
}

// GetResourceOrnsOk returns a tuple with the ResourceOrns field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnersUpdatable) GetResourceOrnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceOrns, true
}

// SetResourceOrns sets field value
func (o *ResourceOwnersUpdatable) SetResourceOrns(v []string) {
	o.ResourceOrns = v
}

func (o ResourceOwnersUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrincipalOrns != nil {
		toSerialize["principalOrns"] = o.PrincipalOrns
	}
	if true {
		toSerialize["resourceOrns"] = o.ResourceOrns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceOwnersUpdatable) UnmarshalJSON(bytes []byte) (err error) {
	varResourceOwnersUpdatable := _ResourceOwnersUpdatable{}

	err = json.Unmarshal(bytes, &varResourceOwnersUpdatable)
	if err == nil {
		*o = ResourceOwnersUpdatable(varResourceOwnersUpdatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "principalOrns")
		delete(additionalProperties, "resourceOrns")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceOwnersUpdatable struct {
	value *ResourceOwnersUpdatable
	isSet bool
}

func (v NullableResourceOwnersUpdatable) Get() *ResourceOwnersUpdatable {
	return v.value
}

func (v *NullableResourceOwnersUpdatable) Set(val *ResourceOwnersUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnersUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnersUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnersUpdatable(val *ResourceOwnersUpdatable) *NullableResourceOwnersUpdatable {
	return &NullableResourceOwnersUpdatable{value: val, isSet: true}
}

func (v NullableResourceOwnersUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnersUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
