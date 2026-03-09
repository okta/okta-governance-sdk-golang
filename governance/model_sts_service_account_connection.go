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

// checks if the StsServiceAccountConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsServiceAccountConnection{}

// StsServiceAccountConnection STS connection to a service account
type StsServiceAccountConnection struct {
	// Type of connection authentication method
	ConnectionType string                          `json:"connectionType"`
	App            ManagedConnectionAppInstance    `json:"app"`
	ServiceAccount ManagedConnectionServiceAccount `json:"serviceAccount"`
	// Unique identifier for the managed connection
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection
	Orn *string `json:"orn,omitempty"`
	// The status of the connection
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StsServiceAccountConnection StsServiceAccountConnection

// NewStsServiceAccountConnection instantiates a new StsServiceAccountConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsServiceAccountConnection(connectionType string, app ManagedConnectionAppInstance, serviceAccount ManagedConnectionServiceAccount) *StsServiceAccountConnection {
	this := StsServiceAccountConnection{}
	return &this
}

// NewStsServiceAccountConnectionWithDefaults instantiates a new StsServiceAccountConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsServiceAccountConnectionWithDefaults() *StsServiceAccountConnection {
	this := StsServiceAccountConnection{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *StsServiceAccountConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *StsServiceAccountConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *StsServiceAccountConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetApp returns the App field value
func (o *StsServiceAccountConnection) GetApp() ManagedConnectionAppInstance {
	if o == nil {
		var ret ManagedConnectionAppInstance
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *StsServiceAccountConnection) GetAppOk() (*ManagedConnectionAppInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *StsServiceAccountConnection) SetApp(v ManagedConnectionAppInstance) {
	o.App = v
}

// GetServiceAccount returns the ServiceAccount field value
func (o *StsServiceAccountConnection) GetServiceAccount() ManagedConnectionServiceAccount {
	if o == nil {
		var ret ManagedConnectionServiceAccount
		return ret
	}

	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *StsServiceAccountConnection) GetServiceAccountOk() (*ManagedConnectionServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value
func (o *StsServiceAccountConnection) SetServiceAccount(v ManagedConnectionServiceAccount) {
	o.ServiceAccount = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StsServiceAccountConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsServiceAccountConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StsServiceAccountConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StsServiceAccountConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *StsServiceAccountConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsServiceAccountConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *StsServiceAccountConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *StsServiceAccountConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StsServiceAccountConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsServiceAccountConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StsServiceAccountConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StsServiceAccountConnection) SetStatus(v string) {
	o.Status = &v
}

func (o StsServiceAccountConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsServiceAccountConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionType"] = o.ConnectionType
	toSerialize["app"] = o.App
	toSerialize["serviceAccount"] = o.ServiceAccount
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StsServiceAccountConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionType",
		"app",
		"serviceAccount",
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

	varStsServiceAccountConnection := _StsServiceAccountConnection{}

	err = json.Unmarshal(data, &varStsServiceAccountConnection)

	if err != nil {
		return err
	}

	*o = StsServiceAccountConnection(varStsServiceAccountConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connectionType")
		delete(additionalProperties, "app")
		delete(additionalProperties, "serviceAccount")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStsServiceAccountConnection struct {
	value *StsServiceAccountConnection
	isSet bool
}

func (v NullableStsServiceAccountConnection) Get() *StsServiceAccountConnection {
	return v.value
}

func (v *NullableStsServiceAccountConnection) Set(val *StsServiceAccountConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableStsServiceAccountConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableStsServiceAccountConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsServiceAccountConnection(val *StsServiceAccountConnection) *NullableStsServiceAccountConnection {
	return &NullableStsServiceAccountConnection{value: val, isSet: true}
}

func (v NullableStsServiceAccountConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsServiceAccountConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
