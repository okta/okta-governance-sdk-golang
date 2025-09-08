# EntitlementBundleAccessScopeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**EntitlementBundles** | [**[]EntitlementBundlesArrayFullInner**](EntitlementBundlesArrayFullInner.md) | A request condition may have a zero item array if it is in an INVALID state. Otherwise, there will be at least one item.  | 

## Methods

### NewEntitlementBundleAccessScopeSettings

`func NewEntitlementBundleAccessScopeSettings(type_ string, entitlementBundles []EntitlementBundlesArrayFullInner, ) *EntitlementBundleAccessScopeSettings`

NewEntitlementBundleAccessScopeSettings instantiates a new EntitlementBundleAccessScopeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementBundleAccessScopeSettingsWithDefaults

`func NewEntitlementBundleAccessScopeSettingsWithDefaults() *EntitlementBundleAccessScopeSettings`

NewEntitlementBundleAccessScopeSettingsWithDefaults instantiates a new EntitlementBundleAccessScopeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntitlementBundleAccessScopeSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementBundleAccessScopeSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementBundleAccessScopeSettings) SetType(v string)`

SetType sets Type field to given value.


### GetEntitlementBundles

`func (o *EntitlementBundleAccessScopeSettings) GetEntitlementBundles() []EntitlementBundlesArrayFullInner`

GetEntitlementBundles returns the EntitlementBundles field if non-nil, zero value otherwise.

### GetEntitlementBundlesOk

`func (o *EntitlementBundleAccessScopeSettings) GetEntitlementBundlesOk() (*[]EntitlementBundlesArrayFullInner, bool)`

GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundles

`func (o *EntitlementBundleAccessScopeSettings) SetEntitlementBundles(v []EntitlementBundlesArrayFullInner)`

SetEntitlementBundles sets EntitlementBundles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


