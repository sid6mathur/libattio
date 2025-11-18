# V2TasksTaskIdPatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeadlineAt** | Pointer to **NullableString** | The deadline of the task, in ISO 8601 format. | [optional] 
**IsCompleted** | Pointer to **bool** | Whether the task has been completed. | [optional] 
**LinkedRecords** | Pointer to [**[]V2TasksPostRequestDataLinkedRecordsInner**](V2TasksPostRequestDataLinkedRecordsInner.md) | Records linked to the task. Creating record links within task content text is not possible via the API at present. | [optional] 
**Assignees** | Pointer to [**[]V2TasksPostRequestDataAssigneesInner**](V2TasksPostRequestDataAssigneesInner.md) | Workspace members assigned to this task. | [optional] 

## Methods

### NewV2TasksTaskIdPatchRequestData

`func NewV2TasksTaskIdPatchRequestData() *V2TasksTaskIdPatchRequestData`

NewV2TasksTaskIdPatchRequestData instantiates a new V2TasksTaskIdPatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TasksTaskIdPatchRequestDataWithDefaults

`func NewV2TasksTaskIdPatchRequestDataWithDefaults() *V2TasksTaskIdPatchRequestData`

NewV2TasksTaskIdPatchRequestDataWithDefaults instantiates a new V2TasksTaskIdPatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadlineAt

`func (o *V2TasksTaskIdPatchRequestData) GetDeadlineAt() string`

GetDeadlineAt returns the DeadlineAt field if non-nil, zero value otherwise.

### GetDeadlineAtOk

`func (o *V2TasksTaskIdPatchRequestData) GetDeadlineAtOk() (*string, bool)`

GetDeadlineAtOk returns a tuple with the DeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineAt

`func (o *V2TasksTaskIdPatchRequestData) SetDeadlineAt(v string)`

SetDeadlineAt sets DeadlineAt field to given value.

### HasDeadlineAt

`func (o *V2TasksTaskIdPatchRequestData) HasDeadlineAt() bool`

HasDeadlineAt returns a boolean if a field has been set.

### SetDeadlineAtNil

`func (o *V2TasksTaskIdPatchRequestData) SetDeadlineAtNil(b bool)`

 SetDeadlineAtNil sets the value for DeadlineAt to be an explicit nil

### UnsetDeadlineAt
`func (o *V2TasksTaskIdPatchRequestData) UnsetDeadlineAt()`

UnsetDeadlineAt ensures that no value is present for DeadlineAt, not even an explicit nil
### GetIsCompleted

`func (o *V2TasksTaskIdPatchRequestData) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *V2TasksTaskIdPatchRequestData) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *V2TasksTaskIdPatchRequestData) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *V2TasksTaskIdPatchRequestData) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### GetLinkedRecords

`func (o *V2TasksTaskIdPatchRequestData) GetLinkedRecords() []V2TasksPostRequestDataLinkedRecordsInner`

GetLinkedRecords returns the LinkedRecords field if non-nil, zero value otherwise.

### GetLinkedRecordsOk

`func (o *V2TasksTaskIdPatchRequestData) GetLinkedRecordsOk() (*[]V2TasksPostRequestDataLinkedRecordsInner, bool)`

GetLinkedRecordsOk returns a tuple with the LinkedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecords

`func (o *V2TasksTaskIdPatchRequestData) SetLinkedRecords(v []V2TasksPostRequestDataLinkedRecordsInner)`

SetLinkedRecords sets LinkedRecords field to given value.

### HasLinkedRecords

`func (o *V2TasksTaskIdPatchRequestData) HasLinkedRecords() bool`

HasLinkedRecords returns a boolean if a field has been set.

### GetAssignees

`func (o *V2TasksTaskIdPatchRequestData) GetAssignees() []V2TasksPostRequestDataAssigneesInner`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *V2TasksTaskIdPatchRequestData) GetAssigneesOk() (*[]V2TasksPostRequestDataAssigneesInner, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *V2TasksTaskIdPatchRequestData) SetAssignees(v []V2TasksPostRequestDataAssigneesInner)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *V2TasksTaskIdPatchRequestData) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


