# ReviewerGroupProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**GroupId** | **string** |  | 
**GroupType** | [**GroupType**](GroupType.md) |  | 

## Methods

### NewReviewerGroupProfile

`func NewReviewerGroupProfile(name string, groupId string, groupType GroupType, ) *ReviewerGroupProfile`

NewReviewerGroupProfile instantiates a new ReviewerGroupProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerGroupProfileWithDefaults

`func NewReviewerGroupProfileWithDefaults() *ReviewerGroupProfile`

NewReviewerGroupProfileWithDefaults instantiates a new ReviewerGroupProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReviewerGroupProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewerGroupProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewerGroupProfile) SetName(v string)`

SetName sets Name field to given value.


### GetGroupId

`func (o *ReviewerGroupProfile) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ReviewerGroupProfile) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ReviewerGroupProfile) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetGroupType

`func (o *ReviewerGroupProfile) GetGroupType() GroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *ReviewerGroupProfile) GetGroupTypeOk() (*GroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *ReviewerGroupProfile) SetGroupType(v GroupType)`

SetGroupType sets GroupType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


