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

// checks if the RequestTypeApprovalSettingsSerialWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeApprovalSettingsSerialWritable{}

// RequestTypeApprovalSettingsSerialWritable Settings for approval
type RequestTypeApprovalSettingsSerialWritable struct {
	// When there are multiple approvals, an approval is not actionable until the previous approval has been approved. A denial will terminate the request and no subsequent approvals will be made actionable.
	Type string `json:"type"`
	// What approval(s) are required to grant access?  All specified approvals are considered required in order to fulfill an access request.  At least one approval must be specified for a request type. The maximum number of approvals is 5.  Approvals are serial.  When a request type has two approvals, and a user creates a request using that request type, then only the first approval will be immediately actionable by an approver.  After an approver has made a decision for that approval, then the second approval will be actionable by its approver.  The approval type of `RESOURCE_OWNER` is only supported for resource type of `GROUP`
	Approvals            []RequestTypeApprovalWritable `json:"approvals"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeApprovalSettingsSerialWritable RequestTypeApprovalSettingsSerialWritable

// NewRequestTypeApprovalSettingsSerialWritable instantiates a new RequestTypeApprovalSettingsSerialWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeApprovalSettingsSerialWritable(type_ string, approvals []RequestTypeApprovalWritable) *RequestTypeApprovalSettingsSerialWritable {
	this := RequestTypeApprovalSettingsSerialWritable{}
	this.Type = type_
	this.Approvals = approvals
	return &this
}

// NewRequestTypeApprovalSettingsSerialWritableWithDefaults instantiates a new RequestTypeApprovalSettingsSerialWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeApprovalSettingsSerialWritableWithDefaults() *RequestTypeApprovalSettingsSerialWritable {
	this := RequestTypeApprovalSettingsSerialWritable{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeApprovalSettingsSerialWritable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalSettingsSerialWritable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeApprovalSettingsSerialWritable) SetType(v string) {
	o.Type = v
}

// GetApprovals returns the Approvals field value
func (o *RequestTypeApprovalSettingsSerialWritable) GetApprovals() []RequestTypeApprovalWritable {
	if o == nil {
		var ret []RequestTypeApprovalWritable
		return ret
	}

	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value
// and a boolean to check if the value has been set.
func (o *RequestTypeApprovalSettingsSerialWritable) GetApprovalsOk() ([]RequestTypeApprovalWritable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvals, true
}

// SetApprovals sets field value
func (o *RequestTypeApprovalSettingsSerialWritable) SetApprovals(v []RequestTypeApprovalWritable) {
	o.Approvals = v
}

func (o RequestTypeApprovalSettingsSerialWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeApprovalSettingsSerialWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["approvals"] = o.Approvals

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeApprovalSettingsSerialWritable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"approvals",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRequestTypeApprovalSettingsSerialWritable := _RequestTypeApprovalSettingsSerialWritable{}

	err = json.Unmarshal(data, &varRequestTypeApprovalSettingsSerialWritable)

	if err != nil {
		return err
	}

	*o = RequestTypeApprovalSettingsSerialWritable(varRequestTypeApprovalSettingsSerialWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "approvals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeApprovalSettingsSerialWritable struct {
	value *RequestTypeApprovalSettingsSerialWritable
	isSet bool
}

func (v NullableRequestTypeApprovalSettingsSerialWritable) Get() *RequestTypeApprovalSettingsSerialWritable {
	return v.value
}

func (v *NullableRequestTypeApprovalSettingsSerialWritable) Set(val *RequestTypeApprovalSettingsSerialWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeApprovalSettingsSerialWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeApprovalSettingsSerialWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeApprovalSettingsSerialWritable(val *RequestTypeApprovalSettingsSerialWritable) *NullableRequestTypeApprovalSettingsSerialWritable {
	return &NullableRequestTypeApprovalSettingsSerialWritable{value: val, isSet: true}
}

func (v NullableRequestTypeApprovalSettingsSerialWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeApprovalSettingsSerialWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
