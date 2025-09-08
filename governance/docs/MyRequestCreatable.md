# MyRequestCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequesterFieldValues** | Pointer to [**[]RequestFieldValue**](RequestFieldValue.md) | The requester input fields required by the approval system.  **Note:** The fields required are determined by the approval system.  For the Okta approval system, the required fields are defined in the approval sequence. Ensure that the requester input fields match up with this definition to avoid request approval flow failure.  For external approval systems, the requester input fields are for recording purposes only and do not affect the approval process.  | [optional] 

## Methods

### NewMyRequestCreatable

`func NewMyRequestCreatable() *MyRequestCreatable`

NewMyRequestCreatable instantiates a new MyRequestCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyRequestCreatableWithDefaults

`func NewMyRequestCreatableWithDefaults() *MyRequestCreatable`

NewMyRequestCreatableWithDefaults instantiates a new MyRequestCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequesterFieldValues

`func (o *MyRequestCreatable) GetRequesterFieldValues() []RequestFieldValue`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *MyRequestCreatable) GetRequesterFieldValuesOk() (*[]RequestFieldValue, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *MyRequestCreatable) SetRequesterFieldValues(v []RequestFieldValue)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.

### HasRequesterFieldValues

`func (o *MyRequestCreatable) HasRequesterFieldValues() bool`

HasRequesterFieldValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


