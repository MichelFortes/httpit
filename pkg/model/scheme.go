package model

type TestScheme struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Tests    []Test `json:"tests"`
}

type Test struct {
	Description    string              `json:"description"`
	Path           string              `json:"path"`
	Method         string              `json:"method"`
	Headers        map[string][]string `json:"headers"`
	QueryParams    map[string][]string `json:"queryParams"`
	Payload        string              `json:"payload"`
	ExpectedResult ExpectedResult      `json:"expectedResult"`
}

type ExpectedResult struct {
	StatusCode uint16              `json:"statusCode"`
	Headers    map[string][]string `json:"headers"`
	Payload    string              `json:"payload"`
}
