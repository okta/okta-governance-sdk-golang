# RcarEntriesListV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RcarEntry**](RcarEntry.md) | Catalog entries list | [optional] 
**Links** | Pointer to [**RcarEntriesLinks**](RcarEntriesLinks.md) |  | [optional] 

## Methods

### NewRcarEntriesListV2

`func NewRcarEntriesListV2() *RcarEntriesListV2`

NewRcarEntriesListV2 instantiates a new RcarEntriesListV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRcarEntriesListV2WithDefaults

`func NewRcarEntriesListV2WithDefaults() *RcarEntriesListV2`

NewRcarEntriesListV2WithDefaults instantiates a new RcarEntriesListV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RcarEntriesListV2) GetData() []RcarEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RcarEntriesListV2) GetDataOk() (*[]RcarEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RcarEntriesListV2) SetData(v []RcarEntry)`

SetData sets Data field to given value.

### HasData

`func (o *RcarEntriesListV2) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *RcarEntriesListV2) GetLinks() RcarEntriesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RcarEntriesListV2) GetLinksOk() (*RcarEntriesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RcarEntriesListV2) SetLinks(v RcarEntriesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RcarEntriesListV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


