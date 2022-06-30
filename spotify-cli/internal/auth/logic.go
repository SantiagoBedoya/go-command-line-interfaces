package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Login() {
	clientId := viper.GetString("client_id")
	clientSecret := viper.GetString("client_secret")
	url := viper.GetString("spotify_account_url")
	request := resty.New().NewRequest()
	auth := fmt.Sprintf("%s:%s", clientId, clientSecret)
	token := fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte(auth)))
	request.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	request.SetHeader("Authorization", token)
	request.SetFormData(map[string]string{"grant_type": "client_credentials"})
	path := fmt.Sprintf("%s/token", url)
	resp, err := request.Post(path)
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(".spotify-cli.json")
	if err != nil {
		panic(err)
	}
	var config map[string]string
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	var respMap map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &respMap); err != nil {
		log.Fatalf("response error: %v", err)
	}
	config["access_token"] = fmt.Sprint(respMap["access_token"])

	file, _ := json.MarshalIndent(config, "", " ")
	ioutil.WriteFile(".spotify-cli.json", file, 0644)
	fmt.Println("successfully logged!")
}
