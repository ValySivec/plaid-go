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
)

// ProcessorBalanceGetRequest ProcessorBalanceGetRequest defines the request schema for `/processor/balance/get`
type ProcessorBalanceGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
	Options *ProcessorBalanceGetRequestOptions `json:"options,omitempty"`
}

// NewProcessorBalanceGetRequest instantiates a new ProcessorBalanceGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorBalanceGetRequest(processorToken string) *ProcessorBalanceGetRequest {
	this := ProcessorBalanceGetRequest{}
	this.ProcessorToken = processorToken
	return &this
}

// NewProcessorBalanceGetRequestWithDefaults instantiates a new ProcessorBalanceGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorBalanceGetRequestWithDefaults() *ProcessorBalanceGetRequest {
	this := ProcessorBalanceGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorBalanceGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorBalanceGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorBalanceGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorBalanceGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorBalanceGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorBalanceGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorBalanceGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorBalanceGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorBalanceGetRequest) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorBalanceGetRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorBalanceGetRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProcessorBalanceGetRequest) GetOptions() ProcessorBalanceGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret ProcessorBalanceGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorBalanceGetRequest) GetOptionsOk() (*ProcessorBalanceGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProcessorBalanceGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ProcessorBalanceGetRequestOptions and assigns it to the Options field.
func (o *ProcessorBalanceGetRequest) SetOptions(v ProcessorBalanceGetRequestOptions) {
	o.Options = &v
}

func (o ProcessorBalanceGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorBalanceGetRequest struct {
	value *ProcessorBalanceGetRequest
	isSet bool
}

func (v NullableProcessorBalanceGetRequest) Get() *ProcessorBalanceGetRequest {
	return v.value
}

func (v *NullableProcessorBalanceGetRequest) Set(val *ProcessorBalanceGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorBalanceGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorBalanceGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorBalanceGetRequest(val *ProcessorBalanceGetRequest) *NullableProcessorBalanceGetRequest {
	return &NullableProcessorBalanceGetRequest{value: val, isSet: true}
}

func (v NullableProcessorBalanceGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorBalanceGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


