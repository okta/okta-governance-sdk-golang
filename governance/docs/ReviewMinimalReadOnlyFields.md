# ReviewMinimalReadOnlyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **string** |  | 
**ResourceId** | **string** |  | 
**EntitlementValue** | Pointer to [**ReviewerEntitlementValue**](ReviewerEntitlementValue.md) |  | [optional] 
**EntitlementBundle** | Pointer to [**ReviewerEntitlementBundle**](ReviewerEntitlementBundle.md) |  | [optional] 
**Decision** | [**Decision**](Decision.md) |  | 
**Decided** | Pointer to **time.Time** |  | [optional] 
**RemediationStatus** | [**RemediationStatus**](RemediationStatus.md) |  | 
**PrincipalProfile** | [**PrincipalProfile**](PrincipalProfile.md) |  | 
**ReviewerProfile** | Pointer to [**PrincipalProfile**](PrincipalProfile.md) |  | [optional] 
**ReviewerType** | [**ReviewersReviewerType**](ReviewersReviewerType.md) |  | 
**ReviewerGroupProfile** | Pointer to [**ReviewerGroupProfile**](ReviewerGroupProfile.md) |  | [optional] 
**CurrentReviewerLevel** | Pointer to [**ReviewerLevelType**](ReviewerLevelType.md) |  | [optional] 
**Links** | [**ReviewLinks**](ReviewLinks.md) |  | 
**RiskRuleConflicts** | Pointer to [**[]RiskRuleConflicts**](RiskRuleConflicts.md) | List of risk rule conflicts caused by this entitlement value. Only applies to review item that has entitlement value. | [optional] 

## Methods

### NewReviewMinimalReadOnlyFields

`func NewReviewMinimalReadOnlyFields(campaignId string, resourceId string, decision Decision, remediationStatus RemediationStatus, principalProfile PrincipalProfile, reviewerType ReviewersReviewerType, links ReviewLinks, ) *ReviewMinimalReadOnlyFields`

NewReviewMinimalReadOnlyFields instantiates a new ReviewMinimalReadOnlyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewMinimalReadOnlyFieldsWithDefaults

`func NewReviewMinimalReadOnlyFieldsWithDefaults() *ReviewMinimalReadOnlyFields`

NewReviewMinimalReadOnlyFieldsWithDefaults instantiates a new ReviewMinimalReadOnlyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *ReviewMinimalReadOnlyFields) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ReviewMinimalReadOnlyFields) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ReviewMinimalReadOnlyFields) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetResourceId

`func (o *ReviewMinimalReadOnlyFields) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ReviewMinimalReadOnlyFields) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ReviewMinimalReadOnlyFields) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetEntitlementValue

`func (o *ReviewMinimalReadOnlyFields) GetEntitlementValue() ReviewerEntitlementValue`

GetEntitlementValue returns the EntitlementValue field if non-nil, zero value otherwise.

### GetEntitlementValueOk

`func (o *ReviewMinimalReadOnlyFields) GetEntitlementValueOk() (*ReviewerEntitlementValue, bool)`

GetEntitlementValueOk returns a tuple with the EntitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValue

`func (o *ReviewMinimalReadOnlyFields) SetEntitlementValue(v ReviewerEntitlementValue)`

SetEntitlementValue sets EntitlementValue field to given value.

### HasEntitlementValue

`func (o *ReviewMinimalReadOnlyFields) HasEntitlementValue() bool`

HasEntitlementValue returns a boolean if a field has been set.

### GetEntitlementBundle

`func (o *ReviewMinimalReadOnlyFields) GetEntitlementBundle() ReviewerEntitlementBundle`

GetEntitlementBundle returns the EntitlementBundle field if non-nil, zero value otherwise.

### GetEntitlementBundleOk

`func (o *ReviewMinimalReadOnlyFields) GetEntitlementBundleOk() (*ReviewerEntitlementBundle, bool)`

GetEntitlementBundleOk returns a tuple with the EntitlementBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundle

`func (o *ReviewMinimalReadOnlyFields) SetEntitlementBundle(v ReviewerEntitlementBundle)`

SetEntitlementBundle sets EntitlementBundle field to given value.

### HasEntitlementBundle

`func (o *ReviewMinimalReadOnlyFields) HasEntitlementBundle() bool`

HasEntitlementBundle returns a boolean if a field has been set.

### GetDecision

`func (o *ReviewMinimalReadOnlyFields) GetDecision() Decision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ReviewMinimalReadOnlyFields) GetDecisionOk() (*Decision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ReviewMinimalReadOnlyFields) SetDecision(v Decision)`

SetDecision sets Decision field to given value.


### GetDecided

`func (o *ReviewMinimalReadOnlyFields) GetDecided() time.Time`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *ReviewMinimalReadOnlyFields) GetDecidedOk() (*time.Time, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *ReviewMinimalReadOnlyFields) SetDecided(v time.Time)`

SetDecided sets Decided field to given value.

### HasDecided

`func (o *ReviewMinimalReadOnlyFields) HasDecided() bool`

HasDecided returns a boolean if a field has been set.

### GetRemediationStatus

`func (o *ReviewMinimalReadOnlyFields) GetRemediationStatus() RemediationStatus`

GetRemediationStatus returns the RemediationStatus field if non-nil, zero value otherwise.

### GetRemediationStatusOk

`func (o *ReviewMinimalReadOnlyFields) GetRemediationStatusOk() (*RemediationStatus, bool)`

GetRemediationStatusOk returns a tuple with the RemediationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationStatus

`func (o *ReviewMinimalReadOnlyFields) SetRemediationStatus(v RemediationStatus)`

SetRemediationStatus sets RemediationStatus field to given value.


### GetPrincipalProfile

`func (o *ReviewMinimalReadOnlyFields) GetPrincipalProfile() PrincipalProfile`

GetPrincipalProfile returns the PrincipalProfile field if non-nil, zero value otherwise.

### GetPrincipalProfileOk

`func (o *ReviewMinimalReadOnlyFields) GetPrincipalProfileOk() (*PrincipalProfile, bool)`

GetPrincipalProfileOk returns a tuple with the PrincipalProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalProfile

`func (o *ReviewMinimalReadOnlyFields) SetPrincipalProfile(v PrincipalProfile)`

SetPrincipalProfile sets PrincipalProfile field to given value.


### GetReviewerProfile

`func (o *ReviewMinimalReadOnlyFields) GetReviewerProfile() PrincipalProfile`

GetReviewerProfile returns the ReviewerProfile field if non-nil, zero value otherwise.

### GetReviewerProfileOk

`func (o *ReviewMinimalReadOnlyFields) GetReviewerProfileOk() (*PrincipalProfile, bool)`

GetReviewerProfileOk returns a tuple with the ReviewerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerProfile

`func (o *ReviewMinimalReadOnlyFields) SetReviewerProfile(v PrincipalProfile)`

SetReviewerProfile sets ReviewerProfile field to given value.

### HasReviewerProfile

`func (o *ReviewMinimalReadOnlyFields) HasReviewerProfile() bool`

HasReviewerProfile returns a boolean if a field has been set.

### GetReviewerType

`func (o *ReviewMinimalReadOnlyFields) GetReviewerType() ReviewersReviewerType`

GetReviewerType returns the ReviewerType field if non-nil, zero value otherwise.

### GetReviewerTypeOk

`func (o *ReviewMinimalReadOnlyFields) GetReviewerTypeOk() (*ReviewersReviewerType, bool)`

GetReviewerTypeOk returns a tuple with the ReviewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerType

`func (o *ReviewMinimalReadOnlyFields) SetReviewerType(v ReviewersReviewerType)`

SetReviewerType sets ReviewerType field to given value.


### GetReviewerGroupProfile

`func (o *ReviewMinimalReadOnlyFields) GetReviewerGroupProfile() ReviewerGroupProfile`

GetReviewerGroupProfile returns the ReviewerGroupProfile field if non-nil, zero value otherwise.

### GetReviewerGroupProfileOk

`func (o *ReviewMinimalReadOnlyFields) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool)`

GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupProfile

`func (o *ReviewMinimalReadOnlyFields) SetReviewerGroupProfile(v ReviewerGroupProfile)`

SetReviewerGroupProfile sets ReviewerGroupProfile field to given value.

### HasReviewerGroupProfile

`func (o *ReviewMinimalReadOnlyFields) HasReviewerGroupProfile() bool`

HasReviewerGroupProfile returns a boolean if a field has been set.

### GetCurrentReviewerLevel

`func (o *ReviewMinimalReadOnlyFields) GetCurrentReviewerLevel() ReviewerLevelType`

GetCurrentReviewerLevel returns the CurrentReviewerLevel field if non-nil, zero value otherwise.

### GetCurrentReviewerLevelOk

`func (o *ReviewMinimalReadOnlyFields) GetCurrentReviewerLevelOk() (*ReviewerLevelType, bool)`

GetCurrentReviewerLevelOk returns a tuple with the CurrentReviewerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReviewerLevel

`func (o *ReviewMinimalReadOnlyFields) SetCurrentReviewerLevel(v ReviewerLevelType)`

SetCurrentReviewerLevel sets CurrentReviewerLevel field to given value.

### HasCurrentReviewerLevel

`func (o *ReviewMinimalReadOnlyFields) HasCurrentReviewerLevel() bool`

HasCurrentReviewerLevel returns a boolean if a field has been set.

### GetLinks

`func (o *ReviewMinimalReadOnlyFields) GetLinks() ReviewLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReviewMinimalReadOnlyFields) GetLinksOk() (*ReviewLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReviewMinimalReadOnlyFields) SetLinks(v ReviewLinks)`

SetLinks sets Links field to given value.


### GetRiskRuleConflicts

`func (o *ReviewMinimalReadOnlyFields) GetRiskRuleConflicts() []RiskRuleConflicts`

GetRiskRuleConflicts returns the RiskRuleConflicts field if non-nil, zero value otherwise.

### GetRiskRuleConflictsOk

`func (o *ReviewMinimalReadOnlyFields) GetRiskRuleConflictsOk() (*[]RiskRuleConflicts, bool)`

GetRiskRuleConflictsOk returns a tuple with the RiskRuleConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRuleConflicts

`func (o *ReviewMinimalReadOnlyFields) SetRiskRuleConflicts(v []RiskRuleConflicts)`

SetRiskRuleConflicts sets RiskRuleConflicts field to given value.

### HasRiskRuleConflicts

`func (o *ReviewMinimalReadOnlyFields) HasRiskRuleConflicts() bool`

HasRiskRuleConflicts returns a boolean if a field has been set.

### SetRiskRuleConflictsNil

`func (o *ReviewMinimalReadOnlyFields) SetRiskRuleConflictsNil(b bool)`

 SetRiskRuleConflictsNil sets the value for RiskRuleConflicts to be an explicit nil

### UnsetRiskRuleConflicts
`func (o *ReviewMinimalReadOnlyFields) UnsetRiskRuleConflicts()`

UnsetRiskRuleConflicts ensures that no value is present for RiskRuleConflicts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


