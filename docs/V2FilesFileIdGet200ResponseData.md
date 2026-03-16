# V2FilesFileIdGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**FileId**](FileId.md) |  | 
**ObjectId** | **string** | The ID of the object the record belongs to. | 
**ObjectSlug** | **string** | The slug of the object the record belongs to. | 
**RecordId** | **string** | The ID of the record the file is linked to. | 
**StorageProvider** | **string** | The storage provider for this file entry. | 
**CreatedByActor** | [**FileCreatedByActor**](FileCreatedByActor.md) |  | 
**CreatedAt** | **string** | Timestamp representing when the file entry was created. | 
**FileType** | **string** | The type of file entry. | 
**Name** | **string** | The name of the folder. | 
**ContentType** | **string** | The content type of the file. | 
**ContentSize** | **float32** | The size of the file in bytes. | 
**ParentFolderId** | **string** | The ID of the parent folder, or null if this is a top-level folder. | 
**HasChildren** | **bool** | Whether the folder contains any child entries. | 
**ExternalProviderFileId** | **string** | The file ID in the external storage provider. | 
**MicrosoftDriveId** | **string** | Microsoft drive ID. This field is only populated for &#x60;microsoft-onedrive&#x60; entries. | 

## Methods

### NewV2FilesFileIdGet200ResponseData

`func NewV2FilesFileIdGet200ResponseData(id FileId, objectId string, objectSlug string, recordId string, storageProvider string, createdByActor FileCreatedByActor, createdAt string, fileType string, name string, contentType string, contentSize float32, parentFolderId string, hasChildren bool, externalProviderFileId string, microsoftDriveId string, ) *V2FilesFileIdGet200ResponseData`

NewV2FilesFileIdGet200ResponseData instantiates a new V2FilesFileIdGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FilesFileIdGet200ResponseDataWithDefaults

`func NewV2FilesFileIdGet200ResponseDataWithDefaults() *V2FilesFileIdGet200ResponseData`

NewV2FilesFileIdGet200ResponseDataWithDefaults instantiates a new V2FilesFileIdGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2FilesFileIdGet200ResponseData) GetId() FileId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2FilesFileIdGet200ResponseData) GetIdOk() (*FileId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2FilesFileIdGet200ResponseData) SetId(v FileId)`

SetId sets Id field to given value.


### GetObjectId

`func (o *V2FilesFileIdGet200ResponseData) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *V2FilesFileIdGet200ResponseData) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *V2FilesFileIdGet200ResponseData) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectSlug

`func (o *V2FilesFileIdGet200ResponseData) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *V2FilesFileIdGet200ResponseData) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *V2FilesFileIdGet200ResponseData) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetRecordId

`func (o *V2FilesFileIdGet200ResponseData) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *V2FilesFileIdGet200ResponseData) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *V2FilesFileIdGet200ResponseData) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStorageProvider

`func (o *V2FilesFileIdGet200ResponseData) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *V2FilesFileIdGet200ResponseData) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *V2FilesFileIdGet200ResponseData) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetCreatedByActor

`func (o *V2FilesFileIdGet200ResponseData) GetCreatedByActor() FileCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2FilesFileIdGet200ResponseData) GetCreatedByActorOk() (*FileCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2FilesFileIdGet200ResponseData) SetCreatedByActor(v FileCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *V2FilesFileIdGet200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2FilesFileIdGet200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2FilesFileIdGet200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileType

`func (o *V2FilesFileIdGet200ResponseData) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *V2FilesFileIdGet200ResponseData) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *V2FilesFileIdGet200ResponseData) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetName

`func (o *V2FilesFileIdGet200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2FilesFileIdGet200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2FilesFileIdGet200ResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetContentType

`func (o *V2FilesFileIdGet200ResponseData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *V2FilesFileIdGet200ResponseData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *V2FilesFileIdGet200ResponseData) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentSize

`func (o *V2FilesFileIdGet200ResponseData) GetContentSize() float32`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *V2FilesFileIdGet200ResponseData) GetContentSizeOk() (*float32, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *V2FilesFileIdGet200ResponseData) SetContentSize(v float32)`

SetContentSize sets ContentSize field to given value.


### GetParentFolderId

`func (o *V2FilesFileIdGet200ResponseData) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *V2FilesFileIdGet200ResponseData) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *V2FilesFileIdGet200ResponseData) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### GetHasChildren

`func (o *V2FilesFileIdGet200ResponseData) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *V2FilesFileIdGet200ResponseData) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *V2FilesFileIdGet200ResponseData) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.


### GetExternalProviderFileId

`func (o *V2FilesFileIdGet200ResponseData) GetExternalProviderFileId() string`

GetExternalProviderFileId returns the ExternalProviderFileId field if non-nil, zero value otherwise.

### GetExternalProviderFileIdOk

`func (o *V2FilesFileIdGet200ResponseData) GetExternalProviderFileIdOk() (*string, bool)`

GetExternalProviderFileIdOk returns a tuple with the ExternalProviderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderFileId

`func (o *V2FilesFileIdGet200ResponseData) SetExternalProviderFileId(v string)`

SetExternalProviderFileId sets ExternalProviderFileId field to given value.


### GetMicrosoftDriveId

`func (o *V2FilesFileIdGet200ResponseData) GetMicrosoftDriveId() string`

GetMicrosoftDriveId returns the MicrosoftDriveId field if non-nil, zero value otherwise.

### GetMicrosoftDriveIdOk

`func (o *V2FilesFileIdGet200ResponseData) GetMicrosoftDriveIdOk() (*string, bool)`

GetMicrosoftDriveIdOk returns a tuple with the MicrosoftDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftDriveId

`func (o *V2FilesFileIdGet200ResponseData) SetMicrosoftDriveId(v string)`

SetMicrosoftDriveId sets MicrosoftDriveId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


