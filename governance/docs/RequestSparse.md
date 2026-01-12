# RequestSparse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Whether the request&#39;s request type has CUSTOM settings or not. | 
**RequestStatus** | [**RequestRequestStatus**](RequestRequestStatus.md) |  | 
**Resolved** | **NullableTime** | The date the request was resolved. The property may transition from having a value to null if the request is reopened. | 
**PermalinkId** | **int32** | The immutable, persistent identifier that always resolves to the request | 
**Links** | [**RequestLinks**](RequestLinks.md) |  | 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**RequestTypeId** | **string** | The request type &#x60;id&#x60;.  | 
**Subject** | **string** | The subject of the request | 
**RequesterUserIds** | **[]string** | A list of requester Okta user &#x60;id&#x60;s. | 

## Methods

### NewRequestSparse

`func NewRequestSparse(type_ string, requestStatus RequestRequestStatus, resolved NullableTime, permalinkId int32, links RequestLinks, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, requestTypeId string, subject string, requesterUserIds []string, ) *RequestSparse`

NewRequestSparse instantiates a new RequestSparse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSparseWithDefaults

`func NewRequestSparseWithDefaults() *RequestSparse`

NewRequestSparseWithDefaults instantiates a new RequestSparse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestSparse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestSparse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestSparse) SetType(v string)`

SetType sets Type field to given value.


### GetRequestStatus

`func (o *RequestSparse) GetRequestStatus() RequestRequestStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *RequestSparse) GetRequestStatusOk() (*RequestRequestStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *RequestSparse) SetRequestStatus(v RequestRequestStatus)`

SetRequestStatus sets RequestStatus field to given value.


### GetResolved

`func (o *RequestSparse) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RequestSparse) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RequestSparse) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.


### SetResolvedNil

`func (o *RequestSparse) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *RequestSparse) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetPermalinkId

`func (o *RequestSparse) GetPermalinkId() int32`

GetPermalinkId returns the PermalinkId field if non-nil, zero value otherwise.

### GetPermalinkIdOk

`func (o *RequestSparse) GetPermalinkIdOk() (*int32, bool)`

GetPermalinkIdOk returns a tuple with the PermalinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalinkId

`func (o *RequestSparse) SetPermalinkId(v int32)`

SetPermalinkId sets PermalinkId field to given value.


### GetLinks

`func (o *RequestSparse) GetLinks() RequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestSparse) GetLinksOk() (*RequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestSparse) SetLinks(v RequestLinks)`

SetLinks sets Links field to given value.


### GetId

`func (o *RequestSparse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestSparse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestSparse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestSparse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestSparse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestSparse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestSparse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestSparse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestSparse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestSparse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestSparse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestSparse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestSparse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestSparse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestSparse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetRequestTypeId

`func (o *RequestSparse) GetRequestTypeId() string`

GetRequestTypeId returns the RequestTypeId field if non-nil, zero value otherwise.

### GetRequestTypeIdOk

`func (o *RequestSparse) GetRequestTypeIdOk() (*string, bool)`

GetRequestTypeIdOk returns a tuple with the RequestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypeId

`func (o *RequestSparse) SetRequestTypeId(v string)`

SetRequestTypeId sets RequestTypeId field to given value.


### GetSubject

`func (o *RequestSparse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestSparse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestSparse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetRequesterUserIds

`func (o *RequestSparse) GetRequesterUserIds() []string`

GetRequesterUserIds returns the RequesterUserIds field if non-nil, zero value otherwise.

### GetRequesterUserIdsOk

`func (o *RequestSparse) GetRequesterUserIdsOk() (*[]string, bool)`

GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterUserIds

`func (o *RequestSparse) SetRequesterUserIds(v []string)`

SetRequesterUserIds sets RequesterUserIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


