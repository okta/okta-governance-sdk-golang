# RequestOnBehalfOfSettingsPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | Indicates that users who can request this resource could also request for another requester of the same resource. This property can only be &#x60;true&#x60;. If request on behalf of should not be allowed then &#x60;requestOnBehalfOfSettings&#x60; should be nullified. | 
**OnlyFor** | Pointer to [**[]RequestOnBehalfOfType**](RequestOnBehalfOfType.md) | Which requesters the resource requester can request on behalf of. If onlyFor is not specified then any requester may request a resource on the behalf of any other user | [optional] 

## Methods

### NewRequestOnBehalfOfSettingsPatchable

`func NewRequestOnBehalfOfSettingsPatchable(allowed bool, ) *RequestOnBehalfOfSettingsPatchable`

NewRequestOnBehalfOfSettingsPatchable instantiates a new RequestOnBehalfOfSettingsPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestOnBehalfOfSettingsPatchableWithDefaults

`func NewRequestOnBehalfOfSettingsPatchableWithDefaults() *RequestOnBehalfOfSettingsPatchable`

NewRequestOnBehalfOfSettingsPatchableWithDefaults instantiates a new RequestOnBehalfOfSettingsPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *RequestOnBehalfOfSettingsPatchable) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *RequestOnBehalfOfSettingsPatchable) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *RequestOnBehalfOfSettingsPatchable) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetOnlyFor

`func (o *RequestOnBehalfOfSettingsPatchable) GetOnlyFor() []RequestOnBehalfOfType`

GetOnlyFor returns the OnlyFor field if non-nil, zero value otherwise.

### GetOnlyForOk

`func (o *RequestOnBehalfOfSettingsPatchable) GetOnlyForOk() (*[]RequestOnBehalfOfType, bool)`

GetOnlyForOk returns a tuple with the OnlyFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyFor

`func (o *RequestOnBehalfOfSettingsPatchable) SetOnlyFor(v []RequestOnBehalfOfType)`

SetOnlyFor sets OnlyFor field to given value.

### HasOnlyFor

`func (o *RequestOnBehalfOfSettingsPatchable) HasOnlyFor() bool`

HasOnlyFor returns a boolean if a field has been set.

### SetOnlyForNil

`func (o *RequestOnBehalfOfSettingsPatchable) SetOnlyForNil(b bool)`

 SetOnlyForNil sets the value for OnlyFor to be an explicit nil

### UnsetOnlyFor
`func (o *RequestOnBehalfOfSettingsPatchable) UnsetOnlyFor()`

UnsetOnlyFor ensures that no value is present for OnlyFor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


