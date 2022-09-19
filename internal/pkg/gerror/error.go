package gerror

import "google.golang.org/grpc/status"

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-11

// Error defines the structure of error for grpc gateway.
type Error struct {
	Code    uint32        `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Details []interface{} `json:"details,omitempty"`
}

func From(status *status.Status) *Error {
	return &Error{
		Code:    uint32(status.Code()),
		Message: status.Message(),
		Details: status.Details(),
	}
}
