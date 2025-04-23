package utils

import (
	"errors"
	"pleasurelove/internal/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ReadRequestParamID(c *fiber.Ctx) (int64, error) {
	idStr := c.Params("id")

	idTmp, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid request id")
	}

	return int64(idTmp), nil
}

// Fungsi untuk mengambil filter dan pagination dari request
//
//	=		|	name=John					|	Mencari nilai yang sama (=)
//	!=		|	status__ne=inactive			|	Tidak sama (!=)
//	<		|	age__lt=30					|	Kurang dari (<)
//	>		|	age__gt=18					|	Lebih dari (>)
//	<=		|	age__lte=30					|	Kurang dari atau sama (<=)
//	>=		|	age__gte=18					|	Lebih dari atau sama (>=)
//	LIKE	|	name__like=Jo				|	Mencari yang mengandung teks (LIKE)
//	IN		|	status__in=active,pending	|	Filter banyak nilai (IN)
func GetFiltersAndPagination(c *fiber.Ctx) *models.GetListStruct {
	filters := make(map[string][2]interface{})

	for k, v := range c.Queries() {
		v = strings.TrimSpace(v) // Hindari spasi yang tidak disengaja
		if k == "page" || k == "limit" {
			continue
		}

		// Pisahkan field dan operator (contoh: age__gte=18)
		parts := strings.Split(k, "__")
		field := parts[0]
		operator := "=" // Default operator

		if len(parts) > 1 {
			switch parts[1] {
			case "ne":
				operator = "!="
			case "lt":
				operator = "<"
			case "gt":
				operator = ">"
			case "lte":
				operator = "<="
			case "gte":
				operator = ">="
			case "like":
				operator = "LIKE"
				v = "%" + v + "%"
			case "in": // Handling operator IN
				operator = "IN"
			default:
				continue // Skip jika operator tidak dikenali
			}
		}

		// Konversi tipe data sesuai kebutuhan
		var value interface{}
		if operator == "IN" {
			value = strings.Split(v, ",") // Konversi ke slice untuk IN condition
		} else if intVal, err := strconv.Atoi(v); err == nil {
			value = intVal
		} else if floatVal, err := strconv.ParseFloat(v, 64); err == nil {
			value = floatVal
		} else if boolVal, err := strconv.ParseBool(v); err == nil {
			value = boolVal
		} else {
			value = v
		}

		filters[field] = [2]interface{}{operator, value}
	}

	// Ambil pagination dengan default
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	// Ambil orderBy dan sortBy
	orderBy := c.Query("orderBy", "id") // Default orderBy adalah "id"
	sortBy := c.Query("sortBy", "desc") // Default sortBy adalah "asc"

	return &models.GetListStruct{
		Filters: filters,
		Page:    page,
		Limit:   limit,
		OrderBy: orderBy,
		SortBy:  sortBy,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ExtractBearerToken(authHeader string) string {
	return strings.TrimPrefix(authHeader, "Bearer ")
}
