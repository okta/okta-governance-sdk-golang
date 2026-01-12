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

// checks if the MyManagedConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MyManagedConnections{}

// MyManagedConnections struct for MyManagedConnections
type MyManagedConnections struct {
	// All connections the agent has established
	Data                 []MyManagedConnection     `json:"data"`
	Links                MyManagedConnectionsLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _MyManagedConnections MyManagedConnections

// NewMyManagedConnections instantiates a new MyManagedConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyManagedConnections(data []MyManagedConnection, links MyManagedConnectionsLinks) *MyManagedConnections {
	this := MyManagedConnections{}
	this.Data = data
	this.Links = links
	return &this
}

// NewMyManagedConnectionsWithDefaults instantiates a new MyManagedConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyManagedConnectionsWithDefaults() *MyManagedConnections {
	this := MyManagedConnections{}
	return &this
}

// GetData returns the Data field value
func (o *MyManagedConnections) GetData() []MyManagedConnection {
	if o == nil {
		var ret []MyManagedConnection
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MyManagedConnections) GetDataOk() ([]MyManagedConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *MyManagedConnections) SetData(v []MyManagedConnection) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *MyManagedConnections) GetLinks() MyManagedConnectionsLinks {
	if o == nil {
		var ret MyManagedConnectionsLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *MyManagedConnections) GetLinksOk() (*MyManagedConnectionsLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *MyManagedConnections) SetLinks(v MyManagedConnectionsLinks) {
	o.Links = v
}

func (o MyManagedConnections) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyManagedConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MyManagedConnections) UnmarshalJSON(data []byte) (err error) {
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

	varMyManagedConnections := _MyManagedConnections{}

	err = json.Unmarshal(data, &varMyManagedConnections)

	if err != nil {
		return err
	}

	*o = MyManagedConnections(varMyManagedConnections)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyManagedConnections struct {
	value *MyManagedConnections
	isSet bool
}

func (v NullableMyManagedConnections) Get() *MyManagedConnections {
	return v.value
}

func (v *NullableMyManagedConnections) Set(val *MyManagedConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableMyManagedConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableMyManagedConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyManagedConnections(val *MyManagedConnections) *NullableMyManagedConnections {
	return &NullableMyManagedConnections{value: val, isSet: true}
}

func (v NullableMyManagedConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyManagedConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
