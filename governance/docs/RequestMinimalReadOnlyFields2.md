# RequestMinimalReadOnlyFields2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RequestStatus**](RequestStatus.md) |  | [optional] 
**Resolved** | Pointer to **NullableTime** | The date the request was resolved. The property may transition from having a value to null if the request is reopened. | [optional] 
**GrantStatus** | Pointer to [**RequestGrantStatus**](RequestGrantStatus.md) |  | [optional] 
**Granted** | Pointer to **NullableTime** | The date the approved access was granted. Only set if request.status is APPROVED. | [optional] 
**RevocationStatus** | Pointer to [**RequestRevocationStatus**](RequestRevocationStatus.md) |  | [optional] 
**Revoked** | Pointer to **NullableTime** | The date the granted access was revoked. Only set if request.grantStatus is GRANTED and request.revocationStatus is REVOKED. | [optional] 
**RequestedBy** | Pointer to [**ClientCredentialPrincipal**](ClientCredentialPrincipal.md) |  | [optional] 
**RequestedFor** | Pointer to [**TargetPrincipal**](TargetPrincipal.md) |  | [optional] 
**Requested** | Pointer to [**Requested**](Requested.md) |  | [optional] 
**Links** | Pointer to [**RequestLinks2**](RequestLinks2.md) |  | [optional] 

## Methods

### NewRequestMinimalReadOnlyFields2

`func NewRequestMinimalReadOnlyFields2() *RequestMinimalReadOnlyFields2`

NewRequestMinimalReadOnlyFields2 instantiates a new RequestMinimalReadOnlyFields2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMinimalReadOnlyFields2WithDefaults

`func NewRequestMinimalReadOnlyFields2WithDefaults() *RequestMinimalReadOnlyFields2`

NewRequestMinimalReadOnlyFields2WithDefaults instantiates a new RequestMinimalReadOnlyFields2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestMinimalReadOnlyFields2) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestMinimalReadOnlyFields2) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestMinimalReadOnlyFields2) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestMinimalReadOnlyFields2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolved

`func (o *RequestMinimalReadOnlyFields2) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RequestMinimalReadOnlyFields2) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RequestMinimalReadOnlyFields2) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *RequestMinimalReadOnlyFields2) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *RequestMinimalReadOnlyFields2) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *RequestMinimalReadOnlyFields2) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetGrantStatus

`func (o *RequestMinimalReadOnlyFields2) GetGrantStatus() RequestGrantStatus`

GetGrantStatus returns the GrantStatus field if non-nil, zero value otherwise.

### GetGrantStatusOk

`func (o *RequestMinimalReadOnlyFields2) GetGrantStatusOk() (*RequestGrantStatus, bool)`

GetGrantStatusOk returns a tuple with the GrantStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantStatus

`func (o *RequestMinimalReadOnlyFields2) SetGrantStatus(v RequestGrantStatus)`

SetGrantStatus sets GrantStatus field to given value.

### HasGrantStatus

`func (o *RequestMinimalReadOnlyFields2) HasGrantStatus() bool`

HasGrantStatus returns a boolean if a field has been set.

### GetGranted

`func (o *RequestMinimalReadOnlyFields2) GetGranted() time.Time`

GetGranted returns the Granted field if non-nil, zero value otherwise.

### GetGrantedOk

`func (o *RequestMinimalReadOnlyFields2) GetGrantedOk() (*time.Time, bool)`

GetGrantedOk returns a tuple with the Granted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranted

`func (o *RequestMinimalReadOnlyFields2) SetGranted(v time.Time)`

SetGranted sets Granted field to given value.

### HasGranted

`func (o *RequestMinimalReadOnlyFields2) HasGranted() bool`

HasGranted returns a boolean if a field has been set.

### SetGrantedNil

`func (o *RequestMinimalReadOnlyFields2) SetGrantedNil(b bool)`

 SetGrantedNil sets the value for Granted to be an explicit nil

### UnsetGranted
`func (o *RequestMinimalReadOnlyFields2) UnsetGranted()`

UnsetGranted ensures that no value is present for Granted, not even an explicit nil
### GetRevocationStatus

`func (o *RequestMinimalReadOnlyFields2) GetRevocationStatus() RequestRevocationStatus`

GetRevocationStatus returns the RevocationStatus field if non-nil, zero value otherwise.

### GetRevocationStatusOk

`func (o *RequestMinimalReadOnlyFields2) GetRevocationStatusOk() (*RequestRevocationStatus, bool)`

GetRevocationStatusOk returns a tuple with the RevocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationStatus

`func (o *RequestMinimalReadOnlyFields2) SetRevocationStatus(v RequestRevocationStatus)`

SetRevocationStatus sets RevocationStatus field to given value.

### HasRevocationStatus

`func (o *RequestMinimalReadOnlyFields2) HasRevocationStatus() bool`

HasRevocationStatus returns a boolean if a field has been set.

### GetRevoked

`func (o *RequestMinimalReadOnlyFields2) GetRevoked() time.Time`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *RequestMinimalReadOnlyFields2) GetRevokedOk() (*time.Time, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *RequestMinimalReadOnlyFields2) SetRevoked(v time.Time)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *RequestMinimalReadOnlyFields2) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### SetRevokedNil

`func (o *RequestMinimalReadOnlyFields2) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *RequestMinimalReadOnlyFields2) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
### GetRequestedBy

`func (o *RequestMinimalReadOnlyFields2) GetRequestedBy() ClientCredentialPrincipal`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *RequestMinimalReadOnlyFields2) GetRequestedByOk() (*ClientCredentialPrincipal, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *RequestMinimalReadOnlyFields2) SetRequestedBy(v ClientCredentialPrincipal)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *RequestMinimalReadOnlyFields2) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetRequestedFor

`func (o *RequestMinimalReadOnlyFields2) GetRequestedFor() TargetPrincipal`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *RequestMinimalReadOnlyFields2) GetRequestedForOk() (*TargetPrincipal, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *RequestMinimalReadOnlyFields2) SetRequestedFor(v TargetPrincipal)`

SetRequestedFor sets RequestedFor field to given value.

### HasRequestedFor

`func (o *RequestMinimalReadOnlyFields2) HasRequestedFor() bool`

HasRequestedFor returns a boolean if a field has been set.

### GetRequested

`func (o *RequestMinimalReadOnlyFields2) GetRequested() Requested`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *RequestMinimalReadOnlyFields2) GetRequestedOk() (*Requested, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *RequestMinimalReadOnlyFields2) SetRequested(v Requested)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *RequestMinimalReadOnlyFields2) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetLinks

`func (o *RequestMinimalReadOnlyFields2) GetLinks() RequestLinks2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestMinimalReadOnlyFields2) GetLinksOk() (*RequestLinks2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestMinimalReadOnlyFields2) SetLinks(v RequestLinks2)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RequestMinimalReadOnlyFields2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


