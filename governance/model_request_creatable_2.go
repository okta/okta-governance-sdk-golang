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

// RequestCreatable2 Properties which are mutable in mutation operations
type RequestCreatable2 struct {
	Requested    RequestResourceCreatable `json:"requested"`
	RequestedBy  *TargetPrincipal         `json:"requestedBy,omitempty"`
	RequestedFor TargetPrincipal          `json:"requestedFor"`
	// The requester input fields required by the approval system.  **Note:** The fields required are determined by the approval system.  For the Okta approval system, the required fields are defined in the approval sequence. Ensure that the requester input fields match up with this definition to avoid request approval flow failure.  For external approval systems, the requester input fields are for recording purposes only and do not affect the approval process.
	RequesterFieldValues []RequestFieldValue `json:"requesterFieldValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestCreatable2 RequestCreatable2

// NewRequestCreatable2 instantiates a new RequestCreatable2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestCreatable2(requested RequestResourceCreatable, requestedFor TargetPrincipal) *RequestCreatable2 {
	this := RequestCreatable2{}
	this.Requested = requested
	this.RequestedFor = requestedFor
	return &this
}

// NewRequestCreatable2WithDefaults instantiates a new RequestCreatable2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestCreatable2WithDefaults() *RequestCreatable2 {
	this := RequestCreatable2{}
	return &this
}

// GetRequested returns the Requested field value
func (o *RequestCreatable2) GetRequested() RequestResourceCreatable {
	if o == nil {
		var ret RequestResourceCreatable
		return ret
	}

	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value
// and a boolean to check if the value has been set.
func (o *RequestCreatable2) GetRequestedOk() (*RequestResourceCreatable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requested, true
}

// SetRequested sets field value
func (o *RequestCreatable2) SetRequested(v RequestResourceCreatable) {
	o.Requested = v
}

// GetRequestedBy returns the RequestedBy field value if set, zero value otherwise.
func (o *RequestCreatable2) GetRequestedBy() TargetPrincipal {
	if o == nil || o.RequestedBy == nil {
		var ret TargetPrincipal
		return ret
	}
	return *o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreatable2) GetRequestedByOk() (*TargetPrincipal, bool) {
	if o == nil || o.RequestedBy == nil {
		return nil, false
	}
	return o.RequestedBy, true
}

// HasRequestedBy returns a boolean if a field has been set.
func (o *RequestCreatable2) HasRequestedBy() bool {
	if o != nil && o.RequestedBy != nil {
		return true
	}

	return false
}

// SetRequestedBy gets a reference to the given TargetPrincipal and assigns it to the RequestedBy field.
func (o *RequestCreatable2) SetRequestedBy(v TargetPrincipal) {
	o.RequestedBy = &v
}

// GetRequestedFor returns the RequestedFor field value
func (o *RequestCreatable2) GetRequestedFor() TargetPrincipal {
	if o == nil {
		var ret TargetPrincipal
		return ret
	}

	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value
// and a boolean to check if the value has been set.
func (o *RequestCreatable2) GetRequestedForOk() (*TargetPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedFor, true
}

// SetRequestedFor sets field value
func (o *RequestCreatable2) SetRequestedFor(v TargetPrincipal) {
	o.RequestedFor = v
}

// GetRequesterFieldValues returns the RequesterFieldValues field value if set, zero value otherwise.
func (o *RequestCreatable2) GetRequesterFieldValues() []RequestFieldValue {
	if o == nil || o.RequesterFieldValues == nil {
		var ret []RequestFieldValue
		return ret
	}
	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreatable2) GetRequesterFieldValuesOk() ([]RequestFieldValue, bool) {
	if o == nil || o.RequesterFieldValues == nil {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// HasRequesterFieldValues returns a boolean if a field has been set.
func (o *RequestCreatable2) HasRequesterFieldValues() bool {
	if o != nil && o.RequesterFieldValues != nil {
		return true
	}

	return false
}

// SetRequesterFieldValues gets a reference to the given []RequestFieldValue and assigns it to the RequesterFieldValues field.
func (o *RequestCreatable2) SetRequesterFieldValues(v []RequestFieldValue) {
	o.RequesterFieldValues = v
}

func (o RequestCreatable2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requested"] = o.Requested
	}
	if o.RequestedBy != nil {
		toSerialize["requestedBy"] = o.RequestedBy
	}
	if true {
		toSerialize["requestedFor"] = o.RequestedFor
	}
	if o.RequesterFieldValues != nil {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestCreatable2) UnmarshalJSON(bytes []byte) (err error) {
	varRequestCreatable2 := _RequestCreatable2{}

	err = json.Unmarshal(bytes, &varRequestCreatable2)
	if err == nil {
		*o = RequestCreatable2(varRequestCreatable2)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requested")
		delete(additionalProperties, "requestedBy")
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requesterFieldValues")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestCreatable2 struct {
	value *RequestCreatable2
	isSet bool
}

func (v NullableRequestCreatable2) Get() *RequestCreatable2 {
	return v.value
}

func (v *NullableRequestCreatable2) Set(val *RequestCreatable2) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCreatable2) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCreatable2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCreatable2(val *RequestCreatable2) *NullableRequestCreatable2 {
	return &NullableRequestCreatable2{value: val, isSet: true}
}

func (v NullableRequestCreatable2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCreatable2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
