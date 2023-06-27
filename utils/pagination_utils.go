package utils

import (
	"Recipe_App/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) models.PaginationInput {
	limit := 2
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return models.PaginationInput{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
