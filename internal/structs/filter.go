package structs

// Filter is the struct for the JSON body of the request
// to filter the results of a database query in Notion.
// It is used in the function FilterDatabaseReturnLen.
type Filter struct {
	Property string `json:"property"`
	Title    Title  `json:"title"`
}

type Title struct {
	Equals string `json:"equals"`
}
