# ConflictCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]Criteria**](Criteria.md) | A conflict occurs when the logical AND evaluation of the two criteria is true | [optional] 

## Methods

### NewConflictCriteria

`func NewConflictCriteria() *ConflictCriteria`

NewConflictCriteria instantiates a new ConflictCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictCriteriaWithDefaults

`func NewConflictCriteriaWithDefaults() *ConflictCriteria`

NewConflictCriteriaWithDefaults instantiates a new ConflictCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *ConflictCriteria) GetAnd() []Criteria`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ConflictCriteria) GetAndOk() (*[]Criteria, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ConflictCriteria) SetAnd(v []Criteria)`

SetAnd sets And field to given value.

### HasAnd

`func (o *ConflictCriteria) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


