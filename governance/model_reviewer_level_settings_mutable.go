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

// checks if the ReviewerLevelSettingsMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerLevelSettingsMutable{}

// ReviewerLevelSettingsMutable struct for ReviewerLevelSettingsMutable
type ReviewerLevelSettingsMutable struct {
	Type        ReviewerType             `json:"type"`
	StartReview ReviewerLevelStartReview `json:"startReview"`
	// This property is required when `reviewerSettings.reviewerLevels.type` is `USER`  The `id` of an Okta user to assign all reviews. Use this reviewer type when you can't specify `reviewerScopeExpression`.
	ReviewerId *string `json:"reviewerId,omitempty"`
	// Required when `reviewerSettings.reviewerLevels.type` is `REVIEWER_EXPRESSION`.  The Okta-specific user expression to fetch the reviewers from a connected identity source:  * If a user is found with the provided expression, that user is assigned as the reviewer. * If a user isn't found with the provided expression, the `fallBackReviewerId` user is assigned as the reviewer. * See [Okta Expression Language (EL)](https://developer.okta.com/docs/reference/okta-expression-language/#okta-user-profile) to build the expression based on user profile attributes.
	ReviewerScopeExpression *string `json:"reviewerScopeExpression,omitempty"`
	// Required when reviewer setting `type` is `REVIEWER_EXPRESSION` or `RESOURCE_OWNER`. The fallback reviewer is assigned as the reviewer if:  * reviewer setting `reviewerScopeExpression` fails to identify reviewers, or * reviewers aren't identified through resource owners
	FallBackReviewerId *string `json:"fallBackReviewerId,omitempty"`
	// Required when `reviewerSettings.reviewerLevels.type` is `GROUP`.  The `id` of the Okta group: * All members of the specified group are assigned as reviewers. * Use this reviewer group assignment to assign more than one reviewer if you can't use `reviewerId` or `reviewerScopeExpression`. * If the Okta group has more than 10 members when the campaign launches, only 10 members are randomly selected as reviewers. * If the Okta group has only one member, then that member is assigned as the reviewer for all reviews, and `reviewerType` is set to `USER` for those reviews.
	ReviewerGroupId *string `json:"reviewerGroupId,omitempty"`
	// If `true`, users can't review their own review items.  > **Note:** This field is deprecated. Use 'selfReviewDisabled'.
	// Deprecated
	IsSelfReviewDisabled *bool `json:"isSelfReviewDisabled,omitempty"`
	// If `true`, users can't review their own review items.  This property must be `true` for resource-centric campaigns when the Okta Admin Console is one of the resources.
	SelfReviewDisabled   *bool `json:"selfReviewDisabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerLevelSettingsMutable ReviewerLevelSettingsMutable

// NewReviewerLevelSettingsMutable instantiates a new ReviewerLevelSettingsMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerLevelSettingsMutable(type_ ReviewerType, startReview ReviewerLevelStartReview) *ReviewerLevelSettingsMutable {
	this := ReviewerLevelSettingsMutable{}
	this.Type = type_
	this.StartReview = startReview
	return &this
}

// NewReviewerLevelSettingsMutableWithDefaults instantiates a new ReviewerLevelSettingsMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerLevelSettingsMutableWithDefaults() *ReviewerLevelSettingsMutable {
	this := ReviewerLevelSettingsMutable{}
	return &this
}

// GetType returns the Type field value
func (o *ReviewerLevelSettingsMutable) GetType() ReviewerType {
	if o == nil {
		var ret ReviewerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelSettingsMutable) GetTypeOk() (*ReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReviewerLevelSettingsMutable) SetType(v ReviewerType) {
	o.Type = v
}

// GetStartReview returns the StartReview field value
func (o *ReviewerLevelSettingsMutable) GetStartReview() ReviewerLevelStartReview {
	if o == nil {
		var ret ReviewerLevelStartReview
		return ret
	}

	return o.StartReview
}

// GetStartReviewOk returns a tuple with the StartReview field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelSettingsMutable) GetStartReviewOk() (*ReviewerLevelStartReview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartReview, true
}

// SetStartReview sets field value
func (o *ReviewerLevelSettingsMutable) SetStartReview(v ReviewerLevelStartReview) {
	o.StartReview = v
}

// GetReviewerId returns the ReviewerId field value if set, zero value otherwise.
func (o *ReviewerLevelSettingsMutable) GetReviewerId() string {
	if o == nil || IsNil(o.ReviewerId) {
		var ret string
		return ret
	}
	return *o.ReviewerId
}

// GetReviewerIdOk returns a tuple with the ReviewerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelSettingsMutable) GetReviewerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerId) {
		return nil, false
	}
	return o.ReviewerId, true
}

// HasReviewerId returns a boolean if a field has been set.
func (o *ReviewerLevelSettingsMutable) HasReviewerId() bool {
	if o != nil && !IsNil(o.ReviewerId) {
		return true
	}

	return false
}

// SetReviewerId gets a reference to the given string and assigns it to the ReviewerId field.
func (o *ReviewerLevelSettingsMutable) SetReviewerId(v string) {
	o.ReviewerId = &v
}

// GetReviewerScopeExpression returns the ReviewerScopeExpression field value if set, zero value otherwise.
func (o *ReviewerLevelSettingsMutable) GetReviewerScopeExpression() string {
	if o == nil || IsNil(o.ReviewerScopeExpression) {
		var ret string
		return ret
	}
	return *o.ReviewerScopeExpression
}

// GetReviewerScopeExpressionOk returns a tuple with the ReviewerScopeExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelSettingsMutable) GetReviewerScopeExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerScopeExpression) {
		return nil, false
	}
	return o.ReviewerScopeExpression, true
}

// HasReviewerScopeExpression returns a boolean if a field has been set.
func (o *ReviewerLevelSettingsMutable) HasReviewerScopeExpression() bool {
	if o != nil && !IsNil(o.ReviewerScopeExpression) {
		return true
	}

	return false
}

// SetReviewerScopeExpression gets a reference to the given string and assigns it to the ReviewerScopeExpression field.
func (o *ReviewerLevelSettingsMutable) SetReviewerScopeExpression(v string) {
	o.ReviewerScopeExpression = &v
}

// GetFallBackReviewerId returns the FallBackReviewerId field value if set, zero value otherwise.
func (o *ReviewerLevelSettingsMutable) GetFallBackReviewerId() string {
	if o == nil || IsNil(o.FallBackReviewerId) {
		var ret string
		return ret
	}
	return *o.FallBackReviewerId
}

// GetFallBackReviewerIdOk returns a tuple with the FallBackReviewerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelSettingsMutable) GetFallBackReviewerIdOk() (*string, bool) {
	if o == nil || IsNil(o.FallBackReviewerId) {
		return nil, false
	}
	return o.FallBackReviewerId, true
}

// HasFallBackReviewerId returns a boolean if a field has been set.
func (o *ReviewerLevelSettingsMutable) HasFallBackReviewerId() bool {
	if o != nil && !IsNil(o.FallBackReviewerId) {
		return true
	}

	return false
}

// SetFallBackReviewerId gets a reference to the given string and assigns it to the FallBackReviewerId field.
func (o *ReviewerLevelSettingsMutable) SetFallBackReviewerId(v string) {
	o.FallBackReviewerId = &v
}

// GetReviewerGroupId returns the ReviewerGroupId field value if set, zero value otherwise.
func (o *ReviewerLevelSettingsMutable) GetReviewerGroupId() string {
	if o == nil || IsNil(o.ReviewerGroupId) {
		var ret string
		return ret
	}
	return *o.ReviewerGroupId
}

// GetReviewerGroupIdOk returns a tuple with the ReviewerGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelSettingsMutable) GetReviewerGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerGroupId) {
		return nil, false
	}
	return o.ReviewerGroupId, true
}

// HasReviewerGroupId returns a boolean if a field has been set.
func (o *ReviewerLevelSettingsMutable) HasReviewerGroupId() bool {
	if o != nil && !IsNil(o.ReviewerGroupId) {
		return true
	}

	return false
}

// SetReviewerGroupId gets a reference to the given string and assigns it to the ReviewerGroupId field.
func (o *ReviewerLevelSettingsMutable) SetReviewerGroupId(v string) {
	o.ReviewerGroupId = &v
}

// GetIsSelfReviewDisabled returns the IsSelfReviewDisabled field value if set, zero value otherwise.
// Deprecated
func (o *ReviewerLevelSettingsMutable) GetIsSelfReviewDisabled() bool {
	if o == nil || IsNil(o.IsSelfReviewDisabled) {
		var ret bool
		return ret
	}
	return *o.IsSelfReviewDisabled
}

// GetIsSelfReviewDisabledOk returns a tuple with the IsSelfReviewDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ReviewerLevelSettingsMutable) GetIsSelfReviewDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSelfReviewDisabled) {
		return nil, false
	}
	return o.IsSelfReviewDisabled, true
}

// HasIsSelfReviewDisabled returns a boolean if a field has been set.
func (o *ReviewerLevelSettingsMutable) HasIsSelfReviewDisabled() bool {
	if o != nil && !IsNil(o.IsSelfReviewDisabled) {
		return true
	}

	return false
}

// SetIsSelfReviewDisabled gets a reference to the given bool and assigns it to the IsSelfReviewDisabled field.
// Deprecated
func (o *ReviewerLevelSettingsMutable) SetIsSelfReviewDisabled(v bool) {
	o.IsSelfReviewDisabled = &v
}

// GetSelfReviewDisabled returns the SelfReviewDisabled field value if set, zero value otherwise.
func (o *ReviewerLevelSettingsMutable) GetSelfReviewDisabled() bool {
	if o == nil || IsNil(o.SelfReviewDisabled) {
		var ret bool
		return ret
	}
	return *o.SelfReviewDisabled
}

// GetSelfReviewDisabledOk returns a tuple with the SelfReviewDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelSettingsMutable) GetSelfReviewDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfReviewDisabled) {
		return nil, false
	}
	return o.SelfReviewDisabled, true
}

// HasSelfReviewDisabled returns a boolean if a field has been set.
func (o *ReviewerLevelSettingsMutable) HasSelfReviewDisabled() bool {
	if o != nil && !IsNil(o.SelfReviewDisabled) {
		return true
	}

	return false
}

// SetSelfReviewDisabled gets a reference to the given bool and assigns it to the SelfReviewDisabled field.
func (o *ReviewerLevelSettingsMutable) SetSelfReviewDisabled(v bool) {
	o.SelfReviewDisabled = &v
}

func (o ReviewerLevelSettingsMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerLevelSettingsMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["startReview"] = o.StartReview
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerLevelSettingsMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"startReview",
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

	varReviewerLevelSettingsMutable := _ReviewerLevelSettingsMutable{}

	err = json.Unmarshal(data, &varReviewerLevelSettingsMutable)

	if err != nil {
		return err
	}

	*o = ReviewerLevelSettingsMutable(varReviewerLevelSettingsMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "startReview")
		delete(additionalProperties, "reviewerId")
		delete(additionalProperties, "reviewerScopeExpression")
		delete(additionalProperties, "fallBackReviewerId")
		delete(additionalProperties, "reviewerGroupId")
		delete(additionalProperties, "isSelfReviewDisabled")
		delete(additionalProperties, "selfReviewDisabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerLevelSettingsMutable struct {
	value *ReviewerLevelSettingsMutable
	isSet bool
}

func (v NullableReviewerLevelSettingsMutable) Get() *ReviewerLevelSettingsMutable {
	return v.value
}

func (v *NullableReviewerLevelSettingsMutable) Set(val *ReviewerLevelSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerLevelSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerLevelSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerLevelSettingsMutable(val *ReviewerLevelSettingsMutable) *NullableReviewerLevelSettingsMutable {
	return &NullableReviewerLevelSettingsMutable{value: val, isSet: true}
}

func (v NullableReviewerLevelSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerLevelSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
