# \SCIMSchemasAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScimV2SchemasGet**](SCIMSchemasAPI.md#ScimV2SchemasGet) | **Get** /scim/v2/Schemas | List SCIM schemas



## ScimV2SchemasGet

> ScimV2SchemasGet200Response ScimV2SchemasGet(ctx).Execute()

List SCIM schemas



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
	resp, r, err := apiClient.SCIMSchemasAPI.ScimV2SchemasGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCIMSchemasAPI.ScimV2SchemasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimV2SchemasGet`: ScimV2SchemasGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SCIMSchemasAPI.ScimV2SchemasGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimV2SchemasGetRequest struct via the builder pattern


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

