package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
)

// InterxManager represents a manager for Interx container and its associated configurations.
type InterxManager struct {
	ContainerConfig        *container.Config
	SekaiHostConfig        *container.HostConfig
	SekaidNetworkingConfig *network.NetworkingConfig
	DockerClient           *docker.DockerManager
}

func NewInterxManager(dockerClient *docker.DockerManager, interxPort, imageName, volumeName, dockerNetworkName, containerName string) (*InterxManager, error) {
	log := logging.Log
	log.Infof("Creating interx manager with port: %s, image: '%s', volume: '%s' in '%s' network\n", interxPort, imageName, volumeName, dockerNetworkName)

	natInterxPort, err := nat.NewPort("tcp", interxPort)
	if err != nil {
		log.Errorf("Creating NAT interx port error: %s\n", err)
		return nil, err
	}

	interxContainerConfig := &container.Config{
		Image:        imageName,
		Cmd:          []string{"/bin/bash"},
		Tty:          true,
		AttachStdin:  true,
		OpenStdin:    true,
		StdinOnce:    true,
		Hostname:     fmt.Sprintf("%s.local", containerName),
		ExposedPorts: nat.PortSet{natInterxPort: struct{}{}},
	}
	interxNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			dockerNetworkName: {},
		},
	}
	interxHostConfig := &container.HostConfig{
		Binds: []string{
			volumeName,
		},
		PortBindings: nat.PortMap{
			natInterxPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: interxPort}},
		},
		Privileged: true,
	}

	return &InterxManager{interxContainerConfig, interxHostConfig, interxNetworkingConfig, dockerClient}, err
}

// SetupInterxContainer sets up the 'interx' container with the specified configurations.
// ctx: The context for the operation.
// sekaidContainerName: The name of the 'sekaid' container.
// interxContainerName: The name of the 'interx' container.
// rpc_port: The RPC port for 'interx' to connect to 'sekaid'.
// grpc_port: The gRPC port for 'interx' to connect to 'sekaid'.
// Returns an error if any issue occurs during the setup process.
func (i *InterxManager) InitInterxBinInContainer(ctx context.Context, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome string) error {
	log := logging.Log
	log.Infoln("Setting up 'interx' container")

	command := fmt.Sprintf(`interx init --rpc="http://%s:%s" --grpc="dns:///%s:%s" -home=%s`, sekaidContainerName, rpc_port, sekaidContainerName, grpc_port, interxHome)
	_, err := i.DockerClient.ExecCommandInContainer(ctx, interxContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}
	// command = fmt.Sprintf(`interx start -home=%s`, interxHome)
	// _, err = i.DockerClient.ExecCommandInContainerInDetachMode(ctx, interxContainerName, []string{`bash`, `-c`, command})
	// if err != nil {
	// 	log.Errorf("Command '%s' execution error: %s\n", command, err)
	// 	return err
	// }

	log.Infoln("interx inited")
	return err
}

func (i *InterxManager) StartInterxBinInContainer(ctx context.Context, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome string) error {
	log := logging.Log

	command := fmt.Sprintf(`interx start -home=%s`, interxHome)
	_, err := i.DockerClient.ExecCommandInContainerInDetachMode(ctx, interxContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}
	log.Infoln("interx started")

	return nil
}

func (i *InterxManager) RunInterxContainer(ctx context.Context, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome string) error {
	log := logging.Log
	err := i.StartInterxBinInContainer(ctx, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome)
	if err != nil {
		log.Errorf("Error while running interex bin in '%s' container: %s\n", sekaidContainerName, err)
	}
	time.Sleep(time.Second * 1)
	check, _, err := i.DockerClient.CheckIfProcessIsRunningInContainer(ctx, "interx", interxContainerName)
	if err != nil {
		log.Errorf("Error while setup '%s' container: %s\n", sekaidContainerName, err)
		return err
	}
	if !check {
		log.Infof("Error starting interx binary first time in '%s' container, initing new instance\n", sekaidContainerName)
		err = i.InitInterxBinInContainer(ctx, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", sekaidContainerName, err)
			return err
		}
		err := i.StartInterxBinInContainer(ctx, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome)
		if err != nil {
			log.Errorf("Error while running interex bin in '%s' container: %s\n", sekaidContainerName, err)
		}
		time.Sleep(time.Second * 1)
		check, _, err := i.DockerClient.CheckIfProcessIsRunningInContainer(ctx, "interx", interxContainerName)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", sekaidContainerName, err)
			return err
		}
		if !check {
			log.Errorf("Error starting interex binary second time in '%s' container\n", sekaidContainerName)
			return fmt.Errorf("cannot start sekaid bin in %s container", interxContainerName)
		}
	}
	return nil
}
