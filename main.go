package main

import (
	"net/http"
	"txtreader/utils"

	"github.com/gin-gonic/gin"
)

type PageRequest struct {
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	FilePath string `json:"f_path"`
}

type PageResponse struct {
	Content    string `json:"content"`
	TotalPages int    `json:"total_pages"`
}

func main() {
	router := gin.Default()
	router.Use(Cors()) // 使用跨域中间件

	router.POST("/read-txt", func(c *gin.Context) {
		var req PageRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		content, totalPages, err := utils.ReadTxtByPage(req.FilePath, req.Page, req.PerPage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		resp := PageResponse{
			Content:    content,
			TotalPages: totalPages,
		}

		c.JSON(http.StatusOK, resp)
	})

	router.Run(":8088")
}

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
