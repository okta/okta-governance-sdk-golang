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

// checks if the ReviewerAiAgentConnectedResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResource{}

// ReviewerAiAgentConnectedResource AI agent connection resource details
type ReviewerAiAgentConnectedResource struct {
	AuthorizationServer  *ReviewerAiAgentConnectedResourceAuthorizationServer `json:"authorizationServer,omitempty"`
	App                  *ReviewerAiAgentConnectedResourceApp                 `json:"app,omitempty"`
	McpServer            *ReviewerAiAgentConnectedResourceMcpServer           `json:"mcpServer,omitempty"`
	ApiServer            *ReviewerAiAgentConnectedResourceApiServer           `json:"apiServer,omitempty"`
	ServiceAccount       *ReviewerAiAgentConnectedResourceServiceAccount      `json:"serviceAccount,omitempty"`
	OpaSecret            *ReviewerAiAgentConnectedResourceSecret              `json:"opaSecret,omitempty"`
	Type                 AiAgentConnectedResourceType                         `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResource ReviewerAiAgentConnectedResource

// NewReviewerAiAgentConnectedResource instantiates a new ReviewerAiAgentConnectedResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResource(type_ AiAgentConnectedResourceType) *ReviewerAiAgentConnectedResource {
	this := ReviewerAiAgentConnectedResource{}
	this.Type = type_
	return &this
}

// NewReviewerAiAgentConnectedResourceWithDefaults instantiates a new ReviewerAiAgentConnectedResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceWithDefaults() *ReviewerAiAgentConnectedResource {
	this := ReviewerAiAgentConnectedResource{}
	return &this
}

// GetAuthorizationServer returns the AuthorizationServer field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResource) GetAuthorizationServer() ReviewerAiAgentConnectedResourceAuthorizationServer {
	if o == nil || IsNil(o.AuthorizationServer) {
		var ret ReviewerAiAgentConnectedResourceAuthorizationServer
		return ret
	}
	return *o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResource) GetAuthorizationServerOk() (*ReviewerAiAgentConnectedResourceAuthorizationServer, bool) {
	if o == nil || IsNil(o.AuthorizationServer) {
		return nil, false
	}
	return o.AuthorizationServer, true
}

// HasAuthorizationServer returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResource) HasAuthorizationServer() bool {
	if o != nil && !IsNil(o.AuthorizationServer) {
		return true
	}

	return false
}

// SetAuthorizationServer gets a reference to the given ReviewerAiAgentConnectedResourceAuthorizationServer and assigns it to the AuthorizationServer field.
func (o *ReviewerAiAgentConnectedResource) SetAuthorizationServer(v ReviewerAiAgentConnectedResourceAuthorizationServer) {
	o.AuthorizationServer = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResource) GetApp() ReviewerAiAgentConnectedResourceApp {
	if o == nil || IsNil(o.App) {
		var ret ReviewerAiAgentConnectedResourceApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResource) GetAppOk() (*ReviewerAiAgentConnectedResourceApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResource) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given ReviewerAiAgentConnectedResourceApp and assigns it to the App field.
func (o *ReviewerAiAgentConnectedResource) SetApp(v ReviewerAiAgentConnectedResourceApp) {
	o.App = &v
}

// GetMcpServer returns the McpServer field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResource) GetMcpServer() ReviewerAiAgentConnectedResourceMcpServer {
	if o == nil || IsNil(o.McpServer) {
		var ret ReviewerAiAgentConnectedResourceMcpServer
		return ret
	}
	return *o.McpServer
}

// GetMcpServerOk returns a tuple with the McpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResource) GetMcpServerOk() (*ReviewerAiAgentConnectedResourceMcpServer, bool) {
	if o == nil || IsNil(o.McpServer) {
		return nil, false
	}
	return o.McpServer, true
}

// HasMcpServer returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResource) HasMcpServer() bool {
	if o != nil && !IsNil(o.McpServer) {
		return true
	}

	return false
}

// SetMcpServer gets a reference to the given ReviewerAiAgentConnectedResourceMcpServer and assigns it to the McpServer field.
func (o *ReviewerAiAgentConnectedResource) SetMcpServer(v ReviewerAiAgentConnectedResourceMcpServer) {
	o.McpServer = &v
}

// GetApiServer returns the ApiServer field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResource) GetApiServer() ReviewerAiAgentConnectedResourceApiServer {
	if o == nil || IsNil(o.ApiServer) {
		var ret ReviewerAiAgentConnectedResourceApiServer
		return ret
	}
	return *o.ApiServer
}

// GetApiServerOk returns a tuple with the ApiServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResource) GetApiServerOk() (*ReviewerAiAgentConnectedResourceApiServer, bool) {
	if o == nil || IsNil(o.ApiServer) {
		return nil, false
	}
	return o.ApiServer, true
}

// HasApiServer returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResource) HasApiServer() bool {
	if o != nil && !IsNil(o.ApiServer) {
		return true
	}

	return false
}

// SetApiServer gets a reference to the given ReviewerAiAgentConnectedResourceApiServer and assigns it to the ApiServer field.
func (o *ReviewerAiAgentConnectedResource) SetApiServer(v ReviewerAiAgentConnectedResourceApiServer) {
	o.ApiServer = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResource) GetServiceAccount() ReviewerAiAgentConnectedResourceServiceAccount {
	if o == nil || IsNil(o.ServiceAccount) {
		var ret ReviewerAiAgentConnectedResourceServiceAccount
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResource) GetServiceAccountOk() (*ReviewerAiAgentConnectedResourceServiceAccount, bool) {
	if o == nil || IsNil(o.ServiceAccount) {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResource) HasServiceAccount() bool {
	if o != nil && !IsNil(o.ServiceAccount) {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given ReviewerAiAgentConnectedResourceServiceAccount and assigns it to the ServiceAccount field.
func (o *ReviewerAiAgentConnectedResource) SetServiceAccount(v ReviewerAiAgentConnectedResourceServiceAccount) {
	o.ServiceAccount = &v
}

// GetOpaSecret returns the OpaSecret field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResource) GetOpaSecret() ReviewerAiAgentConnectedResourceSecret {
	if o == nil || IsNil(o.OpaSecret) {
		var ret ReviewerAiAgentConnectedResourceSecret
		return ret
	}
	return *o.OpaSecret
}

// GetOpaSecretOk returns a tuple with the OpaSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResource) GetOpaSecretOk() (*ReviewerAiAgentConnectedResourceSecret, bool) {
	if o == nil || IsNil(o.OpaSecret) {
		return nil, false
	}
	return o.OpaSecret, true
}

// HasOpaSecret returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResource) HasOpaSecret() bool {
	if o != nil && !IsNil(o.OpaSecret) {
		return true
	}

	return false
}

// SetOpaSecret gets a reference to the given ReviewerAiAgentConnectedResourceSecret and assigns it to the OpaSecret field.
func (o *ReviewerAiAgentConnectedResource) SetOpaSecret(v ReviewerAiAgentConnectedResourceSecret) {
	o.OpaSecret = &v
}

// GetType returns the Type field value
func (o *ReviewerAiAgentConnectedResource) GetType() AiAgentConnectedResourceType {
	if o == nil {
		var ret AiAgentConnectedResourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResource) GetTypeOk() (*AiAgentConnectedResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReviewerAiAgentConnectedResource) SetType(v AiAgentConnectedResourceType) {
	o.Type = v
}

func (o ReviewerAiAgentConnectedResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationServer) {
		toSerialize["authorizationServer"] = o.AuthorizationServer
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.McpServer) {
		toSerialize["mcpServer"] = o.McpServer
	}
	if !IsNil(o.ApiServer) {
		toSerialize["apiServer"] = o.ApiServer
	}
	if !IsNil(o.ServiceAccount) {
		toSerialize["serviceAccount"] = o.ServiceAccount
	}
	if !IsNil(o.OpaSecret) {
		toSerialize["opaSecret"] = o.OpaSecret
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerAiAgentConnectedResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varReviewerAiAgentConnectedResource := _ReviewerAiAgentConnectedResource{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResource)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResource(varReviewerAiAgentConnectedResource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizationServer")
		delete(additionalProperties, "app")
		delete(additionalProperties, "mcpServer")
		delete(additionalProperties, "apiServer")
		delete(additionalProperties, "serviceAccount")
		delete(additionalProperties, "opaSecret")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResource struct {
	value *ReviewerAiAgentConnectedResource
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResource) Get() *ReviewerAiAgentConnectedResource {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResource) Set(val *ReviewerAiAgentConnectedResource) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResource) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResource(val *ReviewerAiAgentConnectedResource) *NullableReviewerAiAgentConnectedResource {
	return &NullableReviewerAiAgentConnectedResource{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
