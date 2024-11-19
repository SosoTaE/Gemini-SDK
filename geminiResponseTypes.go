package main

type Response struct {
	Candidates    []Candidate   `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
	ModelVersion  string        `json:"modelVersion"`
}

type Candidate struct {
	Content          ResponseContent  `json:"content"`
	FinishReason     string           `json:"finishReason"`
	CitationMetadata CitationMetadata `json:"citationMetadata"`
	AvgLogprobs      float64          `json:"avgLogprobs"`
}

type ResponseContent struct {
	Parts []ResponsePart `json:"parts"`
	Role  string         `json:"role"`
}

type ResponsePart struct {
	Text string `json:"text"`
}

type CitationMetadata struct {
	CitationSources []CitationSource `json:"citationSources"`
}

type CitationSource struct {
	StartIndex int    `json:"startIndex"`
	EndIndex   int    `json:"endIndex"`
	URI        string `json:"uri"`
}

type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}
