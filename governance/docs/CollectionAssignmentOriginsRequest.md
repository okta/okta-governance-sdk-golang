# CollectionAssignmentOriginsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | Pointer to **string** | Anchor principal. Provide together with resourceOrns. | [optional] 
**ResourceOrns** | Pointer to **[]string** | Resource ORNs to check for the anchor principal | [optional] 
**ResourceOrn** | Pointer to **string** | Anchor resource. Provide together with principalOrns. | [optional] 
**PrincipalOrns** | Pointer to **[]string** | Principal ORNs to check for the anchor resource | [optional] 

## Methods

### NewCollectionAssignmentOriginsRequest

`func NewCollectionAssignmentOriginsRequest() *CollectionAssignmentOriginsRequest`

NewCollectionAssignmentOriginsRequest instantiates a new CollectionAssignmentOriginsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionAssignmentOriginsRequestWithDefaults

`func NewCollectionAssignmentOriginsRequestWithDefaults() *CollectionAssignmentOriginsRequest`

NewCollectionAssignmentOriginsRequestWithDefaults instantiates a new CollectionAssignmentOriginsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *CollectionAssignmentOriginsRequest) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.

### HasPrincipalOrn

`func (o *CollectionAssignmentOriginsRequest) HasPrincipalOrn() bool`

HasPrincipalOrn returns a boolean if a field has been set.

### GetResourceOrns

`func (o *CollectionAssignmentOriginsRequest) GetResourceOrns() []string`

GetResourceOrns returns the ResourceOrns field if non-nil, zero value otherwise.

### GetResourceOrnsOk

`func (o *CollectionAssignmentOriginsRequest) GetResourceOrnsOk() (*[]string, bool)`

GetResourceOrnsOk returns a tuple with the ResourceOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrns

`func (o *CollectionAssignmentOriginsRequest) SetResourceOrns(v []string)`

SetResourceOrns sets ResourceOrns field to given value.

### HasResourceOrns

`func (o *CollectionAssignmentOriginsRequest) HasResourceOrns() bool`

HasResourceOrns returns a boolean if a field has been set.

### GetResourceOrn

`func (o *CollectionAssignmentOriginsRequest) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionAssignmentOriginsRequest) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionAssignmentOriginsRequest) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.

### HasResourceOrn

`func (o *CollectionAssignmentOriginsRequest) HasResourceOrn() bool`

HasResourceOrn returns a boolean if a field has been set.

### GetPrincipalOrns

`func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrns() []string`

GetPrincipalOrns returns the PrincipalOrns field if non-nil, zero value otherwise.

### GetPrincipalOrnsOk

`func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrnsOk() (*[]string, bool)`

GetPrincipalOrnsOk returns a tuple with the PrincipalOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrns

`func (o *CollectionAssignmentOriginsRequest) SetPrincipalOrns(v []string)`

SetPrincipalOrns sets PrincipalOrns field to given value.

### HasPrincipalOrns

`func (o *CollectionAssignmentOriginsRequest) HasPrincipalOrns() bool`

HasPrincipalOrns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


