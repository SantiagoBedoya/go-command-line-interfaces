package albums

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Get() {
	accessToken := viper.GetString("access_token")
	baseURL := viper.GetString("spotify_api_url")
	request := resty.New().NewRequest()
	url := fmt.Sprintf("%s/v1/albums", baseURL)
	request.SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := request.Get(url)
	if err != nil {
		log.Fatalf("error in albums.get: %v", err)
	}
	fmt.Println(string(resp.Body()))
}
