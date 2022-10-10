package handlers

import (
	"net/http"
	"time"

	"http-server/models"
	"http-server/storage"
	_ "http-server/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func remove(slice []models.Article, s int) []models.Article {
	return append(slice[:s], slice[s+1:]...)
}

// CreateArticle godoc
// @Summary      Create article
// @Description  create a new article
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param	article body models.CreateArticleModel true "article body"
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      201  {object}   models.JSONResponse{data=models.Article}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/article [post]
func CreateArticle(c *gin.Context) {
	var body models.CreateArticleModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		return
	}
	//ToDo - validation should be here

	id := uuid.New()
	err:=storage.AddArticle(id.String(),body)
	if err !=nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	article,err:= storage.GetArticleById(id.String())
	if err!=nil{
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Article | GetByList",
		Data:    article,
	})
}

// GetArticleById godoc
// @Summary      get article by id
// @Description  get an article by id
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param	id path string true "Article ID"
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {object}   models.JSONResponse{data=models.PackedArticleModel}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/article/{id} [get]
func GetArticleById(c *gin.Context) {
	idStr:=c.Param("id")

	article,err:= storage.GetArticleById(idStr)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.JSONErrorResponce{
		Error: err.Error(),
	})
	}
	c.JSON(http.StatusOK, models.JSONResponse{
				Message: "OK",
		 			Data:    article,
})
	// for _, v := range storage.InMemoryArticleData {
	// 	if v.ID == idStr {
	// 		c.JSON(http.StatusOK, models.JSONResponse{
	// 			Message: "Article | GetById",
	// 			Data:    v,
	// 		})
	// 		return
	// 	}
	// }
	
}

// getArticleList godoc
// @Summary      List articles
// @Description  get articles
// @Tags         articles
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.JSONResponse{data=[]models.Article}
// @Router       /v1/article [get]
func GetArticleList(c *gin.Context) {
	articleList, err:= storage.GetArticleList()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
		Error: err.Error(),
	})
}
	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    articleList,
	})
}

// UpdateArticle...
// UpdateArticle godoc
// @Summary      Update article
// @Description  update a new article
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param	article body models.Article true "article body"
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {object}   models.JSONResponse{data=[]models.Article}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/article [put]
func UpdateArticle(c *gin.Context) {
	var article models.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		// gin.H{"error": err.Error()}
		return
	}
	for i, v := range storage.InMemoryArticleData {
		if v.ID == article.ID {
			article.CreatedAt = v.CreatedAt

			t := time.Now()
			article.UpdateAt = &t
			storage.InMemoryArticleData[i] = article
			c.JSON(http.StatusOK, models.JSONResponse{
				Message: "Article | Updaate",
				Data:    storage.InMemoryArticleData,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, models.JSONErrorResponce{
		Error: "Article | Update | Not found",
	})
	// 	 gin.H{
	// 	"message": "Article | Update | Not found",
	// 	"data":    storage.InMemoryArticleData,
	// }
}

// DeleteArticle...
// DeleteArticle godoc
// @Summary      delete article by id
// @Description  delete an article by id
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param	id path string true "Article ID"
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {object}   models.JSONResponse{data=models.Article}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	for i, v := range storage.InMemoryArticleData {
		if v.ID == idStr {
			storage.InMemoryArticleData = remove(storage.InMemoryArticleData, i)
			c.JSON(http.StatusOK, models.JSONResponse{
				Message: "Article | Delete",
				Data:    v,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, models.JSONErrorResponce{
		Error: "Article | Delete | Not found",
	})
	// 	gin.H{
	// 	"message": "Article | Delete | Not found",
	// 	"data":    nil,
	// }

}
