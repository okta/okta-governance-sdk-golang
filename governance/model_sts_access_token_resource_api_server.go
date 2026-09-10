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

// checks if the StsAccessTokenResourceApiServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsAccessTokenResourceApiServer{}

// StsAccessTokenResourceApiServer Resource data for an `STS_ACCESS_TOKEN` connection to a third-party API server
type StsAccessTokenResourceApiServer struct {
	// Type of resource for this `STS_ACCESS_TOKEN` connection
	ResourceType string `json:"resourceType"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the client auth settings
	Orn string `json:"orn"`
	// Display name of the third-party API server
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _StsAccessTokenResourceApiServer StsAccessTokenResourceApiServer

// NewStsAccessTokenResourceApiServer instantiates a new StsAccessTokenResourceApiServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsAccessTokenResourceApiServer(resourceType string, orn string, name string) *StsAccessTokenResourceApiServer {
	this := StsAccessTokenResourceApiServer{}
	this.ResourceType = resourceType
	this.Orn = orn
	this.Name = name
	return &this
}

// NewStsAccessTokenResourceApiServerWithDefaults instantiates a new StsAccessTokenResourceApiServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsAccessTokenResourceApiServerWithDefaults() *StsAccessTokenResourceApiServer {
	this := StsAccessTokenResourceApiServer{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *StsAccessTokenResourceApiServer) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *StsAccessTokenResourceApiServer) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *StsAccessTokenResourceApiServer) SetResourceType(v string) {
	o.ResourceType = v
}

// GetOrn returns the Orn field value
func (o *StsAccessTokenResourceApiServer) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *StsAccessTokenResourceApiServer) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *StsAccessTokenResourceApiServer) SetOrn(v string) {
	o.Orn = v
}

// GetName returns the Name field value
func (o *StsAccessTokenResourceApiServer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StsAccessTokenResourceApiServer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StsAccessTokenResourceApiServer) SetName(v string) {
	o.Name = v
}

func (o StsAccessTokenResourceApiServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsAccessTokenResourceApiServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["orn"] = o.Orn
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StsAccessTokenResourceApiServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceType",
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

	varStsAccessTokenResourceApiServer := _StsAccessTokenResourceApiServer{}

	err = json.Unmarshal(data, &varStsAccessTokenResourceApiServer)

	if err != nil {
		return err
	}

	*o = StsAccessTokenResourceApiServer(varStsAccessTokenResourceApiServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStsAccessTokenResourceApiServer struct {
	value *StsAccessTokenResourceApiServer
	isSet bool
}

func (v NullableStsAccessTokenResourceApiServer) Get() *StsAccessTokenResourceApiServer {
	return v.value
}

func (v *NullableStsAccessTokenResourceApiServer) Set(val *StsAccessTokenResourceApiServer) {
	v.value = val
	v.isSet = true
}

func (v NullableStsAccessTokenResourceApiServer) IsSet() bool {
	return v.isSet
}

func (v *NullableStsAccessTokenResourceApiServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsAccessTokenResourceApiServer(val *StsAccessTokenResourceApiServer) *NullableStsAccessTokenResourceApiServer {
	return &NullableStsAccessTokenResourceApiServer{value: val, isSet: true}
}

func (v NullableStsAccessTokenResourceApiServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsAccessTokenResourceApiServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
