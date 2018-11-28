// Code generated by go-swagger; DO NOT EDIT.

package position

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/bitmexwrap/bitmex/models"
)

// PositionTransferIsolatedMarginReader is a Reader for the PositionTransferIsolatedMargin structure.
type PositionTransferIsolatedMarginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PositionTransferIsolatedMarginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPositionTransferIsolatedMarginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPositionTransferIsolatedMarginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPositionTransferIsolatedMarginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPositionTransferIsolatedMarginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPositionTransferIsolatedMarginOK creates a PositionTransferIsolatedMarginOK with default headers values
func NewPositionTransferIsolatedMarginOK() *PositionTransferIsolatedMarginOK {
	return &PositionTransferIsolatedMarginOK{}
}

/*PositionTransferIsolatedMarginOK handles this case with default header values.

Request was successful
*/
type PositionTransferIsolatedMarginOK struct {
	Payload *models.Position
}

func (o *PositionTransferIsolatedMarginOK) Error() string {
	return fmt.Sprintf("[POST /position/transferMargin][%d] positionTransferIsolatedMarginOK  %+v", 200, o.Payload)
}

func (o *PositionTransferIsolatedMarginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Position)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionTransferIsolatedMarginBadRequest creates a PositionTransferIsolatedMarginBadRequest with default headers values
func NewPositionTransferIsolatedMarginBadRequest() *PositionTransferIsolatedMarginBadRequest {
	return &PositionTransferIsolatedMarginBadRequest{}
}

/*PositionTransferIsolatedMarginBadRequest handles this case with default header values.

Parameter Error
*/
type PositionTransferIsolatedMarginBadRequest struct {
	Payload *models.Error
}

func (o *PositionTransferIsolatedMarginBadRequest) Error() string {
	return fmt.Sprintf("[POST /position/transferMargin][%d] positionTransferIsolatedMarginBadRequest  %+v", 400, o.Payload)
}

func (o *PositionTransferIsolatedMarginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionTransferIsolatedMarginUnauthorized creates a PositionTransferIsolatedMarginUnauthorized with default headers values
func NewPositionTransferIsolatedMarginUnauthorized() *PositionTransferIsolatedMarginUnauthorized {
	return &PositionTransferIsolatedMarginUnauthorized{}
}

/*PositionTransferIsolatedMarginUnauthorized handles this case with default header values.

Unauthorized
*/
type PositionTransferIsolatedMarginUnauthorized struct {
	Payload *models.Error
}

func (o *PositionTransferIsolatedMarginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /position/transferMargin][%d] positionTransferIsolatedMarginUnauthorized  %+v", 401, o.Payload)
}

func (o *PositionTransferIsolatedMarginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionTransferIsolatedMarginNotFound creates a PositionTransferIsolatedMarginNotFound with default headers values
func NewPositionTransferIsolatedMarginNotFound() *PositionTransferIsolatedMarginNotFound {
	return &PositionTransferIsolatedMarginNotFound{}
}

/*PositionTransferIsolatedMarginNotFound handles this case with default header values.

Not Found
*/
type PositionTransferIsolatedMarginNotFound struct {
	Payload *models.Error
}

func (o *PositionTransferIsolatedMarginNotFound) Error() string {
	return fmt.Sprintf("[POST /position/transferMargin][%d] positionTransferIsolatedMarginNotFound  %+v", 404, o.Payload)
}

func (o *PositionTransferIsolatedMarginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
