# V2MeetingsPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the meeting. | 
**Description** | **string** | The description of the meeting. | 
**Start** | [**V2MeetingsPostRequestDataStart**](V2MeetingsPostRequestDataStart.md) |  | 
**End** | [**V2MeetingsPostRequestDataEnd**](V2MeetingsPostRequestDataEnd.md) |  | 
**IsAllDay** | **bool** | Whether or not the meeting is an all day event. All day events may span multiple days. When true, start and end must use date format. When false, start and end must use datetime with timezone format. | 
**Participants** | [**[]V2MeetingsPostRequestDataParticipantsInner**](V2MeetingsPostRequestDataParticipantsInner.md) |  | 
**LinkedRecords** | Pointer to [**[]V2MeetingsPostRequestDataLinkedRecordsInner**](V2MeetingsPostRequestDataLinkedRecordsInner.md) | A list of records to link to the meeting. Each record is specified by its object (slug or UUID) and record ID (UUID). Attio will automatically link the meeting participants&#39; companies to the meeting; this behavior is asynchronous. | [optional] 
**ExternalRef** | [**V2MeetingsPostRequestDataExternalRef**](V2MeetingsPostRequestDataExternalRef.md) |  | 

## Methods

### NewV2MeetingsPostRequestData

`func NewV2MeetingsPostRequestData(title string, description string, start V2MeetingsPostRequestDataStart, end V2MeetingsPostRequestDataEnd, isAllDay bool, participants []V2MeetingsPostRequestDataParticipantsInner, externalRef V2MeetingsPostRequestDataExternalRef, ) *V2MeetingsPostRequestData`

NewV2MeetingsPostRequestData instantiates a new V2MeetingsPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataWithDefaults

`func NewV2MeetingsPostRequestDataWithDefaults() *V2MeetingsPostRequestData`

NewV2MeetingsPostRequestDataWithDefaults instantiates a new V2MeetingsPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V2MeetingsPostRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2MeetingsPostRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2MeetingsPostRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *V2MeetingsPostRequestData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2MeetingsPostRequestData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2MeetingsPostRequestData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStart

`func (o *V2MeetingsPostRequestData) GetStart() V2MeetingsPostRequestDataStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V2MeetingsPostRequestData) GetStartOk() (*V2MeetingsPostRequestDataStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V2MeetingsPostRequestData) SetStart(v V2MeetingsPostRequestDataStart)`

SetStart sets Start field to given value.


### GetEnd

`func (o *V2MeetingsPostRequestData) GetEnd() V2MeetingsPostRequestDataEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V2MeetingsPostRequestData) GetEndOk() (*V2MeetingsPostRequestDataEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V2MeetingsPostRequestData) SetEnd(v V2MeetingsPostRequestDataEnd)`

SetEnd sets End field to given value.


### GetIsAllDay

`func (o *V2MeetingsPostRequestData) GetIsAllDay() bool`

GetIsAllDay returns the IsAllDay field if non-nil, zero value otherwise.

### GetIsAllDayOk

`func (o *V2MeetingsPostRequestData) GetIsAllDayOk() (*bool, bool)`

GetIsAllDayOk returns a tuple with the IsAllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllDay

`func (o *V2MeetingsPostRequestData) SetIsAllDay(v bool)`

SetIsAllDay sets IsAllDay field to given value.


### GetParticipants

`func (o *V2MeetingsPostRequestData) GetParticipants() []V2MeetingsPostRequestDataParticipantsInner`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *V2MeetingsPostRequestData) GetParticipantsOk() (*[]V2MeetingsPostRequestDataParticipantsInner, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *V2MeetingsPostRequestData) SetParticipants(v []V2MeetingsPostRequestDataParticipantsInner)`

SetParticipants sets Participants field to given value.


### GetLinkedRecords

`func (o *V2MeetingsPostRequestData) GetLinkedRecords() []V2MeetingsPostRequestDataLinkedRecordsInner`

GetLinkedRecords returns the LinkedRecords field if non-nil, zero value otherwise.

### GetLinkedRecordsOk

`func (o *V2MeetingsPostRequestData) GetLinkedRecordsOk() (*[]V2MeetingsPostRequestDataLinkedRecordsInner, bool)`

GetLinkedRecordsOk returns a tuple with the LinkedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecords

`func (o *V2MeetingsPostRequestData) SetLinkedRecords(v []V2MeetingsPostRequestDataLinkedRecordsInner)`

SetLinkedRecords sets LinkedRecords field to given value.

### HasLinkedRecords

`func (o *V2MeetingsPostRequestData) HasLinkedRecords() bool`

HasLinkedRecords returns a boolean if a field has been set.

### GetExternalRef

`func (o *V2MeetingsPostRequestData) GetExternalRef() V2MeetingsPostRequestDataExternalRef`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *V2MeetingsPostRequestData) GetExternalRefOk() (*V2MeetingsPostRequestDataExternalRef, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *V2MeetingsPostRequestData) SetExternalRef(v V2MeetingsPostRequestDataExternalRef)`

SetExternalRef sets ExternalRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


