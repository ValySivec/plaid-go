/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.39.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxTransferSimulateRequest Defines the request schema for `/sandbox/transfer/simulate`
type SandboxTransferSimulateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a transfer.
	TransferId string `json:"transfer_id"`
	// The asynchronous event to be simulated. May be: `posted`, `failed`, or `reversed`.  An error will be returned if the event type is incompatible with the current transfer status. Compatible status --> event type transitions include:  `pending` --> `failed`  `pending` --> `posted`  `posted` --> `reversed` 
	EventType string `json:"event_type"`
	FailureReason NullableTransferFailure `json:"failure_reason,omitempty"`
}

// NewSandboxTransferSimulateRequest instantiates a new SandboxTransferSimulateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTransferSimulateRequest(transferId string, eventType string) *SandboxTransferSimulateRequest {
	this := SandboxTransferSimulateRequest{}
	this.TransferId = transferId
	this.EventType = eventType
	return &this
}

// NewSandboxTransferSimulateRequestWithDefaults instantiates a new SandboxTransferSimulateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTransferSimulateRequestWithDefaults() *SandboxTransferSimulateRequest {
	this := SandboxTransferSimulateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxTransferSimulateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferSimulateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxTransferSimulateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxTransferSimulateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxTransferSimulateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTransferSimulateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxTransferSimulateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxTransferSimulateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTransferId returns the TransferId field value
func (o *SandboxTransferSimulateRequest) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferSimulateRequest) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *SandboxTransferSimulateRequest) SetTransferId(v string) {
	o.TransferId = v
}

// GetEventType returns the EventType field value
func (o *SandboxTransferSimulateRequest) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *SandboxTransferSimulateRequest) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *SandboxTransferSimulateRequest) SetEventType(v string) {
	o.EventType = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxTransferSimulateRequest) GetFailureReason() TransferFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret TransferFailure
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxTransferSimulateRequest) GetFailureReasonOk() (*TransferFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *SandboxTransferSimulateRequest) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableTransferFailure and assigns it to the FailureReason field.
func (o *SandboxTransferSimulateRequest) SetFailureReason(v TransferFailure) {
	o.FailureReason.Set(&v)
}
// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *SandboxTransferSimulateRequest) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *SandboxTransferSimulateRequest) UnsetFailureReason() {
	o.FailureReason.Unset()
}

func (o SandboxTransferSimulateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if o.FailureReason.IsSet() {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxTransferSimulateRequest struct {
	value *SandboxTransferSimulateRequest
	isSet bool
}

func (v NullableSandboxTransferSimulateRequest) Get() *SandboxTransferSimulateRequest {
	return v.value
}

func (v *NullableSandboxTransferSimulateRequest) Set(val *SandboxTransferSimulateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTransferSimulateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTransferSimulateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTransferSimulateRequest(val *SandboxTransferSimulateRequest) *NullableSandboxTransferSimulateRequest {
	return &NullableSandboxTransferSimulateRequest{value: val, isSet: true}
}

func (v NullableSandboxTransferSimulateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTransferSimulateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


