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
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *Requested) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
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
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *Requested) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
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
	if o == nil || o.ResourceType == nil {
		var ret ResourceType3
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetResourceTypeOk() (*ResourceType3, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Requested) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
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
	if o == nil || o.AccessScopeId == nil {
		var ret string
		return ret
	}
	return *o.AccessScopeId
}

// GetAccessScopeIdOk returns a tuple with the AccessScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetAccessScopeIdOk() (*string, bool) {
	if o == nil || o.AccessScopeId == nil {
		return nil, false
	}
	return o.AccessScopeId, true
}

// HasAccessScopeId returns a boolean if a field has been set.
func (o *Requested) HasAccessScopeId() bool {
	if o != nil && o.AccessScopeId != nil {
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
	if o == nil || o.AccessScopeType == nil {
		var ret AccessScopeType
		return ret
	}
	return *o.AccessScopeType
}

// GetAccessScopeTypeOk returns a tuple with the AccessScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requested) GetAccessScopeTypeOk() (*AccessScopeType, bool) {
	if o == nil || o.AccessScopeType == nil {
		return nil, false
	}
	return o.AccessScopeType, true
}

// HasAccessScopeType returns a boolean if a field has been set.
func (o *Requested) HasAccessScopeType() bool {
	if o != nil && o.AccessScopeType != nil {
		return true
	}

	return false
}

// SetAccessScopeType gets a reference to the given AccessScopeType and assigns it to the AccessScopeType field.
func (o *Requested) SetAccessScopeType(v AccessScopeType) {
	o.AccessScopeType = &v
}

func (o Requested) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.AccessScopeId != nil {
		toSerialize["accessScopeId"] = o.AccessScopeId
	}
	if o.AccessScopeType != nil {
		toSerialize["accessScopeType"] = o.AccessScopeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Requested) UnmarshalJSON(bytes []byte) (err error) {
	varRequested := _Requested{}

	err = json.Unmarshal(bytes, &varRequested)
	if err == nil {
		*o = Requested(varRequested)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "entryId")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "accessScopeId")
		delete(additionalProperties, "accessScopeType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
