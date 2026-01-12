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

// checks if the ReviewerSettingsMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerSettingsMutable{}

// ReviewerSettingsMutable Reviewer settings for the access certification campaign
type ReviewerSettingsMutable struct {
	Type CampaignReviewerType `json:"type"`
	// Required if `reviewerSettings.type` is `USER`.  The ID of an Okta user to assign all reviews.
	ReviewerId *string `json:"reviewerId,omitempty"`
	// Required when `reviewerSettings.type` is `REVIEWER_EXPRESSION`.  The Okta-specific user expression to fetch the reviewers from a connected identity source:  * If a user is found with the provided expression, that user is assigned as the reviewer. * If a user isn't found with the provided expression, the `reviewerSettings.fallBackReviewerId` user is assigned as the reviewer. * See [Okta Expression Language (EL)](https://developer.okta.com/docs/reference/okta-expression-language/#okta-user-profile) to build the expression base on user profile attributes.
	ReviewerScopeExpression *string `json:"reviewerScopeExpression,omitempty"`
	// Required when reviewer setting `type` is `REVIEWER_EXPRESSION` or `RESOURCE_OWNER`. The fallback reviewer is assigned as the reviewer if:  * reviewer setting `reviewerScopeExpression` fails to identify reviewers, or * reviewers aren't identified through resource owners
	FallBackReviewerId *string `json:"fallBackReviewerId,omitempty"`
	// Required when `reviewerSettings.type` is `GROUP`.  The `id` of the Okta group: * All members of the specified group are assigned as reviewers. * Use this reviewer group assignment if you can't use `reviewerId` or `reviewerScopeExpression`. * If the Okta group has more than 10 members when the campaign launches, only 10 members are randomly selected as reviewers. * If the Okta group has only one member, then that member is assigned as the reviewer for all reviews, and `reviewerType` is set to `USER` for those reviews.
	ReviewerGroupId *string `json:"reviewerGroupId,omitempty"`
	// If `true`, users can't review their own review items.  > **Note:** This field is deprecated. Use ['selfReviewDisabled'](/openapi/governance.api/tag/Campaigns/#tag/Campaigns/operation/createCampaign!path=reviewerSettings/selfReviewDisabled&t=request)
	// Deprecated
	IsSelfReviewDisabled *bool `json:"isSelfReviewDisabled,omitempty"`
	// If true, users won't be able to review their own review items.  This property is required to be `true` for resource-centric campaigns when the Okta Admin Console is one of the resources.
	SelfReviewDisabled *bool `json:"selfReviewDisabled,omitempty"`
	// If true, a justification is required when review items are approved or revoked.  This property must be `true` for resource-centric campaigns that have the Okta Admin Console as one of the resources.
	JustificationRequired *bool `json:"justificationRequired,omitempty"`
	// If true, bulk actions are disabled for approving or revoking review items.
	BulkDecisionDisabled *bool `json:"bulkDecisionDisabled,omitempty"`
	// If true, reassignment is disabled for reviewers.
	ReassignmentDisabled *bool `json:"reassignmentDisabled,omitempty"`
	// Defines the reviewer level in a campaign. A campaign can have a maximum of two reviewer levels.
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
	if o == nil || IsNil(o.ReviewerId) {
		var ret string
		return ret
	}
	return *o.ReviewerId
}

// GetReviewerIdOk returns a tuple with the ReviewerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerId) {
		return nil, false
	}
	return o.ReviewerId, true
}

// HasReviewerId returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerId() bool {
	if o != nil && !IsNil(o.ReviewerId) {
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
	if o == nil || IsNil(o.ReviewerScopeExpression) {
		var ret string
		return ret
	}
	return *o.ReviewerScopeExpression
}

// GetReviewerScopeExpressionOk returns a tuple with the ReviewerScopeExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerScopeExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerScopeExpression) {
		return nil, false
	}
	return o.ReviewerScopeExpression, true
}

// HasReviewerScopeExpression returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerScopeExpression() bool {
	if o != nil && !IsNil(o.ReviewerScopeExpression) {
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
	if o == nil || IsNil(o.FallBackReviewerId) {
		var ret string
		return ret
	}
	return *o.FallBackReviewerId
}

// GetFallBackReviewerIdOk returns a tuple with the FallBackReviewerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetFallBackReviewerIdOk() (*string, bool) {
	if o == nil || IsNil(o.FallBackReviewerId) {
		return nil, false
	}
	return o.FallBackReviewerId, true
}

// HasFallBackReviewerId returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasFallBackReviewerId() bool {
	if o != nil && !IsNil(o.FallBackReviewerId) {
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
	if o == nil || IsNil(o.ReviewerGroupId) {
		var ret string
		return ret
	}
	return *o.ReviewerGroupId
}

// GetReviewerGroupIdOk returns a tuple with the ReviewerGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerGroupId) {
		return nil, false
	}
	return o.ReviewerGroupId, true
}

// HasReviewerGroupId returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerGroupId() bool {
	if o != nil && !IsNil(o.ReviewerGroupId) {
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
	if o == nil || IsNil(o.IsSelfReviewDisabled) {
		var ret bool
		return ret
	}
	return *o.IsSelfReviewDisabled
}

// GetIsSelfReviewDisabledOk returns a tuple with the IsSelfReviewDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ReviewerSettingsMutable) GetIsSelfReviewDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSelfReviewDisabled) {
		return nil, false
	}
	return o.IsSelfReviewDisabled, true
}

// HasIsSelfReviewDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasIsSelfReviewDisabled() bool {
	if o != nil && !IsNil(o.IsSelfReviewDisabled) {
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
	if o == nil || IsNil(o.SelfReviewDisabled) {
		var ret bool
		return ret
	}
	return *o.SelfReviewDisabled
}

// GetSelfReviewDisabledOk returns a tuple with the SelfReviewDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetSelfReviewDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfReviewDisabled) {
		return nil, false
	}
	return o.SelfReviewDisabled, true
}

// HasSelfReviewDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasSelfReviewDisabled() bool {
	if o != nil && !IsNil(o.SelfReviewDisabled) {
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
	if o == nil || IsNil(o.JustificationRequired) {
		var ret bool
		return ret
	}
	return *o.JustificationRequired
}

// GetJustificationRequiredOk returns a tuple with the JustificationRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetJustificationRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.JustificationRequired) {
		return nil, false
	}
	return o.JustificationRequired, true
}

// HasJustificationRequired returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasJustificationRequired() bool {
	if o != nil && !IsNil(o.JustificationRequired) {
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
	if o == nil || IsNil(o.BulkDecisionDisabled) {
		var ret bool
		return ret
	}
	return *o.BulkDecisionDisabled
}

// GetBulkDecisionDisabledOk returns a tuple with the BulkDecisionDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetBulkDecisionDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BulkDecisionDisabled) {
		return nil, false
	}
	return o.BulkDecisionDisabled, true
}

// HasBulkDecisionDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasBulkDecisionDisabled() bool {
	if o != nil && !IsNil(o.BulkDecisionDisabled) {
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
	if o == nil || IsNil(o.ReassignmentDisabled) {
		var ret bool
		return ret
	}
	return *o.ReassignmentDisabled
}

// GetReassignmentDisabledOk returns a tuple with the ReassignmentDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReassignmentDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ReassignmentDisabled) {
		return nil, false
	}
	return o.ReassignmentDisabled, true
}

// HasReassignmentDisabled returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReassignmentDisabled() bool {
	if o != nil && !IsNil(o.ReassignmentDisabled) {
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
	if o == nil || IsNil(o.ReviewerLevels) {
		var ret []ReviewerLevelSettingsMutable
		return ret
	}
	return o.ReviewerLevels
}

// GetReviewerLevelsOk returns a tuple with the ReviewerLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerSettingsMutable) GetReviewerLevelsOk() ([]ReviewerLevelSettingsMutable, bool) {
	if o == nil || IsNil(o.ReviewerLevels) {
		return nil, false
	}
	return o.ReviewerLevels, true
}

// HasReviewerLevels returns a boolean if a field has been set.
func (o *ReviewerSettingsMutable) HasReviewerLevels() bool {
	if o != nil && !IsNil(o.ReviewerLevels) {
		return true
	}

	return false
}

// SetReviewerLevels gets a reference to the given []ReviewerLevelSettingsMutable and assigns it to the ReviewerLevels field.
func (o *ReviewerSettingsMutable) SetReviewerLevels(v []ReviewerLevelSettingsMutable) {
	o.ReviewerLevels = v
}

func (o ReviewerSettingsMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerSettingsMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ReviewerId) {
		toSerialize["reviewerId"] = o.ReviewerId
	}
	if !IsNil(o.ReviewerScopeExpression) {
		toSerialize["reviewerScopeExpression"] = o.ReviewerScopeExpression
	}
	if !IsNil(o.FallBackReviewerId) {
		toSerialize["fallBackReviewerId"] = o.FallBackReviewerId
	}
	if !IsNil(o.ReviewerGroupId) {
		toSerialize["reviewerGroupId"] = o.ReviewerGroupId
	}
	if !IsNil(o.IsSelfReviewDisabled) {
		toSerialize["isSelfReviewDisabled"] = o.IsSelfReviewDisabled
	}
	if !IsNil(o.SelfReviewDisabled) {
		toSerialize["selfReviewDisabled"] = o.SelfReviewDisabled
	}
	if !IsNil(o.JustificationRequired) {
		toSerialize["justificationRequired"] = o.JustificationRequired
	}
	if !IsNil(o.BulkDecisionDisabled) {
		toSerialize["bulkDecisionDisabled"] = o.BulkDecisionDisabled
	}
	if !IsNil(o.ReassignmentDisabled) {
		toSerialize["reassignmentDisabled"] = o.ReassignmentDisabled
	}
	if !IsNil(o.ReviewerLevels) {
		toSerialize["reviewerLevels"] = o.ReviewerLevels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerSettingsMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varReviewerSettingsMutable := _ReviewerSettingsMutable{}

	err = json.Unmarshal(data, &varReviewerSettingsMutable)

	if err != nil {
		return err
	}

	*o = ReviewerSettingsMutable(varReviewerSettingsMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
