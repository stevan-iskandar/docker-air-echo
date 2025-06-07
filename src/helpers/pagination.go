package helpers

import (
	"fmt"
	"docker-air-echo/database"
	"docker-air-echo/structs"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func Pagination(c echo.Context, data interface{}, filterMap map[string]interface{}) (*structs.Pagination, error) {
	db := database.DB()
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(c.QueryParam("page"))          // Page number
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size")) // Items per page
	sorts := c.QueryParams()["sort"]                       // Get an array of sorting parameters

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10 // Default page size
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Clone the base query for counting and data retrieval
	query := db

	// Apply filters
	if filterMap != nil {
		query = query.Where(filterMap)
	}

	// Count total records
	var total int64
	if err := query.Model(data).Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply sorting based on the request parameters
	for _, sortParam := range sorts {
		// Split the sort parameter into field and order
		parts := strings.Split(sortParam, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid sort parameter")
		}

		// Check if the order is valid
		if parts[1] != "asc" && parts[1] != "desc" {
			return nil, fmt.Errorf("invalid sort order")
		}

		field := parts[0]
		order := parts[1]

		// Apply sorting to query
		query = query.Order(field + " " + order)
	}

	// Apply pagination and retrieve data
	if err := query.Offset(offset).Limit(pageSize).Find(data).Error; err != nil {
		return nil, err
	}

	return &structs.Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     data,
	}, nil
}
