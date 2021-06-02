# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/k0k0piotrowski/SEP6-Backend/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PersonGet**](PersonApi.md#PersonGet) | **Get** /person | Search people
[**PersonPersonIdGet**](PersonApi.md#PersonPersonIdGet) | **Get** /person/{personId} | Get person
[**PersonPopularGet**](PersonApi.md#PersonPopularGet) | **Get** /person/popular | Get popular people

# **PersonGet**
> ReturnPeople PersonGet(ctx, optional)
Search people

Returns people by searched keyword

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PersonApiPersonGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonApiPersonGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| The search input | 
 **page** | **optional.Int32**| The page number | 

### Return type

[**ReturnPeople**](ReturnPeople.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PersonPersonIdGet**
> Person PersonPersonIdGet(ctx, personId)
Get person

Returns person object by given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **personId** | **int32**| Numeric ID of the actor to get | 

### Return type

[**Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PersonPopularGet**
> ReturnPeople PersonPopularGet(ctx, optional)
Get popular people

Returns list of people considered popular by The Movie DB Api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PersonApiPersonPopularGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonApiPersonPopularGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The page number | 

### Return type

[**ReturnPeople**](ReturnPeople.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

