// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/bitmexwrap/bitmex/models"
)

// OrderCancelAllAfterReader is a Reader for the OrderCancelAllAfter structure.
type OrderCancelAllAfterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderCancelAllAfterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderCancelAllAfterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderCancelAllAfterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOrderCancelAllAfterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOrderCancelAllAfterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderCancelAllAfterOK creates a OrderCancelAllAfterOK with default headers values
func NewOrderCancelAllAfterOK() *OrderCancelAllAfterOK {
	return &OrderCancelAllAfterOK{}
}

/*OrderCancelAllAfterOK handles this case with default header values.

Request was successful
*/
type OrderCancelAllAfterOK struct {
	Payload interface{}
}

func (o *OrderCancelAllAfterOK) Error() string {
	return fmt.Sprintf("[POST /order/cancelAllAfter][%d] orderCancelAllAfterOK  %+v", 200, o.Payload)
}

func (o *OrderCancelAllAfterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCancelAllAfterBadRequest creates a OrderCancelAllAfterBadRequest with default headers values
func NewOrderCancelAllAfterBadRequest() *OrderCancelAllAfterBadRequest {
	return &OrderCancelAllAfterBadRequest{}
}

/*OrderCancelAllAfterBadRequest handles this case with default header values.

Parameter Error
*/
type OrderCancelAllAfterBadRequest struct {
	Payload *models.Error
}

func (o *OrderCancelAllAfterBadRequest) Error() string {
	return fmt.Sprintf("[POST /order/cancelAllAfter][%d] orderCancelAllAfterBadRequest  %+v", 400, o.Payload)
}

func (o *OrderCancelAllAfterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCancelAllAfterUnauthorized creates a OrderCancelAllAfterUnauthorized with default headers values
func NewOrderCancelAllAfterUnauthorized() *OrderCancelAllAfterUnauthorized {
	return &OrderCancelAllAfterUnauthorized{}
}

/*OrderCancelAllAfterUnauthorized handles this case with default header values.

Unauthorized
*/
type OrderCancelAllAfterUnauthorized struct {
	Payload *models.Error
}

func (o *OrderCancelAllAfterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /order/cancelAllAfter][%d] orderCancelAllAfterUnauthorized  %+v", 401, o.Payload)
}

func (o *OrderCancelAllAfterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCancelAllAfterNotFound creates a OrderCancelAllAfterNotFound with default headers values
func NewOrderCancelAllAfterNotFound() *OrderCancelAllAfterNotFound {
	return &OrderCancelAllAfterNotFound{}
}

/*OrderCancelAllAfterNotFound handles this case with default header values.

Not Found
*/
type OrderCancelAllAfterNotFound struct {
	Payload *models.Error
}

func (o *OrderCancelAllAfterNotFound) Error() string {
	return fmt.Sprintf("[POST /order/cancelAllAfter][%d] orderCancelAllAfterNotFound  %+v", 404, o.Payload)
}

func (o *OrderCancelAllAfterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
