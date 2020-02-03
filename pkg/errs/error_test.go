package errs

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestWriteError(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		err error
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Nil error",
			args{
				&httptest.ResponseRecorder{},
				nil,
			},
			"level=info msg=\"No error specified.\"\n",
		},
		{
			"Error invalid request",
			args{
				&httptest.ResponseRecorder{},
				ErrInvalidRequest,
			},
			"level=info msg=\"Cannot send request.\"\n",
		},
		{
			"Error unauthorized",
			args{
				&httptest.ResponseRecorder{},
				ErrUnauthorized,
			},
			"level=info msg=\"Server error: 401; Message: Authentication failed: check for valid API key.\"\n",
		},
		{
			"Not found",
			args{
				&httptest.ResponseRecorder{},
				ErrNotFound,
			},
			"level=info msg=\"Server error: 404; Message: Not found.\"\n",
		},
		{
			"Internal server error",
			args{
				&httptest.ResponseRecorder{},
				ErrInternalError,
			},
			"level=info msg=\"Server error: 500; Message: Something went wrong.\"\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			logrus.SetFormatter(&logrus.TextFormatter{
				DisableColors:    true,
				DisableTimestamp: true,
			})

			logrus.SetOutput(&buf)
			defer func() {
				logrus.SetOutput(os.Stderr)
				logrus.SetFormatter(&logrus.TextFormatter{})
			}()

			WriteError(tt.args.w, tt.args.err)
			got := buf.String()

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
