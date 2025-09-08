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

// RequestFullApiCompatible Full representation of a Request resource.
type RequestFullApiCompatible struct {
	Approvals []RequestApproval `json:"approvals"`
	// A list of actions. Currently only supports one action per request.
	Actions []RequestAction `json:"actions"`
	// This request is associated with a request type with no `CUSTOM` settings.
	Type string `json:"type"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string       `json:"lastUpdatedBy"`
	Links         RequestLinks `json:"_links"`
	// The Request Type enabling this Request.
	RequestTypeId string `json:"requestTypeId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	// The subject of the request
	Subject string `json:"subject"`
	// A list of requester Okta user `id`s.
	RequesterUserIds []string             `json:"requesterUserIds"`
	RequestStatus    RequestRequestStatus `json:"requestStatus"`
	// The date the request was resolved. The property may transition from having a value to null if the request is reopened.
	Resolved NullableTime `json:"resolved"`
	// Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.
	RequesterFieldValues []FieldValue `json:"requesterFieldValues"`
	AdditionalProperties map[string]interface{}
}

type _RequestFullApiCompatible RequestFullApiCompatible

// NewRequestFullApiCompatible instantiates a new RequestFullApiCompatible object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestFullApiCompatible(approvals []RequestApproval, actions []RequestAction, type_ string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks, requestTypeId string, subject string, requesterUserIds []string, requestStatus RequestRequestStatus, resolved NullableTime, requesterFieldValues []FieldValue) *RequestFullApiCompatible {
	this := RequestFullApiCompatible{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.RequestTypeId = requestTypeId
	this.Subject = subject
	this.RequesterUserIds = requesterUserIds
	this.RequestStatus = requestStatus
	this.Resolved = resolved
	this.RequesterFieldValues = requesterFieldValues
	return &this
}

// NewRequestFullApiCompatibleWithDefaults instantiates a new RequestFullApiCompatible object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFullApiCompatibleWithDefaults() *RequestFullApiCompatible {
	this := RequestFullApiCompatible{}
	return &this
}

// GetApprovals returns the Approvals field value
func (o *RequestFullApiCompatible) GetApprovals() []RequestApproval {
	if o == nil {
		var ret []RequestApproval
		return ret
	}

	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetApprovalsOk() ([]RequestApproval, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvals, true
}

// SetApprovals sets field value
func (o *RequestFullApiCompatible) SetApprovals(v []RequestApproval) {
	o.Approvals = v
}

// GetActions returns the Actions field value
// If the value is explicit nil, the zero value for []RequestAction will be returned
func (o *RequestFullApiCompatible) GetActions() []RequestAction {
	if o == nil {
		var ret []RequestAction
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullApiCompatible) GetActionsOk() ([]RequestAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *RequestFullApiCompatible) SetActions(v []RequestAction) {
	o.Actions = v
}

// GetType returns the Type field value
func (o *RequestFullApiCompatible) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestFullApiCompatible) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *RequestFullApiCompatible) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestFullApiCompatible) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestFullApiCompatible) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestFullApiCompatible) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestFullApiCompatible) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestFullApiCompatible) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestFullApiCompatible) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestFullApiCompatible) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestFullApiCompatible) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestFullApiCompatible) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *RequestFullApiCompatible) GetLinks() RequestLinks {
	if o == nil {
		var ret RequestLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetLinksOk() (*RequestLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestFullApiCompatible) SetLinks(v RequestLinks) {
	o.Links = v
}

// GetRequestTypeId returns the RequestTypeId field value
func (o *RequestFullApiCompatible) GetRequestTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestTypeId
}

// GetRequestTypeIdOk returns a tuple with the RequestTypeId field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetRequestTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTypeId, true
}

// SetRequestTypeId sets field value
func (o *RequestFullApiCompatible) SetRequestTypeId(v string) {
	o.RequestTypeId = v
}

// GetSubject returns the Subject field value
func (o *RequestFullApiCompatible) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *RequestFullApiCompatible) SetSubject(v string) {
	o.Subject = v
}

// GetRequesterUserIds returns the RequesterUserIds field value
func (o *RequestFullApiCompatible) GetRequesterUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequesterUserIds
}

// GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetRequesterUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterUserIds, true
}

// SetRequesterUserIds sets field value
func (o *RequestFullApiCompatible) SetRequesterUserIds(v []string) {
	o.RequesterUserIds = v
}

// GetRequestStatus returns the RequestStatus field value
func (o *RequestFullApiCompatible) GetRequestStatus() RequestRequestStatus {
	if o == nil {
		var ret RequestRequestStatus
		return ret
	}

	return o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiCompatible) GetRequestStatusOk() (*RequestRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestStatus, true
}

// SetRequestStatus sets field value
func (o *RequestFullApiCompatible) SetRequestStatus(v RequestRequestStatus) {
	o.RequestStatus = v
}

// GetResolved returns the Resolved field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RequestFullApiCompatible) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullApiCompatible) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// SetResolved sets field value
func (o *RequestFullApiCompatible) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// GetRequesterFieldValues returns the RequesterFieldValues field value
// If the value is explicit nil, the zero value for []FieldValue will be returned
func (o *RequestFullApiCompatible) GetRequesterFieldValues() []FieldValue {
	if o == nil {
		var ret []FieldValue
		return ret
	}

	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullApiCompatible) GetRequesterFieldValuesOk() ([]FieldValue, bool) {
	if o == nil || o.RequesterFieldValues == nil {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// SetRequesterFieldValues sets field value
func (o *RequestFullApiCompatible) SetRequesterFieldValues(v []FieldValue) {
	o.RequesterFieldValues = v
}

func (o RequestFullApiCompatible) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["approvals"] = o.Approvals
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if true {
		toSerialize["type"] = o.Type
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
		toSerialize["requestTypeId"] = o.RequestTypeId
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["requesterUserIds"] = o.RequesterUserIds
	}
	if true {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if true {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if o.RequesterFieldValues != nil {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestFullApiCompatible) UnmarshalJSON(bytes []byte) (err error) {
	varRequestFullApiCompatible := _RequestFullApiCompatible{}

	err = json.Unmarshal(bytes, &varRequestFullApiCompatible)
	if err == nil {
		*o = RequestFullApiCompatible(varRequestFullApiCompatible)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "approvals")
		delete(additionalProperties, "actions")
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "requestTypeId")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "requesterUserIds")
		delete(additionalProperties, "requestStatus")
		delete(additionalProperties, "resolved")
		delete(additionalProperties, "requesterFieldValues")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestFullApiCompatible struct {
	value *RequestFullApiCompatible
	isSet bool
}

func (v NullableRequestFullApiCompatible) Get() *RequestFullApiCompatible {
	return v.value
}

func (v *NullableRequestFullApiCompatible) Set(val *RequestFullApiCompatible) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFullApiCompatible) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFullApiCompatible) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFullApiCompatible(val *RequestFullApiCompatible) *NullableRequestFullApiCompatible {
	return &NullableRequestFullApiCompatible{value: val, isSet: true}
}

func (v NullableRequestFullApiCompatible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFullApiCompatible) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
