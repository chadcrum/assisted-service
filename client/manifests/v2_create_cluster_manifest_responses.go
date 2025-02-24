// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2CreateClusterManifestReader is a Reader for the V2CreateClusterManifest structure.
type V2CreateClusterManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2CreateClusterManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2CreateClusterManifestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2CreateClusterManifestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2CreateClusterManifestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2CreateClusterManifestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2CreateClusterManifestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2CreateClusterManifestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2CreateClusterManifestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2CreateClusterManifestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2CreateClusterManifestCreated creates a V2CreateClusterManifestCreated with default headers values
func NewV2CreateClusterManifestCreated() *V2CreateClusterManifestCreated {
	return &V2CreateClusterManifestCreated{}
}

/*V2CreateClusterManifestCreated handles this case with default header values.

Success.
*/
type V2CreateClusterManifestCreated struct {
	Payload *models.Manifest
}

func (o *V2CreateClusterManifestCreated) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestCreated  %+v", 201, o.Payload)
}

func (o *V2CreateClusterManifestCreated) GetPayload() *models.Manifest {
	return o.Payload
}

func (o *V2CreateClusterManifestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestBadRequest creates a V2CreateClusterManifestBadRequest with default headers values
func NewV2CreateClusterManifestBadRequest() *V2CreateClusterManifestBadRequest {
	return &V2CreateClusterManifestBadRequest{}
}

/*V2CreateClusterManifestBadRequest handles this case with default header values.

Error.
*/
type V2CreateClusterManifestBadRequest struct {
	Payload *models.Error
}

func (o *V2CreateClusterManifestBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestBadRequest  %+v", 400, o.Payload)
}

func (o *V2CreateClusterManifestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestUnauthorized creates a V2CreateClusterManifestUnauthorized with default headers values
func NewV2CreateClusterManifestUnauthorized() *V2CreateClusterManifestUnauthorized {
	return &V2CreateClusterManifestUnauthorized{}
}

/*V2CreateClusterManifestUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2CreateClusterManifestUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2CreateClusterManifestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestUnauthorized  %+v", 401, o.Payload)
}

func (o *V2CreateClusterManifestUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2CreateClusterManifestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestForbidden creates a V2CreateClusterManifestForbidden with default headers values
func NewV2CreateClusterManifestForbidden() *V2CreateClusterManifestForbidden {
	return &V2CreateClusterManifestForbidden{}
}

/*V2CreateClusterManifestForbidden handles this case with default header values.

Forbidden.
*/
type V2CreateClusterManifestForbidden struct {
	Payload *models.InfraError
}

func (o *V2CreateClusterManifestForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestForbidden  %+v", 403, o.Payload)
}

func (o *V2CreateClusterManifestForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2CreateClusterManifestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestNotFound creates a V2CreateClusterManifestNotFound with default headers values
func NewV2CreateClusterManifestNotFound() *V2CreateClusterManifestNotFound {
	return &V2CreateClusterManifestNotFound{}
}

/*V2CreateClusterManifestNotFound handles this case with default header values.

Error.
*/
type V2CreateClusterManifestNotFound struct {
	Payload *models.Error
}

func (o *V2CreateClusterManifestNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestNotFound  %+v", 404, o.Payload)
}

func (o *V2CreateClusterManifestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestMethodNotAllowed creates a V2CreateClusterManifestMethodNotAllowed with default headers values
func NewV2CreateClusterManifestMethodNotAllowed() *V2CreateClusterManifestMethodNotAllowed {
	return &V2CreateClusterManifestMethodNotAllowed{}
}

/*V2CreateClusterManifestMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2CreateClusterManifestMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2CreateClusterManifestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2CreateClusterManifestMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestConflict creates a V2CreateClusterManifestConflict with default headers values
func NewV2CreateClusterManifestConflict() *V2CreateClusterManifestConflict {
	return &V2CreateClusterManifestConflict{}
}

/*V2CreateClusterManifestConflict handles this case with default header values.

Error.
*/
type V2CreateClusterManifestConflict struct {
	Payload *models.Error
}

func (o *V2CreateClusterManifestConflict) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestConflict  %+v", 409, o.Payload)
}

func (o *V2CreateClusterManifestConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2CreateClusterManifestInternalServerError creates a V2CreateClusterManifestInternalServerError with default headers values
func NewV2CreateClusterManifestInternalServerError() *V2CreateClusterManifestInternalServerError {
	return &V2CreateClusterManifestInternalServerError{}
}

/*V2CreateClusterManifestInternalServerError handles this case with default header values.

Error.
*/
type V2CreateClusterManifestInternalServerError struct {
	Payload *models.Error
}

func (o *V2CreateClusterManifestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/manifests][%d] v2CreateClusterManifestInternalServerError  %+v", 500, o.Payload)
}

func (o *V2CreateClusterManifestInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2CreateClusterManifestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
