# ReviewerSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignReviewerType**](CampaignReviewerType.md) |  | 
**ReviewerId** | Pointer to **string** | This property is required when &#x60;type&#x3D;USER&#x60;  The &#x60;id&#x60; of an Okta user who will be assigned as a reviewer. By specifying &#x60;reviewerId&#x60;, all reviews will be assigned only to that reviewer. This is a conditional property to be provided only if the &#x60;reviewerScopeExpression&#x60; can&#39;t be specified.  | [optional] 
**ReviewerScopeExpression** | Pointer to **string** | This property is required when &#x60;type&#x3D;REVIEWER_EXPRESSION&#x60;  The Okta specific user expression to fetch the reviewers from a connected identity source.  If a user is found using the provided expression, that user is assigned as a reviewer.  If a user is not found using the provided expression, the &#x60;fallbackReviewerId&#x60; is used.  | [optional] 
**FallBackReviewerId** | Pointer to **string** | Required when the &#x60;type&#x3D;REVIEWER_EXPRESSION&#x60; or &#x60;type&#x3D;RESOURCE_OWNER&#x60;  If the reviewer(s) can&#39;t be identified using &#x60;reviewerScopeExpression&#x60; or using the owners of groups (the owners of the group identified through the property &#x60;resourceSettings.type &#x3D; GROUP&#x60;), the fallback reviewer (&#x60;id&#x60; of an Okta user) is assigned as a reviewer.  | [optional] 
**ReviewerGroupId** | Pointer to **string** | This property is &#x60;required&#x60; when &#x60;type&#x3D;GROUP&#x60;  The &#x60;id&#x60; of an Okta group who will be assigned as a reviewer. By specifying &#x60;reviewerGroupId&#x60;, all reviews will be assigned to that group. Any user part of that group will automatically be the reviewers of the campaign. This is a conditional property to be provided, only when one wants to assign more than one reviewer and cannot use &#x60;reviewerId&#x60; or &#x60;reviewerScopeExpression&#x60;.  Note that, if the provided Okta group has more than 10 members, when the campaign gets launched, only a max of 10 members will be pulled from the group. Those 10 members would be the reviewers of the campaign. When fetching members of the group, there is no order guaranteed.  Additionally, if the Okta group specified has only one member, then that member will be assigned as a reviewer to all reviewers and &#x60;type&#x60; for those reviews changes to &#x60;USER&#x60;.  | [optional] 
**IsSelfReviewDisabled** | Pointer to **bool** | If &#x60;true&#x60;, users won&#39;t be able to review their own review items  &gt; **Note:** This field is deprecated. Use [&#39;selfReviewDisabled&#39;](/openapi/governance.api/tag/Campaigns/#tag/Campaigns/operation/createCampaign!path&#x3D;reviewerSettings/selfReviewDisabled&amp;t&#x3D;request)  | [optional] 
**SelfReviewDisabled** | Pointer to **bool** | If &#x60;true&#x60;, users won&#39;t be able to review their own review items  This property is required to be &#x60;true&#x60; for resource-centric campaigns when the Okta Admin Console is one of the resources  | [optional] 
**JustificationRequired** | Pointer to **bool** | When approving or revoking review items, a justification is required if true.  This property is required to be &#x60;true&#x60; for resource-centric campaigns when the Okta Admin Console is one of the resources  | [optional] 
**BulkDecisionDisabled** | Pointer to **bool** | When approving or revoking review items, bulk actions are disabled if true.  | [optional] 
**ReassignmentDisabled** | Pointer to **bool** | reassignment is disabled for reviewers if true.  | [optional] 
**ReviewerLevels** | Pointer to [**[]ReviewerLevelSettingsMutable**](ReviewerLevelSettingsMutable.md) | Definition of reviewer level for a given campaign. Each reviewer level defines the kind of reviewer who is going to review.  Each reviewer level defines the reviewer &#x60;type&#x60;.    - &#x60;type &#x3D; USER&#x60;: a single reviewer reviews at that level. The property &#x60;reviewerId&#x60; is required.   - &#x60;type &#x3D; REVIEWER_EXPRESSION&#x60;: specifies an expression to derive a reviewer for that level. The property &#x60;fallBackReviewerId&#x60; is required. Specify &#x60;fallBackReviewerId&#x60;, if Access Certifications can&#39;t determine the reviewer using the expression. In such a case, the fallback reviewer specified through &#x60;fallBackReviewerId&#x60; becomes automatic reviewer at that level.   - &#x60;type &#x3D; GROUP&#x60;: all the members of an Okta group specified through &#x60;reviewerGroupId&#x60; become reviewers at that level. The property &#x60;reviewerGroupId&#x60; is required.     Note that, if the group specified through &#x60;reviewerGroupId&#x60; has only one member, then that member will be assigned as a reviewer to all reviewers and &#x60;type&#x60; at that level changes to &#x60;USER&#x60;.     Additionally, if the provided Okta group has more than 10 members, when the campaign gets launched, only a maximum of 10 members will be pulled from the group.     Those 10 members would be the reviewers at that level. When fetching members of the group, there is no order guaranteed.   - &#x60;type &#x3D; RESOURCE_OWNER&#x60;: all the owners of an Okta group become reviewers at that level. The property &#x60;fallBackReviewerId&#x60; is required. For &#x60;RESOURCE_OWNER&#x60; to work, the property &#x60;resourceSettings.type&#x60; should be &#x60;GROUP&#x60; and &#x60;resourceSettings.targetResources&#x60; should be an Okta group ID. Access Certifications derives all the owners of the groups specified through &#x60;resourceSettings.targetResources.resourceId&#x60; and assigns them as reviewers to respective reviews of the campaign. Note that, if there is only one owner for such group(s) then that owner will be assigned as a reviewer to respective reviews and type in those reviews changes to &#x60;USER&#x60;. Specify &#x60;fallBackReviewerId&#x60;, if Access Certifications can&#39;t determine the reviewer using the group(s) specified through &#x60;resourceSettings.targetResources.resourceId&#x60;. In such a case, the fallback reviewer specified through &#x60;fallBackReviewerId&#x60; becomes automatic reviewer of the campaign.     Note that, if the Okta group specified has only one member, then that owner will be assigned as a reviewer to all reviewers and &#x60;type&#x60; at that level changes to &#x60;USER&#x60;.     Additionally, if the provided Okta group has more than 10 owners, when the campaign gets launched, only a maximum of 10 owners will be pulled from the group.     Those 10 owners would be the reviewers at that level. When fetching owners of the group, there is no order guaranteed.  Additionally -   - The property &#x60;startReview.onDay&#x60; indicates from which day, the level should start. This property will be &#x60;0&#x60; for &#x60;FIRST&#x60; reviewer level, as the first level will start when the campaign starts. For &#x60;SECOND&#x60; reviewer level specify a value greater than &#x60;0&#x60;.   - The property &#x60;startReview.when&#x60; indicates the condition when the lower level reviews will move to that level. This property should not be set for &#x60;FIRST&#x60; reviewer level. Only &#x60;SECOND&#x60; reviewer level will have the appropriate condition.  | [optional] 

## Methods

### NewReviewerSettingsMutable

`func NewReviewerSettingsMutable(type_ CampaignReviewerType, ) *ReviewerSettingsMutable`

NewReviewerSettingsMutable instantiates a new ReviewerSettingsMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerSettingsMutableWithDefaults

`func NewReviewerSettingsMutableWithDefaults() *ReviewerSettingsMutable`

NewReviewerSettingsMutableWithDefaults instantiates a new ReviewerSettingsMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReviewerSettingsMutable) GetType() CampaignReviewerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReviewerSettingsMutable) GetTypeOk() (*CampaignReviewerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReviewerSettingsMutable) SetType(v CampaignReviewerType)`

SetType sets Type field to given value.


### GetReviewerId

`func (o *ReviewerSettingsMutable) GetReviewerId() string`

GetReviewerId returns the ReviewerId field if non-nil, zero value otherwise.

### GetReviewerIdOk

`func (o *ReviewerSettingsMutable) GetReviewerIdOk() (*string, bool)`

GetReviewerIdOk returns a tuple with the ReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerId

`func (o *ReviewerSettingsMutable) SetReviewerId(v string)`

SetReviewerId sets ReviewerId field to given value.

### HasReviewerId

`func (o *ReviewerSettingsMutable) HasReviewerId() bool`

HasReviewerId returns a boolean if a field has been set.

### GetReviewerScopeExpression

`func (o *ReviewerSettingsMutable) GetReviewerScopeExpression() string`

GetReviewerScopeExpression returns the ReviewerScopeExpression field if non-nil, zero value otherwise.

### GetReviewerScopeExpressionOk

`func (o *ReviewerSettingsMutable) GetReviewerScopeExpressionOk() (*string, bool)`

GetReviewerScopeExpressionOk returns a tuple with the ReviewerScopeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerScopeExpression

`func (o *ReviewerSettingsMutable) SetReviewerScopeExpression(v string)`

SetReviewerScopeExpression sets ReviewerScopeExpression field to given value.

### HasReviewerScopeExpression

`func (o *ReviewerSettingsMutable) HasReviewerScopeExpression() bool`

HasReviewerScopeExpression returns a boolean if a field has been set.

### GetFallBackReviewerId

`func (o *ReviewerSettingsMutable) GetFallBackReviewerId() string`

GetFallBackReviewerId returns the FallBackReviewerId field if non-nil, zero value otherwise.

### GetFallBackReviewerIdOk

`func (o *ReviewerSettingsMutable) GetFallBackReviewerIdOk() (*string, bool)`

GetFallBackReviewerIdOk returns a tuple with the FallBackReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallBackReviewerId

`func (o *ReviewerSettingsMutable) SetFallBackReviewerId(v string)`

SetFallBackReviewerId sets FallBackReviewerId field to given value.

### HasFallBackReviewerId

`func (o *ReviewerSettingsMutable) HasFallBackReviewerId() bool`

HasFallBackReviewerId returns a boolean if a field has been set.

### GetReviewerGroupId

`func (o *ReviewerSettingsMutable) GetReviewerGroupId() string`

GetReviewerGroupId returns the ReviewerGroupId field if non-nil, zero value otherwise.

### GetReviewerGroupIdOk

`func (o *ReviewerSettingsMutable) GetReviewerGroupIdOk() (*string, bool)`

GetReviewerGroupIdOk returns a tuple with the ReviewerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupId

`func (o *ReviewerSettingsMutable) SetReviewerGroupId(v string)`

SetReviewerGroupId sets ReviewerGroupId field to given value.

### HasReviewerGroupId

`func (o *ReviewerSettingsMutable) HasReviewerGroupId() bool`

HasReviewerGroupId returns a boolean if a field has been set.

### GetIsSelfReviewDisabled

`func (o *ReviewerSettingsMutable) GetIsSelfReviewDisabled() bool`

GetIsSelfReviewDisabled returns the IsSelfReviewDisabled field if non-nil, zero value otherwise.

### GetIsSelfReviewDisabledOk

`func (o *ReviewerSettingsMutable) GetIsSelfReviewDisabledOk() (*bool, bool)`

GetIsSelfReviewDisabledOk returns a tuple with the IsSelfReviewDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfReviewDisabled

`func (o *ReviewerSettingsMutable) SetIsSelfReviewDisabled(v bool)`

SetIsSelfReviewDisabled sets IsSelfReviewDisabled field to given value.

### HasIsSelfReviewDisabled

`func (o *ReviewerSettingsMutable) HasIsSelfReviewDisabled() bool`

HasIsSelfReviewDisabled returns a boolean if a field has been set.

### GetSelfReviewDisabled

`func (o *ReviewerSettingsMutable) GetSelfReviewDisabled() bool`

GetSelfReviewDisabled returns the SelfReviewDisabled field if non-nil, zero value otherwise.

### GetSelfReviewDisabledOk

`func (o *ReviewerSettingsMutable) GetSelfReviewDisabledOk() (*bool, bool)`

GetSelfReviewDisabledOk returns a tuple with the SelfReviewDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfReviewDisabled

`func (o *ReviewerSettingsMutable) SetSelfReviewDisabled(v bool)`

SetSelfReviewDisabled sets SelfReviewDisabled field to given value.

### HasSelfReviewDisabled

`func (o *ReviewerSettingsMutable) HasSelfReviewDisabled() bool`

HasSelfReviewDisabled returns a boolean if a field has been set.

### GetJustificationRequired

`func (o *ReviewerSettingsMutable) GetJustificationRequired() bool`

GetJustificationRequired returns the JustificationRequired field if non-nil, zero value otherwise.

### GetJustificationRequiredOk

`func (o *ReviewerSettingsMutable) GetJustificationRequiredOk() (*bool, bool)`

GetJustificationRequiredOk returns a tuple with the JustificationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustificationRequired

`func (o *ReviewerSettingsMutable) SetJustificationRequired(v bool)`

SetJustificationRequired sets JustificationRequired field to given value.

### HasJustificationRequired

`func (o *ReviewerSettingsMutable) HasJustificationRequired() bool`

HasJustificationRequired returns a boolean if a field has been set.

### GetBulkDecisionDisabled

`func (o *ReviewerSettingsMutable) GetBulkDecisionDisabled() bool`

GetBulkDecisionDisabled returns the BulkDecisionDisabled field if non-nil, zero value otherwise.

### GetBulkDecisionDisabledOk

`func (o *ReviewerSettingsMutable) GetBulkDecisionDisabledOk() (*bool, bool)`

GetBulkDecisionDisabledOk returns a tuple with the BulkDecisionDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkDecisionDisabled

`func (o *ReviewerSettingsMutable) SetBulkDecisionDisabled(v bool)`

SetBulkDecisionDisabled sets BulkDecisionDisabled field to given value.

### HasBulkDecisionDisabled

`func (o *ReviewerSettingsMutable) HasBulkDecisionDisabled() bool`

HasBulkDecisionDisabled returns a boolean if a field has been set.

### GetReassignmentDisabled

`func (o *ReviewerSettingsMutable) GetReassignmentDisabled() bool`

GetReassignmentDisabled returns the ReassignmentDisabled field if non-nil, zero value otherwise.

### GetReassignmentDisabledOk

`func (o *ReviewerSettingsMutable) GetReassignmentDisabledOk() (*bool, bool)`

GetReassignmentDisabledOk returns a tuple with the ReassignmentDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignmentDisabled

`func (o *ReviewerSettingsMutable) SetReassignmentDisabled(v bool)`

SetReassignmentDisabled sets ReassignmentDisabled field to given value.

### HasReassignmentDisabled

`func (o *ReviewerSettingsMutable) HasReassignmentDisabled() bool`

HasReassignmentDisabled returns a boolean if a field has been set.

### GetReviewerLevels

`func (o *ReviewerSettingsMutable) GetReviewerLevels() []ReviewerLevelSettingsMutable`

GetReviewerLevels returns the ReviewerLevels field if non-nil, zero value otherwise.

### GetReviewerLevelsOk

`func (o *ReviewerSettingsMutable) GetReviewerLevelsOk() (*[]ReviewerLevelSettingsMutable, bool)`

GetReviewerLevelsOk returns a tuple with the ReviewerLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerLevels

`func (o *ReviewerSettingsMutable) SetReviewerLevels(v []ReviewerLevelSettingsMutable)`

SetReviewerLevels sets ReviewerLevels field to given value.

### HasReviewerLevels

`func (o *ReviewerSettingsMutable) HasReviewerLevels() bool`

HasReviewerLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


