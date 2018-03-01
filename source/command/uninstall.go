package command

import (
	"fmt"
	"github.com/urfave/cli"
)

func Uninstall(context *cli.Context) error {
	fmt.Fprintln(context.App.Writer, "Uninstall")
	return nil
}
