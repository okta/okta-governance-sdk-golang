# AiMessageErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | A unique error code. Error codes are useful to execute additional actions. | 
**ErrorMsg** | **string** | An error message detailing the problem | 
**Args** | Pointer to [**[]ServerMessageArgument**](ServerMessageArgument.md) | Any dynamic arguments that are needed to construct the whole message are supplied as an array of values. | [optional] 

## Methods

### NewAiMessageErrorsInner

`func NewAiMessageErrorsInner(errorCode string, errorMsg string, ) *AiMessageErrorsInner`

NewAiMessageErrorsInner instantiates a new AiMessageErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMessageErrorsInnerWithDefaults

`func NewAiMessageErrorsInnerWithDefaults() *AiMessageErrorsInner`

NewAiMessageErrorsInnerWithDefaults instantiates a new AiMessageErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *AiMessageErrorsInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AiMessageErrorsInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AiMessageErrorsInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMsg

`func (o *AiMessageErrorsInner) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *AiMessageErrorsInner) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *AiMessageErrorsInner) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.


### GetArgs

`func (o *AiMessageErrorsInner) GetArgs() []ServerMessageArgument`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AiMessageErrorsInner) GetArgsOk() (*[]ServerMessageArgument, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AiMessageErrorsInner) SetArgs(v []ServerMessageArgument)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AiMessageErrorsInner) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


