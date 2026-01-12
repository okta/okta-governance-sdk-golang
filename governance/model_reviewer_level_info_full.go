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

// checks if the ReviewerLevelInfoFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerLevelInfoFull{}

// ReviewerLevelInfoFull Full representation of a reviewer level. Applicable for multi level campaigns only.
type ReviewerLevelInfoFull struct {
	ReviewerLevel        ReviewerLevelType         `json:"reviewerLevel"`
	Decision             Decision                  `json:"decision"`
	ReviewerProfile      *PrincipalProfileEnriched `json:"reviewerProfile,omitempty"`
	ReviewerType         ReviewersReviewerType     `json:"reviewerType"`
	ReviewerGroupProfile *ReviewerGroupProfile     `json:"reviewerGroupProfile,omitempty"`
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string `json:"lastUpdatedBy"`
	// Links to related resources
	Links                *map[string]Link `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerLevelInfoFull ReviewerLevelInfoFull

// NewReviewerLevelInfoFull instantiates a new ReviewerLevelInfoFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerLevelInfoFull(reviewerLevel ReviewerLevelType, decision Decision, reviewerType ReviewersReviewerType, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *ReviewerLevelInfoFull {
	this := ReviewerLevelInfoFull{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	return &this
}

// NewReviewerLevelInfoFullWithDefaults instantiates a new ReviewerLevelInfoFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerLevelInfoFullWithDefaults() *ReviewerLevelInfoFull {
	this := ReviewerLevelInfoFull{}
	return &this
}

// GetReviewerLevel returns the ReviewerLevel field value
func (o *ReviewerLevelInfoFull) GetReviewerLevel() ReviewerLevelType {
	if o == nil {
		var ret ReviewerLevelType
		return ret
	}

	return o.ReviewerLevel
}

// GetReviewerLevelOk returns a tuple with the ReviewerLevel field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetReviewerLevelOk() (*ReviewerLevelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerLevel, true
}

// SetReviewerLevel sets field value
func (o *ReviewerLevelInfoFull) SetReviewerLevel(v ReviewerLevelType) {
	o.ReviewerLevel = v
}

// GetDecision returns the Decision field value
func (o *ReviewerLevelInfoFull) GetDecision() Decision {
	if o == nil {
		var ret Decision
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetDecisionOk() (*Decision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *ReviewerLevelInfoFull) SetDecision(v Decision) {
	o.Decision = v
}

// GetReviewerProfile returns the ReviewerProfile field value if set, zero value otherwise.
func (o *ReviewerLevelInfoFull) GetReviewerProfile() PrincipalProfileEnriched {
	if o == nil || IsNil(o.ReviewerProfile) {
		var ret PrincipalProfileEnriched
		return ret
	}
	return *o.ReviewerProfile
}

// GetReviewerProfileOk returns a tuple with the ReviewerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetReviewerProfileOk() (*PrincipalProfileEnriched, bool) {
	if o == nil || IsNil(o.ReviewerProfile) {
		return nil, false
	}
	return o.ReviewerProfile, true
}

// HasReviewerProfile returns a boolean if a field has been set.
func (o *ReviewerLevelInfoFull) HasReviewerProfile() bool {
	if o != nil && !IsNil(o.ReviewerProfile) {
		return true
	}

	return false
}

// SetReviewerProfile gets a reference to the given PrincipalProfileEnriched and assigns it to the ReviewerProfile field.
func (o *ReviewerLevelInfoFull) SetReviewerProfile(v PrincipalProfileEnriched) {
	o.ReviewerProfile = &v
}

// GetReviewerType returns the ReviewerType field value
func (o *ReviewerLevelInfoFull) GetReviewerType() ReviewersReviewerType {
	if o == nil {
		var ret ReviewersReviewerType
		return ret
	}

	return o.ReviewerType
}

// GetReviewerTypeOk returns a tuple with the ReviewerType field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetReviewerTypeOk() (*ReviewersReviewerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerType, true
}

// SetReviewerType sets field value
func (o *ReviewerLevelInfoFull) SetReviewerType(v ReviewersReviewerType) {
	o.ReviewerType = v
}

// GetReviewerGroupProfile returns the ReviewerGroupProfile field value if set, zero value otherwise.
func (o *ReviewerLevelInfoFull) GetReviewerGroupProfile() ReviewerGroupProfile {
	if o == nil || IsNil(o.ReviewerGroupProfile) {
		var ret ReviewerGroupProfile
		return ret
	}
	return *o.ReviewerGroupProfile
}

// GetReviewerGroupProfileOk returns a tuple with the ReviewerGroupProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetReviewerGroupProfileOk() (*ReviewerGroupProfile, bool) {
	if o == nil || IsNil(o.ReviewerGroupProfile) {
		return nil, false
	}
	return o.ReviewerGroupProfile, true
}

// HasReviewerGroupProfile returns a boolean if a field has been set.
func (o *ReviewerLevelInfoFull) HasReviewerGroupProfile() bool {
	if o != nil && !IsNil(o.ReviewerGroupProfile) {
		return true
	}

	return false
}

// SetReviewerGroupProfile gets a reference to the given ReviewerGroupProfile and assigns it to the ReviewerGroupProfile field.
func (o *ReviewerLevelInfoFull) SetReviewerGroupProfile(v ReviewerGroupProfile) {
	o.ReviewerGroupProfile = &v
}

// GetId returns the Id field value
func (o *ReviewerLevelInfoFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewerLevelInfoFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ReviewerLevelInfoFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ReviewerLevelInfoFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *ReviewerLevelInfoFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ReviewerLevelInfoFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *ReviewerLevelInfoFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *ReviewerLevelInfoFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *ReviewerLevelInfoFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *ReviewerLevelInfoFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReviewerLevelInfoFull) GetLinks() map[string]Link {
	if o == nil || IsNil(o.Links) {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerLevelInfoFull) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReviewerLevelInfoFull) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *ReviewerLevelInfoFull) SetLinks(v map[string]Link) {
	o.Links = &v
}

func (o ReviewerLevelInfoFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerLevelInfoFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reviewerLevel"] = o.ReviewerLevel
	toSerialize["decision"] = o.Decision
	if !IsNil(o.ReviewerProfile) {
		toSerialize["reviewerProfile"] = o.ReviewerProfile
	}
	toSerialize["reviewerType"] = o.ReviewerType
	if !IsNil(o.ReviewerGroupProfile) {
		toSerialize["reviewerGroupProfile"] = o.ReviewerGroupProfile
	}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerLevelInfoFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reviewerLevel",
		"decision",
		"reviewerType",
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
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

	varReviewerLevelInfoFull := _ReviewerLevelInfoFull{}

	err = json.Unmarshal(data, &varReviewerLevelInfoFull)

	if err != nil {
		return err
	}

	*o = ReviewerLevelInfoFull(varReviewerLevelInfoFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reviewerLevel")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "reviewerProfile")
		delete(additionalProperties, "reviewerType")
		delete(additionalProperties, "reviewerGroupProfile")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerLevelInfoFull struct {
	value *ReviewerLevelInfoFull
	isSet bool
}

func (v NullableReviewerLevelInfoFull) Get() *ReviewerLevelInfoFull {
	return v.value
}

func (v *NullableReviewerLevelInfoFull) Set(val *ReviewerLevelInfoFull) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerLevelInfoFull) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerLevelInfoFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerLevelInfoFull(val *ReviewerLevelInfoFull) *NullableReviewerLevelInfoFull {
	return &NullableReviewerLevelInfoFull{value: val, isSet: true}
}

func (v NullableReviewerLevelInfoFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerLevelInfoFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
