# ReviewerSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CampaignReviewerType**](CampaignReviewerType.md) |  | 
**ReviewerId** | Pointer to **string** | Required if &#x60;reviewerSettings.type&#x60; is &#x60;USER&#x60;.  The ID of an Okta user to assign all reviews. | [optional] 
**ReviewerScopeExpression** | Pointer to **string** | Required when &#x60;reviewerSettings.type&#x60; is &#x60;REVIEWER_EXPRESSION&#x60;.  The Okta-specific user expression to fetch the reviewers from a connected identity source:  * If a user is found with the provided expression, that user is assigned as the reviewer. * If a user isn&#39;t found with the provided expression, the &#x60;reviewerSettings.fallBackReviewerId&#x60; user is assigned as the reviewer. * See [Okta Expression Language (EL)](https://developer.okta.com/docs/reference/okta-expression-language/#okta-user-profile) to build the expression base on user profile attributes.  | [optional] 
**FallBackReviewerId** | Pointer to **string** | Required when reviewer setting &#x60;type&#x60; is &#x60;REVIEWER_EXPRESSION&#x60; or &#x60;RESOURCE_OWNER&#x60;. The fallback reviewer is assigned as the reviewer if:  * reviewer setting &#x60;reviewerScopeExpression&#x60; fails to identify reviewers, or * reviewers aren&#39;t identified through resource owners  | [optional] 
**ReviewerGroupId** | Pointer to **string** | Required when &#x60;reviewerSettings.type&#x60; is &#x60;GROUP&#x60;.  The &#x60;id&#x60; of the Okta group: * All members of the specified group are assigned as reviewers. * Use this reviewer group assignment if you can&#39;t use &#x60;reviewerId&#x60; or &#x60;reviewerScopeExpression&#x60;. * If the Okta group has more than 10 members when the campaign launches, only 10 members are randomly selected as reviewers. * If the Okta group has only one member, then that member is assigned as the reviewer for all reviews, and &#x60;reviewerType&#x60; is set to &#x60;USER&#x60; for those reviews.  | [optional] 
**IsSelfReviewDisabled** | Pointer to **bool** | If &#x60;true&#x60;, users can&#39;t review their own review items.  &gt; **Note:** This field is deprecated. Use [&#39;selfReviewDisabled&#39;](/openapi/governance.api/tag/Campaigns/#tag/Campaigns/operation/createCampaign!path&#x3D;reviewerSettings/selfReviewDisabled&amp;t&#x3D;request)  | [optional] 
**SelfReviewDisabled** | Pointer to **bool** | If true, users won&#39;t be able to review their own review items.  This property is required to be &#x60;true&#x60; for resource-centric campaigns when the Okta Admin Console is one of the resources.  | [optional] 
**JustificationRequired** | Pointer to **bool** | If true, a justification is required when review items are approved or revoked.  This property must be &#x60;true&#x60; for resource-centric campaigns that have the Okta Admin Console as one of the resources.  | [optional] 
**BulkDecisionDisabled** | Pointer to **bool** | If true, bulk actions are disabled for approving or revoking review items.  | [optional] 
**ReassignmentDisabled** | Pointer to **bool** | If true, reassignment is disabled for reviewers.  | [optional] 
**ReviewerLevels** | Pointer to [**[]ReviewerLevelSettingsMutable**](ReviewerLevelSettingsMutable.md) | Defines the reviewer level in a campaign. A campaign can have a maximum of two reviewer levels.  | [optional] 

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


