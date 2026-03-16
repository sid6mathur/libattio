# \SCIMUsersAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScimV2UsersGet**](SCIMUsersAPI.md#ScimV2UsersGet) | **Get** /scim/v2/Users | List SCIM users
[**ScimV2UsersPost**](SCIMUsersAPI.md#ScimV2UsersPost) | **Post** /scim/v2/Users | Create SCIM user



## ScimV2UsersGet

> ScimV2SchemasGet200Response ScimV2UsersGet(ctx).Execute()

List SCIM users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/sid6mathur/libattio"
)

func main() {

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.SCIMUsersAPI.ScimV2UsersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCIMUsersAPI.ScimV2UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimV2UsersGet`: ScimV2SchemasGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SCIMUsersAPI.ScimV2UsersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimV2UsersGetRequest struct via the builder pattern


### Return type

[**ScimV2SchemasGet200Response**](ScimV2SchemasGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimV2UsersPost

> ScimV2UsersPost201Response ScimV2UsersPost(ctx).Execute()

Create SCIM user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/sid6mathur/libattio"
)

func main() {

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.SCIMUsersAPI.ScimV2UsersPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCIMUsersAPI.ScimV2UsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimV2UsersPost`: ScimV2UsersPost201Response
	fmt.Fprintf(os.Stdout, "Response from `SCIMUsersAPI.ScimV2UsersPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimV2UsersPostRequest struct via the builder pattern


### Return type

[**ScimV2UsersPost201Response**](ScimV2UsersPost201Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

