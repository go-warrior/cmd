package upgrade

import (
	"fmt"

	"github.com/go-warrior/cmd/warrior/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the warrior tools",
	Long:  "Upgrade the warrior tools. Example: warrior upgrade",
	Run:   Run,
}

// Run upgrade the kratos tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall(
		"github.com/go-warrior/cmd/warrior@latest",
		"github.com/go-warrior/cmd/protoc-gen-go-warrior@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"github.com/envoyproxy/protoc-gen-validate@latest",
		"github.com/google/wire/cmd/wire",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
