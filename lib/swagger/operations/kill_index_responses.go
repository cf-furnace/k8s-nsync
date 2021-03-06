package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*KillIndexAccepted Request to kill index accepted

swagger:response killIndexAccepted
*/
type KillIndexAccepted struct {
}

// NewKillIndexAccepted creates KillIndexAccepted with default headers values
func NewKillIndexAccepted() *KillIndexAccepted {
	return &KillIndexAccepted{}
}

// WriteResponse to the client
func (o *KillIndexAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
}

/*KillIndexBadRequest Missing process guid or missing/invalid index

swagger:response killIndexBadRequest
*/
type KillIndexBadRequest struct {
}

// NewKillIndexBadRequest creates KillIndexBadRequest with default headers values
func NewKillIndexBadRequest() *KillIndexBadRequest {
	return &KillIndexBadRequest{}
}

// WriteResponse to the client
func (o *KillIndexBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

/*KillIndexNotFound Process and index combination not found

swagger:response killIndexNotFound
*/
type KillIndexNotFound struct {
}

// NewKillIndexNotFound creates KillIndexNotFound with default headers values
func NewKillIndexNotFound() *KillIndexNotFound {
	return &KillIndexNotFound{}
}

// WriteResponse to the client
func (o *KillIndexNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*KillIndexServiceUnavailable Service unavailable

swagger:response killIndexServiceUnavailable
*/
type KillIndexServiceUnavailable struct {
}

// NewKillIndexServiceUnavailable creates KillIndexServiceUnavailable with default headers values
func NewKillIndexServiceUnavailable() *KillIndexServiceUnavailable {
	return &KillIndexServiceUnavailable{}
}

// WriteResponse to the client
func (o *KillIndexServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
}
