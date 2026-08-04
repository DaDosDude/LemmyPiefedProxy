package lemmy

import (
	"LemmyBeProxy/helper"
	"LemmyBeProxy/http"
	"LemmyBeProxy/router"
	"encoding/json"
	"fmt"
	"io"
	"log"
	goHttp "net/http"
	"strings"
)

const applicationJson = "application/json"

// truncateForLog caps how much of a response body gets logged.
func truncateForLog(body []byte) string {
	const maxLen = 500
	if len(body) > maxLen {
		return string(body[:maxLen]) + "... (truncated)"
	}
	return string(body)
}

// defaultHandler mirrors service/piefed's version exactly — same
// nil-pointer safety on unparseable error bodies, same shape-mismatch
// logging on the success path. Real Lemmy's error shape is simpler than
// Piefed's (just {"error": "code_string"}), which LemmyApiError models —
// a type separate from dto/response/lemmy.ErrorResponse (this proxy's own
// output shape to its callers), since that one has a field named Error,
// not a method, and can't satisfy Go's error interface.
func defaultHandler[TResponse any](
	client *Lemmy,
	urlPath string,
	httpMethod router.HttpMethod,
	request any,
	headers http.Headers,
) (*TResponse, error) {
	resp, err := client.sendRequest(urlPath, httpMethod, request, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != goHttp.StatusOK {
		var errorResponse *LemmyApiError
		if err := json.Unmarshal(body, &errorResponse); err != nil || errorResponse == nil {
			errorResponse = &LemmyApiError{
				ErrorCode: fmt.Sprintf("unparseable_error_response (http %d)", resp.StatusCode),
			}
		}
		errorResponse.StatusCode = resp.StatusCode

		log.Printf(
			"lemmy error response: %s %s -> HTTP %d, error=%q, body=%s",
			httpMethod, resp.Request.URL.String(), resp.StatusCode, errorResponse.ErrorCode, truncateForLog(body),
		)

		return nil, errorResponse
	}

	var response TResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf(
			"lemmy response unmarshal failed: %s %s -> %v, body=%s",
			httpMethod, urlPath, err, truncateForLog(body),
		)
		return nil, err
	}

	return &response, nil
}

// Lemmy is the low-level client for a real Lemmy instance's /api/v3.
// Unlike the Piefed client, this needs almost no shape translation at the
// backend.Backend layer above it, since this proxy's canonical DTOs are
// already modeled on real Lemmy's actual field names.
type Lemmy struct {
	instance string
}

func NewLemmy(instance string) *Lemmy {
	return &Lemmy{instance: instance}
}

func (receiver *Lemmy) url() string {
	return fmt.Sprintf("https://%s/api/v3", receiver.instance)
}

func (receiver *Lemmy) sendRequest(
	path string,
	method router.HttpMethod,
	request any,
	headers http.Headers,
) (*goHttp.Response, error) {
	var body io.Reader
	queryString := ""
	if request != nil {
		if method == router.HttpMethodGet {
			marshalled, err := helper.MarshalToQueryString(request)
			if err != nil {
				return nil, err
			}
			queryString = "?" + marshalled
		} else {
			body = helper.MarshalToReader(request)
		}
	}

	req, err := goHttp.NewRequest(string(method), receiver.url()+path+queryString, body)
	if err != nil {
		return nil, err
	}

	log.Printf("lemmy request: %s %s", method, receiver.url()+path+queryString)

	req.Header = helper.MapMap(helper.MapFilter(headers, func(value, key string) bool {
		lowerKey := strings.ToLower(key)
		// Same fix as the Piefed client: never forward Content-Length or
		// Accept-Encoding, or Go's transport stops managing gzip
		// transparently and a compressed body gets handed straight to
		// json.Unmarshal.
		return lowerKey != "content-length" && lowerKey != "accept-encoding"
	}), func(value, key string) []string {
		return []string{value}
	})
	if body != nil {
		req.Header.Set("Content-Type", applicationJson)
	}

	return goHttp.DefaultClient.Do(req)
}
