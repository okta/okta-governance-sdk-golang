# RequestMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestTypeId** | Pointer to **string** | The request type &#x60;id&#x60;.  | [optional] 
**Subject** | Pointer to **string** | The subject of the request | [optional] 
**RequesterUserIds** | Pointer to **[]string** | A list of requester Okta user &#x60;id&#x60;s. | [optional] 

## Methods

### NewRequestMutable

`func NewRequestMutable() *RequestMutable`

NewRequestMutable instantiates a new RequestMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMutableWithDefaults

`func NewRequestMutableWithDefaults() *RequestMutable`

NewRequestMutableWithDefaults instantiates a new RequestMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestTypeId

`func (o *RequestMutable) GetRequestTypeId() string`

GetRequestTypeId returns the RequestTypeId field if non-nil, zero value otherwise.

### GetRequestTypeIdOk

`func (o *RequestMutable) GetRequestTypeIdOk() (*string, bool)`

GetRequestTypeIdOk returns a tuple with the RequestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypeId

`func (o *RequestMutable) SetRequestTypeId(v string)`

SetRequestTypeId sets RequestTypeId field to given value.

### HasRequestTypeId

`func (o *RequestMutable) HasRequestTypeId() bool`

HasRequestTypeId returns a boolean if a field has been set.

### GetSubject

`func (o *RequestMutable) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestMutable) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestMutable) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RequestMutable) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetRequesterUserIds

`func (o *RequestMutable) GetRequesterUserIds() []string`

GetRequesterUserIds returns the RequesterUserIds field if non-nil, zero value otherwise.

### GetRequesterUserIdsOk

`func (o *RequestMutable) GetRequesterUserIdsOk() (*[]string, bool)`

GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterUserIds

`func (o *RequestMutable) SetRequesterUserIds(v []string)`

SetRequesterUserIds sets RequesterUserIds field to given value.

### HasRequesterUserIds

`func (o *RequestMutable) HasRequesterUserIds() bool`

HasRequesterUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


