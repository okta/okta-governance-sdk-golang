/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2025 - Present Okta, Inc.

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
	"fmt"
)

// checks if the DelegateAppointmentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegateAppointmentList{}

// DelegateAppointmentList struct for DelegateAppointmentList
type DelegateAppointmentList struct {
	// Delegate appointments
	Data                 []DelegateAppointment        `json:"data"`
	Links                DelegateAppointmentListLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _DelegateAppointmentList DelegateAppointmentList

// NewDelegateAppointmentList instantiates a new DelegateAppointmentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegateAppointmentList(data []DelegateAppointment, links DelegateAppointmentListLinks) *DelegateAppointmentList {
	this := DelegateAppointmentList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewDelegateAppointmentListWithDefaults instantiates a new DelegateAppointmentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegateAppointmentListWithDefaults() *DelegateAppointmentList {
	this := DelegateAppointmentList{}
	return &this
}

// GetData returns the Data field value
func (o *DelegateAppointmentList) GetData() []DelegateAppointment {
	if o == nil {
		var ret []DelegateAppointment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DelegateAppointmentList) GetDataOk() ([]DelegateAppointment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DelegateAppointmentList) SetData(v []DelegateAppointment) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *DelegateAppointmentList) GetLinks() DelegateAppointmentListLinks {
	if o == nil {
		var ret DelegateAppointmentListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DelegateAppointmentList) GetLinksOk() (*DelegateAppointmentListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DelegateAppointmentList) SetLinks(v DelegateAppointmentListLinks) {
	o.Links = v
}

func (o DelegateAppointmentList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegateAppointmentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DelegateAppointmentList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"_links",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDelegateAppointmentList := _DelegateAppointmentList{}

	err = json.Unmarshal(data, &varDelegateAppointmentList)

	if err != nil {
		return err
	}

	*o = DelegateAppointmentList(varDelegateAppointmentList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDelegateAppointmentList struct {
	value *DelegateAppointmentList
	isSet bool
}

func (v NullableDelegateAppointmentList) Get() *DelegateAppointmentList {
	return v.value
}

func (v *NullableDelegateAppointmentList) Set(val *DelegateAppointmentList) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegateAppointmentList) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegateAppointmentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegateAppointmentList(val *DelegateAppointmentList) *NullableDelegateAppointmentList {
	return &NullableDelegateAppointmentList{value: val, isSet: true}
}

func (v NullableDelegateAppointmentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegateAppointmentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
