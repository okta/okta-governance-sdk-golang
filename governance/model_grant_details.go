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

// checks if the GrantDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantDetails{}

// GrantDetails struct for GrantDetails
type GrantDetails struct {
	// Grant `id`
	Id                   *string            `json:"id,omitempty" validate:"regexp=gra[0-9a-zA-Z]+"`
	Metadata             *GrantMetadata     `json:"metadata,omitempty"`
	Links                *GrantDetailsLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantDetails GrantDetails

// NewGrantDetails instantiates a new GrantDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantDetails() *GrantDetails {
	this := GrantDetails{}
	return &this
}

// NewGrantDetailsWithDefaults instantiates a new GrantDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantDetailsWithDefaults() *GrantDetails {
	this := GrantDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GrantDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GrantDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GrantDetails) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GrantDetails) GetMetadata() GrantMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret GrantMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantDetails) GetMetadataOk() (*GrantMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GrantDetails) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given GrantMetadata and assigns it to the Metadata field.
func (o *GrantDetails) SetMetadata(v GrantMetadata) {
	o.Metadata = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GrantDetails) GetLinks() GrantDetailsLinks {
	if o == nil || IsNil(o.Links) {
		var ret GrantDetailsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantDetails) GetLinksOk() (*GrantDetailsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GrantDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GrantDetailsLinks and assigns it to the Links field.
func (o *GrantDetails) SetLinks(v GrantDetailsLinks) {
	o.Links = &v
}

func (o GrantDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantDetails) UnmarshalJSON(data []byte) (err error) {
	varGrantDetails := _GrantDetails{}

	err = json.Unmarshal(data, &varGrantDetails)

	if err != nil {
		return err
	}

	*o = GrantDetails(varGrantDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantDetails struct {
	value *GrantDetails
	isSet bool
}

func (v NullableGrantDetails) Get() *GrantDetails {
	return v.value
}

func (v *NullableGrantDetails) Set(val *GrantDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantDetails(val *GrantDetails) *NullableGrantDetails {
	return &NullableGrantDetails{value: val, isSet: true}
}

func (v NullableGrantDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
