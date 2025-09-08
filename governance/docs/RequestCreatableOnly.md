# RequestCreatableOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequesterFieldValues** | Pointer to [**[]FieldValueWritable**](FieldValueWritable.md) | Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.  | [optional] 

## Methods

### NewRequestCreatableOnly

`func NewRequestCreatableOnly() *RequestCreatableOnly`

NewRequestCreatableOnly instantiates a new RequestCreatableOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreatableOnlyWithDefaults

`func NewRequestCreatableOnlyWithDefaults() *RequestCreatableOnly`

NewRequestCreatableOnlyWithDefaults instantiates a new RequestCreatableOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequesterFieldValues

`func (o *RequestCreatableOnly) GetRequesterFieldValues() []FieldValueWritable`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *RequestCreatableOnly) GetRequesterFieldValuesOk() (*[]FieldValueWritable, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *RequestCreatableOnly) SetRequesterFieldValues(v []FieldValueWritable)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.

### HasRequesterFieldValues

`func (o *RequestCreatableOnly) HasRequesterFieldValues() bool`

HasRequesterFieldValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


