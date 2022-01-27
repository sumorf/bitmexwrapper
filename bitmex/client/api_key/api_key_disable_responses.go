// Code generated by go-swagger; DO NOT EDIT.

package api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/frankrap/bitmexwrap/bitmex/models"
)

// APIKeyDisableReader is a Reader for the APIKeyDisable structure.
type APIKeyDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIKeyDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAPIKeyDisableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAPIKeyDisableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAPIKeyDisableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAPIKeyDisableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIKeyDisableOK creates a APIKeyDisableOK with default headers values
func NewAPIKeyDisableOK() *APIKeyDisableOK {
	return &APIKeyDisableOK{}
}

/*APIKeyDisableOK handles this case with default header values.

Request was successful
*/
type APIKeyDisableOK struct {
	Payload *models.APIKey
}

func (o *APIKeyDisableOK) Error() string {
	return fmt.Sprintf("[POST /apiKey/disable][%d] apiKeyDisableOK  %+v", 200, *o.Payload)
}

func (o *APIKeyDisableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeyDisableBadRequest creates a APIKeyDisableBadRequest with default headers values
func NewAPIKeyDisableBadRequest() *APIKeyDisableBadRequest {
	return &APIKeyDisableBadRequest{}
}

/*APIKeyDisableBadRequest handles this case with default header values.

Parameter Error
*/
type APIKeyDisableBadRequest struct {
	Payload *models.Error
}

func (o *APIKeyDisableBadRequest) Error() string {
	return fmt.Sprintf("[POST /apiKey/disable][%d] apiKeyDisableBadRequest  %+v", 400, *o.Payload)
}

func (o *APIKeyDisableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeyDisableUnauthorized creates a APIKeyDisableUnauthorized with default headers values
func NewAPIKeyDisableUnauthorized() *APIKeyDisableUnauthorized {
	return &APIKeyDisableUnauthorized{}
}

/*APIKeyDisableUnauthorized handles this case with default header values.

Unauthorized
*/
type APIKeyDisableUnauthorized struct {
	Payload *models.Error
}

func (o *APIKeyDisableUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apiKey/disable][%d] apiKeyDisableUnauthorized  %+v", 401, o.Payload)
}

func (o *APIKeyDisableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeyDisableNotFound creates a APIKeyDisableNotFound with default headers values
func NewAPIKeyDisableNotFound() *APIKeyDisableNotFound {
	return &APIKeyDisableNotFound{}
}

/*APIKeyDisableNotFound handles this case with default header values.

Not Found
*/
type APIKeyDisableNotFound struct {
	Payload *models.Error
}

func (o *APIKeyDisableNotFound) Error() string {
	return fmt.Sprintf("[POST /apiKey/disable][%d] apiKeyDisableNotFound  %+v", 404, o.Payload)
}

func (o *APIKeyDisableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
