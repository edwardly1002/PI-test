package entity

type Template struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	MimeType string `json:"mimeType"`
	Body     string `json:"body"`
}
