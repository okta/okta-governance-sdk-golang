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
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type PrincipalEntitlementsAPI interface {

	/*
		GetPrincipalEntitlements Retrieve the principal's effective entitlements for a resource

		Retrieves the principal's effective entitlements for a specific resource

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetPrincipalEntitlementsRequest
	*/
	GetPrincipalEntitlements(ctx context.Context) ApiGetPrincipalEntitlementsRequest

	// GetPrincipalEntitlementsExecute executes the request
	//  @return PrincipalEntitlementsList
	GetPrincipalEntitlementsExecute(r ApiGetPrincipalEntitlementsRequest) (*PrincipalEntitlementsList, *APIResponse, error)

	/*
		GetPrincipalEntitlementsChanges Retrieve the principal entitlement changes

		Retrieves the principal's old and new effective entitlement changes for a resource

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param principalEntitlementsChangeId The entitlement change ID, in the format of `{old_effective_grant_id}:{new_effective_grant_id}`.  Only the `{new_effective_grant_id}` value is required if the `{old_effective_grant_id}` value isn't available.
		@return ApiGetPrincipalEntitlementsChangesRequest
	*/
	GetPrincipalEntitlementsChanges(ctx context.Context, principalEntitlementsChangeId string) ApiGetPrincipalEntitlementsChangesRequest

	// GetPrincipalEntitlementsChangesExecute executes the request
	//  @return PrincipalEntitlementsChange
	GetPrincipalEntitlementsChangesExecute(r ApiGetPrincipalEntitlementsChangesRequest) (*PrincipalEntitlementsChange, *APIResponse, error)

	/*
			GetPrincipalEntitlementsHistory Retrieve an entitlement history

			Retrieves an entitlement history log for a specific principal and resource.

		Specify the required principal and resource references in the `filter` query parameter.
		You can optionally specify a time range for the returned history log.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiGetPrincipalEntitlementsHistoryRequest
	*/
	GetPrincipalEntitlementsHistory(ctx context.Context) ApiGetPrincipalEntitlementsHistoryRequest

	// GetPrincipalEntitlementsHistoryExecute executes the request
	//  @return PrincipalEntitlementsHistory
	GetPrincipalEntitlementsHistoryExecute(r ApiGetPrincipalEntitlementsHistoryRequest) (*PrincipalEntitlementsHistory, *APIResponse, error)
}

// PrincipalEntitlementsAPIService PrincipalEntitlementsAPI service
type PrincipalEntitlementsAPIService service

type ApiGetPrincipalEntitlementsRequest struct {
	ctx        context.Context
	ApiService PrincipalEntitlementsAPI
	filter     *string
	retryCount int32
}

// Apply various filters by using supported principal entitlements filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiGetPrincipalEntitlementsRequest) Filter(filter string) ApiGetPrincipalEntitlementsRequest {
	r.filter = &filter
	return r
}

func (r ApiGetPrincipalEntitlementsRequest) Execute() (*PrincipalEntitlementsList, *APIResponse, error) {
	return r.ApiService.GetPrincipalEntitlementsExecute(r)
}

/*
GetPrincipalEntitlements Retrieve the principal's effective entitlements for a resource

Retrieves the principal's effective entitlements for a specific resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPrincipalEntitlementsRequest
*/
func (a *PrincipalEntitlementsAPIService) GetPrincipalEntitlements(ctx context.Context) ApiGetPrincipalEntitlementsRequest {
	return ApiGetPrincipalEntitlementsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return PrincipalEntitlementsList
func (a *PrincipalEntitlementsAPIService) GetPrincipalEntitlementsExecute(r ApiGetPrincipalEntitlementsRequest) (*PrincipalEntitlementsList, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrincipalEntitlementsList
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrincipalEntitlementsAPIService.GetPrincipalEntitlements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/principal-entitlements"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filter == nil {
		return localVarReturnValue, nil, reportError("filter is required and must be specified")
	}

	localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiGetPrincipalEntitlementsChangesRequest struct {
	ctx                           context.Context
	ApiService                    PrincipalEntitlementsAPI
	principalEntitlementsChangeId string
	retryCount                    int32
}

func (r ApiGetPrincipalEntitlementsChangesRequest) Execute() (*PrincipalEntitlementsChange, *APIResponse, error) {
	return r.ApiService.GetPrincipalEntitlementsChangesExecute(r)
}

/*
GetPrincipalEntitlementsChanges Retrieve the principal entitlement changes

Retrieves the principal's old and new effective entitlement changes for a resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param principalEntitlementsChangeId The entitlement change ID, in the format of `{old_effective_grant_id}:{new_effective_grant_id}`.  Only the `{new_effective_grant_id}` value is required if the `{old_effective_grant_id}` value isn't available.
	@return ApiGetPrincipalEntitlementsChangesRequest
*/
func (a *PrincipalEntitlementsAPIService) GetPrincipalEntitlementsChanges(ctx context.Context, principalEntitlementsChangeId string) ApiGetPrincipalEntitlementsChangesRequest {
	return ApiGetPrincipalEntitlementsChangesRequest{
		ApiService:                    a,
		ctx:                           ctx,
		principalEntitlementsChangeId: principalEntitlementsChangeId,
		retryCount:                    0,
	}
}

// Execute executes the request
//
//	@return PrincipalEntitlementsChange
func (a *PrincipalEntitlementsAPIService) GetPrincipalEntitlementsChangesExecute(r ApiGetPrincipalEntitlementsChangesRequest) (*PrincipalEntitlementsChange, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrincipalEntitlementsChange
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrincipalEntitlementsAPIService.GetPrincipalEntitlementsChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/principal-entitlements-changes/{principalEntitlementsChangeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"principalEntitlementsChangeId"+"}", url.PathEscape(parameterToString(r.principalEntitlementsChangeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiGetPrincipalEntitlementsHistoryRequest struct {
	ctx        context.Context
	ApiService PrincipalEntitlementsAPI
	filter     *string
	limit      *int32
	after      *string
	include    *[]string
	retryCount int32
}

// This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;eq&#x60; [operator](https://developer.okta.com/docs/api/#operators) and the following required and optional properties.  **Required**: You must specify both principal and resource references with one of these sets of properties: * &#x60;principal.externalId&#x60;, &#x60;principal.type&#x60;, &#x60;resource.externalId&#x60;, and &#x60;resource.type&#x60; * &#x60;principalOrn&#x60; and &#x60;resourceOrn&#x60; * &#x60;principalId&#x60; and &#x60;resourceId&#x60;  **Optional**: You can optionally filter by a date range with the following properties: * &#x60;startDate&#x60;: Start of the date range (inclusive) in ISO 8601 UTC format. If omitted, data is retrieved from the earliest available records. * &#x60;endDate&#x60;: End of the date range (inclusive) in ISO 8601 UTC format. If omitted, data is retrieved up to the current time.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiGetPrincipalEntitlementsHistoryRequest) Filter(filter string) ApiGetPrincipalEntitlementsHistoryRequest {
	r.filter = &filter
	return r
}

// The maximum number of records returned in a response
func (r ApiGetPrincipalEntitlementsHistoryRequest) Limit(limit int32) ApiGetPrincipalEntitlementsHistoryRequest {
	r.limit = &limit
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiGetPrincipalEntitlementsHistoryRequest) After(after string) ApiGetPrincipalEntitlementsHistoryRequest {
	r.after = &after
	return r
}

// An optional parameter that adds additional properties in the &#x60;metadata&#x60; response object
func (r ApiGetPrincipalEntitlementsHistoryRequest) Include(include []string) ApiGetPrincipalEntitlementsHistoryRequest {
	r.include = &include
	return r
}

func (r ApiGetPrincipalEntitlementsHistoryRequest) Execute() (*PrincipalEntitlementsHistory, *APIResponse, error) {
	return r.ApiService.GetPrincipalEntitlementsHistoryExecute(r)
}

/*
GetPrincipalEntitlementsHistory Retrieve an entitlement history

Retrieves an entitlement history log for a specific principal and resource.

Specify the required principal and resource references in the `filter` query parameter.
You can optionally specify a time range for the returned history log.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPrincipalEntitlementsHistoryRequest
*/
func (a *PrincipalEntitlementsAPIService) GetPrincipalEntitlementsHistory(ctx context.Context) ApiGetPrincipalEntitlementsHistoryRequest {
	return ApiGetPrincipalEntitlementsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return PrincipalEntitlementsHistory
func (a *PrincipalEntitlementsAPIService) GetPrincipalEntitlementsHistoryExecute(r ApiGetPrincipalEntitlementsHistoryRequest) (*PrincipalEntitlementsHistory, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrincipalEntitlementsHistory
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrincipalEntitlementsAPIService.GetPrincipalEntitlementsHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/principal-entitlements/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filter == nil {
		return localVarReturnValue, nil, reportError("filter is required and must be specified")
	}

	localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.include != nil {
		t := *r.include
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("include", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("include", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}
