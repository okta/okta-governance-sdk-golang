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
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MyCatalogsAPI interface {
	/*
			GetMyCatalogEntryRequestFieldsV2 Retrieve an entry's request fields

			Retrieves request fields for my catalog entry

		Request fields are determined by evaluating the entry's associated
		request conditions for the requester.

		The highest priority matching
		condition determines the approval sequence that will be used for the
		requester.

		If that approval sequence has requester fields, then they will
		be returned as a request field.

		If the request can lead to any separation of duty conflicts, then
		the risk assessment is present. The risk assessment indicates whether
		request submission is allowed or restricted and includes rules that
		lead to the possible conflicts. If request submission is allowed, then
		the request fields are determined by the associated approval sequence.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param entryId The ID of the catalog entry
			@return ApiGetMyCatalogEntryRequestFieldsV2Request
	*/
	GetMyCatalogEntryRequestFieldsV2(ctx context.Context, entryId string) ApiGetMyCatalogEntryRequestFieldsV2Request

	// GetMyCatalogEntryRequestFieldsV2Execute executes the request
	//  @return CatalogEntryRequestFields
	GetMyCatalogEntryRequestFieldsV2Execute(r ApiGetMyCatalogEntryRequestFieldsV2Request) (*CatalogEntryRequestFields, *APIResponse, error)

	/*
			GetMyCatalogEntryUserRequestFieldsV2 Retrieve a users request-fields for an entry

			Retrieves the entry's request fields for the specified requester.
		Request fields for the entry are only returned if the entry has request on behalf of enabled, and the authorized user is able to request on behalf of other requesters.

		If the request can lead to any separation of duty conflicts, then
		the risk assessment is present. The risk assessment indicates whether
		request submission is allowed or restricted and includes rules that
		lead to the possible conflicts. If request submission is allowed, then
		the request fields are determined by the associated approval sequence
		for the risk level.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param entryId The ID of the catalog entry
			@param userId The `id` of the user
			@return ApiGetMyCatalogEntryUserRequestFieldsV2Request
	*/
	GetMyCatalogEntryUserRequestFieldsV2(ctx context.Context, entryId string, userId string) ApiGetMyCatalogEntryUserRequestFieldsV2Request

	// GetMyCatalogEntryUserRequestFieldsV2Execute executes the request
	//  @return CatalogEntryRequestFields
	GetMyCatalogEntryUserRequestFieldsV2Execute(r ApiGetMyCatalogEntryUserRequestFieldsV2Request) (*CatalogEntryRequestFields, *APIResponse, error)

	/*
		GetMyEntryV2 Retrieve an entry from my catalog

		Retrieves an entry from my catalog

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param entryId The ID of the catalog entry
		@return ApiGetMyEntryV2Request
	*/
	GetMyEntryV2(ctx context.Context, entryId string) ApiGetMyEntryV2Request

	// GetMyEntryV2Execute executes the request
	//  @return RcarEntryGet
	GetMyEntryV2Execute(r ApiGetMyEntryV2Request) (*RcarEntryGet, *APIResponse, error)

	/*
			ListMyDefaultEntriesV2 List all of my entries for the default access request catalog

			Lists filtered entries for the default access request catalog that you're allowed to request (as the authenticated requestor).

		The following are request examples with query parameters:

		1. Lists at most 20 parent (top-level) entries
		    ```
		    /governance/api/v2/catalogs/default/entries?filter=not(parent%20pr)&limit=20
		    ```
		1. Lists the next 20 results of parent entries after a specific cursor
		    ```
		    /governance/api/v2/catalogs/default/entries?filter=not(parent%20pr)&limit=20&after=cen33e47frfMB93gQ8g6
		    ```
		1. Lists at most 8 parent entries with a fuzzy match for "figma"
		    ```
		    /governance/api/v2/catalogs/default/entries?filter=not(parent%20pr)&match=figma&limit=8
		    ```
		1. Lists at most 8 child entries with a specific parent
		    ```
		    /governance/api/v2/catalogs/default/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&limit=8
		    ```
		1. Lists at most 8 child entries that have "edit" in the name and have a specific parent
		    ```
		    /governance/api/v2/catalogs/default/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&match=edit&limit=8
		    ```

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiListMyDefaultEntriesV2Request
	*/
	ListMyDefaultEntriesV2(ctx context.Context) ApiListMyDefaultEntriesV2Request

	// ListMyDefaultEntriesV2Execute executes the request
	//  @return RcarEntriesListV2
	ListMyDefaultEntriesV2Execute(r ApiListMyDefaultEntriesV2Request) (*RcarEntriesListV2, *APIResponse, error)

	/*
			ListMyEntryUsersV2 List all of my catalog entry users

			Lists all users who match the filtered query and can also view and request the entry.

		A list of users is only returned if the entry has the `requestOnBehalfOfSettings` enabled, a filter is specified, and the authorized user is able to request on behalf of other users.

		**Examples**

		Request examples with query parameters:

		1. Filter users with a last name that starts with "Smi"
		    ```
		    /governance/api/v2/my/catalogs/default/entries/{entryId}/users?filter=lastName%20sw%20%22Smi%22
		    ```
		1. Filter for users with a first name that begins with "John"
		    ```
		    /governance/api/v2/my/catalogs/default/entries/{entryId}/users?filter=firstName%20sw%20%22John%22
		    ```
		1. Search for users with a first or last name that begins with "John"
		    ```
		    /governance/api/v2/my/catalogs/default/entries/{entryId}/users?filter=firstName%20sw%20%22John%22%20OR%20lastName%20sw%20%22John%22
		    ```


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param entryId The ID of the catalog entry
			@return ApiListMyEntryUsersV2Request
	*/
	ListMyEntryUsersV2(ctx context.Context, entryId string) ApiListMyEntryUsersV2Request

	// ListMyEntryUsersV2Execute executes the request
	//  @return EntryUsersList
	ListMyEntryUsersV2Execute(r ApiListMyEntryUsersV2Request) (*EntryUsersList, *APIResponse, error)
}

// MyCatalogsAPIService MyCatalogsAPI service
type MyCatalogsAPIService service

type ApiGetMyCatalogEntryRequestFieldsV2Request struct {
	ctx        context.Context
	ApiService MyCatalogsAPI
	entryId    string
	retryCount int32
}

func (r ApiGetMyCatalogEntryRequestFieldsV2Request) Execute() (*CatalogEntryRequestFields, *APIResponse, error) {
	return r.ApiService.GetMyCatalogEntryRequestFieldsV2Execute(r)
}

/*
GetMyCatalogEntryRequestFieldsV2 Retrieve an entry's request fields

# Retrieves request fields for my catalog entry

Request fields are determined by evaluating the entry's associated
request conditions for the requester.

The highest priority matching
condition determines the approval sequence that will be used for the
requester.

If that approval sequence has requester fields, then they will
be returned as a request field.

If the request can lead to any separation of duty conflicts, then
the risk assessment is present. The risk assessment indicates whether
request submission is allowed or restricted and includes rules that
lead to the possible conflicts. If request submission is allowed, then
the request fields are determined by the associated approval sequence.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryId The ID of the catalog entry
	@return ApiGetMyCatalogEntryRequestFieldsV2Request
*/
func (a *MyCatalogsAPIService) GetMyCatalogEntryRequestFieldsV2(ctx context.Context, entryId string) ApiGetMyCatalogEntryRequestFieldsV2Request {
	return ApiGetMyCatalogEntryRequestFieldsV2Request{
		ApiService: a,
		ctx:        ctx,
		entryId:    entryId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return CatalogEntryRequestFields
func (a *MyCatalogsAPIService) GetMyCatalogEntryRequestFieldsV2Execute(r ApiGetMyCatalogEntryRequestFieldsV2Request) (*CatalogEntryRequestFields, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CatalogEntryRequestFields
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyCatalogsAPIService.GetMyCatalogEntryRequestFieldsV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/catalogs/default/entries/{entryId}/request-fields"
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetMyCatalogEntryUserRequestFieldsV2Request struct {
	ctx        context.Context
	ApiService MyCatalogsAPI
	entryId    string
	userId     string
	retryCount int32
}

func (r ApiGetMyCatalogEntryUserRequestFieldsV2Request) Execute() (*CatalogEntryRequestFields, *APIResponse, error) {
	return r.ApiService.GetMyCatalogEntryUserRequestFieldsV2Execute(r)
}

/*
GetMyCatalogEntryUserRequestFieldsV2 Retrieve a users request-fields for an entry

Retrieves the entry's request fields for the specified requester.
Request fields for the entry are only returned if the entry has request on behalf of enabled, and the authorized user is able to request on behalf of other requesters.

If the request can lead to any separation of duty conflicts, then
the risk assessment is present. The risk assessment indicates whether
request submission is allowed or restricted and includes rules that
lead to the possible conflicts. If request submission is allowed, then
the request fields are determined by the associated approval sequence
for the risk level.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryId The ID of the catalog entry
	@param userId The `id` of the user
	@return ApiGetMyCatalogEntryUserRequestFieldsV2Request
*/
func (a *MyCatalogsAPIService) GetMyCatalogEntryUserRequestFieldsV2(ctx context.Context, entryId string, userId string) ApiGetMyCatalogEntryUserRequestFieldsV2Request {
	return ApiGetMyCatalogEntryUserRequestFieldsV2Request{
		ApiService: a,
		ctx:        ctx,
		entryId:    entryId,
		userId:     userId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return CatalogEntryRequestFields
func (a *MyCatalogsAPIService) GetMyCatalogEntryUserRequestFieldsV2Execute(r ApiGetMyCatalogEntryUserRequestFieldsV2Request) (*CatalogEntryRequestFields, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CatalogEntryRequestFields
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyCatalogsAPIService.GetMyCatalogEntryUserRequestFieldsV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/catalogs/default/entries/{entryId}/users/{userId}/request-fields"
	localVarPath = strings.Replace(localVarPath, "{"+"entryId"+"}", url.PathEscape(parameterToString(r.entryId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.entryId) < 20 {
		return localVarReturnValue, nil, reportError("entryId must have at least 20 elements")
	}
	if strlen(r.entryId) > 20 {
		return localVarReturnValue, nil, reportError("entryId must have less than 20 elements")
	}
	if strlen(r.userId) < 20 {
		return localVarReturnValue, nil, reportError("userId must have at least 20 elements")
	}
	if strlen(r.userId) > 20 {
		return localVarReturnValue, nil, reportError("userId must have less than 20 elements")
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetMyEntryV2Request struct {
	ctx        context.Context
	ApiService MyCatalogsAPI
	entryId    string
	retryCount int32
}

func (r ApiGetMyEntryV2Request) Execute() (*RcarEntryGet, *APIResponse, error) {
	return r.ApiService.GetMyEntryV2Execute(r)
}

/*
GetMyEntryV2 Retrieve an entry from my catalog

Retrieves an entry from my catalog

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryId The ID of the catalog entry
	@return ApiGetMyEntryV2Request
*/
func (a *MyCatalogsAPIService) GetMyEntryV2(ctx context.Context, entryId string) ApiGetMyEntryV2Request {
	return ApiGetMyEntryV2Request{
		ApiService: a,
		ctx:        ctx,
		entryId:    entryId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RcarEntryGet
func (a *MyCatalogsAPIService) GetMyEntryV2Execute(r ApiGetMyEntryV2Request) (*RcarEntryGet, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RcarEntryGet
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyCatalogsAPIService.GetMyEntryV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/catalogs/default/entries/{entryId}"
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListMyDefaultEntriesV2Request struct {
	ctx        context.Context
	ApiService MyCatalogsAPI
	filter     *string
	after      *string
	match      *string
	limit      *int32
	retryCount int32
}

// A required filter expression that returns entries based on the [&#x60;parent&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/parent&amp;t&#x3D;response) property. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;eq&#x60; and &#x60;pr&#x60; [operators](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).
func (r ApiListMyDefaultEntriesV2Request) Filter(filter string) ApiListMyDefaultEntriesV2Request {
	r.filter = &filter
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [&#x60;limit&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in&#x3D;query&amp;path&#x3D;limit&amp;t&#x3D;request) query parameter. If there are more entries to return, the &#x60;_links.next.href&#x60; link contains the &#x60;after&#x60; cursor for the next page of results.
func (r ApiListMyDefaultEntriesV2Request) After(after string) ApiListMyDefaultEntriesV2Request {
	r.after = &after
	return r
}

// Return catalog entries that match a substring value in the [&#x60;name&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/name&amp;t&#x3D;response) or [&#x60;description&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/description&amp;t&#x3D;response) properties. At least three characters are required for fuzzy search.
func (r ApiListMyDefaultEntriesV2Request) Match(match string) ApiListMyDefaultEntriesV2Request {
	r.match = &match
	return r
}

// The maximum number of records returned in a response
func (r ApiListMyDefaultEntriesV2Request) Limit(limit int32) ApiListMyDefaultEntriesV2Request {
	r.limit = &limit
	return r
}

func (r ApiListMyDefaultEntriesV2Request) Execute() (*RcarEntriesListV2, *APIResponse, error) {
	return r.ApiService.ListMyDefaultEntriesV2Execute(r)
}

/*
ListMyDefaultEntriesV2 List all of my entries for the default access request catalog

Lists filtered entries for the default access request catalog that you're allowed to request (as the authenticated requestor).

The following are request examples with query parameters:

 1. Lists at most 20 parent (top-level) entries
    ```
    /governance/api/v2/catalogs/default/entries?filter=not(parent%20pr)&limit=20
    ```

 1. Lists the next 20 results of parent entries after a specific cursor
    ```
    /governance/api/v2/catalogs/default/entries?filter=not(parent%20pr)&limit=20&after=cen33e47frfMB93gQ8g6
    ```

 1. Lists at most 8 parent entries with a fuzzy match for "figma"
    ```
    /governance/api/v2/catalogs/default/entries?filter=not(parent%20pr)&match=figma&limit=8
    ```

 1. Lists at most 8 child entries with a specific parent
    ```
    /governance/api/v2/catalogs/default/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&limit=8
    ```

 1. Lists at most 8 child entries that have "edit" in the name and have a specific parent
    ```
    /governance/api/v2/catalogs/default/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&match=edit&limit=8
    ```

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiListMyDefaultEntriesV2Request
*/
func (a *MyCatalogsAPIService) ListMyDefaultEntriesV2(ctx context.Context) ApiListMyDefaultEntriesV2Request {
	return ApiListMyDefaultEntriesV2Request{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RcarEntriesListV2
func (a *MyCatalogsAPIService) ListMyDefaultEntriesV2Execute(r ApiListMyDefaultEntriesV2Request) (*RcarEntriesListV2, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RcarEntriesListV2
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyCatalogsAPIService.ListMyDefaultEntriesV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/catalogs/default/entries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filter == nil {
		return localVarReturnValue, nil, reportError("filter is required and must be specified")
	}

	localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.match != nil {
		localVarQueryParams.Add("match", parameterToString(*r.match, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListMyEntryUsersV2Request struct {
	ctx        context.Context
	ApiService MyCatalogsAPI
	entryId    string
	filter     *string
	after      *string
	limit      *int32
	retryCount int32
}

// A required filter expression that returns users based on the &#x60;firstName&#x60; or &#x60;lastName&#x60; properties. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).
func (r ApiListMyEntryUsersV2Request) Filter(filter string) ApiListMyEntryUsersV2Request {
	r.filter = &filter
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListMyEntryUsersV2Request) After(after string) ApiListMyEntryUsersV2Request {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListMyEntryUsersV2Request) Limit(limit int32) ApiListMyEntryUsersV2Request {
	r.limit = &limit
	return r
}

func (r ApiListMyEntryUsersV2Request) Execute() (*EntryUsersList, *APIResponse, error) {
	return r.ApiService.ListMyEntryUsersV2Execute(r)
}

/*
ListMyEntryUsersV2 List all of my catalog entry users

Lists all users who match the filtered query and can also view and request the entry.

A list of users is only returned if the entry has the `requestOnBehalfOfSettings` enabled, a filter is specified, and the authorized user is able to request on behalf of other users.

**Examples**

Request examples with query parameters:

 1. Filter users with a last name that starts with "Smi"
    ```
    /governance/api/v2/my/catalogs/default/entries/{entryId}/users?filter=lastName%20sw%20%22Smi%22
    ```

 1. Filter for users with a first name that begins with "John"
    ```
    /governance/api/v2/my/catalogs/default/entries/{entryId}/users?filter=firstName%20sw%20%22John%22
    ```

 1. Search for users with a first or last name that begins with "John"
    ```
    /governance/api/v2/my/catalogs/default/entries/{entryId}/users?filter=firstName%20sw%20%22John%22%20OR%20lastName%20sw%20%22John%22
    ```

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @param entryId The ID of the catalog entry
    @return ApiListMyEntryUsersV2Request
*/
func (a *MyCatalogsAPIService) ListMyEntryUsersV2(ctx context.Context, entryId string) ApiListMyEntryUsersV2Request {
	return ApiListMyEntryUsersV2Request{
		ApiService: a,
		ctx:        ctx,
		entryId:    entryId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return EntryUsersList
func (a *MyCatalogsAPIService) ListMyEntryUsersV2Execute(r ApiListMyEntryUsersV2Request) (*EntryUsersList, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntryUsersList
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyCatalogsAPIService.ListMyEntryUsersV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/catalogs/default/entries/{entryId}/users"
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
	if r.filter == nil {
		return localVarReturnValue, nil, reportError("filter is required and must be specified")
	}

	localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
