# ServerMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Server message with detailed content. | 
**Args** | Pointer to [**[]ServerMessageArgument**](ServerMessageArgument.md) | Dynamic arguments, used to construct the whole message, are supplied as an array of values. | [optional] 

## Methods

### NewServerMessage

`func NewServerMessage(message string, ) *ServerMessage`

NewServerMessage instantiates a new ServerMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMessageWithDefaults

`func NewServerMessageWithDefaults() *ServerMessage`

NewServerMessageWithDefaults instantiates a new ServerMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServerMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetArgs

`func (o *ServerMessage) GetArgs() []ServerMessageArgument`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ServerMessage) GetArgsOk() (*[]ServerMessageArgument, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ServerMessage) SetArgs(v []ServerMessageArgument)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ServerMessage) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


