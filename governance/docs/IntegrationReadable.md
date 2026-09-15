# IntegrationReadable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The integration ID | [optional] 
**Type** | Pointer to [**IntegrationType**](IntegrationType.md) |  | [optional] 
**Status** | Pointer to [**IntegrationStatus**](IntegrationStatus.md) |  | [optional] 

## Methods

### NewIntegrationReadable

`func NewIntegrationReadable() *IntegrationReadable`

NewIntegrationReadable instantiates a new IntegrationReadable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationReadableWithDefaults

`func NewIntegrationReadableWithDefaults() *IntegrationReadable`

NewIntegrationReadableWithDefaults instantiates a new IntegrationReadable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationReadable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationReadable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationReadable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationReadable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *IntegrationReadable) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationReadable) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationReadable) SetType(v IntegrationType)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationReadable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationReadable) GetStatus() IntegrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationReadable) GetStatusOk() (*IntegrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationReadable) SetStatus(v IntegrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationReadable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


