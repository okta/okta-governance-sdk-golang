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

// checks if the StsAccessTokenConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsAccessTokenConnection{}

// StsAccessTokenConnection STS connection for OAuth-based token exchange (app instances, API servers, or third-party MCP servers)
type StsAccessTokenConnection struct {
	// Type of connection authentication method
	ConnectionType string                 `json:"connectionType"`
	Resource       StsAccessTokenResource `json:"resource"`
	// Unique identifier for the resource connection
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the resource connection
	Orn *string `json:"orn,omitempty"`
	// The status of the connection
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StsAccessTokenConnection StsAccessTokenConnection

// NewStsAccessTokenConnection instantiates a new StsAccessTokenConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsAccessTokenConnection(connectionType string, resource StsAccessTokenResource) *StsAccessTokenConnection {
	this := StsAccessTokenConnection{}
	return &this
}

// NewStsAccessTokenConnectionWithDefaults instantiates a new StsAccessTokenConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsAccessTokenConnectionWithDefaults() *StsAccessTokenConnection {
	this := StsAccessTokenConnection{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *StsAccessTokenConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *StsAccessTokenConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *StsAccessTokenConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetResource returns the Resource field value
func (o *StsAccessTokenConnection) GetResource() StsAccessTokenResource {
	if o == nil {
		var ret StsAccessTokenResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *StsAccessTokenConnection) GetResourceOk() (*StsAccessTokenResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *StsAccessTokenConnection) SetResource(v StsAccessTokenResource) {
	o.Resource = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StsAccessTokenConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsAccessTokenConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StsAccessTokenConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StsAccessTokenConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *StsAccessTokenConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsAccessTokenConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *StsAccessTokenConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *StsAccessTokenConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StsAccessTokenConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsAccessTokenConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StsAccessTokenConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StsAccessTokenConnection) SetStatus(v string) {
	o.Status = &v
}

func (o StsAccessTokenConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsAccessTokenConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionType"] = o.ConnectionType
	toSerialize["resource"] = o.Resource
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StsAccessTokenConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionType",
		"resource",
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

	varStsAccessTokenConnection := _StsAccessTokenConnection{}

	err = json.Unmarshal(data, &varStsAccessTokenConnection)

	if err != nil {
		return err
	}

	*o = StsAccessTokenConnection(varStsAccessTokenConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connectionType")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStsAccessTokenConnection struct {
	value *StsAccessTokenConnection
	isSet bool
}

func (v NullableStsAccessTokenConnection) Get() *StsAccessTokenConnection {
	return v.value
}

func (v *NullableStsAccessTokenConnection) Set(val *StsAccessTokenConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableStsAccessTokenConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableStsAccessTokenConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsAccessTokenConnection(val *StsAccessTokenConnection) *NullableStsAccessTokenConnection {
	return &NullableStsAccessTokenConnection{value: val, isSet: true}
}

func (v NullableStsAccessTokenConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsAccessTokenConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
