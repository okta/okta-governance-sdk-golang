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

// checks if the Requested type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Requested{}

// Requested A representation of the resource in request
type Requested struct {
	// The ID of the resource catalog entry
	EntryId *string `json:"entryId,omitempty"`
	// the requested resource ID
	ResourceId   *string        `json:"resourceId,omitempty"`
	ResourceType *ResourceType3 `json:"resourceType,omitempty"`
	// ID of the access scope
	AccessScopeId        *string          `json:"accessScopeId,omitempty"`
	AccessScopeType      *AccessScopeType `json:"accessScopeType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Requested Requested

// NewRequested instantiates a new Requested object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequested() *Requested {
	this := Requested{}
	return &this
}

// NewRequestedWithDefaults instantiates a new Requested object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedWithDefaults() *Requested {
	this := Requested{}
	return &this
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *Requested) GetEntryId() string {
	if o == nil || IsNil(o.EntryId) {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetEntryIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntryId) {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *Requested) HasEntryId() bool {
	if o != nil && !IsNil(o.EntryId) {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *Requested) SetEntryId(v string) {
	o.EntryId = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *Requested) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *Requested) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *Requested) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Requested) GetResourceType() ResourceType3 {
	if o == nil || IsNil(o.ResourceType) {
		var ret ResourceType3
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetResourceTypeOk() (*ResourceType3, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Requested) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given ResourceType3 and assigns it to the ResourceType field.
func (o *Requested) SetResourceType(v ResourceType3) {
	o.ResourceType = &v
}

// GetAccessScopeId returns the AccessScopeId field value if set, zero value otherwise.
func (o *Requested) GetAccessScopeId() string {
	if o == nil || IsNil(o.AccessScopeId) {
		var ret string
		return ret
	}
	return *o.AccessScopeId
}

// GetAccessScopeIdOk returns a tuple with the AccessScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetAccessScopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessScopeId) {
		return nil, false
	}
	return o.AccessScopeId, true
}

// HasAccessScopeId returns a boolean if a field has been set.
func (o *Requested) HasAccessScopeId() bool {
	if o != nil && !IsNil(o.AccessScopeId) {
		return true
	}

	return false
}

// SetAccessScopeId gets a reference to the given string and assigns it to the AccessScopeId field.
func (o *Requested) SetAccessScopeId(v string) {
	o.AccessScopeId = &v
}

// GetAccessScopeType returns the AccessScopeType field value if set, zero value otherwise.
func (o *Requested) GetAccessScopeType() AccessScopeType {
	if o == nil || IsNil(o.AccessScopeType) {
		var ret AccessScopeType
		return ret
	}
	return *o.AccessScopeType
}

// GetAccessScopeTypeOk returns a tuple with the AccessScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetAccessScopeTypeOk() (*AccessScopeType, bool) {
	if o == nil || IsNil(o.AccessScopeType) {
		return nil, false
	}
	return o.AccessScopeType, true
}

// HasAccessScopeType returns a boolean if a field has been set.
func (o *Requested) HasAccessScopeType() bool {
	if o != nil && !IsNil(o.AccessScopeType) {
		return true
	}

	return false
}

// SetAccessScopeType gets a reference to the given AccessScopeType and assigns it to the AccessScopeType field.
func (o *Requested) SetAccessScopeType(v AccessScopeType) {
	o.AccessScopeType = &v
}

func (o Requested) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Requested) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntryId) {
		toSerialize["entryId"] = o.EntryId
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.AccessScopeId) {
		toSerialize["accessScopeId"] = o.AccessScopeId
	}
	if !IsNil(o.AccessScopeType) {
		toSerialize["accessScopeType"] = o.AccessScopeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Requested) UnmarshalJSON(data []byte) (err error) {
	varRequested := _Requested{}

	err = json.Unmarshal(data, &varRequested)

	if err != nil {
		return err
	}

	*o = Requested(varRequested)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entryId")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "accessScopeId")
		delete(additionalProperties, "accessScopeType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequested struct {
	value *Requested
	isSet bool
}

func (v NullableRequested) Get() *Requested {
	return v.value
}

func (v *NullableRequested) Set(val *Requested) {
	v.value = val
	v.isSet = true
}

func (v NullableRequested) IsSet() bool {
	return v.isSet
}

func (v *NullableRequested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequested(val *Requested) *NullableRequested {
	return &NullableRequested{value: val, isSet: true}
}

func (v NullableRequested) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
