package storage

import (
	"errors"
	"http-server/models"
	"time"
)

var InMemoryAuthorData []models.Author

func AddAuthor(id string, entity models.CreateAuthorModel) error {
	var author models.Author

	author.Id = id
	author.Firstname=entity.Firstname
	author.Lastname=entity.Lastname
	author.CreatedAt = time.Now()

	InMemoryAuthorData = append(InMemoryAuthorData, author)

	return nil
}

func GetAuthorById(id string) (models.Author, error) {

	var result models.Author
	for _, v := range InMemoryAuthorData {
		if v.Id == id {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

//GetAuthorList...	
func GetAuthorList() (resp []models.Author,err error){
	resp=InMemoryAuthorData
	return resp, err
}
