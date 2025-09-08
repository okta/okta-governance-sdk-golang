# RequestTypeResourceSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TargetResources** | [**[]OktaEntitlementBundleResource**](OktaEntitlementBundleResource.md) | List of Okta entitlement bundles | 

## Methods

### NewRequestTypeResourceSettingsMutable

`func NewRequestTypeResourceSettingsMutable(type_ string, targetResources []OktaEntitlementBundleResource, ) *RequestTypeResourceSettingsMutable`

NewRequestTypeResourceSettingsMutable instantiates a new RequestTypeResourceSettingsMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeResourceSettingsMutableWithDefaults

`func NewRequestTypeResourceSettingsMutableWithDefaults() *RequestTypeResourceSettingsMutable`

NewRequestTypeResourceSettingsMutableWithDefaults instantiates a new RequestTypeResourceSettingsMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeResourceSettingsMutable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeResourceSettingsMutable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeResourceSettingsMutable) SetType(v string)`

SetType sets Type field to given value.


### GetTargetResources

`func (o *RequestTypeResourceSettingsMutable) GetTargetResources() []OktaEntitlementBundleResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *RequestTypeResourceSettingsMutable) GetTargetResourcesOk() (*[]OktaEntitlementBundleResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *RequestTypeResourceSettingsMutable) SetTargetResources(v []OktaEntitlementBundleResource)`

SetTargetResources sets TargetResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


