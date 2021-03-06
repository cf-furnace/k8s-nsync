package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DesireTaskAccepted Task request accepted

swagger:response desireTaskAccepted
*/
type DesireTaskAccepted struct {
}

// NewDesireTaskAccepted creates DesireTaskAccepted with default headers values
func NewDesireTaskAccepted() *DesireTaskAccepted {
	return &DesireTaskAccepted{}
}

// WriteResponse to the client
func (o *DesireTaskAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
}

/*DesireTaskBadRequest Bad request

swagger:response desireTaskBadRequest
*/
type DesireTaskBadRequest struct {
}

// NewDesireTaskBadRequest creates DesireTaskBadRequest with default headers values
func NewDesireTaskBadRequest() *DesireTaskBadRequest {
	return &DesireTaskBadRequest{}
}

// WriteResponse to the client
func (o *DesireTaskBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}
