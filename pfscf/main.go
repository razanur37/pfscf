package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/razanur37/pfscf/pfscf/cfg"
	"github.com/razanur37/pfscf/pfscf/cmd"
)

var (
	version = "0.16.51"
)

func main() {

	RootCmd := &cobra.Command{
		Use:   "pfscf",
		Short: "The Pathfinder Society Chronicle Filler (v" + version + ")",
	}

	RootCmd.PersistentFlags().BoolVarP(&cfg.Global.Verbose, "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(cmd.GetFillCommand())
	RootCmd.AddCommand(cmd.GetTemplateCommand())
	RootCmd.AddCommand(cmd.GetBatchCommand())
	RootCmd.AddCommand(cmd.GetOpenCommand())

	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
