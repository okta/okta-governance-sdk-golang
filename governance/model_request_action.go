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
// RequestAction - An action taken as part of this request, null if there will never be actions for this request.
type RequestAction struct {
	RequestActionCompleted *RequestActionCompleted
	RequestActionPending   *RequestActionPending
}

// RequestActionCompletedAsRequestAction is a convenience function that returns RequestActionCompleted wrapped in RequestAction
func RequestActionCompletedAsRequestAction(v *RequestActionCompleted) RequestAction {
	return RequestAction{
		RequestActionCompleted: v,
	}
}

// RequestActionPendingAsRequestAction is a convenience function that returns RequestActionPending wrapped in RequestAction
func RequestActionPendingAsRequestAction(v *RequestActionPending) RequestAction {
	return RequestAction{
		RequestActionPending: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *RequestAction) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'COMPLETED'
	if jsonDict["status"] == "COMPLETED" {
		// try to unmarshal JSON data into RequestActionCompleted
		err = json.Unmarshal(data, &dst.RequestActionCompleted)
		if err == nil {
			return nil // data stored in dst.RequestActionCompleted, return on the first match
		} else {
			dst.RequestActionCompleted = nil
			return fmt.Errorf("Failed to unmarshal RequestAction as RequestActionCompleted: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PENDING'
	if jsonDict["status"] == "PENDING" {
		// try to unmarshal JSON data into RequestActionPending
		err = json.Unmarshal(data, &dst.RequestActionPending)
		if err == nil {
			return nil // data stored in dst.RequestActionPending, return on the first match
		} else {
			dst.RequestActionPending = nil
			return fmt.Errorf("Failed to unmarshal RequestAction as RequestActionPending: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-action-completed'
	if jsonDict["status"] == "request-action-completed" {
		// try to unmarshal JSON data into RequestActionCompleted
		err = json.Unmarshal(data, &dst.RequestActionCompleted)
		if err == nil {
			return nil // data stored in dst.RequestActionCompleted, return on the first match
		} else {
			dst.RequestActionCompleted = nil
			return fmt.Errorf("Failed to unmarshal RequestAction as RequestActionCompleted: %s", err.Error())
		}
	}

	// check if the discriminator value is 'request-action-pending'
	if jsonDict["status"] == "request-action-pending" {
		// try to unmarshal JSON data into RequestActionPending
		err = json.Unmarshal(data, &dst.RequestActionPending)
		if err == nil {
			return nil // data stored in dst.RequestActionPending, return on the first match
		} else {
			dst.RequestActionPending = nil
			return fmt.Errorf("Failed to unmarshal RequestAction as RequestActionPending: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestAction) MarshalJSON() ([]byte, error) {
	if src.RequestActionCompleted != nil {
		return json.Marshal(&src.RequestActionCompleted)
	}

	if src.RequestActionPending != nil {
		return json.Marshal(&src.RequestActionPending)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestAction) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestActionCompleted != nil {
		return obj.RequestActionCompleted
	}

	if obj.RequestActionPending != nil {
		return obj.RequestActionPending
	}

	// all schemas are nil
	return nil
}

type NullableRequestAction struct {
	value *RequestAction
	isSet bool
}

func (v NullableRequestAction) Get() *RequestAction {
	return v.value
}

func (v *NullableRequestAction) Set(val *RequestAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestAction(val *RequestAction) *NullableRequestAction {
	return &NullableRequestAction{value: val, isSet: true}
}

func (v NullableRequestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
