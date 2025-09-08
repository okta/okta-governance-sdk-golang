# Requested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | Pointer to **string** | The ID of the resource catalog entry | [optional] 
**ResourceId** | Pointer to **string** | the requested resource ID | [optional] 
**ResourceType** | Pointer to [**ResourceType3**](ResourceType3.md) |  | [optional] 
**AccessScopeId** | Pointer to **string** | ID of the access scope | [optional] 
**AccessScopeType** | Pointer to [**AccessScopeType**](AccessScopeType.md) |  | [optional] 

## Methods

### NewRequested

`func NewRequested() *Requested`

NewRequested instantiates a new Requested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedWithDefaults

`func NewRequestedWithDefaults() *Requested`

NewRequestedWithDefaults instantiates a new Requested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *Requested) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *Requested) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *Requested) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *Requested) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetResourceId

`func (o *Requested) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Requested) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Requested) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Requested) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *Requested) GetResourceType() ResourceType3`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Requested) GetResourceTypeOk() (*ResourceType3, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Requested) SetResourceType(v ResourceType3)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Requested) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetAccessScopeId

`func (o *Requested) GetAccessScopeId() string`

GetAccessScopeId returns the AccessScopeId field if non-nil, zero value otherwise.

### GetAccessScopeIdOk

`func (o *Requested) GetAccessScopeIdOk() (*string, bool)`

GetAccessScopeIdOk returns a tuple with the AccessScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScopeId

`func (o *Requested) SetAccessScopeId(v string)`

SetAccessScopeId sets AccessScopeId field to given value.

### HasAccessScopeId

`func (o *Requested) HasAccessScopeId() bool`

HasAccessScopeId returns a boolean if a field has been set.

### GetAccessScopeType

`func (o *Requested) GetAccessScopeType() AccessScopeType`

GetAccessScopeType returns the AccessScopeType field if non-nil, zero value otherwise.

### GetAccessScopeTypeOk

`func (o *Requested) GetAccessScopeTypeOk() (*AccessScopeType, bool)`

GetAccessScopeTypeOk returns a tuple with the AccessScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScopeType

`func (o *Requested) SetAccessScopeType(v AccessScopeType)`

SetAccessScopeType sets AccessScopeType field to given value.

### HasAccessScopeType

`func (o *Requested) HasAccessScopeType() bool`

HasAccessScopeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


