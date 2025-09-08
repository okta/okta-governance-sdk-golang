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

// ResourceOwnerPrincipal Representation of a principal, which can be a user or a group, that is an owner of a resource.
type ResourceOwnerPrincipal struct {
	// Id of the principal, which is a unique identifier for the principal in Okta.
	Id string `json:"id"`
	// The principal type value from the orn. Examples:- groups, users
	Type string `json:"type"`
	// The Okta user or group `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an user id, or a group id. See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Orn                  string                   `json:"orn"`
	Profile              ExternalPrincipalProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnerPrincipal ResourceOwnerPrincipal

// NewResourceOwnerPrincipal instantiates a new ResourceOwnerPrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnerPrincipal(id string, type_ string, orn string, profile ExternalPrincipalProfile) *ResourceOwnerPrincipal {
	this := ResourceOwnerPrincipal{}
	this.Id = id
	this.Type = type_
	this.Orn = orn
	this.Profile = profile
	return &this
}

// NewResourceOwnerPrincipalWithDefaults instantiates a new ResourceOwnerPrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnerPrincipalWithDefaults() *ResourceOwnerPrincipal {
	this := ResourceOwnerPrincipal{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceOwnerPrincipal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerPrincipal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceOwnerPrincipal) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ResourceOwnerPrincipal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerPrincipal) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceOwnerPrincipal) SetType(v string) {
	o.Type = v
}

// GetOrn returns the Orn field value
func (o *ResourceOwnerPrincipal) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerPrincipal) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *ResourceOwnerPrincipal) SetOrn(v string) {
	o.Orn = v
}

// GetProfile returns the Profile field value
func (o *ResourceOwnerPrincipal) GetProfile() ExternalPrincipalProfile {
	if o == nil {
		var ret ExternalPrincipalProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnerPrincipal) GetProfileOk() (*ExternalPrincipalProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ResourceOwnerPrincipal) SetProfile(v ExternalPrincipalProfile) {
	o.Profile = v
}

func (o ResourceOwnerPrincipal) MarshalJSON() ([]byte, error) {
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

func (o *ResourceOwnerPrincipal) UnmarshalJSON(bytes []byte) (err error) {
	varResourceOwnerPrincipal := _ResourceOwnerPrincipal{}

	err = json.Unmarshal(bytes, &varResourceOwnerPrincipal)
	if err == nil {
		*o = ResourceOwnerPrincipal(varResourceOwnerPrincipal)
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

type NullableResourceOwnerPrincipal struct {
	value *ResourceOwnerPrincipal
	isSet bool
}

func (v NullableResourceOwnerPrincipal) Get() *ResourceOwnerPrincipal {
	return v.value
}

func (v *NullableResourceOwnerPrincipal) Set(val *ResourceOwnerPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnerPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnerPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnerPrincipal(val *ResourceOwnerPrincipal) *NullableResourceOwnerPrincipal {
	return &NullableResourceOwnerPrincipal{value: val, isSet: true}
}

func (v NullableResourceOwnerPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnerPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
