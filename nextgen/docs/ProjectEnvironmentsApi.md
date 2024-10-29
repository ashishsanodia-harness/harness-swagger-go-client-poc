# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvServiceOverride**](ProjectEnvironmentsApi.md#CreateEnvServiceOverride) | **Post** /v1/orgs/{org}/projects/{project}/environments/{environment}/services/{service} | Create Service Override
[**CreateEnvironment**](ProjectEnvironmentsApi.md#CreateEnvironment) | **Post** /v1/orgs/{org}/projects/{project}/environments | Create an Environment
[**DeleteEnvServiceOverride**](ProjectEnvironmentsApi.md#DeleteEnvServiceOverride) | **Delete** /v1/orgs/{org}/projects/{project}/environments/{environment}/services/{service} | Delete Service Override
[**DeleteEnvironment**](ProjectEnvironmentsApi.md#DeleteEnvironment) | **Delete** /v1/orgs/{org}/projects/{project}/environments/{environment} | Delete an Environment
[**GetEnvServiceOverride**](ProjectEnvironmentsApi.md#GetEnvServiceOverride) | **Get** /v1/orgs/{org}/projects/{project}/environments/{environment}/services/{service} | Retrieve a Service Override
[**GetEnvServiceOverrides**](ProjectEnvironmentsApi.md#GetEnvServiceOverrides) | **Get** /v1/orgs/{org}/projects/{project}/environments/{environment}/services | Retrieve Service Overrides list
[**GetEnvironment**](ProjectEnvironmentsApi.md#GetEnvironment) | **Get** /v1/orgs/{org}/projects/{project}/environments/{environment} | Retrieve an Environment
[**GetEnvironments**](ProjectEnvironmentsApi.md#GetEnvironments) | **Get** /v1/orgs/{org}/projects/{project}/environments | List Environments
[**UpdateEnvServiceOverride**](ProjectEnvironmentsApi.md#UpdateEnvServiceOverride) | **Put** /v1/orgs/{org}/projects/{project}/environments/{environment}/services/{service} | Update Service Override
[**UpdateEnvironment**](ProjectEnvironmentsApi.md#UpdateEnvironment) | **Put** /v1/orgs/{org}/projects/{project}/environments/{environment} | Update Environment

# **CreateEnvServiceOverride**
> ServiceOverrideResponse CreateEnvServiceOverride(ctx, org, project, environment, service, optional)
Create Service Override

Creates a Service Override

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
  **service** | **string**| Slug field of the service the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiCreateEnvServiceOverrideOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiCreateEnvServiceOverrideOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of ServiceOverrideRequest**](ServiceOverrideRequest.md)| Create/Update Service Override Request Body | 
 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceOverrideResponse**](ServiceOverrideResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnvironment**
> EnvironmentResponse CreateEnvironment(ctx, body, org, project, optional)
Create an Environment

Creates an Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentRequest**](EnvironmentRequest.md)| Create Environment request body | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiCreateEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiCreateEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvServiceOverride**
> DeleteEnvServiceOverride(ctx, org, project, environment, service, optional)
Delete Service Override

Deletes Service Override

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
  **service** | **string**| Slug field of the service the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiDeleteEnvServiceOverrideOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiDeleteEnvServiceOverrideOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironment**
> DeleteEnvironment(ctx, org, project, environment, optional)
Delete an Environment

Deletes the requested environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiDeleteEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiDeleteEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvServiceOverride**
> ServiceOverrideResponse GetEnvServiceOverride(ctx, org, project, environment, service, optional)
Retrieve a Service Override

Gets a Service Override

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
  **service** | **string**| Slug field of the service the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiGetEnvServiceOverrideOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiGetEnvServiceOverrideOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceOverrideResponse**](ServiceOverrideResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvServiceOverrides**
> []ServiceOverrideResponse GetEnvServiceOverrides(ctx, org, project, environment, optional)
Retrieve Service Overrides list

Gets Service Override List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiGetEnvServiceOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiGetEnvServiceOverridesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]

### Return type

[**[]ServiceOverrideResponse**](array.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironment**
> EnvironmentResponse GetEnvironment(ctx, org, project, environment, optional)
Retrieve an Environment

Retrieves specified Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiGetEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiGetEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironments**
> []EnvironmentResponse GetEnvironments(ctx, org, project, optional)
List Environments

Returns all environments that the current user has view permissions for given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiGetEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiGetEnvironmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **environmentIds** | [**optional.Interface of []string**](string.md)| List of Environment Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specify whether list an access list or not. Access List refers to list of services that current user has permission to use in the pipeline. | [default to false]

### Return type

[**[]EnvironmentResponse**](array.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvServiceOverride**
> ServiceOverrideResponse UpdateEnvServiceOverride(ctx, org, project, environment, service, optional)
Update Service Override

Update Service Override

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
  **service** | **string**| Slug field of the service the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiUpdateEnvServiceOverrideOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiUpdateEnvServiceOverrideOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of ServiceOverrideRequest**](ServiceOverrideRequest.md)| Create/Update Service Override Request Body | 
 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceOverrideResponse**](ServiceOverrideResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironment**
> EnvironmentResponse UpdateEnvironment(ctx, body, org, project, environment, optional)
Update Environment

Update Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentRequest**](EnvironmentRequest.md)| Create Environment request body | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **environment** | **string**| Slug field of the environemnt the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiUpdateEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiUpdateEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

