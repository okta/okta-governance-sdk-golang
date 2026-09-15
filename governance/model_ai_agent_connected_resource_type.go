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

// AiAgentConnectedResourceType Type of resource connected to an AI agent
type AiAgentConnectedResourceType string

// List of ai-agent-connected-resource-type
const (
	AIAGENTCONNECTEDRESOURCETYPE_APP                  AiAgentConnectedResourceType = "APP"
	AIAGENTCONNECTEDRESOURCETYPE_MCP_SERVER           AiAgentConnectedResourceType = "MCP_SERVER"
	AIAGENTCONNECTEDRESOURCETYPE_API_SERVER           AiAgentConnectedResourceType = "API_SERVER"
	AIAGENTCONNECTEDRESOURCETYPE_AUTHORIZATION_SERVER AiAgentConnectedResourceType = "AUTHORIZATION_SERVER"
	AIAGENTCONNECTEDRESOURCETYPE_OPA_SECRET           AiAgentConnectedResourceType = "OPA_SECRET"
	AIAGENTCONNECTEDRESOURCETYPE_APP_SERVICE_ACCOUNT  AiAgentConnectedResourceType = "APP_SERVICE_ACCOUNT"
)

// All allowed values of AiAgentConnectedResourceType enum
var AllowedAiAgentConnectedResourceTypeEnumValues = []AiAgentConnectedResourceType{
	"APP",
	"MCP_SERVER",
	"API_SERVER",
	"AUTHORIZATION_SERVER",
	"OPA_SECRET",
	"APP_SERVICE_ACCOUNT",
}

func (v *AiAgentConnectedResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiAgentConnectedResourceType(value)
	for _, existing := range AllowedAiAgentConnectedResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiAgentConnectedResourceType", value)
}

// NewAiAgentConnectedResourceTypeFromValue returns a pointer to a valid AiAgentConnectedResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiAgentConnectedResourceTypeFromValue(v string) (*AiAgentConnectedResourceType, error) {
	ev := AiAgentConnectedResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiAgentConnectedResourceType: valid values are %v", v, AllowedAiAgentConnectedResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiAgentConnectedResourceType) IsValid() bool {
	for _, existing := range AllowedAiAgentConnectedResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ai-agent-connected-resource-type value
func (v AiAgentConnectedResourceType) Ptr() *AiAgentConnectedResourceType {
	return &v
}

type NullableAiAgentConnectedResourceType struct {
	value *AiAgentConnectedResourceType
	isSet bool
}

func (v NullableAiAgentConnectedResourceType) Get() *AiAgentConnectedResourceType {
	return v.value
}

func (v *NullableAiAgentConnectedResourceType) Set(val *AiAgentConnectedResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAgentConnectedResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAgentConnectedResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAgentConnectedResourceType(val *AiAgentConnectedResourceType) *NullableAiAgentConnectedResourceType {
	return &NullableAiAgentConnectedResourceType{value: val, isSet: true}
}

func (v NullableAiAgentConnectedResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAgentConnectedResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
