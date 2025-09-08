# ListCatalogUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PrincipalProfile**](PrincipalProfile.md) |  | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 

## Methods

### NewListCatalogUsers

`func NewListCatalogUsers() *ListCatalogUsers`

NewListCatalogUsers instantiates a new ListCatalogUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogUsersWithDefaults

`func NewListCatalogUsersWithDefaults() *ListCatalogUsers`

NewListCatalogUsersWithDefaults instantiates a new ListCatalogUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCatalogUsers) GetData() []PrincipalProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCatalogUsers) GetDataOk() (*[]PrincipalProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCatalogUsers) SetData(v []PrincipalProfile)`

SetData sets Data field to given value.

### HasData

`func (o *ListCatalogUsers) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ListCatalogUsers) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListCatalogUsers) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListCatalogUsers) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListCatalogUsers) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


