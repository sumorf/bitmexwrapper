// Code generated by go-swagger; DO NOT EDIT.

package trade

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/frankrap/bitmexwrap/bitmex/models"
)

// TradeGetBucketedReader is a Reader for the TradeGetBucketed structure.
type TradeGetBucketedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TradeGetBucketedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTradeGetBucketedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTradeGetBucketedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTradeGetBucketedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTradeGetBucketedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTradeGetBucketedOK creates a TradeGetBucketedOK with default headers values
func NewTradeGetBucketedOK() *TradeGetBucketedOK {
	return &TradeGetBucketedOK{}
}

/*TradeGetBucketedOK handles this case with default header values.

Request was successful
*/
type TradeGetBucketedOK struct {
	Payload []*models.TradeBin
}

func (o *TradeGetBucketedOK) Error() string {
	return fmt.Sprintf("[GET /trade/bucketed][%d] tradeGetBucketedOK  %+v", 200, o.Payload)
}

func (o *TradeGetBucketedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTradeGetBucketedBadRequest creates a TradeGetBucketedBadRequest with default headers values
func NewTradeGetBucketedBadRequest() *TradeGetBucketedBadRequest {
	return &TradeGetBucketedBadRequest{}
}

/*TradeGetBucketedBadRequest handles this case with default header values.

Parameter Error
*/
type TradeGetBucketedBadRequest struct {
	Payload *models.Error
}

func (o *TradeGetBucketedBadRequest) Error() string {
	return fmt.Sprintf("[GET /trade/bucketed][%d] tradeGetBucketedBadRequest  %+v", 400, *o.Payload.Error)
}

func (o *TradeGetBucketedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTradeGetBucketedUnauthorized creates a TradeGetBucketedUnauthorized with default headers values
func NewTradeGetBucketedUnauthorized() *TradeGetBucketedUnauthorized {
	return &TradeGetBucketedUnauthorized{}
}

/*TradeGetBucketedUnauthorized handles this case with default header values.

Unauthorized
*/
type TradeGetBucketedUnauthorized struct {
	Payload *models.Error
}

func (o *TradeGetBucketedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /trade/bucketed][%d] tradeGetBucketedUnauthorized  %+v", 401, o.Payload.Error)
}

func (o *TradeGetBucketedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTradeGetBucketedNotFound creates a TradeGetBucketedNotFound with default headers values
func NewTradeGetBucketedNotFound() *TradeGetBucketedNotFound {
	return &TradeGetBucketedNotFound{}
}

/*TradeGetBucketedNotFound handles this case with default header values.

Not Found
*/
type TradeGetBucketedNotFound struct {
	Payload *models.Error
}

func (o *TradeGetBucketedNotFound) Error() string {
	return fmt.Sprintf("[GET /trade/bucketed][%d] tradeGetBucketedNotFound  %+v", 404, o.Payload.Error)
}

func (o *TradeGetBucketedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
