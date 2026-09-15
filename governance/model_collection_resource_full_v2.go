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

// checks if the CollectionResourceFullV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResourceFullV2{}

// CollectionResourceFullV2 Full representation of a resource within a resource collection
type CollectionResourceFullV2 struct {
	ResourceProfile       *ResourceProfileV2               `json:"resourceProfile,omitempty"`
	ResourceConfiguration *CollectionResourceConfiguration `json:"resourceConfiguration,omitempty"`
	// Collection of entitlements with associated values
	Entitlements []EntitlementFull `json:"entitlements,omitempty"`
	// List of push groups associated with an app collection resource
	PushGroups []CollectionPushGroup `json:"pushGroups,omitempty"`
	// The governance labels associated with this resource
	Labels []LabelValue `json:"labels,omitempty"`
	// The apps related to this resource. Only populated for `GROUP` type resources. For `GROUP` resources with `hasPushMapping: true`, these are apps that this group is mapped to through push group. For `GROUP` resources with `hasPushMapping: false`, these are apps assigned to this group.
	RelatedApps []RelatedApp `json:"relatedApps,omitempty"`
	// The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.
	ResourceOrn string `json:"resourceOrn"`
	// The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as `GET`/`PUT`/`DELETE /v2/collections/{collectionId}/resources/{resourceId}`.
	ResourceId           *string                   `json:"resourceId,omitempty"`
	Links                CollectionResourceLinksV2 `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _CollectionResourceFullV2 CollectionResourceFullV2

// NewCollectionResourceFullV2 instantiates a new CollectionResourceFullV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResourceFullV2(resourceOrn string, links CollectionResourceLinksV2) *CollectionResourceFullV2 {
	this := CollectionResourceFullV2{}
	this.ResourceOrn = resourceOrn
	this.Links = links
	return &this
}

// NewCollectionResourceFullV2WithDefaults instantiates a new CollectionResourceFullV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceFullV2WithDefaults() *CollectionResourceFullV2 {
	this := CollectionResourceFullV2{}
	return &this
}

// GetResourceProfile returns the ResourceProfile field value if set, zero value otherwise.
func (o *CollectionResourceFullV2) GetResourceProfile() ResourceProfileV2 {
	if o == nil || IsNil(o.ResourceProfile) {
		var ret ResourceProfileV2
		return ret
	}
	return *o.ResourceProfile
}

// GetResourceProfileOk returns a tuple with the ResourceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetResourceProfileOk() (*ResourceProfileV2, bool) {
	if o == nil || IsNil(o.ResourceProfile) {
		return nil, false
	}
	return o.ResourceProfile, true
}

// HasResourceProfile returns a boolean if a field has been set.
func (o *CollectionResourceFullV2) HasResourceProfile() bool {
	if o != nil && !IsNil(o.ResourceProfile) {
		return true
	}

	return false
}

// SetResourceProfile gets a reference to the given ResourceProfileV2 and assigns it to the ResourceProfile field.
func (o *CollectionResourceFullV2) SetResourceProfile(v ResourceProfileV2) {
	o.ResourceProfile = &v
}

// GetResourceConfiguration returns the ResourceConfiguration field value if set, zero value otherwise.
func (o *CollectionResourceFullV2) GetResourceConfiguration() CollectionResourceConfiguration {
	if o == nil || IsNil(o.ResourceConfiguration) {
		var ret CollectionResourceConfiguration
		return ret
	}
	return *o.ResourceConfiguration
}

// GetResourceConfigurationOk returns a tuple with the ResourceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetResourceConfigurationOk() (*CollectionResourceConfiguration, bool) {
	if o == nil || IsNil(o.ResourceConfiguration) {
		return nil, false
	}
	return o.ResourceConfiguration, true
}

// HasResourceConfiguration returns a boolean if a field has been set.
func (o *CollectionResourceFullV2) HasResourceConfiguration() bool {
	if o != nil && !IsNil(o.ResourceConfiguration) {
		return true
	}

	return false
}

// SetResourceConfiguration gets a reference to the given CollectionResourceConfiguration and assigns it to the ResourceConfiguration field.
func (o *CollectionResourceFullV2) SetResourceConfiguration(v CollectionResourceConfiguration) {
	o.ResourceConfiguration = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *CollectionResourceFullV2) GetEntitlements() []EntitlementFull {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementFull
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetEntitlementsOk() ([]EntitlementFull, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *CollectionResourceFullV2) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementFull and assigns it to the Entitlements field.
func (o *CollectionResourceFullV2) SetEntitlements(v []EntitlementFull) {
	o.Entitlements = v
}

// GetPushGroups returns the PushGroups field value if set, zero value otherwise.
func (o *CollectionResourceFullV2) GetPushGroups() []CollectionPushGroup {
	if o == nil || IsNil(o.PushGroups) {
		var ret []CollectionPushGroup
		return ret
	}
	return o.PushGroups
}

// GetPushGroupsOk returns a tuple with the PushGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetPushGroupsOk() ([]CollectionPushGroup, bool) {
	if o == nil || IsNil(o.PushGroups) {
		return nil, false
	}
	return o.PushGroups, true
}

// HasPushGroups returns a boolean if a field has been set.
func (o *CollectionResourceFullV2) HasPushGroups() bool {
	if o != nil && !IsNil(o.PushGroups) {
		return true
	}

	return false
}

// SetPushGroups gets a reference to the given []CollectionPushGroup and assigns it to the PushGroups field.
func (o *CollectionResourceFullV2) SetPushGroups(v []CollectionPushGroup) {
	o.PushGroups = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CollectionResourceFullV2) GetLabels() []LabelValue {
	if o == nil || IsNil(o.Labels) {
		var ret []LabelValue
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetLabelsOk() ([]LabelValue, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CollectionResourceFullV2) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelValue and assigns it to the Labels field.
func (o *CollectionResourceFullV2) SetLabels(v []LabelValue) {
	o.Labels = v
}

// GetRelatedApps returns the RelatedApps field value if set, zero value otherwise.
func (o *CollectionResourceFullV2) GetRelatedApps() []RelatedApp {
	if o == nil || IsNil(o.RelatedApps) {
		var ret []RelatedApp
		return ret
	}
	return o.RelatedApps
}

// GetRelatedAppsOk returns a tuple with the RelatedApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetRelatedAppsOk() ([]RelatedApp, bool) {
	if o == nil || IsNil(o.RelatedApps) {
		return nil, false
	}
	return o.RelatedApps, true
}

// HasRelatedApps returns a boolean if a field has been set.
func (o *CollectionResourceFullV2) HasRelatedApps() bool {
	if o != nil && !IsNil(o.RelatedApps) {
		return true
	}

	return false
}

// SetRelatedApps gets a reference to the given []RelatedApp and assigns it to the RelatedApps field.
func (o *CollectionResourceFullV2) SetRelatedApps(v []RelatedApp) {
	o.RelatedApps = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *CollectionResourceFullV2) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *CollectionResourceFullV2) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *CollectionResourceFullV2) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CollectionResourceFullV2) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *CollectionResourceFullV2) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetLinks returns the Links field value
func (o *CollectionResourceFullV2) GetLinks() CollectionResourceLinksV2 {
	if o == nil {
		var ret CollectionResourceLinksV2
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CollectionResourceFullV2) GetLinksOk() (*CollectionResourceLinksV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CollectionResourceFullV2) SetLinks(v CollectionResourceLinksV2) {
	o.Links = v
}

func (o CollectionResourceFullV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResourceFullV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceProfile) {
		toSerialize["resourceProfile"] = o.ResourceProfile
	}
	if !IsNil(o.ResourceConfiguration) {
		toSerialize["resourceConfiguration"] = o.ResourceConfiguration
	}
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
	toSerialize["resourceOrn"] = o.ResourceOrn
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionResourceFullV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceOrn",
		"_links",
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

	varCollectionResourceFullV2 := _CollectionResourceFullV2{}

	err = json.Unmarshal(data, &varCollectionResourceFullV2)

	if err != nil {
		return err
	}

	*o = CollectionResourceFullV2(varCollectionResourceFullV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceProfile")
		delete(additionalProperties, "resourceConfiguration")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "pushGroups")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "relatedApps")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionResourceFullV2 struct {
	value *CollectionResourceFullV2
	isSet bool
}

func (v NullableCollectionResourceFullV2) Get() *CollectionResourceFullV2 {
	return v.value
}

func (v *NullableCollectionResourceFullV2) Set(val *CollectionResourceFullV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResourceFullV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResourceFullV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResourceFullV2(val *CollectionResourceFullV2) *NullableCollectionResourceFullV2 {
	return &NullableCollectionResourceFullV2{value: val, isSet: true}
}

func (v NullableCollectionResourceFullV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResourceFullV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
