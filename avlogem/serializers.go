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
//
// It extracts the body as a string, and tries to parse it as JSON. If the parse fails, it sets the Body field
// to the original string.
// SerializeHTTPResponse reads and serializes an http.Response into an HTTPResponseData struct.
// It safely reads the response body, restores it for further use, and parses the body content
// using ParsePipeJSON. If reading the body fails, it sets the body to a placeholder error string.
// The function also copies the response status, status code, and headers.
//
// Parameters:
//   - resp: Pointer to an http.Response to be serialized.
//
// Returns:
//   - HTTPResponseData: Struct containing the status, status code, headers, and parsed body of the response.
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
		Body:       parsejson(bodyCopy),
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
//
// It extracts the body as a string, and tries to parse it as JSON. If the parse fails, it sets the Body field
// to the original string.
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
		Body:   parsejson(bodyCopy),
	}
}

// will try to parse orig as string and json and return map
// if it fail it return original data
func parsejson(orig any) any {
	origB, ok := orig.(string)
	if !ok {
		return orig
	}
	var m map[string]any
	if err := json.Unmarshal([]byte(origB), &m); err == nil {
		return m
	}
	return orig
}
