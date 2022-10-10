package main

import (
	"net/http"

	"http-server/docs"
	_ "http-server/docs" // docs is generated by Swag CLI, you have to import it.
	"http-server/handlers"
	"http-server/models"
	"http-server/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @contact.name   API Article
// @contact.url    http://www.johndoe.com
// @contact.email  johndoe@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	docs.SwaggerInfo.Title="Swagger Example info"
	docs.SwaggerInfo.Description="This is a simple server Petstore server"
	docs.SwaggerInfo.Version="2.0"

	err:=storage.AddAuthor("87e5f2ce-d80d-40e6-a77e-f98d58d17481",models.CreateAuthorModel{
		Firstname: "John",
		Lastname: "Doe",
	} )

	err=storage.AddArticle("51800dbe-c098-41de-95ff-dfafa7da9b46", models.CreateArticleModel{
		Content: models.Content{
			Title: "Lorem",
			Body: "Impsume smth smth smtnh",
		},
		AuthorID:"87e5f2ce-d80d-40e6-a77e-f98d58d17481" ,
		// CreatedAt: time.Now(),
	})
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/article", handlers.CreateArticle)
		v1.GET("/article/:id", handlers.GetArticleById)
		v1.GET("/article", handlers.GetArticleList)	
		v1.PUT("/article", handlers.UpdateArticle)
		v1.DELETE("/article/:id", handlers.DeleteArticle)

		v1.POST("/author", handlers.CreateAuthor)
		v1.GET("/author/:id", handlers.GetAuthorById)
		v1.GET("/author", handlers.GetAuthorList)	
		v1.PUT("/author", handlers.UpdateAuthor)
		v1.DELETE("/author/:id", handlers.DeleteAuthor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
