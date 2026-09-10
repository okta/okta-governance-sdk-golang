# CollectionAssignmentOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | **string** |  | 
**ResourceOrn** | **string** |  | 
**CollectionManaged** | **bool** | True if the principal receives this resource through a collection. | 
**Collection** | Pointer to [**CollectionAssignmentOriginCollection**](CollectionAssignmentOriginCollection.md) |  | [optional] 

## Methods

### NewCollectionAssignmentOrigin

`func NewCollectionAssignmentOrigin(principalOrn string, resourceOrn string, collectionManaged bool, ) *CollectionAssignmentOrigin`

NewCollectionAssignmentOrigin instantiates a new CollectionAssignmentOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionAssignmentOriginWithDefaults

`func NewCollectionAssignmentOriginWithDefaults() *CollectionAssignmentOrigin`

NewCollectionAssignmentOriginWithDefaults instantiates a new CollectionAssignmentOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *CollectionAssignmentOrigin) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *CollectionAssignmentOrigin) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *CollectionAssignmentOrigin) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetResourceOrn

`func (o *CollectionAssignmentOrigin) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionAssignmentOrigin) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionAssignmentOrigin) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetCollectionManaged

`func (o *CollectionAssignmentOrigin) GetCollectionManaged() bool`

GetCollectionManaged returns the CollectionManaged field if non-nil, zero value otherwise.

### GetCollectionManagedOk

`func (o *CollectionAssignmentOrigin) GetCollectionManagedOk() (*bool, bool)`

GetCollectionManagedOk returns a tuple with the CollectionManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionManaged

`func (o *CollectionAssignmentOrigin) SetCollectionManaged(v bool)`

SetCollectionManaged sets CollectionManaged field to given value.


### GetCollection

`func (o *CollectionAssignmentOrigin) GetCollection() CollectionAssignmentOriginCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CollectionAssignmentOrigin) GetCollectionOk() (*CollectionAssignmentOriginCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CollectionAssignmentOrigin) SetCollection(v CollectionAssignmentOriginCollection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *CollectionAssignmentOrigin) HasCollection() bool`

HasCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


