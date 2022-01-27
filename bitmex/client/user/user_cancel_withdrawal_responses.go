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

// UserCancelWithdrawalReader is a Reader for the UserCancelWithdrawal structure.
type UserCancelWithdrawalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCancelWithdrawalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCancelWithdrawalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCancelWithdrawalOK creates a UserCancelWithdrawalOK with default headers values
func NewUserCancelWithdrawalOK() *UserCancelWithdrawalOK {
	return &UserCancelWithdrawalOK{}
}

/*UserCancelWithdrawalOK handles this case with default header values.

Request was successful
*/
type UserCancelWithdrawalOK struct {
	Payload *models.Transaction
}

func (o *UserCancelWithdrawalOK) Error() string {
	return fmt.Sprintf("[POST /user/cancelWithdrawal][%d] userCancelWithdrawalOK  %+v", 200, *o.Payload)
}

func (o *UserCancelWithdrawalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Transaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
