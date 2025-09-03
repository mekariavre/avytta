package avlogem

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// HTTPResponseData holds serializable data from an http.Response
type HTTPResponseData struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Header     http.Header `json:"header"`
	Body       any         `json:"body,omitempty"`
}

// SerializeHTTPResponse extracts status, headers, and body from *http.Response without consuming the body.
func SerializeHTTPResponse(resp *http.Response) HTTPResponseData {
	var bodyCopy string
	if resp.Body != nil {
		b, err := io.ReadAll(resp.Body)
		if err == nil {
			bodyCopy = string(b)
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
		} else {
			bodyCopy = "<error reading body>"
		}
	}
	return HTTPResponseData{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       ParsePipeJSON(bodyCopy),
	}
}

// HTTPRequestData holds serializable data from an http.Request
type HTTPRequestData struct {
	Method string      `json:"method"`
	URL    string      `json:"url"`
	Header http.Header `json:"header"`
	Body   any         `json:"body,omitempty"`
}

// SerializeHTTPRequest extracts method, url, header, and body from *http.Request without consuming the body.
func SerializeHTTPRequest(req *http.Request) HTTPRequestData {
	var bodyCopy string
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err == nil {
			bodyCopy = string(b)
			req.Body = io.NopCloser(bytes.NewBuffer(b))
		} else {
			bodyCopy = "<error reading body>"
		}
	}
	return HTTPRequestData{
		Method: req.Method,
		URL:    req.URL.String(),
		Header: req.Header,
		Body:   ParsePipeJSON(bodyCopy),
	}
}

// will try to parse orig as string and json and return map
// if it fail it return original data
func ParsePipeJSON(orig any) any {
	var m map[string]any
	b, err := json.Marshal(orig)
	if err == nil {
		if err := json.Unmarshal(b, &m); err == nil {
			return m
		}
	}
	return orig
}
