# OktaResourceReadable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** | Human readable name of the resource | [optional] 
**ResourceId** | Pointer to **string** | The Okta &#x60;app.id&#x60;, or &#x60;group.id&#x60; of the resource that can be requested with this Request Type.  See [list applications](https://developer.okta.com/docs/reference/api/apps/#list-applications) endpoint for reference on how to retrieve application IDs.  See [list groups](https://developer.okta.com/docs/reference/api/groups/#list-groups) endpoint for reference on how to retrieve group IDs.  | [optional] 
**ResourceType** | Pointer to **string** | The type of resource | [optional] 

## Methods

### NewOktaResourceReadable

`func NewOktaResourceReadable() *OktaResourceReadable`

NewOktaResourceReadable instantiates a new OktaResourceReadable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaResourceReadableWithDefaults

`func NewOktaResourceReadableWithDefaults() *OktaResourceReadable`

NewOktaResourceReadableWithDefaults instantiates a new OktaResourceReadable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *OktaResourceReadable) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *OktaResourceReadable) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *OktaResourceReadable) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *OktaResourceReadable) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceId

`func (o *OktaResourceReadable) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OktaResourceReadable) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OktaResourceReadable) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *OktaResourceReadable) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *OktaResourceReadable) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *OktaResourceReadable) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *OktaResourceReadable) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *OktaResourceReadable) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


