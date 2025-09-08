# ListRiskRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RiskRuleResponse**](RiskRuleResponse.md) | All risk rules on the current page | [optional] 
**Links** | Pointer to [**ListLinks**](ListLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 

## Methods

### NewListRiskRuleResponse

`func NewListRiskRuleResponse() *ListRiskRuleResponse`

NewListRiskRuleResponse instantiates a new ListRiskRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRiskRuleResponseWithDefaults

`func NewListRiskRuleResponseWithDefaults() *ListRiskRuleResponse`

NewListRiskRuleResponseWithDefaults instantiates a new ListRiskRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRiskRuleResponse) GetData() []RiskRuleResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRiskRuleResponse) GetDataOk() (*[]RiskRuleResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRiskRuleResponse) SetData(v []RiskRuleResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ListRiskRuleResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *ListRiskRuleResponse) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListRiskRuleResponse) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListRiskRuleResponse) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListRiskRuleResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ListRiskRuleResponse) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListRiskRuleResponse) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListRiskRuleResponse) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListRiskRuleResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


