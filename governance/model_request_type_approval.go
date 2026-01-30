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

// RequestTypeApproval - request-type-approval-description.md
type RequestTypeApproval struct {
	RequestTypeApprovalManager       *RequestTypeApprovalManager
	RequestTypeApprovalMemberOf      *RequestTypeApprovalMemberOf
	RequestTypeApprovalResourceOwner *RequestTypeApprovalResourceOwner
	RequestTypeApprovalUser          *RequestTypeApprovalUser
}

// RequestTypeApprovalManagerAsRequestTypeApproval is a convenience function that returns RequestTypeApprovalManager wrapped in RequestTypeApproval
func RequestTypeApprovalManagerAsRequestTypeApproval(v *RequestTypeApprovalManager) RequestTypeApproval {
	return RequestTypeApproval{
		RequestTypeApprovalManager: v,
	}
}

// RequestTypeApprovalMemberOfAsRequestTypeApproval is a convenience function that returns RequestTypeApprovalMemberOf wrapped in RequestTypeApproval
func RequestTypeApprovalMemberOfAsRequestTypeApproval(v *RequestTypeApprovalMemberOf) RequestTypeApproval {
	return RequestTypeApproval{
		RequestTypeApprovalMemberOf: v,
	}
}

// RequestTypeApprovalResourceOwnerAsRequestTypeApproval is a convenience function that returns RequestTypeApprovalResourceOwner wrapped in RequestTypeApproval
func RequestTypeApprovalResourceOwnerAsRequestTypeApproval(v *RequestTypeApprovalResourceOwner) RequestTypeApproval {
	return RequestTypeApproval{
		RequestTypeApprovalResourceOwner: v,
	}
}

// RequestTypeApprovalUserAsRequestTypeApproval is a convenience function that returns RequestTypeApprovalUser wrapped in RequestTypeApproval
func RequestTypeApprovalUserAsRequestTypeApproval(v *RequestTypeApprovalUser) RequestTypeApproval {
	return RequestTypeApproval{
		RequestTypeApprovalUser: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestTypeApproval) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'MANAGER'
	if jsonDict["approverType"] == "MANAGER" {
		// try to unmarshal JSON data into RequestTypeApprovalManager
		err = json.Unmarshal(data, &dst.RequestTypeApprovalManager)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalManager, return on the first match
		} else {
			dst.RequestTypeApprovalManager = nil
			return fmt.Errorf("failed to unmarshal RequestTypeApproval as RequestTypeApprovalManager: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MEMBER_OF'
	if jsonDict["approverType"] == "MEMBER_OF" {
		// try to unmarshal JSON data into RequestTypeApprovalMemberOf
		err = json.Unmarshal(data, &dst.RequestTypeApprovalMemberOf)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalMemberOf, return on the first match
		} else {
			dst.RequestTypeApprovalMemberOf = nil
			return fmt.Errorf("failed to unmarshal RequestTypeApproval as RequestTypeApprovalMemberOf: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESOURCE_OWNER'
	if jsonDict["approverType"] == "RESOURCE_OWNER" {
		// try to unmarshal JSON data into RequestTypeApprovalResourceOwner
		err = json.Unmarshal(data, &dst.RequestTypeApprovalResourceOwner)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalResourceOwner, return on the first match
		} else {
			dst.RequestTypeApprovalResourceOwner = nil
			return fmt.Errorf("failed to unmarshal RequestTypeApproval as RequestTypeApprovalResourceOwner: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER'
	if jsonDict["approverType"] == "USER" {
		// try to unmarshal JSON data into RequestTypeApprovalUser
		err = json.Unmarshal(data, &dst.RequestTypeApprovalUser)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalUser, return on the first match
		} else {
			dst.RequestTypeApprovalUser = nil
			return fmt.Errorf("failed to unmarshal RequestTypeApproval as RequestTypeApprovalUser: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTypeApproval) MarshalJSON() ([]byte, error) {
	if src.RequestTypeApprovalManager != nil {
		return json.Marshal(&src.RequestTypeApprovalManager)
	}

	if src.RequestTypeApprovalMemberOf != nil {
		return json.Marshal(&src.RequestTypeApprovalMemberOf)
	}

	if src.RequestTypeApprovalResourceOwner != nil {
		return json.Marshal(&src.RequestTypeApprovalResourceOwner)
	}

	if src.RequestTypeApprovalUser != nil {
		return json.Marshal(&src.RequestTypeApprovalUser)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTypeApproval) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestTypeApprovalManager != nil {
		return obj.RequestTypeApprovalManager
	}

	if obj.RequestTypeApprovalMemberOf != nil {
		return obj.RequestTypeApprovalMemberOf
	}

	if obj.RequestTypeApprovalResourceOwner != nil {
		return obj.RequestTypeApprovalResourceOwner
	}

	if obj.RequestTypeApprovalUser != nil {
		return obj.RequestTypeApprovalUser
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RequestTypeApproval) GetActualInstanceValue() interface{} {
	if obj.RequestTypeApprovalManager != nil {
		return *obj.RequestTypeApprovalManager
	}

	if obj.RequestTypeApprovalMemberOf != nil {
		return *obj.RequestTypeApprovalMemberOf
	}

	if obj.RequestTypeApprovalResourceOwner != nil {
		return *obj.RequestTypeApprovalResourceOwner
	}

	if obj.RequestTypeApprovalUser != nil {
		return *obj.RequestTypeApprovalUser
	}

	// all schemas are nil
	return nil
}

type NullableRequestTypeApproval struct {
	value *RequestTypeApproval
	isSet bool
}

func (v NullableRequestTypeApproval) Get() *RequestTypeApproval {
	return v.value
}

func (v *NullableRequestTypeApproval) Set(val *RequestTypeApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApproval(val *RequestTypeApproval) *NullableRequestTypeApproval {
	return &NullableRequestTypeApproval{value: val, isSet: true}
}

func (v NullableRequestTypeApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
