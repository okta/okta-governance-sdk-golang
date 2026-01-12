# AssignmentCandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | ID of the candidate | 
**Type** | [**CandidateType**](CandidateType.md) |  | 
**Orn** | Pointer to **string** | The Okta user or group &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | [optional] 
**Delegator** | Pointer to [**AssignmentCandidateDelegator**](AssignmentCandidateDelegator.md) |  | [optional] 

## Methods

### NewAssignmentCandidate

`func NewAssignmentCandidate(externalId string, type_ CandidateType, ) *AssignmentCandidate`

NewAssignmentCandidate instantiates a new AssignmentCandidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentCandidateWithDefaults

`func NewAssignmentCandidateWithDefaults() *AssignmentCandidate`

NewAssignmentCandidateWithDefaults instantiates a new AssignmentCandidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *AssignmentCandidate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AssignmentCandidate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AssignmentCandidate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *AssignmentCandidate) GetType() CandidateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignmentCandidate) GetTypeOk() (*CandidateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignmentCandidate) SetType(v CandidateType)`

SetType sets Type field to given value.


### GetOrn

`func (o *AssignmentCandidate) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *AssignmentCandidate) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *AssignmentCandidate) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *AssignmentCandidate) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetDelegator

`func (o *AssignmentCandidate) GetDelegator() AssignmentCandidateDelegator`

GetDelegator returns the Delegator field if non-nil, zero value otherwise.

### GetDelegatorOk

`func (o *AssignmentCandidate) GetDelegatorOk() (*AssignmentCandidateDelegator, bool)`

GetDelegatorOk returns a tuple with the Delegator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegator

`func (o *AssignmentCandidate) SetDelegator(v AssignmentCandidateDelegator)`

SetDelegator sets Delegator field to given value.

### HasDelegator

`func (o *AssignmentCandidate) HasDelegator() bool`

HasDelegator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


