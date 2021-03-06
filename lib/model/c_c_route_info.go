package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CCRouteInfo c c route info

swagger:model CCRouteInfo
*/
type CCRouteInfo map[string]string

// Validate validates this c c route info
func (m CCRouteInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", CCRouteInfo(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
