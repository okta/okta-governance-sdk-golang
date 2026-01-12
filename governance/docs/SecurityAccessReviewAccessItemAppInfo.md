# SecurityAccessReviewAccessItemAppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Okta app ID | 
**Name** | **string** | The app name | 
**LogoUrl** | **string** | The app logo URL | 
**Label** | Pointer to **string** | The app label | [optional] 
**LastAccessDate** | Pointer to **time.Time** | The last time the app was accessed by the user | [optional] 
**AssignedDate** | Pointer to **time.Time** | The date the app was assigned to the user | [optional] 
**AssignmentType** | Pointer to [**AssignmentType**](AssignmentType.md) |  | [optional] 
**ApplicationUsage** | Pointer to **int64** | The number of times the user has used the app in the last 90 days | [optional] 
**LastAccessCertificationReviewedDate** | Pointer to **time.Time** | The last time the user/app pair was reviewed in an access certification campaign | [optional] 
**LastSecurityAccessReviewDate** | Pointer to **time.Time** | The last time an action was taken on this app for the user in a security access review | [optional] 
**LastDeviceLogin** | Pointer to **string** | Information about the last device used to log into the app | [optional] 
**ActiveEntitlements** | Pointer to [**[]EntitlementPropertyFull**](EntitlementPropertyFull.md) | Any active entitlements for this app | [optional] 
**GovernanceLabels** | Pointer to [**[]TargetGovernanceLabel**](TargetGovernanceLabel.md) | All governance labels applied to the app | [optional] 

## Methods

### NewSecurityAccessReviewAccessItemAppInfo

`func NewSecurityAccessReviewAccessItemAppInfo(id string, name string, logoUrl string, ) *SecurityAccessReviewAccessItemAppInfo`

NewSecurityAccessReviewAccessItemAppInfo instantiates a new SecurityAccessReviewAccessItemAppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewAccessItemAppInfoWithDefaults

`func NewSecurityAccessReviewAccessItemAppInfoWithDefaults() *SecurityAccessReviewAccessItemAppInfo`

NewSecurityAccessReviewAccessItemAppInfoWithDefaults instantiates a new SecurityAccessReviewAccessItemAppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReviewAccessItemAppInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReviewAccessItemAppInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityAccessReviewAccessItemAppInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityAccessReviewAccessItemAppInfo) SetName(v string)`

SetName sets Name field to given value.


### GetLogoUrl

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *SecurityAccessReviewAccessItemAppInfo) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### GetLabel

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SecurityAccessReviewAccessItemAppInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SecurityAccessReviewAccessItemAppInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastAccessDate

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessDate() time.Time`

GetLastAccessDate returns the LastAccessDate field if non-nil, zero value otherwise.

### GetLastAccessDateOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessDateOk() (*time.Time, bool)`

GetLastAccessDateOk returns a tuple with the LastAccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessDate

`func (o *SecurityAccessReviewAccessItemAppInfo) SetLastAccessDate(v time.Time)`

SetLastAccessDate sets LastAccessDate field to given value.

### HasLastAccessDate

`func (o *SecurityAccessReviewAccessItemAppInfo) HasLastAccessDate() bool`

HasLastAccessDate returns a boolean if a field has been set.

### GetAssignedDate

`func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignedDate() time.Time`

GetAssignedDate returns the AssignedDate field if non-nil, zero value otherwise.

### GetAssignedDateOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignedDateOk() (*time.Time, bool)`

GetAssignedDateOk returns a tuple with the AssignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDate

`func (o *SecurityAccessReviewAccessItemAppInfo) SetAssignedDate(v time.Time)`

SetAssignedDate sets AssignedDate field to given value.

### HasAssignedDate

`func (o *SecurityAccessReviewAccessItemAppInfo) HasAssignedDate() bool`

HasAssignedDate returns a boolean if a field has been set.

### GetAssignmentType

`func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignmentType() AssignmentType`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignmentTypeOk() (*AssignmentType, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *SecurityAccessReviewAccessItemAppInfo) SetAssignmentType(v AssignmentType)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *SecurityAccessReviewAccessItemAppInfo) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetApplicationUsage

`func (o *SecurityAccessReviewAccessItemAppInfo) GetApplicationUsage() int64`

GetApplicationUsage returns the ApplicationUsage field if non-nil, zero value otherwise.

### GetApplicationUsageOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetApplicationUsageOk() (*int64, bool)`

GetApplicationUsageOk returns a tuple with the ApplicationUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUsage

`func (o *SecurityAccessReviewAccessItemAppInfo) SetApplicationUsage(v int64)`

SetApplicationUsage sets ApplicationUsage field to given value.

### HasApplicationUsage

`func (o *SecurityAccessReviewAccessItemAppInfo) HasApplicationUsage() bool`

HasApplicationUsage returns a boolean if a field has been set.

### GetLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessCertificationReviewedDate() time.Time`

GetLastAccessCertificationReviewedDate returns the LastAccessCertificationReviewedDate field if non-nil, zero value otherwise.

### GetLastAccessCertificationReviewedDateOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessCertificationReviewedDateOk() (*time.Time, bool)`

GetLastAccessCertificationReviewedDateOk returns a tuple with the LastAccessCertificationReviewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewAccessItemAppInfo) SetLastAccessCertificationReviewedDate(v time.Time)`

SetLastAccessCertificationReviewedDate sets LastAccessCertificationReviewedDate field to given value.

### HasLastAccessCertificationReviewedDate

`func (o *SecurityAccessReviewAccessItemAppInfo) HasLastAccessCertificationReviewedDate() bool`

HasLastAccessCertificationReviewedDate returns a boolean if a field has been set.

### GetLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastSecurityAccessReviewDate() time.Time`

GetLastSecurityAccessReviewDate returns the LastSecurityAccessReviewDate field if non-nil, zero value otherwise.

### GetLastSecurityAccessReviewDateOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastSecurityAccessReviewDateOk() (*time.Time, bool)`

GetLastSecurityAccessReviewDateOk returns a tuple with the LastSecurityAccessReviewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewAccessItemAppInfo) SetLastSecurityAccessReviewDate(v time.Time)`

SetLastSecurityAccessReviewDate sets LastSecurityAccessReviewDate field to given value.

### HasLastSecurityAccessReviewDate

`func (o *SecurityAccessReviewAccessItemAppInfo) HasLastSecurityAccessReviewDate() bool`

HasLastSecurityAccessReviewDate returns a boolean if a field has been set.

### GetLastDeviceLogin

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastDeviceLogin() string`

GetLastDeviceLogin returns the LastDeviceLogin field if non-nil, zero value otherwise.

### GetLastDeviceLoginOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetLastDeviceLoginOk() (*string, bool)`

GetLastDeviceLoginOk returns a tuple with the LastDeviceLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeviceLogin

`func (o *SecurityAccessReviewAccessItemAppInfo) SetLastDeviceLogin(v string)`

SetLastDeviceLogin sets LastDeviceLogin field to given value.

### HasLastDeviceLogin

`func (o *SecurityAccessReviewAccessItemAppInfo) HasLastDeviceLogin() bool`

HasLastDeviceLogin returns a boolean if a field has been set.

### GetActiveEntitlements

`func (o *SecurityAccessReviewAccessItemAppInfo) GetActiveEntitlements() []EntitlementPropertyFull`

GetActiveEntitlements returns the ActiveEntitlements field if non-nil, zero value otherwise.

### GetActiveEntitlementsOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetActiveEntitlementsOk() (*[]EntitlementPropertyFull, bool)`

GetActiveEntitlementsOk returns a tuple with the ActiveEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEntitlements

`func (o *SecurityAccessReviewAccessItemAppInfo) SetActiveEntitlements(v []EntitlementPropertyFull)`

SetActiveEntitlements sets ActiveEntitlements field to given value.

### HasActiveEntitlements

`func (o *SecurityAccessReviewAccessItemAppInfo) HasActiveEntitlements() bool`

HasActiveEntitlements returns a boolean if a field has been set.

### GetGovernanceLabels

`func (o *SecurityAccessReviewAccessItemAppInfo) GetGovernanceLabels() []TargetGovernanceLabel`

GetGovernanceLabels returns the GovernanceLabels field if non-nil, zero value otherwise.

### GetGovernanceLabelsOk

`func (o *SecurityAccessReviewAccessItemAppInfo) GetGovernanceLabelsOk() (*[]TargetGovernanceLabel, bool)`

GetGovernanceLabelsOk returns a tuple with the GovernanceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceLabels

`func (o *SecurityAccessReviewAccessItemAppInfo) SetGovernanceLabels(v []TargetGovernanceLabel)`

SetGovernanceLabels sets GovernanceLabels field to given value.

### HasGovernanceLabels

`func (o *SecurityAccessReviewAccessItemAppInfo) HasGovernanceLabels() bool`

HasGovernanceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


