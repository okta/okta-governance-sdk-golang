# CollectionResourceRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOrn** | Pointer to **string** | The ORN of the resource in this collection | [optional] 
**AddedToCollectionDate** | Pointer to **time.Time** | When the resource was added to this collection | [optional] 

## Methods

### NewCollectionResourceRelationship

`func NewCollectionResourceRelationship() *CollectionResourceRelationship`

NewCollectionResourceRelationship instantiates a new CollectionResourceRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceRelationshipWithDefaults

`func NewCollectionResourceRelationshipWithDefaults() *CollectionResourceRelationship`

NewCollectionResourceRelationshipWithDefaults instantiates a new CollectionResourceRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOrn

`func (o *CollectionResourceRelationship) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionResourceRelationship) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionResourceRelationship) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.

### HasResourceOrn

`func (o *CollectionResourceRelationship) HasResourceOrn() bool`

HasResourceOrn returns a boolean if a field has been set.

### GetAddedToCollectionDate

`func (o *CollectionResourceRelationship) GetAddedToCollectionDate() time.Time`

GetAddedToCollectionDate returns the AddedToCollectionDate field if non-nil, zero value otherwise.

### GetAddedToCollectionDateOk

`func (o *CollectionResourceRelationship) GetAddedToCollectionDateOk() (*time.Time, bool)`

GetAddedToCollectionDateOk returns a tuple with the AddedToCollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedToCollectionDate

`func (o *CollectionResourceRelationship) SetAddedToCollectionDate(v time.Time)`

SetAddedToCollectionDate sets AddedToCollectionDate field to given value.

### HasAddedToCollectionDate

`func (o *CollectionResourceRelationship) HasAddedToCollectionDate() bool`

HasAddedToCollectionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


