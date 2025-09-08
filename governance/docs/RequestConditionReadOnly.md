# RequestConditionReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the request condition | 
**Status** | [**RequestConditionStatus**](RequestConditionStatus.md) |  | 
**Priority** | **int32** | The priority of the condition. The smaller the number, the higher the priority. The highest priority is 0. A new condition will default to the lowest priority. | 
**Links** | [**RequestConditionLinks**](RequestConditionLinks.md) |  | 

## Methods

### NewRequestConditionReadOnly

`func NewRequestConditionReadOnly(id string, status RequestConditionStatus, priority int32, links RequestConditionLinks, ) *RequestConditionReadOnly`

NewRequestConditionReadOnly instantiates a new RequestConditionReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestConditionReadOnlyWithDefaults

`func NewRequestConditionReadOnlyWithDefaults() *RequestConditionReadOnly`

NewRequestConditionReadOnlyWithDefaults instantiates a new RequestConditionReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestConditionReadOnly) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestConditionReadOnly) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestConditionReadOnly) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *RequestConditionReadOnly) GetStatus() RequestConditionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestConditionReadOnly) GetStatusOk() (*RequestConditionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestConditionReadOnly) SetStatus(v RequestConditionStatus)`

SetStatus sets Status field to given value.


### GetPriority

`func (o *RequestConditionReadOnly) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequestConditionReadOnly) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequestConditionReadOnly) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetLinks

`func (o *RequestConditionReadOnly) GetLinks() RequestConditionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestConditionReadOnly) GetLinksOk() (*RequestConditionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestConditionReadOnly) SetLinks(v RequestConditionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


