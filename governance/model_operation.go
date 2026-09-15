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
	"time"
)

// checks if the Operation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Operation{}

// Operation Representation of an operation
type Operation struct {
	// ID of the operation
	Id string `json:"id"`
	// The operation type
	Type   string          `json:"type"`
	Status OperationStatus `json:"status"`
	// The ISO 8601 formatted date and time of when the operation was created
	Created *time.Time `json:"created,omitempty"`
	// The ISO 8601 formatted date and time of when the operation completed
	Completed *time.Time `json:"completed,omitempty"`
	// Elapsed time since the start of the operation, in seconds
	TimeElapsedInSeconds *int32         `json:"timeElapsedInSeconds,omitempty"`
	Links                OperationLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _Operation Operation

// NewOperation instantiates a new Operation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperation(id string, type_ string, status OperationStatus, links OperationLinks) *Operation {
	this := Operation{}
	this.Id = id
	this.Type = type_
	this.Status = status
	this.Links = links
	return &this
}

// NewOperationWithDefaults instantiates a new Operation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationWithDefaults() *Operation {
	this := Operation{}
	return &this
}

// GetId returns the Id field value
func (o *Operation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Operation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Operation) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Operation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Operation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Operation) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *Operation) GetStatus() OperationStatus {
	if o == nil {
		var ret OperationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Operation) GetStatusOk() (*OperationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Operation) SetStatus(v OperationStatus) {
	o.Status = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Operation) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Operation) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Operation) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *Operation) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *Operation) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *Operation) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetTimeElapsedInSeconds returns the TimeElapsedInSeconds field value if set, zero value otherwise.
func (o *Operation) GetTimeElapsedInSeconds() int32 {
	if o == nil || IsNil(o.TimeElapsedInSeconds) {
		var ret int32
		return ret
	}
	return *o.TimeElapsedInSeconds
}

// GetTimeElapsedInSecondsOk returns a tuple with the TimeElapsedInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetTimeElapsedInSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeElapsedInSeconds) {
		return nil, false
	}
	return o.TimeElapsedInSeconds, true
}

// HasTimeElapsedInSeconds returns a boolean if a field has been set.
func (o *Operation) HasTimeElapsedInSeconds() bool {
	if o != nil && !IsNil(o.TimeElapsedInSeconds) {
		return true
	}

	return false
}

// SetTimeElapsedInSeconds gets a reference to the given int32 and assigns it to the TimeElapsedInSeconds field.
func (o *Operation) SetTimeElapsedInSeconds(v int32) {
	o.TimeElapsedInSeconds = &v
}

// GetLinks returns the Links field value
func (o *Operation) GetLinks() OperationLinks {
	if o == nil {
		var ret OperationLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *Operation) GetLinksOk() (*OperationLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *Operation) SetLinks(v OperationLinks) {
	o.Links = v
}

func (o Operation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Operation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.TimeElapsedInSeconds) {
		toSerialize["timeElapsedInSeconds"] = o.TimeElapsedInSeconds
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Operation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"status",
		"_links",
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

	varOperation := _Operation{}

	err = json.Unmarshal(data, &varOperation)

	if err != nil {
		return err
	}

	*o = Operation(varOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "timeElapsedInSeconds")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOperation struct {
	value *Operation
	isSet bool
}

func (v NullableOperation) Get() *Operation {
	return v.value
}

func (v *NullableOperation) Set(val *Operation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation(val *Operation) *NullableOperation {
	return &NullableOperation{value: val, isSet: true}
}

func (v NullableOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
