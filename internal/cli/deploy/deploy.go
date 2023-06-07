package deploy

import (
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

const (
	use   = "deploy [ip address]"
	short = "Short description of deploy command"
	long  = "Long description of deploy command"
)

var (
	log   = logging.Log
	nodes = []string{"interx", "sekai"}
)

func Node() *cobra.Command {
	log.Debugln("Adding `deploy` command...")
	nodeCmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Args:    cobra.ExactArgs(1),
		Example: "deploy 127.0.0.1 --priv-key=path/to/priv-key --pub-key=path/to/pub-key --interx=v0.3.16 --sekai=v0.3.46",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	for _, node := range nodes {
		nodeCmd.PersistentFlags().String(node, "", "Provide version to deploy")
	}
	nodeCmd.PersistentFlags().String("priv-key", "", "Path to private key")
	nodeCmd.PersistentFlags().String("pub-key", "", "Path to pub key") // !Can be generated from private

	return nodeCmd
}
