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
)

// RequestTypeApprovalSettingsSerial Settings for approval
type RequestTypeApprovalSettingsSerial struct {
	// When there are multiple approvals, an approval is not actionable until the previous approval has been approved. A denial will terminate the request and no subsequent approvals will be made actionable.
	Type string `json:"type"`
	// What approval(s) are required to grant access?  All specified approvals are considered required in order to fulfill an access request.  At least one approval must be specified for a request type. The maximum number of approvals is 5.  Approvals are serial.  When a request type has two approvals, and a user creates a request using that request type, then only the first approval will be immediately actionable by an approver.  After an approver has made a decision for that approval, then the second approval will be actionable by its approver.  The approval type of `RESOURCE_OWNER` is only supported for resource type of `GROUP`
	Approvals            []RequestTypeApproval `json:"approvals"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalSettingsSerial RequestTypeApprovalSettingsSerial

// NewRequestTypeApprovalSettingsSerial instantiates a new RequestTypeApprovalSettingsSerial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalSettingsSerial(type_ string, approvals []RequestTypeApproval) *RequestTypeApprovalSettingsSerial {
	this := RequestTypeApprovalSettingsSerial{}
	this.Type = type_
	this.Approvals = approvals
	return &this
}

// NewRequestTypeApprovalSettingsSerialWithDefaults instantiates a new RequestTypeApprovalSettingsSerial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalSettingsSerialWithDefaults() *RequestTypeApprovalSettingsSerial {
	this := RequestTypeApprovalSettingsSerial{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeApprovalSettingsSerial) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalSettingsSerial) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeApprovalSettingsSerial) SetType(v string) {
	o.Type = v
}

// GetApprovals returns the Approvals field value
func (o *RequestTypeApprovalSettingsSerial) GetApprovals() []RequestTypeApproval {
	if o == nil {
		var ret []RequestTypeApproval
		return ret
	}

	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalSettingsSerial) GetApprovalsOk() ([]RequestTypeApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvals, true
}

// SetApprovals sets field value
func (o *RequestTypeApprovalSettingsSerial) SetApprovals(v []RequestTypeApproval) {
	o.Approvals = v
}

func (o RequestTypeApprovalSettingsSerial) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["approvals"] = o.Approvals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeApprovalSettingsSerial) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeApprovalSettingsSerial := _RequestTypeApprovalSettingsSerial{}

	err = json.Unmarshal(bytes, &varRequestTypeApprovalSettingsSerial)
	if err == nil {
		*o = RequestTypeApprovalSettingsSerial(varRequestTypeApprovalSettingsSerial)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "approvals")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeApprovalSettingsSerial struct {
	value *RequestTypeApprovalSettingsSerial
	isSet bool
}

func (v NullableRequestTypeApprovalSettingsSerial) Get() *RequestTypeApprovalSettingsSerial {
	return v.value
}

func (v *NullableRequestTypeApprovalSettingsSerial) Set(val *RequestTypeApprovalSettingsSerial) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalSettingsSerial) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalSettingsSerial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalSettingsSerial(val *RequestTypeApprovalSettingsSerial) *NullableRequestTypeApprovalSettingsSerial {
	return &NullableRequestTypeApprovalSettingsSerial{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalSettingsSerial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalSettingsSerial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
