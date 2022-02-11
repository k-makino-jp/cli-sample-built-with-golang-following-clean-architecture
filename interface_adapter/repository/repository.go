package repository

import (
	"encoding/json"

	"github.com/k-makino-jp/cli-built-with-golang-following-clean-architecture/entity"
	"github.com/k-makino-jp/cli-built-with-golang-following-clean-architecture/infrastructure/http"
)

type HTTPClient interface {
	Get(queries map[string]string, headers map[string]string, url string) (statusCode int, respBody []byte, err error)
	Post(queries map[string]string, headers map[string]string, body string, url string) (statusCode int, respBody []byte, err error)
}

type userRepository struct {
	httpClient HTTPClient
}

func NewUserRepository() *userRepository {
	return &userRepository{
		httpClient: http.NewClient(),
	}
}

func (u userRepository) Get() ([]entity.User, error) {
	queries := map[string]string{}
	headers := map[string]string{}
	url := "http://localhost:8080/api/users"
	_, respBody, err := u.httpClient.Get(queries, headers, url)
	if err != nil {
		return nil, err
	}

	var users []entity.User
	err = json.Unmarshal(respBody, &users)
	return users, err
}
