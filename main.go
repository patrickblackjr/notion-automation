package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/patrickblackjr/notion-automation/pkg/gonotion"
)

var databaseID = "e834563bdcbf4dbb9643a54e3111b03d"
var currentDate = generateCurrentDate()

func generateCurrentDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	length, err := gonotion.FilterDatabaseReturnLen(databaseID, "Name", currentDate)
	if err != nil {
		panic(err)
	}
	if length != 0 {
		panic("Page already exists")
	}

	_, err = gonotion.CreatePageInDB(currentDate, databaseID)
	if err != nil {
		panic(err)
	}
}
