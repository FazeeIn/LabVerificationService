package model

type Test struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type TestRequest struct {
	Code  string `json:"code"`
	Tests []Test `json:"tests"`
}

type TestResult struct {
	Passed   bool   `json:"passed"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	Error    string `json:"error"`
}
