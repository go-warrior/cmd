package main

import (
	"log"

	"github.com/go-warrior/cmd/warrior/internal/change"
	"github.com/go-warrior/cmd/warrior/internal/domain"
	"github.com/go-warrior/cmd/warrior/internal/model"
	"github.com/go-warrior/cmd/warrior/internal/project"
	"github.com/go-warrior/cmd/warrior/internal/proto"
	"github.com/go-warrior/cmd/warrior/internal/run"
	"github.com/go-warrior/cmd/warrior/internal/upgrade"
	"github.com/go-warrior/cmd/warrior/internal/wire"

	"github.com/spf13/cobra"
)

var (
	version string = "v0.0.15"

	rootCmd = &cobra.Command{
		Use:     "warrior",
		Short:   "warrior: An elegant toolkit for Go microservices.",
		Long:    `warrior: An elegant toolkit for Go microservices.`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(wire.CmdWire)
	rootCmd.AddCommand(domain.CmdDomain)
	rootCmd.AddCommand(model.CmdModel)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
