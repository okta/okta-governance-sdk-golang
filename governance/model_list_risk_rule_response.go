/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 3.2.0
Contact: devex-public@okta.com
*/

package governance

import (
	"encoding/json"
)

// ListRiskRuleResponse struct for ListRiskRuleResponse
type ListRiskRuleResponse struct {
	// All risk rules on the current page
	Data                 []RiskRuleResponse `json:"data,omitempty"`
	Links                *ListLinks         `json:"_links,omitempty"`
	Metadata             *ListMetadata      `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListRiskRuleResponse ListRiskRuleResponse

// NewListRiskRuleResponse instantiates a new ListRiskRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRiskRuleResponse() *ListRiskRuleResponse {
	this := ListRiskRuleResponse{}
	return &this
}

// NewListRiskRuleResponseWithDefaults instantiates a new ListRiskRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRiskRuleResponseWithDefaults() *ListRiskRuleResponse {
	this := ListRiskRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListRiskRuleResponse) GetData() []RiskRuleResponse {
	if o == nil || o.Data == nil {
		var ret []RiskRuleResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRiskRuleResponse) GetDataOk() ([]RiskRuleResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListRiskRuleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RiskRuleResponse and assigns it to the Data field.
func (o *ListRiskRuleResponse) SetData(v []RiskRuleResponse) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListRiskRuleResponse) GetLinks() ListLinks {
	if o == nil || o.Links == nil {
		var ret ListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRiskRuleResponse) GetLinksOk() (*ListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListRiskRuleResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ListLinks and assigns it to the Links field.
func (o *ListRiskRuleResponse) SetLinks(v ListLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ListRiskRuleResponse) GetMetadata() ListMetadata {
	if o == nil || o.Metadata == nil {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRiskRuleResponse) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ListRiskRuleResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *ListRiskRuleResponse) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

func (o ListRiskRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListRiskRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varListRiskRuleResponse := _ListRiskRuleResponse{}

	err = json.Unmarshal(bytes, &varListRiskRuleResponse)
	if err == nil {
		*o = ListRiskRuleResponse(varListRiskRuleResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableListRiskRuleResponse struct {
	value *ListRiskRuleResponse
	isSet bool
}

func (v NullableListRiskRuleResponse) Get() *ListRiskRuleResponse {
	return v.value
}

func (v *NullableListRiskRuleResponse) Set(val *ListRiskRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRiskRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRiskRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRiskRuleResponse(val *ListRiskRuleResponse) *NullableListRiskRuleResponse {
	return &NullableListRiskRuleResponse{value: val, isSet: true}
}

func (v NullableListRiskRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRiskRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
