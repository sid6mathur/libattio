# \SCIMGroupsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScimV2GroupsGet**](SCIMGroupsAPI.md#ScimV2GroupsGet) | **Get** /scim/v2/Groups | List SCIM groups
[**ScimV2GroupsPost**](SCIMGroupsAPI.md#ScimV2GroupsPost) | **Post** /scim/v2/Groups | Create SCIM group



## ScimV2GroupsGet

> ScimV2SchemasGet200Response ScimV2GroupsGet(ctx).Execute()

List SCIM groups



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
	resp, r, err := apiClient.SCIMGroupsAPI.ScimV2GroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCIMGroupsAPI.ScimV2GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimV2GroupsGet`: ScimV2SchemasGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SCIMGroupsAPI.ScimV2GroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimV2GroupsGetRequest struct via the builder pattern


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


## ScimV2GroupsPost

> ScimV2GroupsPost201Response ScimV2GroupsPost(ctx).Execute()

Create SCIM group



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
	resp, r, err := apiClient.SCIMGroupsAPI.ScimV2GroupsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCIMGroupsAPI.ScimV2GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimV2GroupsPost`: ScimV2GroupsPost201Response
	fmt.Fprintf(os.Stdout, "Response from `SCIMGroupsAPI.ScimV2GroupsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimV2GroupsPostRequest struct via the builder pattern


### Return type

[**ScimV2GroupsPost201Response**](ScimV2GroupsPost201Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

