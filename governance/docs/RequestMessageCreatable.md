# RequestMessageCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Message that will be created for the request. Newline can be specified by characters \&quot;\\n\&quot;. Message will be visible for all users who can view the request. | 

## Methods

### NewRequestMessageCreatable

`func NewRequestMessageCreatable(message string, ) *RequestMessageCreatable`

NewRequestMessageCreatable instantiates a new RequestMessageCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMessageCreatableWithDefaults

`func NewRequestMessageCreatableWithDefaults() *RequestMessageCreatable`

NewRequestMessageCreatableWithDefaults instantiates a new RequestMessageCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RequestMessageCreatable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestMessageCreatable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestMessageCreatable) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


