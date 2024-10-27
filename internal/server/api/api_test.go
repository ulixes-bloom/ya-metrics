package api

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ulixes-bloom/ya-metrics/internal/server/config"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body []byte) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, bytes.NewReader(body))
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestUpdateMetric(t *testing.T) {
	type args struct {
		url          string
		method       string
		expectedCode int
	}
	ts := httptest.NewServer(New(config.Config{}).router)
	defer ts.Close()

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Correct request with gauge metric",
			args: args{
				url:          "/update/gauge/gauge_metric/1.9",
				method:       http.MethodPost,
				expectedCode: http.StatusOK,
			},
		},
		{
			name: "Correct request with counter metric",
			args: args{
				url:          "/update/counter/counter_metric/100",
				method:       http.MethodPost,
				expectedCode: http.StatusOK,
			},
		},
		{
			name: "Wrong method",
			args: args{
				url:          "/update/gauge/some/1",
				method:       http.MethodGet,
				expectedCode: http.StatusMethodNotAllowed,
			},
		},
		{
			name: "Wrong metric type",
			args: args{
				url:          "/update/newmetric/some/1",
				method:       http.MethodPost,
				expectedCode: http.StatusBadRequest,
			},
		},
		{
			name: "Wrong counter metric value",
			args: args{
				url:          "/update/counter/some/1.89",
				method:       http.MethodPost,
				expectedCode: http.StatusBadRequest,
			},
		},
		{
			name: "Return all metrics",
			args: args{
				url:          "/",
				method:       http.MethodGet,
				expectedCode: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, _ := testRequest(t, ts, test.args.method, test.args.url, nil)
			defer resp.Body.Close()

			assert.Equal(t, test.args.expectedCode, resp.StatusCode)
		})
	}
}

func TestUpdateJSONMetric(t *testing.T) {
	type args struct {
		url          string
		method       string
		expectedCode int
		body         []byte
	}
	ts := httptest.NewServer(New(config.Config{}).router)
	defer ts.Close()

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Correct request with gauge metric",
			args: args{
				url:          "/update/",
				method:       http.MethodPost,
				expectedCode: http.StatusOK,
				body:         []byte(`{"id":"gaugeeee","type":"gauge","value":13}`),
			},
		},
		{
			name: "Correct request with counter metric",
			args: args{
				url:          "/update/",
				method:       http.MethodPost,
				expectedCode: http.StatusOK,
				body:         []byte(`{"id":"counterrr","type":"counter","delta":13}`),
			},
		},
		{
			name: "Wrong method",
			args: args{
				url:          "/update/",
				method:       http.MethodGet,
				expectedCode: http.StatusMethodNotAllowed,
			},
		},
		{
			name: "Wrong metric type",
			args: args{
				url:          "/update/",
				method:       http.MethodPost,
				expectedCode: http.StatusBadRequest,
				body:         []byte(`{"id":"counter_metric","type":"some","value":13}`),
			},
		},
		{
			name: "Wrong counter metric value",
			args: args{
				url:          "/update/counter/some/1.89",
				method:       http.MethodPost,
				expectedCode: http.StatusBadRequest,
				body:         []byte(`{"id":"counter_metric","type":"counter","value":13.1234}`),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, _ := testRequest(t, ts, test.args.method, test.args.url, test.args.body)
			defer resp.Body.Close()

			assert.Equal(t, test.args.expectedCode, resp.StatusCode)
		})
	}
}

type Metrics struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

func TestGzipCompression(t *testing.T) {
	type args struct {
		url          string
		method       string
		expectedCode int
		body         []byte
	}
	ts := httptest.NewServer(New(config.Config{}).router)
	defer ts.Close()

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Correct request with gauge metric",
			args: args{
				url:          "/update/",
				method:       http.MethodPost,
				expectedCode: http.StatusOK,
				body:         []byte(`{"id":"Getsetzip","type":"gauge","value":13}`),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			gb := gzip.NewWriter(buf)

			_, err := gb.Write([]byte(test.args.body))
			require.NoError(t, err)

			err = gb.Close()
			require.NoError(t, err)

			// create request
			req, err := http.NewRequest(test.args.method, ts.URL+test.args.url, bytes.NewReader(test.args.body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept-Encoding", "gzip")
			require.NoError(t, err)

			// do request
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			require.Equal(t, test.args.expectedCode, resp.StatusCode)

			//check response
			gr, err := gzip.NewReader(resp.Body)
			require.NoError(t, err)
			respBody, err := io.ReadAll(gr)
			require.NoError(t, err)
			require.Equal(t, respBody, test.args.body)
		})
	}
}