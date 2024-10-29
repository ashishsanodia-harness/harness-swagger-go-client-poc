# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplatesOrg**](OrgTemplateApi.md#CreateTemplatesOrg) | **Post** /v1/orgs/{org}/templates | Create Template
[**DeleteTemplateOrg**](OrgTemplateApi.md#DeleteTemplateOrg) | **Delete** /v1/orgs/{org}/templates/{template}/versions/{version} | Delete Template
[**GetTemplateOrg**](OrgTemplateApi.md#GetTemplateOrg) | **Get** /v1/orgs/{org}/templates/{template}/versions/{version} | Retrieve a Template
[**GetTemplateStableOrg**](OrgTemplateApi.md#GetTemplateStableOrg) | **Get** /v1/orgs/{org}/templates/{template} | Get Stable Template
[**GetTemplatesListOrg**](OrgTemplateApi.md#GetTemplatesListOrg) | **Get** /v1/orgs/{org}/templates | Get Templates List
[**UpdateTemplateOrg**](OrgTemplateApi.md#UpdateTemplateOrg) | **Put** /v1/orgs/{org}/templates/{template}/versions/{version} | Update Template
[**UpdateTemplateStableOrg**](OrgTemplateApi.md#UpdateTemplateStableOrg) | **Put** /v1/orgs/{org}/templates/{template}/versions/{version}/stable | Update Stable Template

# **CreateTemplatesOrg**
> TemplateResponse CreateTemplatesOrg(ctx, org, optional)
Create Template

Creates a Template in the Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization Identifier | 
 **optional** | ***OrgTemplateApiCreateTemplatesOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgTemplateApiCreateTemplatesOrgOpts struct
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

# **DeleteTemplateOrg**
> DeleteTemplateOrg(ctx, template, org, version, optional)
Delete Template

Deletes particular version of Template at Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
  **org** | **string**| Organization Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***OrgTemplateApiDeleteTemplateOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgTemplateApiDeleteTemplateOrgOpts struct
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

# **GetTemplateOrg**
> TemplateWithInputsResponse GetTemplateOrg(ctx, template, org, version, optional)
Retrieve a Template

Retrieves particular version of Template at Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
  **org** | **string**| Organization Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***OrgTemplateApiGetTemplateOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgTemplateApiGetTemplateOrgOpts struct
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

# **GetTemplateStableOrg**
> TemplateWithInputsResponse GetTemplateStableOrg(ctx, org, template, optional)
Get Stable Template

Retrieves stable version of Template at Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization Identifier | 
  **template** | **string**| Template Identifier | 
 **optional** | ***OrgTemplateApiGetTemplateStableOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgTemplateApiGetTemplateStableOrgOpts struct
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

# **GetTemplatesListOrg**
> []TemplateMetadataSummaryResponse GetTemplatesListOrg(ctx, org, optional)
Get Templates List

Retrieves list of Template with meta-data at Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization Identifier | 
 **optional** | ***OrgTemplateApiGetTemplatesListOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgTemplateApiGetTemplatesListOrgOpts struct
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

# **UpdateTemplateOrg**
> TemplateResponse UpdateTemplateOrg(ctx, template, org, version, optional)
Update Template

Updates particular version of Template at Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **template** | **string**| Template Identifier | 
  **org** | **string**| Organization Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***OrgTemplateApiUpdateTemplateOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgTemplateApiUpdateTemplateOrgOpts struct
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

# **UpdateTemplateStableOrg**
> TemplateUpdateStableResponse UpdateTemplateStableOrg(ctx, org, template, version, optional)
Update Stable Template

Updates the stable version of Template at Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization Identifier | 
  **template** | **string**| Template Identifier | 
  **version** | **string**| Version Label for Template | 
 **optional** | ***OrgTemplateApiUpdateTemplateStableOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgTemplateApiUpdateTemplateStableOrgOpts struct
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

