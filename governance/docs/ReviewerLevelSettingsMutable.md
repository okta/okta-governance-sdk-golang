# ReviewerLevelSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ReviewerType**](ReviewerType.md) |  | 
**StartReview** | [**ReviewerLevelStartReview**](ReviewerLevelStartReview.md) |  | 
**ReviewerId** | Pointer to **string** | This property is required when &#x60;reviewerSettings.reviewerLevels.type&#x60; is &#x60;USER&#x60;  The &#x60;id&#x60; of an Okta user to assign all reviews. Use this reviewer type when you can&#39;t specify &#x60;reviewerScopeExpression&#x60;.  | [optional] 
**ReviewerScopeExpression** | Pointer to **string** | Required when &#x60;reviewerSettings.reviewerLevels.type&#x60; is &#x60;REVIEWER_EXPRESSION&#x60;.  The Okta-specific user expression to fetch the reviewers from a connected identity source:  * If a user is found with the provided expression, that user is assigned as the reviewer. * If a user isn&#39;t found with the provided expression, the &#x60;fallBackReviewerId&#x60; user is assigned as the reviewer. * See [Okta Expression Language (EL)](https://developer.okta.com/docs/reference/okta-expression-language/#okta-user-profile) to build the expression based on user profile attributes. | [optional] 
**FallBackReviewerId** | Pointer to **string** | Required when reviewer setting &#x60;type&#x60; is &#x60;REVIEWER_EXPRESSION&#x60; or &#x60;RESOURCE_OWNER&#x60;. The fallback reviewer is assigned as the reviewer if:  * reviewer setting &#x60;reviewerScopeExpression&#x60; fails to identify reviewers, or * reviewers aren&#39;t identified through resource owners  | [optional] 
**ReviewerGroupId** | Pointer to **string** | Required when &#x60;reviewerSettings.reviewerLevels.type&#x60; is &#x60;GROUP&#x60;.  The &#x60;id&#x60; of the Okta group: * All members of the specified group are assigned as reviewers. * Use this reviewer group assignment to assign more than one reviewer if you can&#39;t use &#x60;reviewerId&#x60; or &#x60;reviewerScopeExpression&#x60;. * If the Okta group has more than 10 members when the campaign launches, only 10 members are randomly selected as reviewers. * If the Okta group has only one member, then that member is assigned as the reviewer for all reviews, and &#x60;reviewerType&#x60; is set to &#x60;USER&#x60; for those reviews.  | [optional] 
**IsSelfReviewDisabled** | Pointer to **bool** | If &#x60;true&#x60;, users can&#39;t review their own review items.  &gt; **Note:** This field is deprecated. Use &#39;selfReviewDisabled&#39;.  | [optional] 
**SelfReviewDisabled** | Pointer to **bool** | If &#x60;true&#x60;, users can&#39;t review their own review items.  This property must be &#x60;true&#x60; for resource-centric campaigns when the Okta Admin Console is one of the resources.  | [optional] 

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


