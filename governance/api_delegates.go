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
	"time"
)

type DelegatesAPI interface {

	/*
			ListDelegateAppointments List all delegate appointments

			Lists all delegate appointments in your org.

		You can filter delegate appointments for a specific user with the `filter` query parameter.
		You can also manage the number of records in the response with the `limit` and `after` query parameters.

		The following are request examples with query parameters:

		1. Lists at most 20 delegate appointments
		    ```
		    /governance/api/v1/delegates?limit=20
		    ```
		1. Lists the next 10 results of delegate appointments after a specific cursor
		    ```
		    /governance/api/v1/delegates?limit=10&after=gda123ABCXYZ456DEF
		    ```
		1. Lists at most 2 delegate appointments with `delegatorId` equals `00ub0oNGTSWTBKOLAAAA` or `00u2lxfQaw8WRlkQt0g4`
		    ```
		    /governance/api/v1/delegates?filter=(delegatorId%20eq%20%2200ub0oNGTSWTBKOLGLNR%22%20OR%20delegatorId%20eq%20%2200ub0oNGTSWTBKOLGLNS%22)&limit=2
		    ```

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiListDelegateAppointmentsRequest
	*/
	ListDelegateAppointments(ctx context.Context) ApiListDelegateAppointmentsRequest

	// ListDelegateAppointmentsExecute executes the request
	//  @return DelegateAppointmentList
	ListDelegateAppointmentsExecute(r ApiListDelegateAppointmentsRequest) (*DelegateAppointmentList, *APIResponse, error)
}

// DelegatesAPIService DelegatesAPI service
type DelegatesAPIService service

type ApiListDelegateAppointmentsRequest struct {
	ctx        context.Context
	ApiService DelegatesAPI
	filter     *string
	limit      *int32
	after      *string
	retryCount int32
}

// Filter expression for retrieving delegates based on specific criteria. Only the &#x60;delegatorId&#x60; property is supported with the &#x60;eq&#x60; operator.
func (r ApiListDelegateAppointmentsRequest) Filter(filter string) ApiListDelegateAppointmentsRequest {
	r.filter = &filter
	return r
}

// The maximum number of records returned in a response
func (r ApiListDelegateAppointmentsRequest) Limit(limit int32) ApiListDelegateAppointmentsRequest {
	r.limit = &limit
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListDelegateAppointmentsRequest) After(after string) ApiListDelegateAppointmentsRequest {
	r.after = &after
	return r
}

func (r ApiListDelegateAppointmentsRequest) Execute() (*DelegateAppointmentList, *APIResponse, error) {
	return r.ApiService.ListDelegateAppointmentsExecute(r)
}

/*
ListDelegateAppointments List all delegate appointments

Lists all delegate appointments in your org.

You can filter delegate appointments for a specific user with the `filter` query parameter.
You can also manage the number of records in the response with the `limit` and `after` query parameters.

The following are request examples with query parameters:

 1. Lists at most 20 delegate appointments
    ```
    /governance/api/v1/delegates?limit=20
    ```

 1. Lists the next 10 results of delegate appointments after a specific cursor
    ```
    /governance/api/v1/delegates?limit=10&after=gda123ABCXYZ456DEF
    ```

 1. Lists at most 2 delegate appointments with `delegatorId` equals `00ub0oNGTSWTBKOLAAAA` or `00u2lxfQaw8WRlkQt0g4`
    ```
    /governance/api/v1/delegates?filter=(delegatorId%20eq%20%2200ub0oNGTSWTBKOLGLNR%22%20OR%20delegatorId%20eq%20%2200ub0oNGTSWTBKOLGLNS%22)&limit=2
    ```

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiListDelegateAppointmentsRequest
*/
func (a *DelegatesAPIService) ListDelegateAppointments(ctx context.Context) ApiListDelegateAppointmentsRequest {
	return ApiListDelegateAppointmentsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return DelegateAppointmentList
func (a *DelegatesAPIService) ListDelegateAppointmentsExecute(r ApiListDelegateAppointmentsRequest) (*DelegateAppointmentList, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DelegateAppointmentList
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DelegatesAPIService.ListDelegateAppointments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/delegates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
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
