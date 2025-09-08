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

// GrantFullWithEntitlements Grant response with entitlements
type GrantFullWithEntitlements struct {
	GrantType GrantType `json:"grantType"`
	// The entitlement bundle `id`
	EntitlementBundleId *string `json:"entitlementBundleId,omitempty"`
	// The Okta user `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetPrincipalOrn string              `json:"targetPrincipalOrn"`
	TargetPrincipal    TargetPrincipalFull `json:"targetPrincipal"`
	Action             GrantAction         `json:"action"`
	Actor              GrantActor          `json:"actor"`
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetResourceOrn string         `json:"targetResourceOrn"`
	Target            TargetResource `json:"target"`
	// Collection of entitlements with associated values
	Entitlements     []EntitlementFull          `json:"entitlements"`
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

type _GrantFullWithEntitlements GrantFullWithEntitlements

// NewGrantFullWithEntitlements instantiates a new GrantFullWithEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantFullWithEntitlements(grantType GrantType, targetPrincipalOrn string, targetPrincipal TargetPrincipalFull, action GrantAction, actor GrantActor, targetResourceOrn string, target TargetResource, entitlements []EntitlementFull, links ResourceGrantLinks, status GrantStatus, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *GrantFullWithEntitlements {
	this := GrantFullWithEntitlements{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	return &this
}

// NewGrantFullWithEntitlementsWithDefaults instantiates a new GrantFullWithEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantFullWithEntitlementsWithDefaults() *GrantFullWithEntitlements {
	this := GrantFullWithEntitlements{}
	var action GrantAction = GRANTACTION_ALLOW
	this.Action = action
	var actor GrantActor = GRANTACTOR_API
	this.Actor = actor
	return &this
}

// GetGrantType returns the GrantType field value
func (o *GrantFullWithEntitlements) GetGrantType() GrantType {
	if o == nil {
		var ret GrantType
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *GrantFullWithEntitlements) SetGrantType(v GrantType) {
	o.GrantType = v
}

// GetEntitlementBundleId returns the EntitlementBundleId field value if set, zero value otherwise.
func (o *GrantFullWithEntitlements) GetEntitlementBundleId() string {
	if o == nil || o.EntitlementBundleId == nil {
		var ret string
		return ret
	}
	return *o.EntitlementBundleId
}

// GetEntitlementBundleIdOk returns a tuple with the EntitlementBundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetEntitlementBundleIdOk() (*string, bool) {
	if o == nil || o.EntitlementBundleId == nil {
		return nil, false
	}
	return o.EntitlementBundleId, true
}

// HasEntitlementBundleId returns a boolean if a field has been set.
func (o *GrantFullWithEntitlements) HasEntitlementBundleId() bool {
	if o != nil && o.EntitlementBundleId != nil {
		return true
	}

	return false
}

// SetEntitlementBundleId gets a reference to the given string and assigns it to the EntitlementBundleId field.
func (o *GrantFullWithEntitlements) SetEntitlementBundleId(v string) {
	o.EntitlementBundleId = &v
}

// GetTargetPrincipalOrn returns the TargetPrincipalOrn field value
func (o *GrantFullWithEntitlements) GetTargetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPrincipalOrn
}

// GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetTargetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipalOrn, true
}

// SetTargetPrincipalOrn sets field value
func (o *GrantFullWithEntitlements) SetTargetPrincipalOrn(v string) {
	o.TargetPrincipalOrn = v
}

// GetTargetPrincipal returns the TargetPrincipal field value
func (o *GrantFullWithEntitlements) GetTargetPrincipal() TargetPrincipalFull {
	if o == nil {
		var ret TargetPrincipalFull
		return ret
	}

	return o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetTargetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipal, true
}

// SetTargetPrincipal sets field value
func (o *GrantFullWithEntitlements) SetTargetPrincipal(v TargetPrincipalFull) {
	o.TargetPrincipal = v
}

// GetAction returns the Action field value
func (o *GrantFullWithEntitlements) GetAction() GrantAction {
	if o == nil {
		var ret GrantAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetActionOk() (*GrantAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *GrantFullWithEntitlements) SetAction(v GrantAction) {
	o.Action = v
}

// GetActor returns the Actor field value
func (o *GrantFullWithEntitlements) GetActor() GrantActor {
	if o == nil {
		var ret GrantActor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetActorOk() (*GrantActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *GrantFullWithEntitlements) SetActor(v GrantActor) {
	o.Actor = v
}

// GetTargetResourceOrn returns the TargetResourceOrn field value
func (o *GrantFullWithEntitlements) GetTargetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetResourceOrn
}

// GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetTargetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetResourceOrn, true
}

// SetTargetResourceOrn sets field value
func (o *GrantFullWithEntitlements) SetTargetResourceOrn(v string) {
	o.TargetResourceOrn = v
}

// GetTarget returns the Target field value
func (o *GrantFullWithEntitlements) GetTarget() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetTargetOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *GrantFullWithEntitlements) SetTarget(v TargetResource) {
	o.Target = v
}

// GetEntitlements returns the Entitlements field value
func (o *GrantFullWithEntitlements) GetEntitlements() []EntitlementFull {
	if o == nil {
		var ret []EntitlementFull
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetEntitlementsOk() ([]EntitlementFull, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *GrantFullWithEntitlements) SetEntitlements(v []EntitlementFull) {
	o.Entitlements = v
}

// GetScheduleSettings returns the ScheduleSettings field value if set, zero value otherwise.
func (o *GrantFullWithEntitlements) GetScheduleSettings() ScheduleSettingsWriteable {
	if o == nil || o.ScheduleSettings == nil {
		var ret ScheduleSettingsWriteable
		return ret
	}
	return *o.ScheduleSettings
}

// GetScheduleSettingsOk returns a tuple with the ScheduleSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool) {
	if o == nil || o.ScheduleSettings == nil {
		return nil, false
	}
	return o.ScheduleSettings, true
}

// HasScheduleSettings returns a boolean if a field has been set.
func (o *GrantFullWithEntitlements) HasScheduleSettings() bool {
	if o != nil && o.ScheduleSettings != nil {
		return true
	}

	return false
}

// SetScheduleSettings gets a reference to the given ScheduleSettingsWriteable and assigns it to the ScheduleSettings field.
func (o *GrantFullWithEntitlements) SetScheduleSettings(v ScheduleSettingsWriteable) {
	o.ScheduleSettings = &v
}

// GetLinks returns the Links field value
func (o *GrantFullWithEntitlements) GetLinks() ResourceGrantLinks {
	if o == nil {
		var ret ResourceGrantLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetLinksOk() (*ResourceGrantLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GrantFullWithEntitlements) SetLinks(v ResourceGrantLinks) {
	o.Links = v
}

// GetStatus returns the Status field value
func (o *GrantFullWithEntitlements) GetStatus() GrantStatus {
	if o == nil {
		var ret GrantStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetStatusOk() (*GrantStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GrantFullWithEntitlements) SetStatus(v GrantStatus) {
	o.Status = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GrantFullWithEntitlements) GetMetadata() GrantMetadata {
	if o == nil || o.Metadata == nil {
		var ret GrantMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetMetadataOk() (*GrantMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GrantFullWithEntitlements) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given GrantMetadata and assigns it to the Metadata field.
func (o *GrantFullWithEntitlements) SetMetadata(v GrantMetadata) {
	o.Metadata = &v
}

// GetId returns the Id field value
func (o *GrantFullWithEntitlements) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GrantFullWithEntitlements) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *GrantFullWithEntitlements) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *GrantFullWithEntitlements) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *GrantFullWithEntitlements) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GrantFullWithEntitlements) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *GrantFullWithEntitlements) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *GrantFullWithEntitlements) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *GrantFullWithEntitlements) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *GrantFullWithEntitlements) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *GrantFullWithEntitlements) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

func (o GrantFullWithEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["grantType"] = o.GrantType
	}
	if o.EntitlementBundleId != nil {
		toSerialize["entitlementBundleId"] = o.EntitlementBundleId
	}
	if true {
		toSerialize["targetPrincipalOrn"] = o.TargetPrincipalOrn
	}
	if true {
		toSerialize["targetPrincipal"] = o.TargetPrincipal
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	if true {
		toSerialize["targetResourceOrn"] = o.TargetResourceOrn
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["entitlements"] = o.Entitlements
	}
	if o.ScheduleSettings != nil {
		toSerialize["scheduleSettings"] = o.ScheduleSettings
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GrantFullWithEntitlements) UnmarshalJSON(bytes []byte) (err error) {
	varGrantFullWithEntitlements := _GrantFullWithEntitlements{}

	err = json.Unmarshal(bytes, &varGrantFullWithEntitlements)
	if err == nil {
		*o = GrantFullWithEntitlements(varGrantFullWithEntitlements)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
	} else {
		return err
	}

	return err
}

type NullableGrantFullWithEntitlements struct {
	value *GrantFullWithEntitlements
	isSet bool
}

func (v NullableGrantFullWithEntitlements) Get() *GrantFullWithEntitlements {
	return v.value
}

func (v *NullableGrantFullWithEntitlements) Set(val *GrantFullWithEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantFullWithEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantFullWithEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantFullWithEntitlements(val *GrantFullWithEntitlements) *NullableGrantFullWithEntitlements {
	return &NullableGrantFullWithEntitlements{value: val, isSet: true}
}

func (v NullableGrantFullWithEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantFullWithEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
