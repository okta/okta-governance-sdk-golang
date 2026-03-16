# RequestSparse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | [**RequestLinks2**](RequestLinks2.md) |  | 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**Resolved** | Pointer to **NullableTime** | The date the request was resolved. The property may transition from having a value to null if the request is reopened. | [optional] 
**GrantStatus** | Pointer to [**RequestGrantStatus**](RequestGrantStatus.md) |  | [optional] 
**Granted** | Pointer to **NullableTime** | The date the approved access was granted. Only set if request.status is APPROVED. | [optional] 
**RevocationStatus** | Pointer to [**RequestRevocationStatus**](RequestRevocationStatus.md) |  | [optional] 
**Revoked** | Pointer to **NullableTime** | The date the granted access was revoked. Only set if request.grantStatus is GRANTED and request.revocationStatus is REVOKED. | [optional] 
**RequestedBy** | [**ClientCredentialPrincipal**](ClientCredentialPrincipal.md) |  | 
**RequestedFor** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**Requested** | [**Requested**](Requested.md) |  | 

## Methods

### NewRequestSparse2

`func NewRequestSparse2(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks2, status RequestStatus, requestedBy ClientCredentialPrincipal, requestedFor TargetPrincipal, requested Requested, ) *RequestSparse2`

NewRequestSparse2 instantiates a new RequestSparse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSparse2WithDefaults

`func NewRequestSparse2WithDefaults() *RequestSparse2`

NewRequestSparse2WithDefaults instantiates a new RequestSparse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestSparse2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestSparse2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestSparse2) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestSparse2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestSparse2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestSparse2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestSparse2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestSparse2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestSparse2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestSparse2) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestSparse2) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestSparse2) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestSparse2) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestSparse2) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestSparse2) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *RequestSparse2) GetLinks() RequestLinks2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestSparse2) GetLinksOk() (*RequestLinks2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestSparse2) SetLinks(v RequestLinks2)`

SetLinks sets Links field to given value.


### GetStatus

`func (o *RequestSparse2) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestSparse2) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestSparse2) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetResolved

`func (o *RequestSparse2) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RequestSparse2) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RequestSparse2) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *RequestSparse2) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *RequestSparse2) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *RequestSparse2) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetGrantStatus

`func (o *RequestSparse2) GetGrantStatus() RequestGrantStatus`

GetGrantStatus returns the GrantStatus field if non-nil, zero value otherwise.

### GetGrantStatusOk

`func (o *RequestSparse2) GetGrantStatusOk() (*RequestGrantStatus, bool)`

GetGrantStatusOk returns a tuple with the GrantStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantStatus

`func (o *RequestSparse2) SetGrantStatus(v RequestGrantStatus)`

SetGrantStatus sets GrantStatus field to given value.

### HasGrantStatus

`func (o *RequestSparse2) HasGrantStatus() bool`

HasGrantStatus returns a boolean if a field has been set.

### GetGranted

`func (o *RequestSparse2) GetGranted() time.Time`

GetGranted returns the Granted field if non-nil, zero value otherwise.

### GetGrantedOk

`func (o *RequestSparse2) GetGrantedOk() (*time.Time, bool)`

GetGrantedOk returns a tuple with the Granted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranted

`func (o *RequestSparse2) SetGranted(v time.Time)`

SetGranted sets Granted field to given value.

### HasGranted

`func (o *RequestSparse2) HasGranted() bool`

HasGranted returns a boolean if a field has been set.

### SetGrantedNil

`func (o *RequestSparse2) SetGrantedNil(b bool)`

 SetGrantedNil sets the value for Granted to be an explicit nil

### UnsetGranted
`func (o *RequestSparse2) UnsetGranted()`

UnsetGranted ensures that no value is present for Granted, not even an explicit nil
### GetRevocationStatus

`func (o *RequestSparse2) GetRevocationStatus() RequestRevocationStatus`

GetRevocationStatus returns the RevocationStatus field if non-nil, zero value otherwise.

### GetRevocationStatusOk

`func (o *RequestSparse2) GetRevocationStatusOk() (*RequestRevocationStatus, bool)`

GetRevocationStatusOk returns a tuple with the RevocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationStatus

`func (o *RequestSparse2) SetRevocationStatus(v RequestRevocationStatus)`

SetRevocationStatus sets RevocationStatus field to given value.

### HasRevocationStatus

`func (o *RequestSparse2) HasRevocationStatus() bool`

HasRevocationStatus returns a boolean if a field has been set.

### GetRevoked

`func (o *RequestSparse2) GetRevoked() time.Time`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *RequestSparse2) GetRevokedOk() (*time.Time, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *RequestSparse2) SetRevoked(v time.Time)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *RequestSparse2) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### SetRevokedNil

`func (o *RequestSparse2) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *RequestSparse2) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
### GetRequestedBy

`func (o *RequestSparse2) GetRequestedBy() ClientCredentialPrincipal`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *RequestSparse2) GetRequestedByOk() (*ClientCredentialPrincipal, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *RequestSparse2) SetRequestedBy(v ClientCredentialPrincipal)`

SetRequestedBy sets RequestedBy field to given value.


### GetRequestedFor

`func (o *RequestSparse2) GetRequestedFor() TargetPrincipal`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *RequestSparse2) GetRequestedForOk() (*TargetPrincipal, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *RequestSparse2) SetRequestedFor(v TargetPrincipal)`

SetRequestedFor sets RequestedFor field to given value.


### GetRequested

`func (o *RequestSparse2) GetRequested() Requested`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *RequestSparse2) GetRequestedOk() (*Requested, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *RequestSparse2) SetRequested(v Requested)`

SetRequested sets Requested field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


