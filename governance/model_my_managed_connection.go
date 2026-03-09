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

// MyManagedConnection - struct for MyManagedConnection
type MyManagedConnection struct {
	IdentityAssertionCustomAsConnection *IdentityAssertionCustomAsConnection
	StsServiceAccountConnection         *StsServiceAccountConnection
	StsVaultSecretConnection            *StsVaultSecretConnection
}

// IdentityAssertionCustomAsConnectionAsMyManagedConnection is a convenience function that returns IdentityAssertionCustomAsConnection wrapped in MyManagedConnection
func IdentityAssertionCustomAsConnectionAsMyManagedConnection(v *IdentityAssertionCustomAsConnection) MyManagedConnection {
	return MyManagedConnection{
		IdentityAssertionCustomAsConnection: v,
	}
}

// StsServiceAccountConnectionAsMyManagedConnection is a convenience function that returns StsServiceAccountConnection wrapped in MyManagedConnection
func StsServiceAccountConnectionAsMyManagedConnection(v *StsServiceAccountConnection) MyManagedConnection {
	return MyManagedConnection{
		StsServiceAccountConnection: v,
	}
}

// StsVaultSecretConnectionAsMyManagedConnection is a convenience function that returns StsVaultSecretConnection wrapped in MyManagedConnection
func StsVaultSecretConnectionAsMyManagedConnection(v *StsVaultSecretConnection) MyManagedConnection {
	return MyManagedConnection{
		StsVaultSecretConnection: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MyManagedConnection) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'IDENTITY_ASSERTION_CUSTOM_AS'
	if jsonDict["connectionType"] == "IDENTITY_ASSERTION_CUSTOM_AS" {
		// try to unmarshal JSON data into IdentityAssertionCustomAsConnection
		err = json.Unmarshal(data, &dst.IdentityAssertionCustomAsConnection)
		if err == nil {
			return nil // data stored in dst.IdentityAssertionCustomAsConnection, return on the first match
		} else {
			dst.IdentityAssertionCustomAsConnection = nil
			return fmt.Errorf("failed to unmarshal MyManagedConnection as IdentityAssertionCustomAsConnection: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MyManagedConnection as StsServiceAccountConnection: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal MyManagedConnection as StsVaultSecretConnection: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MyManagedConnection) MarshalJSON() ([]byte, error) {
	if src.IdentityAssertionCustomAsConnection != nil {
		return json.Marshal(&src.IdentityAssertionCustomAsConnection)
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
func (obj *MyManagedConnection) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IdentityAssertionCustomAsConnection != nil {
		return obj.IdentityAssertionCustomAsConnection
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
func (obj MyManagedConnection) GetActualInstanceValue() interface{} {
	if obj.IdentityAssertionCustomAsConnection != nil {
		return *obj.IdentityAssertionCustomAsConnection
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

type NullableMyManagedConnection struct {
	value *MyManagedConnection
	isSet bool
}

func (v NullableMyManagedConnection) Get() *MyManagedConnection {
	return v.value
}

func (v *NullableMyManagedConnection) Set(val *MyManagedConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableMyManagedConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableMyManagedConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyManagedConnection(val *MyManagedConnection) *NullableMyManagedConnection {
	return &NullableMyManagedConnection{value: val, isSet: true}
}

func (v NullableMyManagedConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyManagedConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
