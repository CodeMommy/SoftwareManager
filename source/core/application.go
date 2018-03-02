package core

import (
    "os"
    "log"
    "time"
    "github.com/urfave/cli"
    "github.com/CodeMommy/SoftwareManager/source/config"
)

func Application() {
    application := cli.NewApp()
    application.Name = config.ApplicationName
    application.Version = config.ApplicationVersion
    application.Usage = config.ApplicationSummary
    application.Compiled = time.Now()
    application.Commands = config.Command()
    err := application.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
