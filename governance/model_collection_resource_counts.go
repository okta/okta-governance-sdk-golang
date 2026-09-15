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

// checks if the CollectionResourceCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceCounts{}

// CollectionResourceCounts Counts of related items for this resource. Only non-zero counts are included.
type CollectionResourceCounts struct {
	// The total number of entitlements included for this app resource in this collection
	Entitlements *int32 `json:"entitlements,omitempty"`
	// The total number of push groups associated with this app resource in this collection
	PushGroups *int32 `json:"pushGroups,omitempty"`
	// The total number of governance labels associated with this resource. Only returned when `include=labels` is specified. Not supported in LIST calls.
	Labels *int32 `json:"labels,omitempty"`
	// The total number of related apps for this resource. Only returned when `include=relatedApps` is specified. Not supported in LIST calls.
	RelatedApps          *int32 `json:"relatedApps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceCounts CollectionResourceCounts

// NewCollectionResourceCounts instantiates a new CollectionResourceCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceCounts() *CollectionResourceCounts {
	this := CollectionResourceCounts{}
	return &this
}

// NewCollectionResourceCountsWithDefaults instantiates a new CollectionResourceCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceCountsWithDefaults() *CollectionResourceCounts {
	this := CollectionResourceCounts{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceCounts) GetEntitlements() int32 {
	if o == nil || IsNil(o.Entitlements) {
		var ret int32
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceCounts) GetEntitlementsOk() (*int32, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceCounts) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given int32 and assigns it to the Entitlements field.
func (o *CollectionResourceCounts) SetEntitlements(v int32) {
	o.Entitlements = &v
}

// GetPushGroups returns the PushGroups field value if set, zero value otherwise.
func (o *CollectionResourceCounts) GetPushGroups() int32 {
	if o == nil || IsNil(o.PushGroups) {
		var ret int32
		return ret
	}
	return *o.PushGroups
}

// GetPushGroupsOk returns a tuple with the PushGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceCounts) GetPushGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.PushGroups) {
		return nil, false
	}
	return o.PushGroups, true
}

// HasPushGroups returns a boolean if a field has been set.
func (o *CollectionResourceCounts) HasPushGroups() bool {
	if o != nil && !IsNil(o.PushGroups) {
		return true
	}

	return false
}

// SetPushGroups gets a reference to the given int32 and assigns it to the PushGroups field.
func (o *CollectionResourceCounts) SetPushGroups(v int32) {
	o.PushGroups = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CollectionResourceCounts) GetLabels() int32 {
	if o == nil || IsNil(o.Labels) {
		var ret int32
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceCounts) GetLabelsOk() (*int32, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CollectionResourceCounts) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given int32 and assigns it to the Labels field.
func (o *CollectionResourceCounts) SetLabels(v int32) {
	o.Labels = &v
}

// GetRelatedApps returns the RelatedApps field value if set, zero value otherwise.
func (o *CollectionResourceCounts) GetRelatedApps() int32 {
	if o == nil || IsNil(o.RelatedApps) {
		var ret int32
		return ret
	}
	return *o.RelatedApps
}

// GetRelatedAppsOk returns a tuple with the RelatedApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceCounts) GetRelatedAppsOk() (*int32, bool) {
	if o == nil || IsNil(o.RelatedApps) {
		return nil, false
	}
	return o.RelatedApps, true
}

// HasRelatedApps returns a boolean if a field has been set.
func (o *CollectionResourceCounts) HasRelatedApps() bool {
	if o != nil && !IsNil(o.RelatedApps) {
		return true
	}

	return false
}

// SetRelatedApps gets a reference to the given int32 and assigns it to the RelatedApps field.
func (o *CollectionResourceCounts) SetRelatedApps(v int32) {
	o.RelatedApps = &v
}

func (o CollectionResourceCounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.PushGroups) {
		toSerialize["pushGroups"] = o.PushGroups
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.RelatedApps) {
		toSerialize["relatedApps"] = o.RelatedApps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceCounts) UnmarshalJSON(data []byte) (err error) {
	varCollectionResourceCounts := _CollectionResourceCounts{}

	err = json.Unmarshal(data, &varCollectionResourceCounts)

	if err != nil {
		return err
	}

	*o = CollectionResourceCounts(varCollectionResourceCounts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "pushGroups")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "relatedApps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceCounts struct {
	value *CollectionResourceCounts
	isSet bool
}

func (v NullableCollectionResourceCounts) Get() *CollectionResourceCounts {
	return v.value
}

func (v *NullableCollectionResourceCounts) Set(val *CollectionResourceCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceCounts(val *CollectionResourceCounts) *NullableCollectionResourceCounts {
	return &NullableCollectionResourceCounts{value: val, isSet: true}
}

func (v NullableCollectionResourceCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
