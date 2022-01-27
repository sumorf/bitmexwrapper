// Code generated by go-swagger; DO NOT EDIT.

package execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/frankrap/bitmexwrap/bitmex/models"
)

// ExecutionGetReader is a Reader for the ExecutionGet structure.
type ExecutionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecutionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExecutionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewExecutionGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewExecutionGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewExecutionGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExecutionGetOK creates a ExecutionGetOK with default headers values
func NewExecutionGetOK() *ExecutionGetOK {
	return &ExecutionGetOK{}
}

/*ExecutionGetOK handles this case with default header values.

Request was successful
*/
type ExecutionGetOK struct {
	Payload []*models.Execution
}

func (o *ExecutionGetOK) Error() string {
	return fmt.Sprintf("[GET /execution][%d] executionGetOK  %+v", 200, o.Payload)
}

func (o *ExecutionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionGetBadRequest creates a ExecutionGetBadRequest with default headers values
func NewExecutionGetBadRequest() *ExecutionGetBadRequest {
	return &ExecutionGetBadRequest{}
}

/*ExecutionGetBadRequest handles this case with default header values.

Parameter Error
*/
type ExecutionGetBadRequest struct {
	Payload *models.Error
}

func (o *ExecutionGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /execution][%d] executionGetBadRequest  %+v", 400, *o.Payload)
}

func (o *ExecutionGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionGetUnauthorized creates a ExecutionGetUnauthorized with default headers values
func NewExecutionGetUnauthorized() *ExecutionGetUnauthorized {
	return &ExecutionGetUnauthorized{}
}

/*ExecutionGetUnauthorized handles this case with default header values.

Unauthorized
*/
type ExecutionGetUnauthorized struct {
	Payload *models.Error
}

func (o *ExecutionGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /execution][%d] executionGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ExecutionGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionGetNotFound creates a ExecutionGetNotFound with default headers values
func NewExecutionGetNotFound() *ExecutionGetNotFound {
	return &ExecutionGetNotFound{}
}

/*ExecutionGetNotFound handles this case with default header values.

Not Found
*/
type ExecutionGetNotFound struct {
	Payload *models.Error
}

func (o *ExecutionGetNotFound) Error() string {
	return fmt.Sprintf("[GET /execution][%d] executionGetNotFound  %+v", 404, o.Payload)
}

func (o *ExecutionGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
