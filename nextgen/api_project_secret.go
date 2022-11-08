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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ProjectSecretApiService service

/*
ProjectSecretApiService Create a secret
Creates a new secret
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param org Slug field of the organization the resource is scoped to
 * @param project Slug field of the project the resource is scoped to
 * @param optional nil or *ProjectSecretApiCreateProjectScopedSecretOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "PrivateSecret" (optional.Bool) -  This would be used to define secret as private.
@return SecretResponse
*/

type ProjectSecretApiCreateProjectScopedSecretOpts struct {
	HarnessAccount optional.String
	PrivateSecret  optional.Bool
}

func (a *ProjectSecretApiService) CreateProjectScopedSecret(ctx context.Context, body SecretRequest, org string, project string, localVarOptionals *ProjectSecretApiCreateProjectScopedSecretOpts) (SecretResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SecretResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/secrets"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.PrivateSecret.IsSet() {
		localVarQueryParams.Add("private_secret", parameterToString(localVarOptionals.PrivateSecret.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml", "multipart/form-data"}

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v SecretResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
ProjectSecretApiService Delete a secret
Deletes the information of the secret with the matching secret slug.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Slug field of the organization the resource is scoped to
 * @param project Slug field of the project the resource is scoped to
 * @param secret Slug field of the secret
 * @param optional nil or *ProjectSecretApiDeleteProjectScopedSecretOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return SecretResponse
*/

type ProjectSecretApiDeleteProjectScopedSecretOpts struct {
	HarnessAccount optional.String
}

func (a *ProjectSecretApiService) DeleteProjectScopedSecret(ctx context.Context, org string, project string, secret string, localVarOptionals *ProjectSecretApiDeleteProjectScopedSecretOpts) (SecretResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SecretResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/secrets/{secret}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secret"+"}", fmt.Sprintf("%v", secret), -1)

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SecretResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
ProjectSecretApiService Retrieve a secret
Retrieves the information of the secret.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Slug field of the organization the resource is scoped to
 * @param project Slug field of the project the resource is scoped to
 * @param secret Slug field of the secret
 * @param optional nil or *ProjectSecretApiGetProjectScopedSecretOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return SecretResponse
*/

type ProjectSecretApiGetProjectScopedSecretOpts struct {
	HarnessAccount optional.String
}

func (a *ProjectSecretApiService) GetProjectScopedSecret(ctx context.Context, org string, project string, secret string, localVarOptionals *ProjectSecretApiGetProjectScopedSecretOpts) (SecretResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SecretResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/secrets/{secret}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secret"+"}", fmt.Sprintf("%v", secret), -1)

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SecretResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
ProjectSecretApiService List secrets
Retrieves the information of the secrets.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Slug field of the organization the resource is scoped to
 * @param project Slug field of the project the resource is scoped to
 * @param optional nil or *ProjectSecretApiGetProjectScopedSecretsOpts - Optional Parameters:
     * @param "Secret" (optional.Interface of []string) -  Slug field of secrets
     * @param "Type_" (optional.Interface of []string) -  Secret types on which the filter will be applied
     * @param "Recursive" (optional.Bool) -  Expand current scope to include all child scopes
     * @param "SearchTerm" (optional.String) -  This would be used to filter resources having attributes matching with search term.
     * @param "Page" (optional.Int32) -  Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page
     * @param "Limit" (optional.Int32) -  Pagination: Number of items to return
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return []SecretResponse
*/

type ProjectSecretApiGetProjectScopedSecretsOpts struct {
	Secret         optional.Interface
	Type_          optional.Interface
	Recursive      optional.Bool
	SearchTerm     optional.String
	Page           optional.Int32
	Limit          optional.Int32
	HarnessAccount optional.String
}

func (a *ProjectSecretApiService) GetProjectScopedSecrets(ctx context.Context, org string, project string, localVarOptionals *ProjectSecretApiGetProjectScopedSecretsOpts) ([]SecretResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []SecretResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/secrets"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Secret.IsSet() {
		localVarQueryParams.Add("secret", parameterToString(localVarOptionals.Secret.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Recursive.IsSet() {
		localVarQueryParams.Add("recursive", parameterToString(localVarOptionals.Recursive.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("search_term", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []SecretResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
ProjectSecretApiService Update a secret
Updates the information of the secret with the matching secret slug.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param org Slug field of the organization the resource is scoped to
 * @param project Slug field of the project the resource is scoped to
 * @param secret Slug field of the secret
 * @param optional nil or *ProjectSecretApiUpdateProjectScopedSecretOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return SecretResponse
*/

type ProjectSecretApiUpdateProjectScopedSecretOpts struct {
	HarnessAccount optional.String
}

func (a *ProjectSecretApiService) UpdateProjectScopedSecret(ctx context.Context, body SecretRequest, org string, project string, secret string, localVarOptionals *ProjectSecretApiUpdateProjectScopedSecretOpts) (SecretResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SecretResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/secrets/{secret}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secret"+"}", fmt.Sprintf("%v", secret), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml", "multipart/form-data"}

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SecretResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
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
