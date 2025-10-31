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

// checks if the RequestFullApiIncompatible type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestFullApiIncompatible{}

// RequestFullApiIncompatible A Request where the request type has a CUSTOM setting. This representation is limited and does not contain `actions` and `approvals`.
type RequestFullApiIncompatible struct {
	// This request is associated with a request type with `CUSTOM` settings.
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

type _RequestFullApiIncompatible RequestFullApiIncompatible

// NewRequestFullApiIncompatible instantiates a new RequestFullApiIncompatible object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestFullApiIncompatible(type_ string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks, requestTypeId string, subject string, requesterUserIds []string, requestStatus RequestRequestStatus, resolved NullableTime, requesterFieldValues []FieldValue) *RequestFullApiIncompatible {
	this := RequestFullApiIncompatible{}
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

// NewRequestFullApiIncompatibleWithDefaults instantiates a new RequestFullApiIncompatible object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFullApiIncompatibleWithDefaults() *RequestFullApiIncompatible {
	this := RequestFullApiIncompatible{}
	return &this
}

// GetType returns the Type field value
func (o *RequestFullApiIncompatible) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestFullApiIncompatible) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *RequestFullApiIncompatible) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestFullApiIncompatible) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestFullApiIncompatible) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestFullApiIncompatible) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestFullApiIncompatible) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestFullApiIncompatible) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestFullApiIncompatible) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestFullApiIncompatible) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestFullApiIncompatible) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestFullApiIncompatible) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *RequestFullApiIncompatible) GetLinks() RequestLinks {
	if o == nil {
		var ret RequestLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetLinksOk() (*RequestLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestFullApiIncompatible) SetLinks(v RequestLinks) {
	o.Links = v
}

// GetRequestTypeId returns the RequestTypeId field value
func (o *RequestFullApiIncompatible) GetRequestTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestTypeId
}

// GetRequestTypeIdOk returns a tuple with the RequestTypeId field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetRequestTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestTypeId, true
}

// SetRequestTypeId sets field value
func (o *RequestFullApiIncompatible) SetRequestTypeId(v string) {
	o.RequestTypeId = v
}

// GetSubject returns the Subject field value
func (o *RequestFullApiIncompatible) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *RequestFullApiIncompatible) SetSubject(v string) {
	o.Subject = v
}

// GetRequesterUserIds returns the RequesterUserIds field value
func (o *RequestFullApiIncompatible) GetRequesterUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequesterUserIds
}

// GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetRequesterUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterUserIds, true
}

// SetRequesterUserIds sets field value
func (o *RequestFullApiIncompatible) SetRequesterUserIds(v []string) {
	o.RequesterUserIds = v
}

// GetRequestStatus returns the RequestStatus field value
func (o *RequestFullApiIncompatible) GetRequestStatus() RequestRequestStatus {
	if o == nil {
		var ret RequestRequestStatus
		return ret
	}

	return o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value
// and a boolean to check if the value has been set.
func (o *RequestFullApiIncompatible) GetRequestStatusOk() (*RequestRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestStatus, true
}

// SetRequestStatus sets field value
func (o *RequestFullApiIncompatible) SetRequestStatus(v RequestRequestStatus) {
	o.RequestStatus = v
}

// GetResolved returns the Resolved field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RequestFullApiIncompatible) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullApiIncompatible) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// SetResolved sets field value
func (o *RequestFullApiIncompatible) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// GetRequesterFieldValues returns the RequesterFieldValues field value
// If the value is explicit nil, the zero value for []FieldValue will be returned
func (o *RequestFullApiIncompatible) GetRequesterFieldValues() []FieldValue {
	if o == nil {
		var ret []FieldValue
		return ret
	}

	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullApiIncompatible) GetRequesterFieldValuesOk() ([]FieldValue, bool) {
	if o == nil || IsNil(o.RequesterFieldValues) {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// SetRequesterFieldValues sets field value
func (o *RequestFullApiIncompatible) SetRequesterFieldValues(v []FieldValue) {
	o.RequesterFieldValues = v
}

func (o RequestFullApiIncompatible) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestFullApiIncompatible) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	toSerialize["_links"] = o.Links
	toSerialize["requestTypeId"] = o.RequestTypeId
	toSerialize["subject"] = o.Subject
	toSerialize["requesterUserIds"] = o.RequesterUserIds
	toSerialize["requestStatus"] = o.RequestStatus
	toSerialize["resolved"] = o.Resolved.Get()
	if o.RequesterFieldValues != nil {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestFullApiIncompatible) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"_links",
		"requestTypeId",
		"subject",
		"requesterUserIds",
		"requestStatus",
		"resolved",
		"requesterFieldValues",
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

	varRequestFullApiIncompatible := _RequestFullApiIncompatible{}

	err = json.Unmarshal(data, &varRequestFullApiIncompatible)

	if err != nil {
		return err
	}

	*o = RequestFullApiIncompatible(varRequestFullApiIncompatible)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableRequestFullApiIncompatible struct {
	value *RequestFullApiIncompatible
	isSet bool
}

func (v NullableRequestFullApiIncompatible) Get() *RequestFullApiIncompatible {
	return v.value
}

func (v *NullableRequestFullApiIncompatible) Set(val *RequestFullApiIncompatible) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFullApiIncompatible) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFullApiIncompatible) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFullApiIncompatible(val *RequestFullApiIncompatible) *NullableRequestFullApiIncompatible {
	return &NullableRequestFullApiIncompatible{value: val, isSet: true}
}

func (v NullableRequestFullApiIncompatible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFullApiIncompatible) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
