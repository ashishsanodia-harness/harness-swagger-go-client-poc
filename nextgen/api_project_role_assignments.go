
/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ProjectRoleAssignmentsApiService service
/*
ProjectRoleAssignmentsApiService Create a role assignment
Create a role assignment
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body Role Request body
 * @param org Organization identifier
 * @param project Project identifier
 * @param optional nil or *ProjectRoleAssignmentsApiCreateProjectScopedRoleAssignmentsOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return RoleAssignmentResponse
*/

type ProjectRoleAssignmentsApiCreateProjectScopedRoleAssignmentsOpts struct {
    HarnessAccount optional.String
}

func (a *ProjectRoleAssignmentsApiService) CreateProjectScopedRoleAssignments(ctx context.Context, body RoleAssignment, org string, project string, localVarOptionals *ProjectRoleAssignmentsApiCreateProjectScopedRoleAssignmentsOpts) (RoleAssignmentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RoleAssignmentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/roleassignments"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v RoleAssignmentResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ProjectRoleAssignmentsApiService Delete a role assignment
Deletes the information of the role assignment with the matching role assignment slug.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param roleassignment Role assignment slug
 * @param org Organization identifier
 * @param project Project identifier
 * @param optional nil or *ProjectRoleAssignmentsApiDeleteProjectScopedRoleAssignmentOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return RoleAssignmentResponse
*/

type ProjectRoleAssignmentsApiDeleteProjectScopedRoleAssignmentOpts struct {
    HarnessAccount optional.String
}

func (a *ProjectRoleAssignmentsApiService) DeleteProjectScopedRoleAssignment(ctx context.Context, roleassignment string, org string, project string, localVarOptionals *ProjectRoleAssignmentsApiDeleteProjectScopedRoleAssignmentOpts) (RoleAssignmentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RoleAssignmentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/roleassignments/{roleassignment}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleassignment"+"}", fmt.Sprintf("%v", roleassignment), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v RoleAssignmentResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ProjectRoleAssignmentsApiService Retrieve a role assignment
Retrieves the information of the role assignment with the matching role assignment slug.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param roleassignment Role assignment slug
 * @param org Organization identifier
 * @param project Project identifier
 * @param optional nil or *ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return RoleAssignmentResponse
*/

type ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentOpts struct {
    HarnessAccount optional.String
}

func (a *ProjectRoleAssignmentsApiService) GetProjectScopedRoleAssignment(ctx context.Context, roleassignment string, org string, project string, localVarOptionals *ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentOpts) (RoleAssignmentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RoleAssignmentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/roleassignments/{roleassignment}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleassignment"+"}", fmt.Sprintf("%v", roleassignment), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v RoleAssignmentResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ProjectRoleAssignmentsApiService List role assignments
Retrieves the information of the role assignments
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Organization identifier
 * @param project Project identifier
 * @param optional nil or *ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentsOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "Page" (optional.Int32) -  Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page.
     * @param "Limit" (optional.Int32) -  Pagination: Number of items to return.
     * @param "Sort" (optional.String) -  Parameter on the basis of which sorting is done.
     * @param "Order" (optional.String) -  Order on the basis of which sorting is done.
@return []RoleAssignmentResponse
*/

type ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentsOpts struct {
    HarnessAccount optional.String
    Page optional.Int32
    Limit optional.Int32
    Sort optional.String
    Order optional.String
}

func (a *ProjectRoleAssignmentsApiService) GetProjectScopedRoleAssignments(ctx context.Context, org string, project string, localVarOptionals *ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentsOpts) ([]RoleAssignmentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []RoleAssignmentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/roleassignments"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []RoleAssignmentResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
