# ReviewFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | [**ReviewLinks**](ReviewLinks.md) |  | 
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
**RiskRuleConflicts** | Pointer to [**[]RiskRuleConflicts**](RiskRuleConflicts.md) | List of risk rule conflicts caused by this entitlement value. Only applies to review item that has entitlement value. | [optional] 
**Note** | Pointer to [**Note**](Note.md) |  | [optional] 
**AllReviewerLevels** | Pointer to [**[]ReviewerLevelInfoFull**](ReviewerLevelInfoFull.md) | Applicable only for multi level campaign. Provides details about the reviewer and decisions (if any) made at each reviewer level is captured here. | [optional] 

## Methods

### NewReviewFull

`func NewReviewFull(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links ReviewLinks, campaignId string, resourceId string, decision Decision, remediationStatus RemediationStatus, principalProfile PrincipalProfile, reviewerType ReviewersReviewerType, ) *ReviewFull`

NewReviewFull instantiates a new ReviewFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewFullWithDefaults

`func NewReviewFullWithDefaults() *ReviewFull`

NewReviewFullWithDefaults instantiates a new ReviewFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ReviewFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ReviewFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ReviewFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *ReviewFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReviewFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReviewFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *ReviewFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ReviewFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ReviewFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *ReviewFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ReviewFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ReviewFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *ReviewFull) GetLinks() ReviewLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReviewFull) GetLinksOk() (*ReviewLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReviewFull) SetLinks(v ReviewLinks)`

SetLinks sets Links field to given value.


### GetCampaignId

`func (o *ReviewFull) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ReviewFull) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ReviewFull) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetResourceId

`func (o *ReviewFull) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ReviewFull) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ReviewFull) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetEntitlementValue

`func (o *ReviewFull) GetEntitlementValue() ReviewerEntitlementValue`

GetEntitlementValue returns the EntitlementValue field if non-nil, zero value otherwise.

### GetEntitlementValueOk

`func (o *ReviewFull) GetEntitlementValueOk() (*ReviewerEntitlementValue, bool)`

GetEntitlementValueOk returns a tuple with the EntitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValue

`func (o *ReviewFull) SetEntitlementValue(v ReviewerEntitlementValue)`

SetEntitlementValue sets EntitlementValue field to given value.

### HasEntitlementValue

`func (o *ReviewFull) HasEntitlementValue() bool`

HasEntitlementValue returns a boolean if a field has been set.

### GetEntitlementBundle

`func (o *ReviewFull) GetEntitlementBundle() ReviewerEntitlementBundle`

GetEntitlementBundle returns the EntitlementBundle field if non-nil, zero value otherwise.

### GetEntitlementBundleOk

`func (o *ReviewFull) GetEntitlementBundleOk() (*ReviewerEntitlementBundle, bool)`

GetEntitlementBundleOk returns a tuple with the EntitlementBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundle

`func (o *ReviewFull) SetEntitlementBundle(v ReviewerEntitlementBundle)`

SetEntitlementBundle sets EntitlementBundle field to given value.

### HasEntitlementBundle

`func (o *ReviewFull) HasEntitlementBundle() bool`

HasEntitlementBundle returns a boolean if a field has been set.

### GetDecision

`func (o *ReviewFull) GetDecision() Decision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ReviewFull) GetDecisionOk() (*Decision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ReviewFull) SetDecision(v Decision)`

SetDecision sets Decision field to given value.


### GetDecided

`func (o *ReviewFull) GetDecided() time.Time`

GetDecided returns the Decided field if non-nil, zero value otherwise.

### GetDecidedOk

`func (o *ReviewFull) GetDecidedOk() (*time.Time, bool)`

GetDecidedOk returns a tuple with the Decided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecided

`func (o *ReviewFull) SetDecided(v time.Time)`

SetDecided sets Decided field to given value.

### HasDecided

`func (o *ReviewFull) HasDecided() bool`

HasDecided returns a boolean if a field has been set.

### GetRemediationStatus

`func (o *ReviewFull) GetRemediationStatus() RemediationStatus`

GetRemediationStatus returns the RemediationStatus field if non-nil, zero value otherwise.

### GetRemediationStatusOk

`func (o *ReviewFull) GetRemediationStatusOk() (*RemediationStatus, bool)`

GetRemediationStatusOk returns a tuple with the RemediationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationStatus

`func (o *ReviewFull) SetRemediationStatus(v RemediationStatus)`

SetRemediationStatus sets RemediationStatus field to given value.


### GetPrincipalProfile

`func (o *ReviewFull) GetPrincipalProfile() PrincipalProfile`

GetPrincipalProfile returns the PrincipalProfile field if non-nil, zero value otherwise.

### GetPrincipalProfileOk

`func (o *ReviewFull) GetPrincipalProfileOk() (*PrincipalProfile, bool)`

GetPrincipalProfileOk returns a tuple with the PrincipalProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalProfile

`func (o *ReviewFull) SetPrincipalProfile(v PrincipalProfile)`

SetPrincipalProfile sets PrincipalProfile field to given value.


### GetReviewerProfile

`func (o *ReviewFull) GetReviewerProfile() PrincipalProfile`

GetReviewerProfile returns the ReviewerProfile field if non-nil, zero value otherwise.

### GetReviewerProfileOk

`func (o *ReviewFull) GetReviewerProfileOk() (*PrincipalProfile, bool)`

GetReviewerProfileOk returns a tuple with the ReviewerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerProfile

`func (o *ReviewFull) SetReviewerProfile(v PrincipalProfile)`

SetReviewerProfile sets ReviewerProfile field to given value.

### HasReviewerProfile

`func (o *ReviewFull) HasReviewerProfile() bool`

HasReviewerProfile returns a boolean if a field has been set.

### GetReviewerType

`func (o *ReviewFull) GetReviewerType() ReviewersReviewerType`

GetReviewerType returns the ReviewerType field if non-nil, zero value otherwise.

### GetReviewerTypeOk

`func (o *ReviewFull) GetReviewerTypeOk() (*ReviewersReviewerType, bool)`

GetReviewerTypeOk returns a tuple with the ReviewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerType

`func (o *ReviewFull) SetReviewerType(v ReviewersReviewerType)`

SetReviewerType sets ReviewerType field to given value.


### GetReviewerGroupProfile

`func (o *ReviewFull) GetReviewerGroupProfile() ReviewerGroupProfile`

GetReviewerGroupProfile returns the ReviewerGroupProfile field if non-nil, zero value otherwise.

### GetReviewerGroupProfileOk

`func (o *ReviewFull) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool)`

GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupProfile

`func (o *ReviewFull) SetReviewerGroupProfile(v ReviewerGroupProfile)`

SetReviewerGroupProfile sets ReviewerGroupProfile field to given value.

### HasReviewerGroupProfile

`func (o *ReviewFull) HasReviewerGroupProfile() bool`

HasReviewerGroupProfile returns a boolean if a field has been set.

### GetCurrentReviewerLevel

`func (o *ReviewFull) GetCurrentReviewerLevel() ReviewerLevelType`

GetCurrentReviewerLevel returns the CurrentReviewerLevel field if non-nil, zero value otherwise.

### GetCurrentReviewerLevelOk

`func (o *ReviewFull) GetCurrentReviewerLevelOk() (*ReviewerLevelType, bool)`

GetCurrentReviewerLevelOk returns a tuple with the CurrentReviewerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReviewerLevel

`func (o *ReviewFull) SetCurrentReviewerLevel(v ReviewerLevelType)`

SetCurrentReviewerLevel sets CurrentReviewerLevel field to given value.

### HasCurrentReviewerLevel

`func (o *ReviewFull) HasCurrentReviewerLevel() bool`

HasCurrentReviewerLevel returns a boolean if a field has been set.

### GetRiskRuleConflicts

`func (o *ReviewFull) GetRiskRuleConflicts() []RiskRuleConflicts`

GetRiskRuleConflicts returns the RiskRuleConflicts field if non-nil, zero value otherwise.

### GetRiskRuleConflictsOk

`func (o *ReviewFull) GetRiskRuleConflictsOk() (*[]RiskRuleConflicts, bool)`

GetRiskRuleConflictsOk returns a tuple with the RiskRuleConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRuleConflicts

`func (o *ReviewFull) SetRiskRuleConflicts(v []RiskRuleConflicts)`

SetRiskRuleConflicts sets RiskRuleConflicts field to given value.

### HasRiskRuleConflicts

`func (o *ReviewFull) HasRiskRuleConflicts() bool`

HasRiskRuleConflicts returns a boolean if a field has been set.

### SetRiskRuleConflictsNil

`func (o *ReviewFull) SetRiskRuleConflictsNil(b bool)`

 SetRiskRuleConflictsNil sets the value for RiskRuleConflicts to be an explicit nil

### UnsetRiskRuleConflicts
`func (o *ReviewFull) UnsetRiskRuleConflicts()`

UnsetRiskRuleConflicts ensures that no value is present for RiskRuleConflicts, not even an explicit nil
### GetNote

`func (o *ReviewFull) GetNote() Note`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReviewFull) GetNoteOk() (*Note, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReviewFull) SetNote(v Note)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReviewFull) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAllReviewerLevels

`func (o *ReviewFull) GetAllReviewerLevels() []ReviewerLevelInfoFull`

GetAllReviewerLevels returns the AllReviewerLevels field if non-nil, zero value otherwise.

### GetAllReviewerLevelsOk

`func (o *ReviewFull) GetAllReviewerLevelsOk() (*[]ReviewerLevelInfoFull, bool)`

GetAllReviewerLevelsOk returns a tuple with the AllReviewerLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviewerLevels

`func (o *ReviewFull) SetAllReviewerLevels(v []ReviewerLevelInfoFull)`

SetAllReviewerLevels sets AllReviewerLevels field to given value.

### HasAllReviewerLevels

`func (o *ReviewFull) HasAllReviewerLevels() bool`

HasAllReviewerLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


