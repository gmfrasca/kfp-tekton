// Code generated by go-swagger; DO NOT EDIT.

package experiment_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1ExperimentStorageState v1 experiment storage state
// swagger:model v1ExperimentStorageState
type V1ExperimentStorageState string

const (

	// V1ExperimentStorageStateSTORAGESTATEUNSPECIFIED captures enum value "STORAGESTATE_UNSPECIFIED"
	V1ExperimentStorageStateSTORAGESTATEUNSPECIFIED V1ExperimentStorageState = "STORAGESTATE_UNSPECIFIED"

	// V1ExperimentStorageStateSTORAGESTATEAVAILABLE captures enum value "STORAGESTATE_AVAILABLE"
	V1ExperimentStorageStateSTORAGESTATEAVAILABLE V1ExperimentStorageState = "STORAGESTATE_AVAILABLE"

	// V1ExperimentStorageStateSTORAGESTATEARCHIVED captures enum value "STORAGESTATE_ARCHIVED"
	V1ExperimentStorageStateSTORAGESTATEARCHIVED V1ExperimentStorageState = "STORAGESTATE_ARCHIVED"
)

// for schema
var v1ExperimentStorageStateEnum []interface{}

func init() {
	var res []V1ExperimentStorageState
	if err := json.Unmarshal([]byte(`["STORAGESTATE_UNSPECIFIED","STORAGESTATE_AVAILABLE","STORAGESTATE_ARCHIVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ExperimentStorageStateEnum = append(v1ExperimentStorageStateEnum, v)
	}
}

func (m V1ExperimentStorageState) validateV1ExperimentStorageStateEnum(path, location string, value V1ExperimentStorageState) error {
	if err := validate.Enum(path, location, value, v1ExperimentStorageStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 experiment storage state
func (m V1ExperimentStorageState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ExperimentStorageStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}