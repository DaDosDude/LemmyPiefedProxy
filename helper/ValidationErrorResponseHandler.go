package helper

import (
	"LemmyBeProxy/dto/model"
	lemmyModel "LemmyBeProxy/dto/model/lemmy"
	lemmyResponse "LemmyBeProxy/dto/response/lemmy"
	"LemmyBeProxy/http"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	goHttp "net/http"
	"reflect"
	"strings"
)

func ConvertValidationErrorsToResponse(err error) *http.Response {
	var syntaxError *json.SyntaxError
	ok := errors.As(err, &syntaxError)
	if ok {
		return &http.Response{
			StatusCode: goHttp.StatusBadRequest,
			Body: lemmyResponse.NewErrorResponseWithMessage(
				lemmyModel.ErrorCodeUnknown,
				syntaxError.Error(),
			),
		}
	}

	var validationErrors validator.ValidationErrors
	ok = errors.As(err, &validationErrors)
	if ok {
		return &http.Response{
			StatusCode: goHttp.StatusBadRequest,
			Body:       lemmyResponse.NewErrorResponseWithMessage(lemmyModel.ErrorCodeUnknown, "The provided request body is invalid."),
		}
	}

	var internalValidationError *model.ValidationError
	ok = errors.As(err, &internalValidationError)
	if !ok {
		return http.InternalProxyError()
	}

	parsedStruct := reflect.TypeOf(internalValidationError.Input)
	if parsedStruct.Kind() == reflect.Ptr {
		parsedStruct = parsedStruct.Elem()
	}

	errStr := ""
	for _, violation := range internalValidationError.Err {
		field, ok := parsedStruct.FieldByName(violation.StructField())
		if !ok {
			log.Printf("Failed finding struct field %s in struct %s\n", violation.StructField(), parsedStruct.Name())
			continue
		}
		tag := field.Tag.Get("json")
		parts := strings.Split(tag, ",")
		if parts[0] == "" {
			continue
		}

		errStr += fmt.Sprintf("the value for field '%s'", parts[0])
		if violation.ActualTag() == "required" {
			errStr += " is required"
		} else {
			errStr += " is invalid"
		}
		errStr += ", "
	}
	if len(errStr) > 0 {
		errStr = errStr[:len(errStr)-2]
	} else {
		errStr = "Invalid request payload."
	}

	return &http.Response{
		StatusCode: goHttp.StatusBadRequest,
		Body:       lemmyResponse.NewErrorResponseWithMessage(lemmyModel.ErrorCodeUnknown, errStr),
	}
}
