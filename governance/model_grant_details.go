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
)

// GrantDetails struct for GrantDetails
type GrantDetails struct {
	// Grant `id`
	Id                   *string            `json:"id,omitempty"`
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
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GrantDetails) HasId() bool {
	if o != nil && o.Id != nil {
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
	if o == nil || o.Metadata == nil {
		var ret GrantMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantDetails) GetMetadataOk() (*GrantMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GrantDetails) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
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
	if o == nil || o.Links == nil {
		var ret GrantDetailsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantDetails) GetLinksOk() (*GrantDetailsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GrantDetails) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GrantDetailsLinks and assigns it to the Links field.
func (o *GrantDetails) SetLinks(v GrantDetailsLinks) {
	o.Links = &v
}

func (o GrantDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GrantDetails) UnmarshalJSON(bytes []byte) (err error) {
	varGrantDetails := _GrantDetails{}

	err = json.Unmarshal(bytes, &varGrantDetails)
	if err == nil {
		*o = GrantDetails(varGrantDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
