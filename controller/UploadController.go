package controller

import (
	"LemmyBeProxy/http"
	"LemmyBeProxy/service/backend"
	"bytes"
	"encoding/hex"
	"io"
	"mime"
	"mime/multipart"
	goHttp "net/http"
	"strings"
)

// UploadController is now thin — the raw multipart parsing and the
// hex-token/redirect mechanism (built to work with Piefed's upload API,
// which doesn't return a directly servable URL) both stay here since
// they work identically regardless of which backend actually produced
// the resulting URL. Only the actual upload call is backend-specific.
type UploadController struct {
	backend backend.Backend
}

func NewUploadController(backend backend.Backend) *UploadController {
	return &UploadController{
		backend: backend,
	}
}

// extractCookieJwt pulls the jwt value out of a raw Cookie header, since
// mlmym authenticates its image upload via a Cookie (jwt=...) rather than
// an Authorization: Bearer header like every other request it makes.
func extractCookieJwt(cookieHeader string) string {
	for _, part := range strings.Split(cookieHeader, ";") {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "jwt=") {
			return strings.TrimPrefix(part, "jwt=")
		}
	}
	return ""
}

// UploadImage handles POST /pictrs/image — a route deliberately outside
// /api/v3, since that's where mlmym (and real Lemmy pict-rs) actually sends
// uploads. It parses the incoming multipart body itself (defaultHandler is
// JSON-only), forwards the file via the configured backend, and responds
// in pict-rs's own response shape so mlmym's existing parsing code works
// unmodified regardless of which backend is actually running.
func (receiver *UploadController) UploadImage(request *http.Request) (*http.Response, error) {
	contentType := request.Headers["Content-Type"]
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return &http.Response{
			StatusCode: goHttp.StatusBadRequest,
			Body:       map[string]string{"msg": "invalid content-type for multipart upload"},
		}, nil
	}
	boundary, ok := params["boundary"]
	if !ok {
		return &http.Response{
			StatusCode: goHttp.StatusBadRequest,
			Body:       map[string]string{"msg": "missing multipart boundary"},
		}, nil
	}

	reader := multipart.NewReader(bytes.NewReader(request.Body), boundary)
	var fileBytes []byte
	var filename string
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return &http.Response{
				StatusCode: goHttp.StatusBadRequest,
				Body:       map[string]string{"msg": "malformed multipart body"},
			}, nil
		}
		// mlmym sends the file under the field name "images[]"
		if part.FormName() == "images[]" {
			fileBytes, err = io.ReadAll(part)
			if err != nil {
				return nil, err
			}
			filename = part.FileName()
			break
		}
	}

	if fileBytes == nil {
		return &http.Response{
			StatusCode: goHttp.StatusBadRequest,
			Body:       map[string]string{"msg": "no file found in upload (expected field images[])"},
		}, nil
	}

	jwt := extractCookieJwt(request.Headers["Cookie"])
	if jwt == "" {
		return &http.Response{
			StatusCode: goHttp.StatusUnauthorized,
			Body:       map[string]string{"msg": "missing jwt cookie"},
		}, nil
	}

	fileURL, err := receiver.backend.UploadImage(fileBytes, filename, jwt)
	if err != nil {
		return &http.Response{
			StatusCode: goHttp.StatusBadRequest,
			Body:       map[string]string{"msg": err.Error()},
		}, nil
	}

	// Encode the resulting image URL as a lowercase hex token so it can
	// be handed back in pict-rs's own response shape. mlmym builds the
	// final <img> URL itself as /pictrs/image/{file} — the fake ".jpg"
	// extension is there purely so mlmym's own pictrs-URL regex (which
	// requires [a-z0-9-]+.[a-z]+) recognizes it and applies its
	// thumbnail query params, even though those params have no effect
	// once we redirect to the backend's real image (see ServeImage
	// below). This encoding step is backend-agnostic — it works the
	// same whether the URL came from Piefed or real Lemmy.
	token := hex.EncodeToString([]byte(fileURL))

	return &http.Response{
		StatusCode: goHttp.StatusOK,
		Body: map[string]any{
			"msg": "ok",
			"files": []map[string]string{
				{
					"file":         token + ".jpg",
					"delete_token": "unsupported",
				},
			},
		},
	}, nil
}

// ServeImage handles GET /pictrs/image/{token} — decodes the token mlmym
// requests (built from what UploadImage returned) back into the real
// image URL and redirects there. Against a Piefed backend, this is a
// real, permanent limitation: since we redirect straight to Piefed's
// original image, mlmym's thumbnail query params (?format=jpg&thumbnail=96)
// have no effect — Piefed doesn't understand Lemmy's pict-rs thumbnailing
// convention, so images display at full original size rather than being
// resized down. Against a real Lemmy backend this isn't a limitation at
// all — it's the same pict-rs software, so the thumbnail params redirect
// straight through and work exactly as they would talking to Lemmy directly.
func (receiver *UploadController) ServeImage(request *http.Request) (*http.Response, error) {
	token := request.RouteParams["token"]
	if idx := strings.LastIndex(token, "."); idx != -1 {
		token = token[:idx]
	}

	decoded, err := hex.DecodeString(token)
	if err != nil {
		return http.NotFoundProxyError(), nil
	}

	return &http.Response{
		StatusCode: goHttp.StatusFound,
		Headers: map[string]string{
			"Location": string(decoded),
		},
		Body: "",
	}, nil
}
