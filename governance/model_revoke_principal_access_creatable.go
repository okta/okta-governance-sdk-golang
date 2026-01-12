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

// checks if the RevokePrincipalAccessCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokePrincipalAccessCreatable{}

// RevokePrincipalAccessCreatable struct for RevokePrincipalAccessCreatable
type RevokePrincipalAccessCreatable struct {
	// The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.
	PrincipalOrn string      `json:"principalOrn"`
	Actor        *GrantActor `json:"actor,omitempty"`
	// List of resource [ORNs](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) to revoke access:   * Entitlement value and entitlement bundle resources can be combined in a single request (with a maximum of five resources in a request).   * App resources must be revoked separately (a request can only contain one app ORN).
	RevokeOrns           []string `json:"revokeOrns"`
	AdditionalProperties map[string]interface{}
}

type _RevokePrincipalAccessCreatable RevokePrincipalAccessCreatable

// NewRevokePrincipalAccessCreatable instantiates a new RevokePrincipalAccessCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokePrincipalAccessCreatable(principalOrn string, revokeOrns []string) *RevokePrincipalAccessCreatable {
	this := RevokePrincipalAccessCreatable{}
	this.PrincipalOrn = principalOrn
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	this.RevokeOrns = revokeOrns
	return &this
}

// NewRevokePrincipalAccessCreatableWithDefaults instantiates a new RevokePrincipalAccessCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrincipalAccessCreatableWithDefaults() *RevokePrincipalAccessCreatable {
	this := RevokePrincipalAccessCreatable{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *RevokePrincipalAccessCreatable) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessCreatable) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *RevokePrincipalAccessCreatable) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *RevokePrincipalAccessCreatable) GetActor() GrantActor {
	if o == nil || IsNil(o.Actor) {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessCreatable) GetActorOk() (*GrantActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *RevokePrincipalAccessCreatable) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *RevokePrincipalAccessCreatable) SetActor(v GrantActor) {
	o.Actor = &v
}

// GetRevokeOrns returns the RevokeOrns field value
func (o *RevokePrincipalAccessCreatable) GetRevokeOrns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RevokeOrns
}

// GetRevokeOrnsOk returns a tuple with the RevokeOrns field value
// and a boolean to check if the value has been set.
func (o *RevokePrincipalAccessCreatable) GetRevokeOrnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevokeOrns, true
}

// SetRevokeOrns sets field value
func (o *RevokePrincipalAccessCreatable) SetRevokeOrns(v []string) {
	o.RevokeOrns = v
}

func (o RevokePrincipalAccessCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokePrincipalAccessCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalOrn"] = o.PrincipalOrn
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	toSerialize["revokeOrns"] = o.RevokeOrns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokePrincipalAccessCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalOrn",
		"revokeOrns",
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

	varRevokePrincipalAccessCreatable := _RevokePrincipalAccessCreatable{}

	err = json.Unmarshal(data, &varRevokePrincipalAccessCreatable)

	if err != nil {
		return err
	}

	*o = RevokePrincipalAccessCreatable(varRevokePrincipalAccessCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "revokeOrns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokePrincipalAccessCreatable struct {
	value *RevokePrincipalAccessCreatable
	isSet bool
}

func (v NullableRevokePrincipalAccessCreatable) Get() *RevokePrincipalAccessCreatable {
	return v.value
}

func (v *NullableRevokePrincipalAccessCreatable) Set(val *RevokePrincipalAccessCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokePrincipalAccessCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokePrincipalAccessCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokePrincipalAccessCreatable(val *RevokePrincipalAccessCreatable) *NullableRevokePrincipalAccessCreatable {
	return &NullableRevokePrincipalAccessCreatable{value: val, isSet: true}
}

func (v NullableRevokePrincipalAccessCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokePrincipalAccessCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
