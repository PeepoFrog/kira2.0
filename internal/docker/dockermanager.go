package docker

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"encoding/json"

	"io"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type DockerManager struct {
	Cli *client.Client
}

// NewDockerManager creates a new DockerManager instance based on the provided configuration.
// r: An io.Reader containing the configuration data in JSON format.
// Returns a new DockerManager instance and an error if there is an issue with decoding the configuration or creating the Docker client.
func NewDockerManager(r io.Reader) (*DockerManager, error) {

	log.Println("Creating new DockerManager from config")

	decoder := json.NewDecoder(r)
	config := DockerConfig{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Printf("Failed to decode config: %s", err)
		return nil, fmt.Errorf("Failed to decode config: %w", err)
	}

	cli, err := client.NewClientWithOpts(client.WithHost(config.Host), client.WithVersion(config.APIVersion), client.WithTLSClientConfig(config.CacertPath, config.CertPath, config.KeyPath))
	if err != nil {
		log.Printf("Failed to create Docker client: %s", err)
		return nil, fmt.Errorf("Failed to create Docker client: %w", err)
	}

	log.Println("Successfully created DockerManager")
	return &DockerManager{Cli: cli}, nil
}

// VerifyDockerInstallation verifies if Docker is installed and running by pinging the Docker daemon.
// ctx: The context.Context to use for the ping operation.
// Returns an error if the Docker daemon is not reachable or if there is an error in the ping operation.
func (dm *DockerManager) VerifyDockerInstallation(ctx context.Context) error {
	// Try to ping the Docker daemon to check if it's running
	_, err := dm.Cli.Ping(ctx)
	if err != nil {
		log.Println("Error pinging Docker daemon:", err)
		return fmt.Errorf("Failed to ping Docker daemon: %w", err)
	}

	// If we got here, Docker is installed and running
	log.Println("Docker is installed and running!")
	return nil
}

// PullImage pulls the specified Docker image using the Docker client associated with the DockerManager.
// It streams the image pull output to a buffer and logs the pretified output.
// ctx: The context.Context to use for the image pull operation.
// imageName: The name of the Docker image to pull.
// Returns an error if the image pull fails or if there is an error in copying the image pull output.
func (dm *DockerManager) PullImage(ctx context.Context, imageName string) error {
	options := types.ImagePullOptions{}
	reader, err := dm.Cli.ImagePull(ctx, imageName, options)
	if err != nil {
		return fmt.Errorf("failed to pull image: %w", err)
	}
	defer reader.Close()

	// Create a buffer for the reader
	buf := new(bytes.Buffer)

	// Copy the image pull output to the buffer
	_, err = io.Copy(buf, reader)
	if err != nil {
		return fmt.Errorf("failed to copy image pull output: %w", err)
	}

	// Print the pretified output from the buffer
	log.Printf("Image pull output: %s", buf.String())

	return nil
}

// RunContainer runs a Docker container with the specified image.
// ctx: The context.Context to use for the container operations.
// image: The name of the Docker image to run.
// Returns an error if there is an issue with creating or starting the container, waiting for it to finish, or retrieving the container logs.
func (dm *DockerManager) RunContainer(ctx context.Context, image string) error {
	config := &container.Config{
		Image: image,
		Cmd:   []string{"echo", "hello world"},
		Tty:   false,
	}

	resp, err := dm.Cli.ContainerCreate(ctx, config, nil, nil, nil, "")
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}

	if err := dm.Cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	log.Printf("Container %s started", resp.ID)

	statusCh, errCh := dm.Cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("container wait error: %w", err)
		}
	case <-statusCh:
	}

	log.Printf("Container %s finished", resp.ID)

	out, err := dm.Cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return fmt.Errorf("failed to retrieve container logs: %w", err)
	}
	defer out.Close()

	log.Printf("Container %s logs:", resp.ID)
	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	if err != nil {
		return fmt.Errorf("failed to copy container logs: %w", err)
	}

	return nil
}
