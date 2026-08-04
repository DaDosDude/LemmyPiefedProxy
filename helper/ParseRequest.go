package helper

import (
	"LemmyBeProxy/dto/model"
	"LemmyBeProxy/http"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

var validate = validator.New()

func validateDto(result any) (err error) {
	err = validate.Struct(result)
	if err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			err = &model.ValidationError{
				Err:   validationErrs,
				Input: result,
			}
		}
	}

	return
}

func ParseRequest[T any](request *http.Request) (*T, error) {
	var result T
	err := json.Unmarshal(request.Body, &result)
	if err != nil {
		return nil, err
	}

	err = validateDto(result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ParseRequestQuery[T any](request *http.Request) (*T, error) {
	normalized := make(map[string]any)
	for key, val := range request.QueryParams {
		lowerVal := strings.ToLower(val)

		if lowerVal == "true" || lowerVal == "false" {
			normalized[key] = lowerVal == "true"
		} else if intVal, err := strconv.Atoi(val); err == nil {
			normalized[key] = intVal
		} else if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
			normalized[key] = floatVal
		} else {
			normalized[key] = val
		}
	}

	jsonBytes, err := json.Marshal(normalized)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	err = validateDto(result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
