# RequestCreatable2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | [**RequestResourceCatalogEntryCreatable**](RequestResourceCatalogEntryCreatable.md) |  | 
**RequestedBy** | Pointer to [**TargetPrincipal**](TargetPrincipal.md) |  | [optional] 
**RequestedFor** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**RequesterFieldValues** | Pointer to [**[]RequestFieldValue**](RequestFieldValue.md) | The requester input fields required by the approval system.  **Note:** The fields required are determined by the approval system.  For the Okta approval system, the required fields are defined in the approval sequence. Ensure that the requester input fields match up with this definition to avoid request approval flow failure.  For external approval systems, the requester input fields are for recording purposes only and do not affect the approval process.  | [optional] 

## Methods

### NewRequestCreatable2

`func NewRequestCreatable2(requested RequestResourceCatalogEntryCreatable, requestedFor TargetPrincipal, ) *RequestCreatable2`

NewRequestCreatable2 instantiates a new RequestCreatable2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreatable2WithDefaults

`func NewRequestCreatable2WithDefaults() *RequestCreatable2`

NewRequestCreatable2WithDefaults instantiates a new RequestCreatable2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequested

`func (o *RequestCreatable2) GetRequested() RequestResourceCatalogEntryCreatable`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *RequestCreatable2) GetRequestedOk() (*RequestResourceCatalogEntryCreatable, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *RequestCreatable2) SetRequested(v RequestResourceCatalogEntryCreatable)`

SetRequested sets Requested field to given value.


### GetRequestedBy

`func (o *RequestCreatable2) GetRequestedBy() TargetPrincipal`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *RequestCreatable2) GetRequestedByOk() (*TargetPrincipal, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *RequestCreatable2) SetRequestedBy(v TargetPrincipal)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *RequestCreatable2) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetRequestedFor

`func (o *RequestCreatable2) GetRequestedFor() TargetPrincipal`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *RequestCreatable2) GetRequestedForOk() (*TargetPrincipal, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *RequestCreatable2) SetRequestedFor(v TargetPrincipal)`

SetRequestedFor sets RequestedFor field to given value.


### GetRequesterFieldValues

`func (o *RequestCreatable2) GetRequesterFieldValues() []RequestFieldValue`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *RequestCreatable2) GetRequesterFieldValuesOk() (*[]RequestFieldValue, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *RequestCreatable2) SetRequesterFieldValues(v []RequestFieldValue)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.

### HasRequesterFieldValues

`func (o *RequestCreatable2) HasRequesterFieldValues() bool`

HasRequesterFieldValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


