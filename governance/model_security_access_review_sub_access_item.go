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
)

// checks if the SecurityAccessReviewSubAccessItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewSubAccessItem{}

// SecurityAccessReviewSubAccessItem struct for SecurityAccessReviewSubAccessItem
type SecurityAccessReviewSubAccessItem struct {
	// The ID of the sub-access item
	Id string `json:"id"`
	// The name of the sub-access item
	Name string                                `json:"name"`
	Type SecurityAccessReviewSubAccessItemType `json:"type"`
	// The ID of the resource for the sub-access item
	ResourceId        string                                           `json:"resourceId"`
	Severity          SecurityAccessReviewAccessItemSeverity           `json:"severity"`
	SupportedActions  []SecurityAccessReviewAccessItemSupportedAction  `json:"supportedActions"`
	RemediationStatus *SecurityAccessReviewAccessItemRemediationStatus `json:"remediationStatus,omitempty"`
	// The reasons for manual remediation
	ManualRemediationTypes []SecurityAccessReviewAccessItemManualRemediationType `json:"manualRemediationTypes,omitempty"`
	EntitlementInfo        *SecurityAccessReviewSubAccessItemEntitlementInfo     `json:"entitlementInfo,omitempty"`
	GroupInfo              *SecurityAccessReviewSubAccessItemGroupInfo           `json:"groupInfo,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _SecurityAccessReviewSubAccessItem SecurityAccessReviewSubAccessItem

// NewSecurityAccessReviewSubAccessItem instantiates a new SecurityAccessReviewSubAccessItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewSubAccessItem(id string, name string, type_ SecurityAccessReviewSubAccessItemType, resourceId string, severity SecurityAccessReviewAccessItemSeverity, supportedActions []SecurityAccessReviewAccessItemSupportedAction) *SecurityAccessReviewSubAccessItem {
	this := SecurityAccessReviewSubAccessItem{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.ResourceId = resourceId
	this.Severity = severity
	this.SupportedActions = supportedActions
	return &this
}

// NewSecurityAccessReviewSubAccessItemWithDefaults instantiates a new SecurityAccessReviewSubAccessItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewSubAccessItemWithDefaults() *SecurityAccessReviewSubAccessItem {
	this := SecurityAccessReviewSubAccessItem{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewSubAccessItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewSubAccessItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SecurityAccessReviewSubAccessItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReviewSubAccessItem) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *SecurityAccessReviewSubAccessItem) GetType() SecurityAccessReviewSubAccessItemType {
	if o == nil {
		var ret SecurityAccessReviewSubAccessItemType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetTypeOk() (*SecurityAccessReviewSubAccessItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityAccessReviewSubAccessItem) SetType(v SecurityAccessReviewSubAccessItemType) {
	o.Type = v
}

// GetResourceId returns the ResourceId field value
func (o *SecurityAccessReviewSubAccessItem) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *SecurityAccessReviewSubAccessItem) SetResourceId(v string) {
	o.ResourceId = v
}

// GetSeverity returns the Severity field value
func (o *SecurityAccessReviewSubAccessItem) GetSeverity() SecurityAccessReviewAccessItemSeverity {
	if o == nil {
		var ret SecurityAccessReviewAccessItemSeverity
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetSeverityOk() (*SecurityAccessReviewAccessItemSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *SecurityAccessReviewSubAccessItem) SetSeverity(v SecurityAccessReviewAccessItemSeverity) {
	o.Severity = v
}

// GetSupportedActions returns the SupportedActions field value
func (o *SecurityAccessReviewSubAccessItem) GetSupportedActions() []SecurityAccessReviewAccessItemSupportedAction {
	if o == nil {
		var ret []SecurityAccessReviewAccessItemSupportedAction
		return ret
	}

	return o.SupportedActions
}

// GetSupportedActionsOk returns a tuple with the SupportedActions field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetSupportedActionsOk() ([]SecurityAccessReviewAccessItemSupportedAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedActions, true
}

// SetSupportedActions sets field value
func (o *SecurityAccessReviewSubAccessItem) SetSupportedActions(v []SecurityAccessReviewAccessItemSupportedAction) {
	o.SupportedActions = v
}

// GetRemediationStatus returns the RemediationStatus field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItem) GetRemediationStatus() SecurityAccessReviewAccessItemRemediationStatus {
	if o == nil || IsNil(o.RemediationStatus) {
		var ret SecurityAccessReviewAccessItemRemediationStatus
		return ret
	}
	return *o.RemediationStatus
}

// GetRemediationStatusOk returns a tuple with the RemediationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetRemediationStatusOk() (*SecurityAccessReviewAccessItemRemediationStatus, bool) {
	if o == nil || IsNil(o.RemediationStatus) {
		return nil, false
	}
	return o.RemediationStatus, true
}

// HasRemediationStatus returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItem) HasRemediationStatus() bool {
	if o != nil && !IsNil(o.RemediationStatus) {
		return true
	}

	return false
}

// SetRemediationStatus gets a reference to the given SecurityAccessReviewAccessItemRemediationStatus and assigns it to the RemediationStatus field.
func (o *SecurityAccessReviewSubAccessItem) SetRemediationStatus(v SecurityAccessReviewAccessItemRemediationStatus) {
	o.RemediationStatus = &v
}

// GetManualRemediationTypes returns the ManualRemediationTypes field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItem) GetManualRemediationTypes() []SecurityAccessReviewAccessItemManualRemediationType {
	if o == nil || IsNil(o.ManualRemediationTypes) {
		var ret []SecurityAccessReviewAccessItemManualRemediationType
		return ret
	}
	return o.ManualRemediationTypes
}

// GetManualRemediationTypesOk returns a tuple with the ManualRemediationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetManualRemediationTypesOk() ([]SecurityAccessReviewAccessItemManualRemediationType, bool) {
	if o == nil || IsNil(o.ManualRemediationTypes) {
		return nil, false
	}
	return o.ManualRemediationTypes, true
}

// HasManualRemediationTypes returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItem) HasManualRemediationTypes() bool {
	if o != nil && !IsNil(o.ManualRemediationTypes) {
		return true
	}

	return false
}

// SetManualRemediationTypes gets a reference to the given []SecurityAccessReviewAccessItemManualRemediationType and assigns it to the ManualRemediationTypes field.
func (o *SecurityAccessReviewSubAccessItem) SetManualRemediationTypes(v []SecurityAccessReviewAccessItemManualRemediationType) {
	o.ManualRemediationTypes = v
}

// GetEntitlementInfo returns the EntitlementInfo field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItem) GetEntitlementInfo() SecurityAccessReviewSubAccessItemEntitlementInfo {
	if o == nil || IsNil(o.EntitlementInfo) {
		var ret SecurityAccessReviewSubAccessItemEntitlementInfo
		return ret
	}
	return *o.EntitlementInfo
}

// GetEntitlementInfoOk returns a tuple with the EntitlementInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetEntitlementInfoOk() (*SecurityAccessReviewSubAccessItemEntitlementInfo, bool) {
	if o == nil || IsNil(o.EntitlementInfo) {
		return nil, false
	}
	return o.EntitlementInfo, true
}

// HasEntitlementInfo returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItem) HasEntitlementInfo() bool {
	if o != nil && !IsNil(o.EntitlementInfo) {
		return true
	}

	return false
}

// SetEntitlementInfo gets a reference to the given SecurityAccessReviewSubAccessItemEntitlementInfo and assigns it to the EntitlementInfo field.
func (o *SecurityAccessReviewSubAccessItem) SetEntitlementInfo(v SecurityAccessReviewSubAccessItemEntitlementInfo) {
	o.EntitlementInfo = &v
}

// GetGroupInfo returns the GroupInfo field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItem) GetGroupInfo() SecurityAccessReviewSubAccessItemGroupInfo {
	if o == nil || IsNil(o.GroupInfo) {
		var ret SecurityAccessReviewSubAccessItemGroupInfo
		return ret
	}
	return *o.GroupInfo
}

// GetGroupInfoOk returns a tuple with the GroupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItem) GetGroupInfoOk() (*SecurityAccessReviewSubAccessItemGroupInfo, bool) {
	if o == nil || IsNil(o.GroupInfo) {
		return nil, false
	}
	return o.GroupInfo, true
}

// HasGroupInfo returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItem) HasGroupInfo() bool {
	if o != nil && !IsNil(o.GroupInfo) {
		return true
	}

	return false
}

// SetGroupInfo gets a reference to the given SecurityAccessReviewSubAccessItemGroupInfo and assigns it to the GroupInfo field.
func (o *SecurityAccessReviewSubAccessItem) SetGroupInfo(v SecurityAccessReviewSubAccessItemGroupInfo) {
	o.GroupInfo = &v
}

func (o SecurityAccessReviewSubAccessItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewSubAccessItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["severity"] = o.Severity
	toSerialize["supportedActions"] = o.SupportedActions
	if !IsNil(o.RemediationStatus) {
		toSerialize["remediationStatus"] = o.RemediationStatus
	}
	if !IsNil(o.ManualRemediationTypes) {
		toSerialize["manualRemediationTypes"] = o.ManualRemediationTypes
	}
	if !IsNil(o.EntitlementInfo) {
		toSerialize["entitlementInfo"] = o.EntitlementInfo
	}
	if !IsNil(o.GroupInfo) {
		toSerialize["groupInfo"] = o.GroupInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewSubAccessItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"resourceId",
		"severity",
		"supportedActions",
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

	varSecurityAccessReviewSubAccessItem := _SecurityAccessReviewSubAccessItem{}

	err = json.Unmarshal(data, &varSecurityAccessReviewSubAccessItem)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewSubAccessItem(varSecurityAccessReviewSubAccessItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "supportedActions")
		delete(additionalProperties, "remediationStatus")
		delete(additionalProperties, "manualRemediationTypes")
		delete(additionalProperties, "entitlementInfo")
		delete(additionalProperties, "groupInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewSubAccessItem struct {
	value *SecurityAccessReviewSubAccessItem
	isSet bool
}

func (v NullableSecurityAccessReviewSubAccessItem) Get() *SecurityAccessReviewSubAccessItem {
	return v.value
}

func (v *NullableSecurityAccessReviewSubAccessItem) Set(val *SecurityAccessReviewSubAccessItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSubAccessItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSubAccessItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSubAccessItem(val *SecurityAccessReviewSubAccessItem) *NullableSecurityAccessReviewSubAccessItem {
	return &NullableSecurityAccessReviewSubAccessItem{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSubAccessItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSubAccessItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
