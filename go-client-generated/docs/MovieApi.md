# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/k0k0piotrowski/SEP6-Backend/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MovieGet**](MovieApi.md#MovieGet) | **Get** /movie | Search movies
[**MovieMovieIdGet**](MovieApi.md#MovieMovieIdGet) | **Get** /movie/{movieId} | Get movie by id
[**MoviePopularGet**](MovieApi.md#MoviePopularGet) | **Get** /movie/popular | Get popular movies
[**MovieTopGet**](MovieApi.md#MovieTopGet) | **Get** /movie/top | Get top movies

# **MovieGet**
> ReturnMovies MovieGet(ctx, optional)
Search movies

Returns movies by searched keyword

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MovieApiMovieGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MovieApiMovieGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| The search input | 
 **page** | **optional.Int32**| The page number | 

### Return type

[**ReturnMovies**](ReturnMovies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MovieMovieIdGet**
> Movie MovieMovieIdGet(ctx, movieId)
Get movie by id

Returns movie object (json) that will contain movie title, actors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **movieId** | **int32**| Numeric ID of the movie to get | 

### Return type

[**Movie**](Movie.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoviePopularGet**
> map[string]Movie MoviePopularGet(ctx, optional)
Get popular movies

Returns list of movies considered popular by The Movie DB Api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MovieApiMoviePopularGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MovieApiMoviePopularGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The page number | 

### Return type

[**map[string]Movie**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MovieTopGet**
> ReturnMovies MovieTopGet(ctx, optional)
Get top movies

Returns list of top rated movies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MovieApiMovieTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MovieApiMovieTopGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The page number | 

### Return type

[**ReturnMovies**](ReturnMovies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

