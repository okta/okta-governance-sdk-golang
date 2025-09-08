# RcarEntry

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

### NewRcarEntry

`func NewRcarEntry(id string, name string, requestable bool, links RcarEntryLinks, ) *RcarEntry`

NewRcarEntry instantiates a new RcarEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcarEntryWithDefaults

`func NewRcarEntryWithDefaults() *RcarEntry`

NewRcarEntryWithDefaults instantiates a new RcarEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RcarEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RcarEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RcarEntry) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RcarEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RcarEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RcarEntry) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RcarEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RcarEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RcarEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RcarEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *RcarEntry) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RcarEntry) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RcarEntry) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RcarEntry) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetParent

`func (o *RcarEntry) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RcarEntry) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RcarEntry) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RcarEntry) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRequestable

`func (o *RcarEntry) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *RcarEntry) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *RcarEntry) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.


### GetCounts

`func (o *RcarEntry) GetCounts() RcarEntryCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *RcarEntry) GetCountsOk() (*RcarEntryCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *RcarEntry) SetCounts(v RcarEntryCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *RcarEntry) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetLinks

`func (o *RcarEntry) GetLinks() RcarEntryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RcarEntry) GetLinksOk() (*RcarEntryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RcarEntry) SetLinks(v RcarEntryLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


