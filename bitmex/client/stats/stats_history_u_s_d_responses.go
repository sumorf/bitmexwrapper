// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/bitmexwrap/bitmex/models"
)

// StatsHistoryUSDReader is a Reader for the StatsHistoryUSD structure.
type StatsHistoryUSDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatsHistoryUSDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatsHistoryUSDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewStatsHistoryUSDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewStatsHistoryUSDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStatsHistoryUSDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatsHistoryUSDOK creates a StatsHistoryUSDOK with default headers values
func NewStatsHistoryUSDOK() *StatsHistoryUSDOK {
	return &StatsHistoryUSDOK{}
}

/*StatsHistoryUSDOK handles this case with default header values.

Request was successful
*/
type StatsHistoryUSDOK struct {
	Payload []*models.StatsUSD
}

func (o *StatsHistoryUSDOK) Error() string {
	return fmt.Sprintf("[GET /stats/historyUSD][%d] statsHistoryUSDOK  %+v", 200, o.Payload)
}

func (o *StatsHistoryUSDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsHistoryUSDBadRequest creates a StatsHistoryUSDBadRequest with default headers values
func NewStatsHistoryUSDBadRequest() *StatsHistoryUSDBadRequest {
	return &StatsHistoryUSDBadRequest{}
}

/*StatsHistoryUSDBadRequest handles this case with default header values.

Parameter Error
*/
type StatsHistoryUSDBadRequest struct {
	Payload *models.Error
}

func (o *StatsHistoryUSDBadRequest) Error() string {
	return fmt.Sprintf("[GET /stats/historyUSD][%d] statsHistoryUSDBadRequest  %+v", 400, o.Payload)
}

func (o *StatsHistoryUSDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsHistoryUSDUnauthorized creates a StatsHistoryUSDUnauthorized with default headers values
func NewStatsHistoryUSDUnauthorized() *StatsHistoryUSDUnauthorized {
	return &StatsHistoryUSDUnauthorized{}
}

/*StatsHistoryUSDUnauthorized handles this case with default header values.

Unauthorized
*/
type StatsHistoryUSDUnauthorized struct {
	Payload *models.Error
}

func (o *StatsHistoryUSDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /stats/historyUSD][%d] statsHistoryUSDUnauthorized  %+v", 401, o.Payload)
}

func (o *StatsHistoryUSDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsHistoryUSDNotFound creates a StatsHistoryUSDNotFound with default headers values
func NewStatsHistoryUSDNotFound() *StatsHistoryUSDNotFound {
	return &StatsHistoryUSDNotFound{}
}

/*StatsHistoryUSDNotFound handles this case with default header values.

Not Found
*/
type StatsHistoryUSDNotFound struct {
	Payload *models.Error
}

func (o *StatsHistoryUSDNotFound) Error() string {
	return fmt.Sprintf("[GET /stats/historyUSD][%d] statsHistoryUSDNotFound  %+v", 404, o.Payload)
}

func (o *StatsHistoryUSDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
