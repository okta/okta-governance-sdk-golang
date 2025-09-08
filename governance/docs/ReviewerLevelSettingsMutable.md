# ReviewerLevelSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ReviewerType**](ReviewerType.md) |  | 
**StartReview** | [**ReviewerLevelStartReview**](ReviewerLevelStartReview.md) |  | 
**ReviewerId** | Pointer to **string** | This property is required when &#x60;type&#x3D;USER&#x60;  The &#x60;id&#x60; of an Okta user who will be assigned as a reviewer. By specifying &#x60;reviewerId&#x60;, all reviews will be assigned only to that reviewer. This is a conditional property to be provided only if the &#x60;reviewerScopeExpression&#x60; can&#39;t be specified.  | [optional] 
**ReviewerScopeExpression** | Pointer to **string** | This property is required when &#x60;type&#x3D;REVIEWER_EXPRESSION&#x60;  The Okta specific user expression to fetch the reviewers from a connected identity source.  If a user is found using the provided expression, that user is assigned as a reviewer.  If a user is not found using the provided expression, the &#x60;fallbackReviewerId&#x60; is used.  | [optional] 
**FallBackReviewerId** | Pointer to **string** | Required when the &#x60;type&#x3D;REVIEWER_EXPRESSION&#x60; or &#x60;type&#x3D;RESOURCE_OWNER&#x60;  If the reviewer(s) can&#39;t be identified using &#x60;reviewerScopeExpression&#x60; or using the owners of groups (the owners of the group identified through the property &#x60;resourceSettings.type &#x3D; GROUP&#x60;), the fallback reviewer (&#x60;id&#x60; of an Okta user) is assigned as a reviewer.  | [optional] 
**ReviewerGroupId** | Pointer to **string** | This property is required when &#x60;type&#x3D;GROUP&#x60;  The &#x60;id&#x60; of an Okta group who will be assigned as a reviewer. By specifying &#x60;reviewerGroupId&#x60;, all reviews will be assigned to that group. Any user part of that group will automatically be the reviewers of the campaign. This is a conditional property to be provided, only when one wants to assign more than one reviewer and cannot use &#x60;reviewerId&#x60; or &#x60;reviewerScopeExpression&#x60;.  Note that, if the provided Okta group has more than 10 members, when the campaign gets launched, only a max of 10 members will be pulled from the group. Those 10 members would be the reviewers of the campaign. When fetching members of the group, there is no order guaranteed.  Additionally, if the Okta group specified has only one member, then that member will be assigned as a reviewer to all reviewers and &#x60;type&#x60; for those reviews changes to &#x60;USER&#x60;.  | [optional] 
**IsSelfReviewDisabled** | Pointer to **bool** | if &#39;true&#39;, users will not able to review their own review items  &gt; **Note:** This field is deprecated. Use [&#39;selfReviewDisabled&#39;](/openapi/governance.api/tag/Campaigns/#tag/Campaigns/operation/createCampaign!path&#x3D;reviewerSettings/reviewerLevels/selfReviewDisabled&amp;t&#x3D;request)  | [optional] 
**SelfReviewDisabled** | Pointer to **bool** | If &#x60;true&#x60;, users aren&#39;t able to review their own review items.  This property is required to be &#x60;true&#x60; for resource-centric campaigns when the Okta Admin Console is one of the resources  | [optional] 

## Methods

### NewReviewerLevelSettingsMutable

`func NewReviewerLevelSettingsMutable(type_ ReviewerType, startReview ReviewerLevelStartReview, ) *ReviewerLevelSettingsMutable`

NewReviewerLevelSettingsMutable instantiates a new ReviewerLevelSettingsMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerLevelSettingsMutableWithDefaults

`func NewReviewerLevelSettingsMutableWithDefaults() *ReviewerLevelSettingsMutable`

NewReviewerLevelSettingsMutableWithDefaults instantiates a new ReviewerLevelSettingsMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReviewerLevelSettingsMutable) GetType() ReviewerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReviewerLevelSettingsMutable) GetTypeOk() (*ReviewerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReviewerLevelSettingsMutable) SetType(v ReviewerType)`

SetType sets Type field to given value.


### GetStartReview

`func (o *ReviewerLevelSettingsMutable) GetStartReview() ReviewerLevelStartReview`

GetStartReview returns the StartReview field if non-nil, zero value otherwise.

### GetStartReviewOk

`func (o *ReviewerLevelSettingsMutable) GetStartReviewOk() (*ReviewerLevelStartReview, bool)`

GetStartReviewOk returns a tuple with the StartReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartReview

`func (o *ReviewerLevelSettingsMutable) SetStartReview(v ReviewerLevelStartReview)`

SetStartReview sets StartReview field to given value.


### GetReviewerId

`func (o *ReviewerLevelSettingsMutable) GetReviewerId() string`

GetReviewerId returns the ReviewerId field if non-nil, zero value otherwise.

### GetReviewerIdOk

`func (o *ReviewerLevelSettingsMutable) GetReviewerIdOk() (*string, bool)`

GetReviewerIdOk returns a tuple with the ReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerId

`func (o *ReviewerLevelSettingsMutable) SetReviewerId(v string)`

SetReviewerId sets ReviewerId field to given value.

### HasReviewerId

`func (o *ReviewerLevelSettingsMutable) HasReviewerId() bool`

HasReviewerId returns a boolean if a field has been set.

### GetReviewerScopeExpression

`func (o *ReviewerLevelSettingsMutable) GetReviewerScopeExpression() string`

GetReviewerScopeExpression returns the ReviewerScopeExpression field if non-nil, zero value otherwise.

### GetReviewerScopeExpressionOk

`func (o *ReviewerLevelSettingsMutable) GetReviewerScopeExpressionOk() (*string, bool)`

GetReviewerScopeExpressionOk returns a tuple with the ReviewerScopeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerScopeExpression

`func (o *ReviewerLevelSettingsMutable) SetReviewerScopeExpression(v string)`

SetReviewerScopeExpression sets ReviewerScopeExpression field to given value.

### HasReviewerScopeExpression

`func (o *ReviewerLevelSettingsMutable) HasReviewerScopeExpression() bool`

HasReviewerScopeExpression returns a boolean if a field has been set.

### GetFallBackReviewerId

`func (o *ReviewerLevelSettingsMutable) GetFallBackReviewerId() string`

GetFallBackReviewerId returns the FallBackReviewerId field if non-nil, zero value otherwise.

### GetFallBackReviewerIdOk

`func (o *ReviewerLevelSettingsMutable) GetFallBackReviewerIdOk() (*string, bool)`

GetFallBackReviewerIdOk returns a tuple with the FallBackReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallBackReviewerId

`func (o *ReviewerLevelSettingsMutable) SetFallBackReviewerId(v string)`

SetFallBackReviewerId sets FallBackReviewerId field to given value.

### HasFallBackReviewerId

`func (o *ReviewerLevelSettingsMutable) HasFallBackReviewerId() bool`

HasFallBackReviewerId returns a boolean if a field has been set.

### GetReviewerGroupId

`func (o *ReviewerLevelSettingsMutable) GetReviewerGroupId() string`

GetReviewerGroupId returns the ReviewerGroupId field if non-nil, zero value otherwise.

### GetReviewerGroupIdOk

`func (o *ReviewerLevelSettingsMutable) GetReviewerGroupIdOk() (*string, bool)`

GetReviewerGroupIdOk returns a tuple with the ReviewerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupId

`func (o *ReviewerLevelSettingsMutable) SetReviewerGroupId(v string)`

SetReviewerGroupId sets ReviewerGroupId field to given value.

### HasReviewerGroupId

`func (o *ReviewerLevelSettingsMutable) HasReviewerGroupId() bool`

HasReviewerGroupId returns a boolean if a field has been set.

### GetIsSelfReviewDisabled

`func (o *ReviewerLevelSettingsMutable) GetIsSelfReviewDisabled() bool`

GetIsSelfReviewDisabled returns the IsSelfReviewDisabled field if non-nil, zero value otherwise.

### GetIsSelfReviewDisabledOk

`func (o *ReviewerLevelSettingsMutable) GetIsSelfReviewDisabledOk() (*bool, bool)`

GetIsSelfReviewDisabledOk returns a tuple with the IsSelfReviewDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfReviewDisabled

`func (o *ReviewerLevelSettingsMutable) SetIsSelfReviewDisabled(v bool)`

SetIsSelfReviewDisabled sets IsSelfReviewDisabled field to given value.

### HasIsSelfReviewDisabled

`func (o *ReviewerLevelSettingsMutable) HasIsSelfReviewDisabled() bool`

HasIsSelfReviewDisabled returns a boolean if a field has been set.

### GetSelfReviewDisabled

`func (o *ReviewerLevelSettingsMutable) GetSelfReviewDisabled() bool`

GetSelfReviewDisabled returns the SelfReviewDisabled field if non-nil, zero value otherwise.

### GetSelfReviewDisabledOk

`func (o *ReviewerLevelSettingsMutable) GetSelfReviewDisabledOk() (*bool, bool)`

GetSelfReviewDisabledOk returns a tuple with the SelfReviewDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfReviewDisabled

`func (o *ReviewerLevelSettingsMutable) SetSelfReviewDisabled(v bool)`

SetSelfReviewDisabled sets SelfReviewDisabled field to given value.

### HasSelfReviewDisabled

`func (o *ReviewerLevelSettingsMutable) HasSelfReviewDisabled() bool`

HasSelfReviewDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


