// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/frankrap/bitmexwrap/bitmex/models"
)

// UserUpdateReader is a Reader for the UserUpdate structure.
type UserUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserUpdateOK creates a UserUpdateOK with default headers values
func NewUserUpdateOK() *UserUpdateOK {
	return &UserUpdateOK{}
}

/*UserUpdateOK handles this case with default header values.

Request was successful
*/
type UserUpdateOK struct {
	Payload *models.User
}

func (o *UserUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /user][%d] userUpdateOK  %+v", 200, *o.Payload)
}

func (o *UserUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
