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

// RequestConditionReadOnly Read only properties that are common to sparse and full condition representations.
type RequestConditionReadOnly struct {
	// The ID of the request condition
	Id     string                 `json:"id" validate:"regexp=rco[0-9a-zA-Z]+"`
	Status RequestConditionStatus `json:"status"`
	// The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority.
	Priority             int32                 `json:"priority"`
	Links                RequestConditionLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionReadOnly RequestConditionReadOnly

// NewRequestConditionReadOnly instantiates a new RequestConditionReadOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionReadOnly(id string, status RequestConditionStatus, priority int32, links RequestConditionLinks) *RequestConditionReadOnly {
	this := RequestConditionReadOnly{}
	this.Id = id
	this.Status = status
	this.Priority = priority
	this.Links = links
	return &this
}

// NewRequestConditionReadOnlyWithDefaults instantiates a new RequestConditionReadOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionReadOnlyWithDefaults() *RequestConditionReadOnly {
	this := RequestConditionReadOnly{}
	return &this
}

// GetId returns the Id field value
func (o *RequestConditionReadOnly) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestConditionReadOnly) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestConditionReadOnly) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *RequestConditionReadOnly) GetStatus() RequestConditionStatus {
	if o == nil {
		var ret RequestConditionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestConditionReadOnly) GetStatusOk() (*RequestConditionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestConditionReadOnly) SetStatus(v RequestConditionStatus) {
	o.Status = v
}

// GetPriority returns the Priority field value
func (o *RequestConditionReadOnly) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *RequestConditionReadOnly) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *RequestConditionReadOnly) SetPriority(v int32) {
	o.Priority = v
}

// GetLinks returns the Links field value
func (o *RequestConditionReadOnly) GetLinks() RequestConditionLinks {
	if o == nil {
		var ret RequestConditionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestConditionReadOnly) GetLinksOk() (*RequestConditionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestConditionReadOnly) SetLinks(v RequestConditionLinks) {
	o.Links = v
}

func (o RequestConditionReadOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestConditionReadOnly) UnmarshalJSON(bytes []byte) (err error) {
	varRequestConditionReadOnly := _RequestConditionReadOnly{}

	err = json.Unmarshal(bytes, &varRequestConditionReadOnly)
	if err == nil {
		*o = RequestConditionReadOnly(varRequestConditionReadOnly)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestConditionReadOnly struct {
	value *RequestConditionReadOnly
	isSet bool
}

func (v NullableRequestConditionReadOnly) Get() *RequestConditionReadOnly {
	return v.value
}

func (v *NullableRequestConditionReadOnly) Set(val *RequestConditionReadOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionReadOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionReadOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionReadOnly(val *RequestConditionReadOnly) *NullableRequestConditionReadOnly {
	return &NullableRequestConditionReadOnly{value: val, isSet: true}
}

func (v NullableRequestConditionReadOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionReadOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
