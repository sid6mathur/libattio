# V2MeetingsPostRequestDataParticipantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the participant. New person records and companies will automatically be created based upon the email address values provided. | 
**IsOrganizer** | [**V2MeetingsPostRequestDataParticipantsInnerIsOrganizer**](V2MeetingsPostRequestDataParticipantsInnerIsOrganizer.md) |  | 
**Status** | **string** | The status of the individual meeting participant. | 

## Methods

### NewV2MeetingsPostRequestDataParticipantsInner

`func NewV2MeetingsPostRequestDataParticipantsInner(emailAddress string, isOrganizer V2MeetingsPostRequestDataParticipantsInnerIsOrganizer, status string, ) *V2MeetingsPostRequestDataParticipantsInner`

NewV2MeetingsPostRequestDataParticipantsInner instantiates a new V2MeetingsPostRequestDataParticipantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataParticipantsInnerWithDefaults

`func NewV2MeetingsPostRequestDataParticipantsInnerWithDefaults() *V2MeetingsPostRequestDataParticipantsInner`

NewV2MeetingsPostRequestDataParticipantsInnerWithDefaults instantiates a new V2MeetingsPostRequestDataParticipantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *V2MeetingsPostRequestDataParticipantsInner) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *V2MeetingsPostRequestDataParticipantsInner) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *V2MeetingsPostRequestDataParticipantsInner) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetIsOrganizer

`func (o *V2MeetingsPostRequestDataParticipantsInner) GetIsOrganizer() V2MeetingsPostRequestDataParticipantsInnerIsOrganizer`

GetIsOrganizer returns the IsOrganizer field if non-nil, zero value otherwise.

### GetIsOrganizerOk

`func (o *V2MeetingsPostRequestDataParticipantsInner) GetIsOrganizerOk() (*V2MeetingsPostRequestDataParticipantsInnerIsOrganizer, bool)`

GetIsOrganizerOk returns a tuple with the IsOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizer

`func (o *V2MeetingsPostRequestDataParticipantsInner) SetIsOrganizer(v V2MeetingsPostRequestDataParticipantsInnerIsOrganizer)`

SetIsOrganizer sets IsOrganizer field to given value.


### GetStatus

`func (o *V2MeetingsPostRequestDataParticipantsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2MeetingsPostRequestDataParticipantsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2MeetingsPostRequestDataParticipantsInner) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


