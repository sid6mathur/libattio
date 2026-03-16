# \FilesAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2FilesFileIdDelete**](FilesAPI.md#V2FilesFileIdDelete) | **Delete** /v2/files/{file_id} | Delete a file
[**V2FilesFileIdDownloadGet**](FilesAPI.md#V2FilesFileIdDownloadGet) | **Get** /v2/files/{file_id}/download | Download a file
[**V2FilesFileIdGet**](FilesAPI.md#V2FilesFileIdGet) | **Get** /v2/files/{file_id} | Get a file
[**V2FilesGet**](FilesAPI.md#V2FilesGet) | **Get** /v2/files | List files
[**V2FilesPost**](FilesAPI.md#V2FilesPost) | **Post** /v2/files | Create a folder
[**V2FilesUploadPost**](FilesAPI.md#V2FilesUploadPost) | **Post** /v2/files/upload | Upload a file



## V2FilesFileIdDelete

> map[string]interface{} V2FilesFileIdDelete(ctx, fileId).Execute()

Delete a file



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
	fileId := "a1b2c3d4-e5f6-7890-abcd-ef1234567890" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.V2FilesFileIdDelete(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.V2FilesFileIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2FilesFileIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.V2FilesFileIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2FilesFileIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2FilesFileIdDownloadGet

> V2FilesFileIdDownloadGet(ctx, fileId).Execute()

Download a file



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	r, err := apiClient.FilesAPI.V2FilesFileIdDownloadGet(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.V2FilesFileIdDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2FilesFileIdDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2FilesFileIdGet

> V2FilesFileIdGet200Response V2FilesFileIdGet(ctx, fileId).Execute()

Get a file



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
	fileId := "a1b2c3d4-e5f6-7890-abcd-ef1234567890" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.V2FilesFileIdGet(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.V2FilesFileIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2FilesFileIdGet`: V2FilesFileIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.V2FilesFileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2FilesFileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2FilesFileIdGet200Response**](V2FilesFileIdGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2FilesGet

> V2FilesGet200Response V2FilesGet(ctx).Object(object).RecordId(recordId).StorageProvider(storageProvider).ParentFolderId(parentFolderId).Limit(limit).Cursor(cursor).Execute()

List files



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
	object := "object_example" // string | 
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	storageProvider := "storageProvider_example" // string |  (optional)
	parentFolderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	limit := int32(50) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.V2FilesGet(context.Background()).Object(object).RecordId(recordId).StorageProvider(storageProvider).ParentFolderId(parentFolderId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.V2FilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2FilesGet`: V2FilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.V2FilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2FilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **object** | **string** |  | 
 **recordId** | **string** |  | 
 **storageProvider** | **string** |  | 
 **parentFolderId** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**V2FilesGet200Response**](V2FilesGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2FilesPost

> V2FilesPost200Response V2FilesPost(ctx).V2FilesPostRequest(v2FilesPostRequest).Execute()

Create a folder



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
	v2FilesPostRequest := *libattio.NewV2FilesPostRequest("people", "bf071e1f-6035-429d-b874-d83ea64ea13b", "FileType_example", "Documents", "google-drive", "01ISGXZ5BRAMVD7SEPXNCYS4XGKT3YTOKQ") // V2FilesPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.V2FilesPost(context.Background()).V2FilesPostRequest(v2FilesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.V2FilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2FilesPost`: V2FilesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.V2FilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2FilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2FilesPostRequest** | [**V2FilesPostRequest**](V2FilesPostRequest.md) |  | 

### Return type

[**V2FilesPost200Response**](V2FilesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2FilesUploadPost

> V2FilesUploadPost201Response V2FilesUploadPost(ctx).File(file).Object(object).RecordId(recordId).ParentFolderId(parentFolderId).Execute()

Upload a file



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
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload.
	object := "object_example" // string | The object slug or ID.
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the record to upload the file to.
	parentFolderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional parent folder ID. Omit to upload to the root folder. (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.V2FilesUploadPost(context.Background()).File(file).Object(object).RecordId(recordId).ParentFolderId(parentFolderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.V2FilesUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2FilesUploadPost`: V2FilesUploadPost201Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.V2FilesUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2FilesUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to upload. | 
 **object** | **string** | The object slug or ID. | 
 **recordId** | **string** | The ID of the record to upload the file to. | 
 **parentFolderId** | **string** | Optional parent folder ID. Omit to upload to the root folder. | 

### Return type

[**V2FilesUploadPost201Response**](V2FilesUploadPost201Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

