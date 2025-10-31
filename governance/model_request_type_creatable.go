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

// checks if the RequestTypeCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeCreatable{}

// RequestTypeCreatable The properties expected in an initial Add request
type RequestTypeCreatable struct {
	Status *RequestTypeCreatableStatus `json:"status,omitempty"`
	// The ID of the team that administers this request type.
	OwnerId string `json:"ownerId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	ResourceSettings RequestTypeResourceSettingsMutable `json:"resourceSettings"`
	RequestSettings *RequestTypeRequestSettingsMutable `json:"requestSettings,omitempty"`
	ApprovalSettings RequestTypeApprovalSettingsMutable `json:"approvalSettings"`
	// How long the requester retains access after their request is approved and fulfilled.  Specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).  #### Known limitation  Only single time unit ISO 8601 duration formats (D, H, M) are supported, for units (days, hours, minutes).  ##### Supported  | Unit       | Example | | ---------- | ------- | | D, days    | P40D    | | H, hours   | PT65H   | | M, minutes | PT90M   |  > **Note:** Mixes of units, as well as month/year/week designations, are not supported. For example, `P40DT65H`, `P40M`, `P1W` and `P1Y` are not supported.
	AccessDuration NullableString `json:"accessDuration,omitempty"`
	// Writable unique key on Create. Not modifiable on update.
	Name string `json:"name"`
	// Human readable description.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeCreatable RequestTypeCreatable

// NewRequestTypeCreatable instantiates a new RequestTypeCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeCreatable(ownerId string, resourceSettings RequestTypeResourceSettingsMutable, approvalSettings RequestTypeApprovalSettingsMutable, name string) *RequestTypeCreatable {
	this := RequestTypeCreatable{}
	this.Name = name
	return &this
}

// NewRequestTypeCreatableWithDefaults instantiates a new RequestTypeCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeCreatableWithDefaults() *RequestTypeCreatable {
	this := RequestTypeCreatable{}
	var status RequestTypeCreatableStatus = REQUESTTYPECREATABLESTATUS_DRAFT
	this.Status = &status
	var requestSettings RequestTypeRequestSettingsMutable = {"type":"EVERYONE","requesterFields":[]}
	this.RequestSettings = &requestSettings
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestTypeCreatable) GetStatus() RequestTypeCreatableStatus {
	if o == nil || IsNil(o.Status) {
		var ret RequestTypeCreatableStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeCreatable) GetStatusOk() (*RequestTypeCreatableStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestTypeCreatable) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RequestTypeCreatableStatus and assigns it to the Status field.
func (o *RequestTypeCreatable) SetStatus(v RequestTypeCreatableStatus) {
	o.Status = &v
}

// GetOwnerId returns the OwnerId field value
func (o *RequestTypeCreatable) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *RequestTypeCreatable) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *RequestTypeCreatable) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetResourceSettings returns the ResourceSettings field value
func (o *RequestTypeCreatable) GetResourceSettings() RequestTypeResourceSettingsMutable {
	if o == nil {
		var ret RequestTypeResourceSettingsMutable
		return ret
	}

	return o.ResourceSettings
}

// GetResourceSettingsOk returns a tuple with the ResourceSettings field value
// and a boolean to check if the value has been set.
func (o *RequestTypeCreatable) GetResourceSettingsOk() (*RequestTypeResourceSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSettings, true
}

// SetResourceSettings sets field value
func (o *RequestTypeCreatable) SetResourceSettings(v RequestTypeResourceSettingsMutable) {
	o.ResourceSettings = v
}

// GetRequestSettings returns the RequestSettings field value if set, zero value otherwise.
func (o *RequestTypeCreatable) GetRequestSettings() RequestTypeRequestSettingsMutable {
	if o == nil || IsNil(o.RequestSettings) {
		var ret RequestTypeRequestSettingsMutable
		return ret
	}
	return *o.RequestSettings
}

// GetRequestSettingsOk returns a tuple with the RequestSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeCreatable) GetRequestSettingsOk() (*RequestTypeRequestSettingsMutable, bool) {
	if o == nil || IsNil(o.RequestSettings) {
		return nil, false
	}
	return o.RequestSettings, true
}

// HasRequestSettings returns a boolean if a field has been set.
func (o *RequestTypeCreatable) HasRequestSettings() bool {
	if o != nil && !IsNil(o.RequestSettings) {
		return true
	}

	return false
}

// SetRequestSettings gets a reference to the given RequestTypeRequestSettingsMutable and assigns it to the RequestSettings field.
func (o *RequestTypeCreatable) SetRequestSettings(v RequestTypeRequestSettingsMutable) {
	o.RequestSettings = &v
}

// GetApprovalSettings returns the ApprovalSettings field value
func (o *RequestTypeCreatable) GetApprovalSettings() RequestTypeApprovalSettingsMutable {
	if o == nil {
		var ret RequestTypeApprovalSettingsMutable
		return ret
	}

	return o.ApprovalSettings
}

// GetApprovalSettingsOk returns a tuple with the ApprovalSettings field value
// and a boolean to check if the value has been set.
func (o *RequestTypeCreatable) GetApprovalSettingsOk() (*RequestTypeApprovalSettingsMutable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalSettings, true
}

// SetApprovalSettings sets field value
func (o *RequestTypeCreatable) SetApprovalSettings(v RequestTypeApprovalSettingsMutable) {
	o.ApprovalSettings = v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestTypeCreatable) GetAccessDuration() string {
	if o == nil || IsNil(o.AccessDuration.Get()) {
		var ret string
		return ret
	}
	return *o.AccessDuration.Get()
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestTypeCreatable) GetAccessDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessDuration.Get(), o.AccessDuration.IsSet()
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *RequestTypeCreatable) HasAccessDuration() bool {
	if o != nil && o.AccessDuration.IsSet() {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given NullableString and assigns it to the AccessDuration field.
func (o *RequestTypeCreatable) SetAccessDuration(v string) {
	o.AccessDuration.Set(&v)
}
// SetAccessDurationNil sets the value for AccessDuration to be an explicit nil
func (o *RequestTypeCreatable) SetAccessDurationNil() {
	o.AccessDuration.Set(nil)
}

// UnsetAccessDuration ensures that no value is present for AccessDuration, not even an explicit nil
func (o *RequestTypeCreatable) UnsetAccessDuration() {
	o.AccessDuration.Unset()
}

// GetName returns the Name field value
func (o *RequestTypeCreatable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestTypeCreatable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestTypeCreatable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestTypeCreatable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTypeCreatable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestTypeCreatable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestTypeCreatable) SetDescription(v string) {
	o.Description = &v
}

func (o RequestTypeCreatable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["resourceSettings"] = o.ResourceSettings
	if !IsNil(o.RequestSettings) {
		toSerialize["requestSettings"] = o.RequestSettings
	}
	toSerialize["approvalSettings"] = o.ApprovalSettings
	if o.AccessDuration.IsSet() {
		toSerialize["accessDuration"] = o.AccessDuration.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ownerId",
		"resourceSettings",
		"approvalSettings",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRequestTypeCreatable := _RequestTypeCreatable{}

	err = json.Unmarshal(data, &varRequestTypeCreatable)

	if err != nil {
		return err
	}

	*o = RequestTypeCreatable(varRequestTypeCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "resourceSettings")
		delete(additionalProperties, "requestSettings")
		delete(additionalProperties, "approvalSettings")
		delete(additionalProperties, "accessDuration")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeCreatable struct {
	value *RequestTypeCreatable
	isSet bool
}

func (v NullableRequestTypeCreatable) Get() *RequestTypeCreatable {
	return v.value
}

func (v *NullableRequestTypeCreatable) Set(val *RequestTypeCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeCreatable(val *RequestTypeCreatable) *NullableRequestTypeCreatable {
	return &NullableRequestTypeCreatable{value: val, isSet: true}
}

func (v NullableRequestTypeCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


