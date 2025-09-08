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
// PatchLabelsInner - struct for PatchLabelsInner
type PatchLabelsInner struct {
	PatchLabelOperation      *PatchLabelOperation
	PatchLabelValueOperation *PatchLabelValueOperation
}

// PatchLabelOperationAsPatchLabelsInner is a convenience function that returns PatchLabelOperation wrapped in PatchLabelsInner
func PatchLabelOperationAsPatchLabelsInner(v *PatchLabelOperation) PatchLabelsInner {
	return PatchLabelsInner{
		PatchLabelOperation: v,
	}
}

// PatchLabelValueOperationAsPatchLabelsInner is a convenience function that returns PatchLabelValueOperation wrapped in PatchLabelsInner
func PatchLabelValueOperationAsPatchLabelsInner(v *PatchLabelValueOperation) PatchLabelsInner {
	return PatchLabelsInner{
		PatchLabelValueOperation: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *PatchLabelsInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'LABEL-CATEGORY'
	if jsonDict["refType"] == "LABEL-CATEGORY" {
		// try to unmarshal JSON data into PatchLabelOperation
		err = json.Unmarshal(data, &dst.PatchLabelOperation)
		if err == nil {
			return nil // data stored in dst.PatchLabelOperation, return on the first match
		} else {
			dst.PatchLabelOperation = nil
			return fmt.Errorf("Failed to unmarshal PatchLabelsInner as PatchLabelOperation: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LABEL-VALUE'
	if jsonDict["refType"] == "LABEL-VALUE" {
		// try to unmarshal JSON data into PatchLabelValueOperation
		err = json.Unmarshal(data, &dst.PatchLabelValueOperation)
		if err == nil {
			return nil // data stored in dst.PatchLabelValueOperation, return on the first match
		} else {
			dst.PatchLabelValueOperation = nil
			return fmt.Errorf("Failed to unmarshal PatchLabelsInner as PatchLabelValueOperation: %s", err.Error())
		}
	}

	// check if the discriminator value is 'patch-label-operation'
	if jsonDict["refType"] == "patch-label-operation" {
		// try to unmarshal JSON data into PatchLabelOperation
		err = json.Unmarshal(data, &dst.PatchLabelOperation)
		if err == nil {
			return nil // data stored in dst.PatchLabelOperation, return on the first match
		} else {
			dst.PatchLabelOperation = nil
			return fmt.Errorf("Failed to unmarshal PatchLabelsInner as PatchLabelOperation: %s", err.Error())
		}
	}

	// check if the discriminator value is 'patch-label-value-operation'
	if jsonDict["refType"] == "patch-label-value-operation" {
		// try to unmarshal JSON data into PatchLabelValueOperation
		err = json.Unmarshal(data, &dst.PatchLabelValueOperation)
		if err == nil {
			return nil // data stored in dst.PatchLabelValueOperation, return on the first match
		} else {
			dst.PatchLabelValueOperation = nil
			return fmt.Errorf("Failed to unmarshal PatchLabelsInner as PatchLabelValueOperation: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchLabelsInner) MarshalJSON() ([]byte, error) {
	if src.PatchLabelOperation != nil {
		return json.Marshal(&src.PatchLabelOperation)
	}

	if src.PatchLabelValueOperation != nil {
		return json.Marshal(&src.PatchLabelValueOperation)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchLabelsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PatchLabelOperation != nil {
		return obj.PatchLabelOperation
	}

	if obj.PatchLabelValueOperation != nil {
		return obj.PatchLabelValueOperation
	}

	// all schemas are nil
	return nil
}

type NullablePatchLabelsInner struct {
	value *PatchLabelsInner
	isSet bool
}

func (v NullablePatchLabelsInner) Get() *PatchLabelsInner {
	return v.value
}

func (v *NullablePatchLabelsInner) Set(val *PatchLabelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchLabelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchLabelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchLabelsInner(val *PatchLabelsInner) *NullablePatchLabelsInner {
	return &NullablePatchLabelsInner{value: val, isSet: true}
}

func (v NullablePatchLabelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchLabelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
