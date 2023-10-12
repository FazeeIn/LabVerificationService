package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"

	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func NewContainer(testRequest model.TestRequest) ([]byte, error) {
	// Создание временного Python-файла
	bufCode, err := newCodeFile(testRequest.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to create test file: %s", err)
	}

	// Создание временного файла с тестами
	bufTests, err := newTestsFile(testRequest.Tests)
	if err != nil {
		return nil, fmt.Errorf("failed to create test file: %s", err)
	}

	// Создание и запуск контейнера
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Docker daemon: %s", err)
	}

	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: "python:latest",
		Cmd:   []string{"python", "test.py", "tests.json"},
	}, nil, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %s", err)
	}

	err = cli.CopyToContainer(context.Background(), resp.ID, "/", &bufCode, types.CopyToContainerOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to copy Python file to container: %s", err)
	}

	err = cli.CopyToContainer(context.Background(), resp.ID, "/", &bufTests, types.CopyToContainerOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to copy test file to container: %s", err)
	}

	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %s", err)
	}

	cmd := exec.Command("docker", "wait", resp.ID)
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to wait for container: %s", err)
	}

	out, err := cli.ContainerLogs(context.Background(), resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return nil, fmt.Errorf("failed to get container logs: %s", err)
	}

	resultData, err := io.ReadAll(out)
	if err != nil {
		return nil, fmt.Errorf("failed to read container logs: %s", err)
	}

	return resultData, nil
}

func newCodeFile(code string) (bytes.Buffer, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	err := tw.WriteHeader(&tar.Header{
		Name: "test.py",                // filename
		Mode: 0644,                     // permissions
		Size: int64(len([]byte(code))), // filesize
	})
	if err != nil {
		return buf, err
	}
	tw.Write([]byte(code))
	tw.Close()
	return buf, nil
}

func newTestsFile(tests []model.Test) (bytes.Buffer, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	testData, err := json.Marshal(tests)
	if err != nil {
		return buf, err
	}
	tw.WriteHeader(&tar.Header{
		Name: "tests.json",                 // filename
		Mode: 0644,                         // permissions
		Size: int64(len([]byte(testData))), // filesize
	})
	if err != nil {
		return buf, err
	}
	tw.Write([]byte(testData))
	tw.Close()
	return buf, nil
}
