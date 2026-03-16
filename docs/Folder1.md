# Folder1

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
**ParentFolderId** | **string** | The ID of the parent folder, or null if this is a top-level folder. | 
**HasChildren** | **bool** | Whether the folder contains any child entries. | 

## Methods

### NewFolder1

`func NewFolder1(id FileId, objectId string, objectSlug string, recordId string, storageProvider string, createdByActor FileCreatedByActor, createdAt string, fileType string, name string, parentFolderId string, hasChildren bool, ) *Folder1`

NewFolder1 instantiates a new Folder1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolder1WithDefaults

`func NewFolder1WithDefaults() *Folder1`

NewFolder1WithDefaults instantiates a new Folder1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Folder1) GetId() FileId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Folder1) GetIdOk() (*FileId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Folder1) SetId(v FileId)`

SetId sets Id field to given value.


### GetObjectId

`func (o *Folder1) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *Folder1) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *Folder1) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectSlug

`func (o *Folder1) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *Folder1) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *Folder1) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetRecordId

`func (o *Folder1) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *Folder1) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *Folder1) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStorageProvider

`func (o *Folder1) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *Folder1) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *Folder1) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetCreatedByActor

`func (o *Folder1) GetCreatedByActor() FileCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *Folder1) GetCreatedByActorOk() (*FileCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *Folder1) SetCreatedByActor(v FileCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *Folder1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Folder1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Folder1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileType

`func (o *Folder1) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *Folder1) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *Folder1) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetName

`func (o *Folder1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Folder1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Folder1) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *Folder1) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Folder1) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *Folder1) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### GetHasChildren

`func (o *Folder1) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *Folder1) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *Folder1) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


