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

// StsAccessTokenResource - struct for StsAccessTokenResource
type StsAccessTokenResource struct {
	StsAccessTokenResourceApiServer   *StsAccessTokenResourceApiServer
	StsAccessTokenResourceAppInstance *StsAccessTokenResourceAppInstance
	StsAccessTokenResourceMcpServer   *StsAccessTokenResourceMcpServer
}

// StsAccessTokenResourceApiServerAsStsAccessTokenResource is a convenience function that returns StsAccessTokenResourceApiServer wrapped in StsAccessTokenResource
func StsAccessTokenResourceApiServerAsStsAccessTokenResource(v *StsAccessTokenResourceApiServer) StsAccessTokenResource {
	return StsAccessTokenResource{
		StsAccessTokenResourceApiServer: v,
	}
}

// StsAccessTokenResourceAppInstanceAsStsAccessTokenResource is a convenience function that returns StsAccessTokenResourceAppInstance wrapped in StsAccessTokenResource
func StsAccessTokenResourceAppInstanceAsStsAccessTokenResource(v *StsAccessTokenResourceAppInstance) StsAccessTokenResource {
	return StsAccessTokenResource{
		StsAccessTokenResourceAppInstance: v,
	}
}

// StsAccessTokenResourceMcpServerAsStsAccessTokenResource is a convenience function that returns StsAccessTokenResourceMcpServer wrapped in StsAccessTokenResource
func StsAccessTokenResourceMcpServerAsStsAccessTokenResource(v *StsAccessTokenResourceMcpServer) StsAccessTokenResource {
	return StsAccessTokenResource{
		StsAccessTokenResourceMcpServer: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StsAccessTokenResource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'API_SERVER'
	if jsonDict["resourceType"] == "API_SERVER" {
		// try to unmarshal JSON data into StsAccessTokenResourceApiServer
		err = json.Unmarshal(data, &dst.StsAccessTokenResourceApiServer)
		if err == nil {
			return nil // data stored in dst.StsAccessTokenResourceApiServer, return on the first match
		} else {
			dst.StsAccessTokenResourceApiServer = nil
			return fmt.Errorf("failed to unmarshal StsAccessTokenResource as StsAccessTokenResourceApiServer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'APP_INSTANCE'
	if jsonDict["resourceType"] == "APP_INSTANCE" {
		// try to unmarshal JSON data into StsAccessTokenResourceAppInstance
		err = json.Unmarshal(data, &dst.StsAccessTokenResourceAppInstance)
		if err == nil {
			return nil // data stored in dst.StsAccessTokenResourceAppInstance, return on the first match
		} else {
			dst.StsAccessTokenResourceAppInstance = nil
			return fmt.Errorf("failed to unmarshal StsAccessTokenResource as StsAccessTokenResourceAppInstance: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MCP_SERVER'
	if jsonDict["resourceType"] == "MCP_SERVER" {
		// try to unmarshal JSON data into StsAccessTokenResourceMcpServer
		err = json.Unmarshal(data, &dst.StsAccessTokenResourceMcpServer)
		if err == nil {
			return nil // data stored in dst.StsAccessTokenResourceMcpServer, return on the first match
		} else {
			dst.StsAccessTokenResourceMcpServer = nil
			return fmt.Errorf("failed to unmarshal StsAccessTokenResource as StsAccessTokenResourceMcpServer: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StsAccessTokenResource) MarshalJSON() ([]byte, error) {
	if src.StsAccessTokenResourceApiServer != nil {
		return json.Marshal(&src.StsAccessTokenResourceApiServer)
	}

	if src.StsAccessTokenResourceAppInstance != nil {
		return json.Marshal(&src.StsAccessTokenResourceAppInstance)
	}

	if src.StsAccessTokenResourceMcpServer != nil {
		return json.Marshal(&src.StsAccessTokenResourceMcpServer)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StsAccessTokenResource) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.StsAccessTokenResourceApiServer != nil {
		return obj.StsAccessTokenResourceApiServer
	}

	if obj.StsAccessTokenResourceAppInstance != nil {
		return obj.StsAccessTokenResourceAppInstance
	}

	if obj.StsAccessTokenResourceMcpServer != nil {
		return obj.StsAccessTokenResourceMcpServer
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj StsAccessTokenResource) GetActualInstanceValue() interface{} {
	if obj.StsAccessTokenResourceApiServer != nil {
		return *obj.StsAccessTokenResourceApiServer
	}

	if obj.StsAccessTokenResourceAppInstance != nil {
		return *obj.StsAccessTokenResourceAppInstance
	}

	if obj.StsAccessTokenResourceMcpServer != nil {
		return *obj.StsAccessTokenResourceMcpServer
	}

	// all schemas are nil
	return nil
}

type NullableStsAccessTokenResource struct {
	value *StsAccessTokenResource
	isSet bool
}

func (v NullableStsAccessTokenResource) Get() *StsAccessTokenResource {
	return v.value
}

func (v *NullableStsAccessTokenResource) Set(val *StsAccessTokenResource) {
	v.value = val
	v.isSet = true
}

func (v NullableStsAccessTokenResource) IsSet() bool {
	return v.isSet
}

func (v *NullableStsAccessTokenResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsAccessTokenResource(val *StsAccessTokenResource) *NullableStsAccessTokenResource {
	return &NullableStsAccessTokenResource{value: val, isSet: true}
}

func (v NullableStsAccessTokenResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsAccessTokenResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
