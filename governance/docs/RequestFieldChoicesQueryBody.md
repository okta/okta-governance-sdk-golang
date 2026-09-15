# RequestFieldChoicesQueryBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldValues** | Pointer to [**map[string]RequestFieldValuesMapValue**](RequestFieldValuesMapValue.md) | Previously selected field values, keyed by field ID. | [optional] 
**Filter** | Pointer to **string** | SCIM filter to narrow the returned choices. Supported attribute: &#x60;label&#x60;. Supported operators: &#x60;co&#x60; (contains), &#x60;sw&#x60; (starts with).  | [optional] 

## Methods

### NewRequestFieldChoicesQueryBody

`func NewRequestFieldChoicesQueryBody() *RequestFieldChoicesQueryBody`

NewRequestFieldChoicesQueryBody instantiates a new RequestFieldChoicesQueryBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFieldChoicesQueryBodyWithDefaults

`func NewRequestFieldChoicesQueryBodyWithDefaults() *RequestFieldChoicesQueryBody`

NewRequestFieldChoicesQueryBodyWithDefaults instantiates a new RequestFieldChoicesQueryBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldValues

`func (o *RequestFieldChoicesQueryBody) GetFieldValues() map[string]RequestFieldValuesMapValue`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *RequestFieldChoicesQueryBody) GetFieldValuesOk() (*map[string]RequestFieldValuesMapValue, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *RequestFieldChoicesQueryBody) SetFieldValues(v map[string]RequestFieldValuesMapValue)`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *RequestFieldChoicesQueryBody) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetFilter

`func (o *RequestFieldChoicesQueryBody) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RequestFieldChoicesQueryBody) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RequestFieldChoicesQueryBody) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RequestFieldChoicesQueryBody) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


