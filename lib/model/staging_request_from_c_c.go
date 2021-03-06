package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*StagingRequestFromCC staging request from c c

swagger:model StagingRequestFromCC
*/
type StagingRequestFromCC struct {

	/* app id
	 */
	AppID string `json:"app_id,omitempty"`

	/* completion callback
	 */
	CompletionCallback string `json:"completion_callback,omitempty"`

	/* disk mb
	 */
	DiskMb int64 `json:"disk_mb,omitempty"`

	/* egress rules
	 */
	EgressRules []*SecurityGroupRule `json:"egress_rules,omitempty"`

	/* environment
	 */
	Environment []*EnvironmentVariable `json:"environment,omitempty"`

	/* file descriptors
	 */
	FileDescriptors int64 `json:"file_descriptors,omitempty"`

	/* lifecycle
	 */
	Lifecycle string `json:"lifecycle,omitempty"`

	/* lifecycle data
	 */
	LifecycleData interface{} `json:"lifecycle_data,omitempty"`

	/* log guid
	 */
	LogGUID string `json:"log_guid,omitempty"`

	/* memory mb
	 */
	MemoryMb int64 `json:"memory_mb,omitempty"`

	/* timeout
	 */
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this staging request from c c
func (m *StagingRequestFromCC) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEgressRules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StagingRequestFromCC) validateEgressRules(formats strfmt.Registry) error {

	if swag.IsZero(m.EgressRules) { // not required
		return nil
	}

	for i := 0; i < len(m.EgressRules); i++ {

		if swag.IsZero(m.EgressRules[i]) { // not required
			continue
		}

		if m.EgressRules[i] != nil {

			if err := m.EgressRules[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *StagingRequestFromCC) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	for i := 0; i < len(m.Environment); i++ {

		if swag.IsZero(m.Environment[i]) { // not required
			continue
		}

		if m.Environment[i] != nil {

			if err := m.Environment[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
