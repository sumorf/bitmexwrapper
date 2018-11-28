// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/bitmexwrap/bitmex/models"
)

// UserConfirmReader is a Reader for the UserConfirm structure.
type UserConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserConfirmOK creates a UserConfirmOK with default headers values
func NewUserConfirmOK() *UserConfirmOK {
	return &UserConfirmOK{}
}

/*UserConfirmOK handles this case with default header values.

Request was successful
*/
type UserConfirmOK struct {
	Payload *models.AccessToken
}

func (o *UserConfirmOK) Error() string {
	return fmt.Sprintf("[POST /user/confirmEmail][%d] userConfirmOK  %+v", 200, o.Payload)
}

func (o *UserConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
