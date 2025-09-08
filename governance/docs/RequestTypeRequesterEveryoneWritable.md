# RequestTypeRequesterEveryoneWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**RequesterFields** | Pointer to [**[]FieldWritable**](FieldWritable.md) | A list of fields with which to gather input. The order of the field object controls the order with which the fields are presented to users. | [optional] [default to []]

## Methods

### NewRequestTypeRequesterEveryoneWritable

`func NewRequestTypeRequesterEveryoneWritable(type_ string, ) *RequestTypeRequesterEveryoneWritable`

NewRequestTypeRequesterEveryoneWritable instantiates a new RequestTypeRequesterEveryoneWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeRequesterEveryoneWritableWithDefaults

`func NewRequestTypeRequesterEveryoneWritableWithDefaults() *RequestTypeRequesterEveryoneWritable`

NewRequestTypeRequesterEveryoneWritableWithDefaults instantiates a new RequestTypeRequesterEveryoneWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeRequesterEveryoneWritable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeRequesterEveryoneWritable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeRequesterEveryoneWritable) SetType(v string)`

SetType sets Type field to given value.


### GetRequesterFields

`func (o *RequestTypeRequesterEveryoneWritable) GetRequesterFields() []FieldWritable`

GetRequesterFields returns the RequesterFields field if non-nil, zero value otherwise.

### GetRequesterFieldsOk

`func (o *RequestTypeRequesterEveryoneWritable) GetRequesterFieldsOk() (*[]FieldWritable, bool)`

GetRequesterFieldsOk returns a tuple with the RequesterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFields

`func (o *RequestTypeRequesterEveryoneWritable) SetRequesterFields(v []FieldWritable)`

SetRequesterFields sets RequesterFields field to given value.

### HasRequesterFields

`func (o *RequestTypeRequesterEveryoneWritable) HasRequesterFields() bool`

HasRequesterFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


