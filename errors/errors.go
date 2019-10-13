package climatehero_errors

import (
	"github.com/go-openapi/runtime"
	"net/http"
)

type Error struct {
	StatusCode int
	Message string
}

func NewError(code int, message ...string) *Error {
	var err Error
	err.StatusCode = code
	if message != nil {
		err.Message = message[0]
	}
	return &err
}

func (self *Error) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	w.WriteHeader(self.StatusCode)
	if self.Message != "" {
		w.Write([]byte(self.Message))
	}
}
