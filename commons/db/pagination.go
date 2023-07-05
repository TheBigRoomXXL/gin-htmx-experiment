package db

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const PageDefault = 1
const PageSizeDefault = 20

type PaginateParams struct {
	Page     int `form:"page"  binding:"gte=1"`
	PageSize int `form:"page_size" binding:"gte=1,lte=100"`
}

func GetPaginate(c *gin.Context) (func(db *gorm.DB) *gorm.DB, error) {
	params := PaginateParams{
		Page:     PageDefault,
		PageSize: PageSizeDefault,
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		return nil, err
	}
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.Page - 1) * params.PageSize
		return db.Offset(offset).Limit(params.PageSize)
	}, nil
}
