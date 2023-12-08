package structs

// Page is the struct for the JSON body of the request
// to create a new page in Notion. It is used in the function
// CreatePageInDB. The elements are separated into structs
// for readability and reuse in other functions if necessary.
type Page struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	Name Name     `json:"Name"`
	Tags MultiTag `json:"Tags"`
}

type Name struct {
	Title []TextContent `json:"title"`
}

type TextContent struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type MultiTag struct {
	MultiSelect []ID `json:"multi_select"`
}

type ID struct {
	ID string `json:"id"`
}
