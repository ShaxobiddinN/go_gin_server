package storage

import (
	"errors"
	"http-server/models"

	// "net/http"
	"time"
)

var InMemoryArticleData []models.Article

func AddArticle(id string, entity models.CreateArticleModel) error {
	var article models.Article

	article.ID = id
	article.Content = entity.Content
	article.AuthorID = entity.AuthorID
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

func GetArticleById(id string) (models.PackedArticleModel, error) {

	var result models.PackedArticleModel
	for _, v := range InMemoryArticleData {
		if v.ID == id {

			author, err := GetAuthorById(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.ID = v.ID
			result.Content = v.Content
			result.Author = author
			result.CreatedAt = v.CreatedAt
			result.UpdateAt = v.UpdateAt
			result.DeleteAt = v.DeleteAt
			return result, nil
		}
	}
	return models.PackedArticleModel{}, errors.New("article not found")
}

func GetArticleList() (resp []models.Article, err error) {
	resp = InMemoryArticleData
	return resp, err
}

func UpdateArticle(entity models.UpdateArticleModel) error {
	for i, v := range InMemoryArticleData {
		if v.ID == entity.ID {
			v.Content = entity.Content

			t := time.Now()
			v.UpdateAt = &t

			InMemoryArticleData[i] = v

			return nil
		}
	}
	return errors.New("article not found")
}

func RemoveArticle(id string) error {
	for i, v := range InMemoryArticleData {
		if v.ID == id && v.DeleteAt == nil {
			if v.DeleteAt != nil {
				return errors.New("already deleted")
			}
			t := time.Now()
			v.DeleteAt = &t
			InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found or already deleted")

}
