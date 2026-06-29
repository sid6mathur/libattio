# V2FilesPost200ResponseData

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
**ParentFolderId** | **string** | The ID of the parent folder, or null if this is a top-level folder. | 
**ExternalProviderFileId** | **string** | The file ID in the external storage provider. | 
**MicrosoftDriveId** | **string** | Microsoft drive ID. This field is only populated for &#x60;microsoft-onedrive&#x60; entries. | 

## Methods

### NewV2FilesPost200ResponseData

`func NewV2FilesPost200ResponseData(id FileId, objectId string, objectSlug string, recordId string, storageProvider string, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, createdAt string, fileType string, name string, parentFolderId string, externalProviderFileId string, microsoftDriveId string, ) *V2FilesPost200ResponseData`

NewV2FilesPost200ResponseData instantiates a new V2FilesPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FilesPost200ResponseDataWithDefaults

`func NewV2FilesPost200ResponseDataWithDefaults() *V2FilesPost200ResponseData`

NewV2FilesPost200ResponseDataWithDefaults instantiates a new V2FilesPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2FilesPost200ResponseData) GetId() FileId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2FilesPost200ResponseData) GetIdOk() (*FileId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2FilesPost200ResponseData) SetId(v FileId)`

SetId sets Id field to given value.


### GetObjectId

`func (o *V2FilesPost200ResponseData) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *V2FilesPost200ResponseData) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *V2FilesPost200ResponseData) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectSlug

`func (o *V2FilesPost200ResponseData) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *V2FilesPost200ResponseData) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *V2FilesPost200ResponseData) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetRecordId

`func (o *V2FilesPost200ResponseData) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *V2FilesPost200ResponseData) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *V2FilesPost200ResponseData) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStorageProvider

`func (o *V2FilesPost200ResponseData) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *V2FilesPost200ResponseData) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *V2FilesPost200ResponseData) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetCreatedByActor

`func (o *V2FilesPost200ResponseData) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2FilesPost200ResponseData) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2FilesPost200ResponseData) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *V2FilesPost200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2FilesPost200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2FilesPost200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileType

`func (o *V2FilesPost200ResponseData) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *V2FilesPost200ResponseData) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *V2FilesPost200ResponseData) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetName

`func (o *V2FilesPost200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2FilesPost200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2FilesPost200ResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *V2FilesPost200ResponseData) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *V2FilesPost200ResponseData) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *V2FilesPost200ResponseData) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### GetExternalProviderFileId

`func (o *V2FilesPost200ResponseData) GetExternalProviderFileId() string`

GetExternalProviderFileId returns the ExternalProviderFileId field if non-nil, zero value otherwise.

### GetExternalProviderFileIdOk

`func (o *V2FilesPost200ResponseData) GetExternalProviderFileIdOk() (*string, bool)`

GetExternalProviderFileIdOk returns a tuple with the ExternalProviderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderFileId

`func (o *V2FilesPost200ResponseData) SetExternalProviderFileId(v string)`

SetExternalProviderFileId sets ExternalProviderFileId field to given value.


### GetMicrosoftDriveId

`func (o *V2FilesPost200ResponseData) GetMicrosoftDriveId() string`

GetMicrosoftDriveId returns the MicrosoftDriveId field if non-nil, zero value otherwise.

### GetMicrosoftDriveIdOk

`func (o *V2FilesPost200ResponseData) GetMicrosoftDriveIdOk() (*string, bool)`

GetMicrosoftDriveIdOk returns a tuple with the MicrosoftDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftDriveId

`func (o *V2FilesPost200ResponseData) SetMicrosoftDriveId(v string)`

SetMicrosoftDriveId sets MicrosoftDriveId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


