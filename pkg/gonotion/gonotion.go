package gonotion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/patrickblackjr/notion-automation/internal/structs"
)

var tagID = "3b6f5ceb-099f-4521-a306-a9a385556f80"

// loadEnvironmentVariables loads the environment variables from the .env file
func loadEnvironmentVariables() (string, error) {
	token := os.Getenv("INTEGRATION_TOKEN")
	if token == "" {
		return "", errors.New("INTEGRATION_TOKEN is not set")
	}
	return token, nil
}

// FilterDatabaseReturnLen returns the length of the results array
func FilterDatabaseReturnLen(databaseID, filterProperty, filterValue string) (int, error) {
	token, _ := loadEnvironmentVariables()
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

	body, err := io.ReadAll(resp.Body)
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

func CreatePageInDB(title, databaseID string) (string, error) {
	token, _ := loadEnvironmentVariables()
	headers := generateHeaders(token)

	page := structs.Page{
		Parent: structs.Parent{
			DatabaseID: databaseID,
		},
		Properties: structs.Properties{
			Name: structs.Name{
				Title: []structs.TextContent{
					{Text: structs.Text{Content: title}},
				},
			},
			Tags: structs.MultiTag{
				MultiSelect: []structs.ID{
					{ID: tagID},
				},
			},
		},
	}

	pageJSON, err := json.Marshal(page)
	if err != nil {
		return "", err
	}

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
