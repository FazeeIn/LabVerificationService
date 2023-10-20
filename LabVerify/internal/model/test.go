package model

type Test struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type TestRequest struct {
	Code  []byte `json:"code"`
	Tests []byte `json:"tests"`
}

type TestResult struct {
	Passed   bool   `json:"passed"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	Error    string `json:"error"`
}
