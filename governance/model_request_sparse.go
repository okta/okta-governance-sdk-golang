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

// RequestSparse Sparse representation of a Request resource
type RequestSparse struct {
	// Whether the request's request type has CUSTOM settings or not.
	Type          string               `json:"type"`
	RequestStatus RequestRequestStatus `json:"requestStatus"`
	// The date the request was resolved. The property may transition from having a value to null if the request is reopened.
	Resolved NullableTime `json:"resolved"`
	Links    RequestLinks `json:"_links"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// The request type `id`.
	RequestTypeId string `json:"requestTypeId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	// The subject of the request
	Subject string `json:"subject"`
	// A list of requester Okta user `id`s.
	RequesterUserIds     []string `json:"requesterUserIds"`
	AdditionalProperties map[string]interface{}
}

type _RequestSparse RequestSparse

// NewRequestSparse instantiates a new RequestSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSparse(type_ string, requestStatus RequestRequestStatus, resolved NullableTime, links RequestLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, requestTypeId string, subject string, requesterUserIds []string) *RequestSparse {
	this := RequestSparse{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.RequestTypeId = requestTypeId
	this.Subject = subject
	this.RequesterUserIds = requesterUserIds
	return &this
}

// NewRequestSparseWithDefaults instantiates a new RequestSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSparseWithDefaults() *RequestSparse {
	this := RequestSparse{}
	return &this
}

// GetType returns the Type field value
func (o *RequestSparse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestSparse) SetType(v string) {
	o.Type = v
}

// GetRequestStatus returns the RequestStatus field value
func (o *RequestSparse) GetRequestStatus() RequestRequestStatus {
	if o == nil {
		var ret RequestRequestStatus
		return ret
	}

	return o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetRequestStatusOk() (*RequestRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestStatus, true
}

// SetRequestStatus sets field value
func (o *RequestSparse) SetRequestStatus(v RequestRequestStatus) {
	o.RequestStatus = v
}

// GetResolved returns the Resolved field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RequestSparse) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSparse) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// SetResolved sets field value
func (o *RequestSparse) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// GetLinks returns the Links field value
func (o *RequestSparse) GetLinks() RequestLinks {
	if o == nil {
		var ret RequestLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetLinksOk() (*RequestLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestSparse) SetLinks(v RequestLinks) {
	o.Links = v
}

// GetId returns the Id field value
func (o *RequestSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestSparse) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestSparse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestSparse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestSparse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestSparse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestSparse) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestSparse) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestSparse) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestSparse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetRequestTypeId returns the RequestTypeId field value
func (o *RequestSparse) GetRequestTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestTypeId
}

// GetRequestTypeIdOk returns a tuple with the RequestTypeId field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetRequestTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTypeId, true
}

// SetRequestTypeId sets field value
func (o *RequestSparse) SetRequestTypeId(v string) {
	o.RequestTypeId = v
}

// GetSubject returns the Subject field value
func (o *RequestSparse) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *RequestSparse) SetSubject(v string) {
	o.Subject = v
}

// GetRequesterUserIds returns the RequesterUserIds field value
func (o *RequestSparse) GetRequesterUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequesterUserIds
}

// GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field value
// and a boolean to check if the value has been set.
func (o *RequestSparse) GetRequesterUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterUserIds, true
}

// SetRequesterUserIds sets field value
func (o *RequestSparse) SetRequesterUserIds(v []string) {
	o.RequesterUserIds = v
}

func (o RequestSparse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if true {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if true {
		toSerialize["_links"] = o.Links
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
		toSerialize["requestTypeId"] = o.RequestTypeId
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["requesterUserIds"] = o.RequesterUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestSparse) UnmarshalJSON(bytes []byte) (err error) {
	varRequestSparse := _RequestSparse{}

	err = json.Unmarshal(bytes, &varRequestSparse)
	if err == nil {
		*o = RequestSparse(varRequestSparse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "requestStatus")
		delete(additionalProperties, "resolved")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "requestTypeId")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "requesterUserIds")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestSparse struct {
	value *RequestSparse
	isSet bool
}

func (v NullableRequestSparse) Get() *RequestSparse {
	return v.value
}

func (v *NullableRequestSparse) Set(val *RequestSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSparse(val *RequestSparse) *NullableRequestSparse {
	return &NullableRequestSparse{value: val, isSet: true}
}

func (v NullableRequestSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
