# ReviewerLevelInfoFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewerLevel** | [**ReviewerLevelType**](ReviewerLevelType.md) |  | 
**Decision** | [**Decision**](Decision.md) |  | 
**ReviewerProfile** | Pointer to [**PrincipalProfileEnriched**](PrincipalProfileEnriched.md) |  | [optional] 
**ReviewerType** | [**ReviewersReviewerType**](ReviewersReviewerType.md) |  | 
**ReviewerGroupProfile** | Pointer to [**ReviewerGroupProfile**](ReviewerGroupProfile.md) |  | [optional] 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | Pointer to [**map[string]Link**](Link.md) | Links to related resources | [optional] 

## Methods

### NewReviewerLevelInfoFull

`func NewReviewerLevelInfoFull(reviewerLevel ReviewerLevelType, decision Decision, reviewerType ReviewersReviewerType, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *ReviewerLevelInfoFull`

NewReviewerLevelInfoFull instantiates a new ReviewerLevelInfoFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerLevelInfoFullWithDefaults

`func NewReviewerLevelInfoFullWithDefaults() *ReviewerLevelInfoFull`

NewReviewerLevelInfoFullWithDefaults instantiates a new ReviewerLevelInfoFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewerLevel

`func (o *ReviewerLevelInfoFull) GetReviewerLevel() ReviewerLevelType`

GetReviewerLevel returns the ReviewerLevel field if non-nil, zero value otherwise.

### GetReviewerLevelOk

`func (o *ReviewerLevelInfoFull) GetReviewerLevelOk() (*ReviewerLevelType, bool)`

GetReviewerLevelOk returns a tuple with the ReviewerLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerLevel

`func (o *ReviewerLevelInfoFull) SetReviewerLevel(v ReviewerLevelType)`

SetReviewerLevel sets ReviewerLevel field to given value.


### GetDecision

`func (o *ReviewerLevelInfoFull) GetDecision() Decision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ReviewerLevelInfoFull) GetDecisionOk() (*Decision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ReviewerLevelInfoFull) SetDecision(v Decision)`

SetDecision sets Decision field to given value.


### GetReviewerProfile

`func (o *ReviewerLevelInfoFull) GetReviewerProfile() PrincipalProfileEnriched`

GetReviewerProfile returns the ReviewerProfile field if non-nil, zero value otherwise.

### GetReviewerProfileOk

`func (o *ReviewerLevelInfoFull) GetReviewerProfileOk() (*PrincipalProfileEnriched, bool)`

GetReviewerProfileOk returns a tuple with the ReviewerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerProfile

`func (o *ReviewerLevelInfoFull) SetReviewerProfile(v PrincipalProfileEnriched)`

SetReviewerProfile sets ReviewerProfile field to given value.

### HasReviewerProfile

`func (o *ReviewerLevelInfoFull) HasReviewerProfile() bool`

HasReviewerProfile returns a boolean if a field has been set.

### GetReviewerType

`func (o *ReviewerLevelInfoFull) GetReviewerType() ReviewersReviewerType`

GetReviewerType returns the ReviewerType field if non-nil, zero value otherwise.

### GetReviewerTypeOk

`func (o *ReviewerLevelInfoFull) GetReviewerTypeOk() (*ReviewersReviewerType, bool)`

GetReviewerTypeOk returns a tuple with the ReviewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerType

`func (o *ReviewerLevelInfoFull) SetReviewerType(v ReviewersReviewerType)`

SetReviewerType sets ReviewerType field to given value.


### GetReviewerGroupProfile

`func (o *ReviewerLevelInfoFull) GetReviewerGroupProfile() ReviewerGroupProfile`

GetReviewerGroupProfile returns the ReviewerGroupProfile field if non-nil, zero value otherwise.

### GetReviewerGroupProfileOk

`func (o *ReviewerLevelInfoFull) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool)`

GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupProfile

`func (o *ReviewerLevelInfoFull) SetReviewerGroupProfile(v ReviewerGroupProfile)`

SetReviewerGroupProfile sets ReviewerGroupProfile field to given value.

### HasReviewerGroupProfile

`func (o *ReviewerLevelInfoFull) HasReviewerGroupProfile() bool`

HasReviewerGroupProfile returns a boolean if a field has been set.

### GetId

`func (o *ReviewerLevelInfoFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewerLevelInfoFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewerLevelInfoFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ReviewerLevelInfoFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ReviewerLevelInfoFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ReviewerLevelInfoFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *ReviewerLevelInfoFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReviewerLevelInfoFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReviewerLevelInfoFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *ReviewerLevelInfoFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ReviewerLevelInfoFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ReviewerLevelInfoFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *ReviewerLevelInfoFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ReviewerLevelInfoFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ReviewerLevelInfoFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *ReviewerLevelInfoFull) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReviewerLevelInfoFull) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReviewerLevelInfoFull) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReviewerLevelInfoFull) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


