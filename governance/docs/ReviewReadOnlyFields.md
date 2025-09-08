# ReviewReadOnlyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to [**Note**](Note.md) |  | [optional] 
**ReviewerGroupProfile** | Pointer to [**ReviewerGroupProfile**](ReviewerGroupProfile.md) |  | [optional] 
**AllReviewerLevels** | Pointer to [**[]ReviewerLevelInfoFull**](ReviewerLevelInfoFull.md) | Applicable only for multi level campaign. Provides details about the reviewer and decisions (if any) made at each reviewer level is captured here. | [optional] 

## Methods

### NewReviewReadOnlyFields

`func NewReviewReadOnlyFields() *ReviewReadOnlyFields`

NewReviewReadOnlyFields instantiates a new ReviewReadOnlyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewReadOnlyFieldsWithDefaults

`func NewReviewReadOnlyFieldsWithDefaults() *ReviewReadOnlyFields`

NewReviewReadOnlyFieldsWithDefaults instantiates a new ReviewReadOnlyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *ReviewReadOnlyFields) GetNote() Note`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ReviewReadOnlyFields) GetNoteOk() (*Note, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ReviewReadOnlyFields) SetNote(v Note)`

SetNote sets Note field to given value.

### HasNote

`func (o *ReviewReadOnlyFields) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReviewerGroupProfile

`func (o *ReviewReadOnlyFields) GetReviewerGroupProfile() ReviewerGroupProfile`

GetReviewerGroupProfile returns the ReviewerGroupProfile field if non-nil, zero value otherwise.

### GetReviewerGroupProfileOk

`func (o *ReviewReadOnlyFields) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool)`

GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGroupProfile

`func (o *ReviewReadOnlyFields) SetReviewerGroupProfile(v ReviewerGroupProfile)`

SetReviewerGroupProfile sets ReviewerGroupProfile field to given value.

### HasReviewerGroupProfile

`func (o *ReviewReadOnlyFields) HasReviewerGroupProfile() bool`

HasReviewerGroupProfile returns a boolean if a field has been set.

### GetAllReviewerLevels

`func (o *ReviewReadOnlyFields) GetAllReviewerLevels() []ReviewerLevelInfoFull`

GetAllReviewerLevels returns the AllReviewerLevels field if non-nil, zero value otherwise.

### GetAllReviewerLevelsOk

`func (o *ReviewReadOnlyFields) GetAllReviewerLevelsOk() (*[]ReviewerLevelInfoFull, bool)`

GetAllReviewerLevelsOk returns a tuple with the AllReviewerLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviewerLevels

`func (o *ReviewReadOnlyFields) SetAllReviewerLevels(v []ReviewerLevelInfoFull)`

SetAllReviewerLevels sets AllReviewerLevels field to given value.

### HasAllReviewerLevels

`func (o *ReviewReadOnlyFields) HasAllReviewerLevels() bool`

HasAllReviewerLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


