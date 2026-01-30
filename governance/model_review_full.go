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

// checks if the ReviewFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewFull{}

// ReviewFull Full representation of a Review resource
type ReviewFull struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy        string                     `json:"lastUpdatedBy"`
	Links                ReviewLinks                `json:"_links"`
	CampaignId           string                     `json:"campaignId"`
	ResourceId           string                     `json:"resourceId"`
	EntitlementValue     *ReviewerEntitlementValue  `json:"entitlementValue,omitempty"`
	EntitlementBundle    *ReviewerEntitlementBundle `json:"entitlementBundle,omitempty"`
	Decision             Decision                   `json:"decision"`
	Decided              *time.Time                 `json:"decided,omitempty"`
	RemediationStatus    RemediationStatus          `json:"remediationStatus"`
	PrincipalProfile     PrincipalProfile           `json:"principalProfile"`
	ReviewerProfile      *PrincipalProfile          `json:"reviewerProfile,omitempty"`
	ReviewerType         ReviewersReviewerType      `json:"reviewerType"`
	ReviewerGroupProfile *ReviewerGroupProfile      `json:"reviewerGroupProfile,omitempty"`
	CurrentReviewerLevel *ReviewerLevelType         `json:"currentReviewerLevel,omitempty"`
	// List of risk rule conflicts caused by this entitlement value. Only applies to review item that has entitlement value.
	RiskRuleConflicts []RiskRuleConflicts `json:"riskRuleConflicts,omitempty"`
	Note              *Note               `json:"note,omitempty"`
	// Applicable only for multi level campaign. Provides details about the reviewer and decisions (if any) made at each reviewer level is captured here.
	AllReviewerLevels    []ReviewerLevelInfoFull `json:"allReviewerLevels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewFull ReviewFull

// NewReviewFull instantiates a new ReviewFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewFull(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links ReviewLinks, campaignId string, resourceId string, decision Decision, remediationStatus RemediationStatus, principalProfile PrincipalProfile, reviewerType ReviewersReviewerType) *ReviewFull {
	this := ReviewFull{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.CampaignId = campaignId
	this.ResourceId = resourceId
	this.Decision = decision
	this.RemediationStatus = remediationStatus
	this.PrincipalProfile = principalProfile
	this.ReviewerType = reviewerType
	return &this
}

// NewReviewFullWithDefaults instantiates a new ReviewFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewFullWithDefaults() *ReviewFull {
	this := ReviewFull{}
	return &this
}

// GetId returns the Id field value
func (o *ReviewFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ReviewFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ReviewFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *ReviewFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ReviewFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *ReviewFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *ReviewFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *ReviewFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *ReviewFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *ReviewFull) GetLinks() ReviewLinks {
	if o == nil {
		var ret ReviewLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetLinksOk() (*ReviewLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReviewFull) SetLinks(v ReviewLinks) {
	o.Links = v
}

// GetCampaignId returns the CampaignId field value
func (o *ReviewFull) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *ReviewFull) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetResourceId returns the ResourceId field value
func (o *ReviewFull) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ReviewFull) SetResourceId(v string) {
	o.ResourceId = v
}

// GetEntitlementValue returns the EntitlementValue field value if set, zero value otherwise.
func (o *ReviewFull) GetEntitlementValue() ReviewerEntitlementValue {
	if o == nil || IsNil(o.EntitlementValue) {
		var ret ReviewerEntitlementValue
		return ret
	}
	return *o.EntitlementValue
}

// GetEntitlementValueOk returns a tuple with the EntitlementValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetEntitlementValueOk() (*ReviewerEntitlementValue, bool) {
	if o == nil || IsNil(o.EntitlementValue) {
		return nil, false
	}
	return o.EntitlementValue, true
}

// HasEntitlementValue returns a boolean if a field has been set.
func (o *ReviewFull) HasEntitlementValue() bool {
	if o != nil && !IsNil(o.EntitlementValue) {
		return true
	}

	return false
}

// SetEntitlementValue gets a reference to the given ReviewerEntitlementValue and assigns it to the EntitlementValue field.
func (o *ReviewFull) SetEntitlementValue(v ReviewerEntitlementValue) {
	o.EntitlementValue = &v
}

// GetEntitlementBundle returns the EntitlementBundle field value if set, zero value otherwise.
func (o *ReviewFull) GetEntitlementBundle() ReviewerEntitlementBundle {
	if o == nil || IsNil(o.EntitlementBundle) {
		var ret ReviewerEntitlementBundle
		return ret
	}
	return *o.EntitlementBundle
}

// GetEntitlementBundleOk returns a tuple with the EntitlementBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetEntitlementBundleOk() (*ReviewerEntitlementBundle, bool) {
	if o == nil || IsNil(o.EntitlementBundle) {
		return nil, false
	}
	return o.EntitlementBundle, true
}

// HasEntitlementBundle returns a boolean if a field has been set.
func (o *ReviewFull) HasEntitlementBundle() bool {
	if o != nil && !IsNil(o.EntitlementBundle) {
		return true
	}

	return false
}

// SetEntitlementBundle gets a reference to the given ReviewerEntitlementBundle and assigns it to the EntitlementBundle field.
func (o *ReviewFull) SetEntitlementBundle(v ReviewerEntitlementBundle) {
	o.EntitlementBundle = &v
}

// GetDecision returns the Decision field value
func (o *ReviewFull) GetDecision() Decision {
	if o == nil {
		var ret Decision
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetDecisionOk() (*Decision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *ReviewFull) SetDecision(v Decision) {
	o.Decision = v
}

// GetDecided returns the Decided field value if set, zero value otherwise.
func (o *ReviewFull) GetDecided() time.Time {
	if o == nil || IsNil(o.Decided) {
		var ret time.Time
		return ret
	}
	return *o.Decided
}

// GetDecidedOk returns a tuple with the Decided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetDecidedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Decided) {
		return nil, false
	}
	return o.Decided, true
}

// HasDecided returns a boolean if a field has been set.
func (o *ReviewFull) HasDecided() bool {
	if o != nil && !IsNil(o.Decided) {
		return true
	}

	return false
}

// SetDecided gets a reference to the given time.Time and assigns it to the Decided field.
func (o *ReviewFull) SetDecided(v time.Time) {
	o.Decided = &v
}

// GetRemediationStatus returns the RemediationStatus field value
func (o *ReviewFull) GetRemediationStatus() RemediationStatus {
	if o == nil {
		var ret RemediationStatus
		return ret
	}

	return o.RemediationStatus
}

// GetRemediationStatusOk returns a tuple with the RemediationStatus field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetRemediationStatusOk() (*RemediationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationStatus, true
}

// SetRemediationStatus sets field value
func (o *ReviewFull) SetRemediationStatus(v RemediationStatus) {
	o.RemediationStatus = v
}

// GetPrincipalProfile returns the PrincipalProfile field value
func (o *ReviewFull) GetPrincipalProfile() PrincipalProfile {
	if o == nil {
		var ret PrincipalProfile
		return ret
	}

	return o.PrincipalProfile
}

// GetPrincipalProfileOk returns a tuple with the PrincipalProfile field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetPrincipalProfileOk() (*PrincipalProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalProfile, true
}

// SetPrincipalProfile sets field value
func (o *ReviewFull) SetPrincipalProfile(v PrincipalProfile) {
	o.PrincipalProfile = v
}

// GetReviewerProfile returns the ReviewerProfile field value if set, zero value otherwise.
func (o *ReviewFull) GetReviewerProfile() PrincipalProfile {
	if o == nil || IsNil(o.ReviewerProfile) {
		var ret PrincipalProfile
		return ret
	}
	return *o.ReviewerProfile
}

// GetReviewerProfileOk returns a tuple with the ReviewerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetReviewerProfileOk() (*PrincipalProfile, bool) {
	if o == nil || IsNil(o.ReviewerProfile) {
		return nil, false
	}
	return o.ReviewerProfile, true
}

// HasReviewerProfile returns a boolean if a field has been set.
func (o *ReviewFull) HasReviewerProfile() bool {
	if o != nil && !IsNil(o.ReviewerProfile) {
		return true
	}

	return false
}

// SetReviewerProfile gets a reference to the given PrincipalProfile and assigns it to the ReviewerProfile field.
func (o *ReviewFull) SetReviewerProfile(v PrincipalProfile) {
	o.ReviewerProfile = &v
}

// GetReviewerType returns the ReviewerType field value
func (o *ReviewFull) GetReviewerType() ReviewersReviewerType {
	if o == nil {
		var ret ReviewersReviewerType
		return ret
	}

	return o.ReviewerType
}

// GetReviewerTypeOk returns a tuple with the ReviewerType field value
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetReviewerTypeOk() (*ReviewersReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerType, true
}

// SetReviewerType sets field value
func (o *ReviewFull) SetReviewerType(v ReviewersReviewerType) {
	o.ReviewerType = v
}

// GetReviewerGroupProfile returns the ReviewerGroupProfile field value if set, zero value otherwise.
func (o *ReviewFull) GetReviewerGroupProfile() ReviewerGroupProfile {
	if o == nil || IsNil(o.ReviewerGroupProfile) {
		var ret ReviewerGroupProfile
		return ret
	}
	return *o.ReviewerGroupProfile
}

// GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool) {
	if o == nil || IsNil(o.ReviewerGroupProfile) {
		return nil, false
	}
	return o.ReviewerGroupProfile, true
}

// HasReviewerGroupProfile returns a boolean if a field has been set.
func (o *ReviewFull) HasReviewerGroupProfile() bool {
	if o != nil && !IsNil(o.ReviewerGroupProfile) {
		return true
	}

	return false
}

// SetReviewerGroupProfile gets a reference to the given ReviewerGroupProfile and assigns it to the ReviewerGroupProfile field.
func (o *ReviewFull) SetReviewerGroupProfile(v ReviewerGroupProfile) {
	o.ReviewerGroupProfile = &v
}

// GetCurrentReviewerLevel returns the CurrentReviewerLevel field value if set, zero value otherwise.
func (o *ReviewFull) GetCurrentReviewerLevel() ReviewerLevelType {
	if o == nil || IsNil(o.CurrentReviewerLevel) {
		var ret ReviewerLevelType
		return ret
	}
	return *o.CurrentReviewerLevel
}

// GetCurrentReviewerLevelOk returns a tuple with the CurrentReviewerLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetCurrentReviewerLevelOk() (*ReviewerLevelType, bool) {
	if o == nil || IsNil(o.CurrentReviewerLevel) {
		return nil, false
	}
	return o.CurrentReviewerLevel, true
}

// HasCurrentReviewerLevel returns a boolean if a field has been set.
func (o *ReviewFull) HasCurrentReviewerLevel() bool {
	if o != nil && !IsNil(o.CurrentReviewerLevel) {
		return true
	}

	return false
}

// SetCurrentReviewerLevel gets a reference to the given ReviewerLevelType and assigns it to the CurrentReviewerLevel field.
func (o *ReviewFull) SetCurrentReviewerLevel(v ReviewerLevelType) {
	o.CurrentReviewerLevel = &v
}

// GetRiskRuleConflicts returns the RiskRuleConflicts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewFull) GetRiskRuleConflicts() []RiskRuleConflicts {
	if o == nil {
		var ret []RiskRuleConflicts
		return ret
	}
	return o.RiskRuleConflicts
}

// GetRiskRuleConflictsOk returns a tuple with the RiskRuleConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewFull) GetRiskRuleConflictsOk() ([]RiskRuleConflicts, bool) {
	if o == nil || IsNil(o.RiskRuleConflicts) {
		return nil, false
	}
	return o.RiskRuleConflicts, true
}

// HasRiskRuleConflicts returns a boolean if a field has been set.
func (o *ReviewFull) HasRiskRuleConflicts() bool {
	if o != nil && !IsNil(o.RiskRuleConflicts) {
		return true
	}

	return false
}

// SetRiskRuleConflicts gets a reference to the given []RiskRuleConflicts and assigns it to the RiskRuleConflicts field.
func (o *ReviewFull) SetRiskRuleConflicts(v []RiskRuleConflicts) {
	o.RiskRuleConflicts = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ReviewFull) GetNote() Note {
	if o == nil || IsNil(o.Note) {
		var ret Note
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetNoteOk() (*Note, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ReviewFull) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given Note and assigns it to the Note field.
func (o *ReviewFull) SetNote(v Note) {
	o.Note = &v
}

// GetAllReviewerLevels returns the AllReviewerLevels field value if set, zero value otherwise.
func (o *ReviewFull) GetAllReviewerLevels() []ReviewerLevelInfoFull {
	if o == nil || IsNil(o.AllReviewerLevels) {
		var ret []ReviewerLevelInfoFull
		return ret
	}
	return o.AllReviewerLevels
}

// GetAllReviewerLevelsOk returns a tuple with the AllReviewerLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewFull) GetAllReviewerLevelsOk() ([]ReviewerLevelInfoFull, bool) {
	if o == nil || IsNil(o.AllReviewerLevels) {
		return nil, false
	}
	return o.AllReviewerLevels, true
}

// HasAllReviewerLevels returns a boolean if a field has been set.
func (o *ReviewFull) HasAllReviewerLevels() bool {
	if o != nil && !IsNil(o.AllReviewerLevels) {
		return true
	}

	return false
}

// SetAllReviewerLevels gets a reference to the given []ReviewerLevelInfoFull and assigns it to the AllReviewerLevels field.
func (o *ReviewFull) SetAllReviewerLevels(v []ReviewerLevelInfoFull) {
	o.AllReviewerLevels = v
}

func (o ReviewFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	toSerialize["_links"] = o.Links
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["resourceId"] = o.ResourceId
	if !IsNil(o.EntitlementValue) {
		toSerialize["entitlementValue"] = o.EntitlementValue
	}
	if !IsNil(o.EntitlementBundle) {
		toSerialize["entitlementBundle"] = o.EntitlementBundle
	}
	toSerialize["decision"] = o.Decision
	if !IsNil(o.Decided) {
		toSerialize["decided"] = o.Decided
	}
	toSerialize["remediationStatus"] = o.RemediationStatus
	toSerialize["principalProfile"] = o.PrincipalProfile
	if !IsNil(o.ReviewerProfile) {
		toSerialize["reviewerProfile"] = o.ReviewerProfile
	}
	toSerialize["reviewerType"] = o.ReviewerType
	if !IsNil(o.ReviewerGroupProfile) {
		toSerialize["reviewerGroupProfile"] = o.ReviewerGroupProfile
	}
	if !IsNil(o.CurrentReviewerLevel) {
		toSerialize["currentReviewerLevel"] = o.CurrentReviewerLevel
	}
	if o.RiskRuleConflicts != nil {
		toSerialize["riskRuleConflicts"] = o.RiskRuleConflicts
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.AllReviewerLevels) {
		toSerialize["allReviewerLevels"] = o.AllReviewerLevels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"_links",
		"campaignId",
		"resourceId",
		"decision",
		"remediationStatus",
		"principalProfile",
		"reviewerType",
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

	varReviewFull := _ReviewFull{}

	err = json.Unmarshal(data, &varReviewFull)

	if err != nil {
		return err
	}

	*o = ReviewFull(varReviewFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "campaignId")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "entitlementValue")
		delete(additionalProperties, "entitlementBundle")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "decided")
		delete(additionalProperties, "remediationStatus")
		delete(additionalProperties, "principalProfile")
		delete(additionalProperties, "reviewerProfile")
		delete(additionalProperties, "reviewerType")
		delete(additionalProperties, "reviewerGroupProfile")
		delete(additionalProperties, "currentReviewerLevel")
		delete(additionalProperties, "riskRuleConflicts")
		delete(additionalProperties, "note")
		delete(additionalProperties, "allReviewerLevels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewFull struct {
	value *ReviewFull
	isSet bool
}

func (v NullableReviewFull) Get() *ReviewFull {
	return v.value
}

func (v *NullableReviewFull) Set(val *ReviewFull) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewFull) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewFull(val *ReviewFull) *NullableReviewFull {
	return &NullableReviewFull{value: val, isSet: true}
}

func (v NullableReviewFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
