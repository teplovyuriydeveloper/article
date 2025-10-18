package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teplovyuriydeveloper/article/models"
)

func CreateArticle(c *gin.Context) {
	var article models.Article

	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := article.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "article created", "article": article})
}
