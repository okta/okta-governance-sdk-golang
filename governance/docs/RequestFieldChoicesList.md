# RequestFieldChoicesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RequestFieldChoice**](RequestFieldChoice.md) |  | 
**Links** | [**RequestFieldChoicesListLinks**](RequestFieldChoicesListLinks.md) |  | 

## Methods

### NewRequestFieldChoicesList

`func NewRequestFieldChoicesList(data []RequestFieldChoice, links RequestFieldChoicesListLinks, ) *RequestFieldChoicesList`

NewRequestFieldChoicesList instantiates a new RequestFieldChoicesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFieldChoicesListWithDefaults

`func NewRequestFieldChoicesListWithDefaults() *RequestFieldChoicesList`

NewRequestFieldChoicesListWithDefaults instantiates a new RequestFieldChoicesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RequestFieldChoicesList) GetData() []RequestFieldChoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestFieldChoicesList) GetDataOk() (*[]RequestFieldChoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestFieldChoicesList) SetData(v []RequestFieldChoice)`

SetData sets Data field to given value.


### GetLinks

`func (o *RequestFieldChoicesList) GetLinks() RequestFieldChoicesListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestFieldChoicesList) GetLinksOk() (*RequestFieldChoicesListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestFieldChoicesList) SetLinks(v RequestFieldChoicesListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


