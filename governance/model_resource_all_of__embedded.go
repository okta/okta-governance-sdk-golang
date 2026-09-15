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
)

// checks if the ResourceAllOfEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAllOfEmbedded{}

// ResourceAllOfEmbedded Embedded related resources
type ResourceAllOfEmbedded struct {
	Owners               []EmbeddedItem `json:"owners,omitempty"`
	Apps                 []EmbeddedItem `json:"apps,omitempty"`
	Labels               []EmbeddedItem `json:"labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAllOfEmbedded ResourceAllOfEmbedded

// NewResourceAllOfEmbedded instantiates a new ResourceAllOfEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAllOfEmbedded() *ResourceAllOfEmbedded {
	this := ResourceAllOfEmbedded{}
	return &this
}

// NewResourceAllOfEmbeddedWithDefaults instantiates a new ResourceAllOfEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAllOfEmbeddedWithDefaults() *ResourceAllOfEmbedded {
	this := ResourceAllOfEmbedded{}
	return &this
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *ResourceAllOfEmbedded) GetOwners() []EmbeddedItem {
	if o == nil || IsNil(o.Owners) {
		var ret []EmbeddedItem
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAllOfEmbedded) GetOwnersOk() ([]EmbeddedItem, bool) {
	if o == nil || IsNil(o.Owners) {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *ResourceAllOfEmbedded) HasOwners() bool {
	if o != nil && !IsNil(o.Owners) {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []EmbeddedItem and assigns it to the Owners field.
func (o *ResourceAllOfEmbedded) SetOwners(v []EmbeddedItem) {
	o.Owners = v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ResourceAllOfEmbedded) GetApps() []EmbeddedItem {
	if o == nil || IsNil(o.Apps) {
		var ret []EmbeddedItem
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAllOfEmbedded) GetAppsOk() ([]EmbeddedItem, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ResourceAllOfEmbedded) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []EmbeddedItem and assigns it to the Apps field.
func (o *ResourceAllOfEmbedded) SetApps(v []EmbeddedItem) {
	o.Apps = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ResourceAllOfEmbedded) GetLabels() []EmbeddedItem {
	if o == nil || IsNil(o.Labels) {
		var ret []EmbeddedItem
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAllOfEmbedded) GetLabelsOk() ([]EmbeddedItem, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ResourceAllOfEmbedded) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []EmbeddedItem and assigns it to the Labels field.
func (o *ResourceAllOfEmbedded) SetLabels(v []EmbeddedItem) {
	o.Labels = v
}

func (o ResourceAllOfEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAllOfEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owners) {
		toSerialize["owners"] = o.Owners
	}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAllOfEmbedded) UnmarshalJSON(data []byte) (err error) {
	varResourceAllOfEmbedded := _ResourceAllOfEmbedded{}

	err = json.Unmarshal(data, &varResourceAllOfEmbedded)

	if err != nil {
		return err
	}

	*o = ResourceAllOfEmbedded(varResourceAllOfEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "owners")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAllOfEmbedded struct {
	value *ResourceAllOfEmbedded
	isSet bool
}

func (v NullableResourceAllOfEmbedded) Get() *ResourceAllOfEmbedded {
	return v.value
}

func (v *NullableResourceAllOfEmbedded) Set(val *ResourceAllOfEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAllOfEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAllOfEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAllOfEmbedded(val *ResourceAllOfEmbedded) *NullableResourceAllOfEmbedded {
	return &NullableResourceAllOfEmbedded{value: val, isSet: true}
}

func (v NullableResourceAllOfEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAllOfEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
