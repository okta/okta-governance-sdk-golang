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

// checks if the IdentityAssertionA2aServerConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAssertionA2aServerConnection{}

// IdentityAssertionA2aServerConnection Identity assertion connection for an agent-to-agent (A2A) server
type IdentityAssertionA2aServerConnection struct {
	// Type of connection authentication method
	ConnectionType      string                      `json:"connectionType"`
	A2aServer           ResourceConnectionA2aServer `json:"a2aServer"`
	AuthorizationServer CustomAuthorizationServer   `json:"authorizationServer"`
	// Unique identifier for the resource connection
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the resource connection
	Orn *string `json:"orn,omitempty"`
	// The status of the connection
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAssertionA2aServerConnection IdentityAssertionA2aServerConnection

// NewIdentityAssertionA2aServerConnection instantiates a new IdentityAssertionA2aServerConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAssertionA2aServerConnection(connectionType string, a2aServer ResourceConnectionA2aServer, authorizationServer CustomAuthorizationServer) *IdentityAssertionA2aServerConnection {
	this := IdentityAssertionA2aServerConnection{}
	return &this
}

// NewIdentityAssertionA2aServerConnectionWithDefaults instantiates a new IdentityAssertionA2aServerConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAssertionA2aServerConnectionWithDefaults() *IdentityAssertionA2aServerConnection {
	this := IdentityAssertionA2aServerConnection{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *IdentityAssertionA2aServerConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionA2aServerConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *IdentityAssertionA2aServerConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetA2aServer returns the A2aServer field value
func (o *IdentityAssertionA2aServerConnection) GetA2aServer() ResourceConnectionA2aServer {
	if o == nil {
		var ret ResourceConnectionA2aServer
		return ret
	}

	return o.A2aServer
}

// GetA2aServerOk returns a tuple with the A2aServer field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionA2aServerConnection) GetA2aServerOk() (*ResourceConnectionA2aServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.A2aServer, true
}

// SetA2aServer sets field value
func (o *IdentityAssertionA2aServerConnection) SetA2aServer(v ResourceConnectionA2aServer) {
	o.A2aServer = v
}

// GetAuthorizationServer returns the AuthorizationServer field value
func (o *IdentityAssertionA2aServerConnection) GetAuthorizationServer() CustomAuthorizationServer {
	if o == nil {
		var ret CustomAuthorizationServer
		return ret
	}

	return o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionA2aServerConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationServer, true
}

// SetAuthorizationServer sets field value
func (o *IdentityAssertionA2aServerConnection) SetAuthorizationServer(v CustomAuthorizationServer) {
	o.AuthorizationServer = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityAssertionA2aServerConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionA2aServerConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityAssertionA2aServerConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityAssertionA2aServerConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *IdentityAssertionA2aServerConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionA2aServerConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *IdentityAssertionA2aServerConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *IdentityAssertionA2aServerConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IdentityAssertionA2aServerConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionA2aServerConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentityAssertionA2aServerConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IdentityAssertionA2aServerConnection) SetStatus(v string) {
	o.Status = &v
}

func (o IdentityAssertionA2aServerConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAssertionA2aServerConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionType"] = o.ConnectionType
	toSerialize["a2aServer"] = o.A2aServer
	toSerialize["authorizationServer"] = o.AuthorizationServer
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

func (o *IdentityAssertionA2aServerConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionType",
		"a2aServer",
		"authorizationServer",
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

	varIdentityAssertionA2aServerConnection := _IdentityAssertionA2aServerConnection{}

	err = json.Unmarshal(data, &varIdentityAssertionA2aServerConnection)

	if err != nil {
		return err
	}

	*o = IdentityAssertionA2aServerConnection(varIdentityAssertionA2aServerConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connectionType")
		delete(additionalProperties, "a2aServer")
		delete(additionalProperties, "authorizationServer")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAssertionA2aServerConnection struct {
	value *IdentityAssertionA2aServerConnection
	isSet bool
}

func (v NullableIdentityAssertionA2aServerConnection) Get() *IdentityAssertionA2aServerConnection {
	return v.value
}

func (v *NullableIdentityAssertionA2aServerConnection) Set(val *IdentityAssertionA2aServerConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAssertionA2aServerConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAssertionA2aServerConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAssertionA2aServerConnection(val *IdentityAssertionA2aServerConnection) *NullableIdentityAssertionA2aServerConnection {
	return &NullableIdentityAssertionA2aServerConnection{value: val, isSet: true}
}

func (v NullableIdentityAssertionA2aServerConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAssertionA2aServerConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
