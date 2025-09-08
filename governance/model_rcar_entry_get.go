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

// RcarEntryGet struct for RcarEntryGet
type RcarEntryGet struct {
	// Unique identifier for the entry
	Id string `json:"id"`
	// Name of the entry
	Name string `json:"name"`
	// Description of the entry
	Description *string `json:"description,omitempty"`
	// Label of the entry. Currently either `Application` or `Resource Collection`
	Label *string `json:"label,omitempty"`
	// Parent of the entry
	Parent *string `json:"parent,omitempty"`
	// Is the resource requestable
	Requestable          bool             `json:"requestable"`
	Counts               *RcarEntryCounts `json:"counts,omitempty"`
	Links                RcarEntryLinks   `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _RcarEntryGet RcarEntryGet

// NewRcarEntryGet instantiates a new RcarEntryGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcarEntryGet(id string, name string, requestable bool, links RcarEntryLinks) *RcarEntryGet {
	this := RcarEntryGet{}
	this.Id = id
	this.Name = name
	this.Requestable = requestable
	this.Links = links
	return &this
}

// NewRcarEntryGetWithDefaults instantiates a new RcarEntryGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcarEntryGetWithDefaults() *RcarEntryGet {
	this := RcarEntryGet{}
	return &this
}

// GetId returns the Id field value
func (o *RcarEntryGet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RcarEntryGet) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RcarEntryGet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RcarEntryGet) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RcarEntryGet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RcarEntryGet) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RcarEntryGet) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RcarEntryGet) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RcarEntryGet) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RcarEntryGet) SetLabel(v string) {
	o.Label = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *RcarEntryGet) GetParent() string {
	if o == nil || o.Parent == nil {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetParentOk() (*string, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *RcarEntryGet) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *RcarEntryGet) SetParent(v string) {
	o.Parent = &v
}

// GetRequestable returns the Requestable field value
func (o *RcarEntryGet) GetRequestable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Requestable
}

// GetRequestableOk returns a tuple with the Requestable field value
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetRequestableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requestable, true
}

// SetRequestable sets field value
func (o *RcarEntryGet) SetRequestable(v bool) {
	o.Requestable = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *RcarEntryGet) GetCounts() RcarEntryCounts {
	if o == nil || o.Counts == nil {
		var ret RcarEntryCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetCountsOk() (*RcarEntryCounts, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *RcarEntryGet) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given RcarEntryCounts and assigns it to the Counts field.
func (o *RcarEntryGet) SetCounts(v RcarEntryCounts) {
	o.Counts = &v
}

// GetLinks returns the Links field value
func (o *RcarEntryGet) GetLinks() RcarEntryLinks {
	if o == nil {
		var ret RcarEntryLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RcarEntryGet) GetLinksOk() (*RcarEntryLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RcarEntryGet) SetLinks(v RcarEntryLinks) {
	o.Links = v
}

func (o RcarEntryGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if true {
		toSerialize["requestable"] = o.Requestable
	}
	if o.Counts != nil {
		toSerialize["counts"] = o.Counts
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RcarEntryGet) UnmarshalJSON(bytes []byte) (err error) {
	varRcarEntryGet := _RcarEntryGet{}

	err = json.Unmarshal(bytes, &varRcarEntryGet)
	if err == nil {
		*o = RcarEntryGet(varRcarEntryGet)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "label")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "requestable")
		delete(additionalProperties, "counts")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRcarEntryGet struct {
	value *RcarEntryGet
	isSet bool
}

func (v NullableRcarEntryGet) Get() *RcarEntryGet {
	return v.value
}

func (v *NullableRcarEntryGet) Set(val *RcarEntryGet) {
	v.value = val
	v.isSet = true
}

func (v NullableRcarEntryGet) IsSet() bool {
	return v.isSet
}

func (v *NullableRcarEntryGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcarEntryGet(val *RcarEntryGet) *NullableRcarEntryGet {
	return &NullableRcarEntryGet{value: val, isSet: true}
}

func (v NullableRcarEntryGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcarEntryGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
