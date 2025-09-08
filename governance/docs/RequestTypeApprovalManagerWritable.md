# RequestTypeApprovalManagerWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ApproverFields** | Pointer to [**[]FieldWritable**](FieldWritable.md) | A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.  | [optional] [default to []]

## Methods

### NewRequestTypeApprovalManagerWritable

`func NewRequestTypeApprovalManagerWritable(approverType string, ) *RequestTypeApprovalManagerWritable`

NewRequestTypeApprovalManagerWritable instantiates a new RequestTypeApprovalManagerWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeApprovalManagerWritableWithDefaults

`func NewRequestTypeApprovalManagerWritableWithDefaults() *RequestTypeApprovalManagerWritable`

NewRequestTypeApprovalManagerWritableWithDefaults instantiates a new RequestTypeApprovalManagerWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverType

`func (o *RequestTypeApprovalManagerWritable) GetApproverType() string`

GetApproverType returns the ApproverType field if non-nil, zero value otherwise.

### GetApproverTypeOk

`func (o *RequestTypeApprovalManagerWritable) GetApproverTypeOk() (*string, bool)`

GetApproverTypeOk returns a tuple with the ApproverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverType

`func (o *RequestTypeApprovalManagerWritable) SetApproverType(v string)`

SetApproverType sets ApproverType field to given value.


### GetDescription

`func (o *RequestTypeApprovalManagerWritable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTypeApprovalManagerWritable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTypeApprovalManagerWritable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTypeApprovalManagerWritable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApproverFields

`func (o *RequestTypeApprovalManagerWritable) GetApproverFields() []FieldWritable`

GetApproverFields returns the ApproverFields field if non-nil, zero value otherwise.

### GetApproverFieldsOk

`func (o *RequestTypeApprovalManagerWritable) GetApproverFieldsOk() (*[]FieldWritable, bool)`

GetApproverFieldsOk returns a tuple with the ApproverFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverFields

`func (o *RequestTypeApprovalManagerWritable) SetApproverFields(v []FieldWritable)`

SetApproverFields sets ApproverFields field to given value.

### HasApproverFields

`func (o *RequestTypeApprovalManagerWritable) HasApproverFields() bool`

HasApproverFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


