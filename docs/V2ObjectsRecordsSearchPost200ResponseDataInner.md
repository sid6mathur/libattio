# V2ObjectsRecordsSearchPost200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId.md) |  | 
**RecordText** | **string** | A human-readable label for the record. Present on records from all objects. | 
**RecordImage** | **string** | The image for the record. | 
**ObjectSlug** | **string** | The slug of the object this record belongs to. | 
**EmailAddresses** | **[]string** | The person&#39;s email addresses. | 
**PhoneNumbers** | **[]string** | The person&#39;s phone numbers. | 
**Domains** | **[]string** | The company&#39;s domains. | 

## Methods

### NewV2ObjectsRecordsSearchPost200ResponseDataInner

`func NewV2ObjectsRecordsSearchPost200ResponseDataInner(id V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, recordText string, recordImage string, objectSlug string, emailAddresses []string, phoneNumbers []string, domains []string, ) *V2ObjectsRecordsSearchPost200ResponseDataInner`

NewV2ObjectsRecordsSearchPost200ResponseDataInner instantiates a new V2ObjectsRecordsSearchPost200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsRecordsSearchPost200ResponseDataInnerWithDefaults

`func NewV2ObjectsRecordsSearchPost200ResponseDataInnerWithDefaults() *V2ObjectsRecordsSearchPost200ResponseDataInner`

NewV2ObjectsRecordsSearchPost200ResponseDataInnerWithDefaults instantiates a new V2ObjectsRecordsSearchPost200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetId() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetIdOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) SetId(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId)`

SetId sets Id field to given value.


### GetRecordText

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetRecordText() string`

GetRecordText returns the RecordText field if non-nil, zero value otherwise.

### GetRecordTextOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetRecordTextOk() (*string, bool)`

GetRecordTextOk returns a tuple with the RecordText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordText

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) SetRecordText(v string)`

SetRecordText sets RecordText field to given value.


### GetRecordImage

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetRecordImage() string`

GetRecordImage returns the RecordImage field if non-nil, zero value otherwise.

### GetRecordImageOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetRecordImageOk() (*string, bool)`

GetRecordImageOk returns a tuple with the RecordImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordImage

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) SetRecordImage(v string)`

SetRecordImage sets RecordImage field to given value.


### GetObjectSlug

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetObjectSlug() string`

GetObjectSlug returns the ObjectSlug field if non-nil, zero value otherwise.

### GetObjectSlugOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetObjectSlugOk() (*string, bool)`

GetObjectSlugOk returns a tuple with the ObjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlug

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) SetObjectSlug(v string)`

SetObjectSlug sets ObjectSlug field to given value.


### GetEmailAddresses

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.


### GetPhoneNumbers

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetDomains

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *V2ObjectsRecordsSearchPost200ResponseDataInner) SetDomains(v []string)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


