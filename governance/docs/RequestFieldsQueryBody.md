# RequestFieldsQueryBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldValues** | Pointer to [**map[string]RequestFieldValuesMapValue**](RequestFieldValuesMapValue.md) | Previously selected field values, keyed by field ID. | [optional] 

## Methods

### NewRequestFieldsQueryBody

`func NewRequestFieldsQueryBody() *RequestFieldsQueryBody`

NewRequestFieldsQueryBody instantiates a new RequestFieldsQueryBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFieldsQueryBodyWithDefaults

`func NewRequestFieldsQueryBodyWithDefaults() *RequestFieldsQueryBody`

NewRequestFieldsQueryBodyWithDefaults instantiates a new RequestFieldsQueryBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldValues

`func (o *RequestFieldsQueryBody) GetFieldValues() map[string]RequestFieldValuesMapValue`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *RequestFieldsQueryBody) GetFieldValuesOk() (*map[string]RequestFieldValuesMapValue, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *RequestFieldsQueryBody) SetFieldValues(v map[string]RequestFieldValuesMapValue)`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *RequestFieldsQueryBody) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


