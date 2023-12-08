package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	databaseID := "e834563bdcbf4dbb9643a54e3111b03d"
	currentDate := generateCurrentDate()

	length, err := filterDatabaseReturnLen(databaseID, "Name", currentDate)
	if err != nil {
		panic(err)
	}
	if length != 0 {
		panic("Page already exists")
	}

	_, err = createPageInDB(currentDate, databaseID)
	if err != nil {
		panic(err)
	}
}

func generateCurrentDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func loadEnvironmentVariables() string {
	token := os.Getenv("INTEGRATION_TOKEN")
	return token
}

func filterDatabaseReturnLen(databaseID, filterProperty, filterValue string) (int, error) {
	token := loadEnvironmentVariables()
	headers := generateHeaders(token)

	filter := map[string]interface{}{
		"filter": map[string]interface{}{
			"property": filterProperty,
			"title": map[string]string{
				"equals": filterValue,
			},
		},
	}

	filterJSON, _ := json.Marshal(filter)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", databaseID), bytes.NewBuffer(filterJSON))
	if err != nil {
		return 0, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return len(result["results"].([]interface{})), nil
}

func generateHeaders(token string) map[string]string {
	return map[string]string{
		"Authorization":  fmt.Sprintf("Bearer %s", token),
		"Notion-Version": "2022-06-28",
		"Content-Type":   "application/json",
	}
}

func createPageInDB(title, databaseID string) (string, error) {
	token := loadEnvironmentVariables()
	headers := generateHeaders(token)

	page := map[string]interface{}{
		"parent": map[string]string{
			"database_id": databaseID,
		},
		"properties": map[string]interface{}{
			"Name": map[string]interface{}{
				"title": []map[string]interface{}{
					{"text": map[string]string{"content": title}},
				},
			},
			"Tags": map[string]interface{}{
				"multi_select": []map[string]string{
					{"id": "3b6f5ceb-099f-4521-a306-a9a385556f80"},
				},
			},
		},
	}

	pageJSON, _ := json.Marshal(page)

	req, err := http.NewRequest("POST", "https://api.notion.com/v1/pages", bytes.NewBuffer(pageJSON))
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
