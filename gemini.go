package main

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (g *Gemini) Send(content GeminiRequestBody) (Response, error) {
	jsonBody, err := json.Marshal(content)
	if err != nil {
		fmt.Println(1)
		return Response{}, err
	}

	url := fmt.Sprintf("%s/%s?key=%s", g.url, g.model, g.apiKey)
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return Response{}, err
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
		return Response{}, err
	}

	if resp.Status != "200 OK" {
		return Response{}, errors.New(string(body))
	}

	var responseContent Response
	err = json.Unmarshal(body, &responseContent)
	if err != nil {
		return Response{}, err
	}

	return responseContent, nil
}
