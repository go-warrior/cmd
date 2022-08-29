package upgrade

import (
	"fmt"

	"github.com/go-warrior/cmd/internal/base"

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
		"github.com/go-kratos/kratos/cmd/kratos/v2@latest",
		"github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest",
		"github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}