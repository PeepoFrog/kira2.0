package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
)

const (
	NETWORK_NAME          = "testnet-1"
	SEKAID_HOME           = `/data/.sekai`
	KEYRING_BACKEND       = "test"
	DOCKER_IMAGE_NAME     = "ghcr.io/kiracore/docker/kira-base"
	DOCKER_IMAGE_VERSION  = "v0.13.11"
	DOCKER_NETWORK_NAME   = "kira_network"
	SEKAI_VERSION         = "latest" // or v0.3.16
	INTERX_VERSION        = "latest" // or v0.4.33
	SEKAID_CONTAINER_NAME = "sekaid"
	INTERX_CONTAINER_NAME = "interx"
	VOLUME_NAME           = "kira_volume:/data"
	MNEMONIC_FOLDER       = "~/mnemonics"
	RPC_PORT              = 26657
	GRPC_PORT             = 9090
	INTERX_PORT           = 11000
	MONIKER               = "VALIDATOR"
)

var log = logging.Log

func main() {
	log.SetLevel(logrus.InfoLevel)

	dockerManager, err := docker.NewTestDockerManager()
	if err != nil {
		log.Fatalln("Can't create instance of docker manager", err)
	}
	defer dockerManager.Cli.Close()

	ctx := context.Background()

	err = dockerManager.VerifyDockerInstallation(ctx)
	if err != nil {
		log.Fatalf("Docker is not available: %s\n", err)
	}

	sekaidIP, err := dockerManager.GetIPofContainer(ctx, SEKAID_CONTAINER_NAME, DOCKER_NETWORK_NAME)
	if err != nil {
		log.Fatalf("Can't get IP of '%s' container: %s\n", SEKAID_CONTAINER_NAME, err)
	}
	log.Infof("Sekaid IP: %s", sekaidIP)

	interxIP, err := dockerManager.GetIPofContainer(ctx, INTERX_CONTAINER_NAME, DOCKER_NETWORK_NAME)
	if err != nil {
		log.Fatalf("Can't get IP of '%s' container: %s\n", INTERX_CONTAINER_NAME, err)
	}
	log.Infof("Interx IP: %s", interxIP)
}
