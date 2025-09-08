# DelegateAppointmentDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The Okta user &#x60;id&#x60; | 
**Type** | [**PrincipalType**](PrincipalType.md) |  | 

## Methods

### NewDelegateAppointmentDelegate

`func NewDelegateAppointmentDelegate(externalId string, type_ PrincipalType, ) *DelegateAppointmentDelegate`

NewDelegateAppointmentDelegate instantiates a new DelegateAppointmentDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegateAppointmentDelegateWithDefaults

`func NewDelegateAppointmentDelegateWithDefaults() *DelegateAppointmentDelegate`

NewDelegateAppointmentDelegateWithDefaults instantiates a new DelegateAppointmentDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *DelegateAppointmentDelegate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DelegateAppointmentDelegate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DelegateAppointmentDelegate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *DelegateAppointmentDelegate) GetType() PrincipalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DelegateAppointmentDelegate) GetTypeOk() (*PrincipalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DelegateAppointmentDelegate) SetType(v PrincipalType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


