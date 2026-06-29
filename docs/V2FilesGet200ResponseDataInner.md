# V2FilesGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**FileId**](FileId.md) |  | 
**ObjectId** | **string** | The ID of the object the record belongs to. | 
**ObjectSlug** | **string** | The slug of the object the record belongs to. | 
**RecordId** | **string** | The ID of the record the file is linked to. | 
**StorageProvider** | **string** | The storage provider for this file entry. | 
**CreatedByActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor.md) |  | 
**CreatedAt** | **string** | Timestamp representing when the file entry was created. | 
**FileType** | **string** | The type of file entry. | 
**Name** | **string** | The name of the folder. | 
**ContentType** | **string** | The content type of the file. | 
**ContentSize** | **float32** | The size of the file in bytes. | 
**ParentFolderId** | **string** | The ID of the parent folder, or null if this is a top-level folder. | 
**ExternalProviderFileId** | **string** | The file ID in the external storage provider. | 
**MicrosoftDriveId** | **string** | Microsoft drive ID. This field is only populated for &#x60;microsoft-onedrive&#x60; entries. | 

## Methods

### NewV2FilesGet200ResponseDataInner

`func NewV2FilesGet200ResponseDataInner(id FileId, objectId string, objectSlug string, recordId string, storageProvider string, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, createdAt string, fileType string, name string, contentType string, contentSize float32, parentFolderId string, externalProviderFileId string, microsoftDriveId string, ) *V2FilesGet200ResponseDataInner`

NewV2FilesGet200ResponseDataInner instantiates a new V2FilesGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FilesGet200ResponseDataInnerWithDefaults

`func NewV2FilesGet200ResponseDataInnerWithDefaults() *V2FilesGet200ResponseDataInner`

NewV2FilesGet200ResponseDataInnerWithDefaults instantiates a new V2FilesGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2FilesGet200ResponseDataInner) GetId() FileId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2FilesGet200ResponseDataInner) GetIdOk() (*FileId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2FilesGet200ResponseDataInner) SetId(v FileId)`

SetId sets Id field to given value.


### GetObjectId

`func (o *V2FilesGet200ResponseDataInner) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *V2FilesGet200ResponseDataInner) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *V2FilesGet200ResponseDataInner) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectSlug

`func (o *V2FilesGet200ResponseDataInner) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *V2FilesGet200ResponseDataInner) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *V2FilesGet200ResponseDataInner) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetRecordId

`func (o *V2FilesGet200ResponseDataInner) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *V2FilesGet200ResponseDataInner) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *V2FilesGet200ResponseDataInner) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStorageProvider

`func (o *V2FilesGet200ResponseDataInner) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *V2FilesGet200ResponseDataInner) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *V2FilesGet200ResponseDataInner) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetCreatedByActor

`func (o *V2FilesGet200ResponseDataInner) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2FilesGet200ResponseDataInner) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2FilesGet200ResponseDataInner) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *V2FilesGet200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2FilesGet200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2FilesGet200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileType

`func (o *V2FilesGet200ResponseDataInner) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *V2FilesGet200ResponseDataInner) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *V2FilesGet200ResponseDataInner) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetName

`func (o *V2FilesGet200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2FilesGet200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2FilesGet200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetContentType

`func (o *V2FilesGet200ResponseDataInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *V2FilesGet200ResponseDataInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *V2FilesGet200ResponseDataInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentSize

`func (o *V2FilesGet200ResponseDataInner) GetContentSize() float32`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *V2FilesGet200ResponseDataInner) GetContentSizeOk() (*float32, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *V2FilesGet200ResponseDataInner) SetContentSize(v float32)`

SetContentSize sets ContentSize field to given value.


### GetParentFolderId

`func (o *V2FilesGet200ResponseDataInner) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *V2FilesGet200ResponseDataInner) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *V2FilesGet200ResponseDataInner) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### GetExternalProviderFileId

`func (o *V2FilesGet200ResponseDataInner) GetExternalProviderFileId() string`

GetExternalProviderFileId returns the ExternalProviderFileId field if non-nil, zero value otherwise.

### GetExternalProviderFileIdOk

`func (o *V2FilesGet200ResponseDataInner) GetExternalProviderFileIdOk() (*string, bool)`

GetExternalProviderFileIdOk returns a tuple with the ExternalProviderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderFileId

`func (o *V2FilesGet200ResponseDataInner) SetExternalProviderFileId(v string)`

SetExternalProviderFileId sets ExternalProviderFileId field to given value.


### GetMicrosoftDriveId

`func (o *V2FilesGet200ResponseDataInner) GetMicrosoftDriveId() string`

GetMicrosoftDriveId returns the MicrosoftDriveId field if non-nil, zero value otherwise.

### GetMicrosoftDriveIdOk

`func (o *V2FilesGet200ResponseDataInner) GetMicrosoftDriveIdOk() (*string, bool)`

GetMicrosoftDriveIdOk returns a tuple with the MicrosoftDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftDriveId

`func (o *V2FilesGet200ResponseDataInner) SetMicrosoftDriveId(v string)`

SetMicrosoftDriveId sets MicrosoftDriveId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


