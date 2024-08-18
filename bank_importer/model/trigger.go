package model

type Type string

const (
	Debit  Type = "statement"
	Credit Type = "invoice"
)

type Trigger struct {
	BasePath string `json:"basePath"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
	Day      int    `json:"day"`
	Bank     string `json:"bank"`
	Account  string `json:"account"`
	Type     Type   `json:"type"`
}
