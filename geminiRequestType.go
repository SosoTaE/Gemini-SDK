package main

type RequestPart struct {
	Text string `json:"text"`
}
type RequestContent struct {
	Parts []RequestPart `json:"parts"`
}
type GeminiRequestBody struct {
	Contents []RequestContent `json:"contents"`
}
