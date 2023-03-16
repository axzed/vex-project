package model

import "github.com/gin-gonic/gin"

// Page 分页主要的两个参数构造成结构体 到时候可以直接用context的bind来绑定
type Page struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int64 `json:"pageSize" form:"pageSize"`
}

// Bind 绑定分页参数
func (p *Page) Bind(ctx *gin.Context) {
	_ = ctx.ShouldBindQuery(&p)
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
