package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/FazeeIn/LabVerificationService/LabVerify/internal/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Lang interface {
	GetFiles(testRequest model.TestRequest) ([]bytes.Buffer, error)
}

func NewContainer(testRequest model.TestRequest, lang Lang) ([]byte, error) {

	// Создание необходимого набора файлов (код, тесты, ...)
	files, err := lang.GetFiles(testRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create files: %s", err)
	}
	// Создание и запуск контейнера
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Docker daemon: %s", err)
	}

	config := &container.Config{
		Image: "python:latest",
		Cmd:   []string{"python", "-m", "unittest", "-v", "tests"},
	}

	resp, err := cli.ContainerCreate(context.Background(), config, nil, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %s", err)
	}

	for _, v := range files {
		err = cli.CopyToContainer(context.Background(), resp.ID, "/", &v, types.CopyToContainerOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to copy Python files to container: %s", err)
		}
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
