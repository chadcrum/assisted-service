// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// DownloadClusterISOCloneReader is a Reader for the DownloadClusterISOClone structure.
type DownloadClusterISOCloneReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadClusterISOCloneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadClusterISOCloneOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadClusterISOCloneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDownloadClusterISOCloneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadClusterISOCloneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadClusterISOCloneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDownloadClusterISOCloneMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDownloadClusterISOCloneConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadClusterISOCloneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadClusterISOCloneOK creates a DownloadClusterISOCloneOK with default headers values
func NewDownloadClusterISOCloneOK(writer io.Writer) *DownloadClusterISOCloneOK {
	return &DownloadClusterISOCloneOK{
		Payload: writer,
	}
}

/*DownloadClusterISOCloneOK handles this case with default header values.

Success.
*/
type DownloadClusterISOCloneOK struct {
	Payload io.Writer
}

func (o *DownloadClusterISOCloneOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneOK  %+v", 200, o.Payload)
}

func (o *DownloadClusterISOCloneOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadClusterISOCloneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOCloneBadRequest creates a DownloadClusterISOCloneBadRequest with default headers values
func NewDownloadClusterISOCloneBadRequest() *DownloadClusterISOCloneBadRequest {
	return &DownloadClusterISOCloneBadRequest{}
}

/*DownloadClusterISOCloneBadRequest handles this case with default header values.

Error.
*/
type DownloadClusterISOCloneBadRequest struct {
	Payload *models.Error
}

func (o *DownloadClusterISOCloneBadRequest) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadClusterISOCloneBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISOCloneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOCloneUnauthorized creates a DownloadClusterISOCloneUnauthorized with default headers values
func NewDownloadClusterISOCloneUnauthorized() *DownloadClusterISOCloneUnauthorized {
	return &DownloadClusterISOCloneUnauthorized{}
}

/*DownloadClusterISOCloneUnauthorized handles this case with default header values.

Unauthorized.
*/
type DownloadClusterISOCloneUnauthorized struct {
	Payload *models.InfraError
}

func (o *DownloadClusterISOCloneUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneUnauthorized  %+v", 401, o.Payload)
}

func (o *DownloadClusterISOCloneUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadClusterISOCloneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOCloneForbidden creates a DownloadClusterISOCloneForbidden with default headers values
func NewDownloadClusterISOCloneForbidden() *DownloadClusterISOCloneForbidden {
	return &DownloadClusterISOCloneForbidden{}
}

/*DownloadClusterISOCloneForbidden handles this case with default header values.

Forbidden.
*/
type DownloadClusterISOCloneForbidden struct {
	Payload *models.InfraError
}

func (o *DownloadClusterISOCloneForbidden) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneForbidden  %+v", 403, o.Payload)
}

func (o *DownloadClusterISOCloneForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DownloadClusterISOCloneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOCloneNotFound creates a DownloadClusterISOCloneNotFound with default headers values
func NewDownloadClusterISOCloneNotFound() *DownloadClusterISOCloneNotFound {
	return &DownloadClusterISOCloneNotFound{}
}

/*DownloadClusterISOCloneNotFound handles this case with default header values.

Error.
*/
type DownloadClusterISOCloneNotFound struct {
	Payload *models.Error
}

func (o *DownloadClusterISOCloneNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneNotFound  %+v", 404, o.Payload)
}

func (o *DownloadClusterISOCloneNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISOCloneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOCloneMethodNotAllowed creates a DownloadClusterISOCloneMethodNotAllowed with default headers values
func NewDownloadClusterISOCloneMethodNotAllowed() *DownloadClusterISOCloneMethodNotAllowed {
	return &DownloadClusterISOCloneMethodNotAllowed{}
}

/*DownloadClusterISOCloneMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type DownloadClusterISOCloneMethodNotAllowed struct {
}

func (o *DownloadClusterISOCloneMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneMethodNotAllowed ", 405)
}

func (o *DownloadClusterISOCloneMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadClusterISOCloneConflict creates a DownloadClusterISOCloneConflict with default headers values
func NewDownloadClusterISOCloneConflict() *DownloadClusterISOCloneConflict {
	return &DownloadClusterISOCloneConflict{}
}

/*DownloadClusterISOCloneConflict handles this case with default header values.

Error.
*/
type DownloadClusterISOCloneConflict struct {
	Payload *models.Error
}

func (o *DownloadClusterISOCloneConflict) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneConflict  %+v", 409, o.Payload)
}

func (o *DownloadClusterISOCloneConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISOCloneConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadClusterISOCloneInternalServerError creates a DownloadClusterISOCloneInternalServerError with default headers values
func NewDownloadClusterISOCloneInternalServerError() *DownloadClusterISOCloneInternalServerError {
	return &DownloadClusterISOCloneInternalServerError{}
}

/*DownloadClusterISOCloneInternalServerError handles this case with default header values.

Error.
*/
type DownloadClusterISOCloneInternalServerError struct {
	Payload *models.Error
}

func (o *DownloadClusterISOCloneInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{cluster_id}/downloads/image.iso][%d] downloadClusterISOCloneInternalServerError  %+v", 500, o.Payload)
}

func (o *DownloadClusterISOCloneInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownloadClusterISOCloneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
