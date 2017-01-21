// Package detonate provides simple HTTP errors
package detonate

import (
	"encoding/json"
	"net/http"
)

// Validation is a struct representing additional validation details
type Validation struct {
	Key     string `json:"key"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Detonation is the type for housing code, error, and message
type Detonation struct {
	Code        int          `json:"code"`
	Error       string       `json:"error"`
	Message     string       `json:"message"`
	Validations []Validation `json:"validations"`
}

// ToJSON is a basic wrapper around marshaling
func (d *Detonation) ToJSON() ([]byte, error) {
	j, err := json.Marshal(d)

	return j, err
}

// AddValidation is a methodology to add details to the output
func (d *Detonation) AddValidation(v Validation) {
	d.Validations = append(d.Validations, v)
}

// Trigger blows up the detonation and writes the error to the supplied writer
func (d *Detonation) Trigger(w http.ResponseWriter) {
	j, err := d.ToJSON()
	if err != nil {
		http.Error(w, "System Error", 500)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(d.Code)
	w.Write(j)

	return
}

// Create lets you manually generate a detonation
func Create(code int, name string, message string) Detonation {
	var d = new(Detonation)

	d.Code = code
	d.Error = name
	d.Message = message

	return *d
}

// BadRequest is for generic 400 level errors
func BadRequest(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 400
	d.Error = "Bad Request"
	d.Message = message

	return d
}

// Unauthorized is for 401 errors
func Unauthorized(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 401
	d.Error = "Unauthorized"
	d.Message = message

	return d
}

// PaymentRequired is for 402 errors
func PaymentRequired(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 402
	d.Error = "Payment Required"
	d.Message = message

	return d
}

// Forbidden is for 403 errors
func Forbidden(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 403
	d.Error = "Forbidden"
	d.Message = message

	return d
}

// NotFound is for 404 errors
func NotFound(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 404
	d.Error = "Not Found"
	d.Message = message

	return d
}

// MethodNotAllowed is for 405 errors
func MethodNotAllowed(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 405
	d.Error = "Method Not Allowed"
	d.Message = message

	return d
}

// NotAcceptable is for 406 errors
func NotAcceptable(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 406
	d.Error = "Not Acceptable"
	d.Message = message

	return d
}

// ProxyAuthRequired is for 407 errors
func ProxyAuthRequired(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 407
	d.Error = "Proxy Authentication Required"
	d.Message = message

	return d
}

// ClientTimeout is for 408 errors
func ClientTimeout(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 408
	d.Error = "Request Time-out"
	d.Message = message

	return d
}

// Conflict is for 409 errors
func Conflict(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 409
	d.Error = "Conflict"
	d.Message = message

	return d
}

// ResourceGone is for 410 errors
func ResourceGone(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 410
	d.Error = "Gone"
	d.Message = message

	return d
}

// LengthRequired is for 411 errors
func LengthRequired(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 411
	d.Error = "Length Required"
	d.Message = message

	return d
}

// PreconditionFailed is for 412 errors
func PreconditionFailed(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 412
	d.Error = "Precondition Failed"
	d.Message = message

	return d
}

// EntityTooLarge is for 413 errors
func EntityTooLarge(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 413
	d.Error = "Request Entity Too Large"
	d.Message = message

	return d
}

// UriTooLong is for 414 errors
func UriTooLong(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 414
	d.Error = "Request-URI Too Large"
	d.Message = message

	return d
}

// UnsupportedMediaType is for 415 errors
func UnsupportedMediaType(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 415
	d.Error = "Unsupported Media Type"
	d.Message = message

	return d
}

// RangeNotSatisfiable is for 416 errors
func RangeNotSatisfiable(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 416
	d.Error = "Requested Range Not Satisfiable"
	d.Message = message

	return d
}

// ExpectationFailed is for 417 errors
func ExpectationFailed(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 417
	d.Error = "Expectation Failed"
	d.Message = message

	return d
}

// BadData is for 422 errors
func BadData(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 422
	d.Error = "Unprocessable Entity"
	d.Message = message

	return d
}

// Locked is for 423 errors
func Locked(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 423
	d.Error = "Locked"
	d.Message = message

	return d
}

// PreconditionRequired is for 428 errors
func PreconditionRequired(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 428
	d.Error = "Precondition Required"
	d.Message = message

	return d
}

// TooManyRequests is for 429 errors
func TooManyRequests(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 429
	d.Error = "Too Many Requests"
	d.Message = message

	return d
}

// Illegal is for 451 errors
func Illegal(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 451
	d.Error = "Unavailable For Legal Reasons"
	d.Message = message

	return d
}

// BadImplementation is for 500 errors
func BadImplementation(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 500
	d.Error = "Internal Server Error"
	d.Message = message

	return d
}

// Internal is an alias (to BadImplementation) for 500 errors
func Internal(message string) *Detonation {
	var d = BadImplementation(message)

	return d
}

// NotImplemented is for 501 errors
func NotImplemented(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 501
	d.Error = "Not Implemented"
	d.Message = message

	return d
}

// BadGateway is for 502 errors
func BadGateway(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 502
	d.Error = "Bad Gateway"
	d.Message = message

	return d
}

// ServerUnavailable is for 503 errors
func ServerUnavailable(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 503
	d.Error = "Service Unavailable"
	d.Message = message

	return d
}

// GatewayTimeout is for 504 errors
func GatewayTimeout(message string) *Detonation {
	var d = new(Detonation)

	d.Code = 504
	d.Error = "Gateway Time-out"
	d.Message = message

	return d
}
