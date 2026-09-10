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

// checks if the ResourceCatalogVisibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceCatalogVisibility{}

// ResourceCatalogVisibility <x-lifecycle class=\"ea\"></x-lifecycle><br> Visibility settings for access to the resource catalog.  This setting controls the **Request Access** button visible in the End-User Dashboard. For the Unified requester experience, this setting also controls the **Resource Catalog** links and search options from the Okta Access Request app, Slack, or Microsoft Teams.
type ResourceCatalogVisibility struct {
	// Indicates whether users can access the resource catalog:   * If `false`, no users can access the resource catalog.   * If `true` and `onlyFor` isn't specified, all users in the org can access the resource catalog.
	Visible bool `json:"visible"`
	// Specific user targets for resource catalog visibility: * If this array is specified, only the specified targets can access the resource catalog. * If this array is null and `visible` is `true`, all users in the org can access the resource catalog.
	OnlyFor              []ResourceCatalogVisibilityTarget `json:"onlyFor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceCatalogVisibility ResourceCatalogVisibility

// NewResourceCatalogVisibility instantiates a new ResourceCatalogVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCatalogVisibility(visible bool) *ResourceCatalogVisibility {
	this := ResourceCatalogVisibility{}
	this.Visible = visible
	return &this
}

// NewResourceCatalogVisibilityWithDefaults instantiates a new ResourceCatalogVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCatalogVisibilityWithDefaults() *ResourceCatalogVisibility {
	this := ResourceCatalogVisibility{}
	return &this
}

// GetVisible returns the Visible field value
func (o *ResourceCatalogVisibility) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *ResourceCatalogVisibility) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *ResourceCatalogVisibility) SetVisible(v bool) {
	o.Visible = v
}

// GetOnlyFor returns the OnlyFor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceCatalogVisibility) GetOnlyFor() []ResourceCatalogVisibilityTarget {
	if o == nil {
		var ret []ResourceCatalogVisibilityTarget
		return ret
	}
	return o.OnlyFor
}

// GetOnlyForOk returns a tuple with the OnlyFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceCatalogVisibility) GetOnlyForOk() ([]ResourceCatalogVisibilityTarget, bool) {
	if o == nil || IsNil(o.OnlyFor) {
		return nil, false
	}
	return o.OnlyFor, true
}

// HasOnlyFor returns a boolean if a field has been set.
func (o *ResourceCatalogVisibility) HasOnlyFor() bool {
	if o != nil && !IsNil(o.OnlyFor) {
		return true
	}

	return false
}

// SetOnlyFor gets a reference to the given []ResourceCatalogVisibilityTarget and assigns it to the OnlyFor field.
func (o *ResourceCatalogVisibility) SetOnlyFor(v []ResourceCatalogVisibilityTarget) {
	o.OnlyFor = v
}

func (o ResourceCatalogVisibility) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceCatalogVisibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visible"] = o.Visible
	if o.OnlyFor != nil {
		toSerialize["onlyFor"] = o.OnlyFor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceCatalogVisibility) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"visible",
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

	varResourceCatalogVisibility := _ResourceCatalogVisibility{}

	err = json.Unmarshal(data, &varResourceCatalogVisibility)

	if err != nil {
		return err
	}

	*o = ResourceCatalogVisibility(varResourceCatalogVisibility)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "visible")
		delete(additionalProperties, "onlyFor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceCatalogVisibility struct {
	value *ResourceCatalogVisibility
	isSet bool
}

func (v NullableResourceCatalogVisibility) Get() *ResourceCatalogVisibility {
	return v.value
}

func (v *NullableResourceCatalogVisibility) Set(val *ResourceCatalogVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceCatalogVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceCatalogVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceCatalogVisibility(val *ResourceCatalogVisibility) *NullableResourceCatalogVisibility {
	return &NullableResourceCatalogVisibility{value: val, isSet: true}
}

func (v NullableResourceCatalogVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceCatalogVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
