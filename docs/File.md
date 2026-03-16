# File

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
**Name** | **string** | The name of the file. | 
**ContentType** | **NullableString** | The content type of the file. | 
**ContentSize** | **NullableFloat32** | The size of the file in bytes. | 
**ParentFolderId** | **NullableString** | The ID of the parent folder, or null if this is a top-level file. | 

## Methods

### NewFile

`func NewFile(id FileId, objectId string, objectSlug string, recordId string, storageProvider string, createdByActor FileCreatedByActor, createdAt string, fileType string, name string, contentType NullableString, contentSize NullableFloat32, parentFolderId NullableString, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *File) GetId() FileId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *File) GetIdOk() (*FileId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *File) SetId(v FileId)`

SetId sets Id field to given value.


### GetObjectId

`func (o *File) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *File) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *File) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectSlug

`func (o *File) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *File) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *File) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetRecordId

`func (o *File) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *File) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *File) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStorageProvider

`func (o *File) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *File) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *File) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetCreatedByActor

`func (o *File) GetCreatedByActor() FileCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *File) GetCreatedByActorOk() (*FileCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *File) SetCreatedByActor(v FileCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *File) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *File) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *File) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileType

`func (o *File) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *File) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *File) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetName

`func (o *File) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *File) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *File) SetName(v string)`

SetName sets Name field to given value.


### GetContentType

`func (o *File) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *File) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *File) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### SetContentTypeNil

`func (o *File) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *File) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContentSize

`func (o *File) GetContentSize() float32`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *File) GetContentSizeOk() (*float32, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *File) SetContentSize(v float32)`

SetContentSize sets ContentSize field to given value.


### SetContentSizeNil

`func (o *File) SetContentSizeNil(b bool)`

 SetContentSizeNil sets the value for ContentSize to be an explicit nil

### UnsetContentSize
`func (o *File) UnsetContentSize()`

UnsetContentSize ensures that no value is present for ContentSize, not even an explicit nil
### GetParentFolderId

`func (o *File) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *File) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *File) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### SetParentFolderIdNil

`func (o *File) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *File) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


