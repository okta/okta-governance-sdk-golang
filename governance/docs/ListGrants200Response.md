# ListGrants200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GrantFullWithEntitlements**](GrantFullWithEntitlements.md) |  | [optional] 
**Links** | Pointer to [**GrantListLinks**](GrantListLinks.md) |  | [optional] 

## Methods

### NewListGrants200Response

`func NewListGrants200Response() *ListGrants200Response`

NewListGrants200Response instantiates a new ListGrants200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGrants200ResponseWithDefaults

`func NewListGrants200ResponseWithDefaults() *ListGrants200Response`

NewListGrants200ResponseWithDefaults instantiates a new ListGrants200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListGrants200Response) GetData() []GrantFullWithEntitlements`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListGrants200Response) GetDataOk() (*[]GrantFullWithEntitlements, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListGrants200Response) SetData(v []GrantFullWithEntitlements)`

SetData sets Data field to given value.

### HasData

`func (o *ListGrants200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ListGrants200Response) GetLinks() GrantListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListGrants200Response) GetLinksOk() (*GrantListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListGrants200Response) SetLinks(v GrantListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListGrants200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


