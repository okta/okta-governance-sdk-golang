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

// checks if the RequestTypeSparse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeSparse{}

// RequestTypeSparse Sparse representation of a Request Type resource
type RequestTypeSparse struct {
	Status           RequestTypeStatus `json:"status"`
	LastUpdateSource string            `json:"lastUpdateSource"`
	Links            RequestTypeLinks  `json:"_links"`
	// Writable unique key on Create. Not modifiable on update.
	Name string `json:"name"`
	// Human readable description
	Description string `json:"description"`
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

type _RequestTypeSparse RequestTypeSparse

// NewRequestTypeSparse instantiates a new RequestTypeSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeSparse(status RequestTypeStatus, lastUpdateSource string, links RequestTypeLinks, name string, description string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string) *RequestTypeSparse {
	this := RequestTypeSparse{}
	this.Name = name
	this.Description = description
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	return &this
}

// NewRequestTypeSparseWithDefaults instantiates a new RequestTypeSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeSparseWithDefaults() *RequestTypeSparse {
	this := RequestTypeSparse{}
	return &this
}

// GetStatus returns the Status field value
func (o *RequestTypeSparse) GetStatus() RequestTypeStatus {
	if o == nil {
		var ret RequestTypeStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetStatusOk() (*RequestTypeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestTypeSparse) SetStatus(v RequestTypeStatus) {
	o.Status = v
}

// GetLastUpdateSource returns the LastUpdateSource field value
func (o *RequestTypeSparse) GetLastUpdateSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdateSource
}

// GetLastUpdateSourceOk returns a tuple with the LastUpdateSource field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetLastUpdateSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdateSource, true
}

// SetLastUpdateSource sets field value
func (o *RequestTypeSparse) SetLastUpdateSource(v string) {
	o.LastUpdateSource = v
}

// GetLinks returns the Links field value
func (o *RequestTypeSparse) GetLinks() RequestTypeLinks {
	if o == nil {
		var ret RequestTypeLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetLinksOk() (*RequestTypeLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestTypeSparse) SetLinks(v RequestTypeLinks) {
	o.Links = v
}

// GetName returns the Name field value
func (o *RequestTypeSparse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestTypeSparse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *RequestTypeSparse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RequestTypeSparse) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *RequestTypeSparse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestTypeSparse) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestTypeSparse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestTypeSparse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestTypeSparse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestTypeSparse) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestTypeSparse) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestTypeSparse) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestTypeSparse) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestTypeSparse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestTypeSparse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

func (o RequestTypeSparse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeSparse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["lastUpdateSource"] = o.LastUpdateSource
	toSerialize["_links"] = o.Links
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeSparse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"lastUpdateSource",
		"_links",
		"name",
		"description",
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

	varRequestTypeSparse := _RequestTypeSparse{}

	err = json.Unmarshal(data, &varRequestTypeSparse)

	if err != nil {
		return err
	}

	*o = RequestTypeSparse(varRequestTypeSparse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastUpdateSource")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeSparse struct {
	value *RequestTypeSparse
	isSet bool
}

func (v NullableRequestTypeSparse) Get() *RequestTypeSparse {
	return v.value
}

func (v *NullableRequestTypeSparse) Set(val *RequestTypeSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeSparse(val *RequestTypeSparse) *NullableRequestTypeSparse {
	return &NullableRequestTypeSparse{value: val, isSet: true}
}

func (v NullableRequestTypeSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
