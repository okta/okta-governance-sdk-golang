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

// checks if the ReviewMinimalReadOnlyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewMinimalReadOnlyFields{}

// ReviewMinimalReadOnlyFields Properties contained by default in sparse representation.
type ReviewMinimalReadOnlyFields struct {
	CampaignId           string                     `json:"campaignId"`
	ResourceId           string                     `json:"resourceId"`
	EntitlementValue     *ReviewerEntitlementValue  `json:"entitlementValue,omitempty"`
	EntitlementBundle    *ReviewerEntitlementBundle `json:"entitlementBundle,omitempty"`
	Decision             Decision                   `json:"decision"`
	Decided              *time.Time                 `json:"decided,omitempty"`
	RemediationStatus    RemediationStatus          `json:"remediationStatus"`
	PrincipalProfile     PrincipalProfileEnriched   `json:"principalProfile"`
	ReviewerProfile      *PrincipalProfileEnriched  `json:"reviewerProfile,omitempty"`
	ReviewerType         ReviewersReviewerType      `json:"reviewerType"`
	ReviewerGroupProfile *ReviewerGroupProfile      `json:"reviewerGroupProfile,omitempty"`
	CurrentReviewerLevel *ReviewerLevelType         `json:"currentReviewerLevel,omitempty"`
	Links                ReviewLinks                `json:"_links"`
	// List of risk rule conflicts caused by this entitlement value. Only applies to review item that has entitlement value.
	RiskRuleConflicts []RiskRuleConflicts       `json:"riskRuleConflicts,omitempty"`
	DelegatorProfile  *PrincipalProfileEnriched `json:"delegatorProfile,omitempty"`
	// Specifies if this review was delegated by the original reviewer based on their governance delegate settings
	Delegated            *bool                   `json:"delegated,omitempty"`
	AppServiceAccount    *ReviewerServiceAccount `json:"appServiceAccount,omitempty"`
	OktaServiceAccount   *ReviewerServiceAccount `json:"oktaServiceAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewMinimalReadOnlyFields ReviewMinimalReadOnlyFields

// NewReviewMinimalReadOnlyFields instantiates a new ReviewMinimalReadOnlyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewMinimalReadOnlyFields(campaignId string, resourceId string, decision Decision, remediationStatus RemediationStatus, principalProfile PrincipalProfileEnriched, reviewerType ReviewersReviewerType, links ReviewLinks) *ReviewMinimalReadOnlyFields {
	this := ReviewMinimalReadOnlyFields{}
	this.CampaignId = campaignId
	this.ResourceId = resourceId
	this.Decision = decision
	this.RemediationStatus = remediationStatus
	this.PrincipalProfile = principalProfile
	this.ReviewerType = reviewerType
	this.Links = links
	return &this
}

// NewReviewMinimalReadOnlyFieldsWithDefaults instantiates a new ReviewMinimalReadOnlyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewMinimalReadOnlyFieldsWithDefaults() *ReviewMinimalReadOnlyFields {
	this := ReviewMinimalReadOnlyFields{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *ReviewMinimalReadOnlyFields) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *ReviewMinimalReadOnlyFields) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetResourceId returns the ResourceId field value
func (o *ReviewMinimalReadOnlyFields) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ReviewMinimalReadOnlyFields) SetResourceId(v string) {
	o.ResourceId = v
}

// GetEntitlementValue returns the EntitlementValue field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetEntitlementValue() ReviewerEntitlementValue {
	if o == nil || IsNil(o.EntitlementValue) {
		var ret ReviewerEntitlementValue
		return ret
	}
	return *o.EntitlementValue
}

// GetEntitlementValueOk returns a tuple with the EntitlementValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetEntitlementValueOk() (*ReviewerEntitlementValue, bool) {
	if o == nil || IsNil(o.EntitlementValue) {
		return nil, false
	}
	return o.EntitlementValue, true
}

// HasEntitlementValue returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasEntitlementValue() bool {
	if o != nil && !IsNil(o.EntitlementValue) {
		return true
	}

	return false
}

// SetEntitlementValue gets a reference to the given ReviewerEntitlementValue and assigns it to the EntitlementValue field.
func (o *ReviewMinimalReadOnlyFields) SetEntitlementValue(v ReviewerEntitlementValue) {
	o.EntitlementValue = &v
}

// GetEntitlementBundle returns the EntitlementBundle field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetEntitlementBundle() ReviewerEntitlementBundle {
	if o == nil || IsNil(o.EntitlementBundle) {
		var ret ReviewerEntitlementBundle
		return ret
	}
	return *o.EntitlementBundle
}

// GetEntitlementBundleOk returns a tuple with the EntitlementBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetEntitlementBundleOk() (*ReviewerEntitlementBundle, bool) {
	if o == nil || IsNil(o.EntitlementBundle) {
		return nil, false
	}
	return o.EntitlementBundle, true
}

// HasEntitlementBundle returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasEntitlementBundle() bool {
	if o != nil && !IsNil(o.EntitlementBundle) {
		return true
	}

	return false
}

// SetEntitlementBundle gets a reference to the given ReviewerEntitlementBundle and assigns it to the EntitlementBundle field.
func (o *ReviewMinimalReadOnlyFields) SetEntitlementBundle(v ReviewerEntitlementBundle) {
	o.EntitlementBundle = &v
}

// GetDecision returns the Decision field value
func (o *ReviewMinimalReadOnlyFields) GetDecision() Decision {
	if o == nil {
		var ret Decision
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetDecisionOk() (*Decision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *ReviewMinimalReadOnlyFields) SetDecision(v Decision) {
	o.Decision = v
}

// GetDecided returns the Decided field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetDecided() time.Time {
	if o == nil || IsNil(o.Decided) {
		var ret time.Time
		return ret
	}
	return *o.Decided
}

// GetDecidedOk returns a tuple with the Decided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetDecidedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Decided) {
		return nil, false
	}
	return o.Decided, true
}

// HasDecided returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasDecided() bool {
	if o != nil && !IsNil(o.Decided) {
		return true
	}

	return false
}

// SetDecided gets a reference to the given time.Time and assigns it to the Decided field.
func (o *ReviewMinimalReadOnlyFields) SetDecided(v time.Time) {
	o.Decided = &v
}

// GetRemediationStatus returns the RemediationStatus field value
func (o *ReviewMinimalReadOnlyFields) GetRemediationStatus() RemediationStatus {
	if o == nil {
		var ret RemediationStatus
		return ret
	}

	return o.RemediationStatus
}

// GetRemediationStatusOk returns a tuple with the RemediationStatus field value
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetRemediationStatusOk() (*RemediationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationStatus, true
}

// SetRemediationStatus sets field value
func (o *ReviewMinimalReadOnlyFields) SetRemediationStatus(v RemediationStatus) {
	o.RemediationStatus = v
}

// GetPrincipalProfile returns the PrincipalProfile field value
func (o *ReviewMinimalReadOnlyFields) GetPrincipalProfile() PrincipalProfileEnriched {
	if o == nil {
		var ret PrincipalProfileEnriched
		return ret
	}

	return o.PrincipalProfile
}

// GetPrincipalProfileOk returns a tuple with the PrincipalProfile field value
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetPrincipalProfileOk() (*PrincipalProfileEnriched, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalProfile, true
}

// SetPrincipalProfile sets field value
func (o *ReviewMinimalReadOnlyFields) SetPrincipalProfile(v PrincipalProfileEnriched) {
	o.PrincipalProfile = v
}

// GetReviewerProfile returns the ReviewerProfile field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetReviewerProfile() PrincipalProfileEnriched {
	if o == nil || IsNil(o.ReviewerProfile) {
		var ret PrincipalProfileEnriched
		return ret
	}
	return *o.ReviewerProfile
}

// GetReviewerProfileOk returns a tuple with the ReviewerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetReviewerProfileOk() (*PrincipalProfileEnriched, bool) {
	if o == nil || IsNil(o.ReviewerProfile) {
		return nil, false
	}
	return o.ReviewerProfile, true
}

// HasReviewerProfile returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasReviewerProfile() bool {
	if o != nil && !IsNil(o.ReviewerProfile) {
		return true
	}

	return false
}

// SetReviewerProfile gets a reference to the given PrincipalProfileEnriched and assigns it to the ReviewerProfile field.
func (o *ReviewMinimalReadOnlyFields) SetReviewerProfile(v PrincipalProfileEnriched) {
	o.ReviewerProfile = &v
}

// GetReviewerType returns the ReviewerType field value
func (o *ReviewMinimalReadOnlyFields) GetReviewerType() ReviewersReviewerType {
	if o == nil {
		var ret ReviewersReviewerType
		return ret
	}

	return o.ReviewerType
}

// GetReviewerTypeOk returns a tuple with the ReviewerType field value
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetReviewerTypeOk() (*ReviewersReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerType, true
}

// SetReviewerType sets field value
func (o *ReviewMinimalReadOnlyFields) SetReviewerType(v ReviewersReviewerType) {
	o.ReviewerType = v
}

// GetReviewerGroupProfile returns the ReviewerGroupProfile field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetReviewerGroupProfile() ReviewerGroupProfile {
	if o == nil || IsNil(o.ReviewerGroupProfile) {
		var ret ReviewerGroupProfile
		return ret
	}
	return *o.ReviewerGroupProfile
}

// GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool) {
	if o == nil || IsNil(o.ReviewerGroupProfile) {
		return nil, false
	}
	return o.ReviewerGroupProfile, true
}

// HasReviewerGroupProfile returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasReviewerGroupProfile() bool {
	if o != nil && !IsNil(o.ReviewerGroupProfile) {
		return true
	}

	return false
}

// SetReviewerGroupProfile gets a reference to the given ReviewerGroupProfile and assigns it to the ReviewerGroupProfile field.
func (o *ReviewMinimalReadOnlyFields) SetReviewerGroupProfile(v ReviewerGroupProfile) {
	o.ReviewerGroupProfile = &v
}

// GetCurrentReviewerLevel returns the CurrentReviewerLevel field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetCurrentReviewerLevel() ReviewerLevelType {
	if o == nil || IsNil(o.CurrentReviewerLevel) {
		var ret ReviewerLevelType
		return ret
	}
	return *o.CurrentReviewerLevel
}

// GetCurrentReviewerLevelOk returns a tuple with the CurrentReviewerLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetCurrentReviewerLevelOk() (*ReviewerLevelType, bool) {
	if o == nil || IsNil(o.CurrentReviewerLevel) {
		return nil, false
	}
	return o.CurrentReviewerLevel, true
}

// HasCurrentReviewerLevel returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasCurrentReviewerLevel() bool {
	if o != nil && !IsNil(o.CurrentReviewerLevel) {
		return true
	}

	return false
}

// SetCurrentReviewerLevel gets a reference to the given ReviewerLevelType and assigns it to the CurrentReviewerLevel field.
func (o *ReviewMinimalReadOnlyFields) SetCurrentReviewerLevel(v ReviewerLevelType) {
	o.CurrentReviewerLevel = &v
}

// GetLinks returns the Links field value
func (o *ReviewMinimalReadOnlyFields) GetLinks() ReviewLinks {
	if o == nil {
		var ret ReviewLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetLinksOk() (*ReviewLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReviewMinimalReadOnlyFields) SetLinks(v ReviewLinks) {
	o.Links = v
}

// GetRiskRuleConflicts returns the RiskRuleConflicts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReviewMinimalReadOnlyFields) GetRiskRuleConflicts() []RiskRuleConflicts {
	if o == nil {
		var ret []RiskRuleConflicts
		return ret
	}
	return o.RiskRuleConflicts
}

// GetRiskRuleConflictsOk returns a tuple with the RiskRuleConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReviewMinimalReadOnlyFields) GetRiskRuleConflictsOk() ([]RiskRuleConflicts, bool) {
	if o == nil || IsNil(o.RiskRuleConflicts) {
		return nil, false
	}
	return o.RiskRuleConflicts, true
}

// HasRiskRuleConflicts returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasRiskRuleConflicts() bool {
	if o != nil && !IsNil(o.RiskRuleConflicts) {
		return true
	}

	return false
}

// SetRiskRuleConflicts gets a reference to the given []RiskRuleConflicts and assigns it to the RiskRuleConflicts field.
func (o *ReviewMinimalReadOnlyFields) SetRiskRuleConflicts(v []RiskRuleConflicts) {
	o.RiskRuleConflicts = v
}

// GetDelegatorProfile returns the DelegatorProfile field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetDelegatorProfile() PrincipalProfileEnriched {
	if o == nil || IsNil(o.DelegatorProfile) {
		var ret PrincipalProfileEnriched
		return ret
	}
	return *o.DelegatorProfile
}

// GetDelegatorProfileOk returns a tuple with the DelegatorProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetDelegatorProfileOk() (*PrincipalProfileEnriched, bool) {
	if o == nil || IsNil(o.DelegatorProfile) {
		return nil, false
	}
	return o.DelegatorProfile, true
}

// HasDelegatorProfile returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasDelegatorProfile() bool {
	if o != nil && !IsNil(o.DelegatorProfile) {
		return true
	}

	return false
}

// SetDelegatorProfile gets a reference to the given PrincipalProfileEnriched and assigns it to the DelegatorProfile field.
func (o *ReviewMinimalReadOnlyFields) SetDelegatorProfile(v PrincipalProfileEnriched) {
	o.DelegatorProfile = &v
}

// GetDelegated returns the Delegated field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetDelegated() bool {
	if o == nil || IsNil(o.Delegated) {
		var ret bool
		return ret
	}
	return *o.Delegated
}

// GetDelegatedOk returns a tuple with the Delegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetDelegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Delegated) {
		return nil, false
	}
	return o.Delegated, true
}

// HasDelegated returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasDelegated() bool {
	if o != nil && !IsNil(o.Delegated) {
		return true
	}

	return false
}

// SetDelegated gets a reference to the given bool and assigns it to the Delegated field.
func (o *ReviewMinimalReadOnlyFields) SetDelegated(v bool) {
	o.Delegated = &v
}

// GetAppServiceAccount returns the AppServiceAccount field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetAppServiceAccount() ReviewerServiceAccount {
	if o == nil || IsNil(o.AppServiceAccount) {
		var ret ReviewerServiceAccount
		return ret
	}
	return *o.AppServiceAccount
}

// GetAppServiceAccountOk returns a tuple with the AppServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetAppServiceAccountOk() (*ReviewerServiceAccount, bool) {
	if o == nil || IsNil(o.AppServiceAccount) {
		return nil, false
	}
	return o.AppServiceAccount, true
}

// HasAppServiceAccount returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasAppServiceAccount() bool {
	if o != nil && !IsNil(o.AppServiceAccount) {
		return true
	}

	return false
}

// SetAppServiceAccount gets a reference to the given ReviewerServiceAccount and assigns it to the AppServiceAccount field.
func (o *ReviewMinimalReadOnlyFields) SetAppServiceAccount(v ReviewerServiceAccount) {
	o.AppServiceAccount = &v
}

// GetOktaServiceAccount returns the OktaServiceAccount field value if set, zero value otherwise.
func (o *ReviewMinimalReadOnlyFields) GetOktaServiceAccount() ReviewerServiceAccount {
	if o == nil || IsNil(o.OktaServiceAccount) {
		var ret ReviewerServiceAccount
		return ret
	}
	return *o.OktaServiceAccount
}

// GetOktaServiceAccountOk returns a tuple with the OktaServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewMinimalReadOnlyFields) GetOktaServiceAccountOk() (*ReviewerServiceAccount, bool) {
	if o == nil || IsNil(o.OktaServiceAccount) {
		return nil, false
	}
	return o.OktaServiceAccount, true
}

// HasOktaServiceAccount returns a boolean if a field has been set.
func (o *ReviewMinimalReadOnlyFields) HasOktaServiceAccount() bool {
	if o != nil && !IsNil(o.OktaServiceAccount) {
		return true
	}

	return false
}

// SetOktaServiceAccount gets a reference to the given ReviewerServiceAccount and assigns it to the OktaServiceAccount field.
func (o *ReviewMinimalReadOnlyFields) SetOktaServiceAccount(v ReviewerServiceAccount) {
	o.OktaServiceAccount = &v
}

func (o ReviewMinimalReadOnlyFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewMinimalReadOnlyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	toSerialize["_links"] = o.Links
	if o.RiskRuleConflicts != nil {
		toSerialize["riskRuleConflicts"] = o.RiskRuleConflicts
	}
	if !IsNil(o.DelegatorProfile) {
		toSerialize["delegatorProfile"] = o.DelegatorProfile
	}
	if !IsNil(o.Delegated) {
		toSerialize["delegated"] = o.Delegated
	}
	if !IsNil(o.AppServiceAccount) {
		toSerialize["appServiceAccount"] = o.AppServiceAccount
	}
	if !IsNil(o.OktaServiceAccount) {
		toSerialize["oktaServiceAccount"] = o.OktaServiceAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewMinimalReadOnlyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaignId",
		"resourceId",
		"decision",
		"remediationStatus",
		"principalProfile",
		"reviewerType",
		"_links",
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

	varReviewMinimalReadOnlyFields := _ReviewMinimalReadOnlyFields{}

	err = json.Unmarshal(data, &varReviewMinimalReadOnlyFields)

	if err != nil {
		return err
	}

	*o = ReviewMinimalReadOnlyFields(varReviewMinimalReadOnlyFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
		delete(additionalProperties, "_links")
		delete(additionalProperties, "riskRuleConflicts")
		delete(additionalProperties, "delegatorProfile")
		delete(additionalProperties, "delegated")
		delete(additionalProperties, "appServiceAccount")
		delete(additionalProperties, "oktaServiceAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewMinimalReadOnlyFields struct {
	value *ReviewMinimalReadOnlyFields
	isSet bool
}

func (v NullableReviewMinimalReadOnlyFields) Get() *ReviewMinimalReadOnlyFields {
	return v.value
}

func (v *NullableReviewMinimalReadOnlyFields) Set(val *ReviewMinimalReadOnlyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewMinimalReadOnlyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewMinimalReadOnlyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewMinimalReadOnlyFields(val *ReviewMinimalReadOnlyFields) *NullableReviewMinimalReadOnlyFields {
	return &NullableReviewMinimalReadOnlyFields{value: val, isSet: true}
}

func (v NullableReviewMinimalReadOnlyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewMinimalReadOnlyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
