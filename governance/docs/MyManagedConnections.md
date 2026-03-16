# MyManagedConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MyManagedConnection**](MyManagedConnection.md) | All connections the agent has established | 
**Links** | [**MyManagedConnectionsLinks**](MyManagedConnectionsLinks.md) |  | 

## Methods

### NewMyManagedConnections

`func NewMyManagedConnections(data []MyManagedConnection, links MyManagedConnectionsLinks, ) *MyManagedConnections`

NewMyManagedConnections instantiates a new MyManagedConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyManagedConnectionsWithDefaults

`func NewMyManagedConnectionsWithDefaults() *MyManagedConnections`

NewMyManagedConnectionsWithDefaults instantiates a new MyManagedConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MyManagedConnections) GetData() []MyManagedConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MyManagedConnections) GetDataOk() (*[]MyManagedConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MyManagedConnections) SetData(v []MyManagedConnection)`

SetData sets Data field to given value.


### GetLinks

`func (o *MyManagedConnections) GetLinks() MyManagedConnectionsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MyManagedConnections) GetLinksOk() (*MyManagedConnectionsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MyManagedConnections) SetLinks(v MyManagedConnectionsLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


