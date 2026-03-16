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

// checks if the GrantFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantFull{}

// GrantFull Grant response
type GrantFull struct {
	GrantType GrantType `json:"grantType"`
	// The entitlement bundle `id`
	EntitlementBundleId *string `json:"entitlementBundleId,omitempty" validate:"regexp=enb[0-9a-zA-Z]+"`
	// The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.
	TargetPrincipalOrn string              `json:"targetPrincipalOrn"`
	TargetPrincipal    TargetPrincipalFull `json:"targetPrincipal"`
	Action             GrantAction         `json:"action"`
	Actor              GrantActor          `json:"actor"`
	// The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetResourceOrn string         `json:"targetResourceOrn"`
	Target            TargetResource `json:"target"`
	// Collection of entitlements and associated value identifiers
	Entitlements     []EntitlementCreatable     `json:"entitlements,omitempty"`
	ScheduleSettings *ScheduleSettingsWriteable `json:"scheduleSettings,omitempty"`
	Links            ResourceGrantLinks         `json:"_links"`
	Status           GrantStatus                `json:"status"`
	Metadata         *GrantMetadata             `json:"metadata,omitempty"`
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

type _GrantFull GrantFull

// NewGrantFull instantiates a new GrantFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantFull(grantType GrantType, targetPrincipalOrn string, targetPrincipal TargetPrincipalFull, action GrantAction, actor GrantActor, targetResourceOrn string, target TargetResource, links ResourceGrantLinks, status GrantStatus, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *GrantFull {
	this := GrantFull{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	return &this
}

// NewGrantFullWithDefaults instantiates a new GrantFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantFullWithDefaults() *GrantFull {
	this := GrantFull{}
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = actor
	return &this
}

// GetGrantType returns the GrantType field value
func (o *GrantFull) GetGrantType() GrantType {
	if o == nil {
		var ret GrantType
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *GrantFull) SetGrantType(v GrantType) {
	o.GrantType = v
}

// GetEntitlementBundleId returns the EntitlementBundleId field value if set, zero value otherwise.
func (o *GrantFull) GetEntitlementBundleId() string {
	if o == nil || IsNil(o.EntitlementBundleId) {
		var ret string
		return ret
	}
	return *o.EntitlementBundleId
}

// GetEntitlementBundleIdOk returns a tuple with the EntitlementBundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantFull) GetEntitlementBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementBundleId) {
		return nil, false
	}
	return o.EntitlementBundleId, true
}

// HasEntitlementBundleId returns a boolean if a field has been set.
func (o *GrantFull) HasEntitlementBundleId() bool {
	if o != nil && !IsNil(o.EntitlementBundleId) {
		return true
	}

	return false
}

// SetEntitlementBundleId gets a reference to the given string and assigns it to the EntitlementBundleId field.
func (o *GrantFull) SetEntitlementBundleId(v string) {
	o.EntitlementBundleId = &v
}

// GetTargetPrincipalOrn returns the TargetPrincipalOrn field value
func (o *GrantFull) GetTargetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPrincipalOrn
}

// GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetTargetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipalOrn, true
}

// SetTargetPrincipalOrn sets field value
func (o *GrantFull) SetTargetPrincipalOrn(v string) {
	o.TargetPrincipalOrn = v
}

// GetTargetPrincipal returns the TargetPrincipal field value
func (o *GrantFull) GetTargetPrincipal() TargetPrincipalFull {
	if o == nil {
		var ret TargetPrincipalFull
		return ret
	}

	return o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetTargetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipal, true
}

// SetTargetPrincipal sets field value
func (o *GrantFull) SetTargetPrincipal(v TargetPrincipalFull) {
	o.TargetPrincipal = v
}

// GetAction returns the Action field value
func (o *GrantFull) GetAction() GrantAction {
	if o == nil {
		var ret GrantAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetActionOk() (*GrantAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *GrantFull) SetAction(v GrantAction) {
	o.Action = v
}

// GetActor returns the Actor field value
func (o *GrantFull) GetActor() GrantActor {
	if o == nil {
		var ret GrantActor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetActorOk() (*GrantActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *GrantFull) SetActor(v GrantActor) {
	o.Actor = v
}

// GetTargetResourceOrn returns the TargetResourceOrn field value
func (o *GrantFull) GetTargetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetResourceOrn
}

// GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetTargetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetResourceOrn, true
}

// SetTargetResourceOrn sets field value
func (o *GrantFull) SetTargetResourceOrn(v string) {
	o.TargetResourceOrn = v
}

// GetTarget returns the Target field value
func (o *GrantFull) GetTarget() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetTargetOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *GrantFull) SetTarget(v TargetResource) {
	o.Target = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *GrantFull) GetEntitlements() []EntitlementCreatable {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantFull) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *GrantFull) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *GrantFull) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetScheduleSettings returns the ScheduleSettings field value if set, zero value otherwise.
func (o *GrantFull) GetScheduleSettings() ScheduleSettingsWriteable {
	if o == nil || IsNil(o.ScheduleSettings) {
		var ret ScheduleSettingsWriteable
		return ret
	}
	return *o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantFull) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool) {
	if o == nil || IsNil(o.ScheduleSettings) {
		return nil, false
	}
	return o.ScheduleSettings, true
}

// HasScheduleSettings returns a boolean if a field has been set.
func (o *GrantFull) HasScheduleSettings() bool {
	if o != nil && !IsNil(o.ScheduleSettings) {
		return true
	}

	return false
}

// SetScheduleSettings gets a reference to the given ScheduleSettingsWriteable and assigns it to the ScheduleSettings field.
func (o *GrantFull) SetScheduleSettings(v ScheduleSettingsWriteable) {
	o.ScheduleSettings = &v
}

// GetLinks returns the Links field value
func (o *GrantFull) GetLinks() ResourceGrantLinks {
	if o == nil {
		var ret ResourceGrantLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetLinksOk() (*ResourceGrantLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GrantFull) SetLinks(v ResourceGrantLinks) {
	o.Links = v
}

// GetStatus returns the Status field value
func (o *GrantFull) GetStatus() GrantStatus {
	if o == nil {
		var ret GrantStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetStatusOk() (*GrantStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GrantFull) SetStatus(v GrantStatus) {
	o.Status = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GrantFull) GetMetadata() GrantMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret GrantMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantFull) GetMetadataOk() (*GrantMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GrantFull) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given GrantMetadata and assigns it to the Metadata field.
func (o *GrantFull) SetMetadata(v GrantMetadata) {
	o.Metadata = &v
}

// GetId returns the Id field value
func (o *GrantFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GrantFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *GrantFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *GrantFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *GrantFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GrantFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *GrantFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *GrantFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *GrantFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *GrantFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *GrantFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

func (o GrantFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grantType"] = o.GrantType
	if !IsNil(o.EntitlementBundleId) {
		toSerialize["entitlementBundleId"] = o.EntitlementBundleId
	}
	toSerialize["targetPrincipalOrn"] = o.TargetPrincipalOrn
	toSerialize["targetPrincipal"] = o.TargetPrincipal
	toSerialize["action"] = o.Action
	toSerialize["actor"] = o.Actor
	toSerialize["targetResourceOrn"] = o.TargetResourceOrn
	toSerialize["target"] = o.Target
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.ScheduleSettings) {
		toSerialize["scheduleSettings"] = o.ScheduleSettings
	}
	toSerialize["_links"] = o.Links
	toSerialize["status"] = o.Status
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
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

func (o *GrantFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grantType",
		"targetPrincipalOrn",
		"targetPrincipal",
		"action",
		"actor",
		"targetResourceOrn",
		"target",
		"_links",
		"status",
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

	varGrantFull := _GrantFull{}

	err = json.Unmarshal(data, &varGrantFull)

	if err != nil {
		return err
	}

	*o = GrantFull(varGrantFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "entitlementBundleId")
		delete(additionalProperties, "targetPrincipalOrn")
		delete(additionalProperties, "targetPrincipal")
		delete(additionalProperties, "action")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "targetResourceOrn")
		delete(additionalProperties, "target")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "scheduleSettings")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "status")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantFull struct {
	value *GrantFull
	isSet bool
}

func (v NullableGrantFull) Get() *GrantFull {
	return v.value
}

func (v *NullableGrantFull) Set(val *GrantFull) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantFull) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantFull(val *GrantFull) *NullableGrantFull {
	return &NullableGrantFull{value: val, isSet: true}
}

func (v NullableGrantFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
