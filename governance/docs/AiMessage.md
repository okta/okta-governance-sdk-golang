# AiMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Generated summary message | 
**DeltaMessage** | Pointer to **string** | Generated summary message as a delta (for streaming) | [optional] 
**Errors** | Pointer to [**[]AiMessageErrorsInner**](AiMessageErrorsInner.md) | Whenever summary generation has resulted in an error or blocked, the array of errors will detail the reasons. | [optional] 

## Methods

### NewAiMessage

`func NewAiMessage(message string, ) *AiMessage`

NewAiMessage instantiates a new AiMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMessageWithDefaults

`func NewAiMessageWithDefaults() *AiMessage`

NewAiMessageWithDefaults instantiates a new AiMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDeltaMessage

`func (o *AiMessage) GetDeltaMessage() string`

GetDeltaMessage returns the DeltaMessage field if non-nil, zero value otherwise.

### GetDeltaMessageOk

`func (o *AiMessage) GetDeltaMessageOk() (*string, bool)`

GetDeltaMessageOk returns a tuple with the DeltaMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaMessage

`func (o *AiMessage) SetDeltaMessage(v string)`

SetDeltaMessage sets DeltaMessage field to given value.

### HasDeltaMessage

`func (o *AiMessage) HasDeltaMessage() bool`

HasDeltaMessage returns a boolean if a field has been set.

### GetErrors

`func (o *AiMessage) GetErrors() []AiMessageErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AiMessage) GetErrorsOk() (*[]AiMessageErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AiMessage) SetErrors(v []AiMessageErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AiMessage) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


