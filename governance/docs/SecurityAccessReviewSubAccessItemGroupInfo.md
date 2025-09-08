# SecurityAccessReviewSubAccessItemGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the group | [optional] 
**AssignedDate** | Pointer to **time.Time** | The date the group was assigned to the user | [optional] 
**GroupOwner** | Pointer to **string** | The group owner name | [optional] 
**AssignmentType** | Pointer to [**AssignmentType**](AssignmentType.md) |  | [optional] 
**LastAccessCertificationReviewedDate** | Pointer to **time.Time** | The last time the user/group pair was reviewed in an access certification campaign | [optional] 
**LastSecurityAccessReviewDate** | Pointer to **time.Time** | The last time an action was taken on this group for the user in a security access review | [optional] 
**GovernanceLabels** | Pointer to [**[]GovernanceLabel**](GovernanceLabel.md) | All governance labels applied to the group | [optional] 

## Methods

### NewSecurityAccessReviewSubAccessItemGroupInfo

`func NewSecurityAccessReviewSubAccessItemGroupInfo() *SecurityAccessReviewSubAccessItemGroupInfo`

NewSecurityAccessReviewSubAccessItemGroupInfo instantiates a new SecurityAccessReviewSubAccessItemGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewSubAccessItemGroupInfoWithDefaults

`func NewSecurityAccessReviewSubAccessItemGroupInfoWithDefaults() *SecurityAccessReviewSubAccessItemGroupInfo`

NewSecurityAccessReviewSubAccessItemGroupInfoWithDefaults instantiates a new SecurityAccessReviewSubAccessItemGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAssignedDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignedDate() time.Time`

GetAssignedDate returns the AssignedDate field if non-nil, zero value otherwise.

### GetAssignedDateOk

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignedDateOk() (*time.Time, bool)`

GetAssignedDateOk returns a tuple with the AssignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetAssignedDate(v time.Time)`

SetAssignedDate sets AssignedDate field to given value.

### HasAssignedDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasAssignedDate() bool`

HasAssignedDate returns a boolean if a field has been set.

### GetGroupOwner

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGroupOwner() string`

GetGroupOwner returns the GroupOwner field if non-nil, zero value otherwise.

### GetGroupOwnerOk

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGroupOwnerOk() (*string, bool)`

GetGroupOwnerOk returns a tuple with the GroupOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOwner

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetGroupOwner(v string)`

SetGroupOwner sets GroupOwner field to given value.

### HasGroupOwner

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasGroupOwner() bool`

HasGroupOwner returns a boolean if a field has been set.

### GetAssignmentType

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignmentType() AssignmentType`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetAssignmentTypeOk() (*AssignmentType, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetAssignmentType(v AssignmentType)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastAccessCertificationReviewedDate() time.Time`

GetLastAccessCertificationReviewedDate returns the LastAccessCertificationReviewedDate field if non-nil, zero value otherwise.

### GetLastAccessCertificationReviewedDateOk

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastAccessCertificationReviewedDateOk() (*time.Time, bool)`

GetLastAccessCertificationReviewedDateOk returns a tuple with the LastAccessCertificationReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetLastAccessCertificationReviewedDate(v time.Time)`

SetLastAccessCertificationReviewedDate sets LastAccessCertificationReviewedDate field to given value.

### HasLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasLastAccessCertificationReviewedDate() bool`

HasLastAccessCertificationReviewedDate returns a boolean if a field has been set.

### GetLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastSecurityAccessReviewDate() time.Time`

GetLastSecurityAccessReviewDate returns the LastSecurityAccessReviewDate field if non-nil, zero value otherwise.

### GetLastSecurityAccessReviewDateOk

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetLastSecurityAccessReviewDateOk() (*time.Time, bool)`

GetLastSecurityAccessReviewDateOk returns a tuple with the LastSecurityAccessReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetLastSecurityAccessReviewDate(v time.Time)`

SetLastSecurityAccessReviewDate sets LastSecurityAccessReviewDate field to given value.

### HasLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasLastSecurityAccessReviewDate() bool`

HasLastSecurityAccessReviewDate returns a boolean if a field has been set.

### GetGovernanceLabels

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGovernanceLabels() []GovernanceLabel`

GetGovernanceLabels returns the GovernanceLabels field if non-nil, zero value otherwise.

### GetGovernanceLabelsOk

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) GetGovernanceLabelsOk() (*[]GovernanceLabel, bool)`

GetGovernanceLabelsOk returns a tuple with the GovernanceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceLabels

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) SetGovernanceLabels(v []GovernanceLabel)`

SetGovernanceLabels sets GovernanceLabels field to given value.

### HasGovernanceLabels

`func (o *SecurityAccessReviewSubAccessItemGroupInfo) HasGovernanceLabels() bool`

HasGovernanceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


