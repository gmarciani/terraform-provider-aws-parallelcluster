/*
ParallelCluster

ParallelCluster API

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the UpdateComputeFleetResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComputeFleetResponseContent{}

// UpdateComputeFleetResponseContent struct for UpdateComputeFleetResponseContent
type UpdateComputeFleetResponseContent struct {
	Status ComputeFleetStatus `json:"status"`
	// Timestamp representing the last status update time.
	LastStatusUpdatedTime *time.Time `json:"lastStatusUpdatedTime,omitempty"`
}

type _UpdateComputeFleetResponseContent UpdateComputeFleetResponseContent

// NewUpdateComputeFleetResponseContent instantiates a new UpdateComputeFleetResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComputeFleetResponseContent(status ComputeFleetStatus) *UpdateComputeFleetResponseContent {
	this := UpdateComputeFleetResponseContent{}
	this.Status = status
	return &this
}

// NewUpdateComputeFleetResponseContentWithDefaults instantiates a new UpdateComputeFleetResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComputeFleetResponseContentWithDefaults() *UpdateComputeFleetResponseContent {
	this := UpdateComputeFleetResponseContent{}
	return &this
}

// GetStatus returns the Status field value
func (o *UpdateComputeFleetResponseContent) GetStatus() ComputeFleetStatus {
	if o == nil {
		var ret ComputeFleetStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateComputeFleetResponseContent) GetStatusOk() (*ComputeFleetStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateComputeFleetResponseContent) SetStatus(v ComputeFleetStatus) {
	o.Status = v
}

// GetLastStatusUpdatedTime returns the LastStatusUpdatedTime field value if set, zero value otherwise.
func (o *UpdateComputeFleetResponseContent) GetLastStatusUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastStatusUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastStatusUpdatedTime
}

// GetLastStatusUpdatedTimeOk returns a tuple with the LastStatusUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputeFleetResponseContent) GetLastStatusUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastStatusUpdatedTime) {
		return nil, false
	}
	return o.LastStatusUpdatedTime, true
}

// HasLastStatusUpdatedTime returns a boolean if a field has been set.
func (o *UpdateComputeFleetResponseContent) HasLastStatusUpdatedTime() bool {
	if o != nil && !IsNil(o.LastStatusUpdatedTime) {
		return true
	}

	return false
}

// SetLastStatusUpdatedTime gets a reference to the given time.Time and assigns it to the LastStatusUpdatedTime field.
func (o *UpdateComputeFleetResponseContent) SetLastStatusUpdatedTime(v time.Time) {
	o.LastStatusUpdatedTime = &v
}

func (o UpdateComputeFleetResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComputeFleetResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.LastStatusUpdatedTime) {
		toSerialize["lastStatusUpdatedTime"] = o.LastStatusUpdatedTime
	}
	return toSerialize, nil
}

func (o *UpdateComputeFleetResponseContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateComputeFleetResponseContent := _UpdateComputeFleetResponseContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateComputeFleetResponseContent)

	if err != nil {
		return err
	}

	*o = UpdateComputeFleetResponseContent(varUpdateComputeFleetResponseContent)

	return err
}

type NullableUpdateComputeFleetResponseContent struct {
	value *UpdateComputeFleetResponseContent
	isSet bool
}

func (v NullableUpdateComputeFleetResponseContent) Get() *UpdateComputeFleetResponseContent {
	return v.value
}

func (v *NullableUpdateComputeFleetResponseContent) Set(val *UpdateComputeFleetResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComputeFleetResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComputeFleetResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComputeFleetResponseContent(val *UpdateComputeFleetResponseContent) *NullableUpdateComputeFleetResponseContent {
	return &NullableUpdateComputeFleetResponseContent{value: val, isSet: true}
}

func (v NullableUpdateComputeFleetResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComputeFleetResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

