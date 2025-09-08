# ServerMessageArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Type** | [**ServerMessageArgumentType**](ServerMessageArgumentType.md) |  | 

## Methods

### NewServerMessageArgument

`func NewServerMessageArgument(value string, type_ ServerMessageArgumentType, ) *ServerMessageArgument`

NewServerMessageArgument instantiates a new ServerMessageArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMessageArgumentWithDefaults

`func NewServerMessageArgumentWithDefaults() *ServerMessageArgument`

NewServerMessageArgumentWithDefaults instantiates a new ServerMessageArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ServerMessageArgument) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServerMessageArgument) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServerMessageArgument) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *ServerMessageArgument) GetType() ServerMessageArgumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerMessageArgument) GetTypeOk() (*ServerMessageArgumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerMessageArgument) SetType(v ServerMessageArgumentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


