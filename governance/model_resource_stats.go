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

// checks if the ResourceStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceStats{}

// ResourceStats Statistical information about the resource. Not all properties are applicable to every resource type.
type ResourceStats struct {
	// Number of associated apps
	Apps *int32 `json:"apps,omitempty"`
	// Number of associated groups
	Groups *int32 `json:"groups,omitempty"`
	// Number of associated users
	Users *int32 `json:"users,omitempty"`
	// Number of owners
	Owners *int32 `json:"owners,omitempty"`
	// Number of entitlements
	Entitlements *int32 `json:"entitlements,omitempty"`
	// Number of associated labels
	Labels               *int32 `json:"labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceStats ResourceStats

// NewResourceStats instantiates a new ResourceStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceStats() *ResourceStats {
	this := ResourceStats{}
	return &this
}

// NewResourceStatsWithDefaults instantiates a new ResourceStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceStatsWithDefaults() *ResourceStats {
	this := ResourceStats{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ResourceStats) GetApps() int32 {
	if o == nil || IsNil(o.Apps) {
		var ret int32
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetAppsOk() (*int32, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ResourceStats) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given int32 and assigns it to the Apps field.
func (o *ResourceStats) SetApps(v int32) {
	o.Apps = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ResourceStats) GetGroups() int32 {
	if o == nil || IsNil(o.Groups) {
		var ret int32
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ResourceStats) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given int32 and assigns it to the Groups field.
func (o *ResourceStats) SetGroups(v int32) {
	o.Groups = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ResourceStats) GetUsers() int32 {
	if o == nil || IsNil(o.Users) {
		var ret int32
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ResourceStats) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given int32 and assigns it to the Users field.
func (o *ResourceStats) SetUsers(v int32) {
	o.Users = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *ResourceStats) GetOwners() int32 {
	if o == nil || IsNil(o.Owners) {
		var ret int32
		return ret
	}
	return *o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetOwnersOk() (*int32, bool) {
	if o == nil || IsNil(o.Owners) {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *ResourceStats) HasOwners() bool {
	if o != nil && !IsNil(o.Owners) {
		return true
	}

	return false
}

// SetOwners gets a reference to the given int32 and assigns it to the Owners field.
func (o *ResourceStats) SetOwners(v int32) {
	o.Owners = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *ResourceStats) GetEntitlements() int32 {
	if o == nil || IsNil(o.Entitlements) {
		var ret int32
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetEntitlementsOk() (*int32, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *ResourceStats) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given int32 and assigns it to the Entitlements field.
func (o *ResourceStats) SetEntitlements(v int32) {
	o.Entitlements = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ResourceStats) GetLabels() int32 {
	if o == nil || IsNil(o.Labels) {
		var ret int32
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceStats) GetLabelsOk() (*int32, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ResourceStats) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given int32 and assigns it to the Labels field.
func (o *ResourceStats) SetLabels(v int32) {
	o.Labels = &v
}

func (o ResourceStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Owners) {
		toSerialize["owners"] = o.Owners
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceStats) UnmarshalJSON(data []byte) (err error) {
	varResourceStats := _ResourceStats{}

	err = json.Unmarshal(data, &varResourceStats)

	if err != nil {
		return err
	}

	*o = ResourceStats(varResourceStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apps")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		delete(additionalProperties, "owners")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceStats struct {
	value *ResourceStats
	isSet bool
}

func (v NullableResourceStats) Get() *ResourceStats {
	return v.value
}

func (v *NullableResourceStats) Set(val *ResourceStats) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceStats) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceStats(val *ResourceStats) *NullableResourceStats {
	return &NullableResourceStats{value: val, isSet: true}
}

func (v NullableResourceStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
