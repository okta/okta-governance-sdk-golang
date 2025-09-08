# RequestSequencesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RequestSequence**](RequestSequence.md) | All request sequences | [optional] 
**Links** | Pointer to [**RequestSequencesListLinks**](RequestSequencesListLinks.md) |  | [optional] 

## Methods

### NewRequestSequencesList

`func NewRequestSequencesList() *RequestSequencesList`

NewRequestSequencesList instantiates a new RequestSequencesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSequencesListWithDefaults

`func NewRequestSequencesListWithDefaults() *RequestSequencesList`

NewRequestSequencesListWithDefaults instantiates a new RequestSequencesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RequestSequencesList) GetData() []RequestSequence`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestSequencesList) GetDataOk() (*[]RequestSequence, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestSequencesList) SetData(v []RequestSequence)`

SetData sets Data field to given value.

### HasData

`func (o *RequestSequencesList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *RequestSequencesList) GetLinks() RequestSequencesListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestSequencesList) GetLinksOk() (*RequestSequencesListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestSequencesList) SetLinks(v RequestSequencesListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RequestSequencesList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


