package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"sigs.k8s.io/yaml"

	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
)

// SekaidManager represents a manager for Sekaid container and its associated configurations.
type SekaidManager struct {
	ContainerConfig        *container.Config
	SekaiHostConfig        *container.HostConfig
	SekaidNetworkingConfig *network.NetworkingConfig
	dockerManager          *docker.DockerManager
	config                 *types.Config
}

const timeWaitBetweenBlocks = time.Millisecond * 10700
const validatorAccountName = "validator"

// # Create new config file for sekaidManager
//
//	NetworkName // name of a blockchain name (chandID)
//	SekaidHome // home folder for sekai bin
//	InterxHome // home folder for interx bin
//	KeyringBackend // name of keyring
//	DockerImageName // name of a docker image that will be used to create containers for sekai and interx
//	DockerImageVersion // version of a docker image that will be used to create containers for sekai and interx
//	DockerNetworkName // the name of docker network that will be create and used for sekaid and interx containers
//	SekaiVersion // version of sekai binary
//	InterxVersion // version of interx binary
//	SekaidContainerName // name for sekai container
//	InterxContainerName // name for interx container
//	VolumeName // the name of a docker's volume that will be SekaidContainerName and InterxContainerName will be using
//	MnemonicDir // destination where mnemonics file will be saved
//	RpcPort // sekaid's rpc port
//	GrpcPort // sekaid's grpc port
//	InterxPort // interx endpoint port
//	Moniker // Moniker
func NewConfig(
	NetworkName,
	SekaidHome,
	InterxHome,
	KeyringBackend,
	DockerImageName,
	DockerImageVersion,
	DockerNetworkName,
	SekaiVersion,
	InterxVersion,
	SekaidContainerName,
	InterxContainerName,
	VolumeName,
	MnemonicDir,
	RpcPort,
	GrpcPort,
	InterxPort,
	Moniker string) *types.Config {
	return &types.Config{
		NetworkName:         NetworkName,
		SekaidHome:          SekaidHome,
		InterxHome:          InterxHome,
		KeyringBackend:      KeyringBackend,
		DockerImageName:     DockerImageName,
		DockerImageVersion:  DockerImageVersion,
		DockerNetworkName:   DockerNetworkName,
		SekaiVersion:        SekaiVersion,
		InterxVersion:       InterxVersion,
		SekaidContainerName: SekaidContainerName,
		InterxContainerName: InterxContainerName,
		VolumeName:          VolumeName,
		MnemonicDir:         MnemonicDir,
		RpcPort:             RpcPort,
		GrpcPort:            GrpcPort,
		InterxPort:          InterxPort,
		Moniker:             Moniker}
}

// Returns configured SekaidManager.
//
//	*docker.DockerManager // The pointer for docker.DockerManager instance.
//	*config	//Pointer to config struct, can create new instance by calling NewConfig() function
func NewSekaidManager(dockerManager *docker.DockerManager, config *types.Config) (*SekaidManager, error) {
	log := logging.Log
	log.Infof("Creating sekaid manager with ports: %s, %s, image: '%s', volume: '%s' in '%s' network\n", config.GrpcPort, config.RpcPort, config.DockerImageName, config.VolumeName, config.DockerNetworkName)

	natGrpcPort, err := nat.NewPort("tcp", config.GrpcPort)
	if err != nil {
		log.Errorf("Creating NAT GRPC port error: %s\n", err)
		return nil, err
	}

	natRpcPort, err := nat.NewPort("tcp", config.RpcPort)
	if err != nil {
		log.Errorf("Creating NAT RPC port error: %s\n", err)
		return nil, err
	}

	sekaiContainerConfig := &container.Config{
		Image:       fmt.Sprintf("%s:%s", config.DockerImageName, config.DockerImageVersion),
		Cmd:         []string{"/bin/bash"},
		Tty:         true,
		AttachStdin: true,
		OpenStdin:   true,
		StdinOnce:   true,
		Hostname:    fmt.Sprintf("%s.local", config.SekaidContainerName),
		ExposedPorts: nat.PortSet{
			natGrpcPort: struct{}{},
			natRpcPort:  struct{}{},
		},
	}

	sekaiHostConfig := &container.HostConfig{
		Binds: []string{
			config.VolumeName,
		},
		PortBindings: nat.PortMap{
			natGrpcPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.GrpcPort}},
			natRpcPort:  []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.RpcPort}},
		},
		Privileged: true,
	}

	sekaidNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			config.DockerNetworkName: {},
		},
	}

	return &SekaidManager{sekaiContainerConfig, sekaiHostConfig, sekaidNetworkingConfig, dockerManager, config}, err
}

// InitSekaidBinInContainer sets up the 'sekaid' container with the specified configurations.
// ctx: The context for the operation.
// Moniker: The Moniker for the 'sekaid' container.
// SekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// SekaidHome: The home directory for 'sekaid'.
// KeyringBackend: The keyring backend to use.
// RpcPort: The RPC port for 'sekaid'.
// MnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the init process.
func (s *SekaidManager) InitSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infoln("Setting up 'sekaid' container")

	command := fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`, s.config.NetworkName, s.config.SekaidHome, s.config.Moniker)
	_, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`mkdir %s`, s.config.MnemonicDir)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/sekai.mnemonic`, validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid add-genesis-account %s 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%v --home=%v`, validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid gentx-claim %s --keyring-backend=%s --moniker="%s" --home=%s`, validatorAccountName, s.config.KeyringBackend, s.config.Moniker, s.config.SekaidHome)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	log.Infoln("'sekaid' container started")
	return nil
}

// StartSekaidBinInContainer starts sekaid binary inside SekaidContainerName var
// ctx: The context for the operation.
// Returns an error if any issue occurs during the start process.
func (s *SekaidManager) StartSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infoln("Starting 'sekaid' container")
	command := fmt.Sprintf(`sekaid start --rpc.laddr "tcp://0.0.0.0:%s" --home=%s`, s.config.RpcPort, s.config.SekaidHome)
	_, err := s.dockerManager.ExecCommandInContainerInDetachMode(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
	}

	return nil
}

// Combine SetupSekaidBinInContainer and StartSekaidBinInContainer together.
// First trying to run sekaid bin from previous state if exist.
// Then checking if sekaid bin running inside container.
// If no initing new one.
// Then starting again.
// If no sekaid bin running inside container second time - return error.
// Then starting propagating transactions for permissions as in sekai-env.sh
// ctx: The context for the operation.
// Moniker: The Moniker for the 'sekaid' container.
// SekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// SekaidHome: The home directory for 'sekaid'.
// KeyringBackend: The keyring backend to use.
// RpcPort: The RPC port for 'sekaid'.
// MnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the run process.
func (s *SekaidManager) RunSekaidContainer(ctx context.Context) error {
	log := logging.Log
	err := s.StartSekaidBinInContainer(ctx)
	if err != nil {
		log.Errorf("Cannot start sekaid bin in %s container", s.config.SekaidContainerName)
	}
	time.Sleep(time.Second * 1)
	check, _, err := s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil {
		log.Infof("Error while setup '%s' container: %s\n", s.config.SekaidContainerName, err)
		return err
	}
	if !check {
		log.Infof("Error starting sekaid binary first time in '%s' container, initing new instance\n", s.config.SekaidContainerName)
		err = s.InitSekaidBinInContainer(ctx)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", s.config.SekaidContainerName, err)
			return err
		}
		err := s.StartSekaidBinInContainer(ctx)
		if err != nil {
			log.Errorf("Cannot start sekaid bin in %s container, %s", s.config.SekaidContainerName, err)
		}

		time.Sleep(time.Second * 1)
		log.Printf("%+v\n", s.config)
		check, _, err = s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", s.config.SekaidContainerName, err)
			return err
		}
		if !check {
			log.Errorf("Error starting sekaid bin second time in '%s' container\n", s.config.SekaidContainerName)
			return fmt.Errorf("couldn't start sekaid bin second time")
		}
		err = s.PostGenesisProposals(ctx)
		if err != nil {
			log.Errorf("Error while propagating transaction: %s \n", err)
			return err
		}
		err = s.UpdateIdentityRegistrar(ctx, validatorAccountName)
		if err != nil {
			log.Errorf("Error while updating identity registrar %s \n", err)
			return err
		}

	}
	log.Printf("SEKAID CONTAINER STARTED\n\n\n")
	return nil
}

// Post genesis proposals after launching new network from KM1 await-validator-init.sh file.
// Adding required permitions for validator.
// First getting validator address with GetAddressByName.
// Then in loop calling GivePermissionsToAddress func with delay between calls 10 sec because tx can be propagated once per 10 sec
func (s *SekaidManager) PostGenesisProposals(ctx context.Context) error {
	log := logging.Log
	address, err := s.GetAddressByName(ctx, validatorAccountName)
	if err != nil {
		log.Fatalf("Error while getting address in '%s' container: %s\n", s.config.SekaidContainerName, err)
	}
	permissions := []int{
		types.PermWhitelistAccountPermissionProposal,
		types.PermRemoveWhitelistedAccountPermissionProposal,
		types.PermCreateUpsertTokenAliasProposal,
		types.PermCreateSoftwareUpgradeProposal,
		types.PermVoteWhitelistAccountPermissionProposal,
		types.PermVoteRemoveWhitelistedAccountPermissionProposal,
		types.PermVoteUpsertTokenAliasProposal,
		types.PermVoteSoftwareUpgradeProposal,
	}
	log.Printf("Permissions to add:%v to: %s", permissions, address)
	//waiting 10 sec to first block to be propagated
	time.Sleep(timeWaitBetweenBlocks)
	for _, perm := range permissions {
		log.Printf("\n\n\nAdding permission %v, approximately duration: %v\n", perm, timeWaitBetweenBlocks*2)
		err = s.GivePermissionsToAddress(ctx, perm, address)
		if err != nil {
			log.Errorf("%s\n", err)
		}
		log.Printf("Checking if %s has %v permission\n", address, perm)
		check, err := s.CheckAccountPermission(ctx, perm, address)
		if err != nil {
			log.Errorf("%s\n", err)
		}
		if !check {
			log.Errorf("Could not find  %v with %s\n", perm, address)
		}
		// time.Sleep(timeWaitBetweenBlocks)

	}
	return nil
}

// Getting TX by parsing json output of `sekaid query tx <TXhash>`
func (s *SekaidManager) GetTxQuery(ctx context.Context, transactionHash string) (types.TxData, error) {
	log := logging.Log
	var data types.TxData

	command := fmt.Sprintf(`sekaid query tx %s  --home=%s -output=json`, transactionHash, s.config.SekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("error when checking tx:  %s.  Command: %s. Error:%s\n", transactionHash, command, err)
		return data, err
	}
	err = json.Unmarshal(out, &data)

	if err != nil {
		log.Errorf("error when unmarshaling tx:  %s \ndata to unmarshal:\n%s\n error: %s\n", transactionHash, string(out), err)
		return data, err
	}
	log.Printf("CheckTransactionStatus: \n %+v \n", data)
	return data, nil
}

// Giving permission for chosen address.
// Permissions are ints thats have 0-65 range
//
// Using command: sekaid tx customgov permission whitelist --from "$KM_ACC" --keyring-backend=test --permission="$PERM" --addr="$ADDR" --chain-id=$NETWORK_NAME --home=$SEKAID_HOME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
//
// Then unmarshaling json output and checking sekaid hex of tx
// Then waiting timeWaitBetweenBlocks for tx to propagate in blockchain and checking status code of Tx with GetTxQuery
func (s *SekaidManager) GivePermissionsToAddress(ctx context.Context, permissionToAdd int, address string) error {
	log := logging.Log
	command := fmt.Sprintf(`sekaid tx customgov permission whitelist --from %s --keyring-backend=test --permission=%v --addr=%s --chain-id=%s --home=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, address, permissionToAdd, address, s.config.NetworkName, s.config.SekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("error when giving  %v permission. Command: %s", permissionToAdd, command)
	}
	log.Printf("permission voted to address %s, perm: %v\n", address, permissionToAdd)

	var data types.TxData
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Error unmarshaling:%s \n%s", string(out), err)
		return err
	}
	log.Printf("GivePermissionToAddress: \n %+v \n", data)
	time.Sleep(timeWaitBetweenBlocks)

	txData, err := s.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		log.Errorf("Error Checking transaction status:%s", err)
		return err
	}
	if txData.Code != 0 {
		log.Errorf("ERROR IN PROPAGATING TRANSACTION \nTRANSACTION STATUS:%v\n", txData.Code)
		return fmt.Errorf("error in adding %v permission to %s address.\nTxHash:%s\nCode: %v", permissionToAdd, address, data.Txhash, txData.Code)
	}
	return nil
}

// Checking if account has a specific permission
//
// https://github.com/KiraCore/sekai/blob/master/scripts/sekai-env.sh
//
// sekaid query customgov permissions kira12tptcuw0cp9fccng80vkmqen96npyyrvh2nw5q --output=json --home=/data/.sekai
//
//	permissionToCheck  is a int with 0-65 range
//
// address has to be kira address(not name) : kira12tptcuw0cp9fccng80vkmqen96npyyrvh2nw5q for example, you can get it from local keyring by func GetAddressByName()
func (s *SekaidManager) CheckAccountPermission(ctx context.Context, permissionToCheck int, address string) (bool, error) {
	log := logging.Log
	log.Printf("Looking for %v permission \n", permissionToCheck)
	command := fmt.Sprintf("sekaid query customgov permissions %s --output=json --home=%s", address, s.config.SekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf(`error when executing "%s" command in %s container`, command, s.config.SekaidContainerName)
		return false, err
	}
	var perms types.AddressPermissions
	err = json.Unmarshal(out, &perms)
	if err != nil {
		log.Errorf("Error unmarshaling:%s \n%s", string(out), err)
		return false, err
	}
	log.Printf("CheckAccountPermission: \n %+v \n", perms)
	for _, perm := range perms.WhiteList {
		if permissionToCheck == perm {
			log.Printf("Permission %v was found with %s address. \n", permissionToCheck, address)

			return true, nil
		}
	}
	log.Printf("No %v permission were found with %s address. \n", permissionToCheck, address)
	return false, nil
}

// Getting address from keyring.
//
// sekaid keys show validator --keyring-backend=test --home=test
func (s *SekaidManager) GetAddressByName(ctx context.Context, addressName string) (string, error) {
	log := logging.Log
	command := fmt.Sprintf("sekaid keys show %s --keyring-backend=%s --home=%s", addressName, s.config.KeyringBackend, s.config.SekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Can't get address by %s name. Command: %s. Error: %s", addressName, command, err)
		return "", err
	}
	var key []types.SekaidKey
	log.Println(string(out))
	err = yaml.Unmarshal([]byte(out), &key)
	if err != nil {
		log.Fatalf("error, cannot unmarshal yaml: %v", err)
		return "", err
	}
	if len(key) <= 0 {
		log.Errorf(`no keys were found in keyring with "%s" name`, addressName)
		return "", fmt.Errorf("no keys were found in keyring with %s name", addressName)
	}
	log.Printf("val address: %s\n", key[0].Address)
	return key[0].Address, nil
}

// Updating identity registrar from KM1 await-validator-init.sh file.
func (s *SekaidManager) UpdateIdentityRegistrar(ctx context.Context, accountName string) error {
	log := logging.Log
	nodeStruct, err := s.GetSekaidStatus(s.config.SekaidContainerName, s.config.RpcPort)
	if err != nil {
		log.Errorf("Error when trying to get sekaid status %s\n", err)
		return err
	}

	err = s.UpsertIdentityRecord(ctx, accountName, "description", "This is genesis validator account of the KIRA Team")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "social", "https://tg.kira.network,twitter.kira.network")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "contact", "https://support.kira.network")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "website", "https://kira.network")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "username", "KIRA")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "logo", "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "avatar", "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "pentest1", "<iframe src=javascript:alert(1)>")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "pentest2", "<img/src=x a='' onerror=alert(2)>")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "pentest3", "<img src=1 onerror=alert(3)>")
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}
	err = s.UpsertIdentityRecord(ctx, accountName, "validator_node_id", nodeStruct.Result.NodeInfo.ID)
	if err != nil {
		log.Errorf("Error when upserting identity record: %s\n", err)
		return err
	}

	//not sure if this needed
	// err = s.UpsertIdentityRecord(ctx, "signer", "username", "faucet", SekaidContainerName, SekaidHome, KeyringBackend, NetworkName)
	// if err != nil {
	// 	log.Errorf("Error when upserting identity record: %s\n", err)
	// 	return err
	// }
	// s.UpsertIdentityRecord(ctx, "test", "username", "test", SekaidContainerName, SekaidHome, KeyringBackend, NetworkName).
	// if err != nil {
	// 	log.Errorf("Error when upserting identity record: %s\n", err)
	// 	return err
	// }
	return nil
}

// upsertIdentityRecord  from sekai-utils.sh
func (s *SekaidManager) UpsertIdentityRecord(ctx context.Context, account, key, value string) error {
	log := logging.Log
	address, err := s.GetAddressByName(ctx, account)
	if err != nil {
		log.Errorf("Error while getting kira address from keyring %s\n", err)
		return err
	}
	var out []byte
	if value != "" {
		command := fmt.Sprintf(`sekaid tx customgov register-identity-records --infos-json="{\"%s\":\"%s\"}" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, value, address, s.config.KeyringBackend, s.config.SekaidHome, s.config.NetworkName)
		out, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
		if err != nil {
			log.Errorf("Error while executing command %s\n in %s container", command, s.config.SekaidContainerName)
			return err
		}
		log.Printf("%s\n", string(out))

	} else {
		command := fmt.Sprintf(`sekaid tx customgov delete-identity-records --keys="%s" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, address, s.config.KeyringBackend, s.config.SekaidHome, s.config.NetworkName)
		out, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
		if err != nil {
			log.Errorf("Error while executing command %s\n in %s container", command, s.config.SekaidContainerName)
			return err
		}
		log.Printf("%s\n", string(out))
	}
	var data types.TxData
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Error unmarshaling:%s \n%s", string(out), err)
		return err
	}
	time.Sleep(timeWaitBetweenBlocks)

	txData, err := s.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		log.Errorf("Error while gettind tx data from %s hash\n", data.Txhash)
		return err
	}
	if txData.Code != 0 {
		log.Errorf("the %s transaction was executed with error. Code %v", data.Txhash, txData.Code)
		return fmt.Errorf("the %s transaction was executed with error. Code %v", data.Txhash, txData.Code)
	}
	return nil
}

// func to get status of sekaid node
// same as curl localhost:26657/status (port for sekaid's rpc endpoint)
func (s *SekaidManager) GetSekaidStatus(SekaidContainerName, RpcPort string) (*types.Status, error) {
	log := logging.Log

	url := fmt.Sprintf("http://localhost:%s/status", RpcPort)
	log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Println("Failed to send GET request:", err)
		return &types.Status{}, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to read response body:", err)
		return &types.Status{}, err
	}

	var statusData *types.Status
	if err := json.Unmarshal(body, &statusData); err != nil {
		log.Println("Failed to parse JSON:", err)
		return &types.Status{}, err
	}
	return statusData, nil
}
