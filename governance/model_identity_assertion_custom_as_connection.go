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

// checks if the IdentityAssertionCustomAsConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAssertionCustomAsConnection{}

// IdentityAssertionCustomAsConnection Identity assertion connection for a custom authorization server
type IdentityAssertionCustomAsConnection struct {
	// Type of connection authentication method
	ConnectionType      string                    `json:"connectionType"`
	AuthorizationServer CustomAuthorizationServer `json:"authorizationServer"`
	// Unique identifier for the managed connection
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection
	Orn *string `json:"orn,omitempty"`
	// The status of the connection
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityAssertionCustomAsConnection IdentityAssertionCustomAsConnection

// NewIdentityAssertionCustomAsConnection instantiates a new IdentityAssertionCustomAsConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAssertionCustomAsConnection(connectionType string, authorizationServer CustomAuthorizationServer) *IdentityAssertionCustomAsConnection {
	this := IdentityAssertionCustomAsConnection{}
	return &this
}

// NewIdentityAssertionCustomAsConnectionWithDefaults instantiates a new IdentityAssertionCustomAsConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAssertionCustomAsConnectionWithDefaults() *IdentityAssertionCustomAsConnection {
	this := IdentityAssertionCustomAsConnection{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *IdentityAssertionCustomAsConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomAsConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *IdentityAssertionCustomAsConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetAuthorizationServer returns the AuthorizationServer field value
func (o *IdentityAssertionCustomAsConnection) GetAuthorizationServer() CustomAuthorizationServer {
	if o == nil {
		var ret CustomAuthorizationServer
		return ret
	}

	return o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomAsConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationServer, true
}

// SetAuthorizationServer sets field value
func (o *IdentityAssertionCustomAsConnection) SetAuthorizationServer(v CustomAuthorizationServer) {
	o.AuthorizationServer = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityAssertionCustomAsConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomAsConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityAssertionCustomAsConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityAssertionCustomAsConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *IdentityAssertionCustomAsConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomAsConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *IdentityAssertionCustomAsConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *IdentityAssertionCustomAsConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IdentityAssertionCustomAsConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomAsConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentityAssertionCustomAsConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IdentityAssertionCustomAsConnection) SetStatus(v string) {
	o.Status = &v
}

func (o IdentityAssertionCustomAsConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAssertionCustomAsConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionType"] = o.ConnectionType
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

func (o *IdentityAssertionCustomAsConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionType",
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

	varIdentityAssertionCustomAsConnection := _IdentityAssertionCustomAsConnection{}

	err = json.Unmarshal(data, &varIdentityAssertionCustomAsConnection)

	if err != nil {
		return err
	}

	*o = IdentityAssertionCustomAsConnection(varIdentityAssertionCustomAsConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connectionType")
		delete(additionalProperties, "authorizationServer")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityAssertionCustomAsConnection struct {
	value *IdentityAssertionCustomAsConnection
	isSet bool
}

func (v NullableIdentityAssertionCustomAsConnection) Get() *IdentityAssertionCustomAsConnection {
	return v.value
}

func (v *NullableIdentityAssertionCustomAsConnection) Set(val *IdentityAssertionCustomAsConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAssertionCustomAsConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAssertionCustomAsConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAssertionCustomAsConnection(val *IdentityAssertionCustomAsConnection) *NullableIdentityAssertionCustomAsConnection {
	return &NullableIdentityAssertionCustomAsConnection{value: val, isSet: true}
}

func (v NullableIdentityAssertionCustomAsConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAssertionCustomAsConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
