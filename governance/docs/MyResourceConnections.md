# MyResourceConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MyResourceConnection**](MyResourceConnection.md) | All connections the agent has established | 
**Links** | [**MyResourceConnectionsLinks**](MyResourceConnectionsLinks.md) |  | 

## Methods

### NewMyResourceConnections

`func NewMyResourceConnections(data []MyResourceConnection, links MyResourceConnectionsLinks, ) *MyResourceConnections`

NewMyResourceConnections instantiates a new MyResourceConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyResourceConnectionsWithDefaults

`func NewMyResourceConnectionsWithDefaults() *MyResourceConnections`

NewMyResourceConnectionsWithDefaults instantiates a new MyResourceConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MyResourceConnections) GetData() []MyResourceConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MyResourceConnections) GetDataOk() (*[]MyResourceConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MyResourceConnections) SetData(v []MyResourceConnection)`

SetData sets Data field to given value.


### GetLinks

`func (o *MyResourceConnections) GetLinks() MyResourceConnectionsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MyResourceConnections) GetLinksOk() (*MyResourceConnectionsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MyResourceConnections) SetLinks(v MyResourceConnectionsLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


