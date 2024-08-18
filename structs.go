package main

/*
	Item

TODO: Rename to ItemDB
*/
type Item struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	URL       string `db:"url"`
	ImageURL  string `db:"image_url"`
	Person    string `db:"person"`
	Quantity  int    `db:"quantity"`
	Created   string `db:"created"`
	Deleted   bool   `db:"deleted"`
	Completed bool   `db:"completed"`
}

/* ItemJSON */
type ItemJSON struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	ImageURL  string `json:"image_url"`
	Person    string `json:"person"`
	Quantity  int    `json:"quantity"`
	Created   string `json:"created"`
	Deleted   bool   `json:"deleted"`
	Completed bool   `json:"completed"`
}

/*
	Response

TODO: Split Response struct into Response and ErrorResponse
*/
type Response struct {
	Response string `json:"response"`
	Errno    int    `json:"errno"`
	Error    string `json:"error"`
}
