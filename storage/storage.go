package storage

import (
	"errors"
	"http-server/models"
	"time"
)


var IntMemoryArticleData []models.Article

func AddArticle(id string, entity models.CreateArticleModel) error {
	var article models.Article

	article.ID = id
	article.Content=entity.Content
	article.AuthorID=entity.AuthorID
	article.CreatedAt = time.Now()

	IntMemoryArticleData = append(IntMemoryArticleData, article)

	return nil
}


func GetArticleById(id string) (models.PackedArticleModel, error) {

	var result models.PackedArticleModel
	for _, v := range IntMemoryArticleData {
		if v.ID == id {

			author,err:= GetAuthorById(v.AuthorID)
			if err!=nil{
				return result,err 
			}

			result.Content=v.Content
			result.Author=author
			result.CreatedAt=v.CreatedAt
			result.UpdateAt=v.UpdateAt
			result.DeleteAt=v.DeleteAt
			return result, nil
		}
	}
	return models.PackedArticleModel{}, errors.New("article not found")
}


func GetArticleList() (resp []models.Article,err error){
	resp=IntMemoryArticleData
	return resp, err
}