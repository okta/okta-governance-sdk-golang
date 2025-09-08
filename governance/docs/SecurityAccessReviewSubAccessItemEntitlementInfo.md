# SecurityAccessReviewSubAccessItemEntitlementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SecurityAccessReviewSubAccessItemEntitlementType**](SecurityAccessReviewSubAccessItemEntitlementType.md) |  | 
**Description** | Pointer to **string** | A brief description of the entitlement value or bundle | [optional] 
**EntitlementDescription** | Pointer to **string** | A brief description of the entitlement associated with the value | [optional] 
**AssignedDate** | Pointer to **time.Time** | The date the entitlement or bundle was assigned to the user | [optional] 
**AssignmentType** | Pointer to [**AssignmentType**](AssignmentType.md) |  | [optional] 
**CollectionsAssigning** | Pointer to [**[]CollectionInfoSparse**](CollectionInfoSparse.md) | Collections assigning this resource | [optional] 
**Entitlements** | Pointer to [**[]EntitlementPropertyFull**](EntitlementPropertyFull.md) | If a bundle, these are the entitlements included in the bundle | [optional] 
**LastAccessCertificationReviewedDate** | Pointer to **time.Time** | The last time the user/entitlement pair was reviewed in an access certification campaign | [optional] 
**LastSecurityAccessReviewDate** | Pointer to **time.Time** | The last time an action was taken on this entitlement for the user in a security access review | [optional] 
**GovernanceLabels** | Pointer to [**[]GovernanceLabel**](GovernanceLabel.md) | All governance labels applied to the entitlement value or bundle | [optional] 

## Methods

### NewSecurityAccessReviewSubAccessItemEntitlementInfo

`func NewSecurityAccessReviewSubAccessItemEntitlementInfo(type_ SecurityAccessReviewSubAccessItemEntitlementType, ) *SecurityAccessReviewSubAccessItemEntitlementInfo`

NewSecurityAccessReviewSubAccessItemEntitlementInfo instantiates a new SecurityAccessReviewSubAccessItemEntitlementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewSubAccessItemEntitlementInfoWithDefaults

`func NewSecurityAccessReviewSubAccessItemEntitlementInfoWithDefaults() *SecurityAccessReviewSubAccessItemEntitlementInfo`

NewSecurityAccessReviewSubAccessItemEntitlementInfoWithDefaults instantiates a new SecurityAccessReviewSubAccessItemEntitlementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetType() SecurityAccessReviewSubAccessItemEntitlementType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetTypeOk() (*SecurityAccessReviewSubAccessItemEntitlementType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetType(v SecurityAccessReviewSubAccessItemEntitlementType)`

SetType sets Type field to given value.


### GetDescription

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntitlementDescription

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlementDescription() string`

GetEntitlementDescription returns the EntitlementDescription field if non-nil, zero value otherwise.

### GetEntitlementDescriptionOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlementDescriptionOk() (*string, bool)`

GetEntitlementDescriptionOk returns a tuple with the EntitlementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementDescription

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetEntitlementDescription(v string)`

SetEntitlementDescription sets EntitlementDescription field to given value.

### HasEntitlementDescription

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasEntitlementDescription() bool`

HasEntitlementDescription returns a boolean if a field has been set.

### GetAssignedDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignedDate() time.Time`

GetAssignedDate returns the AssignedDate field if non-nil, zero value otherwise.

### GetAssignedDateOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignedDateOk() (*time.Time, bool)`

GetAssignedDateOk returns a tuple with the AssignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetAssignedDate(v time.Time)`

SetAssignedDate sets AssignedDate field to given value.

### HasAssignedDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasAssignedDate() bool`

HasAssignedDate returns a boolean if a field has been set.

### GetAssignmentType

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignmentType() AssignmentType`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignmentTypeOk() (*AssignmentType, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetAssignmentType(v AssignmentType)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetCollectionsAssigning

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetCollectionsAssigning() []CollectionInfoSparse`

GetCollectionsAssigning returns the CollectionsAssigning field if non-nil, zero value otherwise.

### GetCollectionsAssigningOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetCollectionsAssigningOk() (*[]CollectionInfoSparse, bool)`

GetCollectionsAssigningOk returns a tuple with the CollectionsAssigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionsAssigning

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetCollectionsAssigning(v []CollectionInfoSparse)`

SetCollectionsAssigning sets CollectionsAssigning field to given value.

### HasCollectionsAssigning

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasCollectionsAssigning() bool`

HasCollectionsAssigning returns a boolean if a field has been set.

### GetEntitlements

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlements() []EntitlementPropertyFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlementsOk() (*[]EntitlementPropertyFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetEntitlements(v []EntitlementPropertyFull)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastAccessCertificationReviewedDate() time.Time`

GetLastAccessCertificationReviewedDate returns the LastAccessCertificationReviewedDate field if non-nil, zero value otherwise.

### GetLastAccessCertificationReviewedDateOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastAccessCertificationReviewedDateOk() (*time.Time, bool)`

GetLastAccessCertificationReviewedDateOk returns a tuple with the LastAccessCertificationReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetLastAccessCertificationReviewedDate(v time.Time)`

SetLastAccessCertificationReviewedDate sets LastAccessCertificationReviewedDate field to given value.

### HasLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasLastAccessCertificationReviewedDate() bool`

HasLastAccessCertificationReviewedDate returns a boolean if a field has been set.

### GetLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastSecurityAccessReviewDate() time.Time`

GetLastSecurityAccessReviewDate returns the LastSecurityAccessReviewDate field if non-nil, zero value otherwise.

### GetLastSecurityAccessReviewDateOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastSecurityAccessReviewDateOk() (*time.Time, bool)`

GetLastSecurityAccessReviewDateOk returns a tuple with the LastSecurityAccessReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetLastSecurityAccessReviewDate(v time.Time)`

SetLastSecurityAccessReviewDate sets LastSecurityAccessReviewDate field to given value.

### HasLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasLastSecurityAccessReviewDate() bool`

HasLastSecurityAccessReviewDate returns a boolean if a field has been set.

### GetGovernanceLabels

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetGovernanceLabels() []GovernanceLabel`

GetGovernanceLabels returns the GovernanceLabels field if non-nil, zero value otherwise.

### GetGovernanceLabelsOk

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetGovernanceLabelsOk() (*[]GovernanceLabel, bool)`

GetGovernanceLabelsOk returns a tuple with the GovernanceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceLabels

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetGovernanceLabels(v []GovernanceLabel)`

SetGovernanceLabels sets GovernanceLabels field to given value.

### HasGovernanceLabels

`func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasGovernanceLabels() bool`

HasGovernanceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


