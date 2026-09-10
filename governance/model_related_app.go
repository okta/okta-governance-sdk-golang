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

// checks if the RelatedApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedApp{}

// RelatedApp An app that's related to this resource
type RelatedApp struct {
	// Okta resource instance ID
	Id *string `json:"id,omitempty"`
	// The name of the Okta resource
	Name *string `json:"name,omitempty"`
	// The label of the Okta app. Only populated for `APPLICATION` type resources.
	Label *string `json:"label,omitempty"`
	// The description of the resource
	Description *string `json:"description,omitempty"`
	// List of resource logo resources
	Logo []Link `json:"logo,omitempty"`
	// Indicates whether this app is also a resource in the current collection. Only present for `GROUP` type resources.
	InCollection         *bool `json:"inCollection,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RelatedApp RelatedApp

// NewRelatedApp instantiates a new RelatedApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedApp() *RelatedApp {
	this := RelatedApp{}
	return &this
}

// NewRelatedAppWithDefaults instantiates a new RelatedApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedAppWithDefaults() *RelatedApp {
	this := RelatedApp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelatedApp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedApp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelatedApp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelatedApp) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RelatedApp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedApp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RelatedApp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RelatedApp) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RelatedApp) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedApp) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RelatedApp) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RelatedApp) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RelatedApp) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedApp) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RelatedApp) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RelatedApp) SetDescription(v string) {
	o.Description = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *RelatedApp) GetLogo() []Link {
	if o == nil || IsNil(o.Logo) {
		var ret []Link
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedApp) GetLogoOk() ([]Link, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *RelatedApp) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []Link and assigns it to the Logo field.
func (o *RelatedApp) SetLogo(v []Link) {
	o.Logo = v
}

// GetInCollection returns the InCollection field value if set, zero value otherwise.
func (o *RelatedApp) GetInCollection() bool {
	if o == nil || IsNil(o.InCollection) {
		var ret bool
		return ret
	}
	return *o.InCollection
}

// GetInCollectionOk returns a tuple with the InCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedApp) GetInCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.InCollection) {
		return nil, false
	}
	return o.InCollection, true
}

// HasInCollection returns a boolean if a field has been set.
func (o *RelatedApp) HasInCollection() bool {
	if o != nil && !IsNil(o.InCollection) {
		return true
	}

	return false
}

// SetInCollection gets a reference to the given bool and assigns it to the InCollection field.
func (o *RelatedApp) SetInCollection(v bool) {
	o.InCollection = &v
}

func (o RelatedApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.InCollection) {
		toSerialize["inCollection"] = o.InCollection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RelatedApp) UnmarshalJSON(data []byte) (err error) {
	varRelatedApp := _RelatedApp{}

	err = json.Unmarshal(data, &varRelatedApp)

	if err != nil {
		return err
	}

	*o = RelatedApp(varRelatedApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "inCollection")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRelatedApp struct {
	value *RelatedApp
	isSet bool
}

func (v NullableRelatedApp) Get() *RelatedApp {
	return v.value
}

func (v *NullableRelatedApp) Set(val *RelatedApp) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedApp) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedApp(val *RelatedApp) *NullableRelatedApp {
	return &NullableRelatedApp{value: val, isSet: true}
}

func (v NullableRelatedApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
