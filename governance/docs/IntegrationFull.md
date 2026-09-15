# IntegrationFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The integration ID | [optional] 
**Type** | Pointer to [**IntegrationType**](IntegrationType.md) |  | [optional] 
**Status** | Pointer to [**IntegrationStatus**](IntegrationStatus.md) |  | [optional] 
**Links** | Pointer to [**IntegrationLinks**](IntegrationLinks.md) |  | [optional] 

## Methods

### NewIntegrationFull

`func NewIntegrationFull() *IntegrationFull`

NewIntegrationFull instantiates a new IntegrationFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationFullWithDefaults

`func NewIntegrationFullWithDefaults() *IntegrationFull`

NewIntegrationFullWithDefaults instantiates a new IntegrationFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *IntegrationFull) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationFull) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationFull) SetType(v IntegrationType)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationFull) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationFull) GetStatus() IntegrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationFull) GetStatusOk() (*IntegrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationFull) SetStatus(v IntegrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *IntegrationFull) GetLinks() IntegrationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntegrationFull) GetLinksOk() (*IntegrationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntegrationFull) SetLinks(v IntegrationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IntegrationFull) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


