
/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * The Harness Software Delivery Platform uses OpenAPI Specification v3.0. Harness constantly improves these APIs. Please be aware that some improvements could cause breaking changes. # Introduction     The Harness API allows you to integrate and use all the services and modules we provide on the Harness Platform. If you use client-side SDKs, Harness functionality can be integrated with your client-side automation, helping you reduce manual efforts and deploy code faster.    For more information about how Harness works, read our [documentation](https://developer.harness.io/docs/getting-started) or visit the [Harness Developer Hub](https://developer.harness.io/).  ## How it works    The Harness API is a RESTful API that uses standard HTTP verbs. You can send requests in JSON, YAML, or form-data format. The format of the response matches the format of your request. You must send a single request at a time and ensure that you include your authentication key. For more information about this, go to [Authentication](#section/Introduction/Authentication).  ## Get started    Before you start integrating, get to know our API better by reading the following topics:    * [Harness key concepts](https://developer.harness.io/docs/getting-started/learn-harness-key-concepts/)   * [Authentication](#section/Introduction/Authentication)   * [Requests and responses](#section/Introduction/Requests-and-Responses)   * [Common Parameters](#section/Introduction/Common-Parameters-Beta)   * [Status Codes](#section/Introduction/Status-Codes)   * [Errors](#tag/Error-Response)   * [Versioning](#section/Introduction/Versioning-Beta)   * [Pagination](/#section/Introduction/Pagination-Beta)    The methods you need to integrate with depend on the functionality you want to use. Work with  your Harness Solutions Engineer to determine which methods you need.  ## Authentication  To authenticate with the Harness API, you need to:   1. Generate an API token on the Harness Platform.   2. Send the API token you generate in the `x-api-key` header in each request.  ### Generate an API token  To generate an API token, complete the following steps:   1. Go to the [Harness Platform](https://app.harness.io/).   2. On the left-hand navigation, click **My Profile**.   3. Click **+API Key**, enter a name for your key and then click **Save**.   4. Within the API Key tile, click **+Token**.   5. Enter a name for your token and click **Generate Token**. **Important**: Make sure to save your token securely. Harness does not store the API token for future reference, so make sure to save your token securely before you leave the page.  ### Send the API token in your requests  Send the token you created in the Harness Platform in the x-api-key header. For example:   `x-api-key: YOUR_API_KEY_HERE`  ## Requests and Responses    The structure for each request and response is outlined in the API documentation. We have examples in JSON and YAML for every request and response. You can use our online editor to test the examples.  ## Common Parameters [Beta]  | Field Name | Type    | Default | Description    | |------------|---------|---------|----------------| | identifier | string  | none    | URL-friendly version of the name, used to identify a resource within it's scope and so needs to be unique within the scope.                                                                                                            | | name       | string  | none    | Human-friendly name for the resource.                                                                                       | | org        | string  | none    | Limit to provided org identifiers.                                                                                                                     | | project    | string  | none    | Limit to provided project identifiers.                                                                                                                 | | description| string  | none    | More information about the specific resource.                                                                                    | | tags       | map[string]string  | none    | List of labels applied to the resource.                                                                                                                         | | order      | string  | desc    | Order to use when sorting the specified fields. Type: enum(asc,desc).                                                                                                                                     | | sort       | string  | none    | Fields on which to sort. Note: Specify the fields that you want to use for sorting. When doing so, consider the operational overhead of sorting fields. | | limit      | int     | 30      | Pagination: Number of items to return.                                                                                                                 | | page       | int     | 1       | Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page.                  | | created    | int64   | none    | Unix timestamp that shows when the resource was created (in milliseconds).                                                               | | updated    | int64   | none    | Unix timestamp that shows when the resource was last edited (in milliseconds).                                                           |   ## Status Codes    Harness uses conventional HTTP status codes to indicate the status of an API request.    Generally, 2xx responses are reserved for success and 4xx status codes are reserved for failures. A 5xx response code indicates an error on the Harness server.    | Error Code  | Description |   |-------------|-------------|   | 200         |     OK      |   | 201         |   Created   |   | 202         |   Accepted  |   | 204         |  No Content |   | 400         | Bad Request |   | 401         | Unauthorized |   | 403         | Forbidden |   | 412         | Precondition Failed |   | 415         | Unsupported Media Type |   | 500         | Server Error |    To view our error response structures, go [here](#tag/Error-Response).  ## Versioning [Beta]  ### Harness Version   The current version of our Beta APIs is yet to be announced. The version number will use the date-header format and will be valid only for our Beta APIs.  ### Generation   All our beta APIs are versioned as a Generation, and this version is included in the path to every API resource. For example, v1 beta APIs begin with `app.harness.io/v1/`, where v1 is the API Generation.    The version number represents the core API and does not change frequently. The version number changes only if there is a significant departure from the basic underpinnings of the existing API. For example, when Harness performs a system-wide refactoring of core concepts or resources.  ## Pagination [Beta]  We use pagination to place limits on the number of responses associated with list endpoints. Pagination is achieved by the use of limit query parameters. The limit defaults to 30. Its maximum value is 100.  Following are the pagination headers supported in the response bodies of paginated APIs:   1. X-Total-Elements : Indicates the total number of entries in a paginated response.   2. X-Page-Number : Indicates the page number currently returned for a paginated response.   3. X-Page-Size : Indicates the number of entries per page for a paginated response.  For example:    ``` X-Total-Elements : 30 X-Page-Number : 0 X-Page-Size : 10   ``` 
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
 package swagger

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
 
 type InputSetsApiService service
 /*
 InputSetsApiService Create an Input Set
 Creates an Input Set for a Pipeline.
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param body Input Set create request body.
  * @param pipeline Pipeline identifier for the Input Set.
  * @param org Organization identifier
  * @param project Project identifier
  * @param optional nil or *InputSetsApiCreateInputSetOpts - Optional Parameters:
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
 @return InputSetResponseBody
 */
 
 type InputSetsApiCreateInputSetOpts struct {
	 HarnessAccount optional.String
 }
 
 func (a *InputSetsApiService) CreateInputSet(ctx context.Context, body InputSetCreateRequestBody, pipeline string, org string, project string, localVarOptionals *InputSetsApiCreateInputSetOpts) (InputSetResponseBody, *http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Post")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 localVarReturnValue InputSetResponseBody
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 localVarQueryParams.Add("pipeline", parameterToString(pipeline, ""))
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
			 var v InputSetResponseBody
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
 InputSetsApiService Delete an Input Set
 Deletes an Input Set for a Pipeline.
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param org Organization identifier
  * @param project Project identifier
  * @param inputSet Input Set identifier
  * @param pipeline Pipeline identifier for the Input Set.
  * @param optional nil or *InputSetsApiDeleteInputSetOpts - Optional Parameters:
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
 
 */
 
 type InputSetsApiDeleteInputSetOpts struct {
	 HarnessAccount optional.String
 }
 
 func (a *InputSetsApiService) DeleteInputSet(ctx context.Context, org string, project string, inputSet string, pipeline string, localVarOptionals *InputSetsApiDeleteInputSetOpts) (*http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Delete")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets/{input-set}"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"input-set"+"}", fmt.Sprintf("%v", inputSet), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 localVarQueryParams.Add("pipeline", parameterToString(pipeline, ""))
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
		 return nil, err
	 }
 
	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		 return localVarHttpResponse, err
	 }
 
	 localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	 localVarHttpResponse.Body.Close()
	 if err != nil {
		 return localVarHttpResponse, err
	 }
 
 
	 if localVarHttpResponse.StatusCode >= 300 {
		 newErr := GenericSwaggerError{
			 body: localVarBody,
			 error: localVarHttpResponse.Status,
		 }
		 if localVarHttpResponse.StatusCode == 400 {
			 var v interface{}
			 err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				 if err != nil {
					 newErr.error = err.Error()
					 return localVarHttpResponse, newErr
				 }
				 newErr.model = v
				 return localVarHttpResponse, newErr
		 }
		 return localVarHttpResponse, newErr
	 }
 
	 return localVarHttpResponse, nil
 }
 /*
 InputSetsApiService Retrieve an Input Set
 Retrieves an Input Set for a Pipeline.
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param org Organization identifier
  * @param project Project identifier
  * @param inputSet Input Set identifier
  * @param pipeline Pipeline identifier for the Input Set.
  * @param optional nil or *InputSetsApiGetInputSetOpts - Optional Parameters:
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
	  * @param "BranchName" (optional.String) -  Name of the branch (for Git Experience).
	  * @param "ParentEntityConnectorRef" (optional.String) -  Connector reference for Parent Entity (Pipeline).
	  * @param "ParentEntityRepoName" (optional.String) -  Repository name for Parent Entity (Pipeline).
	  * @param "LoadFromFallbackBranch" (optional.Bool) -  Flag to load the pipeline from the created non default branch
	  * @param "LoadFromCache" (optional.String) -  Specifies whether the remote pipeline should be loaded from Git or Git cache.
 @return InputSetResponseBody
 */
 
 type InputSetsApiGetInputSetOpts struct {
	 HarnessAccount optional.String
	 BranchName optional.String
	 ParentEntityConnectorRef optional.String
	 ParentEntityRepoName optional.String
	 LoadFromFallbackBranch optional.Bool
	 LoadFromCache optional.String
 }
 
 func (a *InputSetsApiService) GetInputSet(ctx context.Context, org string, project string, inputSet string, pipeline string, localVarOptionals *InputSetsApiGetInputSetOpts) (InputSetResponseBody, *http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Get")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 localVarReturnValue InputSetResponseBody
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets/{input-set}"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"input-set"+"}", fmt.Sprintf("%v", inputSet), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 localVarQueryParams.Add("pipeline", parameterToString(pipeline, ""))
	 if localVarOptionals != nil && localVarOptionals.BranchName.IsSet() {
		 localVarQueryParams.Add("branch_name", parameterToString(localVarOptionals.BranchName.Value(), ""))
	 }
	 if localVarOptionals != nil && localVarOptionals.ParentEntityConnectorRef.IsSet() {
		 localVarQueryParams.Add("parent_entity_connector_ref", parameterToString(localVarOptionals.ParentEntityConnectorRef.Value(), ""))
	 }
	 if localVarOptionals != nil && localVarOptionals.ParentEntityRepoName.IsSet() {
		 localVarQueryParams.Add("parent_entity_repo_name", parameterToString(localVarOptionals.ParentEntityRepoName.Value(), ""))
	 }
	 if localVarOptionals != nil && localVarOptionals.LoadFromFallbackBranch.IsSet() {
		 localVarQueryParams.Add("load_from_fallback_branch", parameterToString(localVarOptionals.LoadFromFallbackBranch.Value(), ""))
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
	 if localVarOptionals != nil && localVarOptionals.LoadFromCache.IsSet() {
		 localVarHeaderParams["Load-From-Cache"] = parameterToString(localVarOptionals.LoadFromCache.Value(), "")
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
			 var v InputSetResponseBody
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
 InputSetsApiService Get Input Set YAML from Git Repository
 Fetches InputSet YAML from Git Repository and saves a record for it in Harness
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param pipeline Pipeline identifier for the Input Set.
  * @param org Organization identifier
  * @param project Project identifier
  * @param inputSet Input Set identifier
  * @param optional nil or *InputSetsApiImportInputSetFromGitOpts - Optional Parameters:
	  * @param "Body" (optional.Interface of InputSetImportRequestBody) -  Input Set import request body
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
 @return InputSetImportResponseBody
 */
 
 type InputSetsApiImportInputSetFromGitOpts struct {
	 Body optional.Interface
	 HarnessAccount optional.String
 }
 
 func (a *InputSetsApiService) ImportInputSetFromGit(ctx context.Context, pipeline string, org string, project string, inputSet string, localVarOptionals *InputSetsApiImportInputSetFromGitOpts) (InputSetImportResponseBody, *http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Post")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 localVarReturnValue InputSetImportResponseBody
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets/{input-set}/import"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"input-set"+"}", fmt.Sprintf("%v", inputSet), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 localVarQueryParams.Add("pipeline", parameterToString(pipeline, ""))
	 // to determine the Content-Type header
	 localVarHttpContentTypes := []string{"application/json"}
 
	 // set Content-Type header
	 localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	 if localVarHttpContentType != "" {
		 localVarHeaderParams["Content-Type"] = localVarHttpContentType
	 }
 
	 // to determine the Accept header
	 localVarHttpHeaderAccepts := []string{"application/json"}
 
	 // set Accept header
	 localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	 if localVarHttpHeaderAccept != "" {
		 localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	 }
	 if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		 localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	 }
	 // body params
	 if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		 
		 localVarOptionalBody:= localVarOptionals.Body.Value()
		 localVarPostBody = &localVarOptionalBody
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
			 var v InputSetImportResponseBody
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
 InputSetsApiService Move InputSet YAML from inline to remote
 Creates a remote entity by fetching the input set YAML from Harness.
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param org Organization identifier
  * @param project Project identifier
  * @param inputSet Input Set identifier
  * @param optional nil or *InputSetsApiInputSetsMoveConfigOpts - Optional Parameters:
	  * @param "Body" (optional.Interface of InputSetMoveConfigRequestBody) - 
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
 @return InputSetMoveConfigResponseBody
 */
 
 type InputSetsApiInputSetsMoveConfigOpts struct {
	 Body optional.Interface
	 HarnessAccount optional.String
 }
 
 func (a *InputSetsApiService) InputSetsMoveConfig(ctx context.Context, org string, project string, inputSet string, localVarOptionals *InputSetsApiInputSetsMoveConfigOpts) (InputSetMoveConfigResponseBody, *http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Post")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 localVarReturnValue InputSetMoveConfigResponseBody
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets/{input-set}/move-config"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"input-set"+"}", fmt.Sprintf("%v", inputSet), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 // to determine the Content-Type header
	 localVarHttpContentTypes := []string{"application/json"}
 
	 // set Content-Type header
	 localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	 if localVarHttpContentType != "" {
		 localVarHeaderParams["Content-Type"] = localVarHttpContentType
	 }
 
	 // to determine the Accept header
	 localVarHttpHeaderAccepts := []string{"application/json"}
 
	 // set Accept header
	 localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	 if localVarHttpHeaderAccept != "" {
		 localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	 }
	 if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		 localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	 }
	 // body params
	 if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		 
		 localVarOptionalBody:= localVarOptionals.Body.Value()
		 localVarPostBody = &localVarOptionalBody
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
			 var v InputSetMoveConfigResponseBody
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
 InputSetsApiService List Input Sets
 Returns a List of Input Sets for a Pipeline.
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param org Organization identifier
  * @param project Project identifier
  * @param pipeline Pipeline identifier for the Input Set.
  * @param optional nil or *InputSetsApiListInputSetsOpts - Optional Parameters:
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
	  * @param "Page" (optional.Int32) -  Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page.
	  * @param "Limit" (optional.Int32) -  Pagination: Number of items to return.
	  * @param "SearchTerm" (optional.String) -  This would be used to filter resources having attributes matching the search term.
	  * @param "Sort" (optional.String) -  Parameter on the basis of which sorting is done.
	  * @param "Order" (optional.String) -  Order on the basis of which sorting is done.
 @return []InputSetResponseBody
 */
 
 type InputSetsApiListInputSetsOpts struct {
	 HarnessAccount optional.String
	 Page optional.Int32
	 Limit optional.Int32
	 SearchTerm optional.String
	 Sort optional.String
	 Order optional.String
 }
 
 func (a *InputSetsApiService) ListInputSets(ctx context.Context, org string, project string, pipeline string, localVarOptionals *InputSetsApiListInputSetsOpts) ([]InputSetResponseBody, *http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Get")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 localVarReturnValue []InputSetResponseBody
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 localVarQueryParams.Add("pipeline", parameterToString(pipeline, ""))
	 if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		 localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	 }
	 if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		 localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	 }
	 if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		 localVarQueryParams.Add("search_term", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
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
			 var v []InputSetResponseBody
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
 InputSetsApiService Update an Input Set
 Updates an Input Set for a Pipeline.
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param body Input Set update request body
  * @param pipeline Pipeline identifier for the Input Set.
  * @param org Organization identifier
  * @param project Project identifier
  * @param inputSet Input Set identifier
  * @param optional nil or *InputSetsApiUpdateInputSetOpts - Optional Parameters:
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
 @return InputSetResponseBody
 */
 
 type InputSetsApiUpdateInputSetOpts struct {
	 HarnessAccount optional.String
 }
 
 func (a *InputSetsApiService) UpdateInputSet(ctx context.Context, body InputSetUpdateRequestBody, pipeline string, org string, project string, inputSet string, localVarOptionals *InputSetsApiUpdateInputSetOpts) (InputSetResponseBody, *http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Put")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 localVarReturnValue InputSetResponseBody
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets/{input-set}"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"input-set"+"}", fmt.Sprintf("%v", inputSet), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 localVarQueryParams.Add("pipeline", parameterToString(pipeline, ""))
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
			 var v InputSetResponseBody
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
 InputSetsApiService Update GitMetadata for Remote InputSet
 Update git-metadata in remote inputSet and return the updated inputSet
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param pipeline Pipeline identifier for the Input Set.
  * @param org Organization identifier
  * @param project Project identifier
  * @param inputSet Input Set identifier
  * @param optional nil or *InputSetsApiUpdateInputSetGitMetadataOpts - Optional Parameters:
	  * @param "Body" (optional.Interface of GitMetadataUpdateRequestBody) - 
	  * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to.
 @return GitMetadataUpdateResponseBody
 */
 
 type InputSetsApiUpdateInputSetGitMetadataOpts struct {
	 Body optional.Interface
	 HarnessAccount optional.String
 }
 
 func (a *InputSetsApiService) UpdateInputSetGitMetadata(ctx context.Context, pipeline string, org string, project string, inputSet string, localVarOptionals *InputSetsApiUpdateInputSetGitMetadataOpts) (GitMetadataUpdateResponseBody, *http.Response, error) {
	 var (
		 localVarHttpMethod = strings.ToUpper("Put")
		 localVarPostBody   interface{}
		 localVarFileName   string
		 localVarFileBytes  []byte
		 localVarReturnValue GitMetadataUpdateResponseBody
	 )
 
	 // create path and map variables
	 localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/input-sets/{input-set}/git-metadata"
	 localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	 localVarPath = strings.Replace(localVarPath, "{"+"input-set"+"}", fmt.Sprintf("%v", inputSet), -1)
 
	 localVarHeaderParams := make(map[string]string)
	 localVarQueryParams := url.Values{}
	 localVarFormParams := url.Values{}
 
	 localVarQueryParams.Add("pipeline", parameterToString(pipeline, ""))
	 // to determine the Content-Type header
	 localVarHttpContentTypes := []string{"application/json"}
 
	 // set Content-Type header
	 localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	 if localVarHttpContentType != "" {
		 localVarHeaderParams["Content-Type"] = localVarHttpContentType
	 }
 
	 // to determine the Accept header
	 localVarHttpHeaderAccepts := []string{"application/json"}
 
	 // set Accept header
	 localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	 if localVarHttpHeaderAccept != "" {
		 localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	 }
	 if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		 localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	 }
	 // body params
	 if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		 
		 localVarOptionalBody:= localVarOptionals.Body.Value()
		 localVarPostBody = &localVarOptionalBody
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
			 var v GitMetadataUpdateResponseBody
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
 