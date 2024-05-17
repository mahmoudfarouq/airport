package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_itineraryHandler(t *testing.T) {
	e := echo.New()

	type args struct {
		body string
	}

	type want struct {
		statusCode int
		response   string
	}

	tests := map[string]struct {
		args args
		want want
	}{
		"failure: empty request": {
			args: args{
				body: "[]",
			},
			want: want{
				statusCode: http.StatusBadRequest,
				response:   `{"error":"empty request"}`,
			},
		},
		"failure: some tuples not of length 2": {
			args: args{
				body: `[["A", "B"], ["C"]]`,
			},
			want: want{
				statusCode: http.StatusBadRequest,
				response:   `{"error":"invalid tuple","tuple":["C"]}`,
			},
		},
		"success": {
			args: args{
				body: `[["LAX", "DXB"], ["JFK", "LAX"], ["SFO", "SJC"], ["DXB", "SFO"]]`,
			},
			want: want{
				statusCode: http.StatusOK,
				response:   `["JFK","LAX","DXB","SFO","SJC"]`,
			},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.args.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			err := itineraryHandler(c)
			require.NoError(t, err)

			assert.Equal(t, tt.want.statusCode, rec.Code)
			assert.Equal(t, tt.want.response, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
