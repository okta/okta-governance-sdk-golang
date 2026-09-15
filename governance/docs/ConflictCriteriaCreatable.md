# ConflictCriteriaCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]CriteriaCreatable**](CriteriaCreatable.md) | A conflict occurs when two criteria evaluate to true in a logical &#x60;AND&#x60; evaluation. The criteria are evaluated in order, and the first criterion that evaluates to false causes the entire &#x60;AND&#x60; evaluation to be false. If both criteria evaluate to true, then a conflict is detected. | [optional] 

## Methods

### NewConflictCriteriaCreatable

`func NewConflictCriteriaCreatable() *ConflictCriteriaCreatable`

NewConflictCriteriaCreatable instantiates a new ConflictCriteriaCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictCriteriaCreatableWithDefaults

`func NewConflictCriteriaCreatableWithDefaults() *ConflictCriteriaCreatable`

NewConflictCriteriaCreatableWithDefaults instantiates a new ConflictCriteriaCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *ConflictCriteriaCreatable) GetAnd() []CriteriaCreatable`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ConflictCriteriaCreatable) GetAndOk() (*[]CriteriaCreatable, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ConflictCriteriaCreatable) SetAnd(v []CriteriaCreatable)`

SetAnd sets And field to given value.

### HasAnd

`func (o *ConflictCriteriaCreatable) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


