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

// TransactionStream A grouping of related transactions
type TransactionStream struct {
	// The ID of the account to which the stream belongs
	AccountId string `json:"account_id"`
	// A unique id for the stream
	StreamId string `json:"stream_id"`
	// The ID of the category to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview).
	CategoryId string `json:"category_id"`
	// A hierarchical array of the categories to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview).
	Category []string `json:"category"`
	// A description of the transaction stream.
	Description string `json:"description"`
	// The posted date of the earliest transaction in the stream.
	FirstDate string `json:"first_date"`
	// The posted date of the latest transaction in the stream.
	LastDate string `json:"last_date"`
	Frequency RecurringTransactionFrequency `json:"frequency"`
	// An array of Plaid transaction IDs belonging to the stream, sorted by posted date.
	TransactionIds []string `json:"transaction_ids"`
	AverageAmount TransactionStreamAmount `json:"average_amount"`
	// indicates whether the transaction stream is still live.
	IsActive bool `json:"is_active"`
	AdditionalProperties map[string]interface{}
}

type _TransactionStream TransactionStream

// NewTransactionStream instantiates a new TransactionStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionStream(accountId string, streamId string, categoryId string, category []string, description string, firstDate string, lastDate string, frequency RecurringTransactionFrequency, transactionIds []string, averageAmount TransactionStreamAmount, isActive bool) *TransactionStream {
	this := TransactionStream{}
	this.AccountId = accountId
	this.StreamId = streamId
	this.CategoryId = categoryId
	this.Category = category
	this.Description = description
	this.FirstDate = firstDate
	this.LastDate = lastDate
	this.Frequency = frequency
	this.TransactionIds = transactionIds
	this.AverageAmount = averageAmount
	this.IsActive = isActive
	return &this
}

// NewTransactionStreamWithDefaults instantiates a new TransactionStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionStreamWithDefaults() *TransactionStream {
	this := TransactionStream{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *TransactionStream) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransactionStream) SetAccountId(v string) {
	o.AccountId = v
}

// GetStreamId returns the StreamId field value
func (o *TransactionStream) GetStreamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamId
}

// GetStreamIdOk returns a tuple with the StreamId field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetStreamIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StreamId, true
}

// SetStreamId sets field value
func (o *TransactionStream) SetStreamId(v string) {
	o.StreamId = v
}

// GetCategoryId returns the CategoryId field value
func (o *TransactionStream) GetCategoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetCategoryIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *TransactionStream) SetCategoryId(v string) {
	o.CategoryId = v
}

// GetCategory returns the Category field value
func (o *TransactionStream) GetCategory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetCategoryOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *TransactionStream) SetCategory(v []string) {
	o.Category = v
}

// GetDescription returns the Description field value
func (o *TransactionStream) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransactionStream) SetDescription(v string) {
	o.Description = v
}

// GetFirstDate returns the FirstDate field value
func (o *TransactionStream) GetFirstDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstDate
}

// GetFirstDateOk returns a tuple with the FirstDate field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetFirstDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstDate, true
}

// SetFirstDate sets field value
func (o *TransactionStream) SetFirstDate(v string) {
	o.FirstDate = v
}

// GetLastDate returns the LastDate field value
func (o *TransactionStream) GetLastDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastDate
}

// GetLastDateOk returns a tuple with the LastDate field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetLastDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastDate, true
}

// SetLastDate sets field value
func (o *TransactionStream) SetLastDate(v string) {
	o.LastDate = v
}

// GetFrequency returns the Frequency field value
func (o *TransactionStream) GetFrequency() RecurringTransactionFrequency {
	if o == nil {
		var ret RecurringTransactionFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetFrequencyOk() (*RecurringTransactionFrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *TransactionStream) SetFrequency(v RecurringTransactionFrequency) {
	o.Frequency = v
}

// GetTransactionIds returns the TransactionIds field value
func (o *TransactionStream) GetTransactionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TransactionIds
}

// GetTransactionIdsOk returns a tuple with the TransactionIds field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetTransactionIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionIds, true
}

// SetTransactionIds sets field value
func (o *TransactionStream) SetTransactionIds(v []string) {
	o.TransactionIds = v
}

// GetAverageAmount returns the AverageAmount field value
func (o *TransactionStream) GetAverageAmount() TransactionStreamAmount {
	if o == nil {
		var ret TransactionStreamAmount
		return ret
	}

	return o.AverageAmount
}

// GetAverageAmountOk returns a tuple with the AverageAmount field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetAverageAmountOk() (*TransactionStreamAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AverageAmount, true
}

// SetAverageAmount sets field value
func (o *TransactionStream) SetAverageAmount(v TransactionStreamAmount) {
	o.AverageAmount = v
}

// GetIsActive returns the IsActive field value
func (o *TransactionStream) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *TransactionStream) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *TransactionStream) SetIsActive(v bool) {
	o.IsActive = v
}

func (o TransactionStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["stream_id"] = o.StreamId
	}
	if true {
		toSerialize["category_id"] = o.CategoryId
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["first_date"] = o.FirstDate
	}
	if true {
		toSerialize["last_date"] = o.LastDate
	}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if true {
		toSerialize["transaction_ids"] = o.TransactionIds
	}
	if true {
		toSerialize["average_amount"] = o.AverageAmount
	}
	if true {
		toSerialize["is_active"] = o.IsActive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionStream) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionStream := _TransactionStream{}

	if err = json.Unmarshal(bytes, &varTransactionStream); err == nil {
		*o = TransactionStream(varTransactionStream)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "stream_id")
		delete(additionalProperties, "category_id")
		delete(additionalProperties, "category")
		delete(additionalProperties, "description")
		delete(additionalProperties, "first_date")
		delete(additionalProperties, "last_date")
		delete(additionalProperties, "frequency")
		delete(additionalProperties, "transaction_ids")
		delete(additionalProperties, "average_amount")
		delete(additionalProperties, "is_active")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionStream struct {
	value *TransactionStream
	isSet bool
}

func (v NullableTransactionStream) Get() *TransactionStream {
	return v.value
}

func (v *NullableTransactionStream) Set(val *TransactionStream) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStream) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStream(val *TransactionStream) *NullableTransactionStream {
	return &NullableTransactionStream{value: val, isSet: true}
}

func (v NullableTransactionStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


