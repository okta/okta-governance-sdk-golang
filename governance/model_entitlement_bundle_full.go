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
	"time"
)

// EntitlementBundleFull Full representation of a entitlement bundle resource
type EntitlementBundleFull struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetResourceOrn string                  `json:"targetResourceOrn"`
	Target            TargetResource          `json:"target"`
	Status            EntitlementBundleStatus `json:"status"`
	// Collection of entitlements and associated value identifiers
	Entitlements           []EntitlementCreatable `json:"entitlements,omitempty"`
	EntitlementsObjectType string                 `json:"entitlementsObjectType"`
	Links                  EntitlementBundleLinks `json:"_links"`
	// The unique name of the entitlement bundle
	Name string `json:"name"`
	// The human-readable description
	Description *string `json:"description,omitempty"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy        string `json:"lastUpdatedBy"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementBundleFull EntitlementBundleFull

// NewEntitlementBundleFull instantiates a new EntitlementBundleFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementBundleFull(targetResourceOrn string, target TargetResource, status EntitlementBundleStatus, entitlementsObjectType string, links EntitlementBundleLinks, name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *EntitlementBundleFull {
	this := EntitlementBundleFull{}
	this.Name = name
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	return &this
}

// NewEntitlementBundleFullWithDefaults instantiates a new EntitlementBundleFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementBundleFullWithDefaults() *EntitlementBundleFull {
	this := EntitlementBundleFull{}
	var entitlementsObjectType string = "entitlement-bundle-full"
	this.EntitlementsObjectType = entitlementsObjectType
	return &this
}

// GetTargetResourceOrn returns the TargetResourceOrn field value
func (o *EntitlementBundleFull) GetTargetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetResourceOrn
}

// GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetTargetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetResourceOrn, true
}

// SetTargetResourceOrn sets field value
func (o *EntitlementBundleFull) SetTargetResourceOrn(v string) {
	o.TargetResourceOrn = v
}

// GetTarget returns the Target field value
func (o *EntitlementBundleFull) GetTarget() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetTargetOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *EntitlementBundleFull) SetTarget(v TargetResource) {
	o.Target = v
}

// GetStatus returns the Status field value
func (o *EntitlementBundleFull) GetStatus() EntitlementBundleStatus {
	if o == nil {
		var ret EntitlementBundleStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetStatusOk() (*EntitlementBundleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EntitlementBundleFull) SetStatus(v EntitlementBundleStatus) {
	o.Status = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *EntitlementBundleFull) GetEntitlements() []EntitlementCreatable {
	if o == nil || o.Entitlements == nil {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *EntitlementBundleFull) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *EntitlementBundleFull) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetEntitlementsObjectType returns the EntitlementsObjectType field value
func (o *EntitlementBundleFull) GetEntitlementsObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementsObjectType
}

// GetEntitlementsObjectTypeOk returns a tuple with the EntitlementsObjectType field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetEntitlementsObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementsObjectType, true
}

// SetEntitlementsObjectType sets field value
func (o *EntitlementBundleFull) SetEntitlementsObjectType(v string) {
	o.EntitlementsObjectType = v
}

// GetLinks returns the Links field value
func (o *EntitlementBundleFull) GetLinks() EntitlementBundleLinks {
	if o == nil {
		var ret EntitlementBundleLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetLinksOk() (*EntitlementBundleLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EntitlementBundleFull) SetLinks(v EntitlementBundleLinks) {
	o.Links = v
}

// GetName returns the Name field value
func (o *EntitlementBundleFull) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementBundleFull) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementBundleFull) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementBundleFull) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementBundleFull) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *EntitlementBundleFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementBundleFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *EntitlementBundleFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *EntitlementBundleFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *EntitlementBundleFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *EntitlementBundleFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *EntitlementBundleFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *EntitlementBundleFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *EntitlementBundleFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *EntitlementBundleFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

func (o EntitlementBundleFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["targetResourceOrn"] = o.TargetResourceOrn
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if true {
		toSerialize["entitlementsObjectType"] = o.EntitlementsObjectType
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementBundleFull) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementBundleFull := _EntitlementBundleFull{}

	err = json.Unmarshal(bytes, &varEntitlementBundleFull)
	if err == nil {
		*o = EntitlementBundleFull(varEntitlementBundleFull)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "targetResourceOrn")
		delete(additionalProperties, "target")
		delete(additionalProperties, "status")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "entitlementsObjectType")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementBundleFull struct {
	value *EntitlementBundleFull
	isSet bool
}

func (v NullableEntitlementBundleFull) Get() *EntitlementBundleFull {
	return v.value
}

func (v *NullableEntitlementBundleFull) Set(val *EntitlementBundleFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBundleFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBundleFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBundleFull(val *EntitlementBundleFull) *NullableEntitlementBundleFull {
	return &NullableEntitlementBundleFull{value: val, isSet: true}
}

func (v NullableEntitlementBundleFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBundleFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
