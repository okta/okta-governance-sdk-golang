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

// checks if the CatalogEntryRequestFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogEntryRequestFields{}

// CatalogEntryRequestFields struct for CatalogEntryRequestFields
type CatalogEntryRequestFields struct {
	// List of request fields
	Data                 []RequestField                     `json:"data,omitempty"`
	Metadata             *CatalogEntryRequestFieldsMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogEntryRequestFields CatalogEntryRequestFields

// NewCatalogEntryRequestFields instantiates a new CatalogEntryRequestFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogEntryRequestFields() *CatalogEntryRequestFields {
	this := CatalogEntryRequestFields{}
	return &this
}

// NewCatalogEntryRequestFieldsWithDefaults instantiates a new CatalogEntryRequestFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogEntryRequestFieldsWithDefaults() *CatalogEntryRequestFields {
	this := CatalogEntryRequestFields{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CatalogEntryRequestFields) GetData() []RequestField {
	if o == nil || IsNil(o.Data) {
		var ret []RequestField
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntryRequestFields) GetDataOk() ([]RequestField, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CatalogEntryRequestFields) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RequestField and assigns it to the Data field.
func (o *CatalogEntryRequestFields) SetData(v []RequestField) {
	o.Data = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CatalogEntryRequestFields) GetMetadata() CatalogEntryRequestFieldsMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret CatalogEntryRequestFieldsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntryRequestFields) GetMetadataOk() (*CatalogEntryRequestFieldsMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CatalogEntryRequestFields) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given CatalogEntryRequestFieldsMetadata and assigns it to the Metadata field.
func (o *CatalogEntryRequestFields) SetMetadata(v CatalogEntryRequestFieldsMetadata) {
	o.Metadata = &v
}

func (o CatalogEntryRequestFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogEntryRequestFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogEntryRequestFields) UnmarshalJSON(data []byte) (err error) {
	varCatalogEntryRequestFields := _CatalogEntryRequestFields{}

	err = json.Unmarshal(data, &varCatalogEntryRequestFields)

	if err != nil {
		return err
	}

	*o = CatalogEntryRequestFields(varCatalogEntryRequestFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCatalogEntryRequestFields struct {
	value *CatalogEntryRequestFields
	isSet bool
}

func (v NullableCatalogEntryRequestFields) Get() *CatalogEntryRequestFields {
	return v.value
}

func (v *NullableCatalogEntryRequestFields) Set(val *CatalogEntryRequestFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogEntryRequestFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogEntryRequestFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogEntryRequestFields(val *CatalogEntryRequestFields) *NullableCatalogEntryRequestFields {
	return &NullableCatalogEntryRequestFields{value: val, isSet: true}
}

func (v NullableCatalogEntryRequestFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogEntryRequestFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
