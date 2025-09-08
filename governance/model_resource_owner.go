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

// ResourceOwner struct for ResourceOwner
type ResourceOwner struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn    string                   `json:"parentResourceOrn"`
	Principals           []ResourceOwnerPrincipal `json:"principals,omitempty"`
	Resource             ResourceOwnerResource    `json:"resource"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwner ResourceOwner

// NewResourceOwner instantiates a new ResourceOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwner(parentResourceOrn string, resource ResourceOwnerResource) *ResourceOwner {
	this := ResourceOwner{}
	this.ParentResourceOrn = parentResourceOrn
	this.Resource = resource
	return &this
}

// NewResourceOwnerWithDefaults instantiates a new ResourceOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnerWithDefaults() *ResourceOwner {
	this := ResourceOwner{}
	return &this
}

// GetParentResourceOrn returns the ParentResourceOrn field value
func (o *ResourceOwner) GetParentResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value
// and a boolean to check if the value has been set.
func (o *ResourceOwner) GetParentResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentResourceOrn, true
}

// SetParentResourceOrn sets field value
func (o *ResourceOwner) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = v
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *ResourceOwner) GetPrincipals() []ResourceOwnerPrincipal {
	if o == nil || o.Principals == nil {
		var ret []ResourceOwnerPrincipal
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwner) GetPrincipalsOk() ([]ResourceOwnerPrincipal, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *ResourceOwner) HasPrincipals() bool {
	if o != nil && o.Principals != nil {
		return true
	}

	return false
}

// SetPrincipals gets a reference to the given []ResourceOwnerPrincipal and assigns it to the Principals field.
func (o *ResourceOwner) SetPrincipals(v []ResourceOwnerPrincipal) {
	o.Principals = v
}

// GetResource returns the Resource field value
func (o *ResourceOwner) GetResource() ResourceOwnerResource {
	if o == nil {
		var ret ResourceOwnerResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ResourceOwner) GetResourceOk() (*ResourceOwnerResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ResourceOwner) SetResource(v ResourceOwnerResource) {
	o.Resource = v
}

func (o ResourceOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if true {
		toSerialize["resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceOwner) UnmarshalJSON(bytes []byte) (err error) {
	varResourceOwner := _ResourceOwner{}

	err = json.Unmarshal(bytes, &varResourceOwner)
	if err == nil {
		*o = ResourceOwner(varResourceOwner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "principals")
		delete(additionalProperties, "resource")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceOwner struct {
	value *ResourceOwner
	isSet bool
}

func (v NullableResourceOwner) Get() *ResourceOwner {
	return v.value
}

func (v *NullableResourceOwner) Set(val *ResourceOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwner(val *ResourceOwner) *NullableResourceOwner {
	return &NullableResourceOwner{value: val, isSet: true}
}

func (v NullableResourceOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
