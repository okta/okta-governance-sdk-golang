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
	"time"
)

// checks if the EntitlementDriftFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDriftFull{}

// EntitlementDriftFull A single entitlement drift, in full
type EntitlementDriftFull struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string                `json:"lastUpdatedBy"`
	Links         EntitlementDriftLinks `json:"_links"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn string `json:"resourceOrn"`
	// The Okta user in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format
	PrincipalOrn string `json:"principalOrn"`
	// The `id` of the import job that surfaced this drift
	ImportJobId         string                               `json:"importJobId"`
	DriftType           EntitlementDriftType                 `json:"driftType"`
	Status              EntitlementDriftStatus               `json:"status"`
	Resolution          EntitlementDriftResolution           `json:"resolution"`
	Entitlement         EntitlementDriftEntitlementReference `json:"entitlement"`
	AffectedGrantSource *EntitlementDriftAffectedGrantSource `json:"affectedGrantSource,omitempty"`
	// The `id` of the grant that held the affected entitlement when the drift was detected.
	AffectedGrantId *string `json:"affectedGrantId,omitempty"`
	// The `id` of the grant that resolution created or changed.
	ResolvedGrantId *string `json:"resolvedGrantId,omitempty"`
	// When the drift was reverted. Present only on a reverted drift.
	RevertedAt *time.Time `json:"revertedAt,omitempty"`
	// The principal who reverted the drift. Present only on a reverted drift.
	RevertedBy *string `json:"revertedBy,omitempty"`
	// A message that explains why Okta couldn't resolve the drift. Present only when `status` is `ERROR`.
	ErrorMessage         *string `json:"errorMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDriftFull EntitlementDriftFull

// NewEntitlementDriftFull instantiates a new EntitlementDriftFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDriftFull(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links EntitlementDriftLinks, resourceOrn string, principalOrn string, importJobId string, driftType EntitlementDriftType, status EntitlementDriftStatus, resolution EntitlementDriftResolution, entitlement EntitlementDriftEntitlementReference) *EntitlementDriftFull {
	this := EntitlementDriftFull{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.ResourceOrn = resourceOrn
	this.PrincipalOrn = principalOrn
	this.ImportJobId = importJobId
	this.DriftType = driftType
	this.Status = status
	this.Resolution = resolution
	this.Entitlement = entitlement
	return &this
}

// NewEntitlementDriftFullWithDefaults instantiates a new EntitlementDriftFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDriftFullWithDefaults() *EntitlementDriftFull {
	this := EntitlementDriftFull{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementDriftFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementDriftFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *EntitlementDriftFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *EntitlementDriftFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *EntitlementDriftFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *EntitlementDriftFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *EntitlementDriftFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *EntitlementDriftFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *EntitlementDriftFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *EntitlementDriftFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *EntitlementDriftFull) GetLinks() EntitlementDriftLinks {
	if o == nil {
		var ret EntitlementDriftLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetLinksOk() (*EntitlementDriftLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EntitlementDriftFull) SetLinks(v EntitlementDriftLinks) {
	o.Links = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *EntitlementDriftFull) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *EntitlementDriftFull) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *EntitlementDriftFull) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *EntitlementDriftFull) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetImportJobId returns the ImportJobId field value
func (o *EntitlementDriftFull) GetImportJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImportJobId
}

// GetImportJobIdOk returns a tuple with the ImportJobId field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetImportJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportJobId, true
}

// SetImportJobId sets field value
func (o *EntitlementDriftFull) SetImportJobId(v string) {
	o.ImportJobId = v
}

// GetDriftType returns the DriftType field value
func (o *EntitlementDriftFull) GetDriftType() EntitlementDriftType {
	if o == nil {
		var ret EntitlementDriftType
		return ret
	}

	return o.DriftType
}

// GetDriftTypeOk returns a tuple with the DriftType field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetDriftTypeOk() (*EntitlementDriftType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DriftType, true
}

// SetDriftType sets field value
func (o *EntitlementDriftFull) SetDriftType(v EntitlementDriftType) {
	o.DriftType = v
}

// GetStatus returns the Status field value
func (o *EntitlementDriftFull) GetStatus() EntitlementDriftStatus {
	if o == nil {
		var ret EntitlementDriftStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetStatusOk() (*EntitlementDriftStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EntitlementDriftFull) SetStatus(v EntitlementDriftStatus) {
	o.Status = v
}

// GetResolution returns the Resolution field value
func (o *EntitlementDriftFull) GetResolution() EntitlementDriftResolution {
	if o == nil {
		var ret EntitlementDriftResolution
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetResolutionOk() (*EntitlementDriftResolution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *EntitlementDriftFull) SetResolution(v EntitlementDriftResolution) {
	o.Resolution = v
}

// GetEntitlement returns the Entitlement field value
func (o *EntitlementDriftFull) GetEntitlement() EntitlementDriftEntitlementReference {
	if o == nil {
		var ret EntitlementDriftEntitlementReference
		return ret
	}

	return o.Entitlement
}

// GetEntitlementOk returns a tuple with the Entitlement field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetEntitlementOk() (*EntitlementDriftEntitlementReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entitlement, true
}

// SetEntitlement sets field value
func (o *EntitlementDriftFull) SetEntitlement(v EntitlementDriftEntitlementReference) {
	o.Entitlement = v
}

// GetAffectedGrantSource returns the AffectedGrantSource field value if set, zero value otherwise.
func (o *EntitlementDriftFull) GetAffectedGrantSource() EntitlementDriftAffectedGrantSource {
	if o == nil || IsNil(o.AffectedGrantSource) {
		var ret EntitlementDriftAffectedGrantSource
		return ret
	}
	return *o.AffectedGrantSource
}

// GetAffectedGrantSourceOk returns a tuple with the AffectedGrantSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetAffectedGrantSourceOk() (*EntitlementDriftAffectedGrantSource, bool) {
	if o == nil || IsNil(o.AffectedGrantSource) {
		return nil, false
	}
	return o.AffectedGrantSource, true
}

// HasAffectedGrantSource returns a boolean if a field has been set.
func (o *EntitlementDriftFull) HasAffectedGrantSource() bool {
	if o != nil && !IsNil(o.AffectedGrantSource) {
		return true
	}

	return false
}

// SetAffectedGrantSource gets a reference to the given EntitlementDriftAffectedGrantSource and assigns it to the AffectedGrantSource field.
func (o *EntitlementDriftFull) SetAffectedGrantSource(v EntitlementDriftAffectedGrantSource) {
	o.AffectedGrantSource = &v
}

// GetAffectedGrantId returns the AffectedGrantId field value if set, zero value otherwise.
func (o *EntitlementDriftFull) GetAffectedGrantId() string {
	if o == nil || IsNil(o.AffectedGrantId) {
		var ret string
		return ret
	}
	return *o.AffectedGrantId
}

// GetAffectedGrantIdOk returns a tuple with the AffectedGrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetAffectedGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.AffectedGrantId) {
		return nil, false
	}
	return o.AffectedGrantId, true
}

// HasAffectedGrantId returns a boolean if a field has been set.
func (o *EntitlementDriftFull) HasAffectedGrantId() bool {
	if o != nil && !IsNil(o.AffectedGrantId) {
		return true
	}

	return false
}

// SetAffectedGrantId gets a reference to the given string and assigns it to the AffectedGrantId field.
func (o *EntitlementDriftFull) SetAffectedGrantId(v string) {
	o.AffectedGrantId = &v
}

// GetResolvedGrantId returns the ResolvedGrantId field value if set, zero value otherwise.
func (o *EntitlementDriftFull) GetResolvedGrantId() string {
	if o == nil || IsNil(o.ResolvedGrantId) {
		var ret string
		return ret
	}
	return *o.ResolvedGrantId
}

// GetResolvedGrantIdOk returns a tuple with the ResolvedGrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetResolvedGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResolvedGrantId) {
		return nil, false
	}
	return o.ResolvedGrantId, true
}

// HasResolvedGrantId returns a boolean if a field has been set.
func (o *EntitlementDriftFull) HasResolvedGrantId() bool {
	if o != nil && !IsNil(o.ResolvedGrantId) {
		return true
	}

	return false
}

// SetResolvedGrantId gets a reference to the given string and assigns it to the ResolvedGrantId field.
func (o *EntitlementDriftFull) SetResolvedGrantId(v string) {
	o.ResolvedGrantId = &v
}

// GetRevertedAt returns the RevertedAt field value if set, zero value otherwise.
func (o *EntitlementDriftFull) GetRevertedAt() time.Time {
	if o == nil || IsNil(o.RevertedAt) {
		var ret time.Time
		return ret
	}
	return *o.RevertedAt
}

// GetRevertedAtOk returns a tuple with the RevertedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetRevertedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RevertedAt) {
		return nil, false
	}
	return o.RevertedAt, true
}

// HasRevertedAt returns a boolean if a field has been set.
func (o *EntitlementDriftFull) HasRevertedAt() bool {
	if o != nil && !IsNil(o.RevertedAt) {
		return true
	}

	return false
}

// SetRevertedAt gets a reference to the given time.Time and assigns it to the RevertedAt field.
func (o *EntitlementDriftFull) SetRevertedAt(v time.Time) {
	o.RevertedAt = &v
}

// GetRevertedBy returns the RevertedBy field value if set, zero value otherwise.
func (o *EntitlementDriftFull) GetRevertedBy() string {
	if o == nil || IsNil(o.RevertedBy) {
		var ret string
		return ret
	}
	return *o.RevertedBy
}

// GetRevertedByOk returns a tuple with the RevertedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetRevertedByOk() (*string, bool) {
	if o == nil || IsNil(o.RevertedBy) {
		return nil, false
	}
	return o.RevertedBy, true
}

// HasRevertedBy returns a boolean if a field has been set.
func (o *EntitlementDriftFull) HasRevertedBy() bool {
	if o != nil && !IsNil(o.RevertedBy) {
		return true
	}

	return false
}

// SetRevertedBy gets a reference to the given string and assigns it to the RevertedBy field.
func (o *EntitlementDriftFull) SetRevertedBy(v string) {
	o.RevertedBy = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *EntitlementDriftFull) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftFull) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *EntitlementDriftFull) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *EntitlementDriftFull) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o EntitlementDriftFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDriftFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	toSerialize["_links"] = o.Links
	toSerialize["resourceOrn"] = o.ResourceOrn
	toSerialize["principalOrn"] = o.PrincipalOrn
	toSerialize["importJobId"] = o.ImportJobId
	toSerialize["driftType"] = o.DriftType
	toSerialize["status"] = o.Status
	toSerialize["resolution"] = o.Resolution
	toSerialize["entitlement"] = o.Entitlement
	if !IsNil(o.AffectedGrantSource) {
		toSerialize["affectedGrantSource"] = o.AffectedGrantSource
	}
	if !IsNil(o.AffectedGrantId) {
		toSerialize["affectedGrantId"] = o.AffectedGrantId
	}
	if !IsNil(o.ResolvedGrantId) {
		toSerialize["resolvedGrantId"] = o.ResolvedGrantId
	}
	if !IsNil(o.RevertedAt) {
		toSerialize["revertedAt"] = o.RevertedAt
	}
	if !IsNil(o.RevertedBy) {
		toSerialize["revertedBy"] = o.RevertedBy
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDriftFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"_links",
		"resourceOrn",
		"principalOrn",
		"importJobId",
		"driftType",
		"status",
		"resolution",
		"entitlement",
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

	varEntitlementDriftFull := _EntitlementDriftFull{}

	err = json.Unmarshal(data, &varEntitlementDriftFull)

	if err != nil {
		return err
	}

	*o = EntitlementDriftFull(varEntitlementDriftFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "importJobId")
		delete(additionalProperties, "driftType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "resolution")
		delete(additionalProperties, "entitlement")
		delete(additionalProperties, "affectedGrantSource")
		delete(additionalProperties, "affectedGrantId")
		delete(additionalProperties, "resolvedGrantId")
		delete(additionalProperties, "revertedAt")
		delete(additionalProperties, "revertedBy")
		delete(additionalProperties, "errorMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDriftFull struct {
	value *EntitlementDriftFull
	isSet bool
}

func (v NullableEntitlementDriftFull) Get() *EntitlementDriftFull {
	return v.value
}

func (v *NullableEntitlementDriftFull) Set(val *EntitlementDriftFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftFull(val *EntitlementDriftFull) *NullableEntitlementDriftFull {
	return &NullableEntitlementDriftFull{value: val, isSet: true}
}

func (v NullableEntitlementDriftFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
