# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplatesAcc**](AccountTemplateApi.md#CreateTemplatesAcc) | **Post** /v1/templates | Create Template
[**DeleteTemplateAcc**](AccountTemplateApi.md#DeleteTemplateAcc) | **Delete** /v1/templates/{template}/versions/{version} | Delete Template
[**GetTemplateAcc**](AccountTemplateApi.md#GetTemplateAcc) | **Get** /v1/templates/{template}/versions/{version} | Retrieve a Template
[**GetTemplateStableAcc**](AccountTemplateApi.md#GetTemplateStableAcc) | **Get** /v1/templates/{template} | Get Stable Template
[**GetTemplatesListAcc**](AccountTemplateApi.md#GetTemplatesListAcc) | **Get** /v1/templates | Get Templates List
[**UpdateTemplateAcc**](AccountTemplateApi.md#UpdateTemplateAcc) | **Put** /v1/templates/{template}/versions/{version} | Update Template
[**UpdateTemplateStableAcc**](AccountTemplateApi.md#UpdateTemplateStableAcc) | **Put** /v1/templates/{template}/versions/{version}/stable | Update Stable Template

# **CreateTemplatesAcc**
> TemplateResponse CreateTemplatesAcc(ctx, optional)
Create Template

Creates a Template in the Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountTemplateApiCreateTemplatesAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountTemplateApiCreateTemplatesAccOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TemplateCreateRequestBody**](TemplateCreateRequestBody.md)| Templates Create Request Body | 
 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**TemplateResponse**](TemplateResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTemplateAcc**
> DeleteTemplateAcc(ctx, template, version, optional)
Delete Template

Deletes particular version of Template at Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***AccountTemplateApiDeleteTemplateAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountTemplateApiDeleteTemplateAccOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **comments** | **optional.String**| Specify comment with respect to changes   | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateAcc**
> TemplateWithInputsResponse GetTemplateAcc(ctx, template, version, optional)
Retrieve a Template

Retrieves particular version of Template at Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***AccountTemplateApiGetTemplateAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountTemplateApiGetTemplateAccOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **includeYaml** | **optional.Bool**| Use it to get Template along with Input Set YAML | 
 **branchName** | **optional.String**| Name of the branch | 
 **parentEntityConnectorRef** | **optional.String**| Connector ref of parent template if its remote | 
 **parentEntityRepoName** | **optional.String**| Repo name of parent template if its remote | 
 **parentEntityAccountId** | **optional.String**| Account name of parent template if its remote | 
 **parentEntityOrgId** | **optional.String**| Organization name of parent template if its remote | 
 **parentEntityProjectId** | **optional.String**| Project name of parent entity if its remote | 

### Return type

[**TemplateWithInputsResponse**](TemplateWithInputsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateStableAcc**
> TemplateWithInputsResponse GetTemplateStableAcc(ctx, template, optional)
Get Stable Template

Retrieves stable version of Template at Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
 **optional** | ***AccountTemplateApiGetTemplateStableAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountTemplateApiGetTemplateStableAccOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **includeYaml** | **optional.Bool**| Use it to get Template along with Input Set YAML | 
 **branchName** | **optional.String**| Name of the branch | 
 **parentEntityConnectorRef** | **optional.String**| Connector ref of parent template if its remote | 
 **parentEntityRepoName** | **optional.String**| Repo name of parent template if its remote | 
 **parentEntityAccountId** | **optional.String**| Account name of parent template if its remote | 
 **parentEntityOrgId** | **optional.String**| Organization name of parent template if its remote | 
 **parentEntityProjectId** | **optional.String**| Project name of parent entity if its remote | 

### Return type

[**TemplateWithInputsResponse**](TemplateWithInputsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplatesListAcc**
> []TemplateMetadataSummaryResponse GetTemplatesListAcc(ctx, optional)
Get Templates List

Retrieves list of Template with meta-data at Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountTemplateApiGetTemplatesListAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountTemplateApiGetTemplatesListAccOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **type_** | **optional.String**| Template List Type | 
 **recursive** | **optional.Bool**| Specify true if all accessible Templates are to be included | [default to false]
 **names** | [**optional.Interface of []string**](string.md)| Template names for filtering | 
 **identifiers** | [**optional.Interface of []string**](string.md)| Template Ids for Filtering | 
 **description** | **optional.String**| Filter properties description | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| Type of Template | 
 **childTypes** | [**optional.Interface of []string**](string.md)| Child types describe the type of Step or stage | 

### Return type

[**[]TemplateMetadataSummaryResponse**](array.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTemplateAcc**
> TemplateResponse UpdateTemplateAcc(ctx, template, version, optional)
Update Template

Updates particular version of Template at Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***AccountTemplateApiUpdateTemplateAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountTemplateApiUpdateTemplateAccOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TemplateUpdateRequestBody**](TemplateUpdateRequestBody.md)| Templates Update Request Body | 
 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**TemplateResponse**](TemplateResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTemplateStableAcc**
> TemplateUpdateStableResponse UpdateTemplateStableAcc(ctx, template, version, optional)
Update Stable Template

Updates the stable version of Template at Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***AccountTemplateApiUpdateTemplateStableAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountTemplateApiUpdateTemplateStableAccOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of GitFindDetails**](GitFindDetails.md)| Templates Fetch Request Body | 
 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**TemplateUpdateStableResponse**](TemplateUpdateStableResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

