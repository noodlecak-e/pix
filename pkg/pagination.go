package pkg

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

func NewPaginationFromQuery(ctx *gin.Context) (*Pagination, error) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		return nil, err
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		return nil, err
	}

	return &Pagination{
		Page:  page,
		Limit: limit,
	}, nil
}

func PaginatedResponse(data interface{}, pagination *Pagination) gin.H {
	return gin.H{
		"data":  data,
		"limit": pagination.Limit,
		"page":  pagination.Page,
	}
}
