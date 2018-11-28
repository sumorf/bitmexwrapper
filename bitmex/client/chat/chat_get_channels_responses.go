// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/bitmexwrap/bitmex/models"
)

// ChatGetChannelsReader is a Reader for the ChatGetChannels structure.
type ChatGetChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChatGetChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChatGetChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChatGetChannelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewChatGetChannelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChatGetChannelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChatGetChannelsOK creates a ChatGetChannelsOK with default headers values
func NewChatGetChannelsOK() *ChatGetChannelsOK {
	return &ChatGetChannelsOK{}
}

/*ChatGetChannelsOK handles this case with default header values.

Request was successful
*/
type ChatGetChannelsOK struct {
	Payload []*models.ChatChannel
}

func (o *ChatGetChannelsOK) Error() string {
	return fmt.Sprintf("[GET /chat/channels][%d] chatGetChannelsOK  %+v", 200, o.Payload)
}

func (o *ChatGetChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChatGetChannelsBadRequest creates a ChatGetChannelsBadRequest with default headers values
func NewChatGetChannelsBadRequest() *ChatGetChannelsBadRequest {
	return &ChatGetChannelsBadRequest{}
}

/*ChatGetChannelsBadRequest handles this case with default header values.

Parameter Error
*/
type ChatGetChannelsBadRequest struct {
	Payload *models.Error
}

func (o *ChatGetChannelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /chat/channels][%d] chatGetChannelsBadRequest  %+v", 400, o.Payload)
}

func (o *ChatGetChannelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChatGetChannelsUnauthorized creates a ChatGetChannelsUnauthorized with default headers values
func NewChatGetChannelsUnauthorized() *ChatGetChannelsUnauthorized {
	return &ChatGetChannelsUnauthorized{}
}

/*ChatGetChannelsUnauthorized handles this case with default header values.

Unauthorized
*/
type ChatGetChannelsUnauthorized struct {
	Payload *models.Error
}

func (o *ChatGetChannelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /chat/channels][%d] chatGetChannelsUnauthorized  %+v", 401, o.Payload)
}

func (o *ChatGetChannelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChatGetChannelsNotFound creates a ChatGetChannelsNotFound with default headers values
func NewChatGetChannelsNotFound() *ChatGetChannelsNotFound {
	return &ChatGetChannelsNotFound{}
}

/*ChatGetChannelsNotFound handles this case with default header values.

Not Found
*/
type ChatGetChannelsNotFound struct {
	Payload *models.Error
}

func (o *ChatGetChannelsNotFound) Error() string {
	return fmt.Sprintf("[GET /chat/channels][%d] chatGetChannelsNotFound  %+v", 404, o.Payload)
}

func (o *ChatGetChannelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
