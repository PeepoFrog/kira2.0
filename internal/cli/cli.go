package cli

import (
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/cli/version"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

const (
	use   = "kira2_launcher"
	short = "short descritpion"
	long  = "long description"
)

var log = logging.Log

func NewCLI(cmds []*cobra.Command) *cobra.Command {
	log.Debug("Creating new CLI...")
	rootCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}
	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbosity level. Default: `false` ")
	rootCmd.PersistentFlags().String("log-level", "", fmt.Sprintf("Messages with this level and above will be logged. Valid levels are: %s", strings.Join(logging.ValidLogLevels, ", ")))
	log.Debug(rootCmd.Flags())
	return rootCmd
}

func Start() {
	cmds := []*cobra.Command{version.Version()}
	c := NewCLI(cmds)
	if err := c.Execute(); err != nil {
		panic(err)
	}

	logLevel, _ := c.Flags().GetString("log-level")

	if logLevel != "" {
		logging.SetLevel(logLevel)
	}
}
