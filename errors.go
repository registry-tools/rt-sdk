package sdk

import (
	"errors"
	"fmt"

	"github.com/registry-tools/rt-sdk/generated/models"
)

func IsNotFoundError(err error) bool {
	if modelError, ok := err.(*models.Errors); ok {
		return modelError.ResponseStatusCode == 404
	}
	return false
}

func FormatAPIError(err error) error {
	if err == nil {
		return nil
	}

	result := ""
	if modelError, ok := err.(*models.Errors); ok {
		for _, e := range modelError.GetErrors() {
			if len(result) > 0 {
				result += "\n"
			}
			result += fmt.Sprintf("%s, %s", *e.GetTitle(), *e.GetDetail())
		}
	} else {
		result = fmt.Sprintf("unknown error %q -- This error should have been known by the SDK but was not", err.Error())
	}
	return errors.New(result)
}
