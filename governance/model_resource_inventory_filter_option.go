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

// checks if the ResourceInventoryFilterOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceInventoryFilterOption{}

// ResourceInventoryFilterOption struct for ResourceInventoryFilterOption
type ResourceInventoryFilterOption struct {
	// Unique identifier for the filter option. For owners, this is the Okta user ID.
	Id string `json:"id"`
	// Display name of the filter option
	Name string `json:"name"`
	// URL to a logo image. Only present for app filter options.
	Logo *string `json:"logo,omitempty"`
	// Additional metadata properties for the filter option. Only present for certain filter types.  For the `labels` filter type, metadata contains: * `values`: array of label values, each with `id`, `name`, and `backgroundColor`   * `backgroundColor`: [`red`, `orange`, `yellow`, `green`, `blue`, `purple`, `teal`, `beige`, `gray`]
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceInventoryFilterOption ResourceInventoryFilterOption

// NewResourceInventoryFilterOption instantiates a new ResourceInventoryFilterOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceInventoryFilterOption(id string, name string) *ResourceInventoryFilterOption {
	this := ResourceInventoryFilterOption{}
	this.Id = id
	this.Name = name
	return &this
}

// NewResourceInventoryFilterOptionWithDefaults instantiates a new ResourceInventoryFilterOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceInventoryFilterOptionWithDefaults() *ResourceInventoryFilterOption {
	this := ResourceInventoryFilterOption{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceInventoryFilterOption) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryFilterOption) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceInventoryFilterOption) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResourceInventoryFilterOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceInventoryFilterOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceInventoryFilterOption) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ResourceInventoryFilterOption) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInventoryFilterOption) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ResourceInventoryFilterOption) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ResourceInventoryFilterOption) SetLogo(v string) {
	o.Logo = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ResourceInventoryFilterOption) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceInventoryFilterOption) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ResourceInventoryFilterOption) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ResourceInventoryFilterOption) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ResourceInventoryFilterOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceInventoryFilterOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceInventoryFilterOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varResourceInventoryFilterOption := _ResourceInventoryFilterOption{}

	err = json.Unmarshal(data, &varResourceInventoryFilterOption)

	if err != nil {
		return err
	}

	*o = ResourceInventoryFilterOption(varResourceInventoryFilterOption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceInventoryFilterOption struct {
	value *ResourceInventoryFilterOption
	isSet bool
}

func (v NullableResourceInventoryFilterOption) Get() *ResourceInventoryFilterOption {
	return v.value
}

func (v *NullableResourceInventoryFilterOption) Set(val *ResourceInventoryFilterOption) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceInventoryFilterOption) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceInventoryFilterOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceInventoryFilterOption(val *ResourceInventoryFilterOption) *NullableResourceInventoryFilterOption {
	return &NullableResourceInventoryFilterOption{value: val, isSet: true}
}

func (v NullableResourceInventoryFilterOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceInventoryFilterOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
