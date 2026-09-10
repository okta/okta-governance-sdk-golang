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

// checks if the ReviewerAiAgentConnectedResourceMcpServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResourceMcpServer{}

// ReviewerAiAgentConnectedResourceMcpServer MCP server details for an AI agent connection
type ReviewerAiAgentConnectedResourceMcpServer struct {
	// ID of the connected MCP server
	Id *string `json:"id,omitempty"`
	// Okta Resource Name (ORN) of the connected MCP server
	Orn *string `json:"orn,omitempty"`
	// Name of the connected MCP server
	Name *string `json:"name,omitempty"`
	// Connected MCP server access URL for the AI agent
	EndpointUrl *string `json:"endpointUrl,omitempty"`
	// Connected API server scopes that the AI agent has access to
	Scopes               []string `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResourceMcpServer ReviewerAiAgentConnectedResourceMcpServer

// NewReviewerAiAgentConnectedResourceMcpServer instantiates a new ReviewerAiAgentConnectedResourceMcpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResourceMcpServer() *ReviewerAiAgentConnectedResourceMcpServer {
	this := ReviewerAiAgentConnectedResourceMcpServer{}
	return &this
}

// NewReviewerAiAgentConnectedResourceMcpServerWithDefaults instantiates a new ReviewerAiAgentConnectedResourceMcpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceMcpServerWithDefaults() *ReviewerAiAgentConnectedResourceMcpServer {
	this := ReviewerAiAgentConnectedResourceMcpServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReviewerAiAgentConnectedResourceMcpServer) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ReviewerAiAgentConnectedResourceMcpServer) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewerAiAgentConnectedResourceMcpServer) SetName(v string) {
	o.Name = &v
}

// GetEndpointUrl returns the EndpointUrl field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetEndpointUrl() string {
	if o == nil || IsNil(o.EndpointUrl) {
		var ret string
		return ret
	}
	return *o.EndpointUrl
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointUrl) {
		return nil, false
	}
	return o.EndpointUrl, true
}

// HasEndpointUrl returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) HasEndpointUrl() bool {
	if o != nil && !IsNil(o.EndpointUrl) {
		return true
	}

	return false
}

// SetEndpointUrl gets a reference to the given string and assigns it to the EndpointUrl field.
func (o *ReviewerAiAgentConnectedResourceMcpServer) SetEndpointUrl(v string) {
	o.EndpointUrl = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceMcpServer) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ReviewerAiAgentConnectedResourceMcpServer) SetScopes(v []string) {
	o.Scopes = v
}

func (o ReviewerAiAgentConnectedResourceMcpServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResourceMcpServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.EndpointUrl) {
		toSerialize["endpointUrl"] = o.EndpointUrl
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerAiAgentConnectedResourceMcpServer) UnmarshalJSON(data []byte) (err error) {
	varReviewerAiAgentConnectedResourceMcpServer := _ReviewerAiAgentConnectedResourceMcpServer{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResourceMcpServer)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResourceMcpServer(varReviewerAiAgentConnectedResourceMcpServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "endpointUrl")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResourceMcpServer struct {
	value *ReviewerAiAgentConnectedResourceMcpServer
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResourceMcpServer) Get() *ReviewerAiAgentConnectedResourceMcpServer {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResourceMcpServer) Set(val *ReviewerAiAgentConnectedResourceMcpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResourceMcpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResourceMcpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResourceMcpServer(val *ReviewerAiAgentConnectedResourceMcpServer) *NullableReviewerAiAgentConnectedResourceMcpServer {
	return &NullableReviewerAiAgentConnectedResourceMcpServer{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResourceMcpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResourceMcpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
