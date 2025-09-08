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

// TeamsList struct for TeamsList
type TeamsList struct {
	// All teams on the current page
	Data                 []Team          `json:"data,omitempty"`
	Links                *TeamsListLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TeamsList TeamsList

// NewTeamsList instantiates a new TeamsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsList() *TeamsList {
	this := TeamsList{}
	return &this
}

// NewTeamsListWithDefaults instantiates a new TeamsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsListWithDefaults() *TeamsList {
	this := TeamsList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamsList) GetData() []Team {
	if o == nil || o.Data == nil {
		var ret []Team
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsList) GetDataOk() ([]Team, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamsList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Team and assigns it to the Data field.
func (o *TeamsList) SetData(v []Team) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TeamsList) GetLinks() TeamsListLinks {
	if o == nil || o.Links == nil {
		var ret TeamsListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsList) GetLinksOk() (*TeamsListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamsList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TeamsListLinks and assigns it to the Links field.
func (o *TeamsList) SetLinks(v TeamsListLinks) {
	o.Links = &v
}

func (o TeamsList) MarshalJSON() ([]byte, error) {
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

func (o *TeamsList) UnmarshalJSON(bytes []byte) (err error) {
	varTeamsList := _TeamsList{}

	err = json.Unmarshal(bytes, &varTeamsList)
	if err == nil {
		*o = TeamsList(varTeamsList)
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

type NullableTeamsList struct {
	value *TeamsList
	isSet bool
}

func (v NullableTeamsList) Get() *TeamsList {
	return v.value
}

func (v *NullableTeamsList) Set(val *TeamsList) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsList) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsList(val *TeamsList) *NullableTeamsList {
	return &NullableTeamsList{value: val, isSet: true}
}

func (v NullableTeamsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
