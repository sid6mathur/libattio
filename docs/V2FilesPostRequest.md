# V2FilesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | The object slug or ID. | 
**RecordId** | **string** | The ID of the record to create the file entry on. | 
**FileType** | **string** | Creates a connected file entry. | 
**Name** | **string** | The folder name. | 
**ParentFolderId** | Pointer to **string** | Optional parent folder ID. Omit to create a top-level folder. | [optional] 
**StorageProvider** | **string** | The external storage provider. | 
**ExternalProviderFileId** | **string** | The ID of the file or folder in the external storage provider. | 
**MicrosoftDriveId** | Pointer to **string** | Microsoft drive ID. Only used when &#x60;storage_provider&#x60; is &#x60;microsoft-onedrive&#x60;. | [optional] 

## Methods

### NewV2FilesPostRequest

`func NewV2FilesPostRequest(object string, recordId string, fileType string, name string, storageProvider string, externalProviderFileId string, ) *V2FilesPostRequest`

NewV2FilesPostRequest instantiates a new V2FilesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FilesPostRequestWithDefaults

`func NewV2FilesPostRequestWithDefaults() *V2FilesPostRequest`

NewV2FilesPostRequestWithDefaults instantiates a new V2FilesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *V2FilesPostRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *V2FilesPostRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *V2FilesPostRequest) SetObject(v string)`

SetObject sets Object field to given value.


### GetRecordId

`func (o *V2FilesPostRequest) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *V2FilesPostRequest) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *V2FilesPostRequest) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetFileType

`func (o *V2FilesPostRequest) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *V2FilesPostRequest) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *V2FilesPostRequest) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetName

`func (o *V2FilesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2FilesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2FilesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *V2FilesPostRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *V2FilesPostRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *V2FilesPostRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *V2FilesPostRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetStorageProvider

`func (o *V2FilesPostRequest) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *V2FilesPostRequest) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *V2FilesPostRequest) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetExternalProviderFileId

`func (o *V2FilesPostRequest) GetExternalProviderFileId() string`

GetExternalProviderFileId returns the ExternalProviderFileId field if non-nil, zero value otherwise.

### GetExternalProviderFileIdOk

`func (o *V2FilesPostRequest) GetExternalProviderFileIdOk() (*string, bool)`

GetExternalProviderFileIdOk returns a tuple with the ExternalProviderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderFileId

`func (o *V2FilesPostRequest) SetExternalProviderFileId(v string)`

SetExternalProviderFileId sets ExternalProviderFileId field to given value.


### GetMicrosoftDriveId

`func (o *V2FilesPostRequest) GetMicrosoftDriveId() string`

GetMicrosoftDriveId returns the MicrosoftDriveId field if non-nil, zero value otherwise.

### GetMicrosoftDriveIdOk

`func (o *V2FilesPostRequest) GetMicrosoftDriveIdOk() (*string, bool)`

GetMicrosoftDriveIdOk returns a tuple with the MicrosoftDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftDriveId

`func (o *V2FilesPostRequest) SetMicrosoftDriveId(v string)`

SetMicrosoftDriveId sets MicrosoftDriveId field to given value.

### HasMicrosoftDriveId

`func (o *V2FilesPostRequest) HasMicrosoftDriveId() bool`

HasMicrosoftDriveId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


