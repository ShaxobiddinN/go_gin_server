package handlers

import (
	"net/http"
	"strconv"

	"http-server/models"
	"http-server/storage"
	_ "http-server/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func removee(slice []models.Author, s int) []models.Author {
	return append(slice[:s], slice[s+1:]...)
}

// CreateAuthor godoc
// @Summary      Create author
// @Description  create a new author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param	author body models.CreateAuthorModel true "author body"
// @Success      201  {object}   models.JSONResponse{data=models.Author}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/author [post]
func CreateAuthor(c *gin.Context) {
	var body models.CreateAuthorModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		return
	}
	//ToDo - validation should be here

	id := uuid.New()
	err:=storage.AddAuthor(id.String(),body)
	if err !=nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	author,err:= storage.GetAuthorById(id.String())
	if err!=nil{
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Author | Created",
		Data:    author,
	})
}

// GetAuthorById godoc
// @Summary      get author by id
// @Description  get an author by id
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param	id path string true "Author ID"
// @Success      200  {object}   models.JSONResponse{data=models.Author}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/author/{id} [get]
func GetAuthorById(c *gin.Context) {
	idStr:=c.Param("id")

	author,err:= storage.GetAuthorById(idStr)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.JSONErrorResponce{
		Error: err.Error(),
	})
	}
	c.JSON(http.StatusOK, models.JSONResponse{
				Message: "OK",
		 			Data:    author,
})
	// for _, v := range storage.InMemoryAuthorData {
	// 	if v.ID == idStr {
	// 		c.JSON(http.StatusOK, models.JSONResponse{
	// 			Message: "Author | GetById",
	// 			Data:    v,
	// 		})
	// 		return
	// 	}
	// }
	
}

// getAuthorList godoc
// @Summary      List authors
// @Description  get authors
// @Tags         authors
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.JSONResponse{data=[]models.Author}
// @Router       /v1/author [get]
func GetAuthorList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	authorList, err:= storage.GetAuthorList(offset,limit,searchStr)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
		Error: err.Error(),
	})
	return
}
	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    authorList,
	})
}

// UpdateAuthor...
// UpdateAuthor godoc
// @Summary      Update author
// @Description  update a new author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param	author body models.UpdateAuthorModel true "author body"
// @Success      200  {object}   models.JSONResponse{data=[]models.Author}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/author [put]
func UpdateAuthor(c *gin.Context) {
	var body models.UpdateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		return
	}
	err := storage.UpdateAuthor(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	author, err := storage.GetArticleById(body.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "ok",
		Data: author,
	})
}

// DeleteAuthor...
// DeleteAuthor godoc
// @Summary      delete author by id
// @Description  delete an author by id
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param	id path string true "Author ID"
// @Success      200  {object}   models.JSONResponse{data=models.Author}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/author/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	author, err := storage.GetAuthorById(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	err = storage.RemoveAuthor(author.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    author,
	})

}