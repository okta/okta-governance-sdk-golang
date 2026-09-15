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

// checks if the ResourceConnectionA2aServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceConnectionA2aServer{}

// ResourceConnectionA2aServer Agent-to-agent (A2A) server for the resource connection
type ResourceConnectionA2aServer struct {
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the A2A server
	Orn string `json:"orn"`
	// Display name of the agent-to-agent (A2A) server
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ResourceConnectionA2aServer ResourceConnectionA2aServer

// NewResourceConnectionA2aServer instantiates a new ResourceConnectionA2aServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceConnectionA2aServer(orn string, name string) *ResourceConnectionA2aServer {
	this := ResourceConnectionA2aServer{}
	this.Orn = orn
	this.Name = name
	return &this
}

// NewResourceConnectionA2aServerWithDefaults instantiates a new ResourceConnectionA2aServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceConnectionA2aServerWithDefaults() *ResourceConnectionA2aServer {
	this := ResourceConnectionA2aServer{}
	return &this
}

// GetOrn returns the Orn field value
func (o *ResourceConnectionA2aServer) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *ResourceConnectionA2aServer) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *ResourceConnectionA2aServer) SetOrn(v string) {
	o.Orn = v
}

// GetName returns the Name field value
func (o *ResourceConnectionA2aServer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceConnectionA2aServer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceConnectionA2aServer) SetName(v string) {
	o.Name = v
}

func (o ResourceConnectionA2aServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceConnectionA2aServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orn"] = o.Orn
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceConnectionA2aServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orn",
		"name",
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

	varResourceConnectionA2aServer := _ResourceConnectionA2aServer{}

	err = json.Unmarshal(data, &varResourceConnectionA2aServer)

	if err != nil {
		return err
	}

	*o = ResourceConnectionA2aServer(varResourceConnectionA2aServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceConnectionA2aServer struct {
	value *ResourceConnectionA2aServer
	isSet bool
}

func (v NullableResourceConnectionA2aServer) Get() *ResourceConnectionA2aServer {
	return v.value
}

func (v *NullableResourceConnectionA2aServer) Set(val *ResourceConnectionA2aServer) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceConnectionA2aServer) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceConnectionA2aServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceConnectionA2aServer(val *ResourceConnectionA2aServer) *NullableResourceConnectionA2aServer {
	return &NullableResourceConnectionA2aServer{value: val, isSet: true}
}

func (v NullableResourceConnectionA2aServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceConnectionA2aServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
