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

// RcarEntriesListV2 struct for RcarEntriesListV2
type RcarEntriesListV2 struct {
	// Catalog entries list
	Data                 []RcarEntry       `json:"data,omitempty"`
	Links                *RcarEntriesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RcarEntriesListV2 RcarEntriesListV2

// NewRcarEntriesListV2 instantiates a new RcarEntriesListV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcarEntriesListV2() *RcarEntriesListV2 {
	this := RcarEntriesListV2{}
	return &this
}

// NewRcarEntriesListV2WithDefaults instantiates a new RcarEntriesListV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcarEntriesListV2WithDefaults() *RcarEntriesListV2 {
	this := RcarEntriesListV2{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RcarEntriesListV2) GetData() []RcarEntry {
	if o == nil || o.Data == nil {
		var ret []RcarEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntriesListV2) GetDataOk() ([]RcarEntry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RcarEntriesListV2) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RcarEntry and assigns it to the Data field.
func (o *RcarEntriesListV2) SetData(v []RcarEntry) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RcarEntriesListV2) GetLinks() RcarEntriesLinks {
	if o == nil || o.Links == nil {
		var ret RcarEntriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntriesListV2) GetLinksOk() (*RcarEntriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RcarEntriesListV2) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RcarEntriesLinks and assigns it to the Links field.
func (o *RcarEntriesListV2) SetLinks(v RcarEntriesLinks) {
	o.Links = &v
}

func (o RcarEntriesListV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RcarEntriesListV2) UnmarshalJSON(bytes []byte) (err error) {
	varRcarEntriesListV2 := _RcarEntriesListV2{}

	err = json.Unmarshal(bytes, &varRcarEntriesListV2)
	if err == nil {
		*o = RcarEntriesListV2(varRcarEntriesListV2)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRcarEntriesListV2 struct {
	value *RcarEntriesListV2
	isSet bool
}

func (v NullableRcarEntriesListV2) Get() *RcarEntriesListV2 {
	return v.value
}

func (v *NullableRcarEntriesListV2) Set(val *RcarEntriesListV2) {
	v.value = val
	v.isSet = true
}

func (v NullableRcarEntriesListV2) IsSet() bool {
	return v.isSet
}

func (v *NullableRcarEntriesListV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcarEntriesListV2(val *RcarEntriesListV2) *NullableRcarEntriesListV2 {
	return &NullableRcarEntriesListV2{value: val, isSet: true}
}

func (v NullableRcarEntriesListV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcarEntriesListV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
