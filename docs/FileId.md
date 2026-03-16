# FileId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | The ID of the workspace the file belongs to. | 
**FileId** | **string** | The ID of the file entry. | 

## Methods

### NewFileId

`func NewFileId(workspaceId string, fileId string, ) *FileId`

NewFileId instantiates a new FileId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileIdWithDefaults

`func NewFileIdWithDefaults() *FileId`

NewFileIdWithDefaults instantiates a new FileId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *FileId) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *FileId) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *FileId) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetFileId

`func (o *FileId) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileId) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileId) SetFileId(v string)`

SetFileId sets FileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


