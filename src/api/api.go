package api

import (
	"io"
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"github.com/kjblanchard/baby_names/models"
)

var GENDERS = []string{"boy", "girl", "neutral"}

const API_KEY = "KK4EnrhILpvY9XFMDD3bIA==XYvivUV6hGT1hiWI"


func GetBabyNames(gender int) []models.BabyName {
	params := url.Values{}
	params.Set("gender", GENDERS[gender])
	params.Set("popular_only", "true")
	params.Set("X-Api-Key", API_KEY)
	url := fmt.Sprintf("https://api.api-ninjas.com/v1/babynames?%s", params.Encode())
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var json_data []string
	err = json.Unmarshal(body, &json_data)
	if err != nil {
		panic(err)
	}
	var names []models.BabyName
	for _, name := range json_data {
		names = append(names, models.BabyName{Name: name, Gender: gender})
	}
	return names
}
