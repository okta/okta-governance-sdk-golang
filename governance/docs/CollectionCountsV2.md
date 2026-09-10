# CollectionCountsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalAssignmentCount** | **int32** | Number of principals assigned | 
**ResourceCount** | **int32** | The total number of resources in this collection, across all resource types (apps, groups, and push groups). | 

## Methods

### NewCollectionCountsV2

`func NewCollectionCountsV2(principalAssignmentCount int32, resourceCount int32, ) *CollectionCountsV2`

NewCollectionCountsV2 instantiates a new CollectionCountsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionCountsV2WithDefaults

`func NewCollectionCountsV2WithDefaults() *CollectionCountsV2`

NewCollectionCountsV2WithDefaults instantiates a new CollectionCountsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalAssignmentCount

`func (o *CollectionCountsV2) GetPrincipalAssignmentCount() int32`

GetPrincipalAssignmentCount returns the PrincipalAssignmentCount field if non-nil, zero value otherwise.

### GetPrincipalAssignmentCountOk

`func (o *CollectionCountsV2) GetPrincipalAssignmentCountOk() (*int32, bool)`

GetPrincipalAssignmentCountOk returns a tuple with the PrincipalAssignmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalAssignmentCount

`func (o *CollectionCountsV2) SetPrincipalAssignmentCount(v int32)`

SetPrincipalAssignmentCount sets PrincipalAssignmentCount field to given value.


### GetResourceCount

`func (o *CollectionCountsV2) GetResourceCount() int32`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *CollectionCountsV2) GetResourceCountOk() (*int32, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *CollectionCountsV2) SetResourceCount(v int32)`

SetResourceCount sets ResourceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


