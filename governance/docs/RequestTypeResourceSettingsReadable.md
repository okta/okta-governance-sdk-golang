# RequestTypeResourceSettingsReadable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TargetResources** | [**[]OktaEntitlementBundleResource**](OktaEntitlementBundleResource.md) | List of Okta entitlement bundles | 

## Methods

### NewRequestTypeResourceSettingsReadable

`func NewRequestTypeResourceSettingsReadable(type_ string, targetResources []OktaEntitlementBundleResource, ) *RequestTypeResourceSettingsReadable`

NewRequestTypeResourceSettingsReadable instantiates a new RequestTypeResourceSettingsReadable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeResourceSettingsReadableWithDefaults

`func NewRequestTypeResourceSettingsReadableWithDefaults() *RequestTypeResourceSettingsReadable`

NewRequestTypeResourceSettingsReadableWithDefaults instantiates a new RequestTypeResourceSettingsReadable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeResourceSettingsReadable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeResourceSettingsReadable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeResourceSettingsReadable) SetType(v string)`

SetType sets Type field to given value.


### GetTargetResources

`func (o *RequestTypeResourceSettingsReadable) GetTargetResources() []OktaEntitlementBundleResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *RequestTypeResourceSettingsReadable) GetTargetResourcesOk() (*[]OktaEntitlementBundleResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *RequestTypeResourceSettingsReadable) SetTargetResources(v []OktaEntitlementBundleResource)`

SetTargetResources sets TargetResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


