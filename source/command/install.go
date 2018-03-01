package command

import (
	"fmt"
	"github.com/urfave/cli"
)

func Install(context *cli.Context) error {
	fmt.Fprintln(context.App.Writer, "Install")
	return nil
}
