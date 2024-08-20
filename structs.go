package main

type Node struct {
	ID          int     `db:"id"`
	Type        string  `db:"type"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	Latitude    float64 `db:"lat"`
	Longitude   float64 `db:"long"`
}

type Discussion struct {
	ID        int    `json:"id"`
	MachineID int    `json:"name"`
	Created   string `json:"created"`
	User      string `json:"user"`
	Content   string `json:"content"`
}

type Response struct {
	Response string `json:"response"`
	Errno    int    `json:"errno"`
	Error    string `json:"error"`
}
