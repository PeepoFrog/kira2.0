package version

import (
	"fmt"
	"os"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
	"github.com/spf13/cobra"
)

const (
	use   = "version"
	short = "short description for version command"
	long  = "long description for version command"
)

var (
	KiraLauncherVersion string = types.KiraVersion
	log                        = logging.Log
)

func Version() *cobra.Command {
	log.Debugln("Adding `version` command...")
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			log.Infoln("Version: ", KiraLauncherVersion)
			fmt.Fprintf(os.Stdout, "%v:%v\n", os.Args[0], KiraLauncherVersion)
		},
	}
	return versionCmd
}
