// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// BindUserToClusterRoleV2Reader is a Reader for the BindUserToClusterRoleV2 structure.
type BindUserToClusterRoleV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindUserToClusterRoleV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBindUserToClusterRoleV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBindUserToClusterRoleV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBindUserToClusterRoleV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBindUserToClusterRoleV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBindUserToClusterRoleV2OK creates a BindUserToClusterRoleV2OK with default headers values
func NewBindUserToClusterRoleV2OK() *BindUserToClusterRoleV2OK {
	return &BindUserToClusterRoleV2OK{}
}

/* BindUserToClusterRoleV2OK describes a response with status code 200, with default header values.

ClusterRoleBinding
*/
type BindUserToClusterRoleV2OK struct {
	Payload *models.ClusterRoleBinding
}

func (o *BindUserToClusterRoleV2OK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleV2OK  %+v", 200, o.Payload)
}
func (o *BindUserToClusterRoleV2OK) GetPayload() *models.ClusterRoleBinding {
	return o.Payload
}

func (o *BindUserToClusterRoleV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindUserToClusterRoleV2Unauthorized creates a BindUserToClusterRoleV2Unauthorized with default headers values
func NewBindUserToClusterRoleV2Unauthorized() *BindUserToClusterRoleV2Unauthorized {
	return &BindUserToClusterRoleV2Unauthorized{}
}

/* BindUserToClusterRoleV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type BindUserToClusterRoleV2Unauthorized struct {
}

func (o *BindUserToClusterRoleV2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleV2Unauthorized ", 401)
}

func (o *BindUserToClusterRoleV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToClusterRoleV2Forbidden creates a BindUserToClusterRoleV2Forbidden with default headers values
func NewBindUserToClusterRoleV2Forbidden() *BindUserToClusterRoleV2Forbidden {
	return &BindUserToClusterRoleV2Forbidden{}
}

/* BindUserToClusterRoleV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type BindUserToClusterRoleV2Forbidden struct {
}

func (o *BindUserToClusterRoleV2Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleV2Forbidden ", 403)
}

func (o *BindUserToClusterRoleV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToClusterRoleV2Default creates a BindUserToClusterRoleV2Default with default headers values
func NewBindUserToClusterRoleV2Default(code int) *BindUserToClusterRoleV2Default {
	return &BindUserToClusterRoleV2Default{
		_statusCode: code,
	}
}

/* BindUserToClusterRoleV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type BindUserToClusterRoleV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the bind user to cluster role v2 default response
func (o *BindUserToClusterRoleV2Default) Code() int {
	return o._statusCode
}

func (o *BindUserToClusterRoleV2Default) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleV2 default  %+v", o._statusCode, o.Payload)
}
func (o *BindUserToClusterRoleV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BindUserToClusterRoleV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
