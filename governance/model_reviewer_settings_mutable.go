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
)

// ReviewerSettingsMutable Identifies the kind of reviewer for Access Certification.  1) `type = USER`: a single reviewer reviews the whole campaign. The property `reviewerId` is required for this type.  2) `type = REVIEWER_EXPRESSION`: specifies an expression to derive a reviewer for the whole campaign. The property `fallBackReviewerId` is required for this type. Specify `fallBackReviewerId`, if the system can't determine the reviewer using the expression. In such a case, the fallback reviewer specified through `fallBackReviewerId` becomes the automatic reviewer of the campaign. This type can represent reviews based on the user manager as well. In such a case, the reviewer expression is framed to indicate the user manager.  3) `type = GROUP`: All the members of an Okta group specified through `reviewerGroupId` become reviewers of the campaign. The property `reviewerGroupId` is required for this type. Note that if the group is specified through `reviewerGroupId` and has only one member, then that reviewer is assigned as a reviewer to all reviewers and `type` in those reviews changes to `USER`.  4) `type = RESOURCE_OWNER`: All the owners of an Okta group become reviewers of the campaign. The property `fallBackReviewerId` is required for this type. For `RESOURCE_OWNER` to work, set the property `resourceSettings.type` to `GROUP` and `resourceSettings.targetResources` to an Okta group ID. Access Certifications derives all the owners of the groups specified through`resourceSettings.targetResources.resourceId` and assigns them as reviewers to respective reviews of the campaign. Note that, if there is only one owner for such group(s) then that owner is assigned as a reviewer to respective reviews and `type` in those reviews changes to `USER`. Specify `fallBackReviewerId`, if the system can't determine the reviewer using the group(s) specified through `resourceSettings.targetResources.resourceId`. In such a case, the fallback reviewer specified through `fallBackReviewerId` becomes automatic reviewer of the campaign. Note that, if the provided Okta group has more than 10 members, when the campaign gets launched, only a maximum of 10 members are pulled from the group. Those 10 members become the reviewers of the campaign. When fetching members of the group, there's no order guaranteed. Additionally, if the Okta group specified has only one member, then that member is assigned as a reviewer to all reviewers and `type` for those reviews changes to `USER`.  5) `type = MULTI_LEVEL`: a campaign that can be composed of more than one reviewer level. The property `reviewerLevels` is required. The campaign can be reviewed and passed on from reviewer at one level to the reviewer defined in the next level. This setting provides more than one reviewer who can decide on the reviews before making the final decision. A maximum of two reviewer levels are supported. One has the flexibility to choose a combination of reviewers when defining the campaign. When a campaign is defined as multi level, only the property `reviewerLevels` is required. All other properties are ignored. Note that, if the provided Okta group has more than 10 owners, when the campaign gets launched, only a maximum of 10 owners are pulled from the group. Those 10 owners become the reviewers of the campaign. When fetching owners of the group, there's no order guaranteed. Additionally, if the Okta group specified has only one owner, then that owner is assigned as a reviewer to all reviewers and `type` for those reviews changes to `USER`.
type ReviewerSettingsMutable struct {
	Type CampaignReviewerType `json:"type"`
	// This property is required when `type=USER`  The `id` of an Okta user who will be assigned as a reviewer. By specifying `reviewerId`, all reviews will be assigned only to that reviewer. This is a conditional property to be provided only if the `reviewerScopeExpression` can't be specified.
	ReviewerId *string `json:"reviewerId,omitempty"`
	// This property is required when `type=REVIEWER_EXPRESSION`  The Okta specific user expression to fetch the reviewers from a connected identity source.  If a user is found using the provided expression, that user is assigned as a reviewer.  If a user is not found using the provided expression, the `fallbackReviewerId` is used.
	ReviewerScopeExpression *string `json:"reviewerScopeExpression,omitempty"`
	// Required when the `type=REVIEWER_EXPRESSION` or `type=RESOURCE_OWNER`  If the reviewer(s) can't be identified using `reviewerScopeExpression` or using the owners of groups (the owners of the group identified through the property `resourceSettings.type = GROUP`), the fallback reviewer (`id` of an Okta user) is assigned as a reviewer.
	FallBackReviewerId *string `json:"fallBackReviewerId,omitempty"`
	// This property is `required` when `type=GROUP`  The `id` of an Okta group who will be assigned as a reviewer. By specifying `reviewerGroupId`, all reviews will be assigned to that group. Any user part of that group will automatically be the reviewers of the campaign. This is a conditional property to be provided, only when one wants to assign more than one reviewer and cannot use `reviewerId` or `reviewerScopeExpression`.  Note that, if the provided Okta group has more than 10 members, when the campaign gets launched, only a max of 10 members will be pulled from the group. Those 10 members would be the reviewers of the campaign. When fetching members of the group, there is no order guaranteed.  Additionally, if the Okta group specified has only one member, then that member will be assigned as a reviewer to all reviewers and `type` for those reviews changes to `USER`.
	ReviewerGroupId *string `json:"reviewerGroupId,omitempty"`
	// If `true`, users won't be able to review their own review items  > **Note:** This field is deprecated. Use ['selfReviewDisabled'](/openapi/governance.api/tag/Campaigns/#tag/Campaigns/operation/createCampaign!path=reviewerSettings/selfReviewDisabled&t=request)
	// Deprecated
	IsSelfReviewDisabled *bool `json:"isSelfReviewDisabled,omitempty"`
	// If `true`, users won't be able to review their own review items  This property is required to be `true` for resource-centric campaigns when the Okta Admin Console is one of the resources
	SelfReviewDisabled *bool `json:"selfReviewDisabled,omitempty"`
	// When approving or revoking review items, a justification is required if true.  This property is required to be `true` for resource-centric campaigns when the Okta Admin Console is one of the resources
	JustificationRequired *bool `json:"justificationRequired,omitempty"`
	// When approving or revoking review items, bulk actions are disabled if true.
	BulkDecisionDisabled *bool `json:"bulkDecisionDisabled,omitempty"`
	// reassignment is disabled for reviewers if true.
	ReassignmentDisabled *bool `json:"reassignmentDisabled,omitempty"`
	// Definition of reviewer level for a given campaign. Each reviewer level defines the kind of reviewer who is going to review.  Each reviewer level defines the reviewer `type`.    - `type = USER`: a single reviewer reviews at that level. The property `reviewerId` is required.   - `type = REVIEWER_EXPRESSION`: specifies an expression to derive a reviewer for that level. The property `fallBackReviewerId` is required. Specify `fallBackReviewerId`, if Access Certifications can't determine the reviewer using the expression. In such a case, the fallback reviewer specified through `fallBackReviewerId` becomes automatic reviewer at that level.   - `type = GROUP`: all the members of an Okta group specified through `reviewerGroupId` become reviewers at that level. The property `reviewerGroupId` is required.     Note that, if the group specified through `reviewerGroupId` has only one member, then that member will be assigned as a reviewer to all reviewers and `type` at that level changes to `USER`.     Additionally, if the provided Okta group has more than 10 members, when the campaign gets launched, only a maximum of 10 members will be pulled from the group.     Those 10 members would be the reviewers at that level. When fetching members of the group, there is no order guaranteed.   - `type = RESOURCE_OWNER`: all the owners of an Okta group become reviewers at that level. The property `fallBackReviewerId` is required. For `RESOURCE_OWNER` to work, the property `resourceSettings.type` should be `GROUP` and `resourceSettings.targetResources` should be an Okta group ID. Access Certifications derives all the owners of the groups specified through `resourceSettings.targetResources.resourceId` and assigns them as reviewers to respective reviews of the campaign. Note that, if there is only one owner for such group(s) then that owner will be assigned as a reviewer to respective reviews and type in those reviews changes to `USER`. Specify `fallBackReviewerId`, if Access Certifications can't determine the reviewer using the group(s) specified through `resourceSettings.targetResources.resourceId`. In such a case, the fallback reviewer specified through `fallBackReviewerId` becomes automatic reviewer of the campaign.     Note that, if the Okta group specified has only one member, then that owner will be assigned as a reviewer to all reviewers and `type` at that level changes to `USER`.     Additionally, if the provided Okta group has more than 10 owners, when the campaign gets launched, only a maximum of 10 owners will be pulled from the group.     Those 10 owners would be the reviewers at that level. When fetching owners of the group, there is no order guaranteed.  Additionally -   - The property `startReview.onDay` indicates from which day, the level should start. This property will be `0` for `FIRST` reviewer level, as the first level will start when the campaign starts. For `SECOND` reviewer level specify a value greater than `0`.   - The property `startReview.when` indicates the condition when the lower level reviews will move to that level. This property should not be set for `FIRST` reviewer level. Only `SECOND` reviewer level will have the appropriate condition.
	ReviewerLevels       []ReviewerLevelSettingsMutable `json:"reviewerLevels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerSettingsMutable ReviewerSettingsMutable

// NewReviewerSettingsMutable instantiates a new ReviewerSettingsMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerSettingsMutable(type_ CampaignReviewerType) *ReviewerSettingsMutable {
	this := ReviewerSettingsMutable{}
	this.Type = type_
	return &this
}

// NewReviewerSettingsMutableWithDefaults instantiates a new ReviewerSettingsMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerSettingsMutableWithDefaults() *ReviewerSettingsMutable {
	this := ReviewerSettingsMutable{}
	return &this
}

// GetType returns the Type field value
func (o *ReviewerSettingsMutable) GetType() CampaignReviewerType {
	if o == nil {
		var ret CampaignReviewerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetTypeOk() (*CampaignReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReviewerSettingsMutable) SetType(v CampaignReviewerType) {
	o.Type = v
}

// GetReviewerId returns the ReviewerId field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetReviewerId() string {
	if o == nil || o.ReviewerId == nil {
		var ret string
		return ret
	}
	return *o.ReviewerId
}

// GetReviewerIdOk returns a tuple with the ReviewerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerIdOk() (*string, bool) {
	if o == nil || o.ReviewerId == nil {
		return nil, false
	}
	return o.ReviewerId, true
}

// HasReviewerId returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerId() bool {
	if o != nil && o.ReviewerId != nil {
		return true
	}

	return false
}

// SetReviewerId gets a reference to the given string and assigns it to the ReviewerId field.
func (o *ReviewerSettingsMutable) SetReviewerId(v string) {
	o.ReviewerId = &v
}

// GetReviewerScopeExpression returns the ReviewerScopeExpression field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetReviewerScopeExpression() string {
	if o == nil || o.ReviewerScopeExpression == nil {
		var ret string
		return ret
	}
	return *o.ReviewerScopeExpression
}

// GetReviewerScopeExpressionOk returns a tuple with the ReviewerScopeExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerScopeExpressionOk() (*string, bool) {
	if o == nil || o.ReviewerScopeExpression == nil {
		return nil, false
	}
	return o.ReviewerScopeExpression, true
}

// HasReviewerScopeExpression returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerScopeExpression() bool {
	if o != nil && o.ReviewerScopeExpression != nil {
		return true
	}

	return false
}

// SetReviewerScopeExpression gets a reference to the given string and assigns it to the ReviewerScopeExpression field.
func (o *ReviewerSettingsMutable) SetReviewerScopeExpression(v string) {
	o.ReviewerScopeExpression = &v
}

// GetFallBackReviewerId returns the FallBackReviewerId field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetFallBackReviewerId() string {
	if o == nil || o.FallBackReviewerId == nil {
		var ret string
		return ret
	}
	return *o.FallBackReviewerId
}

// GetFallBackReviewerIdOk returns a tuple with the FallBackReviewerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetFallBackReviewerIdOk() (*string, bool) {
	if o == nil || o.FallBackReviewerId == nil {
		return nil, false
	}
	return o.FallBackReviewerId, true
}

// HasFallBackReviewerId returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasFallBackReviewerId() bool {
	if o != nil && o.FallBackReviewerId != nil {
		return true
	}

	return false
}

// SetFallBackReviewerId gets a reference to the given string and assigns it to the FallBackReviewerId field.
func (o *ReviewerSettingsMutable) SetFallBackReviewerId(v string) {
	o.FallBackReviewerId = &v
}

// GetReviewerGroupId returns the ReviewerGroupId field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetReviewerGroupId() string {
	if o == nil || o.ReviewerGroupId == nil {
		var ret string
		return ret
	}
	return *o.ReviewerGroupId
}

// GetReviewerGroupIdOk returns a tuple with the ReviewerGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerGroupIdOk() (*string, bool) {
	if o == nil || o.ReviewerGroupId == nil {
		return nil, false
	}
	return o.ReviewerGroupId, true
}

// HasReviewerGroupId returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerGroupId() bool {
	if o != nil && o.ReviewerGroupId != nil {
		return true
	}

	return false
}

// SetReviewerGroupId gets a reference to the given string and assigns it to the ReviewerGroupId field.
func (o *ReviewerSettingsMutable) SetReviewerGroupId(v string) {
	o.ReviewerGroupId = &v
}

// GetIsSelfReviewDisabled returns the IsSelfReviewDisabled field value if set, zero value otherwise.
// Deprecated
func (o *ReviewerSettingsMutable) GetIsSelfReviewDisabled() bool {
	if o == nil || o.IsSelfReviewDisabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSelfReviewDisabled
}

// GetIsSelfReviewDisabledOk returns a tuple with the IsSelfReviewDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ReviewerSettingsMutable) GetIsSelfReviewDisabledOk() (*bool, bool) {
	if o == nil || o.IsSelfReviewDisabled == nil {
		return nil, false
	}
	return o.IsSelfReviewDisabled, true
}

// HasIsSelfReviewDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasIsSelfReviewDisabled() bool {
	if o != nil && o.IsSelfReviewDisabled != nil {
		return true
	}

	return false
}

// SetIsSelfReviewDisabled gets a reference to the given bool and assigns it to the IsSelfReviewDisabled field.
// Deprecated
func (o *ReviewerSettingsMutable) SetIsSelfReviewDisabled(v bool) {
	o.IsSelfReviewDisabled = &v
}

// GetSelfReviewDisabled returns the SelfReviewDisabled field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetSelfReviewDisabled() bool {
	if o == nil || o.SelfReviewDisabled == nil {
		var ret bool
		return ret
	}
	return *o.SelfReviewDisabled
}

// GetSelfReviewDisabledOk returns a tuple with the SelfReviewDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetSelfReviewDisabledOk() (*bool, bool) {
	if o == nil || o.SelfReviewDisabled == nil {
		return nil, false
	}
	return o.SelfReviewDisabled, true
}

// HasSelfReviewDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasSelfReviewDisabled() bool {
	if o != nil && o.SelfReviewDisabled != nil {
		return true
	}

	return false
}

// SetSelfReviewDisabled gets a reference to the given bool and assigns it to the SelfReviewDisabled field.
func (o *ReviewerSettingsMutable) SetSelfReviewDisabled(v bool) {
	o.SelfReviewDisabled = &v
}

// GetJustificationRequired returns the JustificationRequired field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetJustificationRequired() bool {
	if o == nil || o.JustificationRequired == nil {
		var ret bool
		return ret
	}
	return *o.JustificationRequired
}

// GetJustificationRequiredOk returns a tuple with the JustificationRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetJustificationRequiredOk() (*bool, bool) {
	if o == nil || o.JustificationRequired == nil {
		return nil, false
	}
	return o.JustificationRequired, true
}

// HasJustificationRequired returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasJustificationRequired() bool {
	if o != nil && o.JustificationRequired != nil {
		return true
	}

	return false
}

// SetJustificationRequired gets a reference to the given bool and assigns it to the JustificationRequired field.
func (o *ReviewerSettingsMutable) SetJustificationRequired(v bool) {
	o.JustificationRequired = &v
}

// GetBulkDecisionDisabled returns the BulkDecisionDisabled field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetBulkDecisionDisabled() bool {
	if o == nil || o.BulkDecisionDisabled == nil {
		var ret bool
		return ret
	}
	return *o.BulkDecisionDisabled
}

// GetBulkDecisionDisabledOk returns a tuple with the BulkDecisionDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetBulkDecisionDisabledOk() (*bool, bool) {
	if o == nil || o.BulkDecisionDisabled == nil {
		return nil, false
	}
	return o.BulkDecisionDisabled, true
}

// HasBulkDecisionDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasBulkDecisionDisabled() bool {
	if o != nil && o.BulkDecisionDisabled != nil {
		return true
	}

	return false
}

// SetBulkDecisionDisabled gets a reference to the given bool and assigns it to the BulkDecisionDisabled field.
func (o *ReviewerSettingsMutable) SetBulkDecisionDisabled(v bool) {
	o.BulkDecisionDisabled = &v
}

// GetReassignmentDisabled returns the ReassignmentDisabled field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetReassignmentDisabled() bool {
	if o == nil || o.ReassignmentDisabled == nil {
		var ret bool
		return ret
	}
	return *o.ReassignmentDisabled
}

// GetReassignmentDisabledOk returns a tuple with the ReassignmentDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReassignmentDisabledOk() (*bool, bool) {
	if o == nil || o.ReassignmentDisabled == nil {
		return nil, false
	}
	return o.ReassignmentDisabled, true
}

// HasReassignmentDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReassignmentDisabled() bool {
	if o != nil && o.ReassignmentDisabled != nil {
		return true
	}

	return false
}

// SetReassignmentDisabled gets a reference to the given bool and assigns it to the ReassignmentDisabled field.
func (o *ReviewerSettingsMutable) SetReassignmentDisabled(v bool) {
	o.ReassignmentDisabled = &v
}

// GetReviewerLevels returns the ReviewerLevels field value if set, zero value otherwise.
func (o *ReviewerSettingsMutable) GetReviewerLevels() []ReviewerLevelSettingsMutable {
	if o == nil || o.ReviewerLevels == nil {
		var ret []ReviewerLevelSettingsMutable
		return ret
	}
	return o.ReviewerLevels
}

// GetReviewerLevelsOk returns a tuple with the ReviewerLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerLevelsOk() ([]ReviewerLevelSettingsMutable, bool) {
	if o == nil || o.ReviewerLevels == nil {
		return nil, false
	}
	return o.ReviewerLevels, true
}

// HasReviewerLevels returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerLevels() bool {
	if o != nil && o.ReviewerLevels != nil {
		return true
	}

	return false
}

// SetReviewerLevels gets a reference to the given []ReviewerLevelSettingsMutable and assigns it to the ReviewerLevels field.
func (o *ReviewerSettingsMutable) SetReviewerLevels(v []ReviewerLevelSettingsMutable) {
	o.ReviewerLevels = v
}

func (o ReviewerSettingsMutable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ReviewerId != nil {
		toSerialize["reviewerId"] = o.ReviewerId
	}
	if o.ReviewerScopeExpression != nil {
		toSerialize["reviewerScopeExpression"] = o.ReviewerScopeExpression
	}
	if o.FallBackReviewerId != nil {
		toSerialize["fallBackReviewerId"] = o.FallBackReviewerId
	}
	if o.ReviewerGroupId != nil {
		toSerialize["reviewerGroupId"] = o.ReviewerGroupId
	}
	if o.IsSelfReviewDisabled != nil {
		toSerialize["isSelfReviewDisabled"] = o.IsSelfReviewDisabled
	}
	if o.SelfReviewDisabled != nil {
		toSerialize["selfReviewDisabled"] = o.SelfReviewDisabled
	}
	if o.JustificationRequired != nil {
		toSerialize["justificationRequired"] = o.JustificationRequired
	}
	if o.BulkDecisionDisabled != nil {
		toSerialize["bulkDecisionDisabled"] = o.BulkDecisionDisabled
	}
	if o.ReassignmentDisabled != nil {
		toSerialize["reassignmentDisabled"] = o.ReassignmentDisabled
	}
	if o.ReviewerLevels != nil {
		toSerialize["reviewerLevels"] = o.ReviewerLevels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewerSettingsMutable) UnmarshalJSON(bytes []byte) (err error) {
	varReviewerSettingsMutable := _ReviewerSettingsMutable{}

	err = json.Unmarshal(bytes, &varReviewerSettingsMutable)
	if err == nil {
		*o = ReviewerSettingsMutable(varReviewerSettingsMutable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "reviewerId")
		delete(additionalProperties, "reviewerScopeExpression")
		delete(additionalProperties, "fallBackReviewerId")
		delete(additionalProperties, "reviewerGroupId")
		delete(additionalProperties, "isSelfReviewDisabled")
		delete(additionalProperties, "selfReviewDisabled")
		delete(additionalProperties, "justificationRequired")
		delete(additionalProperties, "bulkDecisionDisabled")
		delete(additionalProperties, "reassignmentDisabled")
		delete(additionalProperties, "reviewerLevels")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewerSettingsMutable struct {
	value *ReviewerSettingsMutable
	isSet bool
}

func (v NullableReviewerSettingsMutable) Get() *ReviewerSettingsMutable {
	return v.value
}

func (v *NullableReviewerSettingsMutable) Set(val *ReviewerSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerSettingsMutable(val *ReviewerSettingsMutable) *NullableReviewerSettingsMutable {
	return &NullableReviewerSettingsMutable{value: val, isSet: true}
}

func (v NullableReviewerSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
