# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePipeline**](PipelinesApi.md#CreatePipeline) | **Post** /v1/orgs/{org}/projects/{project}/pipelines | Create a Pipeline
[**DeletePipeline**](PipelinesApi.md#DeletePipeline) | **Delete** /v1/orgs/{org}/projects/{project}/pipelines/{pipeline} | Delete a Pipeline
[**GetPipeline**](PipelinesApi.md#GetPipeline) | **Get** /v1/orgs/{org}/projects/{project}/pipelines/{pipeline} | Retrieve a Pipeline
[**ListPipelines**](PipelinesApi.md#ListPipelines) | **Get** /v1/orgs/{org}/projects/{project}/pipelines | List Pipelines
[**UpdatePipeline**](PipelinesApi.md#UpdatePipeline) | **Put** /v1/orgs/{org}/projects/{project}/pipelines/{pipeline} | Update a Pipeline

# **CreatePipeline**
> PipelineCreateResponseBody CreatePipeline(ctx, body, org, project, optional)
Create a Pipeline

Creates a Pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PipelineCreateRequestBody**](PipelineCreateRequestBody.md)| Pipeline request body | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***PipelinesApiCreatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiCreatePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**PipelineCreateResponseBody**](PipelineCreateResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipeline**
> DeletePipeline(ctx, org, project, pipeline, optional)
Delete a Pipeline

Deletes a Pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **pipeline** | **string**| Pipeline identifier | 
 **optional** | ***PipelinesApiDeletePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiDeletePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipeline**
> PipelineGetResponseBody GetPipeline(ctx, org, project, pipeline, optional)
Retrieve a Pipeline

Retrieves a Pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **pipeline** | **string**| Pipeline identifier | 
 **optional** | ***PipelinesApiGetPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiGetPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **branchName** | **optional.String**| Name of the branch (for Git Experience). | 
 **templateApplied** | **optional.Bool**| If true, returns Pipeline YAML with Templates applied on it. | [default to false]
 **connectorRef** | **optional.String**| Identifier of the Harness Connector used for CRUD operations on the Entity (for Git Experience). | 
 **repoName** | **optional.String**| Name of the repository (for Git Experience). | 

### Return type

[**PipelineGetResponseBody**](PipelineGetResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPipelines**
> []PipelineListResponseBody ListPipelines(ctx, org, project, optional)
List Pipelines

Returns a list of Pipelines.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***PipelinesApiListPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiListPipelinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return. | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching the search term. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **module** | **optional.String**| Harness module which is part of the Pipeline. | 
 **filterIdentifier** | **optional.String**| Identifier of a saved Filter. | 
 **pipelineIdentifiers** | [**optional.Interface of []string**](string.md)| List of Pipeline identifiers on the basis of which the Pipelines are filtered. | 
 **name** | **optional.String**| Pipeline Name on the basis of which the Pipelines are filtered. | 
 **description** | **optional.String**| Pipeline Description on the basis of which the Pipelines are filtered. | 
 **tags** | [**optional.Interface of []string**](string.md)| Filter tags as a key:value pair. | 
 **serviceNames** | [**optional.Interface of []string**](string.md)| Service names on the basis of which the Pipelines are filtered. | 
 **envNames** | [**optional.Interface of []string**](string.md)| Names of Environments on the basis of which the Pipelines are filtered. | 
 **deploymentType** | **optional.String**| Deployment type on the basis of which the Pipelines are filtered. | 
 **repository** | **optional.String**| Repository name on the basis of which the Pipelines are filtered. | 

### Return type

[**[]PipelineListResponseBody**](PipelineListResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipeline**
> PipelineCreateResponseBody UpdatePipeline(ctx, body, org, project, pipeline, optional)
Update a Pipeline

Updates a Pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PipelineUpdateRequestBody**](PipelineUpdateRequestBody.md)| Pipeline request body | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **pipeline** | **string**| Pipeline identifier | 
 **optional** | ***PipelinesApiUpdatePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiUpdatePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**PipelineCreateResponseBody**](PipelineCreateResponseBody.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

