# ResourceGrantLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**Link**](Link.md) |  | 
**EntitlementBundle** | Pointer to [**EntitlementBundleLink**](EntitlementBundleLink.md) |  | [optional] 

## Methods

### NewResourceGrantLinks

`func NewResourceGrantLinks(self Link, ) *ResourceGrantLinks`

NewResourceGrantLinks instantiates a new ResourceGrantLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGrantLinksWithDefaults

`func NewResourceGrantLinksWithDefaults() *ResourceGrantLinks`

NewResourceGrantLinksWithDefaults instantiates a new ResourceGrantLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ResourceGrantLinks) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ResourceGrantLinks) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ResourceGrantLinks) SetSelf(v Link)`

SetSelf sets Self field to given value.


### GetEntitlementBundle

`func (o *ResourceGrantLinks) GetEntitlementBundle() EntitlementBundleLink`

GetEntitlementBundle returns the EntitlementBundle field if non-nil, zero value otherwise.

### GetEntitlementBundleOk

`func (o *ResourceGrantLinks) GetEntitlementBundleOk() (*EntitlementBundleLink, bool)`

GetEntitlementBundleOk returns a tuple with the EntitlementBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundle

`func (o *ResourceGrantLinks) SetEntitlementBundle(v EntitlementBundleLink)`

SetEntitlementBundle sets EntitlementBundle field to given value.

### HasEntitlementBundle

`func (o *ResourceGrantLinks) HasEntitlementBundle() bool`

HasEntitlementBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


