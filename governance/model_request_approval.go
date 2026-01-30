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

// RequestApproval - An approval for this request. It may be in a PENDING or COMPLETED state.
type RequestApproval struct {
	RequestApprovalCompleted *RequestApprovalCompleted
	RequestApprovalPending   *RequestApprovalPending
}

// RequestApprovalCompletedAsRequestApproval is a convenience function that returns RequestApprovalCompleted wrapped in RequestApproval
func RequestApprovalCompletedAsRequestApproval(v *RequestApprovalCompleted) RequestApproval {
	return RequestApproval{
		RequestApprovalCompleted: v,
	}
}

// RequestApprovalPendingAsRequestApproval is a convenience function that returns RequestApprovalPending wrapped in RequestApproval
func RequestApprovalPendingAsRequestApproval(v *RequestApprovalPending) RequestApproval {
	return RequestApproval{
		RequestApprovalPending: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestApproval) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'COMPLETED'
	if jsonDict["status"] == "COMPLETED" {
		// try to unmarshal JSON data into RequestApprovalCompleted
		err = json.Unmarshal(data, &dst.RequestApprovalCompleted)
		if err == nil {
			return nil // data stored in dst.RequestApprovalCompleted, return on the first match
		} else {
			dst.RequestApprovalCompleted = nil
			return fmt.Errorf("failed to unmarshal RequestApproval as RequestApprovalCompleted: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PENDING'
	if jsonDict["status"] == "PENDING" {
		// try to unmarshal JSON data into RequestApprovalPending
		err = json.Unmarshal(data, &dst.RequestApprovalPending)
		if err == nil {
			return nil // data stored in dst.RequestApprovalPending, return on the first match
		} else {
			dst.RequestApprovalPending = nil
			return fmt.Errorf("failed to unmarshal RequestApproval as RequestApprovalPending: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestApproval) MarshalJSON() ([]byte, error) {
	if src.RequestApprovalCompleted != nil {
		return json.Marshal(&src.RequestApprovalCompleted)
	}

	if src.RequestApprovalPending != nil {
		return json.Marshal(&src.RequestApprovalPending)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestApproval) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestApprovalCompleted != nil {
		return obj.RequestApprovalCompleted
	}

	if obj.RequestApprovalPending != nil {
		return obj.RequestApprovalPending
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RequestApproval) GetActualInstanceValue() interface{} {
	if obj.RequestApprovalCompleted != nil {
		return *obj.RequestApprovalCompleted
	}

	if obj.RequestApprovalPending != nil {
		return *obj.RequestApprovalPending
	}

	// all schemas are nil
	return nil
}

type NullableRequestApproval struct {
	value *RequestApproval
	isSet bool
}

func (v NullableRequestApproval) Get() *RequestApproval {
	return v.value
}

func (v *NullableRequestApproval) Set(val *RequestApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestApproval(val *RequestApproval) *NullableRequestApproval {
	return &NullableRequestApproval{value: val, isSet: true}
}

func (v NullableRequestApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
