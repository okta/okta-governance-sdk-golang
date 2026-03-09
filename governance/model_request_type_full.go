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

// checks if the RequestTypeFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeFull{}

// RequestTypeFull Full representation of a request type resource
type RequestTypeFull struct {
	// The ID of the team that administers this request type.
	OwnerId          string                              `json:"ownerId" validate:"regexp=^[a-fA-F\\\\d]{24}$"`
	ResourceSettings RequestTypeResourceSettingsReadable `json:"resourceSettings"`
	RequestSettings  RequestTypeRequestSettingsReadable  `json:"requestSettings"`
	ApprovalSettings RequestTypeApprovalSettingsReadable `json:"approvalSettings"`
	// How long the requester retains access after their request is approved and fulfilled.  Specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).  #### Known limitation  Only single time unit ISO 8601 duration formats (D, H, M) are supported, for units (days, hours, minutes).  ##### Supported  | Unit       | Example | | ---------- | ------- | | D, days    | P40D    | | H, hours   | PT65H   | | M, minutes | PT90M   |  > **Note:** Mixes of units, as well as month/year/week designations, are not supported. For example, `P40DT65H`, `P40M`, `P1W` and `P1Y` are not supported.
	AccessDuration   NullableString              `json:"accessDuration"`
	Status           RequestTypeStatus           `json:"status"`
	LastUpdateSource RequestTypeLastUpdateSource `json:"lastUpdateSource"`
	Links            RequestTypeLinks            `json:"_links"`
	// Writable unique key on Create. Not modifiable on update.
	Name string `json:"name"`
	// Human readable description
	Description string `json:"description"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy        string `json:"lastUpdatedBy"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeFull RequestTypeFull

// NewRequestTypeFull instantiates a new RequestTypeFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeFull(ownerId string, resourceSettings RequestTypeResourceSettingsReadable, requestSettings RequestTypeRequestSettingsReadable, approvalSettings RequestTypeApprovalSettingsReadable, accessDuration NullableString, status RequestTypeStatus, lastUpdateSource RequestTypeLastUpdateSource, links RequestTypeLinks, name string, description string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *RequestTypeFull {
	this := RequestTypeFull{}
	this.Name = name
	this.Description = description
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	return &this
}

// NewRequestTypeFullWithDefaults instantiates a new RequestTypeFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeFullWithDefaults() *RequestTypeFull {
	this := RequestTypeFull{}
	return &this
}

// GetOwnerId returns the OwnerId field value
func (o *RequestTypeFull) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *RequestTypeFull) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetResourceSettings returns the ResourceSettings field value
func (o *RequestTypeFull) GetResourceSettings() RequestTypeResourceSettingsReadable {
	if o == nil {
		var ret RequestTypeResourceSettingsReadable
		return ret
	}

	return o.ResourceSettings
}

// GetResourceSettingsOk returns a tuple with the ResourceSettings field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetResourceSettingsOk() (*RequestTypeResourceSettingsReadable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceSettings, true
}

// SetResourceSettings sets field value
func (o *RequestTypeFull) SetResourceSettings(v RequestTypeResourceSettingsReadable) {
	o.ResourceSettings = v
}

// GetRequestSettings returns the RequestSettings field value
func (o *RequestTypeFull) GetRequestSettings() RequestTypeRequestSettingsReadable {
	if o == nil {
		var ret RequestTypeRequestSettingsReadable
		return ret
	}

	return o.RequestSettings
}

// GetRequestSettingsOk returns a tuple with the RequestSettings field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetRequestSettingsOk() (*RequestTypeRequestSettingsReadable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSettings, true
}

// SetRequestSettings sets field value
func (o *RequestTypeFull) SetRequestSettings(v RequestTypeRequestSettingsReadable) {
	o.RequestSettings = v
}

// GetApprovalSettings returns the ApprovalSettings field value
func (o *RequestTypeFull) GetApprovalSettings() RequestTypeApprovalSettingsReadable {
	if o == nil {
		var ret RequestTypeApprovalSettingsReadable
		return ret
	}

	return o.ApprovalSettings
}

// GetApprovalSettingsOk returns a tuple with the ApprovalSettings field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetApprovalSettingsOk() (*RequestTypeApprovalSettingsReadable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalSettings, true
}

// SetApprovalSettings sets field value
func (o *RequestTypeFull) SetApprovalSettings(v RequestTypeApprovalSettingsReadable) {
	o.ApprovalSettings = v
}

// GetAccessDuration returns the AccessDuration field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequestTypeFull) GetAccessDuration() string {
	if o == nil || o.AccessDuration.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessDuration.Get()
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestTypeFull) GetAccessDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessDuration.Get(), o.AccessDuration.IsSet()
}

// SetAccessDuration sets field value
func (o *RequestTypeFull) SetAccessDuration(v string) {
	o.AccessDuration.Set(&v)
}

// GetStatus returns the Status field value
func (o *RequestTypeFull) GetStatus() RequestTypeStatus {
	if o == nil {
		var ret RequestTypeStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetStatusOk() (*RequestTypeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestTypeFull) SetStatus(v RequestTypeStatus) {
	o.Status = v
}

// GetLastUpdateSource returns the LastUpdateSource field value
func (o *RequestTypeFull) GetLastUpdateSource() RequestTypeLastUpdateSource {
	if o == nil {
		var ret RequestTypeLastUpdateSource
		return ret
	}

	return o.LastUpdateSource
}

// GetLastUpdateSourceOk returns a tuple with the LastUpdateSource field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetLastUpdateSourceOk() (*RequestTypeLastUpdateSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdateSource, true
}

// SetLastUpdateSource sets field value
func (o *RequestTypeFull) SetLastUpdateSource(v RequestTypeLastUpdateSource) {
	o.LastUpdateSource = v
}

// GetLinks returns the Links field value
func (o *RequestTypeFull) GetLinks() RequestTypeLinks {
	if o == nil {
		var ret RequestTypeLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetLinksOk() (*RequestTypeLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestTypeFull) SetLinks(v RequestTypeLinks) {
	o.Links = v
}

// GetName returns the Name field value
func (o *RequestTypeFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestTypeFull) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RequestTypeFull) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RequestTypeFull) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *RequestTypeFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestTypeFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestTypeFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestTypeFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestTypeFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestTypeFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestTypeFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestTypeFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestTypeFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestTypeFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestTypeFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

func (o RequestTypeFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["resourceSettings"] = o.ResourceSettings
	toSerialize["requestSettings"] = o.RequestSettings
	toSerialize["approvalSettings"] = o.ApprovalSettings
	toSerialize["accessDuration"] = o.AccessDuration.Get()
	toSerialize["status"] = o.Status
	toSerialize["lastUpdateSource"] = o.LastUpdateSource
	toSerialize["_links"] = o.Links
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ownerId",
		"resourceSettings",
		"requestSettings",
		"approvalSettings",
		"accessDuration",
		"status",
		"lastUpdateSource",
		"_links",
		"name",
		"description",
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
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

	varRequestTypeFull := _RequestTypeFull{}

	err = json.Unmarshal(data, &varRequestTypeFull)

	if err != nil {
		return err
	}

	*o = RequestTypeFull(varRequestTypeFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "resourceSettings")
		delete(additionalProperties, "requestSettings")
		delete(additionalProperties, "approvalSettings")
		delete(additionalProperties, "accessDuration")
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastUpdateSource")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeFull struct {
	value *RequestTypeFull
	isSet bool
}

func (v NullableRequestTypeFull) Get() *RequestTypeFull {
	return v.value
}

func (v *NullableRequestTypeFull) Set(val *RequestTypeFull) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeFull) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeFull(val *RequestTypeFull) *NullableRequestTypeFull {
	return &NullableRequestTypeFull{value: val, isSet: true}
}

func (v NullableRequestTypeFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
