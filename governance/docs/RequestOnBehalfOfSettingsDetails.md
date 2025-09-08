# RequestOnBehalfOfSettingsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | Indicates that users who can request this resource could also request for another requester of the same resource | 
**OnlyFor** | Pointer to [**[]RequestOnBehalfOfType**](RequestOnBehalfOfType.md) | Which requesters the resource requester can request on behalf of. If onlyFor is not specified then any requester may request a resource on the behalf of any other user | [optional] 

## Methods

### NewRequestOnBehalfOfSettingsDetails

`func NewRequestOnBehalfOfSettingsDetails(allowed bool, ) *RequestOnBehalfOfSettingsDetails`

NewRequestOnBehalfOfSettingsDetails instantiates a new RequestOnBehalfOfSettingsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestOnBehalfOfSettingsDetailsWithDefaults

`func NewRequestOnBehalfOfSettingsDetailsWithDefaults() *RequestOnBehalfOfSettingsDetails`

NewRequestOnBehalfOfSettingsDetailsWithDefaults instantiates a new RequestOnBehalfOfSettingsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *RequestOnBehalfOfSettingsDetails) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *RequestOnBehalfOfSettingsDetails) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *RequestOnBehalfOfSettingsDetails) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetOnlyFor

`func (o *RequestOnBehalfOfSettingsDetails) GetOnlyFor() []RequestOnBehalfOfType`

GetOnlyFor returns the OnlyFor field if non-nil, zero value otherwise.

### GetOnlyForOk

`func (o *RequestOnBehalfOfSettingsDetails) GetOnlyForOk() (*[]RequestOnBehalfOfType, bool)`

GetOnlyForOk returns a tuple with the OnlyFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyFor

`func (o *RequestOnBehalfOfSettingsDetails) SetOnlyFor(v []RequestOnBehalfOfType)`

SetOnlyFor sets OnlyFor field to given value.

### HasOnlyFor

`func (o *RequestOnBehalfOfSettingsDetails) HasOnlyFor() bool`

HasOnlyFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


