/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransactionsRecurringGetRequest TransactionsRecurringGetRequest defines the request schema for `/transactions/recurring/get`
type TransactionsRecurringGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A list of `account_ids` to retrieve for the Item  Note: An error will be returned if a provided `account_id` is not associated with the Item.
	AccountIds []string `json:"account_ids"`
}

// NewTransactionsRecurringGetRequest instantiates a new TransactionsRecurringGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRecurringGetRequest(accessToken string, accountIds []string) *TransactionsRecurringGetRequest {
	this := TransactionsRecurringGetRequest{}
	this.AccessToken = accessToken
	this.AccountIds = accountIds
	return &this
}

// NewTransactionsRecurringGetRequestWithDefaults instantiates a new TransactionsRecurringGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRecurringGetRequestWithDefaults() *TransactionsRecurringGetRequest {
	this := TransactionsRecurringGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsRecurringGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsRecurringGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsRecurringGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsRecurringGetRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsRecurringGetRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsRecurringGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsRecurringGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsRecurringGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccountIds returns the AccountIds field value
func (o *TransactionsRecurringGetRequest) GetAccountIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringGetRequest) GetAccountIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIds, true
}

// SetAccountIds sets field value
func (o *TransactionsRecurringGetRequest) SetAccountIds(v []string) {
	o.AccountIds = v
}

func (o TransactionsRecurringGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["account_ids"] = o.AccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRecurringGetRequest struct {
	value *TransactionsRecurringGetRequest
	isSet bool
}

func (v NullableTransactionsRecurringGetRequest) Get() *TransactionsRecurringGetRequest {
	return v.value
}

func (v *NullableTransactionsRecurringGetRequest) Set(val *TransactionsRecurringGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRecurringGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRecurringGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRecurringGetRequest(val *TransactionsRecurringGetRequest) *NullableTransactionsRecurringGetRequest {
	return &NullableTransactionsRecurringGetRequest{value: val, isSet: true}
}

func (v NullableTransactionsRecurringGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRecurringGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


