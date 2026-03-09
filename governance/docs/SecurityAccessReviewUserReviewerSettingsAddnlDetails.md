# SecurityAccessReviewUserReviewerSettingsAddnlDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedUserIds** | Pointer to **[]string** | The list of reviewer user IDs for the security access review | [optional] 
**IncludedUserProfiles** | Pointer to [**[]PrincipalProfileEnriched**](PrincipalProfileEnriched.md) | The list of user reviewers for the security access review | [optional] 

## Methods

### NewSecurityAccessReviewUserReviewerSettingsAddnlDetails

`func NewSecurityAccessReviewUserReviewerSettingsAddnlDetails() *SecurityAccessReviewUserReviewerSettingsAddnlDetails`

NewSecurityAccessReviewUserReviewerSettingsAddnlDetails instantiates a new SecurityAccessReviewUserReviewerSettingsAddnlDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewUserReviewerSettingsAddnlDetailsWithDefaults

`func NewSecurityAccessReviewUserReviewerSettingsAddnlDetailsWithDefaults() *SecurityAccessReviewUserReviewerSettingsAddnlDetails`

NewSecurityAccessReviewUserReviewerSettingsAddnlDetailsWithDefaults instantiates a new SecurityAccessReviewUserReviewerSettingsAddnlDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedUserIds

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserIds() []string`

GetIncludedUserIds returns the IncludedUserIds field if non-nil, zero value otherwise.

### GetIncludedUserIdsOk

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserIdsOk() (*[]string, bool)`

GetIncludedUserIdsOk returns a tuple with the IncludedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserIds

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) SetIncludedUserIds(v []string)`

SetIncludedUserIds sets IncludedUserIds field to given value.

### HasIncludedUserIds

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) HasIncludedUserIds() bool`

HasIncludedUserIds returns a boolean if a field has been set.

### GetIncludedUserProfiles

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserProfiles() []PrincipalProfileEnriched`

GetIncludedUserProfiles returns the IncludedUserProfiles field if non-nil, zero value otherwise.

### GetIncludedUserProfilesOk

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) GetIncludedUserProfilesOk() (*[]PrincipalProfileEnriched, bool)`

GetIncludedUserProfilesOk returns a tuple with the IncludedUserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserProfiles

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) SetIncludedUserProfiles(v []PrincipalProfileEnriched)`

SetIncludedUserProfiles sets IncludedUserProfiles field to given value.

### HasIncludedUserProfiles

`func (o *SecurityAccessReviewUserReviewerSettingsAddnlDetails) HasIncludedUserProfiles() bool`

HasIncludedUserProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


