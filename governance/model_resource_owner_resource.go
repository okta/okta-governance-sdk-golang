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

// ResourceOwnerResource Representation of a resource that can be owned by a principal, such as an application or an entitlement bundle.
type ResourceOwnerResource struct {
	// Id of the resource, which is a unique identifier for the resource in Okta.
	Id string `json:"id"`
	// The resource type value from the orn. Examples:- apps, entitlement-bundles
	Type string `json:"type"`
	// The `id` of the resource in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an app, or a bundle. See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Orn                  string                  `json:"orn"`
	Profile              ExternalResourceProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnerResource ResourceOwnerResource

// NewResourceOwnerResource instantiates a new ResourceOwnerResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnerResource(id string, type_ string, orn string, profile ExternalResourceProfile) *ResourceOwnerResource {
	this := ResourceOwnerResource{}
	this.Id = id
	this.Type = type_
	this.Orn = orn
	this.Profile = profile
	return &this
}

// NewResourceOwnerResourceWithDefaults instantiates a new ResourceOwnerResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnerResourceWithDefaults() *ResourceOwnerResource {
	this := ResourceOwnerResource{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceOwnerResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceOwnerResource) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ResourceOwnerResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceOwnerResource) SetType(v string) {
	o.Type = v
}

// GetOrn returns the Orn field value
func (o *ResourceOwnerResource) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerResource) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *ResourceOwnerResource) SetOrn(v string) {
	o.Orn = v
}

// GetProfile returns the Profile field value
func (o *ResourceOwnerResource) GetProfile() ExternalResourceProfile {
	if o == nil {
		var ret ExternalResourceProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerResource) GetProfileOk() (*ExternalResourceProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ResourceOwnerResource) SetProfile(v ExternalResourceProfile) {
	o.Profile = v
}

func (o ResourceOwnerResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["orn"] = o.Orn
	}
	if true {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceOwnerResource) UnmarshalJSON(bytes []byte) (err error) {
	varResourceOwnerResource := _ResourceOwnerResource{}

	err = json.Unmarshal(bytes, &varResourceOwnerResource)
	if err == nil {
		*o = ResourceOwnerResource(varResourceOwnerResource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceOwnerResource struct {
	value *ResourceOwnerResource
	isSet bool
}

func (v NullableResourceOwnerResource) Get() *ResourceOwnerResource {
	return v.value
}

func (v *NullableResourceOwnerResource) Set(val *ResourceOwnerResource) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnerResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnerResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnerResource(val *ResourceOwnerResource) *NullableResourceOwnerResource {
	return &NullableResourceOwnerResource{value: val, isSet: true}
}

func (v NullableResourceOwnerResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnerResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
