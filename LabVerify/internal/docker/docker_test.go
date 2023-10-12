package docker_test

import (
	"testing"

	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/docker"
	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/model"
)

func TestNewContainer(t *testing.T) {
	testRequest := model.TestRequest{Code: "print('Hello, World!')",
		Tests: []model.Test{{Input: "", Output: "Hello, World!\n"},
			{Input: "test", Output: "Hello, World!\n"}}}

	_, err := docker.NewContainer(testRequest)

	if err != nil {
		t.Errorf("NewContainer('Hello, World!') = %s; want nil", err)
	}
}
