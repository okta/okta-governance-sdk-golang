# PushGroupFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Okta Group ID for the push group | [optional] 
**Name** | Pointer to **string** | The name of the push group | [optional] 
**Description** | Pointer to **string** | The description of the push group | [optional] 
**Logo** | Pointer to [**[]Link**](Link.md) | List of push group logo resources | [optional] 

## Methods

### NewPushGroupFull

`func NewPushGroupFull() *PushGroupFull`

NewPushGroupFull instantiates a new PushGroupFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushGroupFullWithDefaults

`func NewPushGroupFullWithDefaults() *PushGroupFull`

NewPushGroupFullWithDefaults instantiates a new PushGroupFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PushGroupFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushGroupFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushGroupFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PushGroupFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PushGroupFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushGroupFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushGroupFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushGroupFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PushGroupFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PushGroupFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PushGroupFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PushGroupFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogo

`func (o *PushGroupFull) GetLogo() []Link`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PushGroupFull) GetLogoOk() (*[]Link, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PushGroupFull) SetLogo(v []Link)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PushGroupFull) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


