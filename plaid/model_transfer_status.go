/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.44.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TransferStatus The status of the transfer.
type TransferStatus string

// List of TransferStatus
const (
	TRANSFERSTATUS_PENDING TransferStatus = "pending"
	TRANSFERSTATUS_POSTED TransferStatus = "posted"
	TRANSFERSTATUS_CANCELLED TransferStatus = "cancelled"
	TRANSFERSTATUS_FAILED TransferStatus = "failed"
	TRANSFERSTATUS_REVERSED TransferStatus = "reversed"
)

var allowedTransferStatusEnumValues = []TransferStatus{
	"pending",
	"posted",
	"cancelled",
	"failed",
	"reversed",
}

func (v *TransferStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferStatus(value)
	for _, existing := range allowedTransferStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferStatus", value)
}

// NewTransferStatusFromValue returns a pointer to a valid TransferStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferStatusFromValue(v string) (*TransferStatus, error) {
	ev := TransferStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferStatus: valid values are %v", v, allowedTransferStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferStatus) IsValid() bool {
	for _, existing := range allowedTransferStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferStatus value
func (v TransferStatus) Ptr() *TransferStatus {
	return &v
}

type NullableTransferStatus struct {
	value *TransferStatus
	isSet bool
}

func (v NullableTransferStatus) Get() *TransferStatus {
	return v.value
}

func (v *NullableTransferStatus) Set(val *TransferStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferStatus(val *TransferStatus) *NullableTransferStatus {
	return &NullableTransferStatus{value: val, isSet: true}
}

func (v NullableTransferStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

