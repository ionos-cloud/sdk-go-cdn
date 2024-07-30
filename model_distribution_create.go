/*
 * IONOS Cloud - CDN Distribution API
 *
 * This API manages CDN distributions.
 *
 * API version: 0.1.7
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// DistributionCreate struct for DistributionCreate
type DistributionCreate struct {
	// Metadata
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	Properties *DistributionProperties `json:"properties"`
}

// NewDistributionCreate instantiates a new DistributionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionCreate(properties DistributionProperties) *DistributionCreate {
	this := DistributionCreate{}

	this.Properties = &properties

	return &this
}

// NewDistributionCreateWithDefaults instantiates a new DistributionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionCreateWithDefaults() *DistributionCreate {
	this := DistributionCreate{}
	return &this
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *DistributionCreate) GetMetadata() *map[string]interface{} {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionCreate) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *DistributionCreate) SetMetadata(v map[string]interface{}) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *DistributionCreate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for DistributionProperties will be returned
func (o *DistributionCreate) GetProperties() *DistributionProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DistributionCreate) GetPropertiesOk() (*DistributionProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *DistributionCreate) SetProperties(v DistributionProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *DistributionCreate) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o DistributionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableDistributionCreate struct {
	value *DistributionCreate
	isSet bool
}

func (v NullableDistributionCreate) Get() *DistributionCreate {
	return v.value
}

func (v *NullableDistributionCreate) Set(val *DistributionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionCreate(val *DistributionCreate) *NullableDistributionCreate {
	return &NullableDistributionCreate{value: val, isSet: true}
}

func (v NullableDistributionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
