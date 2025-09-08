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
	"time"
)

// RequestConditionFull A condition delineates the criteria governing access requests, specifying the requester's identity, the permitted access scope, and the manner in which access can be requested.
type RequestConditionFull struct {
	RequesterSettings      RequesterSettingsFullRequesterSettings     `json:"requesterSettings"`
	AccessScopeSettings    AccessScopeSettingsFullAccessScopeSettings `json:"accessScopeSettings"`
	AccessDurationSettings *AccessDurationSettingsFull                `json:"accessDurationSettings,omitempty"`
	// If an approval sequence was deleted, then conditions referencing it will become invalid and the approvalSequenceId will not be present.
	ApprovalSequenceId *string `json:"approvalSequenceId,omitempty"`
	// Writable unique key on Create. Modifiable on update.
	Name string `json:"name"`
	// Human readable description.
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

type _RequestConditionFull RequestConditionFull

// NewRequestConditionFull instantiates a new RequestConditionFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestConditionFull(requesterSettings RequesterSettingsFullRequesterSettings, accessScopeSettings AccessScopeSettingsFullAccessScopeSettings, name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestConditionLinks, status RequestConditionStatus, priority int32) *RequestConditionFull {
	this := RequestConditionFull{}
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

// NewRequestConditionFullWithDefaults instantiates a new RequestConditionFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestConditionFullWithDefaults() *RequestConditionFull {
	this := RequestConditionFull{}
	return &this
}

// GetRequesterSettings returns the RequesterSettings field value
func (o *RequestConditionFull) GetRequesterSettings() RequesterSettingsFullRequesterSettings {
	if o == nil {
		var ret RequesterSettingsFullRequesterSettings
		return ret
	}

	return o.RequesterSettings
}

// GetRequesterSettingsOk returns a tuple with the RequesterSettings field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetRequesterSettingsOk() (*RequesterSettingsFullRequesterSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequesterSettings, true
}

// SetRequesterSettings sets field value
func (o *RequestConditionFull) SetRequesterSettings(v RequesterSettingsFullRequesterSettings) {
	o.RequesterSettings = v
}

// GetAccessScopeSettings returns the AccessScopeSettings field value
func (o *RequestConditionFull) GetAccessScopeSettings() AccessScopeSettingsFullAccessScopeSettings {
	if o == nil {
		var ret AccessScopeSettingsFullAccessScopeSettings
		return ret
	}

	return o.AccessScopeSettings
}

// GetAccessScopeSettingsOk returns a tuple with the AccessScopeSettings field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetAccessScopeSettingsOk() (*AccessScopeSettingsFullAccessScopeSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessScopeSettings, true
}

// SetAccessScopeSettings sets field value
func (o *RequestConditionFull) SetAccessScopeSettings(v AccessScopeSettingsFullAccessScopeSettings) {
	o.AccessScopeSettings = v
}

// GetAccessDurationSettings returns the AccessDurationSettings field value if set, zero value otherwise.
func (o *RequestConditionFull) GetAccessDurationSettings() AccessDurationSettingsFull {
	if o == nil || o.AccessDurationSettings == nil {
		var ret AccessDurationSettingsFull
		return ret
	}
	return *o.AccessDurationSettings
}

// GetAccessDurationSettingsOk returns a tuple with the AccessDurationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetAccessDurationSettingsOk() (*AccessDurationSettingsFull, bool) {
	if o == nil || o.AccessDurationSettings == nil {
		return nil, false
	}
	return o.AccessDurationSettings, true
}

// HasAccessDurationSettings returns a boolean if a field has been set.
func (o *RequestConditionFull) HasAccessDurationSettings() bool {
	if o != nil && o.AccessDurationSettings != nil {
		return true
	}

	return false
}

// SetAccessDurationSettings gets a reference to the given AccessDurationSettingsFull and assigns it to the AccessDurationSettings field.
func (o *RequestConditionFull) SetAccessDurationSettings(v AccessDurationSettingsFull) {
	o.AccessDurationSettings = &v
}

// GetApprovalSequenceId returns the ApprovalSequenceId field value if set, zero value otherwise.
func (o *RequestConditionFull) GetApprovalSequenceId() string {
	if o == nil || o.ApprovalSequenceId == nil {
		var ret string
		return ret
	}
	return *o.ApprovalSequenceId
}

// GetApprovalSequenceIdOk returns a tuple with the ApprovalSequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetApprovalSequenceIdOk() (*string, bool) {
	if o == nil || o.ApprovalSequenceId == nil {
		return nil, false
	}
	return o.ApprovalSequenceId, true
}

// HasApprovalSequenceId returns a boolean if a field has been set.
func (o *RequestConditionFull) HasApprovalSequenceId() bool {
	if o != nil && o.ApprovalSequenceId != nil {
		return true
	}

	return false
}

// SetApprovalSequenceId gets a reference to the given string and assigns it to the ApprovalSequenceId field.
func (o *RequestConditionFull) SetApprovalSequenceId(v string) {
	o.ApprovalSequenceId = &v
}

// GetName returns the Name field value
func (o *RequestConditionFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestConditionFull) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestConditionFull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestConditionFull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestConditionFull) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *RequestConditionFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestConditionFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestConditionFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestConditionFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestConditionFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestConditionFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestConditionFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestConditionFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestConditionFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestConditionFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *RequestConditionFull) GetLinks() RequestConditionLinks {
	if o == nil {
		var ret RequestConditionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetLinksOk() (*RequestConditionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestConditionFull) SetLinks(v RequestConditionLinks) {
	o.Links = v
}

// GetStatus returns the Status field value
func (o *RequestConditionFull) GetStatus() RequestConditionStatus {
	if o == nil {
		var ret RequestConditionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetStatusOk() (*RequestConditionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestConditionFull) SetStatus(v RequestConditionStatus) {
	o.Status = v
}

// GetPriority returns the Priority field value
func (o *RequestConditionFull) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *RequestConditionFull) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *RequestConditionFull) SetPriority(v int32) {
	o.Priority = v
}

func (o RequestConditionFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requesterSettings"] = o.RequesterSettings
	}
	if true {
		toSerialize["accessScopeSettings"] = o.AccessScopeSettings
	}
	if o.AccessDurationSettings != nil {
		toSerialize["accessDurationSettings"] = o.AccessDurationSettings
	}
	if o.ApprovalSequenceId != nil {
		toSerialize["approvalSequenceId"] = o.ApprovalSequenceId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestConditionFull) UnmarshalJSON(bytes []byte) (err error) {
	varRequestConditionFull := _RequestConditionFull{}

	err = json.Unmarshal(bytes, &varRequestConditionFull)
	if err == nil {
		*o = RequestConditionFull(varRequestConditionFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
	} else {
		return err
	}

	return err
}

type NullableRequestConditionFull struct {
	value *RequestConditionFull
	isSet bool
}

func (v NullableRequestConditionFull) Get() *RequestConditionFull {
	return v.value
}

func (v *NullableRequestConditionFull) Set(val *RequestConditionFull) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestConditionFull) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestConditionFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestConditionFull(val *RequestConditionFull) *NullableRequestConditionFull {
	return &NullableRequestConditionFull{value: val, isSet: true}
}

func (v NullableRequestConditionFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestConditionFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
