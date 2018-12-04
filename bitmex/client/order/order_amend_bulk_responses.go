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

// OrderAmendBulkReader is a Reader for the OrderAmendBulk structure.
type OrderAmendBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderAmendBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderAmendBulkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderAmendBulkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOrderAmendBulkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOrderAmendBulkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderAmendBulkOK creates a OrderAmendBulkOK with default headers values
func NewOrderAmendBulkOK() *OrderAmendBulkOK {
	return &OrderAmendBulkOK{}
}

/*OrderAmendBulkOK handles this case with default header values.

Request was successful
*/
type OrderAmendBulkOK struct {
	Payload []*models.Order
}

func (o *OrderAmendBulkOK) Error() string {
	return fmt.Sprintf("[PUT /order/bulk][%d] orderAmendBulkOK  %+v", 200, o.Payload)
}

func (o *OrderAmendBulkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderAmendBulkBadRequest creates a OrderAmendBulkBadRequest with default headers values
func NewOrderAmendBulkBadRequest() *OrderAmendBulkBadRequest {
	return &OrderAmendBulkBadRequest{}
}

/*OrderAmendBulkBadRequest handles this case with default header values.

Parameter Error
*/
type OrderAmendBulkBadRequest struct {
	Payload *models.Error
}

func (o *OrderAmendBulkBadRequest) Error() string {
	return fmt.Sprintf("[PUT /order/bulk][%d] orderAmendBulkBadRequest  %+v", 400, *o.Payload)
}

func (o *OrderAmendBulkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderAmendBulkUnauthorized creates a OrderAmendBulkUnauthorized with default headers values
func NewOrderAmendBulkUnauthorized() *OrderAmendBulkUnauthorized {
	return &OrderAmendBulkUnauthorized{}
}

/*OrderAmendBulkUnauthorized handles this case with default header values.

Unauthorized
*/
type OrderAmendBulkUnauthorized struct {
	Payload *models.Error
}

func (o *OrderAmendBulkUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /order/bulk][%d] orderAmendBulkUnauthorized  %+v", 401, o.Payload)
}

func (o *OrderAmendBulkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderAmendBulkNotFound creates a OrderAmendBulkNotFound with default headers values
func NewOrderAmendBulkNotFound() *OrderAmendBulkNotFound {
	return &OrderAmendBulkNotFound{}
}

/*OrderAmendBulkNotFound handles this case with default header values.

Not Found
*/
type OrderAmendBulkNotFound struct {
	Payload *models.Error
}

func (o *OrderAmendBulkNotFound) Error() string {
	return fmt.Sprintf("[PUT /order/bulk][%d] orderAmendBulkNotFound  %+v", 404, o.Payload)
}

func (o *OrderAmendBulkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
