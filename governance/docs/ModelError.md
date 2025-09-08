# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | An error code unique to the error | 
**ErrorId** | **string** | An error identifier useful for troubleshooting an error with support | 
**ErrorSummary** | **string** | An error code description detailing the error | 
**ErrorLink** | Pointer to **string** | An indicator where to look out to troubleshoot the error | [optional] 
**ErrorCauses** | Pointer to [**[]ErrorCause**](ErrorCause.md) | An optional array of string values that explains possible reasons for the error. | [optional] 

## Methods

### NewModelError

`func NewModelError(errorCode string, errorId string, errorSummary string, ) *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ModelError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ModelError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ModelError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorId

`func (o *ModelError) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ModelError) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ModelError) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.


### GetErrorSummary

`func (o *ModelError) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *ModelError) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *ModelError) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.


### GetErrorLink

`func (o *ModelError) GetErrorLink() string`

GetErrorLink returns the ErrorLink field if non-nil, zero value otherwise.

### GetErrorLinkOk

`func (o *ModelError) GetErrorLinkOk() (*string, bool)`

GetErrorLinkOk returns a tuple with the ErrorLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLink

`func (o *ModelError) SetErrorLink(v string)`

SetErrorLink sets ErrorLink field to given value.

### HasErrorLink

`func (o *ModelError) HasErrorLink() bool`

HasErrorLink returns a boolean if a field has been set.

### GetErrorCauses

`func (o *ModelError) GetErrorCauses() []ErrorCause`

GetErrorCauses returns the ErrorCauses field if non-nil, zero value otherwise.

### GetErrorCausesOk

`func (o *ModelError) GetErrorCausesOk() (*[]ErrorCause, bool)`

GetErrorCausesOk returns a tuple with the ErrorCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCauses

`func (o *ModelError) SetErrorCauses(v []ErrorCause)`

SetErrorCauses sets ErrorCauses field to given value.

### HasErrorCauses

`func (o *ModelError) HasErrorCauses() bool`

HasErrorCauses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


