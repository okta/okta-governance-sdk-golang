# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Integration ID | [optional] 
**Type** | Pointer to [**IntegrationType**](IntegrationType.md) |  | [optional] 
**Status** | Pointer to [**IntegrationStatus**](IntegrationStatus.md) |  | [optional] 

## Methods

### NewIntegration

`func NewIntegration() *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Integration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Integration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Integration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Integration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Integration) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Integration) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Integration) SetType(v IntegrationType)`

SetType sets Type field to given value.

### HasType

`func (o *Integration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Integration) GetStatus() IntegrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Integration) GetStatusOk() (*IntegrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Integration) SetStatus(v IntegrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Integration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


