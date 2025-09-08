# RequestTypeResourceSettingsEntitlementBundles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TargetResources** | [**[]OktaEntitlementBundleResource**](OktaEntitlementBundleResource.md) | List of Okta entitlement bundles | 

## Methods

### NewRequestTypeResourceSettingsEntitlementBundles

`func NewRequestTypeResourceSettingsEntitlementBundles(type_ string, targetResources []OktaEntitlementBundleResource, ) *RequestTypeResourceSettingsEntitlementBundles`

NewRequestTypeResourceSettingsEntitlementBundles instantiates a new RequestTypeResourceSettingsEntitlementBundles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeResourceSettingsEntitlementBundlesWithDefaults

`func NewRequestTypeResourceSettingsEntitlementBundlesWithDefaults() *RequestTypeResourceSettingsEntitlementBundles`

NewRequestTypeResourceSettingsEntitlementBundlesWithDefaults instantiates a new RequestTypeResourceSettingsEntitlementBundles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeResourceSettingsEntitlementBundles) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeResourceSettingsEntitlementBundles) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeResourceSettingsEntitlementBundles) SetType(v string)`

SetType sets Type field to given value.


### GetTargetResources

`func (o *RequestTypeResourceSettingsEntitlementBundles) GetTargetResources() []OktaEntitlementBundleResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *RequestTypeResourceSettingsEntitlementBundles) GetTargetResourcesOk() (*[]OktaEntitlementBundleResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *RequestTypeResourceSettingsEntitlementBundles) SetTargetResources(v []OktaEntitlementBundleResource)`

SetTargetResources sets TargetResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


