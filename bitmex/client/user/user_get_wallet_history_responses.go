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

// UserGetWalletHistoryReader is a Reader for the UserGetWalletHistory structure.
type UserGetWalletHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetWalletHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserGetWalletHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserGetWalletHistoryOK creates a UserGetWalletHistoryOK with default headers values
func NewUserGetWalletHistoryOK() *UserGetWalletHistoryOK {
	return &UserGetWalletHistoryOK{}
}

/*UserGetWalletHistoryOK handles this case with default header values.

Request was successful
*/
type UserGetWalletHistoryOK struct {
	Payload []*models.Transaction
}

func (o *UserGetWalletHistoryOK) Error() string {
	return fmt.Sprintf("[GET /user/walletHistory][%d] userGetWalletHistoryOK  %+v", 200, o.Payload)
}

func (o *UserGetWalletHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
