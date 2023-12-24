package helper

import (
	"encoding/json"
	"io"
	"net/http"

	domain "github.com/salamanderman234/project-intro-2023/layanan-kelas/domains"
)

func CallService(url string) (map[string]any, error) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(req)
	if res.StatusCode != 200 {
		return nil, domain.ErrApiResourceNotFound
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, domain.ErrConversionType
	}
	var result map[string]any
	json.Unmarshal(body, &result)
	data, ok := result["data"]
	if !ok {
		return nil, domain.ErrConversionType
	}
	dataMap, ok := data.(map[string]any)
	if !ok {
		return nil, domain.ErrConversionType
	}
	return dataMap, nil
}
