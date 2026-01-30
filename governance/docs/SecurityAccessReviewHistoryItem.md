# SecurityAccessReviewHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SystemGenerated** | **bool** | Indicates if the action or change was made by Okta | 
**Timestamp** | **time.Time** | The date and time of the action or change | 
**Message** | **string** |  | 
**PrincipalProfile** | Pointer to [**PrincipalProfileEnriched**](PrincipalProfileEnriched.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewHistoryItem

`func NewSecurityAccessReviewHistoryItem(id string, systemGenerated bool, timestamp time.Time, message string, ) *SecurityAccessReviewHistoryItem`

NewSecurityAccessReviewHistoryItem instantiates a new SecurityAccessReviewHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewHistoryItemWithDefaults

`func NewSecurityAccessReviewHistoryItemWithDefaults() *SecurityAccessReviewHistoryItem`

NewSecurityAccessReviewHistoryItemWithDefaults instantiates a new SecurityAccessReviewHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReviewHistoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReviewHistoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReviewHistoryItem) SetId(v string)`

SetId sets Id field to given value.


### GetSystemGenerated

`func (o *SecurityAccessReviewHistoryItem) GetSystemGenerated() bool`

GetSystemGenerated returns the SystemGenerated field if non-nil, zero value otherwise.

### GetSystemGeneratedOk

`func (o *SecurityAccessReviewHistoryItem) GetSystemGeneratedOk() (*bool, bool)`

GetSystemGeneratedOk returns a tuple with the SystemGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemGenerated

`func (o *SecurityAccessReviewHistoryItem) SetSystemGenerated(v bool)`

SetSystemGenerated sets SystemGenerated field to given value.


### GetTimestamp

`func (o *SecurityAccessReviewHistoryItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SecurityAccessReviewHistoryItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SecurityAccessReviewHistoryItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMessage

`func (o *SecurityAccessReviewHistoryItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityAccessReviewHistoryItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityAccessReviewHistoryItem) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPrincipalProfile

`func (o *SecurityAccessReviewHistoryItem) GetPrincipalProfile() PrincipalProfileEnriched`

GetPrincipalProfile returns the PrincipalProfile field if non-nil, zero value otherwise.

### GetPrincipalProfileOk

`func (o *SecurityAccessReviewHistoryItem) GetPrincipalProfileOk() (*PrincipalProfileEnriched, bool)`

GetPrincipalProfileOk returns a tuple with the PrincipalProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalProfile

`func (o *SecurityAccessReviewHistoryItem) SetPrincipalProfile(v PrincipalProfileEnriched)`

SetPrincipalProfile sets PrincipalProfile field to given value.

### HasPrincipalProfile

`func (o *SecurityAccessReviewHistoryItem) HasPrincipalProfile() bool`

HasPrincipalProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


