# TargetResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The Okta &#x60;app.id&#x60; of the resource.  See [List applications](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Application/#tag/Application/operation/listApplications) to retrieve app IDs.  | 
**Type** | [**ResourceType2**](ResourceType2.md) |  | 

## Methods

### NewTargetResource

`func NewTargetResource(externalId string, type_ ResourceType2, ) *TargetResource`

NewTargetResource instantiates a new TargetResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetResourceWithDefaults

`func NewTargetResourceWithDefaults() *TargetResource`

NewTargetResourceWithDefaults instantiates a new TargetResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *TargetResource) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TargetResource) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TargetResource) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *TargetResource) GetType() ResourceType2`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetResource) GetTypeOk() (*ResourceType2, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetResource) SetType(v ResourceType2)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


