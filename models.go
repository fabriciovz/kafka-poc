package main

//items
type ItemFileUpdated struct {
	UUID       string `json:"uuid"`
	Timestamp  string `json:"timestamp"`
	Issuer     string `json:"issuer"`
	StatusCode string `json:"statusCode"`
	ReplyTo    string `json:"replyTo"`
}
