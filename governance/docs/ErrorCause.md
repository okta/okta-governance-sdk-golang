# ErrorCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorSummary** | **string** | A more specific summary for the error containing the error cause. | 
**Reason** | Pointer to **string** | An enumerated value to represent the reason why the error occurred. This enumeration allows codes to adapt to different conditions in which an error code can occur. | [optional] 
**Location** | Pointer to **string** | A value that represents the key where the error cause occurred. This is used with &#x60;locationType&#x60; to give a holistic view of where the error cause occurred. For example, if &#x60;locationType&#x60; is body and the location is username and the reason was UNIQUE_CONSTRAINT, you can derive that the username was already taken. | [optional] 
**LocationType** | Pointer to **string** | A value that represents where the error cause occurred. For example, in the body or header of the request. This value is not required for cases where the request is correct, but there was another reason why the error occurred (server-side state conflict, rate limit violation, and so on) | [optional] 
**Domain** | Pointer to **string** | A value that represents the domain of the service in which the error occurs. This value is used to isolate the error cause reason. | [optional] 

## Methods

### NewErrorCause

`func NewErrorCause(errorSummary string, ) *ErrorCause`

NewErrorCause instantiates a new ErrorCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorCauseWithDefaults

`func NewErrorCauseWithDefaults() *ErrorCause`

NewErrorCauseWithDefaults instantiates a new ErrorCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorSummary

`func (o *ErrorCause) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *ErrorCause) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *ErrorCause) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.


### GetReason

`func (o *ErrorCause) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorCause) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorCause) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorCause) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetLocation

`func (o *ErrorCause) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ErrorCause) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ErrorCause) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ErrorCause) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocationType

`func (o *ErrorCause) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *ErrorCause) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *ErrorCause) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *ErrorCause) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetDomain

`func (o *ErrorCause) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ErrorCause) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ErrorCause) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ErrorCause) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


