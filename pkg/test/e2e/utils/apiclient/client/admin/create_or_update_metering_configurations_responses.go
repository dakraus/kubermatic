// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateOrUpdateMeteringConfigurationsReader is a Reader for the CreateOrUpdateMeteringConfigurations structure.
type CreateOrUpdateMeteringConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrUpdateMeteringConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrUpdateMeteringConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateOrUpdateMeteringConfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOrUpdateMeteringConfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOrUpdateMeteringConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrUpdateMeteringConfigurationsOK creates a CreateOrUpdateMeteringConfigurationsOK with default headers values
func NewCreateOrUpdateMeteringConfigurationsOK() *CreateOrUpdateMeteringConfigurationsOK {
	return &CreateOrUpdateMeteringConfigurationsOK{}
}

/* CreateOrUpdateMeteringConfigurationsOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringConfigurationsOK struct {
}

func (o *CreateOrUpdateMeteringConfigurationsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsOK ", 200)
}

func (o *CreateOrUpdateMeteringConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringConfigurationsUnauthorized creates a CreateOrUpdateMeteringConfigurationsUnauthorized with default headers values
func NewCreateOrUpdateMeteringConfigurationsUnauthorized() *CreateOrUpdateMeteringConfigurationsUnauthorized {
	return &CreateOrUpdateMeteringConfigurationsUnauthorized{}
}

/* CreateOrUpdateMeteringConfigurationsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringConfigurationsUnauthorized struct {
}

func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsUnauthorized ", 401)
}

func (o *CreateOrUpdateMeteringConfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringConfigurationsForbidden creates a CreateOrUpdateMeteringConfigurationsForbidden with default headers values
func NewCreateOrUpdateMeteringConfigurationsForbidden() *CreateOrUpdateMeteringConfigurationsForbidden {
	return &CreateOrUpdateMeteringConfigurationsForbidden{}
}

/* CreateOrUpdateMeteringConfigurationsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateOrUpdateMeteringConfigurationsForbidden struct {
}

func (o *CreateOrUpdateMeteringConfigurationsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurationsForbidden ", 403)
}

func (o *CreateOrUpdateMeteringConfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateMeteringConfigurationsDefault creates a CreateOrUpdateMeteringConfigurationsDefault with default headers values
func NewCreateOrUpdateMeteringConfigurationsDefault(code int) *CreateOrUpdateMeteringConfigurationsDefault {
	return &CreateOrUpdateMeteringConfigurationsDefault{
		_statusCode: code,
	}
}

/* CreateOrUpdateMeteringConfigurationsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateOrUpdateMeteringConfigurationsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create or update metering configurations default response
func (o *CreateOrUpdateMeteringConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrUpdateMeteringConfigurationsDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations][%d] createOrUpdateMeteringConfigurations default  %+v", o._statusCode, o.Payload)
}
func (o *CreateOrUpdateMeteringConfigurationsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateOrUpdateMeteringConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}