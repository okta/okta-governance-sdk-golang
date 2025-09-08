/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

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

// model_oneof.mustache
// RequestTypeApprovalWritable - request-type-approval-description.md
type RequestTypeApprovalWritable struct {
	RequestTypeApprovalManagerWritable       *RequestTypeApprovalManagerWritable
	RequestTypeApprovalMemberOfWritable      *RequestTypeApprovalMemberOfWritable
	RequestTypeApprovalResourceOwnerWritable *RequestTypeApprovalResourceOwnerWritable
	RequestTypeApprovalUserWritable          *RequestTypeApprovalUserWritable
}

// RequestTypeApprovalManagerWritableAsRequestTypeApprovalWritable is a convenience function that returns RequestTypeApprovalManagerWritable wrapped in RequestTypeApprovalWritable
func RequestTypeApprovalManagerWritableAsRequestTypeApprovalWritable(v *RequestTypeApprovalManagerWritable) RequestTypeApprovalWritable {
	return RequestTypeApprovalWritable{
		RequestTypeApprovalManagerWritable: v,
	}
}

// RequestTypeApprovalMemberOfWritableAsRequestTypeApprovalWritable is a convenience function that returns RequestTypeApprovalMemberOfWritable wrapped in RequestTypeApprovalWritable
func RequestTypeApprovalMemberOfWritableAsRequestTypeApprovalWritable(v *RequestTypeApprovalMemberOfWritable) RequestTypeApprovalWritable {
	return RequestTypeApprovalWritable{
		RequestTypeApprovalMemberOfWritable: v,
	}
}

// RequestTypeApprovalResourceOwnerWritableAsRequestTypeApprovalWritable is a convenience function that returns RequestTypeApprovalResourceOwnerWritable wrapped in RequestTypeApprovalWritable
func RequestTypeApprovalResourceOwnerWritableAsRequestTypeApprovalWritable(v *RequestTypeApprovalResourceOwnerWritable) RequestTypeApprovalWritable {
	return RequestTypeApprovalWritable{
		RequestTypeApprovalResourceOwnerWritable: v,
	}
}

// RequestTypeApprovalUserWritableAsRequestTypeApprovalWritable is a convenience function that returns RequestTypeApprovalUserWritable wrapped in RequestTypeApprovalWritable
func RequestTypeApprovalUserWritableAsRequestTypeApprovalWritable(v *RequestTypeApprovalUserWritable) RequestTypeApprovalWritable {
	return RequestTypeApprovalWritable{
		RequestTypeApprovalUserWritable: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequestTypeApprovalWritable) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'MANAGER'
	if jsonDict["approverType"] == "MANAGER" {
		// try to unmarshal JSON data into RequestTypeApprovalManagerWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalManagerWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalManagerWritable, return on the first match
		} else {
			dst.RequestTypeApprovalManagerWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalManagerWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MEMBER_OF'
	if jsonDict["approverType"] == "MEMBER_OF" {
		// try to unmarshal JSON data into RequestTypeApprovalMemberOfWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalMemberOfWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalMemberOfWritable, return on the first match
		} else {
			dst.RequestTypeApprovalMemberOfWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalMemberOfWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESOURCE_OWNER'
	if jsonDict["approverType"] == "RESOURCE_OWNER" {
		// try to unmarshal JSON data into RequestTypeApprovalResourceOwnerWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalResourceOwnerWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalResourceOwnerWritable, return on the first match
		} else {
			dst.RequestTypeApprovalResourceOwnerWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalResourceOwnerWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'USER'
	if jsonDict["approverType"] == "USER" {
		// try to unmarshal JSON data into RequestTypeApprovalUserWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalUserWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalUserWritable, return on the first match
		} else {
			dst.RequestTypeApprovalUserWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalUserWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-manager-writable'
	if jsonDict["approverType"] == "request-type-approval-manager-writable" {
		// try to unmarshal JSON data into RequestTypeApprovalManagerWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalManagerWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalManagerWritable, return on the first match
		} else {
			dst.RequestTypeApprovalManagerWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalManagerWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-member-of-writable'
	if jsonDict["approverType"] == "request-type-approval-member-of-writable" {
		// try to unmarshal JSON data into RequestTypeApprovalMemberOfWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalMemberOfWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalMemberOfWritable, return on the first match
		} else {
			dst.RequestTypeApprovalMemberOfWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalMemberOfWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-resource-owner-writable'
	if jsonDict["approverType"] == "request-type-approval-resource-owner-writable" {
		// try to unmarshal JSON data into RequestTypeApprovalResourceOwnerWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalResourceOwnerWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalResourceOwnerWritable, return on the first match
		} else {
			dst.RequestTypeApprovalResourceOwnerWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalResourceOwnerWritable: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-type-approval-user-writable'
	if jsonDict["approverType"] == "request-type-approval-user-writable" {
		// try to unmarshal JSON data into RequestTypeApprovalUserWritable
		err = json.Unmarshal(data, &dst.RequestTypeApprovalUserWritable)
		if err == nil {
			return nil // data stored in dst.RequestTypeApprovalUserWritable, return on the first match
		} else {
			dst.RequestTypeApprovalUserWritable = nil
			return fmt.Errorf("Failed to unmarshal RequestTypeApprovalWritable as RequestTypeApprovalUserWritable: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestTypeApprovalWritable) MarshalJSON() ([]byte, error) {
	if src.RequestTypeApprovalManagerWritable != nil {
		return json.Marshal(&src.RequestTypeApprovalManagerWritable)
	}

	if src.RequestTypeApprovalMemberOfWritable != nil {
		return json.Marshal(&src.RequestTypeApprovalMemberOfWritable)
	}

	if src.RequestTypeApprovalResourceOwnerWritable != nil {
		return json.Marshal(&src.RequestTypeApprovalResourceOwnerWritable)
	}

	if src.RequestTypeApprovalUserWritable != nil {
		return json.Marshal(&src.RequestTypeApprovalUserWritable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestTypeApprovalWritable) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestTypeApprovalManagerWritable != nil {
		return obj.RequestTypeApprovalManagerWritable
	}

	if obj.RequestTypeApprovalMemberOfWritable != nil {
		return obj.RequestTypeApprovalMemberOfWritable
	}

	if obj.RequestTypeApprovalResourceOwnerWritable != nil {
		return obj.RequestTypeApprovalResourceOwnerWritable
	}

	if obj.RequestTypeApprovalUserWritable != nil {
		return obj.RequestTypeApprovalUserWritable
	}

	// all schemas are nil
	return nil
}

type NullableRequestTypeApprovalWritable struct {
	value *RequestTypeApprovalWritable
	isSet bool
}

func (v NullableRequestTypeApprovalWritable) Get() *RequestTypeApprovalWritable {
	return v.value
}

func (v *NullableRequestTypeApprovalWritable) Set(val *RequestTypeApprovalWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalWritable(val *RequestTypeApprovalWritable) *NullableRequestTypeApprovalWritable {
	return &NullableRequestTypeApprovalWritable{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
