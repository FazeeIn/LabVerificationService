package docker_test

import (
	"os"
	"testing"

	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/docker"
	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/model"
)

func TestNewContainer(t *testing.T) {
	testHelloWorld := model.TestRequest{Code: []byte(""),
		Tests: []byte("print('Hello, World!')")}

	sumBuf, _ := os.ReadFile("testdata/code/__init__.py")
	sumTestsBuf, _ := os.ReadFile("testdata/test.py")
	testSum := model.TestRequest{
		Code:  sumBuf,
		Tests: sumTestsBuf}

	testRequests := []struct {
		name string
		test model.TestRequest
	}{
		{name: "HelloWorld", test: testHelloWorld},
		{name: "Sum", test: testSum},
	}

	for _, testRequest := range testRequests {
		_, err := docker.NewContainer(testRequest.test, model.Python{})

		if err != nil {
			t.Errorf("NewContainer(%s) = %s; want nil", testRequest.name, err)
		}
	}
}
