# RequestTypeRequesterEveryone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**RequesterFields** | Pointer to [**[]Field**](Field.md) | A list of fields with which to gather input. The order of the field object controls the order with which the fields are presented to users. | [optional] [default to []]

## Methods

### NewRequestTypeRequesterEveryone

`func NewRequestTypeRequesterEveryone(type_ string, ) *RequestTypeRequesterEveryone`

NewRequestTypeRequesterEveryone instantiates a new RequestTypeRequesterEveryone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeRequesterEveryoneWithDefaults

`func NewRequestTypeRequesterEveryoneWithDefaults() *RequestTypeRequesterEveryone`

NewRequestTypeRequesterEveryoneWithDefaults instantiates a new RequestTypeRequesterEveryone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeRequesterEveryone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeRequesterEveryone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeRequesterEveryone) SetType(v string)`

SetType sets Type field to given value.


### GetRequesterFields

`func (o *RequestTypeRequesterEveryone) GetRequesterFields() []Field`

GetRequesterFields returns the RequesterFields field if non-nil, zero value otherwise.

### GetRequesterFieldsOk

`func (o *RequestTypeRequesterEveryone) GetRequesterFieldsOk() (*[]Field, bool)`

GetRequesterFieldsOk returns a tuple with the RequesterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFields

`func (o *RequestTypeRequesterEveryone) SetRequesterFields(v []Field)`

SetRequesterFields sets RequesterFields field to given value.

### HasRequesterFields

`func (o *RequestTypeRequesterEveryone) HasRequesterFields() bool`

HasRequesterFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


