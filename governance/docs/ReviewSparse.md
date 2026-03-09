# ReviewSparse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**ReviewLinks**](ReviewLinks.md) |  | 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**CampaignId** | **string** |  | 
**ResourceId** | **string** |  | 
**EntitlementValue** | Pointer to [**ReviewerEntitlementValue**](ReviewerEntitlementValue.md) |  | [optional] 
**EntitlementBundle** | Pointer to [**ReviewerEntitlementBundle**](ReviewerEntitlementBundle.md) |  | [optional] 
**Decision** | [**Decision**](Decision.md) |  | 
**Decided** | Pointer to **time.Time** |  | [optional] 
**RemediationStatus** | [**RemediationStatus**](RemediationStatus.md) |  | 
**PrincipalProfile** | [**PrincipalProfileEnriched**](PrincipalProfileEnriched.md) |  | 
**ReviewerProfile** | Pointer to [**PrincipalProfileEnriched**](PrincipalProfileEnriched.md) |  | [optional] 
**ReviewerType** | [**ReviewersReviewerType**](ReviewersReviewerType.md) |  | 
**ReviewerGroupProfile** | Pointer to [**ReviewerGroupProfile**](ReviewerGroupProfile.md) |  | [optional] 
**CurrentReviewerLevel** | Pointer to [**ReviewerLevelType**](ReviewerLevelType.md) |  | [optional] 
**RiskRuleConflicts** | Pointer to [**[]RiskRuleConflicts**](RiskRuleConflicts.md) | List of risk rule conflicts caused by this entitlement value. Only applies to review item that has entitlement value. | [optional] 
**DelegatorProfile** | Pointer to [**PrincipalProfileEnriched**](PrincipalProfileEnriched.md) |  | [optional] 
**Delegated** | Pointer to **bool** | Specifies if this review was delegated by the original reviewer based on their governance delegate settings | [optional] 
**AppServiceAccount** | Pointer to [**ReviewerServiceAccount**](ReviewerServiceAccount.md) |  | [optional] 
**OktaServiceAccount** | Pointer to [**ReviewerServiceAccount**](ReviewerServiceAccount.md) |  | [optional] 

## Methods

### NewReviewSparse

`func NewReviewSparse(links ReviewLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, campaignId string, resourceId string, decision Decision, remediationStatus RemediationStatus, principalProfile PrincipalProfileEnriched, reviewerType ReviewersReviewerType, ) *ReviewSparse`

NewReviewSparse instantiates a new ReviewSparse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewSparseWithDefaults

`func NewReviewSparseWithDefaults() *ReviewSparse`

NewReviewSparseWithDefaults instantiates a new ReviewSparse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ReviewSparse) GetLinks() ReviewLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReviewSparse) GetLinksOk() (*ReviewLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReviewSparse) SetLinks(v ReviewLinks)`

SetLinks sets Links field to given value.


### GetId

`func (o *ReviewSparse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewSparse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewSparse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ReviewSparse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ReviewSparse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ReviewSparse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *ReviewSparse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReviewSparse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReviewSparse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *ReviewSparse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ReviewSparse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ReviewSparse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *ReviewSparse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ReviewSparse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ReviewSparse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetCampaignId

`func (o *ReviewSparse) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ReviewSparse) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ReviewSparse) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetResourceId

`func (o *ReviewSparse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ReviewSparse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ReviewSparse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetEntitlementValue

`func (o *ReviewSparse) GetEntitlementValue() ReviewerEntitlementValue`

GetEntitlementValue returns the EntitlementValue field if non-nil, zero value otherwise.

### GetEntitlementValueOk

`func (o *ReviewSparse) GetEntitlementValueOk() (*ReviewerEntitlementValue, bool)`

GetEntitlementValueOk returns a tuple with the EntitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValue

`func (o *ReviewSparse) SetEntitlementValue(v ReviewerEntitlementValue)`

SetEntitlementValue sets EntitlementValue field to given value.

### HasEntitlementValue

`func (o *ReviewSparse) HasEntitlementValue() bool`

HasEntitlementValue returns a boolean if a field has been set.

### GetEntitlementBundle

`func (o *ReviewSparse) GetEntitlementBundle() ReviewerEntitlementBundle`

GetEntitlementBundle returns the EntitlementBundle field if non-nil, zero value otherwise.

### GetEntitlementBundleOk

`func (o *ReviewSparse) GetEntitlementBundleOk() (*ReviewerEntitlementBundle, bool)`

GetEntitlementBundleOk returns a tuple with the EntitlementBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundle

`func (o *ReviewSparse) SetEntitlementBundle(v ReviewerEntitlementBundle)`

SetEntitlementBundle sets EntitlementBundle field to given value.

### HasEntitlementBundle

`func (o *ReviewSparse) HasEntitlementBundle() bool`

HasEntitlementBundle returns a boolean if a field has been set.

### GetDecision

`func (o *ReviewSparse) GetDecision() Decision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ReviewSparse) GetDecisionOk() (*Decision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ReviewSparse) SetDecision(v Decision)`

SetDecision sets Decision field to given value.


### GetDecided

`func (o *ReviewSparse) GetDecided() time.Time`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *ReviewSparse) GetDecidedOk() (*time.Time, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *ReviewSparse) SetDecided(v time.Time)`

SetDecided sets Decided field to given value.

### HasDecided

`func (o *ReviewSparse) HasDecided() bool`

HasDecided returns a boolean if a field has been set.

### GetRemediationStatus

`func (o *ReviewSparse) GetRemediationStatus() RemediationStatus`

GetRemediationStatus returns the RemediationStatus field if non-nil, zero value otherwise.

### GetRemediationStatusOk

`func (o *ReviewSparse) GetRemediationStatusOk() (*RemediationStatus, bool)`

GetRemediationStatusOk returns a tuple with the RemediationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationStatus

`func (o *ReviewSparse) SetRemediationStatus(v RemediationStatus)`

SetRemediationStatus sets RemediationStatus field to given value.


### GetPrincipalProfile

`func (o *ReviewSparse) GetPrincipalProfile() PrincipalProfileEnriched`

GetPrincipalProfile returns the PrincipalProfile field if non-nil, zero value otherwise.

### GetPrincipalProfileOk

`func (o *ReviewSparse) GetPrincipalProfileOk() (*PrincipalProfileEnriched, bool)`

GetPrincipalProfileOk returns a tuple with the PrincipalProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalProfile

`func (o *ReviewSparse) SetPrincipalProfile(v PrincipalProfileEnriched)`

SetPrincipalProfile sets PrincipalProfile field to given value.


### GetReviewerProfile

`func (o *ReviewSparse) GetReviewerProfile() PrincipalProfileEnriched`

GetReviewerProfile returns the ReviewerProfile field if non-nil, zero value otherwise.

### GetReviewerProfileOk

`func (o *ReviewSparse) GetReviewerProfileOk() (*PrincipalProfileEnriched, bool)`

GetReviewerProfileOk returns a tuple with the ReviewerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerProfile

`func (o *ReviewSparse) SetReviewerProfile(v PrincipalProfileEnriched)`

SetReviewerProfile sets ReviewerProfile field to given value.

### HasReviewerProfile

`func (o *ReviewSparse) HasReviewerProfile() bool`

HasReviewerProfile returns a boolean if a field has been set.

### GetReviewerType

`func (o *ReviewSparse) GetReviewerType() ReviewersReviewerType`

GetReviewerType returns the ReviewerType field if non-nil, zero value otherwise.

### GetReviewerTypeOk

`func (o *ReviewSparse) GetReviewerTypeOk() (*ReviewersReviewerType, bool)`

GetReviewerTypeOk returns a tuple with the ReviewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerType

`func (o *ReviewSparse) SetReviewerType(v ReviewersReviewerType)`

SetReviewerType sets ReviewerType field to given value.


### GetReviewerGroupProfile

`func (o *ReviewSparse) GetReviewerGroupProfile() ReviewerGroupProfile`

GetReviewerGroupProfile returns the ReviewerGroupProfile field if non-nil, zero value otherwise.

### GetReviewerGroupProfileOk

`func (o *ReviewSparse) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool)`

GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupProfile

`func (o *ReviewSparse) SetReviewerGroupProfile(v ReviewerGroupProfile)`

SetReviewerGroupProfile sets ReviewerGroupProfile field to given value.

### HasReviewerGroupProfile

`func (o *ReviewSparse) HasReviewerGroupProfile() bool`

HasReviewerGroupProfile returns a boolean if a field has been set.

### GetCurrentReviewerLevel

`func (o *ReviewSparse) GetCurrentReviewerLevel() ReviewerLevelType`

GetCurrentReviewerLevel returns the CurrentReviewerLevel field if non-nil, zero value otherwise.

### GetCurrentReviewerLevelOk

`func (o *ReviewSparse) GetCurrentReviewerLevelOk() (*ReviewerLevelType, bool)`

GetCurrentReviewerLevelOk returns a tuple with the CurrentReviewerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReviewerLevel

`func (o *ReviewSparse) SetCurrentReviewerLevel(v ReviewerLevelType)`

SetCurrentReviewerLevel sets CurrentReviewerLevel field to given value.

### HasCurrentReviewerLevel

`func (o *ReviewSparse) HasCurrentReviewerLevel() bool`

HasCurrentReviewerLevel returns a boolean if a field has been set.

### GetRiskRuleConflicts

`func (o *ReviewSparse) GetRiskRuleConflicts() []RiskRuleConflicts`

GetRiskRuleConflicts returns the RiskRuleConflicts field if non-nil, zero value otherwise.

### GetRiskRuleConflictsOk

`func (o *ReviewSparse) GetRiskRuleConflictsOk() (*[]RiskRuleConflicts, bool)`

GetRiskRuleConflictsOk returns a tuple with the RiskRuleConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRuleConflicts

`func (o *ReviewSparse) SetRiskRuleConflicts(v []RiskRuleConflicts)`

SetRiskRuleConflicts sets RiskRuleConflicts field to given value.

### HasRiskRuleConflicts

`func (o *ReviewSparse) HasRiskRuleConflicts() bool`

HasRiskRuleConflicts returns a boolean if a field has been set.

### SetRiskRuleConflictsNil

`func (o *ReviewSparse) SetRiskRuleConflictsNil(b bool)`

 SetRiskRuleConflictsNil sets the value for RiskRuleConflicts to be an explicit nil

### UnsetRiskRuleConflicts
`func (o *ReviewSparse) UnsetRiskRuleConflicts()`

UnsetRiskRuleConflicts ensures that no value is present for RiskRuleConflicts, not even an explicit nil
### GetDelegatorProfile

`func (o *ReviewSparse) GetDelegatorProfile() PrincipalProfileEnriched`

GetDelegatorProfile returns the DelegatorProfile field if non-nil, zero value otherwise.

### GetDelegatorProfileOk

`func (o *ReviewSparse) GetDelegatorProfileOk() (*PrincipalProfileEnriched, bool)`

GetDelegatorProfileOk returns a tuple with the DelegatorProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorProfile

`func (o *ReviewSparse) SetDelegatorProfile(v PrincipalProfileEnriched)`

SetDelegatorProfile sets DelegatorProfile field to given value.

### HasDelegatorProfile

`func (o *ReviewSparse) HasDelegatorProfile() bool`

HasDelegatorProfile returns a boolean if a field has been set.

### GetDelegated

`func (o *ReviewSparse) GetDelegated() bool`

GetDelegated returns the Delegated field if non-nil, zero value otherwise.

### GetDelegatedOk

`func (o *ReviewSparse) GetDelegatedOk() (*bool, bool)`

GetDelegatedOk returns a tuple with the Delegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegated

`func (o *ReviewSparse) SetDelegated(v bool)`

SetDelegated sets Delegated field to given value.

### HasDelegated

`func (o *ReviewSparse) HasDelegated() bool`

HasDelegated returns a boolean if a field has been set.

### GetAppServiceAccount

`func (o *ReviewSparse) GetAppServiceAccount() ReviewerServiceAccount`

GetAppServiceAccount returns the AppServiceAccount field if non-nil, zero value otherwise.

### GetAppServiceAccountOk

`func (o *ReviewSparse) GetAppServiceAccountOk() (*ReviewerServiceAccount, bool)`

GetAppServiceAccountOk returns a tuple with the AppServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServiceAccount

`func (o *ReviewSparse) SetAppServiceAccount(v ReviewerServiceAccount)`

SetAppServiceAccount sets AppServiceAccount field to given value.

### HasAppServiceAccount

`func (o *ReviewSparse) HasAppServiceAccount() bool`

HasAppServiceAccount returns a boolean if a field has been set.

### GetOktaServiceAccount

`func (o *ReviewSparse) GetOktaServiceAccount() ReviewerServiceAccount`

GetOktaServiceAccount returns the OktaServiceAccount field if non-nil, zero value otherwise.

### GetOktaServiceAccountOk

`func (o *ReviewSparse) GetOktaServiceAccountOk() (*ReviewerServiceAccount, bool)`

GetOktaServiceAccountOk returns a tuple with the OktaServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaServiceAccount

`func (o *ReviewSparse) SetOktaServiceAccount(v ReviewerServiceAccount)`

SetOktaServiceAccount sets OktaServiceAccount field to given value.

### HasOktaServiceAccount

`func (o *ReviewSparse) HasOktaServiceAccount() bool`

HasOktaServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


