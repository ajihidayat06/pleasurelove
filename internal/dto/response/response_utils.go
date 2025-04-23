package response

import (
	"net/http"
	"pleasurelove/internal/models"

	"github.com/gofiber/fiber/v2"
)

// ListResponse is a generic struct for API responses that return a list of data.
type ListResponse[T any] struct {
	List             []T      `json:"list"`              // The list of items
	TotalCount       int64    `json:"total_count"`       // Total number of items available
	Page             int      `json:"page"`              // Current page number
	PageSize         int      `json:"page_size"`         // Number of items per page
	FiltersAvailable []string `json:"filters_available"` // Filters applied to the list
}

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func SetResponseAPI(c *fiber.Ctx, status int, message, err string, data interface{}) error {
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}

	return c.Status(status).JSON(response)
}

func SetResponseOK(c *fiber.Ctx, message string, data interface{}) error {
	return SetResponseAPI(c, http.StatusOK, message, "", data)
}

func SetResponseBadRequest(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI(c, http.StatusBadRequest, message, err.Error(), nil)
}

func SetResponseInternalServerError(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI(c, http.StatusInternalServerError, message, err.Error(), nil)
}

func SetResponseUnauthorized(c *fiber.Ctx, message string, err string) error {
	return SetResponseAPI(c, http.StatusUnauthorized, message, err, nil)
}

func SetResponseForbiden(c *fiber.Ctx, message string) error {
	return SetResponseAPI(c, http.StatusForbidden, message, "", nil)
}

func SetResponseNotFound(c *fiber.Ctx, message string, err error) error {
	return SetResponseAPI(c, http.StatusNotFound, message, err.Error(), nil)
}

func MapToListResponse[T any](list []T, totalCount int64, listStruct *models.GetListStruct, filtersAvailable []string) ListResponse[T] {
	return ListResponse[T]{
		List:             list,
		TotalCount:       totalCount,
		Page:             listStruct.Page,
		PageSize:         listStruct.Limit,
		FiltersAvailable: filtersAvailable,
	}
}
