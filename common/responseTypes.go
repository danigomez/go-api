package common

type RequestResult struct {
	Data       interface{} `json:"data"`
	StatusCode string      `json:"statusCode"`
}
