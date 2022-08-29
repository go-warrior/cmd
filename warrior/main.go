package main

import (
	"github.com/go-warrior/cmd/warrior/v2/internal/change"
	"github.com/go-warrior/cmd/warrior/v2/internal/project"
	"github.com/go-warrior/cmd/warrior/v2/internal/proto"
	"github.com/go-warrior/cmd/warrior/v2/internal/run"
	"github.com/go-warrior/cmd/warrior/v2/internal/upgrade"
	"github.com/spf13/cobra"
	"log"
)

var (
	release = "v2.5.0"
	rootCmd = &cobra.Command{
	Use:     "Warrior",
	Short:   "Warrior: An elegant toolkit for Go microservices.",
	Long:    `Warrior: An elegant toolkit for Go microservices.`,
	Version: release,
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
