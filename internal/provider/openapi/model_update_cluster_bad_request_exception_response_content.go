/*
ParallelCluster

ParallelCluster API

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateClusterBadRequestExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateClusterBadRequestExceptionResponseContent{}

// UpdateClusterBadRequestExceptionResponseContent This exception is thrown when a client calls the UpdateCluster API with an invalid request. This includes an error due to invalid cluster configuration and unsupported update.
type UpdateClusterBadRequestExceptionResponseContent struct {
	Message *string `json:"message,omitempty"`
	ConfigurationValidationErrors []ConfigValidationMessage `json:"configurationValidationErrors,omitempty"`
	UpdateValidationErrors []UpdateError `json:"updateValidationErrors,omitempty"`
	ChangeSet []Change `json:"changeSet,omitempty"`
}

// NewUpdateClusterBadRequestExceptionResponseContent instantiates a new UpdateClusterBadRequestExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClusterBadRequestExceptionResponseContent() *UpdateClusterBadRequestExceptionResponseContent {
	this := UpdateClusterBadRequestExceptionResponseContent{}
	return &this
}

// NewUpdateClusterBadRequestExceptionResponseContentWithDefaults instantiates a new UpdateClusterBadRequestExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClusterBadRequestExceptionResponseContentWithDefaults() *UpdateClusterBadRequestExceptionResponseContent {
	this := UpdateClusterBadRequestExceptionResponseContent{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateClusterBadRequestExceptionResponseContent) SetMessage(v string) {
	o.Message = &v
}

// GetConfigurationValidationErrors returns the ConfigurationValidationErrors field value if set, zero value otherwise.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetConfigurationValidationErrors() []ConfigValidationMessage {
	if o == nil || IsNil(o.ConfigurationValidationErrors) {
		var ret []ConfigValidationMessage
		return ret
	}
	return o.ConfigurationValidationErrors
}

// GetConfigurationValidationErrorsOk returns a tuple with the ConfigurationValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetConfigurationValidationErrorsOk() ([]ConfigValidationMessage, bool) {
	if o == nil || IsNil(o.ConfigurationValidationErrors) {
		return nil, false
	}
	return o.ConfigurationValidationErrors, true
}

// HasConfigurationValidationErrors returns a boolean if a field has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) HasConfigurationValidationErrors() bool {
	if o != nil && !IsNil(o.ConfigurationValidationErrors) {
		return true
	}

	return false
}

// SetConfigurationValidationErrors gets a reference to the given []ConfigValidationMessage and assigns it to the ConfigurationValidationErrors field.
func (o *UpdateClusterBadRequestExceptionResponseContent) SetConfigurationValidationErrors(v []ConfigValidationMessage) {
	o.ConfigurationValidationErrors = v
}

// GetUpdateValidationErrors returns the UpdateValidationErrors field value if set, zero value otherwise.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetUpdateValidationErrors() []UpdateError {
	if o == nil || IsNil(o.UpdateValidationErrors) {
		var ret []UpdateError
		return ret
	}
	return o.UpdateValidationErrors
}

// GetUpdateValidationErrorsOk returns a tuple with the UpdateValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetUpdateValidationErrorsOk() ([]UpdateError, bool) {
	if o == nil || IsNil(o.UpdateValidationErrors) {
		return nil, false
	}
	return o.UpdateValidationErrors, true
}

// HasUpdateValidationErrors returns a boolean if a field has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) HasUpdateValidationErrors() bool {
	if o != nil && !IsNil(o.UpdateValidationErrors) {
		return true
	}

	return false
}

// SetUpdateValidationErrors gets a reference to the given []UpdateError and assigns it to the UpdateValidationErrors field.
func (o *UpdateClusterBadRequestExceptionResponseContent) SetUpdateValidationErrors(v []UpdateError) {
	o.UpdateValidationErrors = v
}

// GetChangeSet returns the ChangeSet field value if set, zero value otherwise.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetChangeSet() []Change {
	if o == nil || IsNil(o.ChangeSet) {
		var ret []Change
		return ret
	}
	return o.ChangeSet
}

// GetChangeSetOk returns a tuple with the ChangeSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) GetChangeSetOk() ([]Change, bool) {
	if o == nil || IsNil(o.ChangeSet) {
		return nil, false
	}
	return o.ChangeSet, true
}

// HasChangeSet returns a boolean if a field has been set.
func (o *UpdateClusterBadRequestExceptionResponseContent) HasChangeSet() bool {
	if o != nil && !IsNil(o.ChangeSet) {
		return true
	}

	return false
}

// SetChangeSet gets a reference to the given []Change and assigns it to the ChangeSet field.
func (o *UpdateClusterBadRequestExceptionResponseContent) SetChangeSet(v []Change) {
	o.ChangeSet = v
}

func (o UpdateClusterBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateClusterBadRequestExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ConfigurationValidationErrors) {
		toSerialize["configurationValidationErrors"] = o.ConfigurationValidationErrors
	}
	if !IsNil(o.UpdateValidationErrors) {
		toSerialize["updateValidationErrors"] = o.UpdateValidationErrors
	}
	if !IsNil(o.ChangeSet) {
		toSerialize["changeSet"] = o.ChangeSet
	}
	return toSerialize, nil
}

type NullableUpdateClusterBadRequestExceptionResponseContent struct {
	value *UpdateClusterBadRequestExceptionResponseContent
	isSet bool
}

func (v NullableUpdateClusterBadRequestExceptionResponseContent) Get() *UpdateClusterBadRequestExceptionResponseContent {
	return v.value
}

func (v *NullableUpdateClusterBadRequestExceptionResponseContent) Set(val *UpdateClusterBadRequestExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClusterBadRequestExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClusterBadRequestExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClusterBadRequestExceptionResponseContent(val *UpdateClusterBadRequestExceptionResponseContent) *NullableUpdateClusterBadRequestExceptionResponseContent {
	return &NullableUpdateClusterBadRequestExceptionResponseContent{value: val, isSet: true}
}

func (v NullableUpdateClusterBadRequestExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClusterBadRequestExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

