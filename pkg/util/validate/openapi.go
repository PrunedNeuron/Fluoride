package validate

import (
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"go.uber.org/zap"
)

// OpenAPI validates an open api spec
func OpenAPI(file string) (bool, error) {
	doc, err := loads.Spec(file[1:])
	if err != nil {
		zap.S().Error("Error loading spec")
		return false, err
	}

	validator := validate.NewSpecValidator(doc.Schema(), strfmt.Default)
	validator.SetContinueOnErrors(true)  // Set option for this validator
	result, _ := validator.Validate(doc) // Validates spec with default Swagger 2.0 format definitions
	if result.IsValid() {
		zap.S().Info("Valid spec")
		return true, nil
	}

	if result.HasWarnings() {
		zap.S().Info("Spec has some validation warnings")
		return true, nil
	}

	zap.S().Error("Spec has some validation errors")
	return false, fmt.Errorf("Spec has some validation errors")

}
