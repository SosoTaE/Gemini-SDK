package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gemini struct {
	apiKey string
	url    string
	model  string
}

func NewGemini(apiKey, url, model string) *Gemini {
	return &Gemini{apiKey, url, model}
}

func (g *Gemini) Send(content GeminiRequestBody) (ResponseContent, error) {
	jsonBody, err := json.Marshal(content)
	if err != nil {
		return ResponseContent{}, err
	}

	url := fmt.Sprintf("%s/%s?key=%s", g.url, g.model, g.apiKey)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return ResponseContent{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseContent{}, err
	}

	var responseContent ResponseContent
	err = json.Unmarshal(body, &responseContent)
	if err != nil {
		return ResponseContent{}, err
	}

	return responseContent, nil
}
