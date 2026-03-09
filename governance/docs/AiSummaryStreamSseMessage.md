# AiSummaryStreamSseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The event ID | [optional] 
**Event** | Pointer to **string** | The event type | [optional] 
**Data** | **string** | The event data, in JSON string format | 

## Methods

### NewAiSummaryStreamSseMessage

`func NewAiSummaryStreamSseMessage(data string, ) *AiSummaryStreamSseMessage`

NewAiSummaryStreamSseMessage instantiates a new AiSummaryStreamSseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSummaryStreamSseMessageWithDefaults

`func NewAiSummaryStreamSseMessageWithDefaults() *AiSummaryStreamSseMessage`

NewAiSummaryStreamSseMessageWithDefaults instantiates a new AiSummaryStreamSseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiSummaryStreamSseMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiSummaryStreamSseMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiSummaryStreamSseMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiSummaryStreamSseMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvent

`func (o *AiSummaryStreamSseMessage) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AiSummaryStreamSseMessage) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AiSummaryStreamSseMessage) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AiSummaryStreamSseMessage) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetData

`func (o *AiSummaryStreamSseMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiSummaryStreamSseMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiSummaryStreamSseMessage) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


