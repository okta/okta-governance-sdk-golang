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

type CatalogsAPI interface {
	/*
			GetCatalogEntryRequestFieldsV2 Retrieve an entry's request fields

			Retrieves request fields for catalog entry

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
			@param userId The `id` of the user
			@return ApiGetCatalogEntryRequestFieldsV2Request
	*/
	GetCatalogEntryRequestFieldsV2(ctx context.Context, entryId string, userId string) ApiGetCatalogEntryRequestFieldsV2Request

	// GetCatalogEntryRequestFieldsV2Execute executes the request
	//  @return CatalogEntryRequestFields
	GetCatalogEntryRequestFieldsV2Execute(r ApiGetCatalogEntryRequestFieldsV2Request) (*CatalogEntryRequestFields, *APIResponse, error)

	/*
		GetCatalogEntryV2 Retrieve a catalog entry

		Retrieves a catalog entry

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param entryId The ID of the catalog entry
		@return ApiGetCatalogEntryV2Request
	*/
	GetCatalogEntryV2(ctx context.Context, entryId string) ApiGetCatalogEntryV2Request

	// GetCatalogEntryV2Execute executes the request
	//  @return RcarEntryGet
	GetCatalogEntryV2Execute(r ApiGetCatalogEntryV2Request) (*RcarEntryGet, *APIResponse, error)

	/*
			ListAllDefaultEntriesV2 List all entries for the default access request catalog

			Lists entries for the default access request catalog based on a filter.

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
			@return ApiListAllDefaultEntriesV2Request
	*/
	ListAllDefaultEntriesV2(ctx context.Context) ApiListAllDefaultEntriesV2Request

	// ListAllDefaultEntriesV2Execute executes the request
	//  @return RcarEntriesListV2
	ListAllDefaultEntriesV2Execute(r ApiListAllDefaultEntriesV2Request) (*RcarEntriesListV2, *APIResponse, error)

	/*
			ListAllDefaultUserEntriesV2 List all access request catalog entries for a user

			Lists access request catalog entries for a user based on a filter (the `filter` query parameter is required)

		The following are request examples with query parameters:

		1. Lists at most 20 parent (top-level) entries
		    ```
		    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=not(parent%20pr)&limit=20
		    ```
		1. Lists the next 20 results of parent entries after a specific cursor
		    ```
		    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=not(parent%20pr)&limit=20&after=cen33e47frfMB93gQ8g6
		    ```
		1. Lists at most 8 parent entries with a string match for "figma"
		    ```
		    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=not(parent%20pr)&match=figma&limit=8
		    ```
		1. Lists at most 8 child entries with a specific parent
		    ```
		    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&limit=8
		    ```
		1. Lists at most 8 child entries that have "edit" in the name and have a specific parent
		    ```
		    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&match=edit&limit=8
		    ```

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId The `id` of the user
			@return ApiListAllDefaultUserEntriesV2Request
	*/
	ListAllDefaultUserEntriesV2(ctx context.Context, userId string) ApiListAllDefaultUserEntriesV2Request

	// ListAllDefaultUserEntriesV2Execute executes the request
	//  @return RcarEntriesListV2
	ListAllDefaultUserEntriesV2Execute(r ApiListAllDefaultUserEntriesV2Request) (*RcarEntriesListV2, *APIResponse, error)
}

// CatalogsAPIService CatalogsAPI service
type CatalogsAPIService service

type ApiGetCatalogEntryRequestFieldsV2Request struct {
	ctx        context.Context
	ApiService CatalogsAPI
	entryId    string
	userId     string
	retryCount int32
}

func (r ApiGetCatalogEntryRequestFieldsV2Request) Execute() (*CatalogEntryRequestFields, *APIResponse, error) {
	return r.ApiService.GetCatalogEntryRequestFieldsV2Execute(r)
}

/*
GetCatalogEntryRequestFieldsV2 Retrieve an entry's request fields

# Retrieves request fields for catalog entry

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
	@param userId The `id` of the user
	@return ApiGetCatalogEntryRequestFieldsV2Request
*/
func (a *CatalogsAPIService) GetCatalogEntryRequestFieldsV2(ctx context.Context, entryId string, userId string) ApiGetCatalogEntryRequestFieldsV2Request {
	return ApiGetCatalogEntryRequestFieldsV2Request{
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
func (a *CatalogsAPIService) GetCatalogEntryRequestFieldsV2Execute(r ApiGetCatalogEntryRequestFieldsV2Request) (*CatalogEntryRequestFields, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogsAPIService.GetCatalogEntryRequestFieldsV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/catalogs/default/entries/{entryId}/users/{userId}/request-fields"
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

type ApiGetCatalogEntryV2Request struct {
	ctx        context.Context
	ApiService CatalogsAPI
	entryId    string
	retryCount int32
}

func (r ApiGetCatalogEntryV2Request) Execute() (*RcarEntryGet, *APIResponse, error) {
	return r.ApiService.GetCatalogEntryV2Execute(r)
}

/*
GetCatalogEntryV2 Retrieve a catalog entry

Retrieves a catalog entry

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param entryId The ID of the catalog entry
	@return ApiGetCatalogEntryV2Request
*/
func (a *CatalogsAPIService) GetCatalogEntryV2(ctx context.Context, entryId string) ApiGetCatalogEntryV2Request {
	return ApiGetCatalogEntryV2Request{
		ApiService: a,
		ctx:        ctx,
		entryId:    entryId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RcarEntryGet
func (a *CatalogsAPIService) GetCatalogEntryV2Execute(r ApiGetCatalogEntryV2Request) (*RcarEntryGet, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogsAPIService.GetCatalogEntryV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/catalogs/default/entries/{entryId}"
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

type ApiListAllDefaultEntriesV2Request struct {
	ctx        context.Context
	ApiService CatalogsAPI
	filter     *string
	after      *string
	match      *string
	limit      *int32
	retryCount int32
}

// A required filter expression that returns entries based on the [&#x60;parent&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/parent&amp;t&#x3D;response) property. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;eq&#x60; and &#x60;pr&#x60; [operators](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).
func (r ApiListAllDefaultEntriesV2Request) Filter(filter string) ApiListAllDefaultEntriesV2Request {
	r.filter = &filter
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [&#x60;limit&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in&#x3D;query&amp;path&#x3D;limit&amp;t&#x3D;request) query parameter. If there are more entries to return, the &#x60;_links.next.href&#x60; link contains the &#x60;after&#x60; cursor for the next page of results.
func (r ApiListAllDefaultEntriesV2Request) After(after string) ApiListAllDefaultEntriesV2Request {
	r.after = &after
	return r
}

// Return catalog entries that match a substring value in the [&#x60;name&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/name&amp;t&#x3D;response) or [&#x60;description&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/description&amp;t&#x3D;response) properties. At least three characters are required for fuzzy search.
func (r ApiListAllDefaultEntriesV2Request) Match(match string) ApiListAllDefaultEntriesV2Request {
	r.match = &match
	return r
}

// The maximum number of records returned in a response
func (r ApiListAllDefaultEntriesV2Request) Limit(limit int32) ApiListAllDefaultEntriesV2Request {
	r.limit = &limit
	return r
}

func (r ApiListAllDefaultEntriesV2Request) Execute() (*RcarEntriesListV2, *APIResponse, error) {
	return r.ApiService.ListAllDefaultEntriesV2Execute(r)
}

/*
ListAllDefaultEntriesV2 List all entries for the default access request catalog

Lists entries for the default access request catalog based on a filter.

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
    @return ApiListAllDefaultEntriesV2Request
*/
func (a *CatalogsAPIService) ListAllDefaultEntriesV2(ctx context.Context) ApiListAllDefaultEntriesV2Request {
	return ApiListAllDefaultEntriesV2Request{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RcarEntriesListV2
func (a *CatalogsAPIService) ListAllDefaultEntriesV2Execute(r ApiListAllDefaultEntriesV2Request) (*RcarEntriesListV2, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogsAPIService.ListAllDefaultEntriesV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/catalogs/default/entries"

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

type ApiListAllDefaultUserEntriesV2Request struct {
	ctx        context.Context
	ApiService CatalogsAPI
	userId     string
	filter     *string
	after      *string
	match      *string
	limit      *int32
	retryCount int32
}

// A required filter expression that returns entries based on the [&#x60;parent&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/parent&amp;t&#x3D;response) property. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;eq&#x60; and &#x60;pr&#x60; [operators](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).
func (r ApiListAllDefaultUserEntriesV2Request) Filter(filter string) ApiListAllDefaultUserEntriesV2Request {
	r.filter = &filter
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [&#x60;limit&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in&#x3D;query&amp;path&#x3D;limit&amp;t&#x3D;request) query parameter. If there are more entries to return, the &#x60;_links.next.href&#x60; link contains the &#x60;after&#x60; cursor for the next page of results.
func (r ApiListAllDefaultUserEntriesV2Request) After(after string) ApiListAllDefaultUserEntriesV2Request {
	r.after = &after
	return r
}

// Return catalog entries that match a substring value in the [&#x60;name&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/name&amp;t&#x3D;response) or [&#x60;description&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/description&amp;t&#x3D;response) properties. At least three characters are required for fuzzy search.
func (r ApiListAllDefaultUserEntriesV2Request) Match(match string) ApiListAllDefaultUserEntriesV2Request {
	r.match = &match
	return r
}

// The maximum number of records returned in a response
func (r ApiListAllDefaultUserEntriesV2Request) Limit(limit int32) ApiListAllDefaultUserEntriesV2Request {
	r.limit = &limit
	return r
}

func (r ApiListAllDefaultUserEntriesV2Request) Execute() (*RcarEntriesListV2, *APIResponse, error) {
	return r.ApiService.ListAllDefaultUserEntriesV2Execute(r)
}

/*
ListAllDefaultUserEntriesV2 List all access request catalog entries for a user

Lists access request catalog entries for a user based on a filter (the `filter` query parameter is required)

The following are request examples with query parameters:

 1. Lists at most 20 parent (top-level) entries
    ```
    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=not(parent%20pr)&limit=20
    ```

 1. Lists the next 20 results of parent entries after a specific cursor
    ```
    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=not(parent%20pr)&limit=20&after=cen33e47frfMB93gQ8g6
    ```

 1. Lists at most 8 parent entries with a string match for "figma"
    ```
    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=not(parent%20pr)&match=figma&limit=8
    ```

 1. Lists at most 8 child entries with a specific parent
    ```
    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&limit=8
    ```

 1. Lists at most 8 child entries that have "edit" in the name and have a specific parent
    ```
    /governance/api/v2/catalogs/default/user/{userId}/entries?filter=parent%20eq%20%22cen385AlcdqGaY8HE0g2%22&match=edit&limit=8
    ```

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @param userId The `id` of the user
    @return ApiListAllDefaultUserEntriesV2Request
*/
func (a *CatalogsAPIService) ListAllDefaultUserEntriesV2(ctx context.Context, userId string) ApiListAllDefaultUserEntriesV2Request {
	return ApiListAllDefaultUserEntriesV2Request{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RcarEntriesListV2
func (a *CatalogsAPIService) ListAllDefaultUserEntriesV2Execute(r ApiListAllDefaultUserEntriesV2Request) (*RcarEntriesListV2, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogsAPIService.ListAllDefaultUserEntriesV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/catalogs/default/user/{userId}/entries"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.userId) < 20 {
		return localVarReturnValue, nil, reportError("userId must have at least 20 elements")
	}
	if strlen(r.userId) > 20 {
		return localVarReturnValue, nil, reportError("userId must have less than 20 elements")
	}
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
