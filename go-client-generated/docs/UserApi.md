# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/k0k0piotrowski/SEP6-Backend/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserLoginPost**](UserApi.md#UserLoginPost) | **Post** /user/login | Logging in the user
[**UserPlaylistAddToFavouritePost**](UserApi.md#UserPlaylistAddToFavouritePost) | **Post** /user/playlist/addToFavourite | Adding a movie to favourite list of user
[**UserPlaylistGetFavouriteGet**](UserApi.md#UserPlaylistGetFavouriteGet) | **Get** /user/playlist/getFavourite | Get favorite movies list of user
[**UserPlaylistRemoveFromFavouriteMovieIdDelete**](UserApi.md#UserPlaylistRemoveFromFavouriteMovieIdDelete) | **Delete** /user/playlist/removeFromFavourite/{movieId} | Removing movie from favourite list
[**UserRegisterPost**](UserApi.md#UserRegisterPost) | **Post** /user/register | Registering the user

# **UserLoginPost**
> UserLoginPost(ctx, body)
Logging in the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Login**](Login.md)| Object required to send when logging in | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPlaylistAddToFavouritePost**
> UserPlaylistAddToFavouritePost(ctx, body)
Adding a movie to favourite list of user

Add specific movie to the user's list of favourite movies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FavouriteMovie**](FavouriteMovie.md)| Object required to send in order to add a new movie to the favourite movies playlist | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPlaylistGetFavouriteGet**
> map[string]Movie UserPlaylistGetFavouriteGet(ctx, )
Get favorite movies list of user

Returns a whole list of favourite movies for specific user logged in.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]Movie**](map.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPlaylistRemoveFromFavouriteMovieIdDelete**
> UserPlaylistRemoveFromFavouriteMovieIdDelete(ctx, movieId)
Removing movie from favourite list

Removing a movie from list of favourite movies for the specific user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **movieId** | **int32**| Numeric ID of the movie to remove from favourite list | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRegisterPost**
> UserRegisterPost(ctx, body)
Registering the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Register**](Register.md)| Object required to send when registering new user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

