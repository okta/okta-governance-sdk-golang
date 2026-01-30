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

// checks if the RequestConditionSparse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestConditionSparse{}

// RequestConditionSparse A condition delineates the criteria governing access requests, specifying the requester's identity, the permitted access scope, and the manner in which access can be requested.
type RequestConditionSparse struct {
	RequesterSettings      RequesterSettings           `json:"requesterSettings"`
	AccessScopeSettings    AccessScopeSettings         `json:"accessScopeSettings"`
	AccessDurationSettings *AccessDurationSettingsFull `json:"accessDurationSettings,omitempty"`
	// If an approval sequence was deleted, then conditions referencing it will become invalid and the approvalSequenceId will not be present.
	ApprovalSequenceId *string `json:"approvalSequenceId,omitempty"`
	// Writable unique key on create. Modifiable on update.
	Name string `json:"name"`
	// Human readable description
	Description *string `json:"description,omitempty"`
	// The ID of the request condition
	Id string `json:"id" validate:"regexp=rco[0-9a-zA-Z]+"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string                 `json:"lastUpdatedBy"`
	Links         RequestConditionLinks  `json:"_links"`
	Status        RequestConditionStatus `json:"status"`
	// The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority.
	Priority             int32 `json:"priority"`
	AdditionalProperties map[string]interface{}
}

type _RequestConditionSparse RequestConditionSparse

// NewRequestConditionSparse instantiates a new RequestConditionSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionSparse(requesterSettings RequesterSettings, accessScopeSettings AccessScopeSettings, name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestConditionLinks, status RequestConditionStatus, priority int32) *RequestConditionSparse {
	this := RequestConditionSparse{}
	this.Name = name
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.Status = status
	this.Priority = priority
	return &this
}

// NewRequestConditionSparseWithDefaults instantiates a new RequestConditionSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionSparseWithDefaults() *RequestConditionSparse {
	this := RequestConditionSparse{}
	return &this
}

// GetRequesterSettings returns the RequesterSettings field value
func (o *RequestConditionSparse) GetRequesterSettings() RequesterSettings {
	if o == nil {
		var ret RequesterSettings
		return ret
	}

	return o.RequesterSettings
}

// GetRequesterSettingsOk returns a tuple with the RequesterSettings field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetRequesterSettingsOk() (*RequesterSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequesterSettings, true
}

// SetRequesterSettings sets field value
func (o *RequestConditionSparse) SetRequesterSettings(v RequesterSettings) {
	o.RequesterSettings = v
}

// GetAccessScopeSettings returns the AccessScopeSettings field value
func (o *RequestConditionSparse) GetAccessScopeSettings() AccessScopeSettings {
	if o == nil {
		var ret AccessScopeSettings
		return ret
	}

	return o.AccessScopeSettings
}

// GetAccessScopeSettingsOk returns a tuple with the AccessScopeSettings field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetAccessScopeSettingsOk() (*AccessScopeSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessScopeSettings, true
}

// SetAccessScopeSettings sets field value
func (o *RequestConditionSparse) SetAccessScopeSettings(v AccessScopeSettings) {
	o.AccessScopeSettings = v
}

// GetAccessDurationSettings returns the AccessDurationSettings field value if set, zero value otherwise.
func (o *RequestConditionSparse) GetAccessDurationSettings() AccessDurationSettingsFull {
	if o == nil || IsNil(o.AccessDurationSettings) {
		var ret AccessDurationSettingsFull
		return ret
	}
	return *o.AccessDurationSettings
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool) {
	if o == nil || IsNil(o.AccessDurationSettings) {
		return nil, false
	}
	return o.AccessDurationSettings, true
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RequestConditionSparse) HasAccessDurationSettings() bool {
	if o != nil && !IsNil(o.AccessDurationSettings) {
		return true
	}

	return false
}

// SetAccessDurationSettings gets a reference to the given AccessDurationSettingsFull and assigns it to the AccessDurationSettings field.
func (o *RequestConditionSparse) SetAccessDurationSettings(v AccessDurationSettingsFull) {
	o.AccessDurationSettings = &v
}

// GetApprovalSequenceId returns the ApprovalSequenceId field value if set, zero value otherwise.
func (o *RequestConditionSparse) GetApprovalSequenceId() string {
	if o == nil || IsNil(o.ApprovalSequenceId) {
		var ret string
		return ret
	}
	return *o.ApprovalSequenceId
}

// GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetApprovalSequenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovalSequenceId) {
		return nil, false
	}
	return o.ApprovalSequenceId, true
}

// HasApprovalSequenceId returns a boolean if a field has been set.
func (o *RequestConditionSparse) HasApprovalSequenceId() bool {
	if o != nil && !IsNil(o.ApprovalSequenceId) {
		return true
	}

	return false
}

// SetApprovalSequenceId gets a reference to the given string and assigns it to the ApprovalSequenceId field.
func (o *RequestConditionSparse) SetApprovalSequenceId(v string) {
	o.ApprovalSequenceId = &v
}

// GetName returns the Name field value
func (o *RequestConditionSparse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestConditionSparse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestConditionSparse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestConditionSparse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestConditionSparse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *RequestConditionSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestConditionSparse) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestConditionSparse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestConditionSparse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestConditionSparse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestConditionSparse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestConditionSparse) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestConditionSparse) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestConditionSparse) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestConditionSparse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *RequestConditionSparse) GetLinks() RequestConditionLinks {
	if o == nil {
		var ret RequestConditionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetLinksOk() (*RequestConditionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestConditionSparse) SetLinks(v RequestConditionLinks) {
	o.Links = v
}

// GetStatus returns the Status field value
func (o *RequestConditionSparse) GetStatus() RequestConditionStatus {
	if o == nil {
		var ret RequestConditionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetStatusOk() (*RequestConditionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestConditionSparse) SetStatus(v RequestConditionStatus) {
	o.Status = v
}

// GetPriority returns the Priority field value
func (o *RequestConditionSparse) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *RequestConditionSparse) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *RequestConditionSparse) SetPriority(v int32) {
	o.Priority = v
}

func (o RequestConditionSparse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestConditionSparse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requesterSettings"] = o.RequesterSettings
	toSerialize["accessScopeSettings"] = o.AccessScopeSettings
	if !IsNil(o.AccessDurationSettings) {
		toSerialize["accessDurationSettings"] = o.AccessDurationSettings
	}
	if !IsNil(o.ApprovalSequenceId) {
		toSerialize["approvalSequenceId"] = o.ApprovalSequenceId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	toSerialize["_links"] = o.Links
	toSerialize["status"] = o.Status
	toSerialize["priority"] = o.Priority

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestConditionSparse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requesterSettings",
		"accessScopeSettings",
		"name",
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"_links",
		"status",
		"priority",
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

	varRequestConditionSparse := _RequestConditionSparse{}

	err = json.Unmarshal(data, &varRequestConditionSparse)

	if err != nil {
		return err
	}

	*o = RequestConditionSparse(varRequestConditionSparse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requesterSettings")
		delete(additionalProperties, "accessScopeSettings")
		delete(additionalProperties, "accessDurationSettings")
		delete(additionalProperties, "approvalSequenceId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "status")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestConditionSparse struct {
	value *RequestConditionSparse
	isSet bool
}

func (v NullableRequestConditionSparse) Get() *RequestConditionSparse {
	return v.value
}

func (v *NullableRequestConditionSparse) Set(val *RequestConditionSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionSparse(val *RequestConditionSparse) *NullableRequestConditionSparse {
	return &NullableRequestConditionSparse{value: val, isSet: true}
}

func (v NullableRequestConditionSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
