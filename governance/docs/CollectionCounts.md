# CollectionCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalAssignmentCount** | **int32** | Number of principals assigned | 
**ResourceCounts** | [**CollectionCountsResourceCounts**](CollectionCountsResourceCounts.md) |  | 

## Methods

### NewCollectionCounts

`func NewCollectionCounts(principalAssignmentCount int32, resourceCounts CollectionCountsResourceCounts, ) *CollectionCounts`

NewCollectionCounts instantiates a new CollectionCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionCountsWithDefaults

`func NewCollectionCountsWithDefaults() *CollectionCounts`

NewCollectionCountsWithDefaults instantiates a new CollectionCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalAssignmentCount

`func (o *CollectionCounts) GetPrincipalAssignmentCount() int32`

GetPrincipalAssignmentCount returns the PrincipalAssignmentCount field if non-nil, zero value otherwise.

### GetPrincipalAssignmentCountOk

`func (o *CollectionCounts) GetPrincipalAssignmentCountOk() (*int32, bool)`

GetPrincipalAssignmentCountOk returns a tuple with the PrincipalAssignmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalAssignmentCount

`func (o *CollectionCounts) SetPrincipalAssignmentCount(v int32)`

SetPrincipalAssignmentCount sets PrincipalAssignmentCount field to given value.


### GetResourceCounts

`func (o *CollectionCounts) GetResourceCounts() CollectionCountsResourceCounts`

GetResourceCounts returns the ResourceCounts field if non-nil, zero value otherwise.

### GetResourceCountsOk

`func (o *CollectionCounts) GetResourceCountsOk() (*CollectionCountsResourceCounts, bool)`

GetResourceCountsOk returns a tuple with the ResourceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCounts

`func (o *CollectionCounts) SetResourceCounts(v CollectionCountsResourceCounts)`

SetResourceCounts sets ResourceCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


