package command

import (
	"fmt"
	"github.com/urfave/cli"
)

func Update(context *cli.Context) error {
	fmt.Fprintln(context.App.Writer, "Update")
	return nil
}
