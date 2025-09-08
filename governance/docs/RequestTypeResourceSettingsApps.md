# RequestTypeResourceSettingsApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TargetResources** | [**[]OktaApplicationResource**](OktaApplicationResource.md) | List of Okta applications | 

## Methods

### NewRequestTypeResourceSettingsApps

`func NewRequestTypeResourceSettingsApps(type_ string, targetResources []OktaApplicationResource, ) *RequestTypeResourceSettingsApps`

NewRequestTypeResourceSettingsApps instantiates a new RequestTypeResourceSettingsApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeResourceSettingsAppsWithDefaults

`func NewRequestTypeResourceSettingsAppsWithDefaults() *RequestTypeResourceSettingsApps`

NewRequestTypeResourceSettingsAppsWithDefaults instantiates a new RequestTypeResourceSettingsApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeResourceSettingsApps) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeResourceSettingsApps) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeResourceSettingsApps) SetType(v string)`

SetType sets Type field to given value.


### GetTargetResources

`func (o *RequestTypeResourceSettingsApps) GetTargetResources() []OktaApplicationResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *RequestTypeResourceSettingsApps) GetTargetResourcesOk() (*[]OktaApplicationResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *RequestTypeResourceSettingsApps) SetTargetResources(v []OktaApplicationResource)`

SetTargetResources sets TargetResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


