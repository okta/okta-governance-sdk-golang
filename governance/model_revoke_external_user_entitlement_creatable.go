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

// checks if the RevokeExternalUserEntitlementCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokeExternalUserEntitlementCreatable{}

// RevokeExternalUserEntitlementCreatable struct for RevokeExternalUserEntitlementCreatable
type RevokeExternalUserEntitlementCreatable struct {
	// The external/source-system identifier of the unmanaged user whose entitlements are being revoked. Provided by the source app at import time.
	Principal string      `json:"principal"`
	Actor     *GrantActor `json:"actor,omitempty"`
	// List of entitlement-value [ORNs](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) to revoke for this external user
	RevokeOrns           []string `json:"revokeOrns"`
	AdditionalProperties map[string]interface{}
}

type _RevokeExternalUserEntitlementCreatable RevokeExternalUserEntitlementCreatable

// NewRevokeExternalUserEntitlementCreatable instantiates a new RevokeExternalUserEntitlementCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeExternalUserEntitlementCreatable(principal string, revokeOrns []string) *RevokeExternalUserEntitlementCreatable {
	this := RevokeExternalUserEntitlementCreatable{}
	this.Principal = principal
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	this.RevokeOrns = revokeOrns
	return &this
}

// NewRevokeExternalUserEntitlementCreatableWithDefaults instantiates a new RevokeExternalUserEntitlementCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeExternalUserEntitlementCreatableWithDefaults() *RevokeExternalUserEntitlementCreatable {
	this := RevokeExternalUserEntitlementCreatable{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetPrincipal returns the Principal field value
func (o *RevokeExternalUserEntitlementCreatable) GetPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *RevokeExternalUserEntitlementCreatable) GetPrincipalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *RevokeExternalUserEntitlementCreatable) SetPrincipal(v string) {
	o.Principal = v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *RevokeExternalUserEntitlementCreatable) GetActor() GrantActor {
	if o == nil || IsNil(o.Actor) {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeExternalUserEntitlementCreatable) GetActorOk() (*GrantActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *RevokeExternalUserEntitlementCreatable) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *RevokeExternalUserEntitlementCreatable) SetActor(v GrantActor) {
	o.Actor = &v
}

// GetRevokeOrns returns the RevokeOrns field value
func (o *RevokeExternalUserEntitlementCreatable) GetRevokeOrns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RevokeOrns
}

// GetRevokeOrnsOk returns a tuple with the RevokeOrns field value
// and a boolean to check if the value has been set.
func (o *RevokeExternalUserEntitlementCreatable) GetRevokeOrnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevokeOrns, true
}

// SetRevokeOrns sets field value
func (o *RevokeExternalUserEntitlementCreatable) SetRevokeOrns(v []string) {
	o.RevokeOrns = v
}

func (o RevokeExternalUserEntitlementCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokeExternalUserEntitlementCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principal"] = o.Principal
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	toSerialize["revokeOrns"] = o.RevokeOrns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RevokeExternalUserEntitlementCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principal",
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

	varRevokeExternalUserEntitlementCreatable := _RevokeExternalUserEntitlementCreatable{}

	err = json.Unmarshal(data, &varRevokeExternalUserEntitlementCreatable)

	if err != nil {
		return err
	}

	*o = RevokeExternalUserEntitlementCreatable(varRevokeExternalUserEntitlementCreatable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principal")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "revokeOrns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRevokeExternalUserEntitlementCreatable struct {
	value *RevokeExternalUserEntitlementCreatable
	isSet bool
}

func (v NullableRevokeExternalUserEntitlementCreatable) Get() *RevokeExternalUserEntitlementCreatable {
	return v.value
}

func (v *NullableRevokeExternalUserEntitlementCreatable) Set(val *RevokeExternalUserEntitlementCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeExternalUserEntitlementCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeExternalUserEntitlementCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeExternalUserEntitlementCreatable(val *RevokeExternalUserEntitlementCreatable) *NullableRevokeExternalUserEntitlementCreatable {
	return &NullableRevokeExternalUserEntitlementCreatable{value: val, isSet: true}
}

func (v NullableRevokeExternalUserEntitlementCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeExternalUserEntitlementCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
