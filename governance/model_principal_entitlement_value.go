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

// checks if the PrincipalEntitlementValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalEntitlementValue{}

// PrincipalEntitlementValue An entitlement value
type PrincipalEntitlementValue struct {
	// The `id` of the entitlement value
	Id string `json:"id"`
	// The display name for an entitlement value
	Name string `json:"name"`
	// The value of an entitlement property value
	ExternalValue string `json:"externalValue"`
	// The read-only ID of an entitlement property value in the downstream app
	ExternalId *string `json:"externalId,omitempty"`
	// The description of an entitlement value
	Description *string `json:"description,omitempty"`
	// The entitlement value resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)
	Orn *string `json:"orn,omitempty"`
	// The `id` of the Okta user who created the resource
	CreatedBy *string `json:"createdBy,omitempty"`
	// The ISO 8601 formatted date and time when the resource was created
	Created *time.Time `json:"created,omitempty"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy        *string `json:"lastUpdatedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementValue PrincipalEntitlementValue

// NewPrincipalEntitlementValue instantiates a new PrincipalEntitlementValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementValue(id string, name string, externalValue string) *PrincipalEntitlementValue {
	this := PrincipalEntitlementValue{}
	this.Id = id
	this.Name = name
	this.ExternalValue = externalValue
	return &this
}

// NewPrincipalEntitlementValueWithDefaults instantiates a new PrincipalEntitlementValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementValueWithDefaults() *PrincipalEntitlementValue {
	this := PrincipalEntitlementValue{}
	return &this
}

// GetId returns the Id field value
func (o *PrincipalEntitlementValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrincipalEntitlementValue) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PrincipalEntitlementValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrincipalEntitlementValue) SetName(v string) {
	o.Name = v
}

// GetExternalValue returns the ExternalValue field value
func (o *PrincipalEntitlementValue) GetExternalValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetExternalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalValue, true
}

// SetExternalValue sets field value
func (o *PrincipalEntitlementValue) SetExternalValue(v string) {
	o.ExternalValue = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *PrincipalEntitlementValue) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *PrincipalEntitlementValue) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *PrincipalEntitlementValue) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrincipalEntitlementValue) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrincipalEntitlementValue) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PrincipalEntitlementValue) SetDescription(v string) {
	o.Description = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *PrincipalEntitlementValue) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *PrincipalEntitlementValue) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *PrincipalEntitlementValue) SetOrn(v string) {
	o.Orn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *PrincipalEntitlementValue) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PrincipalEntitlementValue) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *PrincipalEntitlementValue) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PrincipalEntitlementValue) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PrincipalEntitlementValue) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PrincipalEntitlementValue) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PrincipalEntitlementValue) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PrincipalEntitlementValue) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PrincipalEntitlementValue) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *PrincipalEntitlementValue) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementValue) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *PrincipalEntitlementValue) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *PrincipalEntitlementValue) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

func (o PrincipalEntitlementValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalEntitlementValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["externalValue"] = o.ExternalValue
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalEntitlementValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"externalValue",
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

	varPrincipalEntitlementValue := _PrincipalEntitlementValue{}

	err = json.Unmarshal(data, &varPrincipalEntitlementValue)

	if err != nil {
		return err
	}

	*o = PrincipalEntitlementValue(varPrincipalEntitlementValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalEntitlementValue struct {
	value *PrincipalEntitlementValue
	isSet bool
}

func (v NullablePrincipalEntitlementValue) Get() *PrincipalEntitlementValue {
	return v.value
}

func (v *NullablePrincipalEntitlementValue) Set(val *PrincipalEntitlementValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementValue(val *PrincipalEntitlementValue) *NullablePrincipalEntitlementValue {
	return &NullablePrincipalEntitlementValue{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
