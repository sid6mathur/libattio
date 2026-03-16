# ConnectedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | The object slug or ID. | 
**RecordId** | **string** | The ID of the record to create the file entry on. | 
**StorageProvider** | **string** | The external storage provider. | 
**ExternalProviderFileId** | **string** | The ID of the file or folder in the external storage provider. | 
**MicrosoftDriveId** | Pointer to **NullableString** | Microsoft drive ID. Only used when &#x60;storage_provider&#x60; is &#x60;microsoft-onedrive&#x60;. | [optional] 
**FileType** | **string** | Creates a connected file entry. | 

## Methods

### NewConnectedFile

`func NewConnectedFile(object string, recordId string, storageProvider string, externalProviderFileId string, fileType string, ) *ConnectedFile`

NewConnectedFile instantiates a new ConnectedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedFileWithDefaults

`func NewConnectedFileWithDefaults() *ConnectedFile`

NewConnectedFileWithDefaults instantiates a new ConnectedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ConnectedFile) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ConnectedFile) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ConnectedFile) SetObject(v string)`

SetObject sets Object field to given value.


### GetRecordId

`func (o *ConnectedFile) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *ConnectedFile) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *ConnectedFile) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStorageProvider

`func (o *ConnectedFile) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ConnectedFile) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ConnectedFile) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.


### GetExternalProviderFileId

`func (o *ConnectedFile) GetExternalProviderFileId() string`

GetExternalProviderFileId returns the ExternalProviderFileId field if non-nil, zero value otherwise.

### GetExternalProviderFileIdOk

`func (o *ConnectedFile) GetExternalProviderFileIdOk() (*string, bool)`

GetExternalProviderFileIdOk returns a tuple with the ExternalProviderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderFileId

`func (o *ConnectedFile) SetExternalProviderFileId(v string)`

SetExternalProviderFileId sets ExternalProviderFileId field to given value.


### GetMicrosoftDriveId

`func (o *ConnectedFile) GetMicrosoftDriveId() string`

GetMicrosoftDriveId returns the MicrosoftDriveId field if non-nil, zero value otherwise.

### GetMicrosoftDriveIdOk

`func (o *ConnectedFile) GetMicrosoftDriveIdOk() (*string, bool)`

GetMicrosoftDriveIdOk returns a tuple with the MicrosoftDriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftDriveId

`func (o *ConnectedFile) SetMicrosoftDriveId(v string)`

SetMicrosoftDriveId sets MicrosoftDriveId field to given value.

### HasMicrosoftDriveId

`func (o *ConnectedFile) HasMicrosoftDriveId() bool`

HasMicrosoftDriveId returns a boolean if a field has been set.

### SetMicrosoftDriveIdNil

`func (o *ConnectedFile) SetMicrosoftDriveIdNil(b bool)`

 SetMicrosoftDriveIdNil sets the value for MicrosoftDriveId to be an explicit nil

### UnsetMicrosoftDriveId
`func (o *ConnectedFile) UnsetMicrosoftDriveId()`

UnsetMicrosoftDriveId ensures that no value is present for MicrosoftDriveId, not even an explicit nil
### GetFileType

`func (o *ConnectedFile) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ConnectedFile) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ConnectedFile) SetFileType(v string)`

SetFileType sets FileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


