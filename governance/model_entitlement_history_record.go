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

// EntitlementHistoryRecord A single entry in the entitlement history, showing all effective entitlements at a specific point in time.
type EntitlementHistoryRecord struct {
	// The start date and time when the entitlements became effective
	StartDate *time.Time `json:"startDate,omitempty"`
	// The end date and time when the entitlements were superseded (if empty, the entitlements are currently effective)
	EndDate *time.Time `json:"endDate,omitempty"`
	// The current status of the entitlements for this history entry
	Lifecycle *string `json:"lifecycle,omitempty"`
	// List of entitlements for this history entry
	Entitlements         []EntitlementDetail `json:"entitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementHistoryRecord EntitlementHistoryRecord

// NewEntitlementHistoryRecord instantiates a new EntitlementHistoryRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementHistoryRecord() *EntitlementHistoryRecord {
	this := EntitlementHistoryRecord{}
	return &this
}

// NewEntitlementHistoryRecordWithDefaults instantiates a new EntitlementHistoryRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementHistoryRecordWithDefaults() *EntitlementHistoryRecord {
	this := EntitlementHistoryRecord{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *EntitlementHistoryRecord) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementHistoryRecord) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *EntitlementHistoryRecord) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *EntitlementHistoryRecord) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *EntitlementHistoryRecord) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementHistoryRecord) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *EntitlementHistoryRecord) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *EntitlementHistoryRecord) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EntitlementHistoryRecord) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementHistoryRecord) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EntitlementHistoryRecord) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EntitlementHistoryRecord) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *EntitlementHistoryRecord) GetEntitlements() []EntitlementDetail {
	if o == nil || o.Entitlements == nil {
		var ret []EntitlementDetail
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementHistoryRecord) GetEntitlementsOk() ([]EntitlementDetail, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *EntitlementHistoryRecord) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementDetail and assigns it to the Entitlements field.
func (o *EntitlementHistoryRecord) SetEntitlements(v []EntitlementDetail) {
	o.Entitlements = v
}

func (o EntitlementHistoryRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementHistoryRecord) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementHistoryRecord := _EntitlementHistoryRecord{}

	err = json.Unmarshal(bytes, &varEntitlementHistoryRecord)
	if err == nil {
		*o = EntitlementHistoryRecord(varEntitlementHistoryRecord)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "lifecycle")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementHistoryRecord struct {
	value *EntitlementHistoryRecord
	isSet bool
}

func (v NullableEntitlementHistoryRecord) Get() *EntitlementHistoryRecord {
	return v.value
}

func (v *NullableEntitlementHistoryRecord) Set(val *EntitlementHistoryRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementHistoryRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementHistoryRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementHistoryRecord(val *EntitlementHistoryRecord) *NullableEntitlementHistoryRecord {
	return &NullableEntitlementHistoryRecord{value: val, isSet: true}
}

func (v NullableEntitlementHistoryRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementHistoryRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
