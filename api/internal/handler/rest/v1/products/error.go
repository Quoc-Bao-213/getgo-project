package products

import (
	"net/http"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/httpserver"
)

const ErrCodeValidationFailed = "validation_failed"

var (
	errInvalidName        = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Product Name"}
	errInvalidDescription = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Product Description"}
	errInvalidPrice       = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "Invalid Product Price"}
)
