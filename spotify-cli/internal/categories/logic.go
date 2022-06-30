package categories

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Get() {
	accessToken := viper.GetString("access_token")
	baseUrl := viper.GetString("spotify_api_url")
	request := resty.New().NewRequest()
	request.SetHeader("authorization", fmt.Sprintf("Bearer %s", accessToken))
	url := fmt.Sprintf("%s/v1/browse/categories", baseUrl)
	response, err := request.Get(url)
	if err != nil {
		log.Fatalf("error in categories.get: %v", err)
	}
	var categoriesResponse CategoriesResponse
	if err := json.Unmarshal(response.Body(), &categoriesResponse); err != nil {
		log.Fatalf("error in categories.get: %v", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Categories")
	t.AppendHeader(table.Row{"Id", "Name", "Href"})
	rows := []table.Row{}
	for _, category := range categoriesResponse.Categories.Items {
		rows = append(rows, table.Row{category.ID, category.Name, category.Href})
	}

	t.AppendRows(rows)
	t.Render()
}
