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
// @Param        q    query     string  false  "name search by q"  Format(email)
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
		Message: "Author | GetByList",
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
// @Param        q    query     string  false  "name search by q"  Format(email)
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
	authorList, err:= storage.GetAuthorList()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
		Error: err.Error(),
	})
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
// @Param	author body models.Author true "author body"
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {object}   models.JSONResponse{data=[]models.Author}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/author [put]
func UpdateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		// gin.H{"error": err.Error()}
		return
	}
	for i, v := range storage.InMemoryAuthorData {
		if v.Id == author.Id {
			author.CreatedAt = v.CreatedAt

			t := time.Now()
			author.UpdateAt = &t
			storage.InMemoryAuthorData[i] = author
			c.JSON(http.StatusOK, models.JSONResponse{
				Message: "Author | Updaate",
				Data:    storage.InMemoryAuthorData,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, models.JSONErrorResponce{
		Error: "Author | Update | Not found",
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
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {object}   models.JSONResponse{data=models.Author}
// @Failure      400  {object}  models.JSONErrorResponce
// @Router       /v1/author/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	for i, v := range storage.InMemoryAuthorData {
		if v.Id == idStr {
			storage.InMemoryAuthorData = removee(storage.InMemoryAuthorData, i)
			c.JSON(http.StatusOK, models.JSONResponse{
				Message: "Author | Delete",
				Data:    v,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, models.JSONErrorResponce{
		Error: "Author | Delete | Not found",
	})


}