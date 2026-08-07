package http

import (
	"LemmyBeProxy/json"
	"log"
	"net/http"
	"strconv"
)

// WriteHttpResponse sets Content-Length explicitly rather than letting
// Go's server fall back to Transfer-Encoding: chunked (its default when
// no length is set). Confirmed directly: chunked responses from this
// server were read as truncated by Rust's reqwest client (lemmyBB),
// while curl read the exact same response correctly — an interop edge
// case in how reqwest's reader handles chunked encoding from this
// specific server, not a content bug. Setting Content-Length sidesteps
// the whole chunked-encoding code path rather than chasing the
// underlying interop cause.
func WriteHttpResponse(response *Response, writer http.ResponseWriter) {
	body := response.Body
	headers := response.Headers
	statusCode := response.StatusCode

	if headers == nil {
		headers = make(map[string]string)
	}
	if body == nil {
		body = make(map[string]string)
	}
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	_, ok := headers["Content-Type"]
	if !ok {
		headers["Content-Type"] = "application/json"
	}

	var err error
	if _, ok = body.(string); !ok {
		body, err = json.ToJson(body)
		if err != nil {
			body, _ = json.ToJson(map[string]string{
				"error": "Internal request error",
			})
			statusCode = http.StatusInternalServerError
			log.Println(err)
		}
	} else {
		body = []byte(body.(string))
	}

	bodyBytes := body.([]byte)
	headers["Content-Length"] = strconv.Itoa(len(bodyBytes))

	for key, value := range headers {
		writer.Header().Set(key, value)
	}
	writer.WriteHeader(statusCode)

	_, err = writer.Write(bodyBytes)
	if err != nil {
		log.Println(err)
	}
}
