package profile

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

func (s *service) Me() {
	accessToken := viper.GetString("access_token")
	baseUrl := viper.GetString("spotify_api_url")
	fmt.Println(accessToken, baseUrl)
	request := resty.New().NewRequest()
	request.SetHeader("authorization", fmt.Sprintf("Bearer %s", accessToken))
	url := fmt.Sprintf("%s/v1/me", baseUrl)
	response, err := request.Get(url)
	if err != nil {
		log.Fatalf("error in profile.me: %v", err)
	}

	fmt.Println(string(response.Body()))
}
