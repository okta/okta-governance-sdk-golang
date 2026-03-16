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

// checks if the StsVaultSecretConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsVaultSecretConnection{}

// StsVaultSecretConnection STS connection to a vaulted secret
type StsVaultSecretConnection struct {
	// Type of connection authentication method
	ConnectionType string                         `json:"connectionType"`
	Secret         ManagedConnectionVaultedSecret `json:"secret"`
	// Unique identifier for the managed connection
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection
	Orn *string `json:"orn,omitempty"`
	// The status of the connection
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StsVaultSecretConnection StsVaultSecretConnection

// NewStsVaultSecretConnection instantiates a new StsVaultSecretConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsVaultSecretConnection(connectionType string, secret ManagedConnectionVaultedSecret) *StsVaultSecretConnection {
	this := StsVaultSecretConnection{}
	return &this
}

// NewStsVaultSecretConnectionWithDefaults instantiates a new StsVaultSecretConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsVaultSecretConnectionWithDefaults() *StsVaultSecretConnection {
	this := StsVaultSecretConnection{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *StsVaultSecretConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *StsVaultSecretConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *StsVaultSecretConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetSecret returns the Secret field value
func (o *StsVaultSecretConnection) GetSecret() ManagedConnectionVaultedSecret {
	if o == nil {
		var ret ManagedConnectionVaultedSecret
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *StsVaultSecretConnection) GetSecretOk() (*ManagedConnectionVaultedSecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *StsVaultSecretConnection) SetSecret(v ManagedConnectionVaultedSecret) {
	o.Secret = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StsVaultSecretConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsVaultSecretConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StsVaultSecretConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StsVaultSecretConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *StsVaultSecretConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsVaultSecretConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *StsVaultSecretConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *StsVaultSecretConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StsVaultSecretConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsVaultSecretConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StsVaultSecretConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StsVaultSecretConnection) SetStatus(v string) {
	o.Status = &v
}

func (o StsVaultSecretConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsVaultSecretConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionType"] = o.ConnectionType
	toSerialize["secret"] = o.Secret
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

func (o *StsVaultSecretConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionType",
		"secret",
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

	varStsVaultSecretConnection := _StsVaultSecretConnection{}

	err = json.Unmarshal(data, &varStsVaultSecretConnection)

	if err != nil {
		return err
	}

	*o = StsVaultSecretConnection(varStsVaultSecretConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connectionType")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStsVaultSecretConnection struct {
	value *StsVaultSecretConnection
	isSet bool
}

func (v NullableStsVaultSecretConnection) Get() *StsVaultSecretConnection {
	return v.value
}

func (v *NullableStsVaultSecretConnection) Set(val *StsVaultSecretConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableStsVaultSecretConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableStsVaultSecretConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsVaultSecretConnection(val *StsVaultSecretConnection) *NullableStsVaultSecretConnection {
	return &NullableStsVaultSecretConnection{value: val, isSet: true}
}

func (v NullableStsVaultSecretConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsVaultSecretConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
