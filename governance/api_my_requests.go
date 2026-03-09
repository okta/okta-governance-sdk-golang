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
	"strings"
	"time"
)

type MyRequestsAPI interface {

	/*
		CreateMyRequestV2 Create a request

		Creates a request for my catalog entry specified by `entryId`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param entryId The ID of the catalog entry
		@return ApiCreateMyRequestV2Request
	*/
	CreateMyRequestV2(ctx context.Context, entryId string) ApiCreateMyRequestV2Request

	// CreateMyRequestV2Execute executes the request
	//  @return RequestFull2
	CreateMyRequestV2Execute(r ApiCreateMyRequestV2Request) (*RequestFull2, *APIResponse, error)

	/*
		GetMyRequestV2 Retrieve my request

		Retrieves a request belonging to the authenticated requester

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param entryId The ID of the catalog entry
		@param requestId The `id` of the request
		@return ApiGetMyRequestV2Request
	*/
	GetMyRequestV2(ctx context.Context, entryId string, requestId string) ApiGetMyRequestV2Request

	// GetMyRequestV2Execute executes the request
	//  @return RequestFull2
	GetMyRequestV2Execute(r ApiGetMyRequestV2Request) (*RequestFull2, *APIResponse, error)
}

// MyRequestsAPIService MyRequestsAPI service
type MyRequestsAPIService service

type ApiCreateMyRequestV2Request struct {
	ctx                context.Context
	ApiService         MyRequestsAPI
	entryId            string
	myRequestCreatable *MyRequestCreatable
	retryCount         int32
}

// Creates a resource access request for a given user.  Use this operation to create access requests managed by access request conditions.  If &#x60;requestedBy&#x60; and &#x60;requestedFor&#x60; aren&#39;t the same, then you must also enable the [&#x60;requestOnBehalfOfSettings&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Request-Settings/#tag/Request-Settings/operation/updateResourceRequestSettingsV2!path&#x3D;requestOnBehalfOfSettings&amp;t&#x3D;request) parameter in the access request settings. See [Update the resource request settings](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Request-Settings/#tag/Request-Settings/operation/updateResourceRequestSettingsV2).  Include the following information in the payload:  - The Okta user ID for the user who requires access. Add the user ID in the &#x60;requestedFor.externalId&#x60; parameter. - The catalog entry ID of the resource required by the user. Add the catalog ID in the &#x60;requested.entryId&#x60; parameter. - If the request conditions include requester input fields, add field information in the &#x60;requesterFieldValues&#x60; array. See [Retrieve the request fields](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/getCatalogEntryRequestFieldsV2). - Optional: The user ID of the person submitting the request. By default, this value is the admin user ID requesting this operation and doesn&#39;t need to be provided. However, to add a different Okta user ID for the request, include the &#x60;requestedBy.externalId&#x60; parameter in the request body.
func (r ApiCreateMyRequestV2Request) MyRequestCreatable(myRequestCreatable MyRequestCreatable) ApiCreateMyRequestV2Request {
	r.myRequestCreatable = &myRequestCreatable
	return r
}

func (r ApiCreateMyRequestV2Request) Execute() (*RequestFull2, *APIResponse, error) {
	return r.ApiService.CreateMyRequestV2Execute(r)
}

/*
CreateMyRequestV2 Create a request

Creates a request for my catalog entry specified by `entryId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryId The ID of the catalog entry
	@return ApiCreateMyRequestV2Request
*/
func (a *MyRequestsAPIService) CreateMyRequestV2(ctx context.Context, entryId string) ApiCreateMyRequestV2Request {
	return ApiCreateMyRequestV2Request{
		ApiService: a,
		ctx:        ctx,
		entryId:    entryId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RequestFull2
func (a *MyRequestsAPIService) CreateMyRequestV2Execute(r ApiCreateMyRequestV2Request) (*RequestFull2, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RequestFull2
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyRequestsAPIService.CreateMyRequestV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/catalogs/default/entries/{entryId}/requests"
	localVarPath = strings.Replace(localVarPath, "{"+"entryId"+"}", url.PathEscape(parameterToString(r.entryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.entryId) < 20 {
		return localVarReturnValue, nil, reportError("entryId must have at least 20 elements")
	}
	if strlen(r.entryId) > 20 {
		return localVarReturnValue, nil, reportError("entryId must have less than 20 elements")
	}
	if r.myRequestCreatable == nil {
		return localVarReturnValue, nil, reportError("myRequestCreatable is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.myRequestCreatable
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetMyRequestV2Request struct {
	ctx        context.Context
	ApiService MyRequestsAPI
	entryId    string
	requestId  string
	retryCount int32
}

func (r ApiGetMyRequestV2Request) Execute() (*RequestFull2, *APIResponse, error) {
	return r.ApiService.GetMyRequestV2Execute(r)
}

/*
GetMyRequestV2 Retrieve my request

Retrieves a request belonging to the authenticated requester

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryId The ID of the catalog entry
	@param requestId The `id` of the request
	@return ApiGetMyRequestV2Request
*/
func (a *MyRequestsAPIService) GetMyRequestV2(ctx context.Context, entryId string, requestId string) ApiGetMyRequestV2Request {
	return ApiGetMyRequestV2Request{
		ApiService: a,
		ctx:        ctx,
		entryId:    entryId,
		requestId:  requestId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RequestFull2
func (a *MyRequestsAPIService) GetMyRequestV2Execute(r ApiGetMyRequestV2Request) (*RequestFull2, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RequestFull2
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyRequestsAPIService.GetMyRequestV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/catalogs/default/entries/{entryId}/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"entryId"+"}", url.PathEscape(parameterToString(r.entryId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterToString(r.requestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.entryId) < 20 {
		return localVarReturnValue, nil, reportError("entryId must have at least 20 elements")
	}
	if strlen(r.entryId) > 20 {
		return localVarReturnValue, nil, reportError("entryId must have less than 20 elements")
	}
	if strlen(r.requestId) < 20 {
		return localVarReturnValue, nil, reportError("requestId must have at least 20 elements")
	}
	if strlen(r.requestId) > 20 {
		return localVarReturnValue, nil, reportError("requestId must have less than 20 elements")
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
		if localVarHTTPResponse.StatusCode == 404 {
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
