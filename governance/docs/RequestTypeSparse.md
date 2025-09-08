# RequestTypeSparse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**RequestTypeStatus**](RequestTypeStatus.md) |  | 
**LastUpdateSource** | **string** |  | 
**Links** | [**RequestTypeLinks**](RequestTypeLinks.md) |  | 
**Name** | **string** | Writable unique key on Create. Not modifiable on update. | 
**Description** | **string** | Human readable description. | 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewRequestTypeSparse

`func NewRequestTypeSparse(status RequestTypeStatus, lastUpdateSource string, links RequestTypeLinks, name string, description string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *RequestTypeSparse`

NewRequestTypeSparse instantiates a new RequestTypeSparse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeSparseWithDefaults

`func NewRequestTypeSparseWithDefaults() *RequestTypeSparse`

NewRequestTypeSparseWithDefaults instantiates a new RequestTypeSparse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestTypeSparse) GetStatus() RequestTypeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestTypeSparse) GetStatusOk() (*RequestTypeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestTypeSparse) SetStatus(v RequestTypeStatus)`

SetStatus sets Status field to given value.


### GetLastUpdateSource

`func (o *RequestTypeSparse) GetLastUpdateSource() string`

GetLastUpdateSource returns the LastUpdateSource field if non-nil, zero value otherwise.

### GetLastUpdateSourceOk

`func (o *RequestTypeSparse) GetLastUpdateSourceOk() (*string, bool)`

GetLastUpdateSourceOk returns a tuple with the LastUpdateSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateSource

`func (o *RequestTypeSparse) SetLastUpdateSource(v string)`

SetLastUpdateSource sets LastUpdateSource field to given value.


### GetLinks

`func (o *RequestTypeSparse) GetLinks() RequestTypeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestTypeSparse) GetLinksOk() (*RequestTypeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestTypeSparse) SetLinks(v RequestTypeLinks)`

SetLinks sets Links field to given value.


### GetName

`func (o *RequestTypeSparse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestTypeSparse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestTypeSparse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestTypeSparse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTypeSparse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTypeSparse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *RequestTypeSparse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestTypeSparse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestTypeSparse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestTypeSparse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestTypeSparse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestTypeSparse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestTypeSparse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestTypeSparse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestTypeSparse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestTypeSparse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestTypeSparse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestTypeSparse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestTypeSparse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestTypeSparse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestTypeSparse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


