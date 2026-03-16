# ConnectedFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | The object slug or ID. | 
**RecordId** | **string** | The ID of the record to create the file entry on. | 
**StorageProvider** | **string** | The external storage provider. | 
**ExternalProviderFileId** | **string** | The ID of the file or folder in the external storage provider. | 
**MicrosoftDriveId** | Pointer to **NullableString** | Microsoft drive ID. Only used when &#x60;storage_provider&#x60; is &#x60;microsoft-onedrive&#x60;. | [optional] 
**FileType** | **string** | Creates a connected folder entry. | 

## Methods

### NewConnectedFolder

`func NewConnectedFolder(object string, recordId string, storageProvider string, externalProviderFileId string, fileType string, ) *ConnectedFolder`

NewConnectedFolder instantiates a new ConnectedFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedFolderWithDefaults

`func NewConnectedFolderWithDefaults() *ConnectedFolder`

NewConnectedFolderWithDefaults instantiates a new ConnectedFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ConnectedFolder) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ConnectedFolder) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ConnectedFolder) SetObject(v string)`

SetObject sets Object field to given value.


### GetRecordId

`func (o *ConnectedFolder) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *ConnectedFolder) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *ConnectedFolder) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStorageProvider

`func (o *ConnectedFolder) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ConnectedFolder) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ConnectedFolder) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetExternalProviderFileId

`func (o *ConnectedFolder) GetExternalProviderFileId() string`

GetExternalProviderFileId returns the ExternalProviderFileId field if non-nil, zero value otherwise.

### GetExternalProviderFileIdOk

`func (o *ConnectedFolder) GetExternalProviderFileIdOk() (*string, bool)`

GetExternalProviderFileIdOk returns a tuple with the ExternalProviderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderFileId

`func (o *ConnectedFolder) SetExternalProviderFileId(v string)`

SetExternalProviderFileId sets ExternalProviderFileId field to given value.


### GetMicrosoftDriveId

`func (o *ConnectedFolder) GetMicrosoftDriveId() string`

GetMicrosoftDriveId returns the MicrosoftDriveId field if non-nil, zero value otherwise.

### GetMicrosoftDriveIdOk

`func (o *ConnectedFolder) GetMicrosoftDriveIdOk() (*string, bool)`

GetMicrosoftDriveIdOk returns a tuple with the MicrosoftDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftDriveId

`func (o *ConnectedFolder) SetMicrosoftDriveId(v string)`

SetMicrosoftDriveId sets MicrosoftDriveId field to given value.

### HasMicrosoftDriveId

`func (o *ConnectedFolder) HasMicrosoftDriveId() bool`

HasMicrosoftDriveId returns a boolean if a field has been set.

### SetMicrosoftDriveIdNil

`func (o *ConnectedFolder) SetMicrosoftDriveIdNil(b bool)`

 SetMicrosoftDriveIdNil sets the value for MicrosoftDriveId to be an explicit nil

### UnsetMicrosoftDriveId
`func (o *ConnectedFolder) UnsetMicrosoftDriveId()`

UnsetMicrosoftDriveId ensures that no value is present for MicrosoftDriveId, not even an explicit nil
### GetFileType

`func (o *ConnectedFolder) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ConnectedFolder) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ConnectedFolder) SetFileType(v string)`

SetFileType sets FileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


