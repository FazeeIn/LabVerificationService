package model

import (
	"bytes"
	"fmt"
)

type Python struct {
}

func (Python) GetFiles(testRequest TestRequest) ([]bytes.Buffer, error) {
	var (
		files []bytes.Buffer
		err   error
	)
	files = make([]bytes.Buffer, 2)
	// Создание временного Python-файла
	files[0], err = newFile("code/__init__.py", testRequest.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to create test file: %s", err)
	}

	// Создание временного файла с тестами
	files[1], err = newFile("tests.py", testRequest.Tests)
	if err != nil {
		return nil, fmt.Errorf("failed to create test file: %s", err)
	}
	return files, nil
}
