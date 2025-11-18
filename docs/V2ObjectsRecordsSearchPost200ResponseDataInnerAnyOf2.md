# V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId.md) |  | 
**RecordText** | **string** | A human-readable label for the record. Present on records from all objects. | 
**RecordImage** | **NullableString** | The image for the record. | 
**ObjectSlug** | **string** | The slug of the object this record belongs to. | 
**Domains** | **[]string** | The company&#39;s domains. | 

## Methods

### NewV2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2

`func NewV2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2(id V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, recordText string, recordImage NullableString, objectSlug string, domains []string, ) *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2`

NewV2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2 instantiates a new V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2WithDefaults

`func NewV2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2WithDefaults() *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2`

NewV2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2WithDefaults instantiates a new V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetId() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetIdOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) SetId(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId)`

SetId sets Id field to given value.


### GetRecordText

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetRecordText() string`

GetRecordText returns the RecordText field if non-nil, zero value otherwise.

### GetRecordTextOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetRecordTextOk() (*string, bool)`

GetRecordTextOk returns a tuple with the RecordText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordText

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) SetRecordText(v string)`

SetRecordText sets RecordText field to given value.


### GetRecordImage

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetRecordImage() string`

GetRecordImage returns the RecordImage field if non-nil, zero value otherwise.

### GetRecordImageOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetRecordImageOk() (*string, bool)`

GetRecordImageOk returns a tuple with the RecordImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordImage

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) SetRecordImage(v string)`

SetRecordImage sets RecordImage field to given value.


### SetRecordImageNil

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) SetRecordImageNil(b bool)`

 SetRecordImageNil sets the value for RecordImage to be an explicit nil

### UnsetRecordImage
`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) UnsetRecordImage()`

UnsetRecordImage ensures that no value is present for RecordImage, not even an explicit nil
### GetObjectSlug

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetDomains

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInnerAnyOf2) SetDomains(v []string)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


