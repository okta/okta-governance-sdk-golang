# CampaignsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CampaignSparse**](CampaignSparse.md) | Sparse representation of a Campaign resource.  | 
**Links** | [**CampaignsListLinks**](CampaignsListLinks.md) |  | 

## Methods

### NewCampaignsList

`func NewCampaignsList(data []CampaignSparse, links CampaignsListLinks, ) *CampaignsList`

NewCampaignsList instantiates a new CampaignsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignsListWithDefaults

`func NewCampaignsListWithDefaults() *CampaignsList`

NewCampaignsListWithDefaults instantiates a new CampaignsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CampaignsList) GetData() []CampaignSparse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignsList) GetDataOk() (*[]CampaignSparse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignsList) SetData(v []CampaignSparse)`

SetData sets Data field to given value.


### GetLinks

`func (o *CampaignsList) GetLinks() CampaignsListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CampaignsList) GetLinksOk() (*CampaignsListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CampaignsList) SetLinks(v CampaignsListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


