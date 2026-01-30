# AssignmentCandidateDelegator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The Okta user &#x60;id&#x60; | 
**Type** | [**PrincipalType**](PrincipalType.md) |  | 
**Orn** | Pointer to **string** | The delegator &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | [optional] 

## Methods

### NewAssignmentCandidateDelegator

`func NewAssignmentCandidateDelegator(externalId string, type_ PrincipalType, ) *AssignmentCandidateDelegator`

NewAssignmentCandidateDelegator instantiates a new AssignmentCandidateDelegator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentCandidateDelegatorWithDefaults

`func NewAssignmentCandidateDelegatorWithDefaults() *AssignmentCandidateDelegator`

NewAssignmentCandidateDelegatorWithDefaults instantiates a new AssignmentCandidateDelegator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *AssignmentCandidateDelegator) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AssignmentCandidateDelegator) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AssignmentCandidateDelegator) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *AssignmentCandidateDelegator) GetType() PrincipalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignmentCandidateDelegator) GetTypeOk() (*PrincipalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignmentCandidateDelegator) SetType(v PrincipalType)`

SetType sets Type field to given value.


### GetOrn

`func (o *AssignmentCandidateDelegator) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *AssignmentCandidateDelegator) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *AssignmentCandidateDelegator) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *AssignmentCandidateDelegator) HasOrn() bool`

HasOrn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


