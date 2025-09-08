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

// ReviewSparse Sparse representation of a Review resource
type ReviewSparse struct {
	Links ReviewLinks `json:"_links"`
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
	RiskRuleConflicts    []RiskRuleConflicts `json:"riskRuleConflicts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewSparse ReviewSparse

// NewReviewSparse instantiates a new ReviewSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSparse(links ReviewLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, campaignId string, resourceId string, decision Decision, remediationStatus RemediationStatus, principalProfile PrincipalProfile, reviewerType ReviewersReviewerType) *ReviewSparse {
	this := ReviewSparse{}
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

// NewReviewSparseWithDefaults instantiates a new ReviewSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSparseWithDefaults() *ReviewSparse {
	this := ReviewSparse{}
	return &this
}

// GetLinks returns the Links field value
func (o *ReviewSparse) GetLinks() ReviewLinks {
	if o == nil {
		var ret ReviewLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetLinksOk() (*ReviewLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReviewSparse) SetLinks(v ReviewLinks) {
	o.Links = v
}

// GetId returns the Id field value
func (o *ReviewSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewSparse) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ReviewSparse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ReviewSparse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *ReviewSparse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ReviewSparse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *ReviewSparse) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *ReviewSparse) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *ReviewSparse) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *ReviewSparse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetCampaignId returns the CampaignId field value
func (o *ReviewSparse) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *ReviewSparse) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetResourceId returns the ResourceId field value
func (o *ReviewSparse) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ReviewSparse) SetResourceId(v string) {
	o.ResourceId = v
}

// GetEntitlementValue returns the EntitlementValue field value if set, zero value otherwise.
func (o *ReviewSparse) GetEntitlementValue() ReviewerEntitlementValue {
	if o == nil || o.EntitlementValue == nil {
		var ret ReviewerEntitlementValue
		return ret
	}
	return *o.EntitlementValue
}

// GetEntitlementValueOk returns a tuple with the EntitlementValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetEntitlementValueOk() (*ReviewerEntitlementValue, bool) {
	if o == nil || o.EntitlementValue == nil {
		return nil, false
	}
	return o.EntitlementValue, true
}

// HasEntitlementValue returns a boolean if a field has been set.
func (o *ReviewSparse) HasEntitlementValue() bool {
	if o != nil && o.EntitlementValue != nil {
		return true
	}

	return false
}

// SetEntitlementValue gets a reference to the given ReviewerEntitlementValue and assigns it to the EntitlementValue field.
func (o *ReviewSparse) SetEntitlementValue(v ReviewerEntitlementValue) {
	o.EntitlementValue = &v
}

// GetEntitlementBundle returns the EntitlementBundle field value if set, zero value otherwise.
func (o *ReviewSparse) GetEntitlementBundle() ReviewerEntitlementBundle {
	if o == nil || o.EntitlementBundle == nil {
		var ret ReviewerEntitlementBundle
		return ret
	}
	return *o.EntitlementBundle
}

// GetEntitlementBundleOk returns a tuple with the EntitlementBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetEntitlementBundleOk() (*ReviewerEntitlementBundle, bool) {
	if o == nil || o.EntitlementBundle == nil {
		return nil, false
	}
	return o.EntitlementBundle, true
}

// HasEntitlementBundle returns a boolean if a field has been set.
func (o *ReviewSparse) HasEntitlementBundle() bool {
	if o != nil && o.EntitlementBundle != nil {
		return true
	}

	return false
}

// SetEntitlementBundle gets a reference to the given ReviewerEntitlementBundle and assigns it to the EntitlementBundle field.
func (o *ReviewSparse) SetEntitlementBundle(v ReviewerEntitlementBundle) {
	o.EntitlementBundle = &v
}

// GetDecision returns the Decision field value
func (o *ReviewSparse) GetDecision() Decision {
	if o == nil {
		var ret Decision
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetDecisionOk() (*Decision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *ReviewSparse) SetDecision(v Decision) {
	o.Decision = v
}

// GetDecided returns the Decided field value if set, zero value otherwise.
func (o *ReviewSparse) GetDecided() time.Time {
	if o == nil || o.Decided == nil {
		var ret time.Time
		return ret
	}
	return *o.Decided
}

// GetDecidedOk returns a tuple with the Decided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetDecidedOk() (*time.Time, bool) {
	if o == nil || o.Decided == nil {
		return nil, false
	}
	return o.Decided, true
}

// HasDecided returns a boolean if a field has been set.
func (o *ReviewSparse) HasDecided() bool {
	if o != nil && o.Decided != nil {
		return true
	}

	return false
}

// SetDecided gets a reference to the given time.Time and assigns it to the Decided field.
func (o *ReviewSparse) SetDecided(v time.Time) {
	o.Decided = &v
}

// GetRemediationStatus returns the RemediationStatus field value
func (o *ReviewSparse) GetRemediationStatus() RemediationStatus {
	if o == nil {
		var ret RemediationStatus
		return ret
	}

	return o.RemediationStatus
}

// GetRemediationStatusOk returns a tuple with the RemediationStatus field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetRemediationStatusOk() (*RemediationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationStatus, true
}

// SetRemediationStatus sets field value
func (o *ReviewSparse) SetRemediationStatus(v RemediationStatus) {
	o.RemediationStatus = v
}

// GetPrincipalProfile returns the PrincipalProfile field value
func (o *ReviewSparse) GetPrincipalProfile() PrincipalProfile {
	if o == nil {
		var ret PrincipalProfile
		return ret
	}

	return o.PrincipalProfile
}

// GetPrincipalProfileOk returns a tuple with the PrincipalProfile field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetPrincipalProfileOk() (*PrincipalProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalProfile, true
}

// SetPrincipalProfile sets field value
func (o *ReviewSparse) SetPrincipalProfile(v PrincipalProfile) {
	o.PrincipalProfile = v
}

// GetReviewerProfile returns the ReviewerProfile field value if set, zero value otherwise.
func (o *ReviewSparse) GetReviewerProfile() PrincipalProfile {
	if o == nil || o.ReviewerProfile == nil {
		var ret PrincipalProfile
		return ret
	}
	return *o.ReviewerProfile
}

// GetReviewerProfileOk returns a tuple with the ReviewerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetReviewerProfileOk() (*PrincipalProfile, bool) {
	if o == nil || o.ReviewerProfile == nil {
		return nil, false
	}
	return o.ReviewerProfile, true
}

// HasReviewerProfile returns a boolean if a field has been set.
func (o *ReviewSparse) HasReviewerProfile() bool {
	if o != nil && o.ReviewerProfile != nil {
		return true
	}

	return false
}

// SetReviewerProfile gets a reference to the given PrincipalProfile and assigns it to the ReviewerProfile field.
func (o *ReviewSparse) SetReviewerProfile(v PrincipalProfile) {
	o.ReviewerProfile = &v
}

// GetReviewerType returns the ReviewerType field value
func (o *ReviewSparse) GetReviewerType() ReviewersReviewerType {
	if o == nil {
		var ret ReviewersReviewerType
		return ret
	}

	return o.ReviewerType
}

// GetReviewerTypeOk returns a tuple with the ReviewerType field value
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetReviewerTypeOk() (*ReviewersReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerType, true
}

// SetReviewerType sets field value
func (o *ReviewSparse) SetReviewerType(v ReviewersReviewerType) {
	o.ReviewerType = v
}

// GetReviewerGroupProfile returns the ReviewerGroupProfile field value if set, zero value otherwise.
func (o *ReviewSparse) GetReviewerGroupProfile() ReviewerGroupProfile {
	if o == nil || o.ReviewerGroupProfile == nil {
		var ret ReviewerGroupProfile
		return ret
	}
	return *o.ReviewerGroupProfile
}

// GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool) {
	if o == nil || o.ReviewerGroupProfile == nil {
		return nil, false
	}
	return o.ReviewerGroupProfile, true
}

// HasReviewerGroupProfile returns a boolean if a field has been set.
func (o *ReviewSparse) HasReviewerGroupProfile() bool {
	if o != nil && o.ReviewerGroupProfile != nil {
		return true
	}

	return false
}

// SetReviewerGroupProfile gets a reference to the given ReviewerGroupProfile and assigns it to the ReviewerGroupProfile field.
func (o *ReviewSparse) SetReviewerGroupProfile(v ReviewerGroupProfile) {
	o.ReviewerGroupProfile = &v
}

// GetCurrentReviewerLevel returns the CurrentReviewerLevel field value if set, zero value otherwise.
func (o *ReviewSparse) GetCurrentReviewerLevel() ReviewerLevelType {
	if o == nil || o.CurrentReviewerLevel == nil {
		var ret ReviewerLevelType
		return ret
	}
	return *o.CurrentReviewerLevel
}

// GetCurrentReviewerLevelOk returns a tuple with the CurrentReviewerLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSparse) GetCurrentReviewerLevelOk() (*ReviewerLevelType, bool) {
	if o == nil || o.CurrentReviewerLevel == nil {
		return nil, false
	}
	return o.CurrentReviewerLevel, true
}

// HasCurrentReviewerLevel returns a boolean if a field has been set.
func (o *ReviewSparse) HasCurrentReviewerLevel() bool {
	if o != nil && o.CurrentReviewerLevel != nil {
		return true
	}

	return false
}

// SetCurrentReviewerLevel gets a reference to the given ReviewerLevelType and assigns it to the CurrentReviewerLevel field.
func (o *ReviewSparse) SetCurrentReviewerLevel(v ReviewerLevelType) {
	o.CurrentReviewerLevel = &v
}

// GetRiskRuleConflicts returns the RiskRuleConflicts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewSparse) GetRiskRuleConflicts() []RiskRuleConflicts {
	if o == nil {
		var ret []RiskRuleConflicts
		return ret
	}
	return o.RiskRuleConflicts
}

// GetRiskRuleConflictsOk returns a tuple with the RiskRuleConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewSparse) GetRiskRuleConflictsOk() ([]RiskRuleConflicts, bool) {
	if o == nil || o.RiskRuleConflicts == nil {
		return nil, false
	}
	return o.RiskRuleConflicts, true
}

// HasRiskRuleConflicts returns a boolean if a field has been set.
func (o *ReviewSparse) HasRiskRuleConflicts() bool {
	if o != nil && o.RiskRuleConflicts != nil {
		return true
	}

	return false
}

// SetRiskRuleConflicts gets a reference to the given []RiskRuleConflicts and assigns it to the RiskRuleConflicts field.
func (o *ReviewSparse) SetRiskRuleConflicts(v []RiskRuleConflicts) {
	o.RiskRuleConflicts = v
}

func (o ReviewSparse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
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
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.EntitlementValue != nil {
		toSerialize["entitlementValue"] = o.EntitlementValue
	}
	if o.EntitlementBundle != nil {
		toSerialize["entitlementBundle"] = o.EntitlementBundle
	}
	if true {
		toSerialize["decision"] = o.Decision
	}
	if o.Decided != nil {
		toSerialize["decided"] = o.Decided
	}
	if true {
		toSerialize["remediationStatus"] = o.RemediationStatus
	}
	if true {
		toSerialize["principalProfile"] = o.PrincipalProfile
	}
	if o.ReviewerProfile != nil {
		toSerialize["reviewerProfile"] = o.ReviewerProfile
	}
	if true {
		toSerialize["reviewerType"] = o.ReviewerType
	}
	if o.ReviewerGroupProfile != nil {
		toSerialize["reviewerGroupProfile"] = o.ReviewerGroupProfile
	}
	if o.CurrentReviewerLevel != nil {
		toSerialize["currentReviewerLevel"] = o.CurrentReviewerLevel
	}
	if o.RiskRuleConflicts != nil {
		toSerialize["riskRuleConflicts"] = o.RiskRuleConflicts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewSparse) UnmarshalJSON(bytes []byte) (err error) {
	varReviewSparse := _ReviewSparse{}

	err = json.Unmarshal(bytes, &varReviewSparse)
	if err == nil {
		*o = ReviewSparse(varReviewSparse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_links")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
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
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewSparse struct {
	value *ReviewSparse
	isSet bool
}

func (v NullableReviewSparse) Get() *ReviewSparse {
	return v.value
}

func (v *NullableReviewSparse) Set(val *ReviewSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSparse(val *ReviewSparse) *NullableReviewSparse {
	return &NullableReviewSparse{value: val, isSet: true}
}

func (v NullableReviewSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
