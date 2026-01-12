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

// checks if the OktaResourceReadable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaResourceReadable{}

// OktaResourceReadable Identifies a unique Okta resource
type OktaResourceReadable struct {
	// Human readable name of the resource
	ResourceName *string `json:"resourceName,omitempty"`
	// The Okta `app.id`, or `group.id` of the resource that can be requested with this Request Type.  * See [List applications](https://developer.okta.com/docs/reference/api/apps/#list-applications) to retrieve app IDs. * See [List groups](https://developer.okta.com/docs/reference/api/groups/#list-groups) to retrieve group IDs.
	ResourceId *string `json:"resourceId,omitempty"`
	// The type of resource
	ResourceType         *string `json:"resourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaResourceReadable OktaResourceReadable

// NewOktaResourceReadable instantiates a new OktaResourceReadable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaResourceReadable() *OktaResourceReadable {
	this := OktaResourceReadable{}
	return &this
}

// NewOktaResourceReadableWithDefaults instantiates a new OktaResourceReadable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaResourceReadableWithDefaults() *OktaResourceReadable {
	this := OktaResourceReadable{}
	return &this
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *OktaResourceReadable) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaResourceReadable) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *OktaResourceReadable) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *OktaResourceReadable) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *OktaResourceReadable) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaResourceReadable) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *OktaResourceReadable) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *OktaResourceReadable) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *OktaResourceReadable) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaResourceReadable) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *OktaResourceReadable) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *OktaResourceReadable) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o OktaResourceReadable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaResourceReadable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaResourceReadable) UnmarshalJSON(data []byte) (err error) {
	varOktaResourceReadable := _OktaResourceReadable{}

	err = json.Unmarshal(data, &varOktaResourceReadable)

	if err != nil {
		return err
	}

	*o = OktaResourceReadable(varOktaResourceReadable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceName")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaResourceReadable struct {
	value *OktaResourceReadable
	isSet bool
}

func (v NullableOktaResourceReadable) Get() *OktaResourceReadable {
	return v.value
}

func (v *NullableOktaResourceReadable) Set(val *OktaResourceReadable) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaResourceReadable) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaResourceReadable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaResourceReadable(val *OktaResourceReadable) *NullableOktaResourceReadable {
	return &NullableOktaResourceReadable{value: val, isSet: true}
}

func (v NullableOktaResourceReadable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaResourceReadable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
