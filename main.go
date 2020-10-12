package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

// GitCommit --
var GitCommit string

func main() {
	var kubeconf string
	var namespace string
	var dependencies cli.StringSlice

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "kamelconfig",
			Usage: "The location of the kamel cli configuration",
			Value: ".kamel/config.yaml",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "kubeconf",
			Usage:       "The path to the kubernetes config file to use for CLI requests",
			Destination: &kubeconf,
			EnvVars:     []string{"KUBECONFIG"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "namespace",
			Aliases:     []string{"n"},
			Usage:       "The `namespace` to use for all operations",
			Destination: &namespace,
		}),
	}

	runFlags := []cli.Flag{
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:        "dependency",
			Usage:       "dependency",
			Destination: &dependencies,
		}),
	}

	app := &cli.App{
		Name:                 "kamel",
		Usage:                "kamel",
		EnableBashCompletion: true,
		Before:               altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("kamelconfig")),
		Flags:                flags,
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run",
				Flags: runFlags,
				Action: func(c *cli.Context) error {
					fmt.Println(dependencies)
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(kubeconf)
			fmt.Println(namespace)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
