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

// checks if the RequestActionCompleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestActionCompleted{}

// RequestActionCompleted A completed access request approval
type RequestActionCompleted struct {
	// The status of a completed approval
	Status string            `json:"status"`
	Action RequestActionEnum `json:"action"`
	// A unique identifier of the action taken in the request
	ActionId string `json:"actionId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	// Whether the action was successful or not
	ActionStatus string `json:"actionStatus"`
	// When the action was attempted
	ActionAttempted time.Time `json:"actionAttempted"`
	// When access was removed if the associated Request Type has an `accessDuration`. null if the Request Type does not have an `accessDuration`, or the duration has not yet expired.
	AccessRemoved NullableTime `json:"accessRemoved,omitempty"`
	// Human readable name of the resource
	ResourceName string `json:"resourceName"`
	// The Okta `app.id`, or `group.id` of the resource that can be requested with this Request Type.  * See [List applications](https://developer.okta.com/docs/reference/api/apps/#list-applications) to retrieve app IDs. * See [List groups](https://developer.okta.com/docs/reference/api/groups/#list-groups) to retrieve group IDs.
	ResourceId string `json:"resourceId"`
	// The type of resource
	ResourceType         string `json:"resourceType"`
	AdditionalProperties map[string]interface{}
}

type _RequestActionCompleted RequestActionCompleted

// NewRequestActionCompleted instantiates a new RequestActionCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestActionCompleted(status string, action RequestActionEnum, actionId string, actionStatus string, actionAttempted time.Time, resourceName string, resourceId string, resourceType string) *RequestActionCompleted {
	this := RequestActionCompleted{}
	this.ResourceName = resourceName
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	return &this
}

// NewRequestActionCompletedWithDefaults instantiates a new RequestActionCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestActionCompletedWithDefaults() *RequestActionCompleted {
	this := RequestActionCompleted{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestActionCompleted) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestActionCompleted) SetStatus(v string) {
	o.Status = v
}

// GetAction returns the Action field value
func (o *RequestActionCompleted) GetAction() RequestActionEnum {
	if o == nil {
		var ret RequestActionEnum
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetActionOk() (*RequestActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RequestActionCompleted) SetAction(v RequestActionEnum) {
	o.Action = v
}

// GetActionId returns the ActionId field value
func (o *RequestActionCompleted) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *RequestActionCompleted) SetActionId(v string) {
	o.ActionId = v
}

// GetActionStatus returns the ActionStatus field value
func (o *RequestActionCompleted) GetActionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetActionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionStatus, true
}

// SetActionStatus sets field value
func (o *RequestActionCompleted) SetActionStatus(v string) {
	o.ActionStatus = v
}

// GetActionAttempted returns the ActionAttempted field value
func (o *RequestActionCompleted) GetActionAttempted() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActionAttempted
}

// GetActionAttemptedOk returns a tuple with the ActionAttempted field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetActionAttemptedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionAttempted, true
}

// SetActionAttempted sets field value
func (o *RequestActionCompleted) SetActionAttempted(v time.Time) {
	o.ActionAttempted = v
}

// GetAccessRemoved returns the AccessRemoved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestActionCompleted) GetAccessRemoved() time.Time {
	if o == nil || IsNil(o.AccessRemoved.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AccessRemoved.Get()
}

// GetAccessRemovedOk returns a tuple with the AccessRemoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestActionCompleted) GetAccessRemovedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessRemoved.Get(), o.AccessRemoved.IsSet()
}

// HasAccessRemoved returns a boolean if a field has been set.
func (o *RequestActionCompleted) HasAccessRemoved() bool {
	if o != nil && o.AccessRemoved.IsSet() {
		return true
	}

	return false
}

// SetAccessRemoved gets a reference to the given NullableTime and assigns it to the AccessRemoved field.
func (o *RequestActionCompleted) SetAccessRemoved(v time.Time) {
	o.AccessRemoved.Set(&v)
}

// SetAccessRemovedNil sets the value for AccessRemoved to be an explicit nil
func (o *RequestActionCompleted) SetAccessRemovedNil() {
	o.AccessRemoved.Set(nil)
}

// UnsetAccessRemoved ensures that no value is present for AccessRemoved, not even an explicit nil
func (o *RequestActionCompleted) UnsetAccessRemoved() {
	o.AccessRemoved.Unset()
}

// GetResourceName returns the ResourceName field value
func (o *RequestActionCompleted) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *RequestActionCompleted) SetResourceName(v string) {
	o.ResourceName = v
}

// GetResourceId returns the ResourceId field value
func (o *RequestActionCompleted) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *RequestActionCompleted) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *RequestActionCompleted) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RequestActionCompleted) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *RequestActionCompleted) SetResourceType(v string) {
	o.ResourceType = v
}

func (o RequestActionCompleted) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestActionCompleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["action"] = o.Action
	toSerialize["actionId"] = o.ActionId
	toSerialize["actionStatus"] = o.ActionStatus
	toSerialize["actionAttempted"] = o.ActionAttempted
	if o.AccessRemoved.IsSet() {
		toSerialize["accessRemoved"] = o.AccessRemoved.Get()
	}
	toSerialize["resourceName"] = o.ResourceName
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceType"] = o.ResourceType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestActionCompleted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"action",
		"actionId",
		"actionStatus",
		"actionAttempted",
		"resourceName",
		"resourceId",
		"resourceType",
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

	varRequestActionCompleted := _RequestActionCompleted{}

	err = json.Unmarshal(data, &varRequestActionCompleted)

	if err != nil {
		return err
	}

	*o = RequestActionCompleted(varRequestActionCompleted)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "action")
		delete(additionalProperties, "actionId")
		delete(additionalProperties, "actionStatus")
		delete(additionalProperties, "actionAttempted")
		delete(additionalProperties, "accessRemoved")
		delete(additionalProperties, "resourceName")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestActionCompleted struct {
	value *RequestActionCompleted
	isSet bool
}

func (v NullableRequestActionCompleted) Get() *RequestActionCompleted {
	return v.value
}

func (v *NullableRequestActionCompleted) Set(val *RequestActionCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestActionCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestActionCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestActionCompleted(val *RequestActionCompleted) *NullableRequestActionCompleted {
	return &NullableRequestActionCompleted{value: val, isSet: true}
}

func (v NullableRequestActionCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestActionCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
