package main

import "fmt"

func main() {
	requestBody := GeminiRequestBody{
		Contents: []RequestContent{
			{
				Parts: []RequestPart{
					{
						Text: "what is 2+2",
					},
				},
			},
		},
	}

	apiKey := "AIzaSyD1YG_xOaK8uzDFHm4hl2rK3pGAHaYhGCk"

	gemini := NewGemini(apiKey, GEMINI_API_URL, GEMINI_15_FLASH_LATEST)

	response, err := gemini.Send(requestBody)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Parts[0].Text)
}
