package piefed

import (
	piefedResponse "LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/helper"
	"LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
	"encoding/json"
	"fmt"
	"io"
	goHttp "net/http"
	"strings"
)

const applicationJson = "application/json"

func defaultHandler[TResponse any](
	piefed *Piefed,
	urlPath string,
	httpMethod router.HttpMethod,
	request any,
	headers http.Headers,
) (*TResponse, error) {
	resp, err := piefed.sendRequest(
		urlPath,
		httpMethod,
		request,
		headers,
	)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != goHttp.StatusOK {
		var errorResponse *piefedResponse.ErrorResponse
		_ = json.Unmarshal(body, &errorResponse)
		errorResponse.StatusCode = resp.StatusCode

		return nil, errorResponse
	}

	var response TResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type Piefed struct {
	instance string
}

func NewPiefed(instance string) *Piefed {
	return &Piefed{
		instance: instance,
	}
}

func (receiver *Piefed) url() string {
	return fmt.Sprintf("https://%s/api/alpha", receiver.instance)
}

func (receiver *Piefed) sendRequest(
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

	req, err := goHttp.NewRequest(
		string(method),
		receiver.url()+path+queryString,
		body,
	)
	if err != nil {
		return nil, err
	}

	req.Header = helper.MapMap(helper.MapFilter(headers, func(value, key string) bool {
		lowerKey := strings.ToLower(key)
		// Content-Length is recalculated by the Go HTTP client for this
		// request's own body, so a stale client-supplied value must not
		// be forwarded. Accept-Encoding must not be forwarded either —
		// if a client-set value is present, Go's transport disables its
		// own automatic gzip request/decompress handling, and a gzip
		// response body then gets handed to json.Unmarshal as raw bytes
		// (visible as "invalid character '\x1f'..." — 0x1f is gzip's
		// magic byte). Leaving this header unset lets Go manage
		// compression transparently, which is what defaultHandler's
		// plain json.Unmarshal on the response body assumes.
		return lowerKey != "content-length" && lowerKey != "accept-encoding"
	}), func(value, key string) []string {
		return []string{value}
	})
	if body != nil {
		req.Header.Set("Content-Type", applicationJson)
	}

	return goHttp.DefaultClient.Do(req)
}
