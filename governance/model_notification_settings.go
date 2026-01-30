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
)

// checks if the NotificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSettings{}

// NotificationSettings Settings for email notifications to be sent to the reviewers at different stages of a campaign. All properties are optional.
type NotificationSettings struct {
	// A Boolean value to indicate whether a notification should be sent to the reviewer when actionable reviews are assigned.
	NotifyReviewerWhenReviewAssigned *bool `json:"notifyReviewerWhenReviewAssigned,omitempty"`
	// A Boolean value to indicate whether a notification should be sent to the reviewers when campaign has come to an end.
	NotifyReviewerAtCampaignEnd *bool `json:"notifyReviewerAtCampaignEnd,omitempty"`
	// Specifies, in seconds, the time a reminder is sent to reviewers before the campaign closes. You can send up to three notifications. For example, the following array, `[86400, 172800, 604800]`, sends reminder notifications 7 days, 2 days, and 1 day before the campaign closes. By default, reminders are sent 2 days and 1 day before the campaign closes.
	RemindersReviewerBeforeCampaignCloseInSecs []int32 `json:"remindersReviewerBeforeCampaignCloseInSecs,omitempty"`
	// A boolean value to indicate whether a notification should be sent to the reviewer when reviews are over due.
	NotifyReviewerWhenOverdue NullableBool `json:"notifyReviewerWhenOverdue,omitempty"`
	// A boolean value to indicate whether a notification should be sent to the reviewer during the midpoint of the review process.
	NotifyReviewerDuringMidpointOfReview NullableBool `json:"notifyReviewerDuringMidpointOfReview,omitempty"`
	// Applicable for multi level campaigns. A boolean value to indicate whether a notification should be sent to the reviewer when a given reviewer level period is about to end.
	NotifyReviewPeriodEnd NullableBool `json:"notifyReviewPeriodEnd,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _NotificationSettings NotificationSettings

// NewNotificationSettings instantiates a new NotificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSettings() *NotificationSettings {
	this := NotificationSettings{}
	var notifyReviewerWhenReviewAssigned bool = false
	this.NotifyReviewerWhenReviewAssigned = &notifyReviewerWhenReviewAssigned
	var notifyReviewerAtCampaignEnd bool = false
	this.NotifyReviewerAtCampaignEnd = &notifyReviewerAtCampaignEnd
	var notifyReviewerWhenOverdue bool = false
	this.NotifyReviewerWhenOverdue = *NewNullableBool(&notifyReviewerWhenOverdue)
	var notifyReviewerDuringMidpointOfReview bool = false
	this.NotifyReviewerDuringMidpointOfReview = *NewNullableBool(&notifyReviewerDuringMidpointOfReview)
	var notifyReviewPeriodEnd bool = false
	this.NotifyReviewPeriodEnd = *NewNullableBool(&notifyReviewPeriodEnd)
	return &this
}

// NewNotificationSettingsWithDefaults instantiates a new NotificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSettingsWithDefaults() *NotificationSettings {
	this := NotificationSettings{}
	var notifyReviewerWhenReviewAssigned bool = false
	this.NotifyReviewerWhenReviewAssigned = &notifyReviewerWhenReviewAssigned
	var notifyReviewerAtCampaignEnd bool = false
	this.NotifyReviewerAtCampaignEnd = &notifyReviewerAtCampaignEnd
	var notifyReviewerWhenOverdue bool = false
	this.NotifyReviewerWhenOverdue = *NewNullableBool(&notifyReviewerWhenOverdue)
	var notifyReviewerDuringMidpointOfReview bool = false
	this.NotifyReviewerDuringMidpointOfReview = *NewNullableBool(&notifyReviewerDuringMidpointOfReview)
	var notifyReviewPeriodEnd bool = false
	this.NotifyReviewPeriodEnd = *NewNullableBool(&notifyReviewPeriodEnd)
	return &this
}

// GetNotifyReviewerWhenReviewAssigned returns the NotifyReviewerWhenReviewAssigned field value if set, zero value otherwise.
func (o *NotificationSettings) GetNotifyReviewerWhenReviewAssigned() bool {
	if o == nil || IsNil(o.NotifyReviewerWhenReviewAssigned) {
		var ret bool
		return ret
	}
	return *o.NotifyReviewerWhenReviewAssigned
}

// GetNotifyReviewerWhenReviewAssignedOk returns a tuple with the NotifyReviewerWhenReviewAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetNotifyReviewerWhenReviewAssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyReviewerWhenReviewAssigned) {
		return nil, false
	}
	return o.NotifyReviewerWhenReviewAssigned, true
}

// HasNotifyReviewerWhenReviewAssigned returns a boolean if a field has been set.
func (o *NotificationSettings) HasNotifyReviewerWhenReviewAssigned() bool {
	if o != nil && !IsNil(o.NotifyReviewerWhenReviewAssigned) {
		return true
	}

	return false
}

// SetNotifyReviewerWhenReviewAssigned gets a reference to the given bool and assigns it to the NotifyReviewerWhenReviewAssigned field.
func (o *NotificationSettings) SetNotifyReviewerWhenReviewAssigned(v bool) {
	o.NotifyReviewerWhenReviewAssigned = &v
}

// GetNotifyReviewerAtCampaignEnd returns the NotifyReviewerAtCampaignEnd field value if set, zero value otherwise.
func (o *NotificationSettings) GetNotifyReviewerAtCampaignEnd() bool {
	if o == nil || IsNil(o.NotifyReviewerAtCampaignEnd) {
		var ret bool
		return ret
	}
	return *o.NotifyReviewerAtCampaignEnd
}

// GetNotifyReviewerAtCampaignEndOk returns a tuple with the NotifyReviewerAtCampaignEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetNotifyReviewerAtCampaignEndOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyReviewerAtCampaignEnd) {
		return nil, false
	}
	return o.NotifyReviewerAtCampaignEnd, true
}

// HasNotifyReviewerAtCampaignEnd returns a boolean if a field has been set.
func (o *NotificationSettings) HasNotifyReviewerAtCampaignEnd() bool {
	if o != nil && !IsNil(o.NotifyReviewerAtCampaignEnd) {
		return true
	}

	return false
}

// SetNotifyReviewerAtCampaignEnd gets a reference to the given bool and assigns it to the NotifyReviewerAtCampaignEnd field.
func (o *NotificationSettings) SetNotifyReviewerAtCampaignEnd(v bool) {
	o.NotifyReviewerAtCampaignEnd = &v
}

// GetRemindersReviewerBeforeCampaignCloseInSecs returns the RemindersReviewerBeforeCampaignCloseInSecs field value if set, zero value otherwise.
func (o *NotificationSettings) GetRemindersReviewerBeforeCampaignCloseInSecs() []int32 {
	if o == nil || IsNil(o.RemindersReviewerBeforeCampaignCloseInSecs) {
		var ret []int32
		return ret
	}
	return o.RemindersReviewerBeforeCampaignCloseInSecs
}

// GetRemindersReviewerBeforeCampaignCloseInSecsOk returns a tuple with the RemindersReviewerBeforeCampaignCloseInSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetRemindersReviewerBeforeCampaignCloseInSecsOk() ([]int32, bool) {
	if o == nil || IsNil(o.RemindersReviewerBeforeCampaignCloseInSecs) {
		return nil, false
	}
	return o.RemindersReviewerBeforeCampaignCloseInSecs, true
}

// HasRemindersReviewerBeforeCampaignCloseInSecs returns a boolean if a field has been set.
func (o *NotificationSettings) HasRemindersReviewerBeforeCampaignCloseInSecs() bool {
	if o != nil && !IsNil(o.RemindersReviewerBeforeCampaignCloseInSecs) {
		return true
	}

	return false
}

// SetRemindersReviewerBeforeCampaignCloseInSecs gets a reference to the given []int32 and assigns it to the RemindersReviewerBeforeCampaignCloseInSecs field.
func (o *NotificationSettings) SetRemindersReviewerBeforeCampaignCloseInSecs(v []int32) {
	o.RemindersReviewerBeforeCampaignCloseInSecs = v
}

// GetNotifyReviewerWhenOverdue returns the NotifyReviewerWhenOverdue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSettings) GetNotifyReviewerWhenOverdue() bool {
	if o == nil || IsNil(o.NotifyReviewerWhenOverdue.Get()) {
		var ret bool
		return ret
	}
	return *o.NotifyReviewerWhenOverdue.Get()
}

// GetNotifyReviewerWhenOverdueOk returns a tuple with the NotifyReviewerWhenOverdue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSettings) GetNotifyReviewerWhenOverdueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyReviewerWhenOverdue.Get(), o.NotifyReviewerWhenOverdue.IsSet()
}

// HasNotifyReviewerWhenOverdue returns a boolean if a field has been set.
func (o *NotificationSettings) HasNotifyReviewerWhenOverdue() bool {
	if o != nil && o.NotifyReviewerWhenOverdue.IsSet() {
		return true
	}

	return false
}

// SetNotifyReviewerWhenOverdue gets a reference to the given NullableBool and assigns it to the NotifyReviewerWhenOverdue field.
func (o *NotificationSettings) SetNotifyReviewerWhenOverdue(v bool) {
	o.NotifyReviewerWhenOverdue.Set(&v)
}

// SetNotifyReviewerWhenOverdueNil sets the value for NotifyReviewerWhenOverdue to be an explicit nil
func (o *NotificationSettings) SetNotifyReviewerWhenOverdueNil() {
	o.NotifyReviewerWhenOverdue.Set(nil)
}

// UnsetNotifyReviewerWhenOverdue ensures that no value is present for NotifyReviewerWhenOverdue, not even an explicit nil
func (o *NotificationSettings) UnsetNotifyReviewerWhenOverdue() {
	o.NotifyReviewerWhenOverdue.Unset()
}

// GetNotifyReviewerDuringMidpointOfReview returns the NotifyReviewerDuringMidpointOfReview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSettings) GetNotifyReviewerDuringMidpointOfReview() bool {
	if o == nil || IsNil(o.NotifyReviewerDuringMidpointOfReview.Get()) {
		var ret bool
		return ret
	}
	return *o.NotifyReviewerDuringMidpointOfReview.Get()
}

// GetNotifyReviewerDuringMidpointOfReviewOk returns a tuple with the NotifyReviewerDuringMidpointOfReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSettings) GetNotifyReviewerDuringMidpointOfReviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyReviewerDuringMidpointOfReview.Get(), o.NotifyReviewerDuringMidpointOfReview.IsSet()
}

// HasNotifyReviewerDuringMidpointOfReview returns a boolean if a field has been set.
func (o *NotificationSettings) HasNotifyReviewerDuringMidpointOfReview() bool {
	if o != nil && o.NotifyReviewerDuringMidpointOfReview.IsSet() {
		return true
	}

	return false
}

// SetNotifyReviewerDuringMidpointOfReview gets a reference to the given NullableBool and assigns it to the NotifyReviewerDuringMidpointOfReview field.
func (o *NotificationSettings) SetNotifyReviewerDuringMidpointOfReview(v bool) {
	o.NotifyReviewerDuringMidpointOfReview.Set(&v)
}

// SetNotifyReviewerDuringMidpointOfReviewNil sets the value for NotifyReviewerDuringMidpointOfReview to be an explicit nil
func (o *NotificationSettings) SetNotifyReviewerDuringMidpointOfReviewNil() {
	o.NotifyReviewerDuringMidpointOfReview.Set(nil)
}

// UnsetNotifyReviewerDuringMidpointOfReview ensures that no value is present for NotifyReviewerDuringMidpointOfReview, not even an explicit nil
func (o *NotificationSettings) UnsetNotifyReviewerDuringMidpointOfReview() {
	o.NotifyReviewerDuringMidpointOfReview.Unset()
}

// GetNotifyReviewPeriodEnd returns the NotifyReviewPeriodEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSettings) GetNotifyReviewPeriodEnd() bool {
	if o == nil || IsNil(o.NotifyReviewPeriodEnd.Get()) {
		var ret bool
		return ret
	}
	return *o.NotifyReviewPeriodEnd.Get()
}

// GetNotifyReviewPeriodEndOk returns a tuple with the NotifyReviewPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSettings) GetNotifyReviewPeriodEndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyReviewPeriodEnd.Get(), o.NotifyReviewPeriodEnd.IsSet()
}

// HasNotifyReviewPeriodEnd returns a boolean if a field has been set.
func (o *NotificationSettings) HasNotifyReviewPeriodEnd() bool {
	if o != nil && o.NotifyReviewPeriodEnd.IsSet() {
		return true
	}

	return false
}

// SetNotifyReviewPeriodEnd gets a reference to the given NullableBool and assigns it to the NotifyReviewPeriodEnd field.
func (o *NotificationSettings) SetNotifyReviewPeriodEnd(v bool) {
	o.NotifyReviewPeriodEnd.Set(&v)
}

// SetNotifyReviewPeriodEndNil sets the value for NotifyReviewPeriodEnd to be an explicit nil
func (o *NotificationSettings) SetNotifyReviewPeriodEndNil() {
	o.NotifyReviewPeriodEnd.Set(nil)
}

// UnsetNotifyReviewPeriodEnd ensures that no value is present for NotifyReviewPeriodEnd, not even an explicit nil
func (o *NotificationSettings) UnsetNotifyReviewPeriodEnd() {
	o.NotifyReviewPeriodEnd.Unset()
}

func (o NotificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotifyReviewerWhenReviewAssigned) {
		toSerialize["notifyReviewerWhenReviewAssigned"] = o.NotifyReviewerWhenReviewAssigned
	}
	if !IsNil(o.NotifyReviewerAtCampaignEnd) {
		toSerialize["notifyReviewerAtCampaignEnd"] = o.NotifyReviewerAtCampaignEnd
	}
	if !IsNil(o.RemindersReviewerBeforeCampaignCloseInSecs) {
		toSerialize["remindersReviewerBeforeCampaignCloseInSecs"] = o.RemindersReviewerBeforeCampaignCloseInSecs
	}
	if o.NotifyReviewerWhenOverdue.IsSet() {
		toSerialize["notifyReviewerWhenOverdue"] = o.NotifyReviewerWhenOverdue.Get()
	}
	if o.NotifyReviewerDuringMidpointOfReview.IsSet() {
		toSerialize["notifyReviewerDuringMidpointOfReview"] = o.NotifyReviewerDuringMidpointOfReview.Get()
	}
	if o.NotifyReviewPeriodEnd.IsSet() {
		toSerialize["notifyReviewPeriodEnd"] = o.NotifyReviewPeriodEnd.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationSettings) UnmarshalJSON(data []byte) (err error) {
	varNotificationSettings := _NotificationSettings{}

	err = json.Unmarshal(data, &varNotificationSettings)

	if err != nil {
		return err
	}

	*o = NotificationSettings(varNotificationSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "notifyReviewerWhenReviewAssigned")
		delete(additionalProperties, "notifyReviewerAtCampaignEnd")
		delete(additionalProperties, "remindersReviewerBeforeCampaignCloseInSecs")
		delete(additionalProperties, "notifyReviewerWhenOverdue")
		delete(additionalProperties, "notifyReviewerDuringMidpointOfReview")
		delete(additionalProperties, "notifyReviewPeriodEnd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationSettings struct {
	value *NotificationSettings
	isSet bool
}

func (v NullableNotificationSettings) Get() *NotificationSettings {
	return v.value
}

func (v *NullableNotificationSettings) Set(val *NotificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSettings(val *NotificationSettings) *NullableNotificationSettings {
	return &NullableNotificationSettings{value: val, isSet: true}
}

func (v NullableNotificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
