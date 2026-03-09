# ServerExtendedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Server message with detailed content. | 
**Args** | Pointer to [**[]ServerMessageArgument**](ServerMessageArgument.md) | Dynamic arguments, used to construct the whole message, are supplied as an array of values. | [optional] 
**MessageCode** | **string** | A unique message code. You can use the message code for additional processes, such as localization. | 

## Methods

### NewServerExtendedMessage

`func NewServerExtendedMessage(message string, messageCode string, ) *ServerExtendedMessage`

NewServerExtendedMessage instantiates a new ServerExtendedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerExtendedMessageWithDefaults

`func NewServerExtendedMessageWithDefaults() *ServerExtendedMessage`

NewServerExtendedMessageWithDefaults instantiates a new ServerExtendedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServerExtendedMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerExtendedMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerExtendedMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetArgs

`func (o *ServerExtendedMessage) GetArgs() []ServerMessageArgument`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ServerExtendedMessage) GetArgsOk() (*[]ServerMessageArgument, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ServerExtendedMessage) SetArgs(v []ServerMessageArgument)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ServerExtendedMessage) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetMessageCode

`func (o *ServerExtendedMessage) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *ServerExtendedMessage) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *ServerExtendedMessage) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


