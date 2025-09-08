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

// RequestMinimalReadOnlyFields Properties contained by default in sparse representation
type RequestMinimalReadOnlyFields struct {
	RequestStatus *RequestRequestStatus `json:"requestStatus,omitempty"`
	// The Request Type enabling this Request.
	RequestTypeId *string `json:"requestTypeId,omitempty" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	// The date the request was resolved. The property may transition from having a value to null if the request is reopened.
	Resolved NullableTime `json:"resolved,omitempty"`
	// Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.
	RequesterFieldValues []FieldValue  `json:"requesterFieldValues,omitempty"`
	Links                *RequestLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestMinimalReadOnlyFields RequestMinimalReadOnlyFields

// NewRequestMinimalReadOnlyFields instantiates a new RequestMinimalReadOnlyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMinimalReadOnlyFields() *RequestMinimalReadOnlyFields {
	this := RequestMinimalReadOnlyFields{}
	return &this
}

// NewRequestMinimalReadOnlyFieldsWithDefaults instantiates a new RequestMinimalReadOnlyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMinimalReadOnlyFieldsWithDefaults() *RequestMinimalReadOnlyFields {
	this := RequestMinimalReadOnlyFields{}
	return &this
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields) GetRequestStatus() RequestRequestStatus {
	if o == nil || o.RequestStatus == nil {
		var ret RequestRequestStatus
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields) GetRequestStatusOk() (*RequestRequestStatus, bool) {
	if o == nil || o.RequestStatus == nil {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields) HasRequestStatus() bool {
	if o != nil && o.RequestStatus != nil {
		return true
	}

	return false
}

// SetRequestStatus gets a reference to the given RequestRequestStatus and assigns it to the RequestStatus field.
func (o *RequestMinimalReadOnlyFields) SetRequestStatus(v RequestRequestStatus) {
	o.RequestStatus = &v
}

// GetRequestTypeId returns the RequestTypeId field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields) GetRequestTypeId() string {
	if o == nil || o.RequestTypeId == nil {
		var ret string
		return ret
	}
	return *o.RequestTypeId
}

// GetRequestTypeIdOk returns a tuple with the RequestTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields) GetRequestTypeIdOk() (*string, bool) {
	if o == nil || o.RequestTypeId == nil {
		return nil, false
	}
	return o.RequestTypeId, true
}

// HasRequestTypeId returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields) HasRequestTypeId() bool {
	if o != nil && o.RequestTypeId != nil {
		return true
	}

	return false
}

// SetRequestTypeId gets a reference to the given string and assigns it to the RequestTypeId field.
func (o *RequestMinimalReadOnlyFields) SetRequestTypeId(v string) {
	o.RequestTypeId = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestMinimalReadOnlyFields) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMinimalReadOnlyFields) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields) HasResolved() bool {
	if o != nil && o.Resolved.IsSet() {
		return true
	}

	return false
}

// SetResolved gets a reference to the given NullableTime and assigns it to the Resolved field.
func (o *RequestMinimalReadOnlyFields) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil
func (o *RequestMinimalReadOnlyFields) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
func (o *RequestMinimalReadOnlyFields) UnsetResolved() {
	o.Resolved.Unset()
}

// GetRequesterFieldValues returns the RequesterFieldValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestMinimalReadOnlyFields) GetRequesterFieldValues() []FieldValue {
	if o == nil {
		var ret []FieldValue
		return ret
	}
	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMinimalReadOnlyFields) GetRequesterFieldValuesOk() ([]FieldValue, bool) {
	if o == nil || o.RequesterFieldValues == nil {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// HasRequesterFieldValues returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields) HasRequesterFieldValues() bool {
	if o != nil && o.RequesterFieldValues != nil {
		return true
	}

	return false
}

// SetRequesterFieldValues gets a reference to the given []FieldValue and assigns it to the RequesterFieldValues field.
func (o *RequestMinimalReadOnlyFields) SetRequesterFieldValues(v []FieldValue) {
	o.RequesterFieldValues = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields) GetLinks() RequestLinks {
	if o == nil || o.Links == nil {
		var ret RequestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields) GetLinksOk() (*RequestLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RequestLinks and assigns it to the Links field.
func (o *RequestMinimalReadOnlyFields) SetLinks(v RequestLinks) {
	o.Links = &v
}

func (o RequestMinimalReadOnlyFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestStatus != nil {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if o.RequestTypeId != nil {
		toSerialize["requestTypeId"] = o.RequestTypeId
	}
	if o.Resolved.IsSet() {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if o.RequesterFieldValues != nil {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestMinimalReadOnlyFields) UnmarshalJSON(bytes []byte) (err error) {
	varRequestMinimalReadOnlyFields := _RequestMinimalReadOnlyFields{}

	err = json.Unmarshal(bytes, &varRequestMinimalReadOnlyFields)
	if err == nil {
		*o = RequestMinimalReadOnlyFields(varRequestMinimalReadOnlyFields)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestStatus")
		delete(additionalProperties, "requestTypeId")
		delete(additionalProperties, "resolved")
		delete(additionalProperties, "requesterFieldValues")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestMinimalReadOnlyFields struct {
	value *RequestMinimalReadOnlyFields
	isSet bool
}

func (v NullableRequestMinimalReadOnlyFields) Get() *RequestMinimalReadOnlyFields {
	return v.value
}

func (v *NullableRequestMinimalReadOnlyFields) Set(val *RequestMinimalReadOnlyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMinimalReadOnlyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMinimalReadOnlyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMinimalReadOnlyFields(val *RequestMinimalReadOnlyFields) *NullableRequestMinimalReadOnlyFields {
	return &NullableRequestMinimalReadOnlyFields{value: val, isSet: true}
}

func (v NullableRequestMinimalReadOnlyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMinimalReadOnlyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
