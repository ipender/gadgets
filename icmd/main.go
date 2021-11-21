package icmd

import (
	"os"

	"code.byted.org/gopkg/logs"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:          "gadgets",
		Short:        "gadgets",
		Long:         "gadgets",
		SilenceUsage: true,
	}

	if err := rootCmd.Execute(); err != nil {
		logs.Errorf("Execute error: %s, pid: %d, ppid: %d, ppname: %s", err.Error())
		os.Exit(-1)
	}
}
