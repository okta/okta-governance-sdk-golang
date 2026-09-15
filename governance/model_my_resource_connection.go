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

// MyResourceConnection - struct for MyResourceConnection
type MyResourceConnection struct {
	IdentityAssertionA2aServerConnection *IdentityAssertionA2aServerConnection
	IdentityAssertionCustomAsConnection  *IdentityAssertionCustomAsConnection
	StsAccessTokenConnection             *StsAccessTokenConnection
	StsServiceAccountConnection          *StsServiceAccountConnection
	StsVaultSecretConnection             *StsVaultSecretConnection
}

// IdentityAssertionA2aServerConnectionAsMyResourceConnection is a convenience function that returns IdentityAssertionA2aServerConnection wrapped in MyResourceConnection
func IdentityAssertionA2aServerConnectionAsMyResourceConnection(v *IdentityAssertionA2aServerConnection) MyResourceConnection {
	return MyResourceConnection{
		IdentityAssertionA2aServerConnection: v,
	}
}

// IdentityAssertionCustomAsConnectionAsMyResourceConnection is a convenience function that returns IdentityAssertionCustomAsConnection wrapped in MyResourceConnection
func IdentityAssertionCustomAsConnectionAsMyResourceConnection(v *IdentityAssertionCustomAsConnection) MyResourceConnection {
	return MyResourceConnection{
		IdentityAssertionCustomAsConnection: v,
	}
}

// StsAccessTokenConnectionAsMyResourceConnection is a convenience function that returns StsAccessTokenConnection wrapped in MyResourceConnection
func StsAccessTokenConnectionAsMyResourceConnection(v *StsAccessTokenConnection) MyResourceConnection {
	return MyResourceConnection{
		StsAccessTokenConnection: v,
	}
}

// StsServiceAccountConnectionAsMyResourceConnection is a convenience function that returns StsServiceAccountConnection wrapped in MyResourceConnection
func StsServiceAccountConnectionAsMyResourceConnection(v *StsServiceAccountConnection) MyResourceConnection {
	return MyResourceConnection{
		StsServiceAccountConnection: v,
	}
}

// StsVaultSecretConnectionAsMyResourceConnection is a convenience function that returns StsVaultSecretConnection wrapped in MyResourceConnection
func StsVaultSecretConnectionAsMyResourceConnection(v *StsVaultSecretConnection) MyResourceConnection {
	return MyResourceConnection{
		StsVaultSecretConnection: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MyResourceConnection) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'IDENTITY_ASSERTION_A2A_SERVER'
	if jsonDict["connectionType"] == "IDENTITY_ASSERTION_A2A_SERVER" {
		// try to unmarshal JSON data into IdentityAssertionA2aServerConnection
		err = json.Unmarshal(data, &dst.IdentityAssertionA2aServerConnection)
		if err == nil {
			return nil // data stored in dst.IdentityAssertionA2aServerConnection, return on the first match
		} else {
			dst.IdentityAssertionA2aServerConnection = nil
			return fmt.Errorf("failed to unmarshal MyResourceConnection as IdentityAssertionA2aServerConnection: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IDENTITY_ASSERTION_CUSTOM_AS'
	if jsonDict["connectionType"] == "IDENTITY_ASSERTION_CUSTOM_AS" {
		// try to unmarshal JSON data into IdentityAssertionCustomAsConnection
		err = json.Unmarshal(data, &dst.IdentityAssertionCustomAsConnection)
		if err == nil {
			return nil // data stored in dst.IdentityAssertionCustomAsConnection, return on the first match
		} else {
			dst.IdentityAssertionCustomAsConnection = nil
			return fmt.Errorf("failed to unmarshal MyResourceConnection as IdentityAssertionCustomAsConnection: %s", err.Error())
		}
	}

	// check if the discriminator value is 'STS_ACCESS_TOKEN'
	if jsonDict["connectionType"] == "STS_ACCESS_TOKEN" {
		// try to unmarshal JSON data into StsAccessTokenConnection
		err = json.Unmarshal(data, &dst.StsAccessTokenConnection)
		if err == nil {
			return nil // data stored in dst.StsAccessTokenConnection, return on the first match
		} else {
			dst.StsAccessTokenConnection = nil
			return fmt.Errorf("failed to unmarshal MyResourceConnection as StsAccessTokenConnection: %s", err.Error())
		}
	}

	// check if the discriminator value is 'STS_SERVICE_ACCOUNT'
	if jsonDict["connectionType"] == "STS_SERVICE_ACCOUNT" {
		// try to unmarshal JSON data into StsServiceAccountConnection
		err = json.Unmarshal(data, &dst.StsServiceAccountConnection)
		if err == nil {
			return nil // data stored in dst.StsServiceAccountConnection, return on the first match
		} else {
			dst.StsServiceAccountConnection = nil
			return fmt.Errorf("failed to unmarshal MyResourceConnection as StsServiceAccountConnection: %s", err.Error())
		}
	}

	// check if the discriminator value is 'STS_VAULT_SECRET'
	if jsonDict["connectionType"] == "STS_VAULT_SECRET" {
		// try to unmarshal JSON data into StsVaultSecretConnection
		err = json.Unmarshal(data, &dst.StsVaultSecretConnection)
		if err == nil {
			return nil // data stored in dst.StsVaultSecretConnection, return on the first match
		} else {
			dst.StsVaultSecretConnection = nil
			return fmt.Errorf("failed to unmarshal MyResourceConnection as StsVaultSecretConnection: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MyResourceConnection) MarshalJSON() ([]byte, error) {
	if src.IdentityAssertionA2aServerConnection != nil {
		return json.Marshal(&src.IdentityAssertionA2aServerConnection)
	}

	if src.IdentityAssertionCustomAsConnection != nil {
		return json.Marshal(&src.IdentityAssertionCustomAsConnection)
	}

	if src.StsAccessTokenConnection != nil {
		return json.Marshal(&src.StsAccessTokenConnection)
	}

	if src.StsServiceAccountConnection != nil {
		return json.Marshal(&src.StsServiceAccountConnection)
	}

	if src.StsVaultSecretConnection != nil {
		return json.Marshal(&src.StsVaultSecretConnection)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MyResourceConnection) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IdentityAssertionA2aServerConnection != nil {
		return obj.IdentityAssertionA2aServerConnection
	}

	if obj.IdentityAssertionCustomAsConnection != nil {
		return obj.IdentityAssertionCustomAsConnection
	}

	if obj.StsAccessTokenConnection != nil {
		return obj.StsAccessTokenConnection
	}

	if obj.StsServiceAccountConnection != nil {
		return obj.StsServiceAccountConnection
	}

	if obj.StsVaultSecretConnection != nil {
		return obj.StsVaultSecretConnection
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj MyResourceConnection) GetActualInstanceValue() interface{} {
	if obj.IdentityAssertionA2aServerConnection != nil {
		return *obj.IdentityAssertionA2aServerConnection
	}

	if obj.IdentityAssertionCustomAsConnection != nil {
		return *obj.IdentityAssertionCustomAsConnection
	}

	if obj.StsAccessTokenConnection != nil {
		return *obj.StsAccessTokenConnection
	}

	if obj.StsServiceAccountConnection != nil {
		return *obj.StsServiceAccountConnection
	}

	if obj.StsVaultSecretConnection != nil {
		return *obj.StsVaultSecretConnection
	}

	// all schemas are nil
	return nil
}

type NullableMyResourceConnection struct {
	value *MyResourceConnection
	isSet bool
}

func (v NullableMyResourceConnection) Get() *MyResourceConnection {
	return v.value
}

func (v *NullableMyResourceConnection) Set(val *MyResourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableMyResourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableMyResourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyResourceConnection(val *MyResourceConnection) *NullableMyResourceConnection {
	return &NullableMyResourceConnection{value: val, isSet: true}
}

func (v NullableMyResourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyResourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
