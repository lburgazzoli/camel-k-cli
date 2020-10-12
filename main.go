package main

import (
	"fmt"
	"github.com/lburgazzoli/camel-k-cli/pkg/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"

	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	var kubeconf string
	var namespace string
	var dependencies cli.StringSlice
	var debug bool

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
			Aliases:     []string{"d"},
			Destination: &dependencies,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug",
			Destination: &debug,
		}),
	}

	app := &cli.App{
		Name:                 "kamel",
		Usage:                "kamel",
		EnableBashCompletion: true,
		Before:               cmd.InitInputSourceWithContext(flags, "kamelconfig", "kamel."),
		Flags:                flags,
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "run",
				Flags:  runFlags,
				Before: cmd.InitInputSourceWithContext(runFlags, "kamelconfig", "kamel.run."),
				Action: func(c *cli.Context) error {
					fmt.Println("dependencies:", dependencies)
					fmt.Println("debug:", debug)
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("kc:", kubeconf)
			fmt.Println("ns:", namespace)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
