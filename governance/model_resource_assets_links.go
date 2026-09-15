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

// checks if the ResourceAssetsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssetsLinks{}

// ResourceAssetsLinks Links available in resource asset response
type ResourceAssetsLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	Parent               *Link `json:"parent,omitempty"`
	Children             *Link `json:"children,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssetsLinks ResourceAssetsLinks

// NewResourceAssetsLinks instantiates a new ResourceAssetsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssetsLinks(self Link) *ResourceAssetsLinks {
	this := ResourceAssetsLinks{}
	this.Self = self
	return &this
}

// NewResourceAssetsLinksWithDefaults instantiates a new ResourceAssetsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssetsLinksWithDefaults() *ResourceAssetsLinks {
	this := ResourceAssetsLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ResourceAssetsLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ResourceAssetsLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ResourceAssetsLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResourceAssetsLinks) GetNext() Link {
	if o == nil || IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetsLinks) GetNextOk() (*Link, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResourceAssetsLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *ResourceAssetsLinks) SetNext(v Link) {
	o.Next = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ResourceAssetsLinks) GetParent() Link {
	if o == nil || IsNil(o.Parent) {
		var ret Link
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetsLinks) GetParentOk() (*Link, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ResourceAssetsLinks) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given Link and assigns it to the Parent field.
func (o *ResourceAssetsLinks) SetParent(v Link) {
	o.Parent = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ResourceAssetsLinks) GetChildren() Link {
	if o == nil || IsNil(o.Children) {
		var ret Link
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssetsLinks) GetChildrenOk() (*Link, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ResourceAssetsLinks) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given Link and assigns it to the Children field.
func (o *ResourceAssetsLinks) SetChildren(v Link) {
	o.Children = &v
}

func (o ResourceAssetsLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssetsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssetsLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varResourceAssetsLinks := _ResourceAssetsLinks{}

	err = json.Unmarshal(data, &varResourceAssetsLinks)

	if err != nil {
		return err
	}

	*o = ResourceAssetsLinks(varResourceAssetsLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "children")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssetsLinks struct {
	value *ResourceAssetsLinks
	isSet bool
}

func (v NullableResourceAssetsLinks) Get() *ResourceAssetsLinks {
	return v.value
}

func (v *NullableResourceAssetsLinks) Set(val *ResourceAssetsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssetsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssetsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssetsLinks(val *ResourceAssetsLinks) *NullableResourceAssetsLinks {
	return &NullableResourceAssetsLinks{value: val, isSet: true}
}

func (v NullableResourceAssetsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssetsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
