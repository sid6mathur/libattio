# V2MeetingsPostRequestDataExternalRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IcalUid** | **string** | The ical uid of the meeting. | 
**Provider** | **string** | The email provider used to sync the meeting. | 
**OriginalStartTime** | Pointer to **string** | The original start time of the meeting. Use a timestamp with a specified offset for all day and non-all day meetings. This property is required for recurring event exceptions and optional otherwise. | [optional] 
**IsRecurring** | **bool** | Whether or not the meeting is recurring. | 

## Methods

### NewV2MeetingsPostRequestDataExternalRef

`func NewV2MeetingsPostRequestDataExternalRef(icalUid string, provider string, isRecurring bool, ) *V2MeetingsPostRequestDataExternalRef`

NewV2MeetingsPostRequestDataExternalRef instantiates a new V2MeetingsPostRequestDataExternalRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataExternalRefWithDefaults

`func NewV2MeetingsPostRequestDataExternalRefWithDefaults() *V2MeetingsPostRequestDataExternalRef`

NewV2MeetingsPostRequestDataExternalRefWithDefaults instantiates a new V2MeetingsPostRequestDataExternalRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcalUid

`func (o *V2MeetingsPostRequestDataExternalRef) GetIcalUid() string`

GetIcalUid returns the IcalUid field if non-nil, zero value otherwise.

### GetIcalUidOk

`func (o *V2MeetingsPostRequestDataExternalRef) GetIcalUidOk() (*string, bool)`

GetIcalUidOk returns a tuple with the IcalUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcalUid

`func (o *V2MeetingsPostRequestDataExternalRef) SetIcalUid(v string)`

SetIcalUid sets IcalUid field to given value.


### GetProvider

`func (o *V2MeetingsPostRequestDataExternalRef) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V2MeetingsPostRequestDataExternalRef) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V2MeetingsPostRequestDataExternalRef) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOriginalStartTime

`func (o *V2MeetingsPostRequestDataExternalRef) GetOriginalStartTime() string`

GetOriginalStartTime returns the OriginalStartTime field if non-nil, zero value otherwise.

### GetOriginalStartTimeOk

`func (o *V2MeetingsPostRequestDataExternalRef) GetOriginalStartTimeOk() (*string, bool)`

GetOriginalStartTimeOk returns a tuple with the OriginalStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartTime

`func (o *V2MeetingsPostRequestDataExternalRef) SetOriginalStartTime(v string)`

SetOriginalStartTime sets OriginalStartTime field to given value.

### HasOriginalStartTime

`func (o *V2MeetingsPostRequestDataExternalRef) HasOriginalStartTime() bool`

HasOriginalStartTime returns a boolean if a field has been set.

### GetIsRecurring

`func (o *V2MeetingsPostRequestDataExternalRef) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *V2MeetingsPostRequestDataExternalRef) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *V2MeetingsPostRequestDataExternalRef) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


