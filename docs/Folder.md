# Folder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | The object slug or ID. | 
**RecordId** | **string** | The ID of the record to create the file entry on. | 
**FileType** | **string** | Creates a native Attio folder entry. | 
**Name** | **string** | The folder name. | 
**ParentFolderId** | Pointer to **string** | Optional parent folder ID. Omit to create a top-level folder. | [optional] 

## Methods

### NewFolder

`func NewFolder(object string, recordId string, fileType string, name string, ) *Folder`

NewFolder instantiates a new Folder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderWithDefaults

`func NewFolderWithDefaults() *Folder`

NewFolderWithDefaults instantiates a new Folder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Folder) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Folder) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Folder) SetObject(v string)`

SetObject sets Object field to given value.


### GetRecordId

`func (o *Folder) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *Folder) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *Folder) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetFileType

`func (o *Folder) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *Folder) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *Folder) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetName

`func (o *Folder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Folder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Folder) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *Folder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Folder) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *Folder) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *Folder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


