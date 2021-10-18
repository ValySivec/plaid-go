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
	"time"
)

// Transfer Represents a transfer within the Transfers API.
type Transfer struct {
	// Plaid’s unique identifier for a transfer.
	Id string `json:"id"`
	AchClass ACHClass `json:"ach_class"`
	// The account ID that should be credited/debited for this transfer.
	AccountId string `json:"account_id"`
	Type TransferType `json:"type"`
	User TransferUserInResponse `json:"user"`
	// The amount of the transfer (decimal string with two digits of precision e.g. “10.00”).
	Amount string `json:"amount"`
	// The description of the transfer.
	Description string `json:"description"`
	// The datetime when this transfer was created. This will be of the form `2006-01-02T15:04:05Z`
	Created time.Time `json:"created"`
	Status TransferStatus `json:"status"`
	Network TransferNetwork `json:"network"`
	// When `true`, you can still cancel this transfer.
	Cancellable bool `json:"cancellable"`
	FailureReason NullableTransferFailure `json:"failure_reason"`
	// The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters 
	Metadata map[string]string `json:"metadata"`
	// Plaid’s unique identifier for the origination account that was used for this transfer.
	OriginationAccountId string `json:"origination_account_id"`
	AdditionalProperties map[string]interface{}
}

type _Transfer Transfer

// NewTransfer instantiates a new Transfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfer(id string, achClass ACHClass, accountId string, type_ TransferType, user TransferUserInResponse, amount string, description string, created time.Time, status TransferStatus, network TransferNetwork, cancellable bool, failureReason NullableTransferFailure, metadata map[string]string, originationAccountId string) *Transfer {
	this := Transfer{}
	this.Id = id
	this.AchClass = achClass
	this.AccountId = accountId
	this.Type = type_
	this.User = user
	this.Amount = amount
	this.Description = description
	this.Created = created
	this.Status = status
	this.Network = network
	this.Cancellable = cancellable
	this.FailureReason = failureReason
	this.Metadata = metadata
	this.OriginationAccountId = originationAccountId
	return &this
}

// NewTransferWithDefaults instantiates a new Transfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferWithDefaults() *Transfer {
	this := Transfer{}
	return &this
}

// GetId returns the Id field value
func (o *Transfer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transfer) SetId(v string) {
	o.Id = v
}

// GetAchClass returns the AchClass field value
func (o *Transfer) GetAchClass() ACHClass {
	if o == nil {
		var ret ACHClass
		return ret
	}

	return o.AchClass
}

// GetAchClassOk returns a tuple with the AchClass field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetAchClassOk() (*ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchClass, true
}

// SetAchClass sets field value
func (o *Transfer) SetAchClass(v ACHClass) {
	o.AchClass = v
}

// GetAccountId returns the AccountId field value
func (o *Transfer) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Transfer) SetAccountId(v string) {
	o.AccountId = v
}

// GetType returns the Type field value
func (o *Transfer) GetType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transfer) SetType(v TransferType) {
	o.Type = v
}

// GetUser returns the User field value
func (o *Transfer) GetUser() TransferUserInResponse {
	if o == nil {
		var ret TransferUserInResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetUserOk() (*TransferUserInResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Transfer) SetUser(v TransferUserInResponse) {
	o.User = v
}

// GetAmount returns the Amount field value
func (o *Transfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Transfer) SetAmount(v string) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *Transfer) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Transfer) SetDescription(v string) {
	o.Description = v
}

// GetCreated returns the Created field value
func (o *Transfer) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Transfer) SetCreated(v time.Time) {
	o.Created = v
}

// GetStatus returns the Status field value
func (o *Transfer) GetStatus() TransferStatus {
	if o == nil {
		var ret TransferStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetStatusOk() (*TransferStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transfer) SetStatus(v TransferStatus) {
	o.Status = v
}

// GetNetwork returns the Network field value
func (o *Transfer) GetNetwork() TransferNetwork {
	if o == nil {
		var ret TransferNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetNetworkOk() (*TransferNetwork, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *Transfer) SetNetwork(v TransferNetwork) {
	o.Network = v
}

// GetCancellable returns the Cancellable field value
func (o *Transfer) GetCancellable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cancellable
}

// GetCancellableOk returns a tuple with the Cancellable field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetCancellableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cancellable, true
}

// SetCancellable sets field value
func (o *Transfer) SetCancellable(v bool) {
	o.Cancellable = v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for TransferFailure will be returned
func (o *Transfer) GetFailureReason() TransferFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret TransferFailure
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transfer) GetFailureReasonOk() (*TransferFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *Transfer) SetFailureReason(v TransferFailure) {
	o.FailureReason.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]string will be returned
func (o *Transfer) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transfer) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Transfer) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
func (o *Transfer) GetOriginationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginationAccountId
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginationAccountId, true
}

// SetOriginationAccountId sets field value
func (o *Transfer) SetOriginationAccountId(v string) {
	o.OriginationAccountId = v
}

func (o Transfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ach_class"] = o.AchClass
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["cancellable"] = o.Cancellable
	}
	if true {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Transfer) UnmarshalJSON(bytes []byte) (err error) {
	varTransfer := _Transfer{}

	if err = json.Unmarshal(bytes, &varTransfer); err == nil {
		*o = Transfer(varTransfer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ach_class")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "description")
		delete(additionalProperties, "created")
		delete(additionalProperties, "status")
		delete(additionalProperties, "network")
		delete(additionalProperties, "cancellable")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "origination_account_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransfer struct {
	value *Transfer
	isSet bool
}

func (v NullableTransfer) Get() *Transfer {
	return v.value
}

func (v *NullableTransfer) Set(val *Transfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfer(val *Transfer) *NullableTransfer {
	return &NullableTransfer{value: val, isSet: true}
}

func (v NullableTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


