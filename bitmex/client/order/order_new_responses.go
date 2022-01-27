// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/frankrap/bitmexwrap/bitmex/models"
)

// OrderNewReader is a Reader for the OrderNew structure.
type OrderNewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderNewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderNewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderNewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOrderNewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOrderNewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderNewOK creates a OrderNewOK with default headers values
func NewOrderNewOK() *OrderNewOK {
	return &OrderNewOK{}
}

/*OrderNewOK handles this case with default header values.

Request was successful
*/
type OrderNewOK struct {
	Payload *models.Order
}

func (o *OrderNewOK) Error() string {
	return fmt.Sprintf("[POST /order][%d] orderNewOK  %+v", 200, *o.Payload)
}

func (o *OrderNewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderNewBadRequest creates a OrderNewBadRequest with default headers values
func NewOrderNewBadRequest() *OrderNewBadRequest {
	return &OrderNewBadRequest{}
}

/*OrderNewBadRequest handles this case with default header values.

Parameter Error
*/
type OrderNewBadRequest struct {
	Payload *models.Error
}

func (o *OrderNewBadRequest) Error() string {
	return fmt.Sprintf("[POST /order][%d] orderNewBadRequest  %+v", 400, o.Payload.Error)
}

func (o *OrderNewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderNewUnauthorized creates a OrderNewUnauthorized with default headers values
func NewOrderNewUnauthorized() *OrderNewUnauthorized {
	return &OrderNewUnauthorized{}
}

/*OrderNewUnauthorized handles this case with default header values.

Unauthorized
*/
type OrderNewUnauthorized struct {
	Payload *models.Error
}

func (o *OrderNewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /order][%d] orderNewUnauthorized  %+v", 401, o.Payload)
}

func (o *OrderNewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderNewNotFound creates a OrderNewNotFound with default headers values
func NewOrderNewNotFound() *OrderNewNotFound {
	return &OrderNewNotFound{}
}

/*OrderNewNotFound handles this case with default header values.

Not Found
*/
type OrderNewNotFound struct {
	Payload *models.Error
}

func (o *OrderNewNotFound) Error() string {
	return fmt.Sprintf("[POST /order][%d] orderNewNotFound  %+v", 404, o.Payload)
}

func (o *OrderNewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
