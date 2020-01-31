package errs

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type HTTPError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (e HTTPError) Error() string {
	return "Server error: " + strconv.Itoa(e.Code) + "; Message: " + e.Msg
}

var (
	ErrInvalidJSON     = errors.New("Invalid JSON.")
	ErrInvalidResponse = errors.New("Cannot read response body.")
	ErrInvalidRequest  = errors.New("Cannot send request.")
)

var (
	ErrBadRequest = HTTPError{
		Code: http.StatusBadRequest,
		Msg:  "Bad request: check query parameters.",
	}
	ErrNotFound = HTTPError{
		Code: http.StatusNotFound,
		Msg:  "Nof found.",
	}
	ErrUnauthorized = HTTPError{
		Code: http.StatusUnauthorized,
		Msg:  "Authentication failed: check for valid API key.",
	}
	ErrForbidden = HTTPError{
		Code: http.StatusForbidden,
		Msg:  "Authentication failed: check for valid API key.",
	}
	ErrInternalError = HTTPError{
		Code: http.StatusInternalServerError,
		Msg:  "Internal error: report bug.",
	}
)

func WriteError(w http.ResponseWriter, err error) {
	if err == nil {
		err = errors.Errorf("No error specified.")
	}

	httpError, ok := err.(HTTPError)
	if ok {
		b, _ := json.Marshal(httpError)
		http.Error(w, string(b), httpError.Code)
		return
	}

	http.Error(w, "Something went wrong.", http.StatusInternalServerError)
}

func FindError(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		log.Fatal(resp)
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusInternalServerError:
		return ErrInternalError
	default:
		return ErrInternalError
	}
}
