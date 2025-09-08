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

// ResourceSettingsMutableExcludedResourcesInner Represents a resource that will be excluded from access certification
type ResourceSettingsMutableExcludedResourcesInner struct {
	// Okta specific resource ID
	ResourceId           *string       `json:"resourceId,omitempty"`
	ResourceType         *ResourceType `json:"resourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSettingsMutableExcludedResourcesInner ResourceSettingsMutableExcludedResourcesInner

// NewResourceSettingsMutableExcludedResourcesInner instantiates a new ResourceSettingsMutableExcludedResourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSettingsMutableExcludedResourcesInner() *ResourceSettingsMutableExcludedResourcesInner {
	this := ResourceSettingsMutableExcludedResourcesInner{}
	return &this
}

// NewResourceSettingsMutableExcludedResourcesInnerWithDefaults instantiates a new ResourceSettingsMutableExcludedResourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSettingsMutableExcludedResourcesInnerWithDefaults() *ResourceSettingsMutableExcludedResourcesInner {
	this := ResourceSettingsMutableExcludedResourcesInner{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *ResourceSettingsMutableExcludedResourcesInner) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutableExcludedResourcesInner) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *ResourceSettingsMutableExcludedResourcesInner) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *ResourceSettingsMutableExcludedResourcesInner) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourceSettingsMutableExcludedResourcesInner) GetResourceType() ResourceType {
	if o == nil || o.ResourceType == nil {
		var ret ResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutableExcludedResourcesInner) GetResourceTypeOk() (*ResourceType, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourceSettingsMutableExcludedResourcesInner) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given ResourceType and assigns it to the ResourceType field.
func (o *ResourceSettingsMutableExcludedResourcesInner) SetResourceType(v ResourceType) {
	o.ResourceType = &v
}

func (o ResourceSettingsMutableExcludedResourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSettingsMutableExcludedResourcesInner) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSettingsMutableExcludedResourcesInner := _ResourceSettingsMutableExcludedResourcesInner{}

	err = json.Unmarshal(bytes, &varResourceSettingsMutableExcludedResourcesInner)
	if err == nil {
		*o = ResourceSettingsMutableExcludedResourcesInner(varResourceSettingsMutableExcludedResourcesInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSettingsMutableExcludedResourcesInner struct {
	value *ResourceSettingsMutableExcludedResourcesInner
	isSet bool
}

func (v NullableResourceSettingsMutableExcludedResourcesInner) Get() *ResourceSettingsMutableExcludedResourcesInner {
	return v.value
}

func (v *NullableResourceSettingsMutableExcludedResourcesInner) Set(val *ResourceSettingsMutableExcludedResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSettingsMutableExcludedResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSettingsMutableExcludedResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSettingsMutableExcludedResourcesInner(val *ResourceSettingsMutableExcludedResourcesInner) *NullableResourceSettingsMutableExcludedResourcesInner {
	return &NullableResourceSettingsMutableExcludedResourcesInner{value: val, isSet: true}
}

func (v NullableResourceSettingsMutableExcludedResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSettingsMutableExcludedResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
