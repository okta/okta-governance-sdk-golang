# RcarEntryGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the entry | 
**Name** | **string** | Name of the entry | 
**Description** | Pointer to **string** | Description of the entry | [optional] 
**Label** | Pointer to **string** | Label of the entry. Currently either &#x60;Application&#x60; or &#x60;Resource Collection&#x60; | [optional] 
**Parent** | Pointer to **string** | Parent of the entry | [optional] 
**Requestable** | **bool** | Is the resource requestable | 
**Counts** | Pointer to [**RcarEntryCounts**](RcarEntryCounts.md) |  | [optional] 
**Links** | [**RcarEntryLinks**](RcarEntryLinks.md) |  | 

## Methods

### NewRcarEntryGet

`func NewRcarEntryGet(id string, name string, requestable bool, links RcarEntryLinks, ) *RcarEntryGet`

NewRcarEntryGet instantiates a new RcarEntryGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcarEntryGetWithDefaults

`func NewRcarEntryGetWithDefaults() *RcarEntryGet`

NewRcarEntryGetWithDefaults instantiates a new RcarEntryGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcarEntryGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcarEntryGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcarEntryGet) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RcarEntryGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RcarEntryGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RcarEntryGet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RcarEntryGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RcarEntryGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RcarEntryGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RcarEntryGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *RcarEntryGet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RcarEntryGet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RcarEntryGet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RcarEntryGet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetParent

`func (o *RcarEntryGet) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RcarEntryGet) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RcarEntryGet) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RcarEntryGet) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRequestable

`func (o *RcarEntryGet) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *RcarEntryGet) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *RcarEntryGet) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.


### GetCounts

`func (o *RcarEntryGet) GetCounts() RcarEntryCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *RcarEntryGet) GetCountsOk() (*RcarEntryCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *RcarEntryGet) SetCounts(v RcarEntryCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *RcarEntryGet) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetLinks

`func (o *RcarEntryGet) GetLinks() RcarEntryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RcarEntryGet) GetLinksOk() (*RcarEntryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RcarEntryGet) SetLinks(v RcarEntryLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


