# AutoRemediationSettingsIncludeOnlyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to [**AutoRemediationResourceType**](AutoRemediationResourceType.md) |  | [optional] 
**ResourceId** | Pointer to **string** | The resource ID of the target resource When &#x60;type &#x3D; GROUP&#x60;, it will point to the group ID.  | [optional] 

## Methods

### NewAutoRemediationSettingsIncludeOnlyInner

`func NewAutoRemediationSettingsIncludeOnlyInner() *AutoRemediationSettingsIncludeOnlyInner`

NewAutoRemediationSettingsIncludeOnlyInner instantiates a new AutoRemediationSettingsIncludeOnlyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRemediationSettingsIncludeOnlyInnerWithDefaults

`func NewAutoRemediationSettingsIncludeOnlyInnerWithDefaults() *AutoRemediationSettingsIncludeOnlyInner`

NewAutoRemediationSettingsIncludeOnlyInnerWithDefaults instantiates a new AutoRemediationSettingsIncludeOnlyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceType() AutoRemediationResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceTypeOk() (*AutoRemediationResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AutoRemediationSettingsIncludeOnlyInner) SetResourceType(v AutoRemediationResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AutoRemediationSettingsIncludeOnlyInner) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceId

`func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AutoRemediationSettingsIncludeOnlyInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AutoRemediationSettingsIncludeOnlyInner) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


