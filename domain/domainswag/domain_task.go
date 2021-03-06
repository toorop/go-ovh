// Code generated by go-swagger; DO NOT EDIT.

package domainswag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainTask Tasks associated to domain
// swagger:model domain.Task
type DomainTask struct {

	// Can accelerate the task
	// Required: true
	// Read Only: true
	CanAccelerate bool `json:"canAccelerate"`

	// Can cancel the task
	// Required: true
	// Read Only: true
	CanCancel bool `json:"canCancel"`

	// Can relaunch the task
	// Required: true
	// Read Only: true
	CanRelaunch bool `json:"canRelaunch"`

	// Comment about the task
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// Creation date of the task
	// Required: true
	// Read Only: true
	CreationDate strfmt.DateTime `json:"creationDate"`

	// Done date of the task
	// Read Only: true
	DoneDate strfmt.DateTime `json:"doneDate,omitempty"`

	// Function of the task
	// Required: true
	// Read Only: true
	Function string `json:"function"`

	// Id of the task
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Last update date of the task
	// Required: true
	// Read Only: true
	LastUpdate strfmt.DateTime `json:"lastUpdate"`

	// Status of the task
	// Required: true
	// Read Only: true
	Status string `json:"status"`

	// Todo date of the task
	// Required: true
	// Read Only: true
	TodoDate strfmt.DateTime `json:"todoDate"`
}

// Validate validates this domain task
func (m *DomainTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanAccelerate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCanCancel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCanRelaunch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDoneDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFunction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastUpdate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTodoDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainTask) validateCanAccelerate(formats strfmt.Registry) error {

	if err := validate.Required("canAccelerate", "body", bool(m.CanAccelerate)); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateCanCancel(formats strfmt.Registry) error {

	if err := validate.Required("canCancel", "body", bool(m.CanCancel)); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateCanRelaunch(formats strfmt.Registry) error {

	if err := validate.Required("canRelaunch", "body", bool(m.CanRelaunch)); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", strfmt.DateTime(m.CreationDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateDoneDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DoneDate) { // not required
		return nil
	}

	if err := validate.FormatOf("doneDate", "body", "date-time", m.DoneDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateFunction(formats strfmt.Registry) error {

	if err := validate.RequiredString("function", "body", string(m.Function)); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateLastUpdate(formats strfmt.Registry) error {

	if err := validate.Required("lastUpdate", "body", strfmt.DateTime(m.LastUpdate)); err != nil {
		return err
	}

	if err := validate.FormatOf("lastUpdate", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

var domainTaskTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cancelled","doing","done","error","todo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainTaskTypeStatusPropEnum = append(domainTaskTypeStatusPropEnum, v)
	}
}

const (

	// DomainTaskStatusCancelled captures enum value "cancelled"
	DomainTaskStatusCancelled string = "cancelled"

	// DomainTaskStatusDoing captures enum value "doing"
	DomainTaskStatusDoing string = "doing"

	// DomainTaskStatusDone captures enum value "done"
	DomainTaskStatusDone string = "done"

	// DomainTaskStatusError captures enum value "error"
	DomainTaskStatusError string = "error"

	// DomainTaskStatusTodo captures enum value "todo"
	DomainTaskStatusTodo string = "todo"
)

// prop value enum
func (m *DomainTask) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, domainTaskTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DomainTask) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *DomainTask) validateTodoDate(formats strfmt.Registry) error {

	if err := validate.Required("todoDate", "body", strfmt.DateTime(m.TodoDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("todoDate", "body", "date-time", m.TodoDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainTask) UnmarshalBinary(b []byte) error {
	var res DomainTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
