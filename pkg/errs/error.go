package errs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type HTTPError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (e HTTPError) Error() string {
	return "Server error: " + strconv.Itoa(e.Code) + "; Message: " + e.Msg
}

var (
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
		Msg:  "Not found.",
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
		Msg:  "Something went wrong.",
	}
)

func WriteError(w http.ResponseWriter, err error) {
	if err == nil {
		err = errors.Errorf("No error specified.")
	}

	logrus.Print(err.Error())
	httpError, ok := err.(HTTPError)
	if ok {
		err = httpError
		w.WriteHeader(httpError.Code)
		_ = json.NewEncoder(w).Encode(httpError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(ErrInternalError)
}

func CheckResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusInternalServerError:
		return ErrInternalError
	default:
		return ErrInternalError
	}
}
