package types

type Test struct {
	Text string
}

type SekaidKey struct {
	Address string `yaml:"address"`
}

type AddressPermissions struct {
	BlackList []int `json:"blacklist"`
	WhiteList []int `json:"whitelist"`
}

type Config struct {
	NetworkName         string
	SekaidHome          string
	InterxHome          string
	KeyringBackend      string
	DockerImageName     string
	DockerImageVersion  string
	DockerNetworkName   string
	SekaiVersion        string
	InterxVersion       string
	SekaidContainerName string
	InterxContainerName string
	VolumeName          string
	MnemonicDir         string
	RpcPort             string
	GrpcPort            string
	InterxPort          string
	Moniker             string
}
