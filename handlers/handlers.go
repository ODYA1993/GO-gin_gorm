package handlers

import (
	"github.com/DmitryOdintsov/GinGormAPI/helpers"
	"github.com/DmitryOdintsov/GinGormAPI/models"
	"github.com/gin-gonic/gin"
)

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	err := models.GetAllArticles(&articles)
	if err != nil {
		helpers.RespondJSON(c, 404, articles)
		return
	}
	helpers.RespondJSON(c, 200, articles)
}

func PostNewArticle(c *gin.Context) {
	var article models.Article
	if err0 := c.BindJSON(&article); err0 != nil {
	} ////////////
	err := models.AddNewArticle(&article)
	if err != nil {
		helpers.RespondJSON(c, 404, article)
		return
	}
	helpers.RespondJSON(c, 201, article)
}

func GetArticleById(c *gin.Context) {
	id := c.Params.ByName("id")
	var article models.Article
	err := models.GetArticleById(&article, id)
	if err != nil {
		helpers.RespondJSON(c, 404, article)
		return
	}
	helpers.RespondJSON(c, 200, article)
}

func UpdateArticleById(c *gin.Context) {
	var article models.Article
	id := c.Params.ByName("id")
	err := models.GetArticleById(&article, id)
	if err != nil {
		helpers.RespondJSON(c, 404, article)
	}
	_ = c.BindJSON(&article)
	err = models.UpdateArticleById(&article, id)
	if err != nil {
		helpers.RespondJSON(c, 404, article)
		return
	}
	helpers.RespondJSON(c, 202, article)
}

func DeleteArticleById(c *gin.Context) {
	var article models.Article
	id := c.Params.ByName("id")
	err := models.DeleteArticleById(&article, id)
	if err != nil {
		helpers.RespondJSON(c, 404, article)
		return
	}
	helpers.RespondJSON(c, 202, article)
}
